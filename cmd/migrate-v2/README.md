# Terraform Cloudflare Provider Migration Tool (v2)

A powerful tool for automatically migrating Terraform configurations between different versions of the Cloudflare provider.

## Quick Start

```bash
# Build the tool
make build

# Migrate a single file
./bin/migrate-v2 migrate -f terraform.tf

# Migrate all files in a directory
./bin/migrate-v2 migrate -d ./infrastructure

# Preview changes without applying them
./bin/migrate-v2 migrate -f terraform.tf --dry-run

# Show available migrations
./bin/migrate-v2 list
```

## Features

- **Automatic Migration**: Transforms both configuration files and state files between provider versions
- **YAML-Driven**: Most migrations require only YAML configuration, no custom code needed
- **Version Chaining**: Automatically finds migration paths (e.g., v4→v5→v6)
- **Safe by Default**: Built-in backup and rollback capabilities
- **Preview Mode**: See what will change before applying migrations
- **Extensible**: Easy to add custom migrations for complex scenarios

## How It Works

The migration tool uses a registry-based system where each resource type registers its migration logic. When you run a migration:

1. The tool identifies resource types and their current versions
2. Finds the appropriate migration path (direct or multi-hop)
3. Applies transformations defined in YAML or custom code
4. Creates backups before making changes
5. Updates both configuration and state files

## Configuration

Migrations are primarily configured through YAML files that define transformations:

```yaml
version: "1.0"
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

# Simple attribute renames
attribute_renames:
  domain: hostname
  
# Convert blocks to lists
blocks_to_lists:
  - cors_headers

# Add default values for new required fields  
default_values:
  type: "self_hosted"
```

See [docs/yaml-configuration.md](docs/yaml-configuration.md) for the full YAML reference.

## Writing Custom Migrations

For complex scenarios that can't be handled by YAML alone, you can write custom migrations:

```go
type MyMigration struct {
    *core.Migration
}

func (m *MyMigration) MigrateConfig(block *hclwrite.Block, ctx *core.MigrationContext) error {
    // Custom transformation logic
    return m.Migration.MigrateConfig(block, ctx) // Apply YAML transformations
}
```

See [docs/custom-migrations.md](docs/custom-migrations.md) for details.

## Available Transformations

### Common Transformations
- **Attribute Renames**: Rename fields while preserving values
- **Field Removals**: Remove deprecated attributes
- **Default Values**: Add defaults for new required fields
- **Type Conversions**: Convert between data types

### Structural Transformations
- **Blocks to Lists**: Convert HCL blocks to list attributes
- **List to Blocks**: Convert list attributes to multiple blocks
- **Flatten Nested**: Flatten nested object structures
- **Split/Merge Objects**: Reorganize object attributes

See [docs/transformations.md](docs/transformations.md) for the complete list.

## Testing

The tool includes comprehensive testing support:

```bash
# Run unit tests
make test

# Run integration tests
make test-integration

# Test a specific migration
go test ./resources/access_application/migrations -v
```

See [docs/testing.md](docs/testing.md) for testing guidelines.

## Documentation

- [Architecture Overview](docs/architecture.md) - System design and components
- [YAML Configuration](docs/yaml-configuration.md) - Complete YAML reference
- [Transformations Guide](docs/transformations.md) - Available transformations
- [Custom Migrations](docs/custom-migrations.md) - Writing custom migration logic
- [Testing Guide](docs/testing.md) - Testing migrations
- [Examples](docs/examples.md) - Real-world migration examples

## Contributing

1. Add your resource migration to `resources/<resource_name>/migrations/`
2. Create a YAML configuration file (e.g., `v4_to_v5.yaml`)
3. Write tests to verify the migration
4. Register the migration in `resources/<resource_name>/registry.go`

## License

Mozilla Public License v2.0