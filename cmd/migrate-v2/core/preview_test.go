package core

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock registry for preview tests
type mockPreviewRegistry struct {
	migrations map[string]ResourceMigration
}

func newMockPreviewRegistry() *mockPreviewRegistry {
	return &mockPreviewRegistry{
		migrations: make(map[string]ResourceMigration),
	}
}

func (r *mockPreviewRegistry) Register(migration ResourceMigration) error {
	key := migration.ResourceType() + ":" + migration.SourceVersion() + ":" + migration.TargetVersion()
	r.migrations[key] = migration
	return nil
}

func (r *mockPreviewRegistry) Get(resourceType, sourceVersion, targetVersion string) (ResourceMigration, error) {
	key := resourceType + ":" + sourceVersion + ":" + targetVersion
	if migration, exists := r.migrations[key]; exists {
		return migration, nil
	}
	return nil, fmt.Errorf("migration not found")
}

func (r *mockPreviewRegistry) GetPath(resourceType, sourceVersion, targetVersion string) ([]ResourceMigration, error) {
	migration, err := r.Get(resourceType, sourceVersion, targetVersion)
	if err != nil {
		return nil, err
	}
	return []ResourceMigration{migration}, nil
}

func (r *mockPreviewRegistry) ListAvailable() []MigrationInfo {
	var infos []MigrationInfo
	for _, migration := range r.migrations {
		infos = append(infos, MigrationInfo{
			ResourceType:  migration.ResourceType(),
			SourceVersion: migration.SourceVersion(),
			TargetVersion: migration.TargetVersion(),
		})
	}
	return infos
}

// Mock migration for testing preview generation
type mockPreviewMigration struct {
	resourceType  string
	sourceVersion string
	targetVersion string
	changes       []func(*hclwrite.Block)
}

func (m *mockPreviewMigration) ResourceType() string  { return m.resourceType }
func (m *mockPreviewMigration) SourceVersion() string { return m.sourceVersion }
func (m *mockPreviewMigration) TargetVersion() string { return m.targetVersion }
func (m *mockPreviewMigration) Validate(block *hclwrite.Block) error { return nil }
func (m *mockPreviewMigration) MigrateState(state map[string]interface{}, ctx *MigrationContext) error {
	return nil
}

func (m *mockPreviewMigration) MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error {
	// Apply mock changes
	for _, change := range m.changes {
		change(block)
	}
	return nil
}

func TestChangeType_Constants(t *testing.T) {
	// Verify change type constants
	assert.Equal(t, ChangeType("ADD"), ChangeTypeAdd)
	assert.Equal(t, ChangeType("REMOVE"), ChangeTypeRemove)
	assert.Equal(t, ChangeType("MODIFY"), ChangeTypeModify)
	assert.Equal(t, ChangeType("RENAME"), ChangeTypeRename)
	assert.Equal(t, ChangeType("RESTRUCTURE"), ChangeTypeRestructure)
}

func TestImpactLevel_Constants(t *testing.T) {
	// Verify impact level constants
	assert.Equal(t, ImpactLevel("LOW"), ImpactLow)
	assert.Equal(t, ImpactLevel("MEDIUM"), ImpactMedium)
	assert.Equal(t, ImpactLevel("HIGH"), ImpactHigh)
	assert.Equal(t, ImpactLevel("BREAKING"), ImpactBreaking)
}

func TestPreviewGenerator_GeneratePreview(t *testing.T) {
	registry := newMockPreviewRegistry()
	pg := NewPreviewGenerator(registry)

	// Register a mock migration
	migration := &mockPreviewMigration{
		resourceType:  "cloudflare_test",
		sourceVersion: "v4",
		targetVersion: "v5",
		changes: []func(*hclwrite.Block){
			func(block *hclwrite.Block) {
				// Remove an attribute
				block.Body().RemoveAttribute("old_attr")
				// Add a new attribute
				block.Body().SetAttributeRaw("new_attr", hclwrite.Tokens{
					{Type: hclsyntax.TokenIdent, Bytes: []byte(`"value"`)},
				})
			},
		},
	}
	registry.Register(migration)

	// Create test blocks
	file, _ := hclwrite.ParseConfig([]byte(`
resource "cloudflare_test" "example" {
	name = "test"
	old_attr = "old_value"
}

resource "cloudflare_other" "example" {
	name = "other"
}
`), "test.tf", hcl.InitialPos)

	blocks := file.Body().Blocks()

	// Generate preview
	preview, err := pg.GeneratePreview(blocks, "v4", "v5")
	require.NoError(t, err)
	require.NotNil(t, preview)

	// Verify preview metadata
	assert.Equal(t, "v4", preview.FromVersion)
	assert.Equal(t, "v5", preview.ToVersion)
	
	// Should only preview the resource with available migration
	assert.Len(t, preview.Resources, 1)
	assert.Equal(t, "cloudflare_test", preview.Resources[0].Type)
	assert.Equal(t, "example", preview.Resources[0].Name)
	
	// Verify changes detected
	assert.NotEmpty(t, preview.Resources[0].ConfigChanges)
}

func TestPreviewGenerator_detectConfigChanges(t *testing.T) {
	pg := &PreviewGenerator{}

	tests := []struct {
		name     string
		original string
		modified string
		expected []ChangeType
	}{
		{
			name: "attribute removed",
			original: `resource "test" "example" {
				name = "test"
				old_attr = "value"
			}`,
			modified: `resource "test" "example" {
				name = "test"
			}`,
			expected: []ChangeType{ChangeTypeRemove},
		},
		{
			name: "attribute added",
			original: `resource "test" "example" {
				name = "test"
			}`,
			modified: `resource "test" "example" {
				name = "test"
				new_attr = "value"
			}`,
			expected: []ChangeType{ChangeTypeAdd},
		},
		{
			name: "attribute modified",
			original: `resource "test" "example" {
				name = "old_name"
			}`,
			modified: `resource "test" "example" {
				name = "new_name"
			}`,
			expected: []ChangeType{ChangeTypeModify},
		},
		{
			name: "multiple changes",
			original: `resource "test" "example" {
				name = "test"
				old_attr = "value"
				modify_attr = "old"
			}`,
			modified: `resource "test" "example" {
				name = "test"
				modify_attr = "new"
				new_attr = "added"
			}`,
			expected: []ChangeType{ChangeTypeRemove, ChangeTypeModify, ChangeTypeAdd},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origFile, _ := hclwrite.ParseConfig([]byte(tt.original), "test.tf", hcl.InitialPos)
			modFile, _ := hclwrite.ParseConfig([]byte(tt.modified), "test.tf", hcl.InitialPos)

			origBlock := origFile.Body().Blocks()[0]
			modBlock := modFile.Body().Blocks()[0]

			changes := pg.detectConfigChanges(origBlock, modBlock)

			// Verify we got the expected change types
			changeTypes := make(map[ChangeType]bool)
			for _, change := range changes {
				changeTypes[change.Type] = true
			}

			for _, expectedType := range tt.expected {
				assert.True(t, changeTypes[expectedType], "Expected change type %s not found", expectedType)
			}
		})
	}
}

func TestPreviewGenerator_calculateImpact(t *testing.T) {
	pg := &PreviewGenerator{}

	tests := []struct {
		name     string
		changes  []Change
		expected ImpactLevel
	}{
		{
			name:     "no changes",
			changes:  []Change{},
			expected: ImpactLow,
		},
		{
			name: "only renames",
			changes: []Change{
				{Type: ChangeTypeRename, Path: "attr1"},
			},
			expected: ImpactLow,
		},
		{
			name: "additions",
			changes: []Change{
				{Type: ChangeTypeAdd, Path: "new_attr"},
			},
			expected: ImpactMedium,
		},
		{
			name: "removals",
			changes: []Change{
				{Type: ChangeTypeRemove, Path: "old_attr"},
			},
			expected: ImpactMedium,
		},
		{
			name: "structural changes",
			changes: []Change{
				{Type: ChangeTypeRestructure, Path: "block"},
			},
			expected: ImpactHigh,
		},
		{
			name: "mixed changes with restructure",
			changes: []Change{
				{Type: ChangeTypeAdd, Path: "new_attr"},
				{Type: ChangeTypeRemove, Path: "old_attr"},
				{Type: ChangeTypeRestructure, Path: "block"},
			},
			expected: ImpactHigh, // Restructure takes precedence
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impact := pg.calculateImpact(tt.changes)
			assert.Equal(t, tt.expected, impact)
		})
	}
}

func TestPreviewGenerator_calculateSummary(t *testing.T) {
	pg := &PreviewGenerator{}

	resources := []ResourcePreview{
		{
			Type: "cloudflare_test1",
			Name: "example1",
			ConfigChanges: []Change{
				{Type: ChangeTypeAdd},
				{Type: ChangeTypeRemove},
			},
			Impact: ImpactMedium,
		},
		{
			Type: "cloudflare_test2",
			Name: "example2",
			ConfigChanges: []Change{
				{Type: ChangeTypeModify},
			},
			Impact: ImpactLow,
		},
		{
			Type:          "cloudflare_test3",
			Name:          "example3",
			ConfigChanges: []Change{},
			Impact:        ImpactLow,
		},
	}

	summary := pg.calculateSummary(resources)

	assert.Equal(t, 3, summary.TotalResources)
	assert.Equal(t, 2, summary.ChangedResources) // Only 2 have changes
	assert.Equal(t, 3, summary.TotalChanges)      // 2 + 1 changes
	assert.Equal(t, 1, summary.ImpactCounts["MEDIUM"])
	assert.Equal(t, 2, summary.ImpactCounts["LOW"])
	assert.Equal(t, 1, summary.TypeCounts["ADD"])
	assert.Equal(t, 1, summary.TypeCounts["REMOVE"])
	assert.Equal(t, 1, summary.TypeCounts["MODIFY"])
}

func TestMigrationPreview_RenderDiff(t *testing.T) {
	preview := &MigrationPreview{
		FromVersion: "v4",
		ToVersion:   "v5",
		Resources: []ResourcePreview{
			{
				Type: "cloudflare_test",
				Name: "example",
				ConfigChanges: []Change{
					{
						Type:        ChangeTypeRemove,
						Path:        "old_attr",
						OldValue:    "old_value",
						Description: "Attribute 'old_attr' will be removed",
					},
					{
						Type:        ChangeTypeAdd,
						Path:        "new_attr",
						NewValue:    "new_value",
						Description: "Attribute 'new_attr' will be added",
					},
				},
				Impact: ImpactMedium,
				Warnings: []string{
					"Manual review required",
				},
			},
		},
		Summary: PreviewSummary{
			TotalResources:   1,
			ChangedResources: 1,
			TotalChanges:     2,
			ImpactCounts:     map[string]int{"MEDIUM": 1},
			TypeCounts:       map[string]int{"ADD": 1, "REMOVE": 1},
		},
	}

	output := preview.RenderDiff()

	// Verify key elements are present in output
	assert.Contains(t, output, "Migration Preview: v4 → v5")
	assert.Contains(t, output, "Resources to migrate: 1")
	assert.Contains(t, output, "Resources with changes: 1")
	assert.Contains(t, output, "Total changes: 2")
	assert.Contains(t, output, "cloudflare_test.example")
	assert.Contains(t, output, "Impact: MEDIUM")
	assert.Contains(t, output, "old_attr")
	assert.Contains(t, output, "new_attr")
	assert.Contains(t, output, "Manual review required")
	assert.Contains(t, output, "MEDIUM: 1 resources")
}

func TestFormatChange(t *testing.T) {
	tests := []struct {
		name     string
		change   Change
		expected string
	}{
		{
			name: "add change",
			change: Change{
				Type:     ChangeTypeAdd,
				Path:     "new_attr",
				NewValue: "value",
			},
			expected: "+ new_attr: value",
		},
		{
			name: "remove change",
			change: Change{
				Type:     ChangeTypeRemove,
				Path:     "old_attr",
				OldValue: "value",
			},
			expected: "- old_attr: value",
		},
		{
			name: "modify change",
			change: Change{
				Type:     ChangeTypeModify,
				Path:     "attr",
				OldValue: "old",
				NewValue: "new",
			},
			expected: "~ attr: old → new",
		},
		{
			name: "rename change",
			change: Change{
				Type:     ChangeTypeRename,
				Path:     "resource_type",
				OldValue: "old_type",
				NewValue: "new_type",
			},
			expected: "→ resource_type: old_type → new_type",
		},
		{
			name: "restructure change",
			change: Change{
				Type:        ChangeTypeRestructure,
				Path:        "block",
				Description: "Block structure changed",
			},
			expected: "※ block: Block structure changed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := formatChange(tt.change)
			output = strings.TrimSpace(output)
			assert.Contains(t, output, tt.expected)
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"short", 10, "short"},
		{"this is a long string", 10, "this is..."},
		{"exact", 5, "exact"},
		{"", 5, ""},
	}

	for _, tt := range tests {
		result := truncate(tt.input, tt.maxLen)
		assert.Equal(t, tt.expected, result)
	}
}

func TestPreviewGenerator_getAttributeValue(t *testing.T) {
	pg := &PreviewGenerator{}

	// Create an attribute
	file, _ := hclwrite.ParseConfig([]byte(`
		resource "test" "example" {
			name = "test_value"
			number = 42
			bool_val = true
		}
	`), "test.tf", hcl.InitialPos)

	block := file.Body().Blocks()[0]
	body := block.Body()

	// Test string attribute
	nameAttr := body.GetAttribute("name")
	require.NotNil(t, nameAttr)
	value := pg.getAttributeValue(nameAttr)
	assert.Equal(t, `"test_value"`, value)

	// Test number attribute
	numAttr := body.GetAttribute("number")
	require.NotNil(t, numAttr)
	value = pg.getAttributeValue(numAttr)
	assert.Equal(t, "42", value)

	// Test boolean attribute
	boolAttr := body.GetAttribute("bool_val")
	require.NotNil(t, boolAttr)
	value = pg.getAttributeValue(boolAttr)
	assert.Equal(t, "true", value)
}

func TestPreviewGenerator_getBlocksByType(t *testing.T) {
	pg := &PreviewGenerator{}

	file, _ := hclwrite.ParseConfig([]byte(`
		resource "test" "example" {
			network {
				name = "net1"
			}
			network {
				name = "net2"
			}
			rule {
				action = "allow"
			}
		}
	`), "test.tf", hcl.InitialPos)

	block := file.Body().Blocks()[0]
	counts := pg.getBlocksByType(block.Body())

	assert.Equal(t, 2, counts["network"])
	assert.Equal(t, 1, counts["rule"])
	assert.Equal(t, 0, counts["nonexistent"])
}

func TestPreviewGenerator_cloneBlock(t *testing.T) {
	pg := &PreviewGenerator{}

	// Create original block
	file, _ := hclwrite.ParseConfig([]byte(`
		resource "test" "example" {
			name = "original"
			value = 123
		}
	`), "test.tf", hcl.InitialPos)

	original := file.Body().Blocks()[0]
	
	// Clone the block
	cloned := pg.cloneBlock(original)
	
	// Modify the clone
	cloned.Body().SetAttributeRaw("name", hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`"modified"`)},
	})
	
	// Verify original is unchanged
	origName := original.Body().GetAttribute("name")
	require.NotNil(t, origName)
	origValue := pg.getAttributeValue(origName)
	assert.Contains(t, origValue, "original")
	
	// Verify clone was modified
	clonedName := cloned.Body().GetAttribute("name")
	require.NotNil(t, clonedName)
	clonedValue := pg.getAttributeValue(clonedName)
	assert.Contains(t, clonedValue, "modified")
}

func TestResourcePreview_WithNewType(t *testing.T) {
	preview := ResourcePreview{
		Type:    "old_resource_type",
		Name:    "example",
		NewType: "new_resource_type",
		ConfigChanges: []Change{
			{
				Type:        ChangeTypeRename,
				Path:        "resource_type",
				OldValue:    "old_resource_type",
				NewValue:    "new_resource_type",
				Description: "Resource type changed",
			},
		},
		Impact: ImpactHigh,
	}

	// Verify the preview contains both old and new type
	assert.Equal(t, "old_resource_type", preview.Type)
	assert.Equal(t, "new_resource_type", preview.NewType)
	assert.Len(t, preview.ConfigChanges, 1)
	assert.Equal(t, ChangeTypeRename, preview.ConfigChanges[0].Type)
}

func TestPreviewGenerator_EmptyResources(t *testing.T) {
	registry := newMockPreviewRegistry()
	pg := NewPreviewGenerator(registry)

	// Generate preview with no blocks
	preview, err := pg.GeneratePreview([]*hclwrite.Block{}, "v4", "v5")
	require.NoError(t, err)
	require.NotNil(t, preview)

	assert.Equal(t, 0, preview.Summary.TotalResources)
	assert.Equal(t, 0, preview.Summary.ChangedResources)
	assert.Equal(t, 0, preview.Summary.TotalChanges)
	assert.Empty(t, preview.Resources)
}

func TestPreviewGenerator_NonResourceBlocks(t *testing.T) {
	registry := newMockPreviewRegistry()
	pg := NewPreviewGenerator(registry)

	// Create blocks that aren't resources
	file, _ := hclwrite.ParseConfig([]byte(`
		data "cloudflare_zones" "example" {
			filter {
				name = "example.com"
			}
		}
		
		locals {
			zone_id = data.cloudflare_zones.example.zones[0].id
		}
	`), "test.tf", hcl.InitialPos)

	blocks := file.Body().Blocks()

	// Generate preview - should skip non-resource blocks
	preview, err := pg.GeneratePreview(blocks, "v4", "v5")
	require.NoError(t, err)
	require.NotNil(t, preview)

	assert.Equal(t, 0, preview.Summary.TotalResources)
	assert.Empty(t, preview.Resources)
}