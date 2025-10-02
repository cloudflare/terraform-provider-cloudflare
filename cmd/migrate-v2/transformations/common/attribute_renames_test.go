package common

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttributeRenamer_ComplexCases(t *testing.T) {
	tests := []struct {
		name     string
		hcl      string
		mappings map[string]string
		validate func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "rename with complex values",
			hcl: `resource "test" "example" {
  old_list = ["a", "b", "c"]
  old_map = {
    key1 = "value1"
    key2 = "value2"
  }
  old_bool = true
  old_number = 42
}`,
			mappings: map[string]string{
				"old_list":   "new_list",
				"old_map":    "new_map",
				"old_bool":   "new_bool",
				"old_number": "new_number",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Check new attributes exist
				assert.NotNil(t, body.GetAttribute("new_list"))
				assert.NotNil(t, body.GetAttribute("new_map"))
				assert.NotNil(t, body.GetAttribute("new_bool"))
				assert.NotNil(t, body.GetAttribute("new_number"))
				
				// Check old attributes don't exist
				assert.Nil(t, body.GetAttribute("old_list"))
				assert.Nil(t, body.GetAttribute("old_map"))
				assert.Nil(t, body.GetAttribute("old_bool"))
				assert.Nil(t, body.GetAttribute("old_number"))
				
				// Verify values are preserved
				listValue := string(body.GetAttribute("new_list").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, listValue, `["a", "b", "c"]`)
				
				mapValue := string(body.GetAttribute("new_map").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, mapValue, "key1")
				assert.Contains(t, mapValue, "value1")
			},
		},
		{
			name: "rename with references",
			hcl: `resource "test" "example" {
  old_ref = var.some_variable
  old_resource_ref = aws_instance.example.id
  old_local_ref = local.some_local
}`,
			mappings: map[string]string{
				"old_ref":          "new_ref",
				"old_resource_ref": "new_resource_ref",
				"old_local_ref":    "new_local_ref",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Check references are preserved
				refValue := string(body.GetAttribute("new_ref").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, refValue, "var.some_variable")
				
				resourceRefValue := string(body.GetAttribute("new_resource_ref").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, resourceRefValue, "aws_instance.example.id")
				
				localRefValue := string(body.GetAttribute("new_local_ref").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, localRefValue, "local.some_local")
			},
		},
		{
			name: "rename with functions",
			hcl: `resource "test" "example" {
  old_func = jsonencode({
    key = "value"
  })
  old_interpolation = "prefix-${var.name}-suffix"
}`,
			mappings: map[string]string{
				"old_func":          "new_func",
				"old_interpolation": "new_interpolation",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Check function calls are preserved
				funcValue := string(body.GetAttribute("new_func").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, funcValue, "jsonencode")
				
				interpValue := string(body.GetAttribute("new_interpolation").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, interpValue, "${var.name}")
			},
		},
		{
			name: "empty mappings",
			hcl: `resource "test" "example" {
  keep1 = "value1"
  keep2 = "value2"
}`,
			mappings: map[string]string{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// All attributes should remain unchanged
				assert.NotNil(t, body.GetAttribute("keep1"))
				assert.NotNil(t, body.GetAttribute("keep2"))
			},
		},
		{
			name: "overlapping renames (swap names)",
			hcl: `resource "test" "example" {
  first = "value1"
  second = "value2"
}`,
			mappings: map[string]string{
				"first":  "temp_first",
				"second": "first",
				// Note: This would need two passes to swap properly
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// This demonstrates order dependency in renames
				assert.NotNil(t, body.GetAttribute("temp_first"))
				assert.NotNil(t, body.GetAttribute("first"))
				assert.Nil(t, body.GetAttribute("second"))
			},
		},
		{
			name: "rename with multiline strings",
			hcl: `resource "test" "example" {
  old_multiline = <<EOF
line1
line2
line3
EOF
}`,
			mappings: map[string]string{
				"old_multiline": "new_multiline",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.NotNil(t, body.GetAttribute("new_multiline"))
				assert.Nil(t, body.GetAttribute("old_multiline"))
				
				// Check multiline content is preserved
				value := string(body.GetAttribute("new_multiline").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "line1")
				assert.Contains(t, value, "line2")
				assert.Contains(t, value, "line3")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}

			transformer := AttributeRenamer(tt.mappings)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestAttributeRenamer_EdgeCases(t *testing.T) {
	t.Run("nil mappings", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  attr = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := AttributeRenamer(nil)
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Should not panic and attribute should remain
		assert.NotNil(t, block.Body().GetAttribute("attr"))
	})
	
	t.Run("rename to same name", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  same = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		mappings := map[string]string{
			"same": "same",
		}
		
		transformer := AttributeRenamer(mappings)
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Attribute should still exist with same value
		attr := block.Body().GetAttribute("same")
		assert.NotNil(t, attr)
		value := string(attr.Expr().BuildTokens(nil).Bytes())
		assert.Contains(t, value, "value")
	})
}

func BenchmarkAttributeRenamer(b *testing.B) {
	hclContent := `resource "test" "example" {
  attr1 = "value1"
  attr2 = "value2"
  attr3 = "value3"
  attr4 = "value4"
  attr5 = "value5"
}`
	
	mappings := map[string]string{
		"attr1": "new_attr1",
		"attr2": "new_attr2",
		"attr3": "new_attr3",
		"attr4": "new_attr4",
		"attr5": "new_attr5",
	}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := AttributeRenamer(mappings)
		_ = transformer(block, ctx)
	}
}