package main

import (
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isTieredCacheResource checks if a block is a cloudflare_tiered_cache resource
func isTieredCacheResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_tiered_cache"
}

// transformTieredCacheBlock returns new blocks if the resource should be migrated to argo_tiered_caching
// Returns nil if the block should remain as-is (handled by string transformation)
func transformTieredCacheBlock(block *hclwrite.Block) []*hclwrite.Block {
	body := block.Body()

	// Check if we have a value attribute with "generic"
	// (after string transformation has renamed cache_type to value)
	valueAttr := body.GetAttribute("value")
	if valueAttr == nil {
		return nil
	}

	// Check if the value is "generic" by examining the tokens
	tokens := valueAttr.Expr().BuildTokens(nil)
	isGeneric := false
	for _, token := range tokens {
		tokenStr := strings.Trim(string(token.Bytes), `"`)
		if tokenStr == "generic" {
			isGeneric = true
			break
		}
	}

	if !isGeneric {
		return nil // Not a generic type, no transformation needed
	}

	// Create new argo_tiered_caching resource
	resourceName := block.Labels()[1]
	newResource := hclwrite.NewBlock("resource", []string{"cloudflare_argo_tiered_caching", resourceName})
	newBody := newResource.Body()

	// Copy all attributes in deterministic order
	for _, attrInfo := range AttributesOrdered(body) {
		if attrInfo.Name == "value" {
			// Set value to "on" for argo_tiered_caching
			newBody.SetAttributeValue("value", cty.StringVal("on"))
		} else {
			// Copy other attributes as-is
			tokens := attrInfo.Attribute.Expr().BuildTokens(nil)
			newBody.SetAttributeRaw(attrInfo.Name, tokens)
		}
	}

	// Copy all nested blocks (like lifecycle)
	for _, nestedBlock := range body.Blocks() {
		newNestedBlock := hclwrite.NewBlock(nestedBlock.Type(), nestedBlock.Labels())
		// Copy the content of the nested block in deterministic order
		for _, attrInfo := range AttributesOrdered(nestedBlock.Body()) {
			tokens := attrInfo.Attribute.Expr().BuildTokens(nil)
			newNestedBlock.Body().SetAttributeRaw(attrInfo.Name, tokens)
		}
		// Recursively copy any deeper nested blocks
		for _, deeperBlock := range nestedBlock.Body().Blocks() {
			copyBlock(newNestedBlock.Body(), deeperBlock)
		}
		newBody.AppendBlock(newNestedBlock)
	}

	// Create moved block
	movedBlock := createMovedBlock(
		"cloudflare_tiered_cache."+resourceName,
		"cloudflare_argo_tiered_caching."+resourceName,
	)

	return []*hclwrite.Block{newResource, movedBlock}
}

// Helper function to recursively copy blocks
func copyBlock(targetBody *hclwrite.Body, sourceBlock *hclwrite.Block) {
	newBlock := hclwrite.NewBlock(sourceBlock.Type(), sourceBlock.Labels())

	// Copy attributes in deterministic order
	for _, attrInfo := range AttributesOrdered(sourceBlock.Body()) {
		tokens := attrInfo.Attribute.Expr().BuildTokens(nil)
		newBlock.Body().SetAttributeRaw(attrInfo.Name, tokens)
	}

	// Recursively copy nested blocks
	for _, deeperBlock := range sourceBlock.Body().Blocks() {
		copyBlock(newBlock.Body(), deeperBlock)
	}
	targetBody.AppendBlock(newBlock)
}

// transformTieredCacheValues transforms tiered_cache attribute values at the string level
// This handles both the attribute rename (cache_type -> value) and value transformation
// NOTE: We don't transform "generic" values here - they're handled by HCL transformation
func transformTieredCacheValues(content string) string {
	// First, handle resources that already have "value" attribute (from transformations)
	// Pattern to match value = "smart" in tiered_cache resources
	valueSmartPattern := regexp.MustCompile(`(resource\s+"cloudflare_tiered_cache"[^{]+\{[^}]*\n\s*value\s*=\s*)"smart"`)
	content = valueSmartPattern.ReplaceAllString(content, `${1}"on"`)

	// Don't transform "generic" - it will be handled by HCL transformation to create argo_tiered_caching

	// Also handle cache_type in case transformations haven't run
	cacheTypeSmartPattern := regexp.MustCompile(`(resource\s+"cloudflare_tiered_cache"[^{]+\{[^}]*\n\s*)cache_type(\s*=\s*)"smart"`)
	content = cacheTypeSmartPattern.ReplaceAllString(content, `${1}value${2}"on"`)

	// Keep "generic" as-is for HCL transformation
	cacheTypeGenericPattern := regexp.MustCompile(`(resource\s+"cloudflare_tiered_cache"[^{]+\{[^}]*\n\s*)cache_type(\s*=\s*)"generic"`)
	content = cacheTypeGenericPattern.ReplaceAllString(content, `${1}value${2}"generic"`)

	cacheTypeOffPattern := regexp.MustCompile(`(resource\s+"cloudflare_tiered_cache"[^{]+\{[^}]*\n\s*)cache_type(\s*=\s*)"off"`)
	content = cacheTypeOffPattern.ReplaceAllString(content, `${1}value${2}"off"`)

	// Also handle any remaining cache_type that isn't smart/generic/off (like variables)
	cacheTypeVarPattern := regexp.MustCompile(`(resource\s+"cloudflare_tiered_cache"[^{]+\{[^}]*\n\s*)cache_type(\s*=\s*)`)
	content = cacheTypeVarPattern.ReplaceAllString(content, `${1}value${2}`)

	return content
}
