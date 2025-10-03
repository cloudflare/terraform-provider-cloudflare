# Terraform Cloudflare Provider Migration Tool (v2)

Automatically migrate Terraform configurations between Cloudflare provider versions (v4 → v5).

## Quick Start

```bash
# Build the tool
go build -o ./bin/migrate-v2 .

# Migrate all .tf files in a directory
./bin/migrate-v2 -config ./terraform

# Preview changes without applying
./bin/migrate-v2 -config ./terraform -dry-run -preview

# Migrate with state file
./bin/migrate-v2 -config ./terraform -state terraform.tfstate
```

## Features

- **YAML-Driven**: Define migrations with simple YAML configuration
- **Automatic Resource Discovery**: Resources auto-register when added to `resources/imports.go`
- **Safe by Default**: Preview mode and automatic backups
- **Resource Splitting**: Handle complex cases like splitting one resource into multiple
- **State Migration**: Updates both configuration and state files

## Adding New Resources

1. Create your resource structure under `resources/<resource_name>/`
2. Define migration rules in YAML (e.g., `migrations/v4_to_v5.yaml`)
3. Add two lines to `resources/imports.go`
4. That's it! See [resources/README.md](resources/README.md) for details.

## Command Line Options

```
-config string     Directory with Terraform files (required)
-state string      Path to Terraform state file (optional)
-dry-run          Preview changes without applying
-preview          Show detailed change preview
-verbose          Enable verbose output
-backup           Create backups (default: true)
-from string      Source version (default: "v4")
-to string        Target version (default: "v5")
```

## Testing

```bash
# Run all tests
go test ./...

# Test specific resource
go test ./resources/argo/...

# Run with verbose output
go test -v ./resources/access_application/migrations/...
```

## Documentation

- [Adding Resources](resources/README.md) - How to add new resource migrations
- [Architecture](docs/architecture.md) - System design overview
- [YAML Schema](docs/yaml-configuration.md) - YAML configuration reference

## License

Mozilla Public License v2.0