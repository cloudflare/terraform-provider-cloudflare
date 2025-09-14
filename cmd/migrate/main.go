package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
)

func main() {
	// Define flags
	dryRun := flag.Bool("dryrun", false, "Show what changes would be made without actually modifying files")
	configDir := flag.String("config", "", "Directory containing Terraform files to migrate (defaults to current directory)")
	stateFile := flag.String("state", "", "Terraform state file to migrate (defaults to first .tfstate file in current directory)")
	useGrit := flag.Bool("grit", true, "Use grit for initial migrations (default: false)")
	useTransformer := flag.Bool("transformer", false, "Use Go-based YAML transformations (default: false)")
	transformerConfig := flag.String("transformer-dir", "", "Path to directory containing transformer YAML configs (defaults to embedded configs)")
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

	// Run Go-based YAML transformations if enabled
	if *useTransformer {
		fmt.Println("Running Go-based YAML transformations...")

		// Determine the transformer config directory
		transformerDir := *transformerConfig
		if transformerDir == "" {
			// Use default embedded configs from GitHub
			transformerDir = "https://github.com/cloudflare/terraform-provider-cloudflare/tree/grit-to-go-transformations/cmd/migrate/transformations/config"
			fmt.Println("Using embedded transformer configs from GitHub")
		} else {
			fmt.Printf("Using local transformer configs from: %s\n", transformerDir)
		}

		// Run the YAML-based transformations
		if err := runYAMLTransformations(*configDir, *stateFile, transformerDir, *dryRun); err != nil {
			log.Fatalf("Error running YAML transformations: %v", err)
		}
		fmt.Println("YAML transformations completed")
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
	fmt.Printf("    DEBUG: Processing file %s\n", filename)

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
		fmt.Printf("    DEBUG: No changes needed for %s\n", filename)
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

		// TODO declare a map from resource types to transform functions
		// instead of having all these if statement

		if isZoneSettingsOverrideResource(block) {
			blocksToRemove = append(blocksToRemove, block)
			newBlocks = append(newBlocks, transformZoneSettingsBlock(block)...)
		}

		if isRegionalHostnameResource(block) {
			transformRegionalHostnameBlock(block)
		}

		if isLoadBalancerPoolResource(block) {
			transformLoadBalancerPoolBlock(block, diags)
		}

		if isLoadBalancerResource(block) {
			transformLoadBalancerBlock(block, diags)
		}

		if isCloudflareRulesetResource(block) {
			transformCloudflareRulesetBlock(block, diags)
		}

		if isAccessPolicyResource(block) {
			// TOOD eventually pass diags through to all resource transformers,
			// not just accessPolicyBlock
			transformAccessPolicyBlock(block, diags)
		}

		if isAccessApplicationResource(block) {
			transformAccessApplicationBlock(block, diags)
		}

		if isZoneResource(block) {
			transformZoneBlock(block, diags)
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

		if isArgoResource(block) {
			// Transform cloudflare_argo to separate resources
			blocksToRemove = append(blocksToRemove, block)
			newBlocks = append(newBlocks, transformArgoBlock(block)...)
		}

		if isZeroTrustAccessMTLSHostnameSettingsResource(block) {
			transformZeroTrustAccessMTLSHostnameSettingsBlock(block, diags)
		}

		if isAccessMutualTLSHostnameSettingsResource(block) {
			transformZeroTrustAccessMTLSHostnameSettingsBlock(block, diags)
		}

		if isManagedTransformsResource(block) {
			transformManagedTransformsBlock(block)
		}

		if isAccessGroupResource(block) {
			transformAccessGroupBlock(block, diags)
		}

		if isDNSRecordResource(block) {
			// Process DNS record to fix CAA flags
			ProcessDNSRecordConfig(file)
		}

		if isSnippetResource(block) {
			transformSnippetBlock(block, diags)
		}

		if isSnippetRulesResource(block) {
			transformSnippetRulesBlock(block, diags)
		}

		if isSpectrumApplicationResource(block) {
			transformSpectrumApplicationBlock(block, diags)
		}

		if isWorkersRouteResource(block) {
			transformWorkersRouteBlock(block, diags)
		}

		if isWorkersScriptResource(block) {
			transformWorkersScriptBlock(block, diags)
		}

		if isWorkersCronTriggerResource(block) {
			transformWorkersCronTriggerBlock(block, diags)
		}

		if isWorkersDomainResource(block) {
			transformWorkersDomainBlock(block, diags)
		}

		// Note: workers_secret resources are handled by cross-resource migration below

		if isCloudflareListResource(block) {
			// Transform cloudflare_list item blocks to items attribute
			// Handles both static and dynamic blocks
			transformCloudflareListBlock(block)
		}
	}

	// Merge cloudflare_list_item resources into their parent lists
	// This must happen after processing all blocks to ensure we've seen all list and list_item resources
	listItemBlocksToRemove := mergeListItemResources(blocks)
	blocksToRemove = append(blocksToRemove, listItemBlocksToRemove...)

	// Remove old blocks
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}

	// Add new blocks
	for _, block := range newBlocks {
		body.AppendBlock(block)
	}

	// Perform cross-resource migration for workers_secret -> workers_script secret_text bindings
	// This must happen after all individual block transformations are complete
	migrateWorkersSecretsToBindings(file, diags)

	// Generate moved blocks for policy transitions if we have application-policy mappings
	if len(applicationPolicyMapping) > 0 {
		movedBlocks := generateMovedBlocks()
		for _, block := range movedBlocks {
			body.AppendBlock(block)
		}
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
