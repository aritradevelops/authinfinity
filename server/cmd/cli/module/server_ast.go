package module

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

// updateServerAST modifies internal/pkg/server/server.go to register the new module routes.
func updateServerAST(mi moduleInput) error {
	serverPath := filepath.Join("internal", "pkg", "server", "server.go")
	serverCode, err := os.ReadFile(serverPath)
	if err != nil {
		logger.Error().Err(err).Str("path", serverPath).Msg("failed to read server.go")
		return err
	}
	logger.Info().Str("path", serverPath).Msg("parsed server for AST update")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, serverPath, string(serverCode), parser.ParseComments)
	if err != nil {
		logger.Error().Err(err).Str("path", serverPath).Msg("failed to parse server.go")
		return err
	}

	importPath := fmt.Sprintf("%s/internal/app/modules/%s", mi.ModulePath, mi.Package)
	if astutil.AddImport(fset, f, importPath) {
		logger.Info().Str("import", importPath).Msg("added module import")
	} else {
		logger.Debug().Str("import", importPath).Msg("module import already present")
	}

	var targetFunc *ast.FuncDecl
	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "setupRoutes" {
			targetFunc = fn
			return false
		}
		return true
	})

	if targetFunc == nil {
		logger.Error().Msg("function 'setupRoutes' not found in server.go")
		return errors.New("function 'setupRoutes' not found in server.go")
	}

	newStatement := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: mi.Package},
				Sel: &ast.Ident{Name: "RegisterRoutes"},
			},
			Args: []ast.Expr{&ast.Ident{Name: "apiV1"}},
		},
	}

	targetFunc.Body.List = append(targetFunc.Body.List, newStatement)
	logger.Info().Str("module", mi.Package).Msg("registered routes in server.go")

	file, err := os.Create(serverPath)
	if err != nil {
		logger.Error().Err(err).Str("path", serverPath).Msg("failed to open server.go for writing")
		return err
	}
	defer file.Close()

	if err := printer.Fprint(file, fset, f); err != nil {
		logger.Error().Err(err).Str("path", serverPath).Msg("failed to write updated server.go")
		return err
	}

	logger.Info().Str("path", serverPath).Msg("updated server.go successfully")
	return nil
}

// removeRegisterRoutesCall finds and removes <pkg>.RegisterRoutes(apiV1) from setupRoutes()
func removeRegisterRoutesCall(f *ast.File, pkgName string) {
	ast.Inspect(f, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "setupRoutes" {
			return true
		}

		newBody := make([]ast.Stmt, 0, len(fn.Body.List))
		for _, stmt := range fn.Body.List {
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

			if pkgIdent.Name == pkgName && sel.Sel.Name == "RegisterRoutes" {
				logger.Info().Str("module", pkgName).Msg("removed RegisterRoutes call")
				continue
			}

			newBody = append(newBody, stmt)
		}
		fn.Body.List = newBody
		return false
	})
}
