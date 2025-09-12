package transformations

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResourceRenameTransformer(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "resource_rename_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file
	configPath := filepath.Join(tempDir, "test_resource_rename_config.yaml")
	configContent := `
version: "1.0"
description: "Test resource renames"
resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
  cloudflare_access_policy: cloudflare_zero_trust_access_policy
  cloudflare_teams_account: cloudflare_zero_trust_gateway_settings
  cloudflare_tunnel: cloudflare_zero_trust_tunnel_cloudflared
  cloudflare_managed_headers: cloudflare_managed_transforms
resource_categories:
  zero_trust_access:
    - cloudflare_access_application
    - cloudflare_access_policy
  zero_trust_gateway:
    - cloudflare_teams_account
  zero_trust_tunnel:
    - cloudflare_tunnel
  other:
    - cloudflare_managed_headers
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
			name: "rename access application resource",
			input: `resource "cloudflare_access_application" "example" {
  name   = "example-app"
  domain = "example.com"
}`,
			expected: `resource "cloudflare_zero_trust_access_application" "example" {
  name   = "example-app"
  domain = "example.com"
}`,
		},
		{
			name: "rename access policy resource",
			input: `resource "cloudflare_access_policy" "example" {
  name     = "example-policy"
  decision = "allow"
}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "example" {
  decision = "allow"
  name     = "example-policy"
}`,
		},
		{
			name: "rename teams account resource",
			input: `resource "cloudflare_teams_account" "example" {
  account_id = "123456"
}`,
			expected: `resource "cloudflare_zero_trust_gateway_settings" "example" {
  account_id = "123456"
}`,
		},
		{
			name: "rename tunnel resource",
			input: `resource "cloudflare_tunnel" "example" {
  name       = "example-tunnel"
  account_id = "123456"
}`,
			expected: `resource "cloudflare_zero_trust_tunnel_cloudflared" "example" {
  name       = "example-tunnel"
  account_id = "123456"
}`,
		},
		{
			name: "rename managed headers resource",
			input: `resource "cloudflare_managed_headers" "example" {
  zone_id = "zone123"
}`,
			expected: `resource "cloudflare_managed_transforms" "example" {
  zone_id = "zone123"
}`,
		},
		{
			name: "preserve resource with no rename",
			input: `resource "cloudflare_zone" "example" {
  zone = "example.com"
  plan = "enterprise"
}`,
			expected: `resource "cloudflare_zone" "example" {
  zone = "example.com"
  plan = "enterprise"
}`,
		},
		{
			name: "preserve data sources",
			input: `data "cloudflare_access_application" "example" {
  name = "example-app"
}`,
			expected: `data "cloudflare_access_application" "example" {
  name = "example-app"
}`,
		},
		{
			name: "rename multiple resources in same file",
			input: `resource "cloudflare_access_application" "app" {
  name = "app"
}

resource "cloudflare_access_policy" "policy" {
  name = "policy"
}

resource "cloudflare_zone" "zone" {
  zone = "example.com"
}`,
			expected: `resource "cloudflare_zero_trust_access_application" "app" {
  name = "app"
}

resource "cloudflare_zero_trust_access_policy" "policy" {
  name = "policy"
}

resource "cloudflare_zone" "zone" {
  zone = "example.com"
}`,
		},
		{
			name: "preserve nested blocks during rename",
			input: `resource "cloudflare_access_application" "example" {
  name   = "example-app"
  domain = "example.com"
  
  cors_headers {
    allowed_methods = ["GET", "POST"]
    allowed_origins = ["https://example.com"]
  }
  
  footer_links {
    name = "Support"
    url  = "https://support.example.com"
  }
}`,
			expected: `resource "cloudflare_zero_trust_access_application" "example" {
  name   = "example-app"
  domain = "example.com"
  cors_headers {
    allowed_methods = ["GET", "POST"]
    allowed_origins = ["https://example.com"]
  }
  footer_links {
    name = "Support"
    url  = "https://support.example.com"
  }
}`,
		},
	}

	// Create transformer
	transformer, err := NewResourceRenameTransformer(configPath)
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

			// Use semantic comparison instead of string comparison
			match, diff, err := CompareHCLResources(tt.expected, string(output))
			if err != nil {
				t.Fatalf("Failed to compare HCL: %v", err)
			}
			if !match {
				t.Errorf("Transformation mismatch: %s\nExpected:\n%s\n\nGot:\n%s", diff, tt.expected, string(output))
			}
		})
	}
}

func TestLoadResourceRenamesConfig(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "resource_config_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
description: "Test config"
notes:
  - "Test note 1"
  - "Test note 2"
resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
  cloudflare_tunnel: cloudflare_zero_trust_tunnel_cloudflared
resource_categories:
  zero_trust_access:
    - cloudflare_access_application
  zero_trust_tunnel:
    - cloudflare_tunnel
migration_guidance:
  description: "Test guidance"
  steps:
    - "Step 1"
    - "Step 2"
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the configuration
	config, err := LoadResourceRenamesConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify the configuration
	if config.Version != "1.0" {
		t.Errorf("Expected version 1.0, got %s", config.Version)
	}

	if len(config.Notes) != 2 {
		t.Errorf("Expected 2 notes, got %d", len(config.Notes))
	}

	if rename, exists := config.ResourceRenames["cloudflare_access_application"]; !exists || rename != "cloudflare_zero_trust_access_application" {
		t.Error("Expected cloudflare_access_application to be renamed")
	}

	if category, exists := config.ResourceCategories["zero_trust_access"]; !exists || len(category) != 1 {
		t.Error("Expected zero_trust_access category with 1 resource")
	}

	if config.MigrationGuidance.Description != "Test guidance" {
		t.Errorf("Expected migration guidance description 'Test guidance', got %s", config.MigrationGuidance.Description)
	}

	if len(config.MigrationGuidance.Steps) != 2 {
		t.Errorf("Expected 2 migration steps, got %d", len(config.MigrationGuidance.Steps))
	}
}

func TestGetResourceRename(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "get_rename_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
  cloudflare_tunnel: cloudflare_zero_trust_tunnel_cloudflared
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	transformer, err := NewResourceRenameTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	tests := []struct {
		resourceType string
		expectRename bool
		expectedName string
	}{
		{"cloudflare_access_application", true, "cloudflare_zero_trust_access_application"},
		{"cloudflare_tunnel", true, "cloudflare_zero_trust_tunnel_cloudflared"},
		{"cloudflare_zone", false, ""},
		{"unknown_resource", false, ""},
	}

	for _, tt := range tests {
		newName, exists := transformer.GetResourceRename(tt.resourceType)
		if exists != tt.expectRename {
			t.Errorf("GetResourceRename(%s) exists = %v, want %v", tt.resourceType, exists, tt.expectRename)
		}
		if exists && newName != tt.expectedName {
			t.Errorf("GetResourceRename(%s) = %s, want %s", tt.resourceType, newName, tt.expectedName)
		}
	}
}

func TestGetResourceCategory(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "get_category_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
resource_categories:
  zero_trust_access:
    - cloudflare_access_application
    - cloudflare_access_policy
  zero_trust_tunnel:
    - cloudflare_tunnel
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	transformer, err := NewResourceRenameTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	tests := []struct {
		resourceType     string
		expectedCategory string
	}{
		{"cloudflare_access_application", "zero_trust_access"},
		{"cloudflare_access_policy", "zero_trust_access"},
		{"cloudflare_tunnel", "zero_trust_tunnel"},
		{"cloudflare_zone", "other"},
		{"unknown_resource", "other"},
	}

	for _, tt := range tests {
		category := transformer.GetResourceCategory(tt.resourceType)
		if category != tt.expectedCategory {
			t.Errorf("GetResourceCategory(%s) = %s, want %s", tt.resourceType, category, tt.expectedCategory)
		}
	}
}

func TestResourceRenameTransformDirectory(t *testing.T) {
	// Create a temporary directory structure
	tempDir, err := os.MkdirTemp("", "resource_rename_dir_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
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
		filepath.Join(tempDir, "main.tf"): `resource "cloudflare_access_application" "main" {
  name = "main-app"
}`,
		filepath.Join(subDir, "module.tf"): `resource "cloudflare_access_application" "module" {
  name = "module-app"
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
	transformer, err := NewResourceRenameTransformer(configPath)
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
	if !strings.Contains(string(mainContent), "cloudflare_zero_trust_access_application") {
		t.Error("main.tf was not transformed")
	}

	// Check module.tf was NOT transformed (non-recursive)
	moduleContent, err := os.ReadFile(filepath.Join(subDir, "module.tf"))
	if err != nil {
		t.Fatalf("Failed to read module.tf: %v", err)
	}
	if strings.Contains(string(moduleContent), "cloudflare_zero_trust_access_application") {
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
	if !strings.Contains(string(moduleContent), "cloudflare_zero_trust_access_application") {
		t.Error("module.tf was not transformed in recursive mode")
	}
}