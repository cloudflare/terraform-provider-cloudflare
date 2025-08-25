package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ListItemMigrationInfo holds information about a list item migration
type ListItemMigrationInfo struct {
	ListName    string
	OldItemName string // e.g., "salesforce_ips_item_0"
	NewItemName string // e.g., "salesforce_ips_items"
	ItemValue   string // The actual value (IP, ASN, etc.) to use as the key
	IsForEach   bool   // Whether the new resource uses for_each
}

// GenerateListMovedBlocks generates moved blocks for cloudflare_list_item migrations
// It analyzes the upgraded config and state to determine the mapping between old static items
// and new for_each items, then appends the moved blocks to the config.
//
// This function should be called after both config and state migrations are complete.
// It's particularly useful when dynamic blocks have been converted to for_each resources.
//
// Returns the updated config with moved blocks appended.
func GenerateListMovedBlocks(upgradedConfig []byte, upgradedStateJSON []byte) ([]byte, error) {
	// Parse the upgraded config to understand the resource structure
	configFile, diags := hclwrite.ParseConfig(upgradedConfig, "config.tf", hcl.InitialPos)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse config: %v", diags)
	}

	// Parse the upgraded state to get actual item values
	var state CFListStateV5
	if err := json.Unmarshal(upgradedStateJSON, &state); err != nil {
		return nil, fmt.Errorf("failed to parse state: %w", err)
	}

	// Collect migration info
	migrations := collectListItemMigrations(configFile, &state)

	// Generate moved blocks
	var movedBlocks []*hclwrite.Block
	for _, migration := range migrations {
		if migration.IsForEach && migration.ItemValue != "" {
			movedBlock := createListMovedBlock(migration)
			movedBlocks = append(movedBlocks, movedBlock)
		}
	}

	// Append moved blocks to the config
	for _, block := range movedBlocks {
		configFile.Body().AppendBlock(block)
	}

	// Format and return the updated config
	return hclwrite.Format(configFile.Bytes()), nil
}

// collectListItemMigrations analyzes config and state to find migrations needed
func collectListItemMigrations(configFile *hclwrite.File, state *CFListStateV5) []ListItemMigrationInfo {
	var migrations []ListItemMigrationInfo

	// Track which resources exist in the config
	configResources := make(map[string]bool)    // resource_type.resource_name -> exists
	forEachResources := make(map[string]string) // list_name -> for_each_resource_name

	// Analyze the config to understand what resources exist
	for _, block := range configFile.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 {
			resourceType := block.Labels()[0]
			resourceName := block.Labels()[1]

			// Track all cloudflare_list_item resources
			if resourceType == "cloudflare_list_item" {
				configResources[fmt.Sprintf("%s.%s", resourceType, resourceName)] = true

				// Check if it has for_each
				if block.Body().GetAttribute("for_each") != nil {
					// This is a for_each resource
					// Extract the list name from the resource name
					// Convention: list_name_items -> list_name
					listName := extractListNameFromItemResource(resourceName)
					if listName != "" {
						forEachResources[listName] = resourceName
					}
				}
			}
		}
	}

	// Now look in the state for list_item resources
	for _, resource := range state.Resources {
		if resource.Type == "cloudflare_list_item" {
			stateResourceID := fmt.Sprintf("%s.%s", resource.Type, resource.Name)

			// Check if this resource exists in the config
			if !configResources[stateResourceID] {
				// This state resource doesn't exist in config - might need a moved block

				// Check if it's a static item that should map to a for_each resource
				listName := extractListNameFromStaticItem(resource.Name)
				if listName != "" {
					// Check if there's a corresponding for_each resource in the config
					if forEachResourceName, exists := forEachResources[listName]; exists {
						// Get the item value from state to use as the for_each key
						itemValue := getItemValueFromState(state, resource.Name)
						if itemValue != "" {
							migrations = append(migrations, ListItemMigrationInfo{
								ListName:    listName,
								OldItemName: resource.Name,
								NewItemName: forEachResourceName,
								ItemValue:   itemValue,
								IsForEach:   true,
							})
						}
					}
				}
			}
		}
	}

	return migrations
}

// createListMovedBlock creates a moved block for a list item migration
func createListMovedBlock(migration ListItemMigrationInfo) *hclwrite.Block {
	block := hclwrite.NewBlock("moved", nil)
	body := block.Body()

	// Set 'from' attribute - the old static resource
	fromTraversal := hcl.Traversal{
		hcl.TraverseRoot{Name: "cloudflare_list_item"},
		hcl.TraverseAttr{Name: migration.OldItemName},
	}
	body.SetAttributeTraversal("from", fromTraversal)

	// Set 'to' attribute - the new for_each resource with key
	// We need to create a traversal with an index like: cloudflare_list_item.salesforce_ips_items["128.0.0.1"]
	toTokens := buildIndexedResourceReference("cloudflare_list_item", migration.NewItemName, migration.ItemValue)
	body.SetAttributeRaw("to", toTokens)

	return block
}

// buildIndexedResourceReference builds tokens for a reference like resource.name["key"]
func buildIndexedResourceReference(resourceType, resourceName, indexKey string) hclwrite.Tokens {
	var tokens hclwrite.Tokens

	// Add resource type
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte(resourceType),
	})

	// Add dot
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenDot,
		Bytes: []byte("."),
	})

	// Add resource name
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte(resourceName),
	})

	// Add index - ["key"]
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenOBrack,
		Bytes: []byte("["),
	})
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenOQuote,
		Bytes: []byte("\""),
	})
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenQuotedLit,
		Bytes: []byte(indexKey),
	})
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenCQuote,
		Bytes: []byte("\""),
	})
	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrack,
		Bytes: []byte("]"),
	})

	return tokens
}

// extractListNameFromItemResource extracts the list name from a for_each item resource name
// e.g., "salesforce_ips_items" -> "salesforce_ips"
func extractListNameFromItemResource(resourceName string) string {
	if strings.HasSuffix(resourceName, "_items") {
		return strings.TrimSuffix(resourceName, "_items")
	}
	return ""
}

// extractListNameFromStaticItem extracts the list name from a static item resource name
// e.g., "salesforce_ips_item_0" -> "salesforce_ips"
func extractListNameFromStaticItem(resourceName string) string {
	// Look for pattern: list_name_item_N
	parts := strings.Split(resourceName, "_")
	if len(parts) >= 3 {
		// Check if it ends with item_N pattern
		if parts[len(parts)-2] == "item" {
			// Check if last part is a number
			if isNumeric(parts[len(parts)-1]) {
				// Reconstruct the list name (everything except _item_N)
				return strings.Join(parts[:len(parts)-2], "_")
			}
		}
	}
	return ""
}

// isNumeric checks if a string represents a number
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}

// findStaticItemsForList finds all static item resources for a given list
func findStaticItemsForList(configFile *hclwrite.File, listName string) []string {
	var staticItems []string

	for _, block := range configFile.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 {
			resourceType := block.Labels()[0]
			resourceName := block.Labels()[1]

			if resourceType == "cloudflare_list_item" {
				// Check if this is a static item for our list
				if extractListNameFromStaticItem(resourceName) == listName {
					// Make sure it doesn't have for_each (is truly static)
					if block.Body().GetAttribute("for_each") == nil {
						staticItems = append(staticItems, resourceName)
					}
				}
			}
		}
	}

	return staticItems
}

// getItemValueFromState extracts the actual value of a list item from state
func getItemValueFromState(state *CFListStateV5, itemResourceName string) string {
	// Find the resource in state
	resources := state.FindResource("cloudflare_list_item", itemResourceName)
	if len(resources) == 0 || len(resources[0].Instances) == 0 {
		return ""
	}

	// Get the attributes
	instance := resources[0].Instances[0]
	attrs, err := instance.GetListItemAttributes()
	if err != nil {
		return ""
	}

	// Extract the value based on what's present
	if attrs.IP != nil && *attrs.IP != "" {
		return *attrs.IP
	}
	if attrs.ASN != nil {
		return fmt.Sprintf("%d", *attrs.ASN)
	}
	if attrs.Hostname != nil && attrs.Hostname.URLHostname != "" {
		return attrs.Hostname.URLHostname
	}
	if attrs.Redirect != nil {
		// Use source URL as the key for redirects
		return attrs.Redirect.SourceURL
	}

	return ""
}

// GenerateMovedBlocksFromMigration generates moved blocks by comparing before and after states
// This is an alternative approach that works directly with state changes.
// Note: This returns just the blocks, not an updated config. Use GenerateListMovedBlocks
// if you need to append the blocks to an existing config.
func GenerateMovedBlocksFromMigration(beforeStateJSON, afterStateJSON []byte) ([]*hclwrite.Block, error) {
	// Parse before state (v4)
	var beforeState CFListStateV4
	if err := json.Unmarshal(beforeStateJSON, &beforeState); err != nil {
		return nil, fmt.Errorf("failed to parse before state: %w", err)
	}

	// Parse after state (v5)
	var afterState CFListStateV5
	if err := json.Unmarshal(afterStateJSON, &afterState); err != nil {
		return nil, fmt.Errorf("failed to parse after state: %w", err)
	}

	var movedBlocks []*hclwrite.Block

	// For each cloudflare_list in the before state that had items
	for _, listResource := range beforeState.GetCloudflareListResources() {
		if len(listResource.Instances) == 0 {
			continue
		}

		listName := listResource.Name
		listAttrs := listResource.Instances[0].Attributes

		// Check if this list had items
		if len(listAttrs.Item) > 0 {
			// Look for corresponding for_each resource in after state
			forEachResourceName := fmt.Sprintf("%s_items", listName)
			forEachResources := afterState.FindResource("cloudflare_list_item", forEachResourceName)

			if len(forEachResources) > 0 {
				// This list was migrated to for_each
				// Create moved blocks for each static item
				for i, item := range listAttrs.Item {
					staticResourceName := fmt.Sprintf("%s_item_%d", listName, i)

					// Get the item value to use as the key
					itemValue := getItemValueFromV4Item(item, listAttrs.Kind)
					if itemValue != "" {
						migration := ListItemMigrationInfo{
							ListName:    listName,
							OldItemName: staticResourceName,
							NewItemName: forEachResourceName,
							ItemValue:   itemValue,
							IsForEach:   true,
						}
						movedBlocks = append(movedBlocks, createListMovedBlock(migration))
					}
				}
			}
		}
	}

	return movedBlocks, nil
}

// getItemValueFromV4Item extracts the value from a v4 list item
func getItemValueFromV4Item(item ListItemV4, kind string) string {
	if len(item.Value) == 0 {
		return ""
	}

	v := item.Value[0]
	switch kind {
	case "ip":
		if v.IP != nil {
			return fmt.Sprintf("%v", v.IP)
		}
	case "asn":
		if v.ASN != nil {
			switch asn := v.ASN.(type) {
			case float64:
				return fmt.Sprintf("%d", int64(asn))
			case int:
				return fmt.Sprintf("%d", asn)
			case int64:
				return fmt.Sprintf("%d", asn)
			case string:
				return asn
			}
		}
	case "hostname":
		if len(v.Hostname) > 0 {
			return v.Hostname[0].URLHostname
		}
	case "redirect":
		if len(v.Redirect) > 0 {
			return v.Redirect[0].SourceURL
		}
	}

	return ""
}
