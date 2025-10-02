package structural

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

func TestListToBlocksConverter(t *testing.T) {
	tests := []struct {
		name     string
		hcl      string
		attrName string
		validate func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic list to blocks conversion",
			hcl: `resource "test" "example" {
  destinations = [
    {
      uri = "https://example.com"
      priority = 1
    },
    {
      uri = "https://backup.com"
      priority = 2
    }
  ]
}`,
			attrName: "destinations",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original list attribute should be removed
				assert.Nil(t, body.GetAttribute("destinations"))
				
				// Should have blocks instead
				blocks := body.Blocks()
				destinationBlocks := 0
				for _, b := range blocks {
					if b.Type() == "destinations" {
						destinationBlocks++
						// Check that block has attributes
						assert.NotNil(t, b.Body().GetAttribute("uri"))
						assert.NotNil(t, b.Body().GetAttribute("priority"))
					}
				}
				assert.Equal(t, 2, destinationBlocks, "Should have 2 destination blocks")
			},
		},
		{
			name: "single item list",
			hcl: `resource "test" "example" {
  rules = [
    {
      action = "allow"
      source = "10.0.0.0/8"
    }
  ]
}`,
			attrName: "rules",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("rules"))
				
				blocks := body.Blocks()
				ruleBlocks := 0
				for _, b := range blocks {
					if b.Type() == "rules" {
						ruleBlocks++
						assert.NotNil(t, b.Body().GetAttribute("action"))
						assert.NotNil(t, b.Body().GetAttribute("source"))
					}
				}
				assert.Equal(t, 1, ruleBlocks, "Should have 1 rule block")
			},
		},
		{
			name: "empty list",
			hcl: `resource "test" "example" {
  items = []
}`,
			attrName: "items",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Empty list should be removed
				assert.Nil(t, body.GetAttribute("items"))
				
				// No blocks should be created
				blocks := body.Blocks()
				for _, b := range blocks {
					assert.NotEqual(t, "items", b.Type())
				}
			},
		},
		{
			name: "list with simple values",
			hcl: `resource "test" "example" {
  tags = ["tag1", "tag2", "tag3"]
}`,
			attrName: "tags",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Simple list should be converted to blocks with value attribute
				assert.Nil(t, body.GetAttribute("tags"))
				
				blocks := body.Blocks()
				tagBlocks := 0
				for _, b := range blocks {
					if b.Type() == "tags" {
						tagBlocks++
						// Simple values should be wrapped in a value attribute
						assert.NotNil(t, b.Body().GetAttribute("value"))
					}
				}
				assert.Equal(t, 3, tagBlocks, "Should have 3 tag blocks")
			},
		},
		{
			name: "non-existent attribute",
			hcl: `resource "test" "example" {
  other = "value"
}`,
			attrName: "missing",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should not affect other attributes
				assert.NotNil(t, body.GetAttribute("other"))
				
				// No blocks should be created
				blocks := body.Blocks()
				for _, b := range blocks {
					assert.NotEqual(t, "missing", b.Type())
				}
			},
		},
		{
			name: "list with nested objects",
			hcl: `resource "test" "example" {
  headers = [
    {
      name = "Content-Type"
      value = "application/json"
      metadata = {
        required = true
        description = "JSON content"
      }
    },
    {
      name = "Authorization"
      value = "Bearer token"
      metadata = {
        required = false
        description = "Auth header"
      }
    }
  ]
}`,
			attrName: "headers",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("headers"))
				
				blocks := body.Blocks()
				headerBlocks := 0
				for _, b := range blocks {
					if b.Type() == "headers" {
						headerBlocks++
						assert.NotNil(t, b.Body().GetAttribute("name"))
						assert.NotNil(t, b.Body().GetAttribute("value"))
						assert.NotNil(t, b.Body().GetAttribute("metadata"))
					}
				}
				assert.Equal(t, 2, headerBlocks, "Should have 2 header blocks")
			},
		},
		{
			name: "list with mixed types",
			hcl: `resource "test" "example" {
  configs = [
    {
      type = "string"
      value = "text"
    },
    {
      type = "number"
      value = 42
    },
    {
      type = "bool"
      value = true
    }
  ]
}`,
			attrName: "configs",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("configs"))
				
				blocks := body.Blocks()
				configBlocks := 0
				for _, b := range blocks {
					if b.Type() == "configs" {
						configBlocks++
						assert.NotNil(t, b.Body().GetAttribute("type"))
						assert.NotNil(t, b.Body().GetAttribute("value"))
					}
				}
				assert.Equal(t, 3, configBlocks, "Should have 3 config blocks")
			},
		},
		{
			name: "attribute is not a list",
			hcl: `resource "test" "example" {
  single = {
    field = "value"
  }
}`,
			attrName: "single",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Non-list attribute should remain unchanged
				assert.NotNil(t, body.GetAttribute("single"))
				
				// No blocks should be created
				blocks := body.Blocks()
				for _, b := range blocks {
					assert.NotEqual(t, "single", b.Type())
				}
			},
		},
		{
			name: "list with references",
			hcl: `resource "test" "example" {
  policies = [
    {
      id = aws_iam_policy.read.arn
      effect = "Allow"
    },
    {
      id = aws_iam_policy.write.arn
      effect = "Allow"
    }
  ]
}`,
			attrName: "policies",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("policies"))
				
				blocks := body.Blocks()
				policyBlocks := 0
				for _, b := range blocks {
					if b.Type() == "policies" {
						policyBlocks++
						assert.NotNil(t, b.Body().GetAttribute("id"))
						assert.NotNil(t, b.Body().GetAttribute("effect"))
					}
				}
				assert.Equal(t, 2, policyBlocks, "Should have 2 policy blocks")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := ListToBlocksConverter(tt.attrName)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestListToBlocksConverter_ComplexScenarios(t *testing.T) {
	t.Run("preserves existing blocks", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  existing_block {
    field = "value"
  }
  
  items = [
    {
      name = "item1"
    }
  ]
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}

		// Count existing blocks before transformation
		existingCount := 0
		for _, b := range block.Body().Blocks() {
			if b.Type() == "existing_block" {
				existingCount++
			}
		}
		assert.Equal(t, 1, existingCount)

		transformer := ListToBlocksConverter("items")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Existing blocks should still be there
		existingAfter := 0
		itemBlocks := 0
		for _, b := range block.Body().Blocks() {
			if b.Type() == "existing_block" {
				existingAfter++
			}
			if b.Type() == "items" {
				itemBlocks++
			}
		}
		assert.Equal(t, 1, existingAfter, "Existing block should be preserved")
		assert.Equal(t, 1, itemBlocks, "Item block should be created")
	})

	t.Run("handles dynamic expressions", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  rules = var.firewall_rules
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}

		transformer := ListToBlocksConverter("rules")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Dynamic expression should remain as-is since we can't parse it
		assert.NotNil(t, block.Body().GetAttribute("rules"))
	})
}

func TestListToBlocksWithMapping(t *testing.T) {
	tests := []struct {
		name     string
		hcl      string
		attrName string
		mapping  map[string]string
		validate func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic field mapping",
			hcl: `resource "test" "example" {
  items = [
    { old_field = "value1", another = "data1" },
    { old_field = "value2", another = "data2" }
  ]
}`,
			attrName: "items",
			mapping: map[string]string{
				"old_field": "new_field",
				"another":   "renamed",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original list should be removed
				assert.Nil(t, body.GetAttribute("items"))
				
				// Count blocks
				blocks := body.Blocks()
				itemBlocks := 0
				for _, b := range blocks {
					if b.Type() == "items" {
						itemBlocks++
						// Check renamed fields
						assert.NotNil(t, b.Body().GetAttribute("new_field"))
						assert.NotNil(t, b.Body().GetAttribute("renamed"))
						// Original field names should not exist
						assert.Nil(t, b.Body().GetAttribute("old_field"))
						assert.Nil(t, b.Body().GetAttribute("another"))
					}
				}
				assert.Equal(t, 2, itemBlocks)
			},
		},
		{
			name: "partial field mapping",
			hcl: `resource "test" "example" {
  configs = [
    { host = "localhost", port = 8080, enabled = true }
  ]
}`,
			attrName: "configs",
			mapping: map[string]string{
				"host": "hostname",
				// port and enabled are not mapped
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				blocks := body.Blocks()
				for _, b := range blocks {
					if b.Type() == "configs" {
						// Mapped field
						assert.NotNil(t, b.Body().GetAttribute("hostname"))
						assert.Nil(t, b.Body().GetAttribute("host"))
						// Unmapped fields keep original names
						assert.NotNil(t, b.Body().GetAttribute("port"))
						assert.NotNil(t, b.Body().GetAttribute("enabled"))
					}
				}
			},
		},
		{
			name: "empty mapping",
			hcl: `resource "test" "example" {
  rules = [
    { action = "allow", priority = 1 }
  ]
}`,
			attrName: "rules",
			mapping:  map[string]string{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				blocks := body.Blocks()
				for _, b := range blocks {
					if b.Type() == "rules" {
						// All fields keep original names
						assert.NotNil(t, b.Body().GetAttribute("action"))
						assert.NotNil(t, b.Body().GetAttribute("priority"))
					}
				}
			},
		},
		{
			name: "nil mapping",
			hcl: `resource "test" "example" {
  items = [{ field = "value" }]
}`,
			attrName: "items",
			mapping:  nil,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				blocks := body.Blocks()
				for _, b := range blocks {
					if b.Type() == "items" {
						// Field keeps original name
						assert.NotNil(t, b.Body().GetAttribute("field"))
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := ListToBlocksWithMapping(tt.attrName, tt.mapping)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestListToBlocksConverter_EdgeCases(t *testing.T) {
	t.Run("handles malformed list", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  broken = [
    {
      field = 
    }
  ]
}`
		// This will fail parsing, not in the transformer
		_, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		assert.True(t, diags.HasErrors())
	})

	t.Run("empty attribute name", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  items = []
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}

		transformer := ListToBlocksConverter("")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Should not crash, but also shouldn't do anything
		assert.NotNil(t, block.Body().GetAttribute("items"))
	})
}

func BenchmarkListToBlocksConverter(b *testing.B) {
	hclContent := `resource "test" "example" {
  destinations = [
    {
      uri = "https://example1.com"
      priority = 1
    },
    {
      uri = "https://example2.com"
      priority = 2
    },
    {
      uri = "https://example3.com"
      priority = 3
    },
    {
      uri = "https://example4.com"
      priority = 4
    },
    {
      uri = "https://example5.com"
      priority = 5
    }
  ]
}`

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}
		
		transformer := ListToBlocksConverter("destinations")
		_ = transformer(block, ctx)
	}
}

func BenchmarkListToBlocksConverter_LargeList(b *testing.B) {
	// Build HCL with many list items
	hclContent := `resource "test" "example" {
  items = [`
	
	for i := 0; i < 50; i++ {
		if i > 0 {
			hclContent += ","
		}
		hclContent += `
    {
      id = "item` + string(rune(i)) + `"
      value = ` + string(rune(i)) + `
    }`
	}
	
	hclContent += `
  ]
}`

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}
		
		transformer := ListToBlocksConverter("items")
		_ = transformer(block, ctx)
	}
}