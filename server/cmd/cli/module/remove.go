package module

import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/ast/astutil"
)

func removeModule(mi moduleInput) error {
	baseDir := filepath.Join("internal", "app", "modules", mi.Package)
	logger.Info().Str("dir", baseDir).Msg("removing module directory")

	if err := os.RemoveAll(baseDir); err != nil {
		return fmt.Errorf("remove module directory: %w", err)
	}

	serverPath := filepath.Join("internal", "pkg", "server", "server.go")
	logger.Info().Str("path", serverPath).Msg("updating server AST to remove routes")

	fset := token.NewFileSet()
	src, err := os.ReadFile(serverPath)
	if err != nil {
		return fmt.Errorf("read server.go: %w", err)
	}

	f, err := parser.ParseFile(fset, serverPath, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse server.go: %w", err)
	}

	importPath := fmt.Sprintf("%s/internal/app/modules/%s", mi.ModulePath, mi.Package)
	if astutil.DeleteImport(fset, f, importPath) {
		logger.Info().Str("import", importPath).Msg("removed import")
	}

	removeRegisterRoutesCall(f, mi.Package)

	file, err := os.Create(serverPath)
	if err != nil {
		return fmt.Errorf("open for write: %w", err)
	}
	defer file.Close()

	if err := printer.Fprint(file, fset, f); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	logger.Info().Msg("updated server.go after removing module")
	return nil
}

func newRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove:modules <module_names...>",
		Short: "Remove one or more modules cleanly",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info().Int("count", len(args)).Msg("starting module removal")

			for _, raw := range args {
				mi, err := buildNames(raw)
				if err != nil {
					return err
				}

				if err := removeModule(mi); err != nil {
					return err
				}

				if err := refreshAtlasProvider(mi); err != nil {
					return err
				}

				logger.Info().
					Str("entity", mi.Entity).
					Str("module", mi.Package).
					Msg("module removed")
			}

			logger.Info().Msg("module removal finished")
			return nil
		},
	}
}
