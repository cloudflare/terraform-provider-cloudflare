package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/davecgh/go-spew/spew"
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

	// Generate moved blocks for list resources after both config and state are migrated
	if *configDir != "" && *configDir != "false" && *stateFile != "" && *stateFile != "false" && !*dryRun {
		if err := generateMovedBlocksForListDirectory(*configDir, *stateFile); err != nil {
			// Don't fail the migration if moved blocks generation fails, just warn
			log.Printf("Warning: Failed to generate moved blocks: %v\n", err)
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
	// Create new diagnostics collector for this file
	diags := ast.NewDiagnostics()
	// First, transform header blocks in load_balancer_pool resources at the string level
	// This is needed because grit leaves header blocks inside origins lists, which causes parsing issues
	contentStr := string(content)
	contentStr = transformLoadBalancerPoolHeaders(contentStr)
	// Also transform tiered_cache values at the string level
	contentStr = transformTieredCacheValues(contentStr)
	content = []byte(contentStr)

	file, hcl_diags := hclwrite.ParseConfig(content, filename, hcl.InitialPos)
	if hcl_diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL in %s: %s", filename, hcl_diags.Error())
	}
	diags.HclDiagnostics.Extend(hcl_diags)

	body := file.Body()
	blocks := body.Blocks()

	// Track blocks to remove and new blocks to add
	var blocksToRemove []*hclwrite.Block
	var newBlocks []*hclwrite.Block

	for _, block := range blocks {
		applyRenames(block)

		// TODO declare a map from resource types to transform functions instead of having all these if statement
		if isZoneSettingsOverrideResource(block) {
			blocksToRemove = append(blocksToRemove, block)
			newBlocks = append(newBlocks, transformZoneSettingsBlock(block)...)
		}

		if isLoadBalancerPoolResource(block) {
			transformLoadBalancerPoolBlock(block)
		}

		if isAccessPolicyResource(block) {
			// TOOD eventually pass diags through to all resource transformers, not just accessPolicyBlock
			transformAccessPolicyBlock(block, diags)
		}

		if isAccessApplicationResource(block) {
			transformAccessApplicationBlock(block, diags)
		}

		if isZeroTrustAccessIdentityProviderResource(block) {
			transformZeroTrustAccessIdentityProviderBlock(block, diags)
		}

		if isTieredCacheResource(block) {
			newTieredCacheBlocks := transformTieredCacheBlock(block)
			if newTieredCacheBlocks != nil {
				// This resource needs to be replaced with argo_tiered_caching
				blocksToRemove = append(blocksToRemove, block)
				newBlocks = append(newBlocks, newTieredCacheBlocks...)
			}
		}

		if isZeroTrustAccessMTLSHostnameSettingsResource(block) {
			transformZeroTrustAccessMTLSHostnameSettingsBlock(block, diags)
		}

		if isManagedTransformsResource(block) {
			transformManagedTransformsBlock(block)
		}

		if IsCloudflareListResource(block) {
			// Check if the block has item blocks (static or dynamic) that need to be split out
			hasItems := false
			hasDynamicItems := false
			for _, itemBlock := range block.Body().Blocks() {
				if itemBlock.Type() == "item" {
					hasItems = true
				}
				// Also check for dynamic "item" blocks
				if itemBlock.Type() == "dynamic" && len(itemBlock.Labels()) > 0 && itemBlock.Labels()[0] == "item" {
					hasItems = true
					hasDynamicItems = true
				}
			}

			if hasItems {
				// Transform the list resource: split items into separate resources
				blocksToRemove = append(blocksToRemove, block)
				transformedBlocks := TransformCloudflareListBlock(block)
				newBlocks = append(newBlocks, transformedBlocks...)

				// TODO: If there are dynamic items, generate moved blocks
				// This requires coordination with state migration
				// For now, the moved blocks can be generated separately using GenerateListMovedBlocks
				_ = hasDynamicItems
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

	// Report diagnostics
	// TODO make this controlled by a flag or something
	for _, d := range diags.HclDiagnostics {
		log.Println(strings.Join([]string{d.Summary, d.Detail}, "\n"))
	}
	for _, e := range diags.ComplicatedHCL {
		spew.Dump("Gave up dealing with expression:", e)
	}

	return formatted, nil
}

// generateMovedBlocksForListDirectory generates moved blocks for cloudflare_list resources after migration
func generateMovedBlocksForListDirectory(configDir string, stateFile string) error {
	// Read the state file
	stateData, err := os.ReadFile(stateFile)
	if err != nil {
		return fmt.Errorf("failed to read state file: %v", err)
	}

	// Walk through the directory to find config files with list_item resources
	err = filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-.tf files
		if info.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".tf") {
			return nil
		}

		// Read the config file
		configData, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Failed to read config file %s: %v", path, err)
			return nil // Continue with other files
		}

		// Check if this file has any cloudflare_list_item resources
		// which indicates it was migrated from a cloudflare_list with items
		if !strings.Contains(string(configData), "cloudflare_list_item") {
			return nil
		}

		fmt.Printf("  Generating moved blocks for %s\n", path)

		// Generate moved blocks for this config
		updatedConfig, err := GenerateListMovedBlocks(configData, stateData)
		if err != nil {
			log.Printf("Failed to generate moved blocks for %s: %v", path, err)
			return nil // Continue with other files
		}

		// Write the updated config back
		if err := os.WriteFile(path, updatedConfig, 0644); err != nil {
			log.Printf("Failed to write updated config to %s: %v", path, err)
			return nil // Continue with other files
		}

		fmt.Printf("    ✓ Added moved blocks to %s\n", path)
		return nil
	})

	return err
}
