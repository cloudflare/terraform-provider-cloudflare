package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatHCLForDiff(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "format valid HCL",
			input: `resource "cloudflare_zone" "test" {
zone = "example.com"
plan = "free"
}`,
			expected: `resource "cloudflare_zone" "test" {
  zone = "example.com"
  plan = "free"
}`,
		},
		{
			name: "format HCL fragment",
			input: `  attribute   =   "value"  
  another  = true   `,
			expected: `attribute = "value"
another   = true   `,
		},
		{
			name: "handle empty input",
			input:    "",
			expected: "",
		},
		{
			name: "format with comments",
			input: `# Comment here
resource "test" "example" {
# Another comment
value = 1
}`,
			expected: `# Comment here
resource "test" "example" {
  # Another comment
  value = 1
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatHCLForDiff(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNormalizeHCLWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "normalize multiple spaces",
			input: `attribute    =    "value"
another   attribute   =   true`,
			expected: `attribute = "value"
another attribute = true`,
		},
		{
			name: "remove empty lines",
			input: `line1

line2

line3`,
			expected: `line1
line2
line3`,
		},
		{
			name: "trim leading and trailing whitespace",
			input: `   attribute = "value"   
	  another = true	  `,
			expected: `attribute = "value"
another = true`,
		},
		{
			name: "handle tabs",
			input: `attribute	=	"value"
		another	=	true`,
			expected: `attribute = "value"
another = true`,
		},
		{
			name: "empty input",
			input:    "",
			expected: "",
		},
		{
			name: "only whitespace",
			input:    "   \n\t\n   ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeHCLWhitespace(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
