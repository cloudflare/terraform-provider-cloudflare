package transformations

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestBlockToAttributeTransformations tests the block-to-attribute transformation functionality
func TestBlockToAttributeTransformations(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "block_to_attribute_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file for block transformations
	configPath := filepath.Join(tempDir, "test_config.yaml")
	configContent := `
version: "1.0"
description: "Test block to attribute transformations"
transformations:
  cloudflare_load_balancer_pool:
    to_map:
      - load_shedding
      - origin_steering
    to_list:
      - origins
  cloudflare_access_application:
    to_map:
      - cors_headers
    to_list:
      - footer_links
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
			name: "transform load balancer pool origins to list",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins {
    name    = "origin-1"
    address = "192.0.2.1"
    enabled = true
  }

  origins {
    name    = "origin-2"
    address = "192.0.2.2"
    enabled = false
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins = [{
    address = "192.0.2.1"
    enabled = true
    name    = "origin-1"
  }, {
    address = "192.0.2.2"
    enabled = false
    name    = "origin-2"
  }]
}`,
		},
		{
			name: "transform load balancer pool load_shedding to map",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  load_shedding {
    default_percent = 10
    default_policy  = "random"
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  load_shedding = {
    default_percent = 10
    default_policy  = "random"
  }
}`,
		},
		{
			name: "transform access application cors_headers to map",
			input: `resource "cloudflare_access_application" "example" {
  name = "example-app"

  cors_headers {
    allowed_methods   = ["GET", "POST"]
    allowed_origins   = ["https://example.com"]
    allow_credentials = true
  }
}`,
			expected: `resource "cloudflare_access_application" "example" {
  name = "example-app"

  cors_headers = {
    allow_credentials = true
    allowed_methods   = ["GET", "POST"]
    allowed_origins   = ["https://example.com"]
  }
}`,
		},
		{
			name: "transform access application footer_links to list",
			input: `resource "cloudflare_access_application" "example" {
  name = "example-app"

  footer_links {
    name = "Link 1"
    url  = "https://example.com/1"
  }

  footer_links {
    name = "Link 2"
    url  = "https://example.com/2"
  }
}`,
			expected: `resource "cloudflare_access_application" "example" {
  name = "example-app"

  footer_links = [{
    name = "Link 1"
    url  = "https://example.com/1"
  }, {
    name = "Link 2"
    url  = "https://example.com/2"
  }]
}`,
		},
		{
			name: "handle multiple transformations in same resource",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins {
    name    = "origin-1"
    address = "192.0.2.1"
  }

  load_shedding {
    default_percent = 10
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  load_shedding = {
    default_percent = 10
  }
  origins = [{
    address = "192.0.2.1"
    name    = "origin-1"
  }]
}`,
		},
		{
			name: "preserve attributes not in transformation list",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name            = "example-pool"
  enabled         = true
  minimum_origins = 1

  origins {
    name    = "origin-1"
    address = "192.0.2.1"
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name            = "example-pool"
  enabled         = true
  minimum_origins = 1

  origins = [{
    address = "192.0.2.1"
    name    = "origin-1"
  }]
}`,
		},
		{
			name: "handle resources with no transformations",
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

			// Normalize whitespace for comparison
			expectedNorm := normalizeWhitespace(tt.expected)
			outputNorm := normalizeWhitespace(string(output))

			if expectedNorm != outputNorm {
				t.Errorf("Transformation mismatch\nExpected:\n%s\n\nGot:\n%s", tt.expected, string(output))
			}
		})
	}
}

// TestBlockTransformDirectory tests directory-level block transformations
func TestBlockTransformDirectory(t *testing.T) {
	// Create a temporary directory structure
	tempDir, err := os.MkdirTemp("", "block_transform_dir_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
transformations:
  cloudflare_load_balancer_pool:
    to_list:
      - origins
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Create directory structure with .tf files
	subDir := filepath.Join(tempDir, "modules")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Create test files
	testFiles := map[string]string{
		filepath.Join(tempDir, "main.tf"): `resource "cloudflare_load_balancer_pool" "main" {
  origins {
    name = "origin-1"
  }
}`,
		filepath.Join(subDir, "module.tf"): `resource "cloudflare_load_balancer_pool" "module" {
  origins {
    name = "origin-2"
  }
}`,
		filepath.Join(tempDir, "variables.tf"): `variable "example" {
  default = "test"
}`,
	}

	for path, content := range testFiles {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to write test file %s: %v", path, err)
		}
	}

	// Create transformer and process directory
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	// Test non-recursive processing
	if err := transformer.TransformDirectory(tempDir, false); err != nil {
		t.Fatalf("Failed to transform directory: %v", err)
	}

	// Check main.tf was transformed
	mainContent, err := os.ReadFile(filepath.Join(tempDir, "main.tf"))
	if err != nil {
		t.Fatalf("Failed to read main.tf: %v", err)
	}
	if !strings.Contains(string(mainContent), "origins = [") {
		t.Error("main.tf was not transformed")
	}

	// Check module.tf was NOT transformed (non-recursive)
	moduleContent, err := os.ReadFile(filepath.Join(subDir, "module.tf"))
	if err != nil {
		t.Fatalf("Failed to read module.tf: %v", err)
	}
	if strings.Contains(string(moduleContent), "origins = [") {
		t.Error("module.tf should not have been transformed in non-recursive mode")
	}

	// Test recursive processing
	if err := transformer.TransformDirectory(tempDir, true); err != nil {
		t.Fatalf("Failed to transform directory recursively: %v", err)
	}

	// Check module.tf was transformed
	moduleContent, err = os.ReadFile(filepath.Join(subDir, "module.tf"))
	if err != nil {
		t.Fatalf("Failed to read module.tf: %v", err)
	}
	if !strings.Contains(string(moduleContent), "origins = [") {
		t.Error("module.tf was not transformed in recursive mode")
	}
}

// TestGetBlockTransformationType tests the transformation type retrieval
func TestGetBlockTransformationType(t *testing.T) {
	// Create a temporary config file
	tempDir, err := os.MkdirTemp("", "block_transform_type_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
transformations:
  cloudflare_test_resource:
    to_map:
      - map_attr
    to_list:
      - list_attr
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
		blockName    string
		expected     string
	}{
		{"cloudflare_test_resource", "map_attr", "map"},
		{"cloudflare_test_resource", "list_attr", "list"},
		{"cloudflare_test_resource", "unknown_attr", ""},
		{"unknown_resource", "any_attr", ""},
	}

	for _, tt := range tests {
		result := transformer.GetTransformationType(tt.resourceType, tt.blockName)
		if result != tt.expected {
			t.Errorf("GetTransformationType(%s, %s) = %s, want %s",
				tt.resourceType, tt.blockName, result, tt.expected)
		}
	}
}