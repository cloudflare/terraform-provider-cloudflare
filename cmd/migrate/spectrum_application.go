package main

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isSpectrumApplicationResource checks if a block is a cloudflare_spectrum_application resource
func isSpectrumApplicationResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_spectrum_application"
}

// transformSpectrumApplicationBlock transforms various attributes in spectrum_application resources
//
// Example transformations:
// 1. Remove optional id attribute from configuration (V4 → V5)
// 2. Block to object syntax changes are handled by transformations
// 3. Convert origin_port_range block to origin_port string format (V4 → V5)
func transformSpectrumApplicationBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Remove the id attribute if present (V4 allowed optional, V5 is computed-only)
	body := block.Body()
	if body.GetAttribute("id") != nil {
		body.RemoveAttribute("id")
	}
	
	// Handle origin_port_range to origin_port conversion
	// V4: origin_port_range { start = 80, end = 85 }
	// V5: origin_port = "80-85"
	transformOriginPortRange(body, diags)
}

// transformOriginPortRange converts origin_port_range block to origin_port string
func transformOriginPortRange(body *hclwrite.Body, diags ast.Diagnostics) {
	// Look for origin_port_range blocks
	for _, block := range body.Blocks() {
		if block.Type() == "origin_port_range" {
			blockBody := block.Body()
			
			// Extract start and end attributes
			startAttr := blockBody.GetAttribute("start")
			endAttr := blockBody.GetAttribute("end")
			
			if startAttr != nil && endAttr != nil {
				// Parse the values (they should be numeric tokens)
				startTokens := startAttr.Expr().BuildTokens(nil)
				endTokens := endAttr.Expr().BuildTokens(nil)
				
				if len(startTokens) >= 1 && len(endTokens) >= 1 {
					startValue := string(startTokens[0].Bytes)
					endValue := string(endTokens[0].Bytes)
					
					// Create the range string: "start-end"
					rangeValue := fmt.Sprintf(`"%s-%s"`, startValue, endValue)
					
					// Set origin_port attribute with the range string
					body.SetAttributeRaw("origin_port", []*hclwrite.Token{
						{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(rangeValue)},
					})
				}
			}
			
			// Remove the origin_port_range block
			body.RemoveBlock(block)
			break // Only handle the first one (should be MaxItems: 1 anyway)
		}
	}
}