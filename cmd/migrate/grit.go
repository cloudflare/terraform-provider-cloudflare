package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// checkGritInstalled checks if grit is installed and available
func checkGritInstalled() error {
	cmd := exec.Command("grit", "--help")
	cmd.Stdout = nil // Suppress output during check
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("grit is not installed. Please install it with: npm install -g @getgrit/cli")
	}
	return nil
}

// runGritMigrations runs the grit migrations for Cloudflare Terraform provider v5
func runGritMigrations(configDir string, stateFile string, patternsDir string, dryRun bool) error {
	// Check for patterns directory from environment variable first (for local development and acceptance testing)
	useLocalPatterns := !(patternsDir == "")

	// Define the grit patterns to apply
	patterns := []struct {
		pattern string
		target  string
	}{}

	if useLocalPatterns {
		// Use the patterns from local directory
		patterns = []struct {
			pattern string
			target  string
		}{
			{"cloudflare_terraform_v5", "config"},
			{"cloudflare_terraform_v5_attribute_renames_state", "state"},
			{"cloudflare_terraform_v5_resource_renames_configuration", "config"},
			{"cloudflare_terraform_v5_resource_renames_state", "state"},
		}
		fmt.Printf("Using local grit patterns from: %s\n", patternsDir)
	} else {
		// Use GitHub patterns
		patterns = []struct {
			pattern string
			target  string
		}{
			{"github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_v5", "config"},
			{"github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_v5_attribute_renames_state", "state"},
			{"github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_v5_resource_renames_configuration", "config"},
			{"github.com/cloudflare/terraform-provider-cloudflare#cloudflare_terraform_v5_resource_renames_state", "state"},
		}
		fmt.Println("Warning: Using GitHub grit patterns (local patterns not found)")
	}

	for _, p := range patterns {
		// Determine the target path based on pattern type
		targetPath := configDir
		workingDir := configDir
		if p.target == "state" && stateFile != "" {
			targetPath = stateFile
			workingDir = filepath.Dir(stateFile)
		}

		// Skip config patterns if no config directory is specified
		if p.target == "config" && configDir == "" {
			fmt.Printf("Skipping %s (config transformations disabled)\n", p.pattern)
			continue
		}

		// Skip state patterns if no state file is specified
		if p.target == "state" && stateFile == "" {
			fmt.Printf("Skipping %s (state transformations disabled)\n", p.pattern)
			continue
		}

		args := []string{"apply"}
		if dryRun {
			args = append(args, "--dry-run")
		}
		args = append(args, "--verbose")
		args = append(args, "--force")
		if useLocalPatterns {
			args = append(args, "--grit-dir", patternsDir)
		}
		args = append(args, p.pattern)
		args = append(args, targetPath)

		fmt.Printf("grit %s\n", strings.Join(args, " "))
		cmd := exec.Command("grit", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// Grit warns about untracked changes in the working tree, so we need to make sure it's in
		// the correct working directory (i.e. the directory that contains the target config/state)
		// When using local patterns, run from the patterns directory
		cmd.Dir = workingDir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to run grit pattern %s: %w", p.pattern, err)
		}
	}

	return nil
}
