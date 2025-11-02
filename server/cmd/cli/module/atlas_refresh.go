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
	"strings"

	"github.com/ettle/strcase"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/aritradevelops/authinfinity/server/cmd/cli/internal/cliutil"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

type mod struct{ pkg, entity, importPath string }

// refreshAtlasProvider updates cmd/atlas-provider/main.go to include all module models.
func refreshAtlasProvider(mi moduleInput) error {
	repoRoot, err := cliutil.FindRepoRoot(".")
	if err != nil {
		return err
	}

	modulesDir := filepath.Join(repoRoot, "internal", "app", "modules")
	logger.Info().Str("path", modulesDir).Msg("refreshing migrate models")

	dirs, err := os.ReadDir(modulesDir)
	if err != nil {
		return err
	}

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

		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), "_model.go") {
				entity := strings.TrimSuffix(f.Name(), "_model.go")
				found = append(found, mod{
					pkg:        name,
					entity:     strcase.ToPascal(entity),
					importPath: fmt.Sprintf("%s/internal/app/modules/%s", mi.ModulePath, name),
				})
				logger.Info().Str("pkg", name).Str("entity", strcase.ToPascal(entity)).Msg("found module model")
				break
			}
		}
	}

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

	// Clean old imports
	for _, imp := range f.Imports {
		p := strings.Trim(imp.Path.Value, "\"")
		if strings.Contains(p, "/internal/app/modules/") {
			astutil.DeleteImport(fset, f, p)
		}
	}

	// Add new imports
	for _, m := range found {
		astutil.AddImport(fset, f, m.importPath)
	}

	// Update models slice
	updated := updateModelsSlice(f, found)
	if !updated {
		return errors.New("did not find models slice to rewrite in migrate file")
	}

	out, err := os.Create(migratePath)
	if err != nil {
		return err
	}
	defer out.Close()

	if err := printer.Fprint(out, fset, f); err != nil {
		return err
	}
	logger.Info().Str("path", migratePath).Msg("updated migrate file")
	return nil
}

func updateModelsSlice(f *ast.File, found []mod) bool {
	rewrote := false
	ast.Inspect(f, func(n ast.Node) bool {
		as, ok := n.(*ast.GenDecl)
		if !ok || len(as.Specs) != 1 {
			return true
		}
		valSpec, ok := as.Specs[0].(*ast.ValueSpec)
		if !ok || valSpec.Names[0].Name != "models" {
			return true
		}

		var elts []ast.Expr
		for _, m := range found {
			elts = append(elts, &ast.UnaryExpr{
				Op: token.AND,
				X: &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: m.pkg},
						Sel: &ast.Ident{Name: m.entity},
					},
				},
			})
		}

		if orig, ok := valSpec.Values[0].(*ast.CompositeLit); ok {
			orig.Elts = elts
			rewrote = true
		}
		return false
	})
	return rewrote
}
