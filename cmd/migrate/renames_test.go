package main

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestCustomPagesTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "simple attributes",
			Config: `
resource "cloudflare_custom_pages" "example" {
  account_id = "987654321fedcba0123456789abcdef0"
  type       = "basic_challenge"
  state      = "active"
  url        = "https://example.com/404"
}`,
			Expected: []string{`
resource "cloudflare_custom_pages" "example" {
  account_id = "987654321fedcba0123456789abcdef0"
  identifier = "basic_challenge"
  state      = "active"
  url        = "https://example.com/404"
}`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

func TestResourceReferenceRename(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "rename_access_policy_references",
			input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_access_policy.allow.id,
    cloudflare_access_policy.deny.id
  ]
}`,
			expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    cloudflare_zero_trust_access_policy.deny.id
  ]
}`,
		},
		{
			name: "rename_access_application_references",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  
  application_id = cloudflare_access_application.app.id
}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  
  application_id = cloudflare_zero_trust_access_application.app.id
}`,
		},
		{
			name: "rename_access_group_references",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  
  include = [{
    group = [cloudflare_access_group.engineering.id]
  }]
}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  
  include = [{
    group = [cloudflare_zero_trust_access_group.engineering.id]
  }]
}`,
		},
		{
			name: "multiple_references_in_one_line",
			input: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  
  policies = [
    cloudflare_access_policy.p1.id,
    cloudflare_access_policy.p2.id,
    cloudflare_access_policy.p3.id
  ]
}`,
			expected: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  
  policies = [
    cloudflare_zero_trust_access_policy.p1.id,
    cloudflare_zero_trust_access_policy.p2.id,
    cloudflare_zero_trust_access_policy.p3.id
  ]
}`,
		},
		{
			name: "do_not_rename_other_resources",
			input: `resource "cloudflare_zone" "test" {
  account_id = "abc123"
  zone       = "example.com"
  
  settings_override_id = cloudflare_zone_settings_override.settings.id
}`,
			expected: `resource "cloudflare_zone" "test" {
  account_id = "abc123"
  zone       = "example.com"
  
  settings_override_id = cloudflare_zone_settings_override.settings.id
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse input: %v", diags.Error())
			}

			// Apply the transformation
			for _, block := range file.Body().Blocks() {
				applyRenames(block)
			}

			// Format the output
			output := string(hclwrite.Format(file.Bytes()))

			// Normalize whitespace for comparison
			expected := normalizeWhitespace(tt.expected)
			actual := normalizeWhitespace(output)

			if expected != actual {
				t.Errorf("Transformation mismatch.\nExpected:\n%s\n\nGot:\n%s\n\nRaw output:\n%s", expected, actual, output)
			}
		})
	}
}
