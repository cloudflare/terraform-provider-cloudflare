package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func isLoadBalancerResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_load_balancer"
}

// transformLoadBalancerFile applies cloudflare_load_balancer specific transformations to an HCL file
func transformLoadBalancerFile(file *hclwrite.File) {
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_load_balancer" {
			diags := ast.Diagnostics{}
			transformLoadBalancerBlock(block, diags)
		}
	}
}

// transformLoadBalancerBlock handles block-level transformations for cloudflare_load_balancer
func transformLoadBalancerBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Transform pool blocks (region_pools, country_pools, pop_pools) to maps
	// v4: Multiple blocks with same name
	// v5: Single attribute as a map
	transformPoolBlocksToMap(block, "region_pools", "region")
	transformPoolBlocksToMap(block, "country_pools", "country")
	transformPoolBlocksToMap(block, "pop_pools", "pop")
}

// transformPoolBlocksToMap converts multiple blocks to a single map attribute
func transformPoolBlocksToMap(block *hclwrite.Block, blockName string, keyField string) {
	blocks := block.Body().Blocks()
	poolBlocks := []*hclwrite.Block{}
	
	// Collect all blocks with the given name
	for _, b := range blocks {
		if b.Type() == blockName {
			poolBlocks = append(poolBlocks, b)
		}
	}
	
	if len(poolBlocks) == 0 {
		return
	}
	
	// Build the map value
	poolMap := make(map[string]cty.Value)
	for _, poolBlock := range poolBlocks {
		// Get the key value (region, country, or pop)
		keyAttr := poolBlock.Body().GetAttribute(keyField)
		if keyAttr == nil {
			continue
		}
		
		// Get the key as a string
		keyTokens := keyAttr.Expr().BuildTokens(nil)
		keyStr := string(hclwrite.TokensForValue(cty.StringVal(string(keyTokens.Bytes()))).Bytes())
		// Remove quotes from the key
		if len(keyStr) >= 2 && keyStr[0] == '"' && keyStr[len(keyStr)-1] == '"' {
			keyStr = keyStr[1:len(keyStr)-1]
		}
		
		// Get the pool_ids value
		poolIDsAttr := poolBlock.Body().GetAttribute("pool_ids")
		if poolIDsAttr != nil {
			// Keep the pool_ids expression as-is
			poolMap[keyStr] = cty.DynamicVal // We'll use the raw expression
		}
	}
	
	// Remove all the blocks
	for _, poolBlock := range poolBlocks {
		block.Body().RemoveBlock(poolBlock)
	}
	
	// Create the new map attribute
	if len(poolMap) > 0 {
		// Build the map expression manually to preserve references
		tokens := hclwrite.Tokens{
			{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		}
		
		first := true
		for _, poolBlock := range poolBlocks {
			keyAttr := poolBlock.Body().GetAttribute(keyField)
			poolIDsAttr := poolBlock.Body().GetAttribute("pool_ids")
			if keyAttr == nil || poolIDsAttr == nil {
				continue
			}
			
			if !first {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			}
			first = false
			
			// Add indentation
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
			
			// Add the key
			keyTokens := keyAttr.Expr().BuildTokens(nil)
			keyStr := string(keyTokens.Bytes())
			// Remove quotes if present and re-add them as identifiers
			if len(keyStr) >= 2 && keyStr[0] == '"' && keyStr[len(keyStr)-1] == '"' {
				keyStr = keyStr[1:len(keyStr)-1]
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(keyStr)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			
			// Add the pool_ids value
			poolIDsTokens := poolIDsAttr.Expr().BuildTokens(nil)
			tokens = append(tokens, poolIDsTokens...)
		}
		
		tokens = append(tokens,
			&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")},
			&hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
		)
		
		block.Body().SetAttributeRaw(blockName, tokens)
	}
}

