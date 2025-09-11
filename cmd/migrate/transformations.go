package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/transformations"
)

// runYAMLTransformations runs the YAML-based transformations for config and state files
func runYAMLTransformations(configDir, stateFile, transformerDir string, dryRun bool) error {
	// List of transformation configs to apply
	transformationConfigs := []struct {
		configFile  string
		target      string // "config" or "state"
		description string
	}{
		// Configuration file transformations
		{"cloudflare_terraform_v5_block_to_attribute_configuration.yaml", "config", "Block to attribute transformations"},
		{"cloudflare_terraform_v5_attribute_renames_configuration.yaml", "config", "Configuration attribute renames"},
		{"cloudflare_terraform_v5_resource_renames_configuration.yaml", "config", "Configuration resource renames"},

		// State file transformations
		{"cloudflare_terraform_v5_attribute_renames_state.yaml", "state", "State attribute renames"},
		{"cloudflare_terraform_v5_resource_renames_state.yaml", "state", "State resource renames"},
	}

	// If transformerDir is a URL, we need to handle it differently
	// For now, we'll use local files
	if strings.HasPrefix(transformerDir, "http") {
		// For embedded configs, use the local transformations/config directory
		// This should be packaged with the binary or fetched from GitHub
		execPath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("failed to get executable path: %v", err)
		}
		transformerDir = filepath.Join(filepath.Dir(execPath), "transformations", "config")

		// Check if the directory exists
		if _, err := os.Stat(transformerDir); os.IsNotExist(err) {
			// Try relative to the current working directory
			transformerDir = filepath.Join("transformations", "config")
			if _, err := os.Stat(transformerDir); os.IsNotExist(err) {
				return fmt.Errorf("transformer config directory not found. Please specify --transformer-dir")
			}
		}
	}

	// Apply each transformation
	for _, t := range transformationConfigs {
		// Skip config transformations if no config directory is specified
		if t.target == "config" && (configDir == "" || configDir == "false") {
			fmt.Printf("  Skipping %s (config transformations disabled)\n", t.description)
			continue
		}

		// Skip state transformations if no state file is specified
		if t.target == "state" && (stateFile == "" || stateFile == "false") {
			fmt.Printf("  Skipping %s (state transformations disabled)\n", t.description)
			continue
		}

		configPath := filepath.Join(transformerDir, t.configFile)

		// Check if the config file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("  ⚠ Warning: Config file %s not found, skipping %s\n", t.configFile, t.description)
			continue
		}

		fmt.Printf("  Applying %s...\n", t.description)

		if t.target == "config" {
			// Apply configuration transformations
			if err := applyConfigTransformation(configDir, configPath, t.configFile, dryRun); err != nil {
				fmt.Printf("    ⚠ Warning: Failed to apply %s: %v\n", t.description, err)
				continue
			}
		} else if t.target == "state" {
			// Apply state transformations
			if err := applyStateTransformation(stateFile, configPath, t.configFile, dryRun); err != nil {
				fmt.Printf("    ⚠ Warning: Failed to apply %s: %v\n", t.description, err)
				continue
			}
		}
	}

	return nil
}

// applyConfigTransformation applies a specific transformation to configuration files
func applyConfigTransformation(configDir, configPath, configFile string, dryRun bool) error {
	// Determine which transformer to use based on the config file name
	switch {
	case strings.Contains(configFile, "block_to_attribute"):
		// Use HCL transformer for block-to-attribute transformations
		transformer, err := transformations.NewHCLTransformer(configPath)
		if err != nil {
			return fmt.Errorf("failed to create HCL transformer: %v", err)
		}
		if dryRun {
			fmt.Printf("    ✗ Would transform config files in %s\n", configDir)
			return nil
		}
		return transformer.TransformDirectory(configDir, false)

	case strings.Contains(configFile, "attribute_renames_configuration"):
		// Use HCL transformer for attribute renames
		transformer, err := transformations.NewHCLTransformer(configPath)
		if err != nil {
			return fmt.Errorf("failed to create HCL transformer: %v", err)
		}
		if dryRun {
			fmt.Printf("    ✗ Would transform config files in %s\n", configDir)
			return nil
		}
		return transformer.TransformDirectory(configDir, false)

	case strings.Contains(configFile, "resource_renames_configuration"):
		// Use resource rename transformer
		transformer, err := transformations.NewResourceRenameTransformer(configPath)
		if err != nil {
			return fmt.Errorf("failed to create resource rename transformer: %v", err)
		}
		if dryRun {
			fmt.Printf("    ✗ Would transform config files in %s\n", configDir)
			return nil
		}
		return transformer.TransformDirectory(configDir, false)

	default:
		return fmt.Errorf("unknown configuration transformation type: %s", configFile)
	}
}

// applyStateTransformation applies a specific transformation to state files
func applyStateTransformation(stateFile, configPath, configFile string, dryRun bool) error {
	// Determine which transformer to use based on the config file name
	switch {
	case strings.Contains(configFile, "attribute_renames_state"):
		// Use state attribute transformer
		transformer, err := transformations.NewStateTransformer(configPath)
		if err != nil {
			return fmt.Errorf("failed to create state transformer: %v", err)
		}
		if dryRun {
			fmt.Printf("    ✗ Would transform state file %s\n", stateFile)
			return nil
		}
		return transformer.TransformFile(stateFile, stateFile)

	case strings.Contains(configFile, "resource_renames_state"):
		// Use state resource rename transformer
		transformer, err := transformations.NewStateResourceRenameTransformer(configPath)
		if err != nil {
			return fmt.Errorf("failed to create state resource rename transformer: %v", err)
		}
		if dryRun {
			fmt.Printf("    ✗ Would transform state file %s\n", stateFile)
			return nil
		}
		return transformer.TransformStateFile(stateFile, stateFile)

	default:
		return fmt.Errorf("unknown state transformation type: %s", configFile)
	}
}