package module

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ettle/strcase"
	pluralizer "github.com/gertd/go-pluralize"
	"github.com/spf13/cobra"

	"github.com/aritradevelops/authinfinity/server/cmd/cli/internal/cliutil"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

var pluralize = pluralizer.NewClient()

//go:embed templates/*
var templatesFS embed.FS

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

// buildNames derives consistent module naming conventions.
func buildNames(raw string) (moduleInput, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return moduleInput{}, errors.New("empty module name")
	}

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

// renderModule renders the module from templates and updates AST references.
func renderModule(mi moduleInput) error {
	baseDir := filepath.Join("internal", "app", "modules", mi.Package)
	logger.Info().Str("dir", baseDir).Msg("creating module directory")

	if _, err := os.Stat(baseDir); err == nil {
		logger.Warn().Str("dir", baseDir).Msg("module already exists, skipping creation")
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
		if err := renderTemplate(tmplName, outPath, data); err != nil {
			return err
		}
		logger.Info().Str("file", outPath).Str("template", tmplName).Msg("generated file from template")
	}

	if err := updateServerAST(mi); err != nil {
		return err
	}

	if err := refreshAtlasProvider(mi); err != nil {
		logger.Error().Err(err).Msg("failed to refresh migrate models")
		return err
	}

	return nil
}

// renderTemplate is a safe helper for rendering embedded templates.
func renderTemplate(tmplName, outPath string, data any) error {
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
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("render %s: %w", outPath, err)
	}
	return nil
}

// newGenerateCmd defines the CLI entrypoint for module generation.
func newGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate:modules <module_names...>",
		Short: "Generate modules with CRUD boilerplate",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info().Int("count", len(args)).Msg("starting module generation")

			for _, raw := range args {
				mi, err := buildNames(raw)
				if err != nil {
					return err
				}
				logger.Info().Str("name", mi.Package).Str("path", mi.ModulePath).Msg("resolved module input")

				if err := renderModule(mi); err != nil {
					return err
				}

				logger.Info().
					Str("entity", mi.Entity).
					Str("dir", filepath.Join("internal", "app", "modules", mi.Package)).
					Msg("module generated")
			}

			logger.Info().Msg("module generation completed successfully")
			return nil
		},
	}
	return cmd
}
