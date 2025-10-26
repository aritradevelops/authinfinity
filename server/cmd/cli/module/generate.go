package module

import (
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/go/ast/astutil"

	pluralizer "github.com/gertd/go-pluralize"
)

//go:embed templates/*
var templatesFS embed.FS
var pluralize = pluralizer.NewClient()

type moduleInput struct {
	Raw        string
	Package    string
	Entity     string
	Collection string
	RouteGroup string
}

func buildNames(raw string) moduleInput {
	trimmed := strings.TrimSpace(raw)
	pkg := strings.ToLower(trimmed)
	entity := cases.Title(language.English).String(strings.ToLower(trimmed))
	collection := pluralize.Plural(pkg)
	return moduleInput{
		Raw:        trimmed,
		Package:    pkg,
		Entity:     entity,
		Collection: collection,
		RouteGroup: collection,
	}
}

// interactive TUI removed; module names are taken only from CLI args

func renderModule(mi moduleInput) error {
	baseDir := filepath.Join("internal", "modules", mi.Package)
	// Check if directory already exists
	if _, err := os.Stat(baseDir); err == nil {
		// Directory exists — terminate early
		return fmt.Errorf("module directory already exists: %s", baseDir)
	} else if !os.IsNotExist(err) {
		// Some other error (e.g., permission issue)
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
	}

	serverPath := filepath.Join("internal", "server", "server.go")
	serverCode, err := os.ReadFile(serverPath)
	if err != nil {
		log.Fatal(err)
	}
	serverCodeString := string(serverCode)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, serverPath, serverCodeString, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// ✅ Add the import if not already present
	if astutil.AddImport(fset, f, fmt.Sprintf("github.com/aritradevelops/crudpaglu/internal/modules/%s", mi.Package)) {
		log.Println(fmt.Sprintf("Added import for %s package", mi.Package))
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
		log.Fatal("Function 'setupRoutes' not found")
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

	// ✅ Write the modified AST back to file
	file, err := os.Create(serverPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := printer.Fprint(file, fset, f); err != nil {
		log.Fatal(err)
	}

	log.Println("Updated server.go successfully!")
	return nil
}

func newGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate <module_names...>",
		Short: "Generate modules with CRUD boilerplate",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			modules := args
			for _, raw := range modules {
				mi := buildNames(raw)
				if err := renderModule(mi); err != nil {
					return err
				}
				fmt.Printf("Generated module %s at internal/modules/%s\n", mi.Entity, mi.Package)
			}
			return nil
		},
	}
	return cmd
}
