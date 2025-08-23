package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isRegionalHostnameResource checks if the given block is a cloudflare_regional_hostname resource
func isRegionalHostnameResource(block *hclwrite.Block) bool {
	if block.Type() != "resource" {
		return false
	}
	labels := block.Labels()
	return len(labels) >= 1 && labels[0] == "cloudflare_regional_hostname"
}

// transformRegionalHostnameBlock removes timeouts blocks from regional hostname resources
// since v5 provider doesn't support them
func transformRegionalHostnameBlock(block *hclwrite.Block) {
	body := block.Body()
	
	// Find and remove timeouts blocks
	var blocksToRemove []*hclwrite.Block
	for _, nestedBlock := range body.Blocks() {
		if nestedBlock.Type() == "timeouts" {
			blocksToRemove = append(blocksToRemove, nestedBlock)
		}
	}
	
	// Remove the timeouts blocks
	for _, blockToRemove := range blocksToRemove {
		body.RemoveBlock(blockToRemove)
	}
}