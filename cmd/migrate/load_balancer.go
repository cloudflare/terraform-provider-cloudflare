package main

import (
	"fmt"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
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
	
	// Transform rules attribute to ensure region_pools.region is a list
	transformLoadBalancerRules(block, diags)
}

// transformPoolBlocksToMap converts multiple blocks or arrays to a single map attribute
func transformPoolBlocksToMap(block *hclwrite.Block, blockName string, keyField string) {
	blocks := block.Body().Blocks()
	poolBlocks := []*hclwrite.Block{}
	
	// Collect all blocks with the given name
	for _, b := range blocks {
		if b.Type() == blockName {
			poolBlocks = append(poolBlocks, b)
		}
	}
	
	// If we have blocks, transform them to map
	if len(poolBlocks) > 0 {
		transformPoolBlocksToMapFromBlocks(block, blockName, keyField, poolBlocks)
		return
	}
	
	// Check if there's an array attribute instead (from Grit transformation)
	attr := block.Body().GetAttribute(blockName)
	if attr != nil {
		transformPoolArrayToMap(block, blockName, keyField, attr)
	}
}

// transformPoolBlocksToMapFromBlocks handles the block-to-map conversion
func transformPoolBlocksToMapFromBlocks(block *hclwrite.Block, blockName string, keyField string, poolBlocks []*hclwrite.Block) {
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
	
	// Remove all the blocks
	for _, poolBlock := range poolBlocks {
		block.Body().RemoveBlock(poolBlock)
	}
	
	block.Body().SetAttributeRaw(blockName, tokens)
}

// transformPoolArrayToMap handles array-to-map conversion (from Grit transformations)
func transformPoolArrayToMap(block *hclwrite.Block, blockName string, keyField string, attr *hclwrite.Attribute) {
	// Parse the array expression to get the AST
	diags := ast.NewDiagnostics()
	arrayExpr := ast.WriteExpr2Expr(attr.Expr(), diags)
	
	// Check if it's a tuple (array)
	tup, ok := arrayExpr.(*hclsyntax.TupleConsExpr)
	if !ok {
		// Not an array - continue execution, nothing to transform
		return
	}
	
	// Build the map expression by reconstructing from original tokens
	// This preserves the original expressions
	mapStr := "{\n"
	
	validItems := []string{}
	for _, itemExpr := range tup.Exprs {
		// Each item should be an object with keyField and pool_ids
		itemObj, ok := itemExpr.(*hclsyntax.ObjectConsExpr)
		if !ok {
			continue
		}
		
		var keyValue string
		var poolIDsExpr hclsyntax.Expression
		for _, item := range itemObj.Items {
			keyExpr, ok := item.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)
			if !ok {
				continue
			}
			
			keyName := ast.Expr2S(keyExpr, diags)
			if keyName == keyField {
				keyValue = ast.Expr2S(item.ValueExpr, diags)
			} else if keyName == "pool_ids" {
				poolIDsExpr = item.ValueExpr
			}
		}
		
		if keyValue != "" && poolIDsExpr != nil {
			// Clean up the key
			cleanKey := keyValue
			if len(cleanKey) >= 2 && cleanKey[0] == '"' && cleanKey[len(cleanKey)-1] == '"' {
				cleanKey = cleanKey[1:len(cleanKey)-1]
			}
			
			// Get the pool_ids expression as string (preserves references)
			poolIDsStr := ast.Expr2S(poolIDsExpr, diags)
			
			validItems = append(validItems, fmt.Sprintf("    %s = %s", cleanKey, poolIDsStr))
		}
	}
	
	// Only proceed if we have valid items to transform
	if len(validItems) == 0 {
		return
	}
	
	// Add all valid items
	for i, item := range validItems {
		mapStr += item
		if i < len(validItems)-1 {
			mapStr += "\n"
		}
	}
	mapStr += "\n  }"
	
	// Parse the new map expression and set it
	tempConfig := fmt.Sprintf("temp = %s", mapStr)
	tempFile, err := hclwrite.ParseConfig([]byte(tempConfig), "temp.hcl", hcl.InitialPos)
	if err == nil && tempFile.Body().GetAttribute("temp") != nil {
		newTokens := tempFile.Body().GetAttribute("temp").Expr().BuildTokens(nil)
		block.Body().SetAttributeRaw(blockName, newTokens)
	}
}


// transformLoadBalancerRules transforms the rules attribute to ensure region_pools.region is a list
func transformLoadBalancerRules(block *hclwrite.Block, diags ast.Diagnostics) {
	// Get the rules attribute
	rulesAttr := block.Body().GetAttribute("rules")
	if rulesAttr == nil {
		return
	}
	
	// Parse the rules expression
	rulesExpr := ast.WriteExpr2Expr(rulesAttr.Expr(), diags)
	
	// Check if it's a tuple (list of rules)
	tup, ok := rulesExpr.(*hclsyntax.TupleConsExpr)
	if !ok {
		// Can't parse - keep as is since we don't know the structure
		return
	}
	
	// Process each rule in the list
	modified := false
	for _, ruleExpr := range tup.Exprs {
		// Check if the rule is an object
		ruleObj, ok := ruleExpr.(*hclsyntax.ObjectConsExpr)
		if !ok {
			continue
		}
		
		// Find the overrides attribute
		for _, item := range ruleObj.Items {
			keyExpr, ok := item.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)
			if !ok {
				continue
			}
			
			if ast.Expr2S(keyExpr, diags) == "overrides" {
				// Process the overrides object
				overridesObj, ok := item.ValueExpr.(*hclsyntax.ObjectConsExpr)
				if !ok {
					continue
				}
				
				// Find region_pools in overrides
				for _, overrideItem := range overridesObj.Items {
					overrideKeyExpr, ok := overrideItem.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)
					if !ok {
						continue
					}
					
					if ast.Expr2S(overrideKeyExpr, diags) == "region_pools" {
						// Transform region_pools
						if transformRegionPools(&overrideItem.ValueExpr, diags) {
							modified = true
						}
					}
				}
			}
		}
	}
	
	// Only update if we made changes
	if modified {
		// Convert the modified expression back to a string and reparse as HCL
		exprStr := ast.Expr2S(rulesExpr, diags)
		// Use hclwrite to parse and format the expression properly
		newExpr, err := hclwrite.ParseConfig([]byte("temp = "+exprStr), "temp.hcl", hcl.InitialPos)
		if err == nil && newExpr.Body().Attributes()["temp"] != nil {
			block.Body().SetAttributeRaw("rules", newExpr.Body().Attributes()["temp"].Expr().BuildTokens(nil))
		}
	}
}

// transformRegionPools ensures that each region_pools element has region as a list
func transformRegionPools(expr *hclsyntax.Expression, diags ast.Diagnostics) bool {
	modified := false
	
	// Check if it's a tuple (list of region_pools)
	if tup, ok := (*expr).(*hclsyntax.TupleConsExpr); ok {
		// Process each region_pool in the list
		for _, poolExpr := range tup.Exprs {
			if transformSingleRegionPool(poolExpr, diags) {
				modified = true
			}
		}
		return modified
	}
	
	// Check if it's a single object (region_pools = { ... })
	if obj, ok := (*expr).(*hclsyntax.ObjectConsExpr); ok {
		if transformSingleRegionPool(obj, diags) {
			modified = true
		}
		return modified
	}
	
	return false
}

// transformSingleRegionPool transforms a single region pool object
func transformSingleRegionPool(poolExpr hclsyntax.Expression, diags ast.Diagnostics) bool {
	// Check if the pool is an object
	poolObj, ok := poolExpr.(*hclsyntax.ObjectConsExpr)
	if !ok {
		return false
	}
	
	modified := false
	// Find the region attribute
	for i, item := range poolObj.Items {
		keyExpr, ok := item.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)
		if !ok {
			continue
		}
		
		if ast.Expr2S(keyExpr, diags) == "region" {
			// Check if region is already a list
			if _, isList := item.ValueExpr.(*hclsyntax.TupleConsExpr); !isList {
				// Convert single string to list
				poolObj.Items[i].ValueExpr = &hclsyntax.TupleConsExpr{
					Exprs: []hclsyntax.Expression{item.ValueExpr},
				}
				modified = true
			}
		}
	}
	return modified
}

