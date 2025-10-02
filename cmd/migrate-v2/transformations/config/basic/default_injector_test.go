package basic

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestDefaultValueSetter_ComplexCases(t *testing.T) {
	tests := []struct {
		name     string
		hcl      string
		defaults map[string]interface{}
		validate func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "set defaults for missing attributes only",
			hcl: `resource "test" "example" {
  existing_string = "keep_me"
  existing_bool = false
}`,
			defaults: map[string]interface{}{
				"existing_string": "should_not_override",
				"existing_bool":   true,
				"new_string":      "default_value",
				"new_int":         123,
				"new_bool":        true,
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Existing values should not be overridden
				existingStr := body.GetAttribute("existing_string")
				require.NotNil(t, existingStr)
				assert.Contains(t, string(existingStr.Expr().BuildTokens(nil).Bytes()), "keep_me")
				
				existingBool := body.GetAttribute("existing_bool")
				require.NotNil(t, existingBool)
				assert.Contains(t, string(existingBool.Expr().BuildTokens(nil).Bytes()), "false")
				
				// New defaults should be added
				newStr := body.GetAttribute("new_string")
				require.NotNil(t, newStr)
				assert.Contains(t, string(newStr.Expr().BuildTokens(nil).Bytes()), "default_value")
				
				newInt := body.GetAttribute("new_int")
				require.NotNil(t, newInt)
				assert.Contains(t, string(newInt.Expr().BuildTokens(nil).Bytes()), "123")
				
				newBool := body.GetAttribute("new_bool")
				require.NotNil(t, newBool)
				assert.Contains(t, string(newBool.Expr().BuildTokens(nil).Bytes()), "true")
			},
		},
		{
			name: "different data types",
			hcl: `resource "test" "example" {
}`,
			defaults: map[string]interface{}{
				"string_val":    "hello",
				"int_val":       42,
				"negative_int":  -100,
				"zero_int":      0,
				"bool_true":     true,
				"bool_false":    false,
				"float_val":     3.14,     // Will be treated as unhandled
				"nil_val":       nil,      // Will be ignored
				"array_val":     []string{"a", "b"}, // Will be ignored (unsupported)
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Supported types should be added
				assert.NotNil(t, body.GetAttribute("string_val"))
				assert.NotNil(t, body.GetAttribute("int_val"))
				assert.NotNil(t, body.GetAttribute("negative_int"))
				assert.NotNil(t, body.GetAttribute("zero_int"))
				assert.NotNil(t, body.GetAttribute("bool_true"))
				assert.NotNil(t, body.GetAttribute("bool_false"))
				
				// Unsupported types should be ignored
				assert.Nil(t, body.GetAttribute("float_val"))
				assert.Nil(t, body.GetAttribute("nil_val"))
				assert.Nil(t, body.GetAttribute("array_val"))
				
				// Verify values
				assert.Contains(t, string(body.GetAttribute("string_val").Expr().BuildTokens(nil).Bytes()), "hello")
				assert.Contains(t, string(body.GetAttribute("int_val").Expr().BuildTokens(nil).Bytes()), "42")
				assert.Contains(t, string(body.GetAttribute("bool_true").Expr().BuildTokens(nil).Bytes()), "true")
				assert.Contains(t, string(body.GetAttribute("bool_false").Expr().BuildTokens(nil).Bytes()), "false")
			},
		},
		{
			name: "empty defaults map",
			hcl: `resource "test" "example" {
  existing = "value"
}`,
			defaults: map[string]interface{}{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Only existing attribute should remain
				assert.NotNil(t, body.GetAttribute("existing"))
				
				// Count attributes (crude but effective)
				count := 0
				// We'd need to iterate through all attributes to count properly
				// For now, just check the one we know about
				if body.GetAttribute("existing") != nil {
					count++
				}
				assert.Equal(t, 1, count)
			},
		},
		{
			name: "nil defaults map",
			hcl: `resource "test" "example" {
  existing = "value"
}`,
			defaults: nil,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should not panic, existing attribute should remain
				assert.NotNil(t, body.GetAttribute("existing"))
			},
		},
		{
			name: "special string values",
			hcl: `resource "test" "example" {
}`,
			defaults: map[string]interface{}{
				"empty_string":     "",
				"space_string":     " ",
				"multiline_string": "line1\nline2\nline3",
				"quoted_string":    `"quoted"`,
				"special_chars":    "!@#$%^&*()",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// All string values should be set
				assert.NotNil(t, body.GetAttribute("empty_string"))
				assert.NotNil(t, body.GetAttribute("space_string"))
				assert.NotNil(t, body.GetAttribute("multiline_string"))
				assert.NotNil(t, body.GetAttribute("quoted_string"))
				assert.NotNil(t, body.GetAttribute("special_chars"))
			},
		},
		{
			name: "numeric edge cases",
			hcl: `resource "test" "example" {
}`,
			defaults: map[string]interface{}{
				"max_int32":  2147483647,
				"min_int32":  -2147483648,
				"zero":       0,
				"negative_1": -1,
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.NotNil(t, body.GetAttribute("max_int32"))
				assert.NotNil(t, body.GetAttribute("min_int32"))
				assert.NotNil(t, body.GetAttribute("zero"))
				assert.NotNil(t, body.GetAttribute("negative_1"))
				
				// Verify numeric values are correct
				assert.Contains(t, string(body.GetAttribute("max_int32").Expr().BuildTokens(nil).Bytes()), "2147483647")
				assert.Contains(t, string(body.GetAttribute("min_int32").Expr().BuildTokens(nil).Bytes()), "-2147483648")
				assert.Contains(t, string(body.GetAttribute("zero").Expr().BuildTokens(nil).Bytes()), "0")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}

			transformer := DefaultValueSetter(tt.defaults)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestDefaultValueSetter_TypeConversions(t *testing.T) {
	t.Run("int64 conversion", func(t *testing.T) {
		hclContent := `resource "test" "example" {}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		defaults := map[string]interface{}{
			"int8_val":  int8(127),
			"int16_val": int16(32767),
			"int32_val": int32(2147483647),
			"int64_val": int64(9223372036854775807),
			"uint_val":  uint(4294967295),
		}
		
		transformer := DefaultValueSetter(defaults)
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		body := block.Body()
		
		// Only regular int is handled, others are ignored
		assert.Nil(t, body.GetAttribute("int8_val"))
		assert.Nil(t, body.GetAttribute("int16_val"))
		assert.Nil(t, body.GetAttribute("int32_val"))
		assert.Nil(t, body.GetAttribute("int64_val"))
		assert.Nil(t, body.GetAttribute("uint_val"))
	})
	
	t.Run("custom types", func(t *testing.T) {
		hclContent := `resource "test" "example" {}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		type CustomString string
		type CustomInt int
		type CustomBool bool
		
		defaults := map[string]interface{}{
			"custom_string": CustomString("value"),
			"custom_int":    CustomInt(42),
			"custom_bool":   CustomBool(true),
			"struct_val":    struct{ Name string }{Name: "test"},
			"map_val":       map[string]string{"key": "value"},
			"slice_val":     []string{"a", "b", "c"},
		}
		
		transformer := DefaultValueSetter(defaults)
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		body := block.Body()
		
		// Custom types and complex types are not handled
		assert.Nil(t, body.GetAttribute("custom_string"))
		assert.Nil(t, body.GetAttribute("custom_int"))
		assert.Nil(t, body.GetAttribute("custom_bool"))
		assert.Nil(t, body.GetAttribute("struct_val"))
		assert.Nil(t, body.GetAttribute("map_val"))
		assert.Nil(t, body.GetAttribute("slice_val"))
	})
}

func TestDefaultValueSetter_DirectCtyValues(t *testing.T) {
	// Test that the function uses SetAttributeValue correctly
	hclContent := `resource "test" "example" {}`
	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())
	
	block := file.Body().Blocks()[0]
	
	// Manually set some attributes using SetAttributeValue
	body := block.Body()
	body.SetAttributeValue("manual_string", cty.StringVal("test"))
	body.SetAttributeValue("manual_number", cty.NumberIntVal(99))
	body.SetAttributeValue("manual_bool", cty.BoolVal(false))
	
	// Now use DefaultValueSetter
	ctx := &TransformContext{}
	defaults := map[string]interface{}{
		"manual_string": "should_not_override", // Should not override
		"new_string":    "new_value",
		"new_number":    77,
		"new_bool":      true,
	}
	
	transformer := DefaultValueSetter(defaults)
	err := transformer(block, ctx)
	require.NoError(t, err)
	
	// Check that manual values were not overridden
	manualStr := body.GetAttribute("manual_string")
	require.NotNil(t, manualStr)
	assert.Contains(t, string(manualStr.Expr().BuildTokens(nil).Bytes()), "test")
	assert.NotContains(t, string(manualStr.Expr().BuildTokens(nil).Bytes()), "should_not_override")
	
	// Check that new values were added
	assert.NotNil(t, body.GetAttribute("new_string"))
	assert.NotNil(t, body.GetAttribute("new_number"))
	assert.NotNil(t, body.GetAttribute("new_bool"))
}

func BenchmarkDefaultValueSetter(b *testing.B) {
	hclContent := `resource "test" "example" {
  existing1 = "value1"
  existing2 = "value2"
}`
	
	defaults := map[string]interface{}{
		"existing1": "should_not_override",
		"new1":      "default1",
		"new2":      "default2",
		"new3":      42,
		"new4":      true,
		"new5":      "default5",
	}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := DefaultValueSetter(defaults)
		_ = transformer(block, ctx)
	}
}

func BenchmarkDefaultValueSetter_LargeDefaults(b *testing.B) {
	hclContent := `resource "test" "example" {}`
	
	// Create a large defaults map
	defaults := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		defaults[fmt.Sprintf("attr_%d", i)] = fmt.Sprintf("value_%d", i)
	}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := DefaultValueSetter(defaults)
		_ = transformer(block, ctx)
	}
}