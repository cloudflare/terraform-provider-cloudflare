package internal

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMigration(t *testing.T) {
	tests := []struct {
		name        string
		yamlContent string
		expectError bool
	}{
		{
			name: "valid yaml configuration",
			yamlContent: `
resource_type: cloudflare_test_resource
source_version: v4
target_version: v5
description: Test migration
config:
  attribute_renames:
    old_name: new_name
`,
			expectError: false,
		},
		{
			name:        "invalid yaml",
			yamlContent: `invalid: yaml: content:`,
			expectError: true,
		},
		{
			name:        "empty yaml",
			yamlContent: ``,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migration, err := NewMigration([]byte(tt.yamlContent))
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, migration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, migration)
			}
		})
	}
}

func TestMigration_AttributeRenames(t *testing.T) {
	yamlContent := `
resource_type: test_resource
source_version: v4
target_version: v5
config:
  attribute_renames:
    old_attr1: new_attr1
    old_attr2: new_attr2
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	// Create test HCL block
	hclContent := `
resource "test_resource" "test" {
  old_attr1 = "value1"
  old_attr2 = "value2"
  unchanged = "value3"
}
`
	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Apply migration
	err = migration.MigrateConfig(block, ctx)
	require.NoError(t, err)

	// Check that attributes were renamed
	body := block.Body()
	assert.Nil(t, body.GetAttribute("old_attr1"))
	assert.Nil(t, body.GetAttribute("old_attr2"))
	assert.NotNil(t, body.GetAttribute("new_attr1"))
	assert.NotNil(t, body.GetAttribute("new_attr2"))
	assert.NotNil(t, body.GetAttribute("unchanged"))
}

func TestMigration_AttributeRemovals(t *testing.T) {
	yamlContent := `
resource_type: test_resource
source_version: v4
target_version: v5
config:
  attribute_removals:
    - deprecated_attr
    - unused_attr
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	hclContent := `
resource "test_resource" "test" {
  deprecated_attr = "value1"
  unused_attr = "value2"
  keep_attr = "value3"
}
`
	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Apply migration
	err = migration.MigrateConfig(block, ctx)
	require.NoError(t, err)

	// Check that attributes were removed
	body := block.Body()
	assert.Nil(t, body.GetAttribute("deprecated_attr"))
	assert.Nil(t, body.GetAttribute("unused_attr"))
	assert.NotNil(t, body.GetAttribute("keep_attr"))
}

func TestMigration_ConditionalRemovals(t *testing.T) {
	yamlContent := `
resource_type: test_resource
source_version: v4
target_version: v5
config:
  conditional_removals:
    - attribute: conditional_attr
      condition:
        type: "type_a"
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	tests := []struct {
		name         string
		hclContent   string
		shouldRemove bool
	}{
		{
			name: "condition matches - should remove",
			hclContent: `
resource "test_resource" "test" {
  type = "type_a"
  conditional_attr = "value"
}`,
			shouldRemove: true,
		},
		{
			name: "condition does not match - should not remove",
			hclContent: `
resource "test_resource" "test" {
  type = "type_b"
  conditional_attr = "value"
}`,
			shouldRemove: false,
		},
		{
			name: "condition attribute missing - should not remove",
			hclContent: `
resource "test_resource" "test" {
  conditional_attr = "value"
}`,
			shouldRemove: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hclContent), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &MigrationContext{
				Diagnostics: []Diagnostic{},
				Metrics:     &MigrationMetrics{},
			}

			// Apply migration
			err = migration.MigrateConfig(block, ctx)
			require.NoError(t, err)

			// Check if attribute was removed based on condition
			body := block.Body()
			if tt.shouldRemove {
				assert.Nil(t, body.GetAttribute("conditional_attr"))
			} else {
				assert.NotNil(t, body.GetAttribute("conditional_attr"))
			}
		})
	}
}

func TestMigration_DefaultValues(t *testing.T) {
	yamlContent := `
resource_type: test_resource
source_version: v4
target_version: v5
config:
  default_values:
    new_string: "default_value"
    new_number: 42
    new_bool: true
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	hclContent := `
resource "test_resource" "test" {
  existing = "value"
}
`
	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Apply migration
	err = migration.MigrateConfig(block, ctx)
	require.NoError(t, err)

	// Check that default values were added
	body := block.Body()
	assert.NotNil(t, body.GetAttribute("new_string"))
	assert.NotNil(t, body.GetAttribute("new_number"))
	assert.NotNil(t, body.GetAttribute("new_bool"))
	assert.NotNil(t, body.GetAttribute("existing"))
}

func TestMigration_BlocksToLists(t *testing.T) {
	yamlContent := `
resource_type: test_resource
source_version: v4
target_version: v5
config:
  blocks_to_lists:
    - item
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	hclContent := `
resource "test_resource" "test" {
  name = "test"
  
  item {
    key = "value1"
  }
  
  item {
    key = "value2"
  }
}
`
	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Apply migration
	err = migration.MigrateConfig(block, ctx)
	require.NoError(t, err)

	// Check that blocks were converted to list attribute
	body := block.Body()
	assert.NotNil(t, body.GetAttribute("item"))

	// Check that blocks were removed
	hasItemBlock := false
	for _, b := range body.Blocks() {
		if b.Type() == "item" {
			hasItemBlock = true
		}
	}
	assert.False(t, hasItemBlock)
}

func TestMigration_StateMigration(t *testing.T) {
	yamlContent := `
resource_type: test_resource
source_version: v4
target_version: v5
state:
  attribute_renames:
    old_state_attr: new_state_attr
  schema_version: 2
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	state := map[string]interface{}{
		"id": "resource-123",
		"attributes": map[string]interface{}{
			"old_state_attr": "value",
			"keep_attr":      "keep",
		},
	}

	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Apply state migration
	err = migration.MigrateState(state, ctx)
	require.NoError(t, err)

	// Check state changes - attributes are nested under "attributes" key
	attrs := state["attributes"].(map[string]interface{})
	assert.Equal(t, "value", attrs["new_state_attr"])
	assert.Nil(t, attrs["old_state_attr"])
	assert.Equal(t, "keep", attrs["keep_attr"])
	assert.Equal(t, 2, state["schema_version"])
}

func TestMigration_ComplexYAML(t *testing.T) {
	yamlContent := `
resource_type: complex_resource
source_version: v4
target_version: v5
description: Complex migration test
config:
  attribute_renames:
    old1: new1
    old2: new2
  attribute_removals:
    - deprecated
  conditional_removals:
    - attribute: conditional
      condition:
        enabled: "false"
  type_conversions:
    - attribute: set_attr
      from_type: set
      to_type: list
  blocks_to_lists:
    - nested
  default_values:
    new_field: "default"
state:
  attribute_renames:
    state_old: state_new
  schema_version: 3
`
	migration, err := NewMigration([]byte(yamlContent))
	require.NoError(t, err)

	assert.Equal(t, "complex_resource", migration.ResourceType())
	assert.Equal(t, "v4", migration.SourceVersion())
	assert.Equal(t, "v5", migration.TargetVersion())

	// Verify complex config was parsed correctly
	assert.Equal(t, 2, len(migration.config.Config.AttributeRenames))
	assert.Equal(t, 1, len(migration.config.Config.AttributeRemovals))
	assert.Equal(t, 1, len(migration.config.Config.ConditionalRemovals))
	assert.Equal(t, 1, len(migration.config.Config.TypeConversions))
	assert.Equal(t, 1, len(migration.config.Config.BlocksToLists))
	assert.Equal(t, 1, len(migration.config.Config.DefaultValues))
	assert.Equal(t, 1, len(migration.config.State.AttributeRenames))
	assert.Equal(t, 3, migration.config.State.SchemaVersion)
}
