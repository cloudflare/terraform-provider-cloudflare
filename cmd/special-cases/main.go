package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define flags
	dryRun := flag.Bool("dryrun", false, "Show what changes would be made without actually modifying files")
	configDir := flag.String("config", "", "Directory containing Terraform files to migrate (defaults to current directory)")
	caseType := flag.String("case", "", "Special case to handle (required): zone-settings-modules")
	flag.Parse()

	if *caseType == "" {
		fmt.Fprintf(os.Stderr, "Error: --case flag is required\n")
		fmt.Fprintf(os.Stderr, "Available cases: zone-settings-modules\n")
		os.Exit(1)
	}

	// Default config directory to current working directory if not specified
	if *configDir == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current working directory: %v", err)
		}
		*configDir = currentDir
	}

	if *dryRun {
		fmt.Println("DRY RUN MODE - No files will be modified")
		fmt.Println(strings.Repeat("=", 50))
	}

	switch *caseType {
	case "zone-settings-modules":
		if err := handleZoneSettingsModules(*configDir, *dryRun); err != nil {
			log.Fatalf("Error handling zone settings modules: %v", err)
		}
	default:
		log.Fatalf("Unknown case type: %s", *caseType)
	}

	if !*dryRun {
		fmt.Println("Special case migration completed successfully!")
	}
}

func handleZoneSettingsModules(configDir string, dryRun bool) error {
	fmt.Printf("Handling zone settings modules in: %s\n", configDir)
	
	// First, discover all modules that use cloudflare_zone_settings_override
	modules, err := findZoneSettingsModules(configDir)
	if err != nil {
		return fmt.Errorf("failed to find zone settings modules: %v", err)
	}
	
	if len(modules) == 0 {
		fmt.Println("No zone settings modules found")
		return nil
	}
	
	fmt.Printf("Found %d zone settings modules:\n", len(modules))
	for _, module := range modules {
		fmt.Printf("  - %s (source: %s)\n", module.Name, module.Source)
	}
	
	// Process each root module file that uses these zone settings modules
	return filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories and non-tf files
		if info.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".tf") {
			return nil
		}
		
		// Check if this file uses any zone settings modules
		usesZoneSettingsModules, err := fileUsesZoneSettingsModules(path, modules)
		if err != nil {
			log.Printf("Error checking file %s: %v", path, err)
			return nil
		}
		
		if !usesZoneSettingsModules {
			return nil
		}
		
		// Transform this file
		if err := transformFileWithModuleExpansion(path, modules, dryRun); err != nil {
			log.Printf("Error transforming file %s: %v", path, err)
		}
		
		return nil
	})
}