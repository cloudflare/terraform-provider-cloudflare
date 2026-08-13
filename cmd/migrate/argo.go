package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isArgoResource checks if a block is a cloudflare_argo resource
func isArgoResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_argo"
}

// transformArgoBlock transforms cloudflare_argo resource into separate 
// cloudflare_argo_smart_routing and/or cloudflare_argo_tiered_caching resources
func transformArgoBlock(block *hclwrite.Block) []*hclwrite.Block {
	body := block.Body()
	resourceName := block.Labels()[1]
	
	var newBlocks []*hclwrite.Block
	
	// Check for smart_routing attribute
	smartRoutingAttr := body.GetAttribute("smart_routing")
	if smartRoutingAttr != nil {
		// Create cloudflare_argo_smart_routing resource
		smartRoutingBlock := hclwrite.NewBlock("resource", []string{"cloudflare_argo_smart_routing", resourceName})
		smartRoutingBody := smartRoutingBlock.Body()
		
		// Copy zone_id
		if zoneIdAttr := body.GetAttribute("zone_id"); zoneIdAttr != nil {
			tokens := zoneIdAttr.Expr().BuildTokens(nil)
			smartRoutingBody.SetAttributeRaw("zone_id", tokens)
		}
		
		// Rename smart_routing to value
		tokens := smartRoutingAttr.Expr().BuildTokens(nil)
		smartRoutingBody.SetAttributeRaw("value", tokens)
		
		// Copy lifecycle and other nested blocks
		for _, nestedBlock := range body.Blocks() {
			copyBlock(smartRoutingBody, nestedBlock)
		}
		
		newBlocks = append(newBlocks, smartRoutingBlock)
		
		// Create moved block for smart_routing
		movedBlockSmart := createMovedBlock(
			"cloudflare_argo." + resourceName,
			"cloudflare_argo_smart_routing." + resourceName,
		)
		newBlocks = append(newBlocks, movedBlockSmart)
	}
	
	// Check for tiered_caching attribute
	tieredCachingAttr := body.GetAttribute("tiered_caching")
	if tieredCachingAttr != nil {
		// Create cloudflare_argo_tiered_caching resource with a different name to avoid conflicts
		// when both smart_routing and tiered_caching exist
		tieredResourceName := resourceName
		if smartRoutingAttr != nil {
			// Only append suffix if we have both attributes to avoid name collision
			tieredResourceName = resourceName + "_tiered"
		}
		
		tieredCachingBlock := hclwrite.NewBlock("resource", []string{"cloudflare_argo_tiered_caching", tieredResourceName})
		tieredCachingBody := tieredCachingBlock.Body()
		
		// Copy zone_id
		if zoneIdAttr := body.GetAttribute("zone_id"); zoneIdAttr != nil {
			tokens := zoneIdAttr.Expr().BuildTokens(nil)
			tieredCachingBody.SetAttributeRaw("zone_id", tokens)
		}
		
		// Rename tiered_caching to value
		tokens := tieredCachingAttr.Expr().BuildTokens(nil)
		tieredCachingBody.SetAttributeRaw("value", tokens)
		
		// Copy lifecycle and other nested blocks
		for _, nestedBlock := range body.Blocks() {
			copyBlock(tieredCachingBody, nestedBlock)
		}
		
		newBlocks = append(newBlocks, tieredCachingBlock)
		
		// Create moved block for tiered_caching
		movedBlockTiered := createMovedBlock(
			"cloudflare_argo." + resourceName,
			"cloudflare_argo_tiered_caching." + tieredResourceName,
		)
		newBlocks = append(newBlocks, movedBlockTiered)
	}
	
	// If neither attribute exists, just create a smart_routing resource with value = "off"
	if smartRoutingAttr == nil && tieredCachingAttr == nil {
		smartRoutingBlock := hclwrite.NewBlock("resource", []string{"cloudflare_argo_smart_routing", resourceName})
		smartRoutingBody := smartRoutingBlock.Body()
		
		// Copy zone_id
		if zoneIdAttr := body.GetAttribute("zone_id"); zoneIdAttr != nil {
			tokens := zoneIdAttr.Expr().BuildTokens(nil)
			smartRoutingBody.SetAttributeRaw("zone_id", tokens)
		}
		
		// Default to "off"
		smartRoutingBody.SetAttributeValue("value", cty.StringVal("off"))
		
		// Copy lifecycle and other nested blocks
		for _, nestedBlock := range body.Blocks() {
			copyBlock(smartRoutingBody, nestedBlock)
		}
		
		newBlocks = append(newBlocks, smartRoutingBlock)
		
		// Create moved block
		movedBlock := createMovedBlock(
			"cloudflare_argo." + resourceName,
			"cloudflare_argo_smart_routing." + resourceName,
		)
		newBlocks = append(newBlocks, movedBlock)
	}
	
	return newBlocks
}