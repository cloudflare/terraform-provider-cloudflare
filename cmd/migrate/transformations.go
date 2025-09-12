package main

import (
	"fmt"
	"io"
	"net/http"
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

	// If transformerDir is a URL, download the configs from GitHub
	isRemote := strings.HasPrefix(transformerDir, "http")

	if isRemote {
		// Create a temporary directory to store downloaded configs
		var err error
		tempDir, err := os.MkdirTemp("", "migrate-configs-*")
		if err != nil {
			return fmt.Errorf("failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tempDir) // Clean up temp dir when done

		// Download each config file from GitHub
		baseURL := "https://raw.githubusercontent.com/cloudflare/terraform-provider-cloudflare/refs/heads/next/cmd/migrate/transformations/config"

		for _, t := range transformationConfigs {
			url := fmt.Sprintf("%s/%s", baseURL, t.configFile)
			destPath := filepath.Join(tempDir, t.configFile)

			fmt.Printf("  Downloading %s...\n", t.configFile)
			if err := downloadFile(url, destPath); err != nil {
				fmt.Printf("    ⚠ Warning: Failed to download %s: %v\n", t.configFile, err)
				// Continue with other files even if one fails
			}
		}

		// Use the temp directory as the transformer directory
		transformerDir = tempDir
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

// downloadFile downloads a file from a URL and saves it to the specified path
func downloadFile(url, destPath string) error {
	// Create the file
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
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
		fmt.Printf("    DEBUG: Calling TransformDirectory on %s (recursive=true)\n", configDir)
		return transformer.TransformDirectory(configDir, true)

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
		fmt.Printf("    DEBUG: Calling TransformDirectory on %s (recursive=true)\n", configDir)
		return transformer.TransformDirectory(configDir, true)

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
		fmt.Printf("    DEBUG: Calling TransformDirectory on %s (recursive=true)\n", configDir)
		return transformer.TransformDirectory(configDir, true)

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
