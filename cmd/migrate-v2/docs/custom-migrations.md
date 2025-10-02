# Creating Custom Migrations

This guide explains how to create custom migrations for complex scenarios that can't be handled by YAML alone.

## When to Use Custom Migrations

Use custom Go migrations when you need:
- Complex conditional logic
- JSON encoding/decoding
- Dynamic field generation
- Advanced validation

For simple renames, removals, and structural changes, use YAML configuration instead.

## Quick Example

```go
package migrations

import (
    _ "embed"
    "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/core"
)

//go:embed v4_to_v5.yaml
var migrationConfig []byte

type V4ToV5Migration struct {
    *core.Migration
}

func NewV4ToV5Migration() (*V4ToV5Migration, error) {
    migration, err := core.NewMigration(migrationConfig)
    if err != nil {
        return nil, err
    }
    
    return &V4ToV5Migration{
        Migration: migration,
    }, nil
}

func (m *V4ToV5Migration) MigrateConfig(block *hclwrite.Block, ctx *core.MigrationContext) error {
    // Apply YAML transformations first
    if err := m.Migration.MigrateConfig(block, ctx); err != nil {
        return err
    }
    
    // Add custom logic here
    // Example: Convert settings block to JSON
    if err := m.convertSettingsToJSON(block, ctx); err != nil {
        return err
    }
    
    return nil
}
```

## Complex Migration Example

Here's a real-world example converting a settings block to JSON:

```go
func (m *V4ToV5Migration) convertSettingsToJSON(block *hclwrite.Block, ctx *core.MigrationContext) error {
    body := block.Body()
    
    // Find and process settings blocks
    for _, b := range body.Blocks() {
        if b.Type() == "settings" {
            // Extract block content
            data := extractBlockData(b)
            
            // Convert to jsonencode function call
            body.SetAttributeRaw("settings_json", 
                hclwrite.TokensForFunctionCall("jsonencode", data))
            
            // Remove old block
            body.RemoveBlock(b)
            
            ctx.AddInfo("Converted settings block to JSON", "", m.ResourceType())
        }
    }
    
    return nil
}
```

## Testing

```go
func TestV4ToV5Migration(t *testing.T) {
    migration, err := NewV4ToV5Migration()
    require.NoError(t, err)
    
    // Test configuration migration
    config := parseHCL(`
        resource "cloudflare_example" "test" {
            enabled = true
            settings {
                key = "value"
            }
        }
    `)
    
    ctx := core.NewMigrationContext()
    err = migration.MigrateConfig(config, ctx)
    require.NoError(t, err)
    
    // Verify transformations applied
    assert.NotNil(t, config.Body().GetAttribute("is_enabled"))
    assert.NotNil(t, config.Body().GetAttribute("settings_json"))
}
```

## Best Practices

1. **Use YAML first** - Most migrations can be handled with YAML configuration
2. **Keep it simple** - Only add custom code for truly complex logic
3. **Test thoroughly** - Include edge cases and real configurations
4. **Handle errors gracefully** - Use warnings for non-critical issues
5. **Preserve formatting** - Don't disrupt user's code style