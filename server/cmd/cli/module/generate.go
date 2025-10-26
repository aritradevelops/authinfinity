package module

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/go/ast/astutil"

	pluralizer "github.com/gertd/go-pluralize"

	"github.com/aritradevelops/authinfinity/server/cmd/cli/internal/cliutil"
	"github.com/aritradevelops/authinfinity/server/cmd/cli/internal/logx"
)

//go:embed templates/*
var templatesFS embed.FS
var pluralize = pluralizer.NewClient()
var lg = logx.New()

type moduleInput struct {
	Raw        string
	Package    string
	Entity     string
	Collection string
	RouteGroup string
	ModulePath string
}

func buildNames(raw string) (moduleInput, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return moduleInput{}, errors.New("empty module name")
	}
	// derive module path from go.mod
	repoRoot, err := cliutil.FindRepoRoot(".")
	if err != nil {
		return moduleInput{}, err
	}
	mp, err := cliutil.ModulePath(repoRoot)
	if err != nil {
		return moduleInput{}, err
	}
	pkg := strings.ToLower(trimmed)
	entity := cases.Title(language.English).String(strings.ToLower(trimmed))
	collection := pluralize.Plural(pkg)
	return moduleInput{
		Raw:        trimmed,
		Package:    pkg,
		Entity:     entity,
		Collection: collection,
		RouteGroup: collection,
		ModulePath: mp,
	}, nil
}

// interactive TUI removed; module names are taken only from CLI args

func renderModule(mi moduleInput) error {
	baseDir := filepath.Join("internal", "app", "modules", mi.Package)
	lg.Info("creating module directory", slog.String("dir", baseDir))
	// Check if directory already exists
	if _, err := os.Stat(baseDir); err == nil {
		lg.Warn("module already exists", slog.String("dir", baseDir))
		return fmt.Errorf("module directory already exists: %s", baseDir)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check module directory: %w", err)
	}
	if err := os.MkdirAll(baseDir, 0o755); err != nil {
		return err
	}
	data := map[string]string{
		"Package":    mi.Package,
		"Entity":     mi.Entity,
		"Collection": mi.Collection,
		"RouteGroup": mi.RouteGroup,
		"ModulePath": mi.ModulePath,
	}

	files := map[string]string{
		"model.tmpl":      filepath.Join(baseDir, mi.Package+"_model.go"),
		"repository.tmpl": filepath.Join(baseDir, mi.Package+"_repository.go"),
		"service.tmpl":    filepath.Join(baseDir, mi.Package+"_service.go"),
		"controller.tmpl": filepath.Join(baseDir, mi.Package+"_controller.go"),
		"router.tmpl":     filepath.Join(baseDir, mi.Package+"_router.go"),
	}

	for tmplName, outPath := range files {
		b, err := templatesFS.ReadFile("templates/" + tmplName)
		if err != nil {
			return fmt.Errorf("read template %s: %w", tmplName, err)
		}
		t, err := template.New(tmplName).Parse(string(b))
		if err != nil {
			return fmt.Errorf("parse template %s: %w", tmplName, err)
		}
		f, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("create %s: %w", outPath, err)
		}
		if err := t.Execute(f, data); err != nil {
			_ = f.Close()
			return fmt.Errorf("render %s: %w", outPath, err)
		}
		if err := f.Close(); err != nil {
			return err
		}
		lg.Info("wrote file", slog.String("file", outPath), slog.String("template", tmplName))
	}

	serverPath := filepath.Join("internal", "pkg", "server", "server.go")
	serverCode, err := os.ReadFile(serverPath)
	if err != nil {
		lg.Error("read server.go", slog.String("path", serverPath), slog.Any("err", err))
		return err
	}
	lg.Info("parsed server for AST update", slog.String("path", serverPath))
	serverCodeString := string(serverCode)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, serverPath, serverCodeString, parser.ParseComments)
	if err != nil {
		lg.Error("parse server.go", slog.Any("err", err))
		return err
	}

	// ✅ Add the import if not already present
	importPath := fmt.Sprintf("%s/internal/app/modules/%s", mi.ModulePath, mi.Package)
	if astutil.AddImport(fset, f, importPath) {
		lg.Info("added import", slog.String("import", importPath))
	} else {
		lg.Info("import already present", slog.String("import", importPath))
	}

	// ✅ Find setupRoutes function
	var targetFunc *ast.FuncDecl
	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "setupRoutes" {
			targetFunc = fn
			return false
		}
		return true
	})

	if targetFunc == nil {
		return errors.New("function 'setupRoutes' not found in server.go")
	}

	// ✅ Create statement: user.RegisterRoutes(apiV1)
	newStatement := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: mi.Package},
				Sel: &ast.Ident{Name: "RegisterRoutes"},
			},
			Args: []ast.Expr{
				&ast.Ident{Name: "apiV1"},
			},
		},
	}

	// Add the new statement at the end of the function body
	targetFunc.Body.List = append(targetFunc.Body.List, newStatement)
	lg.Info("registered routes in server", slog.String("module", mi.Package))

	// ✅ Write the modified AST back to file
	file, err := os.Create(serverPath)
	if err != nil {
		lg.Error("open server.go for write", slog.Any("err", err))
		return err
	}
	defer file.Close()

	if err := printer.Fprint(file, fset, f); err != nil {
		lg.Error("write server.go", slog.Any("err", err))
		return err
	}

	lg.Info("updated server.go successfully")
	return nil
}

func newGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate <module_names...>",
		Short: "Generate modules with CRUD boilerplate",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			lg.Info("starting module generation", slog.Int("count", len(args)))
			for _, raw := range args {
				mi, err := buildNames(raw)
				if err != nil {
					return err
				}
				lg.Info("resolved module", slog.String("name", mi.Package), slog.String("modulePath", mi.ModulePath))
				if err := renderModule(mi); err != nil {
					return err
				}
				lg.Info("module generated", slog.String("entity", mi.Entity), slog.String("path", filepath.Join("internal", "app", "modules", mi.Package)))
			}
			lg.Info("module generation finished")
			return nil
		},
	}
	return cmd
}
