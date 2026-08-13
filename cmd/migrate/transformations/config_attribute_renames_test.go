package transformations

import (
	"os"
	"path/filepath"
	"testing"
)

// TestConfigAttributeRenames tests the attribute rename functionality for configuration files
func TestConfigAttributeRenames(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "config_attr_renames_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file with attribute renames
	configPath := filepath.Join(tempDir, "test_renames_config.yaml")
	configContent := `
version: "1.0"
description: "Test attribute renames"
attribute_renames:
  cloudflare_load_balancer:
    fallback_pool_id: fallback_pool
    default_pool_ids: default_pools
  cloudflare_api_token:
    policy: policies
  cloudflare_workers_kv:
    key: key_name
attribute_removals:
  cloudflare_access_policy:
    - application_id
    - precedence
  cloudflare_dns_record:
    - hostname
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "rename load balancer attributes",
			input: `resource "cloudflare_load_balancer" "example" {
  name             = "example-lb"
  fallback_pool_id = "pool-1"
  default_pool_ids = ["pool-2", "pool-3"]
}`,
			expected: `resource "cloudflare_load_balancer" "example" {
  name          = "example-lb"
  default_pools = ["pool-2", "pool-3"]
  fallback_pool = "pool-1"
}`,
		},
		{
			name: "rename api token policy to policies",
			input: `resource "cloudflare_api_token" "example" {
  name   = "example-token"
  policy = "read-only"
}`,
			expected: `resource "cloudflare_api_token" "example" {
  name     = "example-token"
  policies = "read-only"
}`,
		},
		{
			name: "rename workers_kv key to key_name",
			input: `resource "cloudflare_workers_kv" "example" {
  namespace_id = "namespace-1"
  key          = "my-key"
  value        = "my-value"
}`,
			expected: `resource "cloudflare_workers_kv" "example" {
  namespace_id = "namespace-1"
  value        = "my-value"
  key_name     = "my-key"
}`,
		},
		{
			name: "remove access policy attributes",
			input: `resource "cloudflare_access_policy" "example" {
  name           = "example-policy"
  application_id = "app-123"
  precedence     = 1
  decision       = "allow"
}`,
			expected: `resource "cloudflare_access_policy" "example" {
  name     = "example-policy"
  decision = "allow"
}`,
		},
		{
			name: "remove dns record hostname",
			input: `resource "cloudflare_dns_record" "example" {
  zone_id  = "zone-123"
  name     = "example"
  type     = "A"
  value    = "192.0.2.1"
  hostname = "example.com"
}`,
			expected: `resource "cloudflare_dns_record" "example" {
  zone_id = "zone-123"
  name    = "example"
  type    = "A"
  value   = "192.0.2.1"
}`,
		},
		{
			name: "handle resource with no renames or removals",
			input: `resource "cloudflare_zone" "example" {
  zone = "example.com"
  plan = "enterprise"
}`,
			expected: `resource "cloudflare_zone" "example" {
  zone = "example.com"
  plan = "enterprise"
}`,
		},
	}

	// Create transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input file
			inputPath := filepath.Join(tempDir, "input.tf")
			if err := os.WriteFile(inputPath, []byte(tt.input), 0644); err != nil {
				t.Fatalf("Failed to write input file: %v", err)
			}

			// Transform the file
			outputPath := filepath.Join(tempDir, "output.tf")
			if err := transformer.TransformFile(inputPath, outputPath); err != nil {
				t.Fatalf("Failed to transform file: %v", err)
			}

			// Read the output
			output, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			// Use semantic comparison that ignores attribute order
			if !compareHCLBlocks(t, tt.expected, string(output)) {
				// If semantic comparison fails, show the actual vs expected for debugging
				t.Errorf("Transformation mismatch\nExpected:\n%s\n\nGot:\n%s", tt.expected, string(output))
			}
		})
	}
}

// TestCombinedConfigTransformations tests combined block and attribute transformations
func TestCombinedConfigTransformations(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "config_combined_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file with both block and attribute transformations
	configPath := filepath.Join(tempDir, "test_combined_config.yaml")
	configContent := `
version: "1.0"
description: "Test combined transformations"
transformations:
  cloudflare_load_balancer_pool:
    to_list:
      - origins
attribute_renames:
  cloudflare_load_balancer_pool:
    minimum_origins: min_origins
attribute_removals:
  cloudflare_load_balancer_pool:
    - check_regions
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test case
	input := `resource "cloudflare_load_balancer_pool" "example" {
  name            = "example-pool"
  minimum_origins = 1
  check_regions   = ["WNAM"]

  origins {
    name    = "origin-1"
    address = "192.0.2.1"
  }

  origins {
    name    = "origin-2"
    address = "192.0.2.2"
  }
}`

	expected := `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"


  origins = [{
    address = "192.0.2.1"
    name    = "origin-1"
    }, {
    address = "192.0.2.2"
    name    = "origin-2"
  }]
  min_origins = 1
}`

	// Create transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	// Create input file
	inputPath := filepath.Join(tempDir, "input.tf")
	if err := os.WriteFile(inputPath, []byte(input), 0644); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	// Transform the file
	outputPath := filepath.Join(tempDir, "output.tf")
	if err := transformer.TransformFile(inputPath, outputPath); err != nil {
		t.Fatalf("Failed to transform file: %v", err)
	}

	// Read the output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Use semantic comparison that ignores attribute order
	if !compareHCLBlocks(t, expected, string(output)) {
		// If semantic comparison fails, show the actual vs expected for debugging
		t.Errorf("Transformation mismatch\nExpected:\n%s\n\nGot:\n%s", expected, string(output))
	}
}

// TestHasConfigAttributeRename tests checking for attribute renames
func TestHasConfigAttributeRename(t *testing.T) {
	// Create a temporary config file
	tempDir, err := os.MkdirTemp("", "config_attr_rename_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
attribute_renames:
  cloudflare_api_token:
    policy: policies
  cloudflare_workers_kv:
    key: key_name
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	tests := []struct {
		resourceType string
		attrName     string
		expectedName string
		expectedBool bool
	}{
		{"cloudflare_api_token", "policy", "policies", true},
		{"cloudflare_workers_kv", "key", "key_name", true},
		{"cloudflare_api_token", "name", "", false},
		{"unknown_resource", "policy", "", false},
	}

	for _, tt := range tests {
		newName, hasRename := transformer.HasAttributeRename(tt.resourceType, tt.attrName)
		if newName != tt.expectedName || hasRename != tt.expectedBool {
			t.Errorf("HasAttributeRename(%s, %s) = (%s, %v), want (%s, %v)",
				tt.resourceType, tt.attrName, newName, hasRename, tt.expectedName, tt.expectedBool)
		}
	}
}

// TestShouldRemoveConfigAttribute tests checking for attribute removals
func TestShouldRemoveConfigAttribute(t *testing.T) {
	// Create a temporary config file
	tempDir, err := os.MkdirTemp("", "config_attr_removal_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
attribute_removals:
  cloudflare_access_policy:
    - application_id
    - precedence
  cloudflare_dns_record:
    - hostname
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	tests := []struct {
		resourceType string
		attrName     string
		expected     bool
	}{
		{"cloudflare_access_policy", "application_id", true},
		{"cloudflare_access_policy", "precedence", true},
		{"cloudflare_dns_record", "hostname", true},
		{"cloudflare_access_policy", "name", false},
		{"unknown_resource", "application_id", false},
	}

	for _, tt := range tests {
		result := transformer.ShouldRemoveAttribute(tt.resourceType, tt.attrName)
		if result != tt.expected {
			t.Errorf("ShouldRemoveAttribute(%s, %s) = %v, want %v",
				tt.resourceType, tt.attrName, result, tt.expected)
		}
	}
}

