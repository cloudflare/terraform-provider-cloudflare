package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func main() {
	// Define flags
	dryRun := flag.Bool("dryrun", false, "Show what changes would be made without actually modifying files")
	configDir := flag.String("config", "", "Directory containing Terraform files to migrate (defaults to current directory)")
	stateFile := flag.String("state", "", "Terraform state file to migrate (defaults to first .tfstate file in current directory)")
	useGrit := flag.Bool("grit", true, "Use grit for initial migrations (default: true)")
	patternsDir := flag.String("patterns-dir", "", "Local directory to get patterns from (otherwise they're pulled from github)")
	flag.Parse()

	// Default config directory to current working directory if not specified
	if *configDir == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current working directory: %v", err)
		}
		*configDir = currentDir
	}

	// Default state file to first .tfstate file in current directory if not specified
	if *stateFile == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current working directory: %v", err)
		}

		entries, err := os.ReadDir(currentDir)
		if err != nil {
			log.Fatalf("Failed to read current directory: %v", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".tfstate") {
				*stateFile = filepath.Join(currentDir, entry.Name())
				break
			}
		}
	}

	if *dryRun {
		fmt.Println("DRY RUN MODE - No files will be modified")
		fmt.Println(strings.Repeat("=", 50))
	}

	// Run grit migrations if enabled
	if *useGrit {
		fmt.Println("Running grit migrations...")
		if err := checkGritInstalled(); err != nil {
			log.Fatalf("Error: %v", err)
		}

		// Pass empty strings to grit if explicitly disabled
		gritConfigDir := *configDir
		gritStateFile := *stateFile
		if gritConfigDir == "false" {
			gritConfigDir = ""
		}
		if gritStateFile == "false" {
			gritStateFile = ""
		}

		if err := runGritMigrations(gritConfigDir, gritStateFile, *patternsDir, *dryRun); err != nil {
			log.Fatalf("Error running grit migrations: %v", err)
		}
		fmt.Println("Grit migrations completed")
		fmt.Println(strings.Repeat("-", 50))
	}

	// Process config directory if specified and not explicitly disabled
	if *configDir != "" && *configDir != "false" {
		fmt.Println("Running Go transformations on configuration files...")
		if err := processConfigDirectory(*configDir, *dryRun); err != nil {
			log.Fatalf("Error processing config directory: %v", err)
		}
	}

	// Process state file if specified and not explicitly disabled
	if *stateFile != "" && *stateFile != "false" {
		if err := processStateFile(*stateFile, *dryRun); err != nil {
			log.Fatalf("Error processing state file: %v", err)
		}
	}

	if !*dryRun {
		fmt.Println("Migration completed successfully!")
	}
}

func processConfigDirectory(directory string, dryRun bool) error {
	// Check if the directory exists
	info, err := os.Stat(directory)
	if err != nil {
		return fmt.Errorf("failed to access directory %s: %v", directory, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", directory)
	}

	fmt.Printf("Processing config directory: %s\n", directory)

	// Walk through the directory recursively
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process .tf files
		if !strings.HasSuffix(strings.ToLower(path), ".tf") {
			return nil
		}

		// Process the file
		if err := processFile(path, dryRun); err != nil {
			log.Printf("Error processing file %s: %v", path, err)
			// return err
		}

		return nil
	})
}

func processStateFile(stateFile string, dryRun bool) error {
	// Check if the state file exists
	info, err := os.Stat(stateFile)
	if err != nil {
		return fmt.Errorf("failed to access state file %s: %v", stateFile, err)
	}
	if info.IsDir() {
		return fmt.Errorf("%s is a directory, expected a file", stateFile)
	}

	// Check if it's a .tfstate file
	if !strings.HasSuffix(strings.ToLower(stateFile), ".tfstate") {
		return fmt.Errorf("%s is not a .tfstate file", stateFile)
	}

	fmt.Printf("Processing state file: %s\n", stateFile)

	if dryRun {
		fmt.Printf("  ✗ Would update state file %s\n", stateFile)
		fmt.Printf("    State file size: %d bytes\n", info.Size())
	} else {
		// Actually transform the state file
		if err := transformStateFile(stateFile); err != nil {
			return fmt.Errorf("failed to transform state file: %v", err)
		}
		fmt.Printf("  ✓ Updated state file %s\n", stateFile)
	}

	return nil
}

func processFile(filename string, dryRun bool) error {
	// Read the file
	originalBytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filename, err)
	}

	// Transform the file
	transformedBytes, err := transformFile(originalBytes, filename)
	if err != nil {
		return err
	}

	if string(originalBytes) == string(transformedBytes) {
		return nil
	}

	if dryRun {
		fmt.Printf("  ✗ Would update %s\n", filename)
	} else {
		// Write the result back to the file
		err = os.WriteFile(filename, transformedBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %v", filename, err)
		}
		fmt.Printf("  ✓ Updated %s\n", filename)
	}

	return nil
}

func transformFile(content []byte, filename string) ([]byte, error) {
	// First, transform header blocks in load_balancer_pool resources at the string level
	// This is needed because grit leaves header blocks inside origins lists, which causes parsing issues
	contentStr := string(content)
	contentStr = transformLoadBalancerPoolHeaders(contentStr)
	// Also transform tiered_cache values at the string level
	contentStr = transformTieredCacheValues(contentStr)
	content = []byte(contentStr)

	file, diags := hclwrite.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL in %s: %s", filename, diags.Error())
	}

	body := file.Body()
	blocks := body.Blocks()

	// Track blocks to remove and new blocks to add
	var blocksToRemove []*hclwrite.Block
	var newBlocks []*hclwrite.Block

	for _, block := range blocks {
		applyRenames(block)

		if isZoneSettingsOverrideResource(block) {
			blocksToRemove = append(blocksToRemove, block)
			newBlocks = append(newBlocks, transformZoneSettingsBlock(block)...)
		}

		if isLoadBalancerPoolResource(block) {
			transformLoadBalancerPoolBlock(block)
		}

		if isAccessPolicyResource(block) {
			transformAccessPolicyBlock(block)
		}

		if isAccessApplicationResource(block) {
			transformAccessApplicationBlock(block)
		}

		if isTieredCacheResource(block) {
			newTieredCacheBlocks := transformTieredCacheBlock(block)
			if newTieredCacheBlocks != nil {
				// This resource needs to be replaced with argo_tiered_caching
				blocksToRemove = append(blocksToRemove, block)
				newBlocks = append(newBlocks, newTieredCacheBlocks...)
			}
		}
	}

	// Remove old blocks
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}

	// Add new blocks
	for _, block := range newBlocks {
		body.AppendBlock(block)
	}

	formatted := hclwrite.Format(file.Bytes())
	return formatted, nil
}
