package module

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/ast/astutil"
)

func deleteModule(mi moduleInput) error {
	baseDir := filepath.Join("internal", "app", "modules", mi.Package)
	lg.Info("removing module directory", slog.String("dir", baseDir))
	if err := os.RemoveAll(baseDir); err != nil {
		return err
	}

	serverPath := filepath.Join("internal", "pkg", "server", "server.go")
	lg.Info("updating server AST to remove routes", slog.String("path", serverPath))
	fset := token.NewFileSet()

	src, err := os.ReadFile(serverPath)
	if err != nil {
		return fmt.Errorf("read server.go: %w", err)
	}

	f, err := parser.ParseFile(fset, serverPath, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse server.go: %w", err)
	}

	// ✅ 1. Remove import for that module
	importPath := fmt.Sprintf("%s/internal/app/modules/%s", mi.ModulePath, mi.Package)
	if astutil.DeleteImport(fset, f, importPath) {
		lg.Info("removed import", slog.String("import", importPath))
	}

	// ✅ 2. Remove <mi.Package>.RegisterRoutes(apiV1) call inside setupRoutes()
	ast.Inspect(f, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "setupRoutes" {
			return true
		}

		newBody := make([]ast.Stmt, 0, len(fn.Body.List))
		for _, stmt := range fn.Body.List {
			// Check if it's a call expression like `pkg.RegisterRoutes(apiV1)`
			call, ok := stmt.(*ast.ExprStmt)
			if !ok {
				newBody = append(newBody, stmt)
				continue
			}

			callExpr, ok := call.X.(*ast.CallExpr)
			if !ok {
				newBody = append(newBody, stmt)
				continue
			}

			sel, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				newBody = append(newBody, stmt)
				continue
			}

			pkgIdent, ok := sel.X.(*ast.Ident)
			if !ok {
				newBody = append(newBody, stmt)
				continue
			}

			// Match <mi.Package>.RegisterRoutes(...)
			if pkgIdent.Name == mi.Package && sel.Sel.Name == "RegisterRoutes" {
				lg.Info("removed RegisterRoutes call", slog.String("module", mi.Package))
				continue // skip adding it to the new body
			}

			newBody = append(newBody, stmt)
		}
		fn.Body.List = newBody
		return false
	})

	// ✅ Write modified AST back to file
	file, err := os.Create(serverPath)
	if err != nil {
		return fmt.Errorf("open for write: %w", err)
	}
	defer file.Close()

	if err := printer.Fprint(file, fset, f); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	lg.Info("updated server.go after removing module")
	return nil
}

func newRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove:modules <module_names...>",
		Short: "Remove modules",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			lg.Info("starting module removal", slog.Int("count", len(args)))
			for _, raw := range args {
				mi, err := buildNames(raw)
				if err != nil {
					return err
				}
				if err := deleteModule(mi); err != nil {
					return err
				}
				// Refresh migrate models/imports after removal
				if err := refreshMigrateModels(mi); err != nil {
					return err
				}
				lg.Info("module removed", slog.String("entity", mi.Entity), slog.String("module", mi.Package))
			}
			lg.Info("module removal finished")
			return nil
		},
	}
}
