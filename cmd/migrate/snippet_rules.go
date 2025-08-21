package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isSnippetRulesResource checks if the given block is a cloudflare_snippet_rules resource
func isSnippetRulesResource(block *hclwrite.Block) bool {
	if block.Type() != "resource" {
		return false
	}
	labels := block.Labels()
	return len(labels) >= 1 && labels[0] == "cloudflare_snippet_rules"
}

// transformSnippetRulesBlock handles the migration from v4 to v5 for snippet_rules
// In v4: rules were defined as blocks (rules { ... })
// In v5: rules are defined as attributes (rules = [{ ... }])
// The Grit patterns handle this transformation, so no special Go transformation needed
func transformSnippetRulesBlock(block *hclwrite.Block) {
	// The v4 to v5 migration for snippet_rules is handled by Grit patterns
	// which convert blocks to attributes automatically.
	// This function is a placeholder for any future Go-specific transformations
	// that might be needed beyond what Grit can handle.
}