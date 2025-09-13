package transformations

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// StateResourceRenamesConfig represents the YAML configuration for state resource renames
type StateResourceRenamesConfig struct {
	Version                  string                     `yaml:"version"`
	Description              string                     `yaml:"description"`
	Notes                    []string                   `yaml:"notes"`
	StateResourceRenames     map[string]string          `yaml:"state_resource_renames"`
	DependencyRenamePatterns DependencyRenamePatterns  `yaml:"dependency_rename_patterns"`
	StateStructure           StateStructure             `yaml:"state_structure"`
	SafetyChecks             []string                   `yaml:"safety_checks"`
}

// DependencyRenamePatterns defines patterns for updating dependencies
type DependencyRenamePatterns struct {
	Description string                  `yaml:"description"`
	Patterns    []DependencyPattern     `yaml:"patterns"`
}

// DependencyPattern defines a single dependency rename pattern
type DependencyPattern struct {
	FromPrefix string `yaml:"from_prefix"`
	ToPrefix   string `yaml:"to_prefix"`
}

// StateStructure describes the state file structure
type StateStructure struct {
	Description    string   `yaml:"description"`
	AffectedFields []string `yaml:"affected_fields"`
}

// LoadStateResourceRenamesConfig loads the state resource renames configuration
func LoadStateResourceRenamesConfig(filename string) (*StateResourceRenamesConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config StateResourceRenamesConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}

// StateResourceRenameTransformer handles resource rename transformations for state files
type StateResourceRenameTransformer struct {
	config *StateResourceRenamesConfig
}

// NewStateResourceRenameTransformer creates a new state resource rename transformer
func NewStateResourceRenameTransformer(configPath string) (*StateResourceRenameTransformer, error) {
	config, err := LoadStateResourceRenamesConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &StateResourceRenameTransformer{
		config: config,
	}, nil
}

// TransformStateFile transforms resource types in a Terraform state file
func (srt *StateResourceRenameTransformer) TransformStateFile(inputPath, outputPath string) error {
	// Read the state file
	stateData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read state file: %w", err)
	}

	// Parse the state file
	var state TerraformState
	if err := json.Unmarshal(stateData, &state); err != nil {
		return fmt.Errorf("failed to parse state file: %w", err)
	}

	// Apply resource type renames
	for i := range state.Resources {
		resource := &state.Resources[i]
		
		// Only process managed resources
		if resource.Mode != "managed" {
			continue
		}

		// Check if this resource type needs to be renamed
		if newType, exists := srt.config.StateResourceRenames[resource.Type]; exists {
			resource.Type = newType
		}

		// Update dependencies for all instances
		for j := range resource.Instances {
			instance := &resource.Instances[j]
			srt.updateInstanceDependencies(instance)
		}
	}

	// Write the transformed state file
	output, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}

	if outputPath == "" || outputPath == inputPath {
		outputPath = inputPath
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// updateInstanceDependencies updates dependencies in a resource instance
func (srt *StateResourceRenameTransformer) updateInstanceDependencies(instance *TerraformStateInstance) {
	if instance.Dependencies == nil || len(instance.Dependencies) == 0 {
		return
	}

	// Update each dependency based on rename patterns
	for i, dep := range instance.Dependencies {
		instance.Dependencies[i] = srt.updateDependency(dep)
	}
}

// updateDependency updates a single dependency string based on rename patterns
func (srt *StateResourceRenameTransformer) updateDependency(dep string) string {
	// Apply each dependency rename pattern
	for _, pattern := range srt.config.DependencyRenamePatterns.Patterns {
		if strings.HasPrefix(dep, pattern.FromPrefix) {
			// Replace the prefix
			return pattern.ToPrefix + strings.TrimPrefix(dep, pattern.FromPrefix)
		}
	}
	
	// If no pattern matches, return unchanged
	return dep
}

// GetStateResourceRename checks if a resource type should be renamed in state
func (srt *StateResourceRenameTransformer) GetStateResourceRename(resourceType string) (string, bool) {
	newType, exists := srt.config.StateResourceRenames[resourceType]
	return newType, exists
}

// TransformStateFileWithBackup creates a backup before transforming
func (srt *StateResourceRenameTransformer) TransformStateFileWithBackup(statePath string) error {
	// Create backup
	backupPath := statePath + ".backup"
	stateData, err := os.ReadFile(statePath)
	if err != nil {
		return fmt.Errorf("failed to read state file for backup: %w", err)
	}
	
	if err := os.WriteFile(backupPath, stateData, 0644); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}
	
	fmt.Printf("Created backup at: %s\n", backupPath)
	
	// Transform the state file
	return srt.TransformStateFile(statePath, statePath)
}

// ValidateStateFile performs basic validation on a state file
func ValidateStateFile(statePath string) error {
	stateData, err := os.ReadFile(statePath)
	if err != nil {
		return fmt.Errorf("failed to read state file: %w", err)
	}

	var state TerraformState
	if err := json.Unmarshal(stateData, &state); err != nil {
		return fmt.Errorf("invalid state file format: %w", err)
	}

	// Basic validation
	if state.Version == 0 {
		return fmt.Errorf("invalid state version: %d", state.Version)
	}

	return nil
}

// CombinedStateTransformer combines attribute renames and resource renames
type CombinedStateTransformer struct {
	attributeConfig *StateTransformationConfig
	resourceConfig  *StateResourceRenamesConfig
}

// NewCombinedStateTransformer creates a transformer that handles both attribute and resource renames
func NewCombinedStateTransformer(attributeConfigPath, resourceConfigPath string) (*CombinedStateTransformer, error) {
	var transformer CombinedStateTransformer
	var err error

	if attributeConfigPath != "" {
		transformer.attributeConfig, err = LoadStateConfig(attributeConfigPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load attribute config: %w", err)
		}
	}

	if resourceConfigPath != "" {
		transformer.resourceConfig, err = LoadStateResourceRenamesConfig(resourceConfigPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load resource config: %w", err)
		}
	}

	return &transformer, nil
}

// TransformStateFile applies both attribute and resource transformations
func (cst *CombinedStateTransformer) TransformStateFile(inputPath, outputPath string) error {
	// Read the state file
	stateData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read state file: %w", err)
	}

	// Parse the state file
	var state TerraformState
	if err := json.Unmarshal(stateData, &state); err != nil {
		return fmt.Errorf("failed to parse state file: %w", err)
	}

	// Apply transformations to each resource
	for i := range state.Resources {
		resource := &state.Resources[i]
		
		// Only process managed resources
		if resource.Mode != "managed" {
			continue
		}

		// Check if this is a Cloudflare resource
		if !isCloudflareResource(resource.Type) && !isCloudflareResource(cst.getOriginalType(resource.Type)) {
			continue
		}

		// Apply resource type rename first
		originalType := resource.Type
		if cst.resourceConfig != nil {
			if newType, exists := cst.resourceConfig.StateResourceRenames[resource.Type]; exists {
				resource.Type = newType
			}
		}

		// Apply attribute transformations using the original type for lookups
		if cst.attributeConfig != nil {
			// Apply schema version reset
			if cst.attributeConfig.SchemaVersionReset.AllCloudflareResources {
				for j := range resource.Instances {
					resource.Instances[j].SchemaVersion = 0
				}
			}

			// Apply attribute renames (use original type for config lookup)
			if renames, exists := cst.attributeConfig.StateAttributeRenames[originalType]; exists {
				for j := range resource.Instances {
					applyStateAttributeRenames(&resource.Instances[j], renames)
				}
			}

			// Apply attribute removals
			if removals, exists := cst.attributeConfig.StateAttributeRemovals[originalType]; exists {
				for j := range resource.Instances {
					applyStateAttributeRemovals(&resource.Instances[j], removals)
				}
			}

			// Apply special transformations
			if special, exists := cst.attributeConfig.StateSpecialTransformations[originalType]; exists {
				for j := range resource.Instances {
					applySpecialTransformations(&resource.Instances[j], special)
				}
			}
		}

		// Update dependencies
		if cst.resourceConfig != nil {
			for j := range resource.Instances {
				cst.updateInstanceDependencies(&resource.Instances[j])
			}
		}
	}

	// Write the transformed state file
	output, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}

	if outputPath == "" || outputPath == inputPath {
		outputPath = inputPath
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// getOriginalType returns the original type if it was renamed, otherwise returns the input
func (cst *CombinedStateTransformer) getOriginalType(resourceType string) string {
	if cst.resourceConfig == nil {
		return resourceType
	}
	
	// Reverse lookup to find original type
	for oldType, newType := range cst.resourceConfig.StateResourceRenames {
		if newType == resourceType {
			return oldType
		}
	}
	
	return resourceType
}

// updateInstanceDependencies updates dependencies using resource config
func (cst *CombinedStateTransformer) updateInstanceDependencies(instance *TerraformStateInstance) {
	if cst.resourceConfig == nil || instance.Dependencies == nil || len(instance.Dependencies) == 0 {
		return
	}

	// Update each dependency based on rename patterns
	for i, dep := range instance.Dependencies {
		for _, pattern := range cst.resourceConfig.DependencyRenamePatterns.Patterns {
			if strings.HasPrefix(dep, pattern.FromPrefix) {
				instance.Dependencies[i] = pattern.ToPrefix + strings.TrimPrefix(dep, pattern.FromPrefix)
				break
			}
		}
	}
}