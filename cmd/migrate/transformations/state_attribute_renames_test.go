package transformations

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestStateTransformer(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "state_transformer_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file
	configPath := filepath.Join(tempDir, "test_state_config.yaml")
	configContent := `
version: "1.0"
description: "Test state transformations"
schema_version_reset:
  all_cloudflare_resources: true
state_attribute_renames:
  cloudflare_api_token:
    policy: policies
  cloudflare_account_member:
    role_ids: roles
    email_address: email
  cloudflare_queue:
    name: queue_name
    id_duplicate_as: queue_id
  cloudflare_turnstile_widget:
    id_duplicate_as: sitekey
state_attribute_removals:
  cloudflare_access_policy:
    - application_id
    - precedence
  cloudflare_tunnel:
    - secret
    - cname
  cloudflare_page_rule:
    - disable_railgun
state_special_transformations:
  cloudflare_page_rule:
    empty_to_null:
      - '""'
      - '[]'
      - 'false'
      - '0'
    field_transformations:
      forwarding_url:
        from: '[]'
        to: 'null'
      actions:
        from: '[$action]'
        to: '$action'
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test cases
	tests := []struct {
		name     string
		input    TerraformState
		expected TerraformState
	}{
		{
			name: "schema version reset",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_api_token",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 2,
								Attributes:    map[string]interface{}{},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_api_token",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes:    map[string]interface{}{},
							},
						},
					},
				},
			},
		},
		{
			name: "attribute renames",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_api_token",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"policy": "read-only",
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_api_token",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"policies": "read-only",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "multiple attribute renames",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_account_member",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"role_ids":      []interface{}{"admin", "user"},
									"email_address": "test@example.com",
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_account_member",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"roles": []interface{}{"admin", "user"},
									"email": "test@example.com",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "id duplication",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_queue",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"id":   "queue-123",
									"name": "my-queue",
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_queue",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"id":         "queue-123",
									"queue_id":   "queue-123",
									"queue_name": "my-queue",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "attribute removals",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_access_policy",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name":           "test-policy",
									"application_id": "app-123",
									"precedence":     1,
									"decision":       "allow",
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_access_policy",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"name":     "test-policy",
									"decision": "allow",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "page rule empty to null",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"target":         "example.com/*",
									"forwarding_url": []interface{}{},
									"cache_level":    "",
									"ssl":            false,
									"priority":       0,
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"target":         "example.com/*",
									"forwarding_url": nil,
									"cache_level":    nil,
									"ssl":            nil,
									"priority":       nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "page rule unwrap single element",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"forwarding_url": []interface{}{
										map[string]interface{}{
											"status_code": 301,
											"url":         "https://new.example.com",
										},
									},
									"actions": []interface{}{
										map[string]interface{}{
											"cache_level": "aggressive",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"forwarding_url": map[string]interface{}{
										"status_code": 301,
										"url":         "https://new.example.com",
									},
									"actions": map[string]interface{}{
										"cache_level": "aggressive",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "page rule remove disable_railgun attribute",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"disable_railgun": true,
									"actions": map[string]interface{}{
										"cache_level": "aggressive",
										"disable_railgun": "on",
										"ssl": "flexible",
									},
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"actions": map[string]interface{}{
										"cache_level": "aggressive",
										"ssl": "flexible",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "page rule forwarding_url unwrap in actions",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"actions": map[string]interface{}{
										"forwarding_url": []interface{}{
											map[string]interface{}{
												"status_code": 301,
												"url": "https://new.example.com",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"actions": map[string]interface{}{
										"forwarding_url": map[string]interface{}{
											"status_code": 301,
											"url": "https://new.example.com",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "page rule browser_cache_ttl empty string to null",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"actions": map[string]interface{}{
										"cache_level": "aggressive",
										"browser_cache_ttl": "",
										"edge_cache_ttl": "0",
										"ssl": "flexible",
									},
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_page_rule",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 0,
								Attributes: map[string]interface{}{
									"target": "example.com/*",
									"actions": map[string]interface{}{
										"cache_level": "aggressive",
										"browser_cache_ttl": nil,
										"edge_cache_ttl": nil,
										"ssl": "flexible",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "non-cloudflare resources unchanged",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "aws_instance",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"instance_type": "t2.micro",
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "aws_instance",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"instance_type": "t2.micro",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "data sources unchanged",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "data",
						Type: "cloudflare_zone",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name": "example.com",
								},
							},
						},
					},
				},
			},
			expected: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "data",
						Type: "cloudflare_zone",
						Name: "test",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name": "example.com",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create transformer
	transformer, err := NewStateTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input file
			inputPath := filepath.Join(tempDir, "input.tfstate")
			inputData, err := json.MarshalIndent(tt.input, "", "  ")
			if err != nil {
				t.Fatalf("Failed to marshal input: %v", err)
			}
			if err := os.WriteFile(inputPath, inputData, 0644); err != nil {
				t.Fatalf("Failed to write input file: %v", err)
			}

			// Transform the file
			outputPath := filepath.Join(tempDir, "output.tfstate")
			if err := transformer.TransformFile(inputPath, outputPath); err != nil {
				t.Fatalf("Failed to transform file: %v", err)
			}

			// Read the output
			outputData, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			var output TerraformState
			if err := json.Unmarshal(outputData, &output); err != nil {
				t.Fatalf("Failed to parse output: %v", err)
			}

			// Compare the output with expected
			if !compareStates(t, output, tt.expected) {
				expectedJSON, _ := json.MarshalIndent(tt.expected, "", "  ")
				outputJSON, _ := json.MarshalIndent(output, "", "  ")
				t.Errorf("State mismatch\nExpected:\n%s\n\nGot:\n%s", expectedJSON, outputJSON)
			}
		})
	}
}

func TestLoadStateConfig(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "state_config_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
description: "Test config"
schema_version_reset:
  all_cloudflare_resources: true
state_attribute_renames:
  cloudflare_api_token:
    policy: policies
state_attribute_removals:
  cloudflare_access_policy:
    - application_id
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the configuration
	config, err := LoadStateConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify the configuration
	if config.Version != "1.0" {
		t.Errorf("Expected version 1.0, got %s", config.Version)
	}

	if !config.SchemaVersionReset.AllCloudflareResources {
		t.Error("Expected AllCloudflareResources to be true")
	}

	if rename, exists := config.StateAttributeRenames["cloudflare_api_token"]["policy"]; !exists || rename != "policies" {
		t.Error("Expected policy to be renamed to policies")
	}

	if removals, exists := config.StateAttributeRemovals["cloudflare_access_policy"]; !exists || len(removals) != 1 || removals[0] != "application_id" {
		t.Error("Expected application_id to be in removals")
	}
}

func TestIsCloudflareResource(t *testing.T) {
	tests := []struct {
		resourceType string
		expected     bool
	}{
		{"cloudflare_zone", true},
		{"cloudflare_access_policy", true},
		{"aws_instance", false},
		{"google_compute_instance", false},
		{"cloudflare", false}, // Too short
		{"", false},
	}

	for _, tt := range tests {
		result := isCloudflareResource(tt.resourceType)
		if result != tt.expected {
			t.Errorf("isCloudflareResource(%s) = %v, want %v", tt.resourceType, result, tt.expected)
		}
	}
}

func TestShouldConvertToNull(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		pattern  string
		expected bool
	}{
		{"empty string", "", `""`, true},
		{"non-empty string", "value", `""`, false},
		{"empty array", []interface{}{}, "[]", true},
		{"non-empty array", []interface{}{"item"}, "[]", false},
		{"false bool", false, "false", true},
		{"true bool", true, "false", false},
		{"zero float", float64(0), "0", true},
		{"non-zero float", float64(1), "0", false},
		{"zero int", 0, "0", true},
		{"non-zero int", 1, "0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shouldConvertToNull(tt.value, tt.pattern)
			if result != tt.expected {
				t.Errorf("shouldConvertToNull(%v, %s) = %v, want %v", tt.value, tt.pattern, result, tt.expected)
			}
		})
	}
}

// compareStates compares two Terraform states for equality
func compareStates(t *testing.T, a, b TerraformState) bool {
	if a.Version != b.Version {
		return false
	}

	if len(a.Resources) != len(b.Resources) {
		return false
	}

	for i := range a.Resources {
		if !compareResources(t, a.Resources[i], b.Resources[i]) {
			return false
		}
	}

	return true
}

// compareResources compares two Terraform state resources
func compareResources(t *testing.T, a, b TerraformStateResource) bool {
	if a.Mode != b.Mode || a.Type != b.Type || a.Name != b.Name {
		return false
	}

	if len(a.Instances) != len(b.Instances) {
		return false
	}

	for i := range a.Instances {
		if !compareInstances(t, a.Instances[i], b.Instances[i]) {
			return false
		}
	}

	return true
}

// compareInstances compares two Terraform state instances
func compareInstances(t *testing.T, a, b TerraformStateInstance) bool {
	if a.SchemaVersion != b.SchemaVersion {
		return false
	}

	// Compare attributes
	if len(a.Attributes) != len(b.Attributes) {
		return false
	}

	for key, valA := range a.Attributes {
		valB, exists := b.Attributes[key]
		if !exists {
			return false
		}

		// Handle nil comparison
		if valA == nil && valB == nil {
			continue
		}
		if valA == nil || valB == nil {
			return false
		}

		// Convert to JSON for deep comparison
		jsonA, _ := json.Marshal(valA)
		jsonB, _ := json.Marshal(valB)
		if string(jsonA) != string(jsonB) {
			return false
		}
	}

	return true
}