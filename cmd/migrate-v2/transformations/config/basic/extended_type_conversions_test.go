package basic

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtendedTypeConverter(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		conversions    []TypeConversion
		expectedOutput string
		checkFunc      func(t *testing.T, output string)
	}{
		{
			name: "number to string conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  port = 8080
  timeout = 30
}`,
			conversions: []TypeConversion{
				{
					Attribute: "port",
					FromType:  "number",
					ToType:    "string",
				},
				{
					Attribute: "timeout",
					FromType:  "int",
					ToType:    "string",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `port    = "8080"`)
				assert.Contains(t, output, `timeout = "30"`)
			},
		},
		{
			name: "string to number conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  max_size = "1024"
  ratio = "0.75"
}`,
			conversions: []TypeConversion{
				{
					Attribute: "max_size",
					FromType:  "string",
					ToType:    "int",
				},
				{
					Attribute: "ratio",
					FromType:  "string",
					ToType:    "float",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "max_size = 1024")
				assert.NotContains(t, output, `max_size = "1024"`)
				assert.Contains(t, output, "ratio    = 0.75")
			},
		},
		{
			name: "boolean to string conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  enabled = true
  active = false
}`,
			conversions: []TypeConversion{
				{
					Attribute: "enabled",
					FromType:  "bool",
					ToType:    "string",
				},
				{
					Attribute: "active",
					FromType:  "boolean",
					ToType:    "string",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `enabled = "true"`)
				assert.Contains(t, output, `active  = "false"`)
			},
		},
		{
			name: "string to boolean conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  ssl = "enabled"
  debug = "0"
  verbose = "yes"
}`,
			conversions: []TypeConversion{
				{
					Attribute: "ssl",
					FromType:  "string",
					ToType:    "bool",
				},
				{
					Attribute: "debug",
					FromType:  "string",
					ToType:    "boolean",
				},
				{
					Attribute: "verbose",
					FromType:  "string",
					ToType:    "bool",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "ssl     = true")
				assert.Contains(t, output, "debug   = false")
				assert.Contains(t, output, "verbose = true")
			},
		},
		{
			name: "single value to list conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  ip = "192.168.1.1"
  port = 443
}`,
			conversions: []TypeConversion{
				{
					Attribute: "ip",
					FromType:  "string",
					ToType:    "list",
				},
				{
					Attribute: "port",
					FromType:  "number",
					ToType:    "list",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `ip      = ["192.168.1.1"]`)
				assert.Contains(t, output, "port    = [443]")
			},
		},
		{
			name: "list to single value conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  servers = ["server1", "server2"]
  ports = [80]
}`,
			conversions: []TypeConversion{
				{
					Attribute: "servers",
					FromType:  "list",
					ToType:    "single",
				},
				{
					Attribute: "ports",
					FromType:  "list",
					ToType:    "single",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `servers = "server1"`)
				assert.NotContains(t, output, "server2")
				assert.Contains(t, output, "ports   = 80")
				assert.NotContains(t, output, "[80]")
			},
		},
		{
			name: "number to boolean conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  count = 0
  retries = 3
}`,
			conversions: []TypeConversion{
				{
					Attribute: "count",
					FromType:  "number",
					ToType:    "bool",
				},
				{
					Attribute: "retries",
					FromType:  "int",
					ToType:    "boolean",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "count   = false")
				assert.Contains(t, output, "retries = true")
			},
		},
		{
			name: "boolean to number conversion",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  enabled = true
  disabled = false
}`,
			conversions: []TypeConversion{
				{
					Attribute: "enabled",
					FromType:  "bool",
					ToType:    "number",
				},
				{
					Attribute: "disabled",
					FromType:  "boolean",
					ToType:    "int",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "enabled  = 1")
				assert.Contains(t, output, "disabled = 0")
			},
		},
		{
			name: "multiple conversions on same resource",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  port = "8080"
  ssl = true
  servers = ["primary"]
}`,
			conversions: []TypeConversion{
				{
					Attribute: "port",
					FromType:  "string",
					ToType:    "number",
				},
				{
					Attribute: "ssl",
					FromType:  "bool",
					ToType:    "string",
				},
				{
					Attribute: "servers",
					FromType:  "list",
					ToType:    "single",
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "port    = 8080")
				assert.Contains(t, output, `ssl     = "true"`)
				assert.Contains(t, output, `servers = "primary"`)
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
			ctx := &TransformContext{}
			converter := ExtendedTypeConverter(tt.conversions)
			
			err := converter(block, ctx)
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

func TestExtendedTypeConverterForState(t *testing.T) {
	tests := []struct {
		name          string
		state         map[string]interface{}
		conversions   []TypeConversion
		expectedState map[string]interface{}
	}{
		{
			name: "number to string in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"port":    8080,
					"timeout": 30,
				},
			},
			conversions: []TypeConversion{
				{
					Attribute: "port",
					FromType:  "number",
					ToType:    "string",
				},
				{
					Attribute: "timeout",
					FromType:  "int",
					ToType:    "string",
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"port":    "8080",
					"timeout": "30",
				},
			},
		},
		{
			name: "string to number in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":  "abc123",
					"max_size": "1024",
					"ratio":    "0.75",
				},
			},
			conversions: []TypeConversion{
				{
					Attribute: "max_size",
					FromType:  "string",
					ToType:    "int",
				},
				{
					Attribute: "ratio",
					FromType:  "string",
					ToType:    "float",
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":  "abc123",
					"max_size": 1024,
					"ratio":    0.75,
				},
			},
		},
		{
			name: "boolean conversions in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"enabled": true,
					"debug":   "yes",
					"count":   0,
				},
			},
			conversions: []TypeConversion{
				{
					Attribute: "enabled",
					FromType:  "bool",
					ToType:    "string",
				},
				{
					Attribute: "debug",
					FromType:  "string",
					ToType:    "bool",
				},
				{
					Attribute: "count",
					FromType:  "number",
					ToType:    "bool",
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"enabled": "true",
					"debug":   true,
					"count":   false,
				},
			},
		},
		{
			name: "list conversions in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"ip":      "192.168.1.1",
					"servers": []interface{}{"server1", "server2"},
				},
			},
			conversions: []TypeConversion{
				{
					Attribute: "ip",
					FromType:  "string",
					ToType:    "list",
				},
				{
					Attribute: "servers",
					FromType:  "list",
					ToType:    "single",
				},
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"ip":      []interface{}{"192.168.1.1"},
					"servers": "server1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply state transformation
			converter := ExtendedTypeConverterForState(tt.conversions)
			err := converter(tt.state)
			require.NoError(t, err)

			// Compare states
			assert.Equal(t, tt.expectedState, tt.state)
		})
	}
}

func TestParseListElements(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple list",
			input:    `"item1", "item2", "item3"`,
			expected: []string{`"item1"`, `"item2"`, `"item3"`},
		},
		{
			name:     "list with spaces",
			input:    `  "item1"  ,  "item2"  `,
			expected: []string{`"item1"`, `"item2"`},
		},
		{
			name:     "mixed quotes",
			input:    `"item1", 'item2', "item3"`,
			expected: []string{`"item1"`, `'item2'`, `"item3"`},
		},
		{
			name:     "list with commas in values",
			input:    `"item,1", "item,2"`,
			expected: []string{`"item,1"`, `"item,2"`},
		},
		{
			name:     "numbers in list",
			input:    `1, 2, 3`,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "mixed types",
			input:    `"string", 123, true`,
			expected: []string{`"string"`, "123", "true"},
		},
		{
			name:     "empty list",
			input:    ``,
			expected: []string{},
		},
		{
			name:     "single element",
			input:    `"single"`,
			expected: []string{`"single"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseListElements(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConversionEdgeCases(t *testing.T) {
	t.Run("empty conversions list", func(t *testing.T) {
		converter := ExtendedTypeConverter([]TypeConversion{})
		
		input := `resource "test" "example" { zone_id = "abc" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		err := converter(block, ctx)
		assert.NoError(t, err)
	})
	
	t.Run("non-existent attribute", func(t *testing.T) {
		conversions := []TypeConversion{
			{
				Attribute: "nonexistent",
				FromType:  "string",
				ToType:    "number",
			},
		}
		
		input := `resource "test" "example" { zone_id = "abc" }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		converter := ExtendedTypeConverter(conversions)
		err := converter(block, ctx)
		assert.NoError(t, err) // Should not error on missing attributes
	})
	
	t.Run("invalid conversion", func(t *testing.T) {
		conversions := []TypeConversion{
			{
				Attribute: "invalid",
				FromType:  "string",
				ToType:    "number",
			},
		}
		
		input := `resource "test" "example" { 
  zone_id = "abc"
  invalid = "not_a_number"
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		converter := ExtendedTypeConverter(conversions)
		err := converter(block, ctx)
		assert.NoError(t, err) // Should not error but log diagnostic
		assert.NotEmpty(t, ctx.Diagnostics) // Should have diagnostic message
	})
	
	t.Run("empty list to single conversion", func(t *testing.T) {
		conversions := []TypeConversion{
			{
				Attribute: "empty_list",
				FromType:  "list",
				ToType:    "single",
			},
		}
		
		input := `resource "test" "example" { 
  zone_id = "abc"
  empty_list = []
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		converter := ExtendedTypeConverter(conversions)
		err := converter(block, ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, ctx.Diagnostics) // Should have diagnostic about empty list
	})
	
	t.Run("unknown conversion type", func(t *testing.T) {
		conversions := []TypeConversion{
			{
				Attribute: "test",
				FromType:  "string",
				ToType:    "unknown_type",
			},
		}
		
		input := `resource "test" "example" { 
  zone_id = "abc"
  test = "value"
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		converter := ExtendedTypeConverter(conversions)
		err := converter(block, ctx)
		assert.NoError(t, err) // Should not error on unknown types
		
		// Value should remain unchanged
		output := string(hclwrite.Format(file.Bytes()))
		assert.Contains(t, output, `test    = "value"`)
	})
}