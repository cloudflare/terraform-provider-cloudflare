# Migration Tool v2 - Technical Specification

## Overview

The migrate-v2 tool provides automated migration of Terraform configurations between different versions of the Cloudflare provider. It uses a combination of YAML-driven transformations and optional custom code to handle schema changes, attribute renames, and structural reorganizations.

## Architecture

```
cmd/migrate-v2/
├── main.go                          # CLI entry point
├── core/                            # Core migration engine
│   ├── interfaces.go                # Core interfaces and types
│   ├── registry.go                  # Migration registry
│   ├── orchestrator.go              # Migration orchestration
│   ├── migration.go                 # Base migration implementation
│   ├── context.go                   # Migration context
│   ├── preview.go                   # Change preview generation
│   ├── backup.go                    # Backup/rollback system
│   └── progress.go                  # Progress tracking
├── transformations/                 # Transformation functions
│   ├── config/                      # HCL configuration transformations
│   │   ├── basic/                   # Simple transformations
│   │   │   ├── attribute_renames.go
│   │   │   ├── field_removals.go
│   │   │   ├── default_injector.go
│   │   │   └── type_conversions.go
│   │   ├── structural/              # Complex structural transformations
│   │   │   ├── blocks_to_lists.go
│   │   │   ├── list_to_blocks.go
│   │   │   ├── flatten_nested.go
│   │   │   ├── split_object.go
│   │   │   └── merge_attributes.go
│   │   └── conditional/             # Conditional transformations
│   │       └── conditional_attribute_remover.go
│   └── state/                       # State transformations
│       ├── attribute_renames.go
│       ├── field_removals.go
│       └── schema_updater.go
└── resources/                       # Resource-specific migrations
    └── <resource_name>/
        └── migrations/
            ├── v4_to_v5.yaml        # YAML configuration
            └── v4_to_v5_test.go     # Tests

```

## Core Components

### Migration Interface

```go
type ResourceMigration interface {
    ResourceType() string
    TargetResourceType() string  
    Version() MigrationVersion
    MigrateConfig(*hclwrite.Block, *MigrationContext) error
    MigrateState(map[string]interface{}, *MigrationContext) error
}
```

### Transformation Pipeline

Transformations are applied in a specific order to ensure predictable results:

1. **Structural Changes** (blocks_to_lists, list_to_blocks)
2. **Attribute Operations** (renames, removals)
3. **Type Conversions** 
4. **Default Values** (applied last)

### YAML Configuration

The YAML configuration drives most migrations without requiring custom code:

```yaml
version: "1.0"
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"

# Structural changes
blocks_to_lists:
  - cors_headers
  
lists_to_blocks:
  - destinations

# Attribute operations  
attribute_renames:
  domain: hostname
  
remove_fields:
  - deprecated_field

# Complex transformations
structural_changes:
  - type: flatten_nested
    source: address
    parameters:
      separator: "_"
      max_depth: 2
      
  - type: split_object
    source: config
    parameters:
      attributes:
        - host
        - port
      prefix: "server_"

# Default values
default_values:
  type: "self_hosted"
  enabled: true
```

## Transformation Types

### Config Transformations

#### Basic (`config/basic`)
- **AttributeRenamer**: Renames attributes while preserving values
- **AttributeRemover**: Removes specified attributes
- **DefaultValueSetter**: Sets default values for missing attributes
- **SetToListConverter**: Converts `toset()` to list syntax

#### Structural (`config/structural`)
- **BlocksToListConverter**: Converts HCL blocks to list attributes
- **ListToBlocksConverter**: Converts list attributes to multiple blocks
- **ListToBlocksWithMapping**: Converts lists to blocks with field renaming
- **FlattenNestedTransformer**: Flattens nested object structures
- **SplitObjectTransformer**: Splits object attributes into flat attributes
- **MergeAttributesTransformer**: Merges multiple attributes into objects/lists

#### Conditional (`config/conditional`)
- **ConditionalAttributeRemover**: Removes attributes based on conditions

### State Transformations (`state`)
- **AttributeRenamer**: Renames fields in state JSON
- **FieldRemover**: Removes fields from state
- **SchemaVersionUpdater**: Updates schema version
- **DefaultValueSetter**: Sets defaults for missing state fields

## Migration Process

1. **Discovery**: Identify resource types and versions
2. **Path Finding**: Determine migration path (direct or multi-hop)
3. **Backup**: Create backup of original files
4. **Transformation**: Apply configured transformations
5. **Validation**: Verify transformed output
6. **Writing**: Save transformed files

## Testing

Each migration should include comprehensive tests:

```go
func TestMigration_V4toV5(t *testing.T) {
    // Test configuration migration
    config := LoadTestConfig("testdata/v4_config.tf")
    migration := NewV4toV5Migration()
    
    err := migration.MigrateConfig(config, ctx)
    assert.NoError(t, err)
    
    // Verify transformations
    assert.Equal(t, "hostname", config.GetAttribute("hostname"))
    assert.Nil(t, config.GetAttribute("domain"))
}
```

## Extension Points

The system is designed for extensibility:

1. **Custom Transformers**: Implement `TransformerFunc` interface
2. **Custom Migrations**: Extend `Migration` for complex logic
3. **New Resource Types**: Add to `resources/` directory
4. **New Transformation Types**: Add to `transformations/` package

## Performance Considerations

- Transformations operate on AST, avoiding string manipulation
- Parallel processing for multiple files
- Minimal memory footprint through streaming
- Efficient token-based parsing for HCL

## Error Handling

- Non-destructive by default (backups before changes)
- Comprehensive diagnostics collection
- Rollback capability on failure
- Clear error messages with context