package basic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttributeRemover_ComplexCases(t *testing.T) {
	tests := []struct {
		name          string
		hcl           string
		removeAttrs   []string
		shouldExist   []string
		shouldNotExist []string
	}{
		{
			name: "remove multiple attributes",
			hcl: `resource "test" "example" {
  remove1 = "value1"
  remove2 = ["list", "values"]
  remove3 = {
    key = "value"
  }
  keep1 = "keep_value1"
  keep2 = true
}`,
			removeAttrs:    []string{"remove1", "remove2", "remove3"},
			shouldExist:    []string{"keep1", "keep2"},
			shouldNotExist: []string{"remove1", "remove2", "remove3"},
		},
		{
			name: "remove non-existent attributes",
			hcl: `resource "test" "example" {
  existing1 = "value1"
  existing2 = "value2"
}`,
			removeAttrs:    []string{"non_existent1", "non_existent2", "non_existent3"},
			shouldExist:    []string{"existing1", "existing2"},
			shouldNotExist: []string{"non_existent1", "non_existent2", "non_existent3"},
		},
		{
			name: "remove all attributes",
			hcl: `resource "test" "example" {
  attr1 = "value1"
  attr2 = "value2"
  attr3 = "value3"
}`,
			removeAttrs:    []string{"attr1", "attr2", "attr3"},
			shouldExist:    []string{},
			shouldNotExist: []string{"attr1", "attr2", "attr3"},
		},
		{
			name: "empty removal list",
			hcl: `resource "test" "example" {
  attr1 = "value1"
  attr2 = "value2"
}`,
			removeAttrs:    []string{},
			shouldExist:    []string{"attr1", "attr2"},
			shouldNotExist: []string{},
		},
		{
			name: "remove attributes with special names",
			hcl: `resource "test" "example" {
  under_score = "value1"
  dash-name = "value2"
  number123 = "value3"
  CamelCase = "value4"
  keep = "value5"
}`,
			removeAttrs:    []string{"under_score", "dash-name", "number123", "CamelCase"},
			shouldExist:    []string{"keep"},
			shouldNotExist: []string{"under_score", "dash-name", "number123", "CamelCase"},
		},
		{
			name: "remove duplicate entries in list",
			hcl: `resource "test" "example" {
  attr1 = "value1"
  attr2 = "value2"
  attr3 = "value3"
}`,
			removeAttrs:    []string{"attr1", "attr1", "attr2", "attr2"}, // Duplicates
			shouldExist:    []string{"attr3"},
			shouldNotExist: []string{"attr1", "attr2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}

			transformer := AttributeRemover(tt.removeAttrs...)
			err := transformer(block, ctx)
			require.NoError(t, err)

			body := block.Body()
			
			// Check attributes that should exist
			for _, attr := range tt.shouldExist {
				assert.NotNil(t, body.GetAttribute(attr), "Expected attribute %s to exist", attr)
			}
			
			// Check attributes that should not exist
			for _, attr := range tt.shouldNotExist {
				assert.Nil(t, body.GetAttribute(attr), "Expected attribute %s to not exist", attr)
			}
		})
	}
}

func TestConditionalRemover_ComplexConditions(t *testing.T) {
	tests := []struct {
		name         string
		hcl          string
		attribute    string
		condition    func(*hclwrite.Block) bool
		shouldRemove bool
		description  string
	}{
		{
			name: "condition based on attribute value",
			hcl: `resource "test" "example" {
  type = "special"
  conditional = "remove_me"
  other = "keep"
}`,
			attribute: "conditional",
			condition: func(block *hclwrite.Block) bool {
				attr := block.Body().GetAttribute("type")
				if attr == nil {
					return false
				}
				value := string(attr.Expr().BuildTokens(nil).Bytes())
				return strings.Contains(value, "special")
			},
			shouldRemove: true,
			description:  "Should remove when type contains 'special'",
		},
		{
			name: "condition based on multiple attributes",
			hcl: `resource "test" "example" {
  env = "production"
  region = "us-west-2"
  debug = true
}`,
			attribute: "debug",
			condition: func(block *hclwrite.Block) bool {
				envAttr := block.Body().GetAttribute("env")
				regionAttr := block.Body().GetAttribute("region")
				
				if envAttr == nil || regionAttr == nil {
					return false
				}
				
				envValue := string(envAttr.Expr().BuildTokens(nil).Bytes())
				regionValue := string(regionAttr.Expr().BuildTokens(nil).Bytes())
				
				// Remove debug in production US regions
				return strings.Contains(envValue, "production") && 
				       strings.Contains(regionValue, "us-")
			},
			shouldRemove: true,
			description:  "Should remove debug in production US regions",
		},
		{
			name: "condition with complex logic",
			hcl: `resource "test" "example" {
  version = "2.0"
  legacy_field = "old_value"
  new_field = "new_value"
}`,
			attribute: "legacy_field",
			condition: func(block *hclwrite.Block) bool {
				versionAttr := block.Body().GetAttribute("version")
				newFieldAttr := block.Body().GetAttribute("new_field")
				
				if versionAttr == nil {
					return false
				}
				
				versionValue := string(versionAttr.Expr().BuildTokens(nil).Bytes())
				
				// Remove legacy field if version >= 2.0 AND new_field exists
				hasHighVersion := strings.Contains(versionValue, "2.") || 
				                 strings.Contains(versionValue, "3.")
				hasNewField := newFieldAttr != nil
				
				return hasHighVersion && hasNewField
			},
			shouldRemove: true,
			description:  "Should remove legacy field when version >= 2.0 and new field exists",
		},
		{
			name: "condition that always returns true",
			hcl: `resource "test" "example" {
  always_remove = "value"
  other = "keep"
}`,
			attribute: "always_remove",
			condition: func(block *hclwrite.Block) bool {
				return true
			},
			shouldRemove: true,
			description:  "Should always remove",
		},
		{
			name: "condition that always returns false",
			hcl: `resource "test" "example" {
  never_remove = "value"
  other = "keep"
}`,
			attribute: "never_remove",
			condition: func(block *hclwrite.Block) bool {
				return false
			},
			shouldRemove: false,
			description:  "Should never remove",
		},
		{
			name: "condition checking attribute count",
			hcl: `resource "test" "example" {
  attr1 = "value1"
  attr2 = "value2"
  attr3 = "value3"
  optional = "remove_if_many_attrs"
}`,
			attribute: "optional",
			condition: func(block *hclwrite.Block) bool {
				// Remove if there are more than 3 other attributes
				count := 0
				body := block.Body()
				// Check for specific attributes (simplified)
				if body.GetAttribute("attr1") != nil { count++ }
				if body.GetAttribute("attr2") != nil { count++ }
				if body.GetAttribute("attr3") != nil { count++ }
				return count >= 3
			},
			shouldRemove: true,
			description:  "Should remove when too many attributes",
		},
		{
			name: "condition with nil block body",
			hcl: `resource "test" "example" {
  target = "value"
}`,
			attribute: "target",
			condition: func(block *hclwrite.Block) bool {
				// Simulate checking a non-existent nested block
				// This should handle gracefully
				return block == nil // Will be false since block is not nil
			},
			shouldRemove: false,
			description:  "Should handle nil checks gracefully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors(), tt.description)

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}

			transformer := ConditionalRemover(tt.attribute, tt.condition)
			err := transformer(block, ctx)
			require.NoError(t, err)

			body := block.Body()
			if tt.shouldRemove {
				assert.Nil(t, body.GetAttribute(tt.attribute), tt.description)
			} else {
				assert.NotNil(t, body.GetAttribute(tt.attribute), tt.description)
			}
		})
	}
}

func TestAttributeRemover_EdgeCases(t *testing.T) {
	t.Run("nil varargs", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  attr = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		// Call with no arguments
		transformer := AttributeRemover()
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Attribute should remain
		assert.NotNil(t, block.Body().GetAttribute("attr"))
	})
	
	t.Run("remove same attribute multiple times", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  target = "value"
  keep = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		// Apply remover multiple times
		transformer := AttributeRemover("target")
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Apply again (should not error)
		err = transformer(block, ctx)
		require.NoError(t, err)
		
		// Target should be removed, keep should remain
		assert.Nil(t, block.Body().GetAttribute("target"))
		assert.NotNil(t, block.Body().GetAttribute("keep"))
	})
}

func TestConditionalRemover_PanicSafety(t *testing.T) {
	t.Run("nil condition function", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  attr = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		// This will panic if not handled properly
		transformer := ConditionalRemover("attr", nil)
		
		// Should panic when trying to execute
		assert.Panics(t, func() {
			_ = transformer(block, ctx)
		})
	})
	
	t.Run("condition function that panics", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  attr = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		panicCondition := func(block *hclwrite.Block) bool {
			panic("intentional panic in condition")
		}
		
		transformer := ConditionalRemover("attr", panicCondition)
		
		// Should panic
		assert.Panics(t, func() {
			_ = transformer(block, ctx)
		})
	})
}

func BenchmarkAttributeRemover(b *testing.B) {
	hclContent := `resource "test" "example" {
  remove1 = "value1"
  remove2 = "value2"
  remove3 = "value3"
  keep1 = "value4"
  keep2 = "value5"
}`
	
	removeAttrs := []string{"remove1", "remove2", "remove3"}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := AttributeRemover(removeAttrs...)
		_ = transformer(block, ctx)
	}
}

func BenchmarkConditionalRemover(b *testing.B) {
	hclContent := `resource "test" "example" {
  type = "special"
  conditional = "value"
  other = "keep"
}`
	
	condition := func(block *hclwrite.Block) bool {
		attr := block.Body().GetAttribute("type")
		if attr == nil {
			return false
		}
		value := string(attr.Expr().BuildTokens(nil).Bytes())
		return strings.Contains(value, "special")
	}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := ConditionalRemover("conditional", condition)
		_ = transformer(block, ctx)
	}
}

func BenchmarkAttributeRemover_ManyAttributes(b *testing.B) {
	// Create HCL with many attributes
	var hclBuilder strings.Builder
	hclBuilder.WriteString(`resource "test" "example" {`)
	for i := 0; i < 100; i++ {
		hclBuilder.WriteString(fmt.Sprintf("\n  attr%d = \"value%d\"", i, i))
	}
	hclBuilder.WriteString("\n}")
	
	hclContent := hclBuilder.String()
	
	// Remove first 50 attributes
	var removeAttrs []string
	for i := 0; i < 50; i++ {
		removeAttrs = append(removeAttrs, fmt.Sprintf("attr%d", i))
	}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := AttributeRemover(removeAttrs...)
		_ = transformer(block, ctx)
	}
}