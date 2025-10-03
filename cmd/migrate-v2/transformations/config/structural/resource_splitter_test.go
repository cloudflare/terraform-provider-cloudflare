package structural

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

func TestResourceSplitter(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		split          basic.ResourceSplit
		checkFunc      func(t *testing.T, ctx *basic.TransformContext, output string)
	}{
		{
			name: "split resource based on attribute existence",
			input: `resource "cloudflare_firewall" "example" {
  zone_id = "abc123"
  expression = "ip.src eq 1.2.3.4"
  action = "allow"
  waf_settings = {
    enabled = true
    rules = ["100015"]
  }
}`,
			split: basic.ResourceSplit{
				SourceResource: "cloudflare_firewall",
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "waf_settings",
						CreateResource:      "cloudflare_waf_rule",
						AttributeMappings: map[string]string{
							"waf_settings": "settings",
						},
						CopyAttributes: []string{"zone_id"},
						ResourceNameSuffix: "waf",
					},
				},
			},
			checkFunc: func(t *testing.T, ctx *basic.TransformContext, output string) {
				// Should have created a new block
				assert.NotNil(t, ctx.NewBlocks)
				assert.Len(t, ctx.NewBlocks, 1)
				
				// Check new block
				newBlock := ctx.NewBlocks[0]
				assert.Equal(t, "resource", newBlock.Type())
				labels := newBlock.Labels()
				assert.Equal(t, "cloudflare_waf_rule", labels[0])
				assert.Equal(t, "example_waf", labels[1])
				
				// Original should not have waf_settings anymore
				assert.NotContains(t, output, "waf_settings")
			},
		},
		{
			name: "split with fallback rule",
			input: `resource "cloudflare_access" "example" {
  zone_id = "abc123"
  name = "example"
}`,
			split: basic.ResourceSplit{
				SourceResource: "cloudflare_access",
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "policy_settings",
						CreateResource:      "cloudflare_access_policy",
						ResourceNameSuffix:  "policy",
					},
				},
				Fallback: &basic.SplitRule{
					ChangeResourceType: "cloudflare_zero_trust_access",
					SetAttributes: map[string]interface{}{
						"type": "default",
					},
				},
			},
			checkFunc: func(t *testing.T, ctx *basic.TransformContext, output string) {
				// Fallback should change resource type
				assert.NotNil(t, ctx.ResourceTypeChanges)
				assert.Len(t, ctx.ResourceTypeChanges, 1)
				
				// Should have set default type
				assert.Contains(t, output, `type    = "default"`)
			},
		},
		{
			name: "split with attribute mappings and new attributes",
			input: `resource "cloudflare_combined" "example" {
  zone_id = "abc123"
  name = "combined_resource"
  feature_a = "value_a"
  feature_b = "value_b"
}`,
			split: basic.ResourceSplit{
				SourceResource: "cloudflare_combined",
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "feature_b",
						CreateResource:      "cloudflare_feature_b",
						AttributeMappings: map[string]string{
							"feature_b": "configuration",
						},
						SetAttributes: map[string]interface{}{
							"enabled": true,
							"version": 2,
						},
						CopyAttributes: []string{"zone_id", "name"},
						ResourceNameSuffix: "feature_b",
					},
				},
				RemoveOriginal: true,
			},
			checkFunc: func(t *testing.T, ctx *basic.TransformContext, output string) {
				// Should create new block
				assert.Len(t, ctx.NewBlocks, 1)
				
				// Should mark original for removal
				assert.NotNil(t, ctx.BlocksToRemove)
				assert.Len(t, ctx.BlocksToRemove, 1)
				
				// New block should have mapped and set attributes
				newBlock := ctx.NewBlocks[0]
				body := newBlock.Body()
				
				// Check for set attributes
				enabledAttr := body.GetAttribute("enabled")
				assert.NotNil(t, enabledAttr)
				
				versionAttr := body.GetAttribute("version")
				assert.NotNil(t, versionAttr)
			},
		},
		{
			name: "split with moved block generation",
			input: `resource "cloudflare_old_resource" "example" {
  zone_id = "abc123"
  settings = "config"
}`,
			split: basic.ResourceSplit{
				SourceResource: "cloudflare_old_resource",
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "settings",
						CreateResource:      "cloudflare_new_resource",
						AttributeMappings: map[string]string{
							"settings": "configuration",
						},
						ResourceNameSuffix: "new",
					},
				},
				GenerateMovedBlocks: true,
			},
			checkFunc: func(t *testing.T, ctx *basic.TransformContext, output string) {
				// Should generate moved blocks
				assert.NotNil(t, ctx.MovedBlocks)
				assert.Len(t, ctx.MovedBlocks, 1)
				
				// Check moved block mapping
				from := "cloudflare_old_resource.example"
				to := "cloudflare_new_resource.example_new"
				assert.Equal(t, to, ctx.MovedBlocks[from])
			},
		},
		{
			name: "multiple split rules",
			input: `resource "cloudflare_monolith" "example" {
  zone_id = "abc123"
  feature_a = "config_a"
  feature_b = "config_b"
  common = "shared"
}`,
			split: basic.ResourceSplit{
				SourceResource: "cloudflare_monolith",
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "feature_a",
						CreateResource:      "cloudflare_feature_a",
						AttributeMappings: map[string]string{
							"feature_a": "config",
						},
						CopyAttributes: []string{"zone_id", "common"},
						ResourceNameSuffix: "a",
					},
					{
						WhenAttributeExists: "feature_b",
						CreateResource:      "cloudflare_feature_b",
						AttributeMappings: map[string]string{
							"feature_b": "config",
						},
						CopyAttributes: []string{"zone_id", "common"},
						ResourceNameSuffix: "b",
					},
				},
			},
			checkFunc: func(t *testing.T, ctx *basic.TransformContext, output string) {
				// Should create two new blocks
				assert.Len(t, ctx.NewBlocks, 2)
				
				// Check resource types
				block1 := ctx.NewBlocks[0]
				assert.Equal(t, "cloudflare_feature_a", block1.Labels()[0])
				
				block2 := ctx.NewBlocks[1]
				assert.Equal(t, "cloudflare_feature_b", block2.Labels()[0])
				
				// Original should not have feature_a or feature_b
				assert.NotContains(t, output, "feature_a")
				assert.NotContains(t, output, "feature_b")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// Get the first block
			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Create context and apply transformation
			ctx := &basic.TransformContext{}
			splitter := ResourceSplitter(tt.split)
			
			err := splitter(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, ctx, output)
			}
		})
	}
}

func TestResourceSplitterForState(t *testing.T) {
	tests := []struct {
		name          string
		state         map[string]interface{}
		split         basic.ResourceSplit
		checkFunc     func(t *testing.T, state map[string]interface{})
	}{
		{
			name: "split state based on attribute",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"name": "example",
					"waf_settings": map[string]interface{}{
						"enabled": true,
						"rules": []string{"100015"},
					},
				},
			},
			split: basic.ResourceSplit{
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "waf_settings",
						CreateResource:      "cloudflare_waf_rule",
						AttributeMappings: map[string]string{
							"waf_settings": "settings",
						},
						CopyAttributes: []string{"zone_id"},
						ResourceNameSuffix: "waf",
					},
				},
			},
			checkFunc: func(t *testing.T, state map[string]interface{}) {
				// Should have split resources metadata
				splitResources, ok := state["_split_resources"].([]map[string]interface{})
				assert.True(t, ok)
				assert.Len(t, splitResources, 1)
				
				// Check split resource
				splitResource := splitResources[0]
				attrs := splitResource["attributes"].(map[string]interface{})
				assert.Equal(t, "abc123", attrs["zone_id"])
				assert.NotNil(t, attrs["settings"])
				
				// Original should not have waf_settings
				originalAttrs := state["attributes"].(map[string]interface{})
				assert.Nil(t, originalAttrs["waf_settings"])
			},
		},
		{
			name: "fallback rule in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"name": "example",
				},
			},
			split: basic.ResourceSplit{
				Splits: []basic.SplitRule{
					{
						WhenAttributeExists: "nonexistent",
						CreateResource:      "cloudflare_new",
					},
				},
				Fallback: &basic.SplitRule{
					SetAttributes: map[string]interface{}{
						"type": "default",
						"version": 1,
					},
				},
			},
			checkFunc: func(t *testing.T, state map[string]interface{}) {
				attrs := state["attributes"].(map[string]interface{})
				
				// Fallback should have added attributes
				assert.Equal(t, "default", attrs["type"])
				assert.Equal(t, 1, attrs["version"])
			},
		},
		{
			name: "attribute mappings in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"old_setting": "value",
				},
			},
			split: basic.ResourceSplit{
				Fallback: &basic.SplitRule{
					AttributeMappings: map[string]string{
						"old_setting": "new_setting",
					},
				},
			},
			checkFunc: func(t *testing.T, state map[string]interface{}) {
				attrs := state["attributes"].(map[string]interface{})
				
				// Should have renamed attribute
				assert.Equal(t, "value", attrs["new_setting"])
				assert.Nil(t, attrs["old_setting"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply state transformation
			splitter := ResourceSplitterForState(tt.split)
			err := splitter(tt.state)
			require.NoError(t, err)

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, tt.state)
			}
		})
	}
}

func TestResourceSplitterEdgeCases(t *testing.T) {
	t.Run("non-resource block", func(t *testing.T) {
		input := `data "cloudflare_zone" "example" { name = "example.com" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		split := basic.ResourceSplit{
			SourceResource: "cloudflare_zone",
		}
		
		ctx := &basic.TransformContext{}
		splitter := ResourceSplitter(split)
		err := splitter(block, ctx)
		
		assert.NoError(t, err)
		assert.Len(t, ctx.NewBlocks, 0) // Should not process data blocks
	})
	
	t.Run("resource type mismatch", func(t *testing.T) {
		input := `resource "cloudflare_other" "example" { zone_id = "abc" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		split := basic.ResourceSplit{
			SourceResource: "cloudflare_different",
			Splits: []basic.SplitRule{
				{CreateResource: "cloudflare_new"},
			},
		}
		
		ctx := &basic.TransformContext{}
		splitter := ResourceSplitter(split)
		err := splitter(block, ctx)
		
		assert.NoError(t, err)
		assert.Len(t, ctx.NewBlocks, 0) // Should not process mismatched resource types
	})
	
	t.Run("empty split rules", func(t *testing.T) {
		input := `resource "cloudflare_test" "example" { zone_id = "abc" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		split := basic.ResourceSplit{
			SourceResource: "cloudflare_test",
			Splits: []basic.SplitRule{},
		}
		
		ctx := &basic.TransformContext{}
		splitter := ResourceSplitter(split)
		err := splitter(block, ctx)
		
		assert.NoError(t, err)
	})
	
	t.Run("reference values in set attributes", func(t *testing.T) {
		input := `resource "cloudflare_test" "example" { 
  zone_id = "abc"
  trigger = true
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		split := basic.ResourceSplit{
			SourceResource: "cloudflare_test",
			Splits: []basic.SplitRule{
				{
					WhenAttributeExists: "trigger",
					CreateResource: "cloudflare_new",
					SetAttributes: map[string]interface{}{
						"reference": "var.my_variable",
						"data_ref": "data.cloudflare_zone.example.id",
					},
				},
			},
		}
		
		ctx := &basic.TransformContext{}
		splitter := ResourceSplitter(split)
		err := splitter(block, ctx)
		
		assert.NoError(t, err)
		assert.Len(t, ctx.NewBlocks, 1)
		
		// Check that references are set as traversals
		newBlock := ctx.NewBlocks[0]
		body := newBlock.Body()
		
		refAttr := body.GetAttribute("reference")
		assert.NotNil(t, refAttr)
		
		dataRefAttr := body.GetAttribute("data_ref")
		assert.NotNil(t, dataRefAttr)
	})
}