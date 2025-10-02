# Architecture Overview

## System Design

The migration tool follows a modular, extensible architecture:

```
┌─────────────┐     ┌──────────────┐     ┌─────────────┐
│     CLI     │────▶│ Orchestrator │────▶│  Registry   │
└─────────────┘     └──────────────┘     └─────────────┘
                            │                     │
                            ▼                     ▼
                    ┌──────────────┐     ┌─────────────┐
                    │  Migration   │◀────│  Resources  │
                    │    Engine    │     │ (v4→v5, etc)│
                    └──────────────┘     └─────────────┘
                            │
                            ▼
                    ┌──────────────┐
                    │Transformers  │
                    │ (HCL & JSON) │
                    └──────────────┘
```

## Core Components

### CLI
Entry point that handles:
- Command-line arguments
- File discovery
- Progress reporting
- Error handling

### Orchestrator
Coordinates the migration process:
- Identifies resources and versions
- Finds migration paths (direct or multi-hop)
- Manages backup and rollback
- Applies migrations in order

### Registry
Maintains available migrations:
- Resource type → migration mapping
- Version routing (v4→v5, v5→v6)
- Path finding for multi-hop migrations

### Migration Engine
Executes transformations:
- Loads YAML configuration
- Builds transformation pipeline
- Applies transformers in order
- Validates results

### Transformers
Reusable transformation functions:
- **Common**: Work on both config and state
- **Config**: HCL-specific transformations
- **State**: JSON state file transformations

## Data Flow

1. **Input**: Terraform configuration files (.tf)
2. **Parsing**: Convert to HCL AST using hclwrite
3. **Transformation**: Apply migrations via transformers
4. **Output**: Write transformed configuration
5. **State**: Optionally update state files

## Extension Points

### Adding New Resources

1. Create directory: `resources/<resource_name>/migrations/`
2. Add YAML configuration: `v4_to_v5.yaml`
3. Write tests: `v4_to_v5_test.go`
4. Register in: `resources/<resource_name>/registry.go`

### Adding New Transformers

1. Implement `TransformerFunc` interface
2. Add to appropriate package (`common/`, `config/`, or `state/`)
3. Register in migration pipeline
4. Add tests

### Custom Migrations

For complex scenarios beyond YAML:

```go
type CustomMigration struct {
    *core.Migration
}

func (m *CustomMigration) MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error {
    // Apply YAML transformations first
    if err := m.Migration.MigrateConfig(block, ctx); err != nil {
        return err
    }
    
    // Add custom logic
    // ...
    
    return nil
}
```

## Design Principles

1. **Non-destructive**: Always backup before changes
2. **Predictable**: Transformations apply in consistent order
3. **Extensible**: Easy to add new resources and transformations
4. **Testable**: Comprehensive test coverage
5. **User-friendly**: Clear error messages and progress reporting