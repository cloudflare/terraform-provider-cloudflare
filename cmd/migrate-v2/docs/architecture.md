# Architecture Overview

## System Design

```
┌─────────┐     ┌──────────────┐     ┌──────────┐
│   CLI   │────▶│  Orchestrator │────▶│ Registry │
└─────────┘     └──────────────┘     └──────────┘
                        │                    │
                        ▼                    ▼
            ┌──────────────────┐     ┌──────────┐
            │ Resource Imports │────▶│Migration │
            └──────────────────┘     └──────────┘
```

## Core Components

### CLI (`main.go`)
- Parses command-line arguments
- Calls `resources.RegisterAll()` to auto-register resources
- Creates migration context and runs orchestrator

### Orchestrator
- Finds and processes all .tf files
- Applies file-level transformations (for resource splitting)
- Handles backup and rollback
- Generates migration reports

### Registry & Auto-Registration
- Resources auto-register via `resources/imports.go`
- Each resource implements `RegisterMigrations()`
- No need to modify main.go for new resources

### Migrations
- YAML configuration defines transformation rules
- Custom Go code only when necessary (e.g., resource splitting)
- Support both config and state transformations

## Extension Points

### Adding New Resources
1. Create `resources/<name>/` with YAML config
2. Add import and registration to `resources/imports.go`
3. Done! See [resources/README.md](../resources/README.md)

### Custom Transformations
When YAML isn't enough:
```go
type CustomMigration struct {
    *internal.Migration
}

func (m *CustomMigration) TransformFile(file *hclwrite.File) error {
    // File-level transformations (e.g., resource splitting)
}
```

## Design Principles
- **YAML-first**: Minimize custom code
- **Safe**: Preview mode and automatic backups
- **Modular**: Each resource manages its own migrations
- **Testable**: Unit and integration tests for all migrations