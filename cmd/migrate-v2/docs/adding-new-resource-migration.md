# Adding a New Resource Migration

This guide explains how to add migration support for a new Terraform resource.

## Directory Structure

For each resource that needs migration support, create the following structure:

```
resources/
└── <resource_name>/
    ├── migrations/
    │   ├── v4_to_v5.go      # Migration implementation
    │   ├── v4_to_v5.yaml     # Migration configuration
    │   └── v4_to_v5_test.go  # Migration tests
    └── integration/
        └── testdata/
            ├── v4/           # v4 test configurations
            └── v5/           # Expected v5 outputs
```

## Step 1: Create the Migration Configuration (YAML)

Create a YAML file defining the transformation rules:

```yaml
# resources/<resource_name>/migrations/v4_to_v5.yaml
resource_type: cloudflare_<resource_name>
source_version: v4
target_version: v5
description: "Migrate <resource_name> from v4 to v5"

config:
  attribute_renames:
    old_name: new_name
    
  attribute_removals:
    - deprecated_field
    
  lists_to_blocks:
    - configuration_list
    
  blocks_to_lists:
    - settings_block
    
  default_values:
    new_field: "default_value"
    
  structural_changes:
    - type: "flatten_nested"
      source: "nested_config"
      parameters:
        separator: "_"
        depth: 2

state:
  attribute_renames:
    old_name: new_name
  schema_version: 1
```

## Step 2: Create the Migration Implementation

Use the generic pattern for creating migrations:

```go
// resources/<resource_name>/migrations/v4_to_v5.go
package migrations

import (
    _ "embed"
    "fmt"
    
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

//go:embed v4_to_v5.yaml
var v4ToV5Config []byte

// <ResourceName>Migration handles migrations for cloudflare_<resource_name>
type <ResourceName>Migration struct {
    *internal.Migration
}

// New<ResourceName>Migration creates a new migration from config
func New<ResourceName>Migration(config []byte) (*<ResourceName>Migration, error) {
    base, err := internal.NewMigration(config)
    if err != nil {
        return nil, fmt.Errorf("failed to create base migration: %w", err)
    }
    
    return &<ResourceName>Migration{
        Migration: base,
    }, nil
}

// NewV4ToV5Migration creates the v4 to v5 migration
func NewV4ToV5Migration() (*<ResourceName>Migration, error) {
    return New<ResourceName>Migration(v4ToV5Config)
}

// RegisterV4ToV5 registers the migration with a registry
func RegisterV4ToV5(registry internal.MigrationRegistry) error {
    migration, err := NewV4ToV5Migration()
    if err != nil {
        return err
    }
    return registry.Register(migration)
}
```

## Step 3: Add Custom Migration Logic (Optional)

If you need custom transformation logic beyond what YAML supports:

```go
// Override MigrateConfig for custom config transformations
func (m *<ResourceName>Migration) MigrateConfig(block *hclwrite.Block, ctx *internal.MigrationContext) error {
    // First apply standard YAML-based transformations
    if err := m.Migration.MigrateConfig(block, ctx); err != nil {
        return err
    }
    
    // Add custom logic here
    // Example: Complex conditional transformation
    body := block.Body()
    if attr := body.GetAttribute("special_field"); attr != nil {
        // Custom transformation logic
    }
    
    return nil
}

// Override MigrateState for custom state transformations
func (m *<ResourceName>Migration) MigrateState(state map[string]interface{}, ctx *internal.MigrationContext) error {
    // First apply standard YAML-based transformations
    if err := m.Migration.MigrateState(state, ctx); err != nil {
        return err
    }
    
    // Add custom state migration logic
    if val, exists := state["complex_field"]; exists {
        // Transform the state value
        state["complex_field"] = transformComplexField(val)
    }
    
    return nil
}
```

## Step 4: Write Tests

Create comprehensive tests for your migration:

```go
// resources/<resource_name>/migrations/v4_to_v5_test.go
package migrations

import (
    "testing"
    
    "github.com/hashicorp/hcl/v2/hclwrite"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

func TestV4ToV5Migration_BasicConfig(t *testing.T) {
    migration, err := NewV4ToV5Migration()
    require.NoError(t, err)
    
    // Create test input
    block := hclwrite.NewBlock("resource", []string{"cloudflare_<resource_name>", "test"})
    body := block.Body()
    body.SetAttributeValue("old_name", cty.StringVal("test_value"))
    
    // Apply migration
    ctx := internal.NewMigrationContext()
    err = migration.MigrateConfig(block, ctx)
    require.NoError(t, err)
    
    // Verify transformation
    assert.Nil(t, body.GetAttribute("old_name"))
    assert.NotNil(t, body.GetAttribute("new_name"))
}

func TestV4ToV5Migration_State(t *testing.T) {
    migration, err := NewV4ToV5Migration()
    require.NoError(t, err)
    
    // Create test state
    state := map[string]interface{}{
        "old_name": "value",
        "id": "test-id",
    }
    
    // Apply migration
    ctx := internal.NewMigrationContext()
    err = migration.MigrateState(state, ctx)
    require.NoError(t, err)
    
    // Verify state transformation
    assert.NotContains(t, state, "old_name")
    assert.Equal(t, "value", state["new_name"])
    assert.Equal(t, 1, state["schema_version"])
}
```

## Step 5: Integration Testing

Add integration test configurations:

```hcl
# resources/<resource_name>/integration/testdata/v4/basic.tf
resource "cloudflare_<resource_name>" "test" {
  old_name = "test_value"
  deprecated_field = "remove_me"
  
  settings_block {
    key = "value"
  }
}
```

```hcl
# resources/<resource_name>/integration/testdata/v5/basic.tf
resource "cloudflare_<resource_name>" "test" {
  new_name = "test_value"
  new_field = "default_value"
  
  settings_block = [{
    key = "value"
  }]
}
```

## Step 6: Register the Migration

Add your migration to the main registry during application initialization:

```go
// In your main application or registry initialization
registry := internal.NewDefaultRegistry()
if err := migrations.RegisterV4ToV5(registry); err != nil {
    log.Fatal("Failed to register migration:", err)
}
```

## Common Transformation Patterns

### 1. Blocks to Lists
Used when v5 changes repeated blocks to list attributes:
```yaml
blocks_to_lists:
  - rule
  - policy
```

### 2. Lists to Blocks  
Used when v5 changes list attributes to repeated blocks:
```yaml
lists_to_blocks:
  - configurations
```

### 3. Flatten Nested Objects
Used when v5 flattens nested structures:
```yaml
structural_changes:
  - type: "flatten_nested"
    source: "nested_settings"
    parameters:
      separator: "_"
      depth: 2
```

### 4. String Arrays to Object Arrays
Common for ID lists becoming object lists:
```yaml
structural_changes:
  - transform: "string_list_to_object_list"
    source: "policy_ids"
    parameters:
      object_key: "id"
```

## Testing Your Migration

Run unit tests:
```bash
go test ./resources/<resource_name>/migrations/...
```

Run integration tests:
```bash
go test ./resources/<resource_name>/integration/...
```

Test with real configurations:
```bash
./bin/migrate-v2 -config ./test-configs -from v4 -to v5
```

## Debugging Tips

1. Use `-verbose` flag to see detailed transformation steps
2. Use `-dry-run` to preview changes without modifying files
3. Check the context diagnostics for warnings and errors
4. Use `-preview` to see a diff of changes before applying

## Supporting Multiple Versions

To support multiple migration paths (e.g., v4→v5 and v5→v6):

```go
//go:embed v4_to_v5.yaml
var v4ToV5Config []byte

//go:embed v5_to_v6.yaml  
var v5ToV6Config []byte

func NewV4ToV5Migration() (*<ResourceName>Migration, error) {
    return New<ResourceName>Migration(v4ToV5Config)
}

func NewV5ToV6Migration() (*<ResourceName>Migration, error) {
    return New<ResourceName>Migration(v5ToV6Config)
}

// Register all migrations
func RegisterMigrations(registry internal.MigrationRegistry) error {
    if err := RegisterV4ToV5(registry); err != nil {
        return err
    }
    if err := RegisterV5ToV6(registry); err != nil {
        return err
    }
    return nil
}
```