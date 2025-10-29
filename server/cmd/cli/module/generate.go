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

	"github.com/ettle/strcase"
	"github.com/spf13/cobra"
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
	File       string
	Var        string
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
	entity := trimmed
	fileName := strcase.ToSnake(trimmed)
	collection := pluralize.Plural(fileName)
	routeGroup := strcase.ToKebab(collection)
	varName := strcase.ToCamel(entity)
	return moduleInput{
		Raw:        trimmed,
		Package:    pkg,
		Entity:     entity,
		Collection: collection,
		RouteGroup: routeGroup,
		File:       fileName,
		Var:        varName,
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
		"Var":        mi.Var,
		"File":       mi.File,
	}

	files := map[string]string{
		"model.tmpl":      filepath.Join(baseDir, mi.File+"_model.go"),
		"repository.tmpl": filepath.Join(baseDir, mi.File+"_repository.go"),
		"service.tmpl":    filepath.Join(baseDir, mi.File+"_service.go"),
		"controller.tmpl": filepath.Join(baseDir, mi.File+"_controller.go"),
		"router.tmpl":     filepath.Join(baseDir, mi.File+"_router.go"),
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
	// Refresh migrate models/imports dynamically
	if err := refreshAtlasProvider(mi); err != nil {
		lg.Error("refresh migrate models failed", slog.Any("err", err))
		return err
	}
	return nil
}

// refreshAtlasProvider scans internal/app/modules and rewrites cmd/cli/migrate/migrate.go
// to import each module and populate the models slice accordingly.
func refreshAtlasProvider(mi moduleInput) error {
	repoRoot, err := cliutil.FindRepoRoot(".")
	if err != nil {
		return err
	}
	modulesDir := filepath.Join(repoRoot, "internal", "app", "modules")
	lg.Info("refreshing migrate models", slog.String("modulesDir", modulesDir))
	dirs, err := os.ReadDir(modulesDir)
	if err != nil {
		return err
	}
	lg.Info("scanned modules directory", slog.Int("entries", len(dirs)))
	// collect found modules
	type mod struct{ pkg, entity, importPath string }
	var found []mod
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		name := d.Name()
		files, err := os.ReadDir(filepath.Join(modulesDir, name))
		if err != nil {
			return err
		}

		modelFileName := ""
		model := ""
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if strings.HasSuffix(f.Name(), "_model.go") {
				modelFileName = f.Name()
				model = strings.TrimSuffix(f.Name(), "_model.go")
			}
		}
		if model == "" {
			continue
		}

		modelPath := filepath.Join(modulesDir, name, modelFileName)
		if _, err := os.Stat(modelPath); err == nil {
			found = append(found, mod{pkg: name, entity: strcase.ToPascal(model), importPath: fmt.Sprintf("%s/internal/app/modules/%s", mi.ModulePath, name)})
			lg.Info("found model", slog.String("pkg", name), slog.String("entity", strcase.ToPascal(model)))
		}
	}
	// open atlas provider
	migratePath := filepath.Join(repoRoot, "cmd", "atlas-provider", "main.go")
	fset := token.NewFileSet()
	srcBytes, err := os.ReadFile(migratePath)
	if err != nil {
		return err
	}
	f, err := parser.ParseFile(fset, migratePath, string(srcBytes), parser.ParseComments)
	if err != nil {
		return err
	}
	lg.Info("parsed migrate file", slog.String("path", migratePath))
	// remove existing module imports
	removed := 0
	for _, imp := range f.Imports {
		p := strings.Trim(imp.Path.Value, "\"")
		if strings.Contains(p, "/internal/app/modules/") {
			if astutil.DeleteImport(fset, f, p) {
				removed++
			}
		}
	}
	lg.Info("removed module imports", slog.Int("count", removed))
	// add new imports
	added := 0
	for _, m := range found {
		if astutil.AddImport(fset, f, m.importPath) {
			added++
		}
	}
	lg.Info("added module imports", slog.Int("count", added))
	// rewrite models slice
	rewrote := false
	ast.Inspect(f, func(n ast.Node) bool {
		as, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}
		if len(as.Specs) != 1 {
			return true
		}
		valSpec, ok := as.Specs[0].(*ast.ValueSpec)
		if !ok {
			return true
		}

		if len(valSpec.Names) != 1 || len(valSpec.Values) != 1 {
			return true
		}

		if valSpec.Names[0].Name != "models" {
			return true
		}

		// build elements: &pkg.Entity{}
		var elts []ast.Expr
		for _, m := range found {
			elts = append(elts, &ast.UnaryExpr{Op: token.AND, X: &ast.CompositeLit{Type: &ast.SelectorExpr{X: &ast.Ident{Name: m.pkg}, Sel: &ast.Ident{Name: m.entity}}}})
		}
		eltsOrginal, ok := valSpec.Values[0].(*ast.CompositeLit)
		if !ok {
			return true
		}
		eltsOrginal.Elts = elts
		rewrote = true
		return false
	})
	if !rewrote {
		return errors.New("did not find models slice to rewrite in generate.go")
	}
	lg.Info("rewrote models slice", slog.Int("models", len(found)))
	// write back
	out, err := os.Create(migratePath)
	if err != nil {
		return err
	}
	defer out.Close()
	if err := printer.Fprint(out, fset, f); err != nil {
		return err
	}
	lg.Info("updated migrate file", slog.String("path", migratePath))
	return nil
}

func newGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate:modules <module_names...>",
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
