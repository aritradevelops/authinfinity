package module

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/ast/astutil"
)

func deleteModule(mi moduleInput) error {
	baseDir := filepath.Join("internal", "modules", mi.Package)
	if err := os.RemoveAll(baseDir); err != nil {
		return err
	}

	serverPath := filepath.Join("internal", "server", "server.go")
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
	importPath := fmt.Sprintf("github.com/aritradevelops/crudpaglu/internal/modules/%s", mi.Package)
	if astutil.DeleteImport(fset, f, importPath) {
		fmt.Printf("Removed import: %s\n", importPath)
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
				fmt.Printf("Removed call: %s.RegisterRoutes(...)\n", mi.Package)
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

	fmt.Println("Successfully updated server.go after removing module.")
	return nil
}

func newRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove <module_names...>",
		Short: "Remove modules",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, raw := range args {
				mi := buildNames(raw)
				if err := deleteModule(mi); err != nil {
					return err
				}
				fmt.Printf("Removed module %s at internal/modules/%s\n", mi.Entity, mi.Package)
			}
			return nil
		},
	}
}
