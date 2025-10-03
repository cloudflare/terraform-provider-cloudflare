package conditional

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

func TestConditionalTransformer(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		transforms     []basic.ConditionalTransform
		expectedOutput string
		checkFunc      func(t *testing.T, output string)
	}{
		{
			name: "condition equals with then action",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  type    = "premium"
  enabled = true
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "type",
						Operator:  "equals",
						Value:     "premium",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "gold",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `tier    = "gold"`)
				assert.Contains(t, output, `type    = "premium"`)
			},
		},
		{
			name: "condition with else action",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  type    = "basic"
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "type",
						Operator:  "equals",
						Value:     "premium",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "gold",
						},
					},
					Else: &basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "silver",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `tier    = "silver"`)
				assert.NotContains(t, output, `tier = "gold"`)
			},
		},
		{
			name: "condition exists operator",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  optional_field = "value"
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "optional_field",
						Operator:  "exists",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"has_optional": "true",
						},
						RenameAttributes: map[string]string{
							"optional_field": "required_field",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `has_optional   = true`)
				assert.Contains(t, output, `required_field = "value"`)
				assert.NotContains(t, output, "optional_field")
			},
		},
		{
			name: "condition not_exists operator",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "missing_field",
						Operator:  "not_exists",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"default_field": "default_value",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `default_field = "default_value"`)
			},
		},
		{
			name: "condition with remove attributes",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  environment = "dev"
  debug = true
  verbose = true
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "environment",
						Operator:  "equals",
						Value:     "dev",
					},
					Then: basic.TransformActions{
						RemoveAttributes: []string{"debug", "verbose"},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "debug")
				assert.NotContains(t, output, "verbose")
				assert.Contains(t, output, "environment")
			},
		},
		{
			name: "condition contains operator",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  description = "This is a premium feature"
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "description",
						Operator:  "contains",
						Value:     "premium",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "premium",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `tier        = "premium"`)
			},
		},
		{
			name: "multiple conditions",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  type = "basic"
  region = "us"
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "type",
						Operator:  "equals",
						Value:     "basic",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "free",
						},
					},
				},
				{
					Condition: basic.TransformCondition{
						Attribute: "region",
						Operator:  "equals",
						Value:     "us",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"datacenter": "us-west-1",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `tier       = "free"`)
				assert.Contains(t, output, `datacenter = "us-west-1"`)
			},
		},
		{
			name: "complex rename and set operations",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  old_setting = "value1"
  deprecated = "old"
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "deprecated",
						Operator:  "exists",
					},
					Then: basic.TransformActions{
						RenameAttributes: map[string]string{
							"old_setting": "new_setting",
							"deprecated":  "legacy_mode",
						},
						SetAttributes: map[string]string{
							"migration_version": "2",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "new_setting")
				assert.Contains(t, output, "legacy_mode")
				assert.Contains(t, output, "migration_version")
				assert.NotContains(t, output, "old_setting")
				assert.NotContains(t, output, `deprecated =`)
			},
		},
		{
			name: "is_empty operator",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  items = []
}`,
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "items",
						Operator:  "is_empty",
					},
					Then: basic.TransformActions{
						RemoveAttributes: []string{"items"},
						SetAttributes: map[string]string{
							"has_items": "false",
						},
					},
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "\n  items")  // Check that the items attribute was removed (at start of line)
				assert.NotContains(t, output, "items = []")  
				assert.Contains(t, output, `has_items = false`)
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
			transformer := ConditionalTransformer(tt.transforms)
			
			err := transformer(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestConditionalTransformerForState(t *testing.T) {
	tests := []struct {
		name          string
		state         map[string]interface{}
		transforms    []basic.ConditionalTransform
		expectedState map[string]interface{}
	}{
		{
			name: "simple condition with then action",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"type":    "premium",
				},
			},
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "type",
						Operator:  "equals",
						Value:     "premium",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "gold",
						},
					},
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"type":    "premium",
					"tier":    "gold",
				},
			},
		},
		{
			name: "condition with else in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"type":    "basic",
				},
			},
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "type",
						Operator:  "equals",
						Value:     "premium",
					},
					Then: basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "gold",
						},
					},
					Else: &basic.TransformActions{
						SetAttributes: map[string]string{
							"tier": "silver",
						},
					},
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"type":    "basic",
					"tier":    "silver",
				},
			},
		},
		{
			name: "is_empty on array in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"items":   []interface{}{},
				},
			},
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "items",
						Operator:  "is_empty",
					},
					Then: basic.TransformActions{
						RemoveAttributes: []string{"items"},
						SetAttributes: map[string]string{
							"has_items": "false",
						},
					},
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":   "abc123",
					"has_items": false,  // parseStateValue returns boolean false
				},
			},
		},
		{
			name: "rename attributes in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":     "abc123",
					"old_setting": "value",
					"deprecated":  true,
				},
			},
			transforms: []basic.ConditionalTransform{
				{
					Condition: basic.TransformCondition{
						Attribute: "deprecated",
						Operator:  "exists",
					},
					Then: basic.TransformActions{
						RenameAttributes: map[string]string{
							"old_setting": "new_setting",
							"deprecated":  "legacy_mode",
						},
					},
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":      "abc123",
					"new_setting":  "value",
					"legacy_mode":  true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply state transformation
			transformer := ConditionalTransformerForState(tt.transforms)
			err := transformer(tt.state)
			require.NoError(t, err)

			// Compare states
			assert.Equal(t, tt.expectedState, tt.state)
		})
	}
}

func TestEvaluateConditionOperators(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		condition basic.TransformCondition
		expected  bool
	}{
		{
			name: "starts_with operator",
			input: `resource "test" "example" {
  name = "prefix_value"
}`,
			condition: basic.TransformCondition{
				Attribute: "name",
				Operator:  "starts_with",
				Value:     "prefix",
			},
			expected: true,
		},
		{
			name: "ends_with operator",
			input: `resource "test" "example" {
  name = "value_suffix"
}`,
			condition: basic.TransformCondition{
				Attribute: "name",
				Operator:  "ends_with",
				Value:     "suffix",
			},
			expected: true,
		},
		{
			name: "not_equals operator",
			input: `resource "test" "example" {
  type = "basic"
}`,
			condition: basic.TransformCondition{
				Attribute: "type",
				Operator:  "not_equals",
				Value:     "premium",
			},
			expected: true,
		},
		{
			name: "is_not_empty operator",
			input: `resource "test" "example" {
  items = ["one", "two"]
}`,
			condition: basic.TransformCondition{
				Attribute: "items",
				Operator:  "is_not_empty",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			
			body := blocks[0].Body()
			result := evaluateCondition(body, tt.condition)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConditionalTransformerEdgeCases(t *testing.T) {
	t.Run("empty transforms list", func(t *testing.T) {
		transformer := ConditionalTransformer([]basic.ConditionalTransform{})
		
		input := `resource "test" "example" { zone_id = "abc" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &basic.TransformContext{}
		err := transformer(block, ctx)
		assert.NoError(t, err)
	})
	
	t.Run("nil else clause", func(t *testing.T) {
		transforms := []basic.ConditionalTransform{
			{
				Condition: basic.TransformCondition{
					Attribute: "nonexistent",
					Operator:  "exists",
				},
				Then: basic.TransformActions{
					SetAttributes: map[string]string{"new": "value"},
				},
				// Else is nil
			},
		}
		
		input := `resource "test" "example" { zone_id = "abc" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &basic.TransformContext{}
		transformer := ConditionalTransformer(transforms)
		err := transformer(block, ctx)
		assert.NoError(t, err)
		
		// Should not have added the attribute since condition was false and no else
		output := string(file.Bytes())
		assert.NotContains(t, output, "new")
	})
}