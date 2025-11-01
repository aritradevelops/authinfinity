package migrate

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const migrationsDir = "migrations"

// runAtlasCommand executes an Atlas CLI command
func runAtlasCommand(args ...string) error {
	godotenv.Load()
	cmd := exec.Command("atlas", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// newMigrateDiffCommand creates migration diff files using Atlas
func newMigrateDiffCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate:diff <name>",
		Short: "Generate migration diff from GORM models using Atlas",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			migrationName := strings.ToLower(strings.ReplaceAll(args[0], " ", "_"))

			// Create migrations directory if it doesn't exist
			if err := os.MkdirAll(migrationsDir, 0o755); err != nil {
				return fmt.Errorf("failed to create migrations directory: %w", err)
			}

			fmt.Printf("Generating migration: %s\n", migrationName)

			// Run Atlas migrate diff command
			return runAtlasCommand(
				"migrate",
				"diff",
				migrationName,
				"--config", "file://atlas.hcl",
				"--env", "gorm",
			)
		},
	}
	return cmd
}

// newMigrateApplyCommand applies pending migrations
func newMigrateApplyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate:apply",
		Short: "Apply pending migrations to the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, _ := cmd.Flags().GetString("env")
			fmt.Printf("Applying migrations (env: %s)...\n", env)
			return runAtlasCommand(
				"migrate",
				"apply",
				"--config", "file://atlas.hcl",
				"--env", env,
			)
		},
	}
	cmd.Flags().String("env", "dev", "Environment to use (local, dev, prod)")
	return cmd
}

// newMigrateStatusCommand shows migration status
func newMigrateStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate:status",
		Short: "Show current migration status",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, _ := cmd.Flags().GetString("env")
			return runAtlasCommand(
				"migrate",
				"status",
				"--config", "file://atlas.hcl",
				"--env", env,
			)
		},
	}
	cmd.Flags().String("env", "local", "Environment to use (local, dev, prod)")
	return cmd
}

// newMigrateValidateCommand validates migration files
func newMigrateValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate:validate",
		Short: "Validate migration files integrity",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, _ := cmd.Flags().GetString("env")
			return runAtlasCommand(
				"migrate",
				"validate",
				"--config", "file://atlas.hcl",
				"--env", env,
			)
		},
	}
	cmd.Flags().String("env", "local", "Environment to use (local, dev, prod)")
	return cmd
}

// newSchemaInspectCommand inspects the current database schema
func newSchemaInspectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema:inspect",
		Short: "Inspect current database schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, _ := cmd.Flags().GetString("env")
			return runAtlasCommand(
				"schema",
				"inspect",
				"--config", "file://atlas.hcl",
				"--env", env,
			)
		},
	}
	cmd.Flags().String("env", "local", "Environment to use (local, dev, prod)")
	return cmd
}

// newSchemaApplyCommand applies schema directly (for dev only)
func newSchemaApplyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema:apply",
		Short: "Apply schema changes directly (dev only - bypasses migrations)",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, _ := cmd.Flags().GetString("env")
			if env == "prod" {
				return fmt.Errorf("schema:apply is not allowed in production - use migrations instead")
			}
			autoApprove, _ := cmd.Flags().GetBool("auto-approve")
			args = []string{"schema", "apply", "--config", "file://atlas.hcl", "--env", env}
			if autoApprove {
				args = append(args, "--auto-approve")
			}
			return runAtlasCommand(args...)
		},
	}
	cmd.Flags().String("env", "gorm", "Environment to use (gorm, local)")
	cmd.Flags().Bool("auto-approve", false, "Auto approve changes without prompt")
	return cmd
}
func RegisterCommand(parent *cobra.Command) {
	// Register all migration commands
	parent.AddCommand(newMigrateDiffCommand())
	parent.AddCommand(newMigrateApplyCommand())
	parent.AddCommand(newMigrateStatusCommand())
	parent.AddCommand(newMigrateValidateCommand())
	parent.AddCommand(newSchemaInspectCommand())
	parent.AddCommand(newSchemaApplyCommand())
}
