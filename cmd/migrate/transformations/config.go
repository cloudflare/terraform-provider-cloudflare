package transformations

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// TransformationConfig represents the YAML configuration for transformations
type TransformationConfig struct {
	Version         string                       `yaml:"version"`
	Description     string                       `yaml:"description"`
	Transformations map[string]ResourceTransform `yaml:"transformations"`

	// Fields for attribute renames configuration
	AttributeRenames       map[string]map[string]string `yaml:"attribute_renames"`
	AttributeRemovals      map[string][]string          `yaml:"attribute_removals"`
	ComplexTransformations map[string]interface{}       `yaml:"complex_transformations"`
}

// ResourceTransform defines transformations for a specific resource type
type ResourceTransform struct {
	ToMap             []string          `yaml:"to_map,omitempty"`
	ToList            []string          `yaml:"to_list,omitempty"`
	Notes             string            `yaml:"notes,omitempty"`
	DuplicateHandling map[string]string `yaml:"duplicate_handling,omitempty"`
}

// LoadConfig loads the transformation configuration from a YAML file
func LoadConfig(filename string) (*TransformationConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config TransformationConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}