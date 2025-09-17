package main

import (
	"fmt"
	"regexp"
	"strings"
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
	
	// Transform dynamic rules blocks to for expressions
	// v4: dynamic "rules" { for_each = ...; content { ... } }
	// v5: rules = [for ... : { ... }]
	transformDynamicRulesBlocksToAttribute(block, diags)
	
	// Transform rules attribute to ensure region_pools is a map
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
	
	// Build the map by constructing a complete map string and parsing it
	// This is much safer than manual token manipulation
	mapItems := []string{}
	
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
			
			mapItems = append(mapItems, fmt.Sprintf("    %s = %s", cleanKey, poolIDsStr))
		}
	}
	
	// Only proceed if we have valid items to transform
	if len(mapItems) == 0 {
		return
	}
	
	// Build the complete map string
	mapStr := "{\n" + strings.Join(mapItems, "\n") + "\n  }"
	
	// Parse the map as a complete unit to preserve expression structure
	tempConfig := fmt.Sprintf("temp = %s", mapStr)
	tempFile, err := hclwrite.ParseConfig([]byte(tempConfig), "temp.hcl", hcl.InitialPos)
	if err == nil && tempFile.Body().GetAttribute("temp") != nil {
		// Use the parsed attribute's tokens which preserve expression structure
		newTokens := tempFile.Body().GetAttribute("temp").Expr().BuildTokens(nil)
		block.Body().SetAttributeRaw(blockName, newTokens)
	}
}


// transformLoadBalancerRulesString applies string-level transformations to consolidate region_pools
func transformLoadBalancerRulesString(block *hclwrite.Block) {
	// Get the rules attribute
	rulesAttr := block.Body().GetAttribute("rules")
	if rulesAttr == nil {
		return
	}
	
	// Get the tokens as a string
	tokens := rulesAttr.Expr().BuildTokens(nil)
	content := string(tokens.Bytes())
	
	// Check if we have region_pools to process
	if !strings.Contains(content, "region_pools") {
		return
	}
	
	// Process each overrides block separately to avoid mixing rules
	// Find all overrides blocks
	overridesPattern := `overrides\s*=\s*\{[^{}]*(?:\{[^}]*\}[^{}]*)*\}`
	overridesMatches := regexp.MustCompile(overridesPattern).FindAllString(content, -1)
	
	for _, overridesBlock := range overridesMatches {
		// Check if this overrides block has region_pools
		if !strings.Contains(overridesBlock, "region_pools") {
			continue
		}
		
		// Find all region_pools within this specific overrides block
		regionPoolsPattern := `region_pools\s*=\s*\{[^}]*?\}`
		regionPoolsMatches := regexp.MustCompile(regionPoolsPattern).FindAllString(overridesBlock, -1)
		
		if len(regionPoolsMatches) == 0 {
			continue
		}
		
		// Parse each region_pools to extract region and pool_ids
		var regionPoolData [][]string
		for _, rpMatch := range regionPoolsMatches {
			// Extract region
			regionPattern := `region\s*=\s*"([^"]+)"`
			regionMatch := regexp.MustCompile(regionPattern).FindStringSubmatch(rpMatch)
			
			// Extract pool_ids
			poolIDsPattern := `pool_ids\s*=\s*(\[(?:[^\[\]]*|\[[^\]]*\])*\])`
			poolIDsMatch := regexp.MustCompile(poolIDsPattern).FindStringSubmatch(rpMatch)
			
			if regionMatch != nil && poolIDsMatch != nil {
				regionPoolData = append(regionPoolData, []string{rpMatch, regionMatch[1], poolIDsMatch[1]})
			}
		}
		
		// If we found region_pools to consolidate
		if len(regionPoolData) > 0 {
			// Build the consolidated map
			consolidatedMap := "region_pools = {\n"
			for i, data := range regionPoolData {
				region := data[1]
				poolIds := data[2]
				if i < len(regionPoolData)-1 {
					consolidatedMap += fmt.Sprintf("      %s = %s,\n", region, poolIds)
				} else {
					consolidatedMap += fmt.Sprintf("      %s = %s\n", region, poolIds)
				}
			}
			consolidatedMap += "    }"
			
			// Create modified overrides block
			modifiedOverrides := overridesBlock
			
			// Replace first region_pools with consolidated map
			modifiedOverrides = strings.Replace(modifiedOverrides, regionPoolData[0][0], consolidatedMap, 1)
			
			// Remove all other region_pools blocks
			for i := 1; i < len(regionPoolData); i++ {
				modifiedOverrides = strings.Replace(modifiedOverrides, regionPoolData[i][0], "", 1)
			}
			
			// Clean up extra whitespace and commas
			modifiedOverrides = regexp.MustCompile(`\n\s*\n\s*\n`).ReplaceAllString(modifiedOverrides, "\n")
			
			// Replace the original overrides block with the modified one
			content = strings.Replace(content, overridesBlock, modifiedOverrides, 1)
		}
	}
	
	// Parse the modified content and set it back
	tempConfig := fmt.Sprintf("temp = %s", content)
	tempFile, err := hclwrite.ParseConfig([]byte(tempConfig), "temp.hcl", hcl.InitialPos)
	if err == nil && tempFile.Body().GetAttribute("temp") != nil {
		block.Body().SetAttributeRaw("rules", tempFile.Body().GetAttribute("temp").Expr().BuildTokens(nil))
	}
}


// transformLoadBalancerRules transforms the rules attribute to:
// 1. Consolidate multiple region_pools objects into a single map
// 2. Fix the structure from { region = "X", pool_ids = [...] } to { "X" = [...] }
func transformLoadBalancerRules(block *hclwrite.Block, diags ast.Diagnostics) {
	// Just use the string-level transformation which handles the region_pools consolidation
	// This preserves complex expressions that can't be parsed
	transformLoadBalancerRulesString(block)
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

// extractRegionPoolIntoMap extracts region and pool_ids from a region_pool object and adds to map
func extractRegionPoolIntoMap(poolObj *hclsyntax.ObjectConsExpr, regionPoolsMap map[string]hclsyntax.Expression, diags ast.Diagnostics) {
	var regionValues []string
	var poolIDsExpr hclsyntax.Expression
	
	for _, item := range poolObj.Items {
		keyExpr, ok := item.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)
		if !ok {
			continue
		}
		
		keyName := ast.Expr2S(keyExpr, diags)
		switch keyName {
		case "region":
			// Extract region value(s)
			
			// Convert the expression to string regardless of type
			regionStr := ast.Expr2S(item.ValueExpr, diags)
			
			// Clean up the value - remove quotes and whitespace
			regionStr = strings.TrimSpace(regionStr)
			if len(regionStr) >= 2 && regionStr[0] == '"' && regionStr[len(regionStr)-1] == '"' {
				regionStr = regionStr[1 : len(regionStr)-1]
			}
			
			if regionStr != "" {
				regionValues = append(regionValues, regionStr)
			}
		case "pool_ids":
			poolIDsExpr = item.ValueExpr
		}
	}
	
	// Add to map for each region
	for _, region := range regionValues {
		if poolIDsExpr != nil {
			regionPoolsMap[region] = poolIDsExpr
		}
	}
}

// extractPoolIntoMap extracts key field and pool_ids from a pool object and adds to map
func extractPoolIntoMap(poolObj *hclsyntax.ObjectConsExpr, keyField string, poolMap map[string]hclsyntax.Expression, diags ast.Diagnostics) {
	var keyValue string
	var poolIDsExpr hclsyntax.Expression
	
	for _, item := range poolObj.Items {
		keyExpr, ok := item.KeyExpr.(*hclsyntax.ObjectConsKeyExpr)
		if !ok {
			continue
		}
		
		keyName := ast.Expr2S(keyExpr, diags)
		if keyName == keyField {
			// Extract key value
			if litExpr, ok := item.ValueExpr.(*hclsyntax.LiteralValueExpr); ok {
				if litExpr.Val.Type().FriendlyName() == "string" {
					keyValue = litExpr.Val.AsString()
				}
			}
		} else if keyName == "pool_ids" {
			poolIDsExpr = item.ValueExpr
		}
	}
	
	// Add to map
	if keyValue != "" && poolIDsExpr != nil {
		poolMap[keyValue] = poolIDsExpr
	}
}

// transformDynamicRulesBlocksToAttribute converts dynamic "rules" blocks to a rules attribute with for expression
func transformDynamicRulesBlocksToAttribute(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()
	dynamicRulesBlocks := []*hclwrite.Block{}
	
	// Find all dynamic blocks with label "rules"
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "dynamic" && len(childBlock.Labels()) > 0 && childBlock.Labels()[0] == "rules" {
			dynamicRulesBlocks = append(dynamicRulesBlocks, childBlock)
		}
	}
	
	if len(dynamicRulesBlocks) == 0 {
		return
	}
	
	// Process each dynamic rules block
	for _, dynBlock := range dynamicRulesBlocks {
		dynBody := dynBlock.Body()
		
		// Get the for_each expression
		forEachAttr := dynBody.GetAttribute("for_each")
		if forEachAttr == nil {
			continue
		}
		
		// Find the content block
		var contentBlock *hclwrite.Block
		for _, cb := range dynBody.Blocks() {
			if cb.Type() == "content" {
				contentBlock = cb
				break
			}
		}
		
		if contentBlock == nil {
			continue
		}
		
		// Get the for_each expression tokens
		forEachTokens := forEachAttr.Expr().BuildTokens(nil)
		
		// Build the for expression
		// rules = [for rule in <for_each> : { <content> }]
		tokens := hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("for")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" rules")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" in")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" ")},
		}
		
		// Add the for_each expression
		tokens = append(tokens, forEachTokens...)
		tokens = append(tokens,
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" ")},
			&hclwrite.Token{Type: hclsyntax.TokenColon, Bytes: []byte(":")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" ")},
			&hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
			&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		)
		
		// Add content block attributes as object properties
		contentBody := contentBlock.Body()
		
		// Process attributes in the content block
		attrs := contentBody.Attributes()
		first := true
		for name, attr := range attrs {
			if !first {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			}
			first = false
			
			// Add indentation
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
			
			// Add attribute name
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			
			// Get the attribute value expression tokens
			attrTokens := attr.Expr().BuildTokens(nil)
			
			// For rules, we use a comprehensive replacement since "rules" is used as the entire iterator
			// not "rules.key" and "rules.value" like in the origins case
			// We need to replace all instances of "rules.value" -> "rules" and "rules.key" -> "rules"
			attrStr := string(attrTokens.Bytes())
			
			// Use regex to replace rules.value and rules.key more comprehensively
			// This handles cases like [rules.value], ${rules.value}, (rules.value), etc.
			rulesValuePattern := regexp.MustCompile(`\brules\.value\b`)
			rulesKeyPattern := regexp.MustCompile(`\brules\.key\b`)
			attrStr = rulesValuePattern.ReplaceAllString(attrStr, "rules")
			attrStr = rulesKeyPattern.ReplaceAllString(attrStr, "rules")
			
			// Parse the modified expression
			tempConfig := fmt.Sprintf("temp = %s", attrStr)
			tempFile, parseErr := hclwrite.ParseConfig([]byte(tempConfig), "temp.hcl", hcl.InitialPos)
			if parseErr == nil && tempFile.Body().GetAttribute("temp") != nil {
				tokens = append(tokens, tempFile.Body().GetAttribute("temp").Expr().BuildTokens(nil)...)
			} else {
				// Fall back to original tokens if parsing fails
				tokens = append(tokens, attrTokens...)
			}
		}
		
		// Close the object and list
		tokens = append(tokens,
			&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")},
			&hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
			&hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
		)
		
		// Remove the dynamic block
		body.RemoveBlock(dynBlock)
		
		// Add the rules attribute
		body.SetAttributeRaw("rules", tokens)
	}
}

