package migrations

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

func TestArgoMigrationConfig(t *testing.T) {
	// Test that the YAML config loads correctly
	migration, err := NewMigration(v4ToV5Config)
	require.NoError(t, err, "Failed to load migration config")
	
	assert.Equal(t, "cloudflare_argo", migration.ResourceType())
	assert.Equal(t, "v4", migration.SourceVersion())
	assert.Equal(t, "v5", migration.TargetVersion())
	
	// Parse the YAML directly to test the extended configuration
	var extConfig internal.ExtendedMigrationConfig
	err = yaml.Unmarshal(v4ToV5Config, &extConfig)
	require.NoError(t, err, "Failed to parse extended config")
	
	assert.Contains(t, extConfig.Description, "Split cloudflare_argo")
	
	// Verify resource splits configuration
	assert.Len(t, extConfig.Config.ResourceSplits, 1)
	split := extConfig.Config.ResourceSplits[0]
	
	assert.Equal(t, "cloudflare_argo", split.SourceResource)
	assert.Len(t, split.Splits, 2)
	assert.NotNil(t, split.Fallback)
	assert.True(t, split.GenerateMovedBlocks)
	assert.True(t, split.RemoveOriginal)
	
	// Check first split rule (smart_routing)
	assert.Equal(t, "smart_routing", split.Splits[0].WhenAttributeExists)
	assert.Equal(t, "cloudflare_argo_smart_routing", split.Splits[0].CreateResource)
	assert.Equal(t, "value", split.Splits[0].AttributeMappings["smart_routing"])
	assert.Contains(t, split.Splits[0].CopyAttributes, "zone_id")
	
	// Check second split rule (tiered_caching)
	assert.Equal(t, "tiered_caching", split.Splits[1].WhenAttributeExists)
	assert.Equal(t, "cloudflare_tiered_cache", split.Splits[1].CreateResource)
	assert.Equal(t, "cache_type", split.Splits[1].AttributeMappings["tiered_caching"])
	assert.Contains(t, split.Splits[1].CopyAttributes, "zone_id")
	assert.Equal(t, "_tiered", split.Splits[1].ResourceNameSuffix)
	
	// Check fallback rule
	assert.Equal(t, "cloudflare_argo_smart_routing", split.Fallback.CreateResource)
	assert.Equal(t, "off", split.Fallback.SetAttributes["value"])
	assert.Contains(t, split.Fallback.CopyAttributes, "zone_id")
}

func TestArgoFileTransformation(t *testing.T) {
	migration, err := NewMigration(v4ToV5Config)
	require.NoError(t, err)
	
	// Cast to ArgoMigration to access TransformFile
	argoMigration, ok := migration.(*ArgoMigration)
	require.True(t, ok, "Expected ArgoMigration type")
	
	tests := []struct {
		name           string
		input          string
		expectedBlocks int
		validateFunc   func(t *testing.T, blocks []*hclwrite.Block)
	}{
		{
			name: "splits argo resource with both attributes",
			input: `resource "cloudflare_argo" "example" {
				zone_id = "abc123"
				smart_routing = "on"
				tiered_caching = "on"
			}`,
			expectedBlocks: 4, // 2 resources + 2 moved blocks
			validateFunc: func(t *testing.T, blocks []*hclwrite.Block) {
				// Check first resource is smart_routing
				assert.Equal(t, "resource", blocks[0].Type())
				assert.Equal(t, []string{"cloudflare_argo_smart_routing", "example"}, blocks[0].Labels())
				
				// Check second resource is tiered_cache
				assert.Equal(t, "resource", blocks[1].Type())
				assert.Equal(t, []string{"cloudflare_tiered_cache", "example_tiered"}, blocks[1].Labels())
				
				// Check moved blocks
				assert.Equal(t, "moved", blocks[2].Type())
				assert.Equal(t, "moved", blocks[3].Type())
			},
		},
		{
			name: "splits argo resource with smart_routing only",
			input: `resource "cloudflare_argo" "example" {
				zone_id = "abc123"
				smart_routing = "on"
			}`,
			expectedBlocks: 2, // 1 resource + 1 moved block
			validateFunc: func(t *testing.T, blocks []*hclwrite.Block) {
				// Check resource is smart_routing
				assert.Equal(t, "resource", blocks[0].Type())
				assert.Equal(t, []string{"cloudflare_argo_smart_routing", "example"}, blocks[0].Labels())
				
				// Check moved block
				assert.Equal(t, "moved", blocks[1].Type())
			},
		},
		{
			name: "splits argo resource with tiered_caching only",
			input: `resource "cloudflare_argo" "example" {
				zone_id = "abc123"
				tiered_caching = "on"
			}`,
			expectedBlocks: 2, // 1 resource + 1 moved block
			validateFunc: func(t *testing.T, blocks []*hclwrite.Block) {
				// Check resource is tiered_cache
				assert.Equal(t, "resource", blocks[0].Type())
				assert.Equal(t, []string{"cloudflare_tiered_cache", "example"}, blocks[0].Labels())
				
				// Check moved block
				assert.Equal(t, "moved", blocks[1].Type())
			},
		},
		{
			name: "creates default smart_routing with no attributes",
			input: `resource "cloudflare_argo" "example" {
				zone_id = "abc123"
			}`,
			expectedBlocks: 2, // 1 resource + 1 moved block
			validateFunc: func(t *testing.T, blocks []*hclwrite.Block) {
				// Check resource is smart_routing with default value
				assert.Equal(t, "resource", blocks[0].Type())
				assert.Equal(t, []string{"cloudflare_argo_smart_routing", "example"}, blocks[0].Labels())
				
				// Check moved block
				assert.Equal(t, "moved", blocks[1].Type())
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the HCL
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())
			
			// Run file transformation
			err := argoMigration.TransformFile(file)
			assert.NoError(t, err)
			
			// Check results
			blocks := file.Body().Blocks()
			assert.Len(t, blocks, tt.expectedBlocks)
			
			if tt.validateFunc != nil {
				tt.validateFunc(t, blocks)
			}
		})
	}
}

func TestArgoStateTransformation(t *testing.T) {
	migration, err := NewMigration(v4ToV5Config)
	require.NoError(t, err)
	
	tests := []struct {
		name          string
		inputState    map[string]interface{}
		expectedType  string
		expectedAttrs map[string]interface{}
	}{
		{
			name: "smart_routing only",
			inputState: map[string]interface{}{
				"type": "cloudflare_argo",
				"attributes": map[string]interface{}{
					"zone_id":       "abc123",
					"smart_routing": "on",
				},
			},
			expectedType: "cloudflare_argo",  // Type change would need custom logic
			expectedAttrs: map[string]interface{}{
				"zone_id": "abc123",
				"value":   "on",  // Renamed from smart_routing
			},
		},
		{
			name: "tiered_caching only",
			inputState: map[string]interface{}{
				"type": "cloudflare_argo",
				"attributes": map[string]interface{}{
					"zone_id":        "abc123",
					"tiered_caching": "on",
				},
			},
			expectedType: "cloudflare_argo",
			expectedAttrs: map[string]interface{}{
				"zone_id":    "abc123",
				"cache_type": "on",  // Renamed from tiered_caching
			},
		},
		{
			name: "no attributes",
			inputState: map[string]interface{}{
				"type": "cloudflare_argo",
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
				},
			},
			expectedType: "cloudflare_argo",
			expectedAttrs: map[string]interface{}{
				"zone_id": "abc123",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the state to avoid mutation
			state := make(map[string]interface{})
			for k, v := range tt.inputState {
				if k == "attributes" {
					attrs := make(map[string]interface{})
					for ak, av := range v.(map[string]interface{}) {
						attrs[ak] = av
					}
					state[k] = attrs
				} else {
					state[k] = v
				}
			}
			
			ctx := internal.NewMigrationContext()
			
			// Run state migration
			err := migration.MigrateState(state, ctx)
			assert.NoError(t, err)
			
			// Check type (note: actual type change requires custom logic)
			assert.Equal(t, tt.expectedType, state["type"])
			
			// Check attributes based on the renames defined in YAML
			attrs := state["attributes"].(map[string]interface{})
			
			// The base migration should handle attribute renames
			// smart_routing -> value, tiered_caching -> cache_type
			if tt.inputState["attributes"].(map[string]interface{})["smart_routing"] != nil {
				assert.Equal(t, tt.expectedAttrs["value"], attrs["value"])
				assert.Nil(t, attrs["smart_routing"])
			}
			
			if tt.inputState["attributes"].(map[string]interface{})["tiered_caching"] != nil {
				assert.Equal(t, tt.expectedAttrs["cache_type"], attrs["cache_type"])
				assert.Nil(t, attrs["tiered_caching"])
			}
			
			// Zone ID should always be preserved
			assert.Equal(t, tt.expectedAttrs["zone_id"], attrs["zone_id"])
		})
	}
}

func TestResourceSplitConfiguration(t *testing.T) {
	// Test the resource split configuration structure
	var extConfig internal.ExtendedMigrationConfig
	err := yaml.Unmarshal(v4ToV5Config, &extConfig)
	require.NoError(t, err)
	
	// Verify the configuration can be used with the ResourceSplitter
	split := extConfig.Config.ResourceSplits[0]
	
	// Test smart routing split rule
	smartRoutingRule := split.Splits[0]
	assert.Equal(t, "smart_routing", smartRoutingRule.WhenAttributeExists)
	assert.Equal(t, "cloudflare_argo_smart_routing", smartRoutingRule.CreateResource)
	assert.Contains(t, smartRoutingRule.AttributeMappings, "smart_routing")
	assert.Equal(t, "value", smartRoutingRule.AttributeMappings["smart_routing"])
	assert.Contains(t, smartRoutingRule.CopyAttributes, "zone_id")
	assert.Equal(t, "", smartRoutingRule.ResourceNameSuffix)
	
	// Test tiered caching split rule
	tieredCachingRule := split.Splits[1]
	assert.Equal(t, "tiered_caching", tieredCachingRule.WhenAttributeExists)
	assert.Equal(t, "cloudflare_tiered_cache", tieredCachingRule.CreateResource)
	assert.Contains(t, tieredCachingRule.AttributeMappings, "tiered_caching")
	assert.Equal(t, "cache_type", tieredCachingRule.AttributeMappings["tiered_caching"])
	assert.Contains(t, tieredCachingRule.CopyAttributes, "zone_id")
	assert.Equal(t, "_tiered", tieredCachingRule.ResourceNameSuffix)
	
	// Test fallback configuration
	assert.NotNil(t, split.Fallback)
	assert.Equal(t, "cloudflare_argo_smart_routing", split.Fallback.CreateResource)
	assert.Contains(t, split.Fallback.SetAttributes, "value")
	assert.Equal(t, "off", split.Fallback.SetAttributes["value"])
	assert.Contains(t, split.Fallback.CopyAttributes, "zone_id")
	
	// Test flags
	assert.True(t, split.GenerateMovedBlocks)
	assert.True(t, split.RemoveOriginal)
}

func TestValueMappings(t *testing.T) {
	var extConfig internal.ExtendedMigrationConfig
	err := yaml.Unmarshal(v4ToV5Config, &extConfig)
	require.NoError(t, err)
	
	// Check value mappings configuration
	assert.NotEmpty(t, extConfig.Config.ValueMappings)
	
	// Find the cache_type mapping
	var cacheTypeMapping *internal.ValueMapping
	for i, mapping := range extConfig.Config.ValueMappings {
		if mapping.Attribute == "cache_type" {
			cacheTypeMapping = &extConfig.Config.ValueMappings[i]
			break
		}
	}
	
	require.NotNil(t, cacheTypeMapping, "cache_type value mapping not found")
	assert.Equal(t, "smart", cacheTypeMapping.Mappings["on"])
	assert.Equal(t, "off", cacheTypeMapping.Mappings["off"])
}

// NewMigration creates a migration from the embedded YAML for testing
func NewMigration(yamlContent []byte) (internal.ResourceMigration, error) {
	baseMigration, err := internal.NewMigration(yamlContent)
	if err != nil {
		return nil, err
	}
	
	// Return the ArgoMigration wrapper to include custom behavior
	// This is defined in v4_to_v5.go
	return &ArgoMigration{
		Migration: baseMigration,
	}, nil
}