package main

import (
	"regexp"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

// transformLoadBalancerPoolBlock transforms a load_balancer_pool resource block
func transformLoadBalancerPoolBlock(block *hclwrite.Block) {
	// This function is called after the full file transformation
	// No additional processing needed here since we handle it at the file level
}

// transformHeadersInOrigins transforms header blocks in the origins string
func transformHeadersInOrigins(origins string) string {
	// Pattern to match header blocks: header { header = "Host" values = [...] }
	// This regex looks for header blocks and captures the values array
	headerPattern := regexp.MustCompile(`header\s*\{\s*header\s*=\s*"Host"\s*values\s*=\s*(\[[^\]]+\])\s*\}`)

	// Replace with the v5 format: header = { host = [...] }
	result := headerPattern.ReplaceAllString(origins, `header = { host = $1 }`)

	return result
}

// isLoadBalancerPoolResource checks if a block is a load_balancer_pool resource
func isLoadBalancerPoolResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_load_balancer_pool"
}

// transformLoadBalancerPoolHeaders transforms header blocks in load_balancer_pool resources
// This is done at the string level before HCL parsing to avoid syntax errors
func transformLoadBalancerPoolHeaders(content string) string {
	// Pattern to match header blocks inside origins
	// This handles the case where grit has converted origins to a list but left header as a block
	// The pattern needs to be flexible with whitespace and indentation
	headerBlockPattern := regexp.MustCompile(`(?m)([ \t]*)header\s*\{\s*\n[ \t]*header\s*=\s*"Host"\s*\n[ \t]*values\s*=\s*(\[[^\]]+\])\s*\n[ \t]*\}`)

	// Replace header blocks with header attributes
	result := headerBlockPattern.ReplaceAllString(content, `${1}header = { host = ${2} }`)

	return result
}

