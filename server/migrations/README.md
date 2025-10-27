# Atlas Migrations

This directory contains database migrations managed by [Atlas](https://atlasgo.io/).

## Prerequisites

Install Atlas CLI:

```bash
# macOS
brew install ariga/tap/atlas

# Linux/WSL
curl -sSf https://atlasgo.sh | sh

# Windows (via Scoop)
scoop install atlas
```

## Quick Start

### 1. Generate Migration from Model Changes

After modifying GORM models:

```bash
go run cmd/cli/main.go migrate:diff "add_user_email_index"
```

This creates a new migration file in this directory (e.g., `20240127120000_add_user_email_index.sql`).

### 2. Apply Migrations

```bash
go run cmd/cli/main.go migrate:apply
```

### 3. Check Status

```bash
go run cmd/cli/main.go migrate:status
```

## Common Workflows

### Development Workflow

```bash
# 1. Modify GORM models in internal/app/modules/
# 2. Generate migration
go run cmd/cli/main.go migrate:diff "your_change_description"

# 3. Review generated SQL in migrations/
# 4. Apply to local database
go run cmd/cli/main.go migrate:apply
```

### Quick Iteration (Dev Only)

For rapid prototyping, you can skip migrations and apply schema directly:

```bash
go run cmd/cli/main.go schema:apply --auto-approve
```

⚠️ **Warning**: This bypasses version control. Use only in development.

## Environment Configuration

Migrations require two environment variables:

```bash
# Main database connection
export DB_CONNECTION_URI="postgresql://postgres:admin@localhost:5432/authinfinity?sslmode=disable"

# Dev database for schema diffing (optional - defaults to DB_CONNECTION_URI)
export DEV_DB_URL="postgresql://postgres:admin@localhost:5432/authinfinity?sslmode=disable"
```

**Note**: For local development, both can point to the same database. Atlas uses the dev database as a temporary workspace to compute schema diffs.

Atlas supports multiple environments (configured in `atlas.hcl`):
- `local` - Default for development
- `dev` - Shared development database
- `prod` - Production (with safety checks)

Use `--env` flag to specify:

```bash
go run cmd/cli/main.go migrate:apply --env=dev
```

## Migration File Format

Migrations follow Atlas's versioned format:

```
20240127120000_add_user_email_index.sql
└─┬──┘ └────┬────┘
  │         └─ Description
  └─ Timestamp version
```

## Adding New Models

When creating a new module:

1. Update `cmd/cli/atlas-provider/main.go`:
   ```go
   var models = []any{
       // existing models...
       &yournewmodule.YourModel{},
   }
   ```

2. Generate migration:
   ```bash
   go run cmd/cli/main.go migrate:diff "add_yourmodule_table"
   ```

## Troubleshooting

### Migration Conflicts

If you get checksum errors:

```bash
# Validate migrations
go run cmd/cli/main.go migrate:validate

# Check what's applied
go run cmd/cli/main.go migrate:status
```

### Inspect Database

Compare database with models:

```bash
go run cmd/cli/main.go schema:inspect
```

## CI/CD Integration

In your CI pipeline:

```bash
# Validate migrations
go run cmd/cli/main.go migrate:validate

# Check status
go run cmd/cli/main.go migrate:status --env=prod

# Apply (with approval)
go run cmd/cli/main.go migrate:apply --env=prod
```

## Learn More

- [Atlas Documentation](https://atlasgo.io/docs)
- [GORM Atlas Provider](https://github.com/ariga/atlas-provider-gorm)
- [Project WARP.md](../WARP.md) - Full project documentation
