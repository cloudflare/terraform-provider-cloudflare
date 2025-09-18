package transformations

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// TestPageRuleStatusDefault tests that status = "active" is added when not present
func TestPageRuleStatusDefault(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "page_rule_status_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file - minimal config just for page_rule
	configPath := filepath.Join(tempDir, "test_config.yaml")
	configContent := `
version: "1.0"
description: "Test page rule status default"
attribute_renames: {}
attribute_removals: {}
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
			name: "page_rule without status attribute should get status = active",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "12345"
  target  = "example.com/*"
  priority = 1

  actions {
    cache_level = "bypass"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "12345"
  target  = "example.com/*"
  priority = 1
  status = "active"

  actions {
    cache_level = "bypass"
  }
}`,
		},
		{
			name: "page_rule with existing status should be unchanged",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "12345"
  target  = "example.com/*"
  priority = 1
  status = "disabled"

  actions {
    cache_level = "bypass"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "12345"
  target  = "example.com/*"
  priority = 1
  status = "disabled"

  actions {
    cache_level = "bypass"
  }
}`,
		},
		{
			name: "multiple page_rules without status",
			input: `resource "cloudflare_page_rule" "rule1" {
  zone_id = "12345"
  target  = "example.com/api/*"
  priority = 1

  actions {
    cache_level = "bypass"
  }
}

resource "cloudflare_page_rule" "rule2" {
  zone_id = "12345"
  target  = "example.com/images/*"
  priority = 2

  actions {
    cache_level = "aggressive"
  }
}`,
			expected: `resource "cloudflare_page_rule" "rule1" {
  zone_id = "12345"
  target  = "example.com/api/*"
  priority = 1
  status = "active"

  actions {
    cache_level = "bypass"
  }
}

resource "cloudflare_page_rule" "rule2" {
  zone_id = "12345"
  target  = "example.com/images/*"
  priority = 2
  status = "active"

  actions {
    cache_level = "aggressive"
  }
}`,
		},
		{
			name: "non-page_rule resources should not get status",
			input: `resource "cloudflare_zone" "example" {
  zone = "example.com"
  plan = "free"
}`,
			expected: `resource "cloudflare_zone" "example" {
  zone = "example.com"
  plan = "free"
}`,
		},
	}

	// Create a transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write test input to file
			testFile := filepath.Join(tempDir, "test.tf")
			if err := os.WriteFile(testFile, []byte(tt.input), 0644); err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			// Transform the file
			if err := transformer.TransformFile(testFile, testFile); err != nil {
				t.Fatalf("Failed to transform file: %v", err)
			}

			// Read the result
			result, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("Failed to read result file: %v", err)
			}

			// Parse and compare
			expectedFile, diags := hclwrite.ParseConfig([]byte(tt.expected), "expected.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse expected config: %v", diags)
			}

			actualFile, diags := hclwrite.ParseConfig(result, "actual.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse actual config: %v", diags)
			}

			// Compare the formatted outputs
			actualFormatted := string(hclwrite.Format(actualFile.Bytes()))

			// Check that expected content is in actual (because actual may have additional formatting)
			actualLines := strings.Split(strings.TrimSpace(actualFormatted), "\n")

			// For page_rule tests, specifically check for status attribute
			if strings.Contains(tt.input, "cloudflare_page_rule") && !strings.Contains(tt.input, "status") {
				// Should have added status = "active"
				found := false
				for _, line := range actualLines {
					if strings.Contains(line, `status = "active"`) || strings.Contains(line, `status  = "active"`) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected status = \"active\" to be added, but it wasn't found in output:\n%s", actualFormatted)
				}
			}

			// For page_rules with existing status, check it wasn't changed
			if strings.Contains(tt.input, `status = "disabled"`) {
				found := false
				for _, line := range actualLines {
					if strings.Contains(line, `status`) && strings.Contains(line, `"disabled"`) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected existing status = \"disabled\" to be preserved, but it wasn't found in output:\n%s", actualFormatted)
				}
			}

			// For non-page_rule resources, ensure status wasn't added
			if !strings.Contains(tt.input, "cloudflare_page_rule") {
				for _, line := range actualLines {
					if strings.Contains(line, "status") {
						t.Errorf("Status attribute should not be added to non-page_rule resources, but found in output:\n%s", actualFormatted)
					}
				}
			}

			// General structure check - ensure we have the same number of resource blocks
			expectedBlocks := expectedFile.Body().Blocks()
			actualBlocks := actualFile.Body().Blocks()
			if len(expectedBlocks) != len(actualBlocks) {
				t.Errorf("Block count mismatch: expected %d, got %d", len(expectedBlocks), len(actualBlocks))
			}
		})
	}
}

// TestPageRuleStatusWithRenames tests that status default works alongside other renames
func TestPageRuleStatusWithRenames(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "page_rule_combined_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file with some page_rule renames
	configPath := filepath.Join(tempDir, "test_config.yaml")
	configContent := `
version: "1.0"
description: "Test page rule with other transformations"
attribute_renames:
  cloudflare_page_rule:
    old_attr: new_attr
attribute_removals:
  cloudflare_page_rule:
    - remove_me
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test that status default works with other transformations
	input := `resource "cloudflare_page_rule" "example" {
  zone_id = "12345"
  target = "example.com/*"
  priority = 1
  old_attr = "value"
  remove_me = "should be gone"

  actions {
    cache_level = "bypass"
  }
}`

	// Create a transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	// Write test input to file
	testFile := filepath.Join(tempDir, "test.tf")
	if err := os.WriteFile(testFile, []byte(input), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Transform the file
	if err := transformer.TransformFile(testFile, testFile); err != nil {
		t.Fatalf("Failed to transform file: %v", err)
	}

	// Read the result
	result, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read result file: %v", err)
	}

	// Parse the result
	actualFile, diags := hclwrite.ParseConfig(result, "actual.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Failed to parse actual config: %v", diags)
	}

	actualFormatted := string(hclwrite.Format(actualFile.Bytes()))

	// Check that status was added
	statusFound := false
	lines := strings.Split(actualFormatted, "\n")
	for _, line := range lines {
		if strings.Contains(line, "status") && strings.Contains(line, `"active"`) {
			statusFound = true
			break
		}
	}
	if !statusFound {
		t.Errorf("Expected status = \"active\" to be added, but it wasn't found in output:\n%s", actualFormatted)
	}

	// Check that rename happened
	if !strings.Contains(actualFormatted, "new_attr") {
		t.Errorf("Expected old_attr to be renamed to new_attr, but new_attr wasn't found in output:\n%s", actualFormatted)
	}

	// Check that removal happened
	if strings.Contains(actualFormatted, "remove_me") {
		t.Errorf("Expected remove_me attribute to be removed, but it was found in output:\n%s", actualFormatted)
	}

	// Check that old_attr is gone
	if strings.Contains(actualFormatted, "old_attr") {
		t.Errorf("Expected old_attr to be renamed, but old_attr was still found in output:\n%s", actualFormatted)
	}
}