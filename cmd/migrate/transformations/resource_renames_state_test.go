package transformations

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestStateResourceRenameTransformer(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "state_resource_rename_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file
	configPath := filepath.Join(tempDir, "test_state_resource_config.yaml")
	configContent := `
version: "1.0"
description: "Test state resource renames"
state_resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
  cloudflare_access_policy: cloudflare_zero_trust_access_policy
  cloudflare_teams_account: cloudflare_zero_trust_gateway_settings
  cloudflare_tunnel: cloudflare_zero_trust_tunnel_cloudflared
  cloudflare_managed_headers: cloudflare_managed_transforms
dependency_rename_patterns:
  description: "Test dependency patterns"
  patterns:
    - from_prefix: "cloudflare_access_application."
      to_prefix: "cloudflare_zero_trust_access_application."
    - from_prefix: "cloudflare_access_policy."
      to_prefix: "cloudflare_zero_trust_access_policy."
    - from_prefix: "cloudflare_tunnel."
      to_prefix: "cloudflare_zero_trust_tunnel_cloudflared."
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
			name: "rename access application resource",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode:     "managed",
						Type:     "cloudflare_access_application",
						Name:     "example",
						Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name": "example-app",
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
						Mode:     "managed",
						Type:     "cloudflare_zero_trust_access_application",
						Name:     "example",
						Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name": "example-app",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "rename teams account resource",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_teams_account",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"account_id": "123456",
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
						Type: "cloudflare_zero_trust_gateway_settings",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"account_id": "123456",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "update dependencies when renaming",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_access_policy",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name": "example-policy",
								},
								Dependencies: []string{
									"cloudflare_access_application.example",
									"cloudflare_zone.example",
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
						Type: "cloudflare_zero_trust_access_policy",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
								SchemaVersion: 1,
								Attributes: map[string]interface{}{
									"name": "example-policy",
								},
								Dependencies: []string{
									"cloudflare_zero_trust_access_application.example",
									"cloudflare_zone.example",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "rename multiple resources",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "cloudflare_access_application",
						Name: "app",
						Instances: []TerraformStateInstance{
							{
								Attributes: map[string]interface{}{
									"name": "app",
								},
							},
						},
					},
					{
						Mode: "managed",
						Type: "cloudflare_tunnel",
						Name: "tunnel",
						Instances: []TerraformStateInstance{
							{
								Attributes: map[string]interface{}{
									"name": "tunnel",
								},
								Dependencies: []string{
									"cloudflare_access_application.app",
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
						Type: "cloudflare_zero_trust_access_application",
						Name: "app",
						Instances: []TerraformStateInstance{
							{
								Attributes: map[string]interface{}{
									"name": "app",
								},
							},
						},
					},
					{
						Mode: "managed",
						Type: "cloudflare_zero_trust_tunnel_cloudflared",
						Name: "tunnel",
						Instances: []TerraformStateInstance{
							{
								Attributes: map[string]interface{}{
									"name": "tunnel",
								},
								Dependencies: []string{
									"cloudflare_zero_trust_access_application.app",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "preserve non-cloudflare resources",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "managed",
						Type: "aws_instance",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
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
						Name: "example",
						Instances: []TerraformStateInstance{
							{
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
			name: "preserve data sources",
			input: TerraformState{
				Version: 4,
				Resources: []TerraformStateResource{
					{
						Mode: "data",
						Type: "cloudflare_access_application",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
								Attributes: map[string]interface{}{
									"name": "example-app",
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
						Type: "cloudflare_access_application",
						Name: "example",
						Instances: []TerraformStateInstance{
							{
								Attributes: map[string]interface{}{
									"name": "example-app",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create transformer
	transformer, err := NewStateResourceRenameTransformer(configPath)
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
			if err := transformer.TransformStateFile(inputPath, outputPath); err != nil {
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

func TestCombinedStateTransformer(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "combined_state_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create attribute renames config
	attrConfigPath := filepath.Join(tempDir, "attr_config.yaml")
	attrConfigContent := `
version: "1.0"
schema_version_reset:
  all_cloudflare_resources: true
state_attribute_renames:
  cloudflare_access_application:
    domain: domain_name
  cloudflare_access_policy:
    application_id: app_id
state_attribute_removals:
  cloudflare_access_policy:
    - precedence
`
	if err := os.WriteFile(attrConfigPath, []byte(attrConfigContent), 0644); err != nil {
		t.Fatalf("Failed to write attribute config file: %v", err)
	}

	// Create resource renames config
	resourceConfigPath := filepath.Join(tempDir, "resource_config.yaml")
	resourceConfigContent := `
version: "1.0"
state_resource_renames:
  cloudflare_access_application: cloudflare_zero_trust_access_application
  cloudflare_access_policy: cloudflare_zero_trust_access_policy
dependency_rename_patterns:
  patterns:
    - from_prefix: "cloudflare_access_application."
      to_prefix: "cloudflare_zero_trust_access_application."
`
	if err := os.WriteFile(resourceConfigPath, []byte(resourceConfigContent), 0644); err != nil {
		t.Fatalf("Failed to write resource config file: %v", err)
	}

	// Test case: combined transformations
	input := TerraformState{
		Version: 4,
		Resources: []TerraformStateResource{
			{
				Mode: "managed",
				Type: "cloudflare_access_application",
				Name: "app",
				Instances: []TerraformStateInstance{
					{
						SchemaVersion: 2,
						Attributes: map[string]interface{}{
							"name":   "app",
							"domain": "example.com",
						},
					},
				},
			},
			{
				Mode: "managed",
				Type: "cloudflare_access_policy",
				Name: "policy",
				Instances: []TerraformStateInstance{
					{
						SchemaVersion: 1,
						Attributes: map[string]interface{}{
							"name":           "policy",
							"application_id": "app-123",
							"precedence":     1,
						},
						Dependencies: []string{
							"cloudflare_access_application.app",
						},
					},
				},
			},
		},
	}

	expected := TerraformState{
		Version: 4,
		Resources: []TerraformStateResource{
			{
				Mode: "managed",
				Type: "cloudflare_zero_trust_access_application",
				Name: "app",
				Instances: []TerraformStateInstance{
					{
						SchemaVersion: 0, // Reset by attribute config
						Attributes: map[string]interface{}{
							"name":        "app",
							"domain_name": "example.com", // Renamed by attribute config
						},
					},
				},
			},
			{
				Mode: "managed",
				Type: "cloudflare_zero_trust_access_policy",
				Name: "policy",
				Instances: []TerraformStateInstance{
					{
						SchemaVersion: 0, // Reset by attribute config
						Attributes: map[string]interface{}{
							"name":   "policy",
							"app_id": "app-123", // Renamed by attribute config
							// precedence removed by attribute config
						},
						Dependencies: []string{
							"cloudflare_zero_trust_access_application.app", // Updated by resource config
						},
					},
				},
			},
		},
	}

	// Create combined transformer
	transformer, err := NewCombinedStateTransformer(attrConfigPath, resourceConfigPath)
	if err != nil {
		t.Fatalf("Failed to create combined transformer: %v", err)
	}

	// Create input file
	inputPath := filepath.Join(tempDir, "input.tfstate")
	inputData, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal input: %v", err)
	}
	if err := os.WriteFile(inputPath, inputData, 0644); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	// Transform the file
	outputPath := filepath.Join(tempDir, "output.tfstate")
	if err := transformer.TransformStateFile(inputPath, outputPath); err != nil {
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
	if !compareStates(t, output, expected) {
		expectedJSON, _ := json.MarshalIndent(expected, "", "  ")
		outputJSON, _ := json.MarshalIndent(output, "", "  ")
		t.Errorf("State mismatch\nExpected:\n%s\n\nGot:\n%s", expectedJSON, outputJSON)
	}
}

func TestLoadStateResourceRenamesConfig(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "state_resource_config_test")
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
  - "Note 1"
  - "Note 2"
state_resource_renames:
  cloudflare_tunnel: cloudflare_zero_trust_tunnel_cloudflared
dependency_rename_patterns:
  description: "Test patterns"
  patterns:
    - from_prefix: "cloudflare_tunnel."
      to_prefix: "cloudflare_zero_trust_tunnel_cloudflared."
safety_checks:
  - "Check 1"
  - "Check 2"
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the configuration
	config, err := LoadStateResourceRenamesConfig(configPath)
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

	if rename, exists := config.StateResourceRenames["cloudflare_tunnel"]; !exists || rename != "cloudflare_zero_trust_tunnel_cloudflared" {
		t.Error("Expected cloudflare_tunnel to be renamed")
	}

	if len(config.DependencyRenamePatterns.Patterns) != 1 {
		t.Errorf("Expected 1 dependency pattern, got %d", len(config.DependencyRenamePatterns.Patterns))
	}

	if len(config.SafetyChecks) != 2 {
		t.Errorf("Expected 2 safety checks, got %d", len(config.SafetyChecks))
	}
}

func TestValidateStateFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "validate_state_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test valid state file
	validState := TerraformState{
		Version: 4,
		Resources: []TerraformStateResource{
			{
				Mode: "managed",
				Type: "cloudflare_zone",
				Name: "example",
			},
		},
	}

	validPath := filepath.Join(tempDir, "valid.tfstate")
	validData, _ := json.MarshalIndent(validState, "", "  ")
	if err := os.WriteFile(validPath, validData, 0644); err != nil {
		t.Fatalf("Failed to write valid state file: %v", err)
	}

	if err := ValidateStateFile(validPath); err != nil {
		t.Errorf("ValidateStateFile failed for valid state: %v", err)
	}

	// Test invalid state file (version 0)
	invalidState := TerraformState{
		Version: 0,
	}

	invalidPath := filepath.Join(tempDir, "invalid.tfstate")
	invalidData, _ := json.MarshalIndent(invalidState, "", "  ")
	if err := os.WriteFile(invalidPath, invalidData, 0644); err != nil {
		t.Fatalf("Failed to write invalid state file: %v", err)
	}

	if err := ValidateStateFile(invalidPath); err == nil {
		t.Error("ValidateStateFile should have failed for invalid state")
	}

	// Test non-JSON file
	nonJSONPath := filepath.Join(tempDir, "notjson.tfstate")
	if err := os.WriteFile(nonJSONPath, []byte("not json"), 0644); err != nil {
		t.Fatalf("Failed to write non-JSON file: %v", err)
	}

	if err := ValidateStateFile(nonJSONPath); err == nil {
		t.Error("ValidateStateFile should have failed for non-JSON file")
	}
}

func TestTransformStateFileWithBackup(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "backup_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create config
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
state_resource_renames:
  cloudflare_tunnel: cloudflare_zero_trust_tunnel_cloudflared
dependency_rename_patterns:
  patterns: []
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Create state file
	state := TerraformState{
		Version: 4,
		Resources: []TerraformStateResource{
			{
				Mode: "managed",
				Type: "cloudflare_tunnel",
				Name: "example",
			},
		},
	}

	statePath := filepath.Join(tempDir, "terraform.tfstate")
	stateData, _ := json.MarshalIndent(state, "", "  ")
	if err := os.WriteFile(statePath, stateData, 0644); err != nil {
		t.Fatalf("Failed to write state file: %v", err)
	}

	// Create transformer
	transformer, err := NewStateResourceRenameTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	// Transform with backup
	if err := transformer.TransformStateFileWithBackup(statePath); err != nil {
		t.Fatalf("Failed to transform with backup: %v", err)
	}

	// Check backup exists
	backupPath := statePath + ".backup"
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		t.Error("Backup file was not created")
	}

	// Check backup contents match original
	backupData, err := os.ReadFile(backupPath)
	if err != nil {
		t.Fatalf("Failed to read backup: %v", err)
	}

	var backupState TerraformState
	if err := json.Unmarshal(backupData, &backupState); err != nil {
		t.Fatalf("Failed to parse backup: %v", err)
	}

	if backupState.Resources[0].Type != "cloudflare_tunnel" {
		t.Error("Backup should contain original resource type")
	}

	// Check transformed file
	transformedData, err := os.ReadFile(statePath)
	if err != nil {
		t.Fatalf("Failed to read transformed file: %v", err)
	}

	var transformedState TerraformState
	if err := json.Unmarshal(transformedData, &transformedState); err != nil {
		t.Fatalf("Failed to parse transformed file: %v", err)
	}

	if transformedState.Resources[0].Type != "cloudflare_zero_trust_tunnel_cloudflared" {
		t.Error("State file should be transformed")
	}
}