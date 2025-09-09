package main

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isCloudflareListResource checks if a block is a cloudflare_list resource
func isCloudflareListResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_list"
}

// transformCloudflareListBlock is the main entry point for transforming cloudflare_list resources
// Handles both static item blocks and dynamic blocks, transforming them to the v5 items attribute
func transformCloudflareListBlock(block *hclwrite.Block) {
	body := block.Body()
	diags := ast.NewDiagnostics()
	
	// Get the list kind to determine item structure
	kindAttr := body.GetAttribute("kind")
	if kindAttr == nil {
		return // Can't transform without knowing the kind
	}
	
	kindValue := extractStringValue(*kindAttr.Expr())
	if kindValue == "" {
		return // Can't determine kind
	}
	
	// Check for problematic patterns first
	checkAndWarnProblematicPatterns(block, diags)
	
	// Collect all item-related blocks
	var staticItemBlocks []*hclwrite.Block
	var dynamicItemBlocks []*hclwrite.Block
	
	for _, b := range body.Blocks() {
		if b.Type() == "item" {
			staticItemBlocks = append(staticItemBlocks, b)
		} else if b.Type() == "dynamic" && len(b.Labels()) > 0 && b.Labels()[0] == "item" {
			dynamicItemBlocks = append(dynamicItemBlocks, b)
		}
	}
	
	// Handle transformation based on what we found
	if len(dynamicItemBlocks) > 0 {
		// Has dynamic blocks - use AST-based transformation
		transformListWithDynamicBlocks(body, staticItemBlocks, dynamicItemBlocks, kindValue, diags)
	} else if len(staticItemBlocks) > 0 {
		// Only static blocks - use simple transformation
		transformStaticItemBlocks(body, staticItemBlocks, kindValue)
	}
	
	// Add any accumulated warnings
	if len(diags.ComplicatedHCL) > 0 {
		addDiagnosticsAsComments(body, diags)
	}
}

// transformListWithDynamicBlocks handles lists with dynamic blocks using AST manipulation
func transformListWithDynamicBlocks(body *hclwrite.Body, staticBlocks, dynamicBlocks []*hclwrite.Block, kind string, diags ast.Diagnostics) {
	// Parse the body to get AST representation
	bodyBytes := hclwrite.Format(body.BuildTokens(nil).Bytes())
	syntaxBody := ast.ParseIntoSyntaxBody(bodyBytes, "cloudflare_list", diags)
	if syntaxBody == nil {
		return
	}
	
	// Build items expression
	var itemsExpr hclsyntax.Expression
	
	if len(dynamicBlocks) == 1 && len(staticBlocks) == 0 {
		// Single dynamic block - create for expression
		itemsExpr = buildForExpressionFromDynamic(dynamicBlocks[0], kind, syntaxBody, diags)
	} else if len(dynamicBlocks) > 0 || len(staticBlocks) > 0 {
		// Mixed or multiple - create concat expression
		itemsExpr = buildConcatExpression(staticBlocks, dynamicBlocks, kind, syntaxBody, diags)
	}
	
	if itemsExpr == nil {
		// Couldn't build expression - add warning comment
		comment := fmt.Sprintf("\n  # MIGRATION WARNING: Could not transform dynamic blocks automatically")
		tokens := hclwrite.Tokens{
			&hclwrite.Token{
				Type:  hclsyntax.TokenComment,
				Bytes: []byte(comment),
			},
		}
		body.AppendUnstructuredTokens(tokens)
		return
	}
	
	// Remove all item-related blocks
	for _, block := range staticBlocks {
		body.RemoveBlock(block)
	}
	for _, block := range dynamicBlocks {
		body.RemoveBlock(block)
	}
	
	// Add items attribute using the expression
	addItemsAttributeFromExpression(body, itemsExpr, diags)
}

// buildForExpressionFromDynamic creates a for expression from a dynamic block
func buildForExpressionFromDynamic(dynBlock *hclwrite.Block, kind string, parentBody *hclsyntax.Body, diags ast.Diagnostics) hclsyntax.Expression {
	dynBody := dynBlock.Body()
	
	// Get for_each expression
	forEachAttr := dynBody.GetAttribute("for_each")
	if forEachAttr == nil {
		return nil
	}
	
	// Parse for_each expression to AST
	forEachBytes := forEachAttr.Expr().BuildTokens(nil).Bytes()
	forEachExpr, d := hclsyntax.ParseExpression(forEachBytes, "for_each", hcl.InitialPos)
	if d.HasErrors() {
		diags.HclDiagnostics.Extend(d)
		return nil
	}
	
	// Get iterator name (default is the block label)
	iteratorName := dynBlock.Labels()[0]
	if iterAttr := dynBody.GetAttribute("iterator"); iterAttr != nil {
		iteratorName = extractStringValue(*iterAttr.Expr())
	}
	
	// Find content block
	var contentBlock *hclwrite.Block
	for _, b := range dynBody.Blocks() {
		if b.Type() == "content" {
			contentBlock = b
			break
		}
	}
	
	if contentBlock == nil {
		return nil
	}
	
	// Build object expression from content
	objExpr := buildObjectFromContentBlock(contentBlock, kind, iteratorName, diags)
	if objExpr == nil {
		return nil
	}
	
	// Create the for expression
	return &hclsyntax.ForExpr{
		KeyVar:   "",  // No key variable for list comprehension
		ValVar:   iteratorName,
		CollExpr: forEachExpr,
		ValExpr:  objExpr,
	}
}

// buildObjectFromContentBlock creates an object expression from a content block
func buildObjectFromContentBlock(contentBlock *hclwrite.Block, kind string, iteratorName string, diags ast.Diagnostics) hclsyntax.Expression {
	contentBody := contentBlock.Body()
	items := []hclsyntax.ObjectConsItem{}
	
	// Process value block first (to get ip, asn, hostname, redirect before comment)
	for _, vBlock := range contentBody.Blocks() {
		if vBlock.Type() == "value" {
			valueItems := extractValueBlockItems(vBlock, kind, iteratorName, diags)
			items = append(items, valueItems...)
		}
	}
	
	// Then add comment attribute
	if commentAttr := contentBody.GetAttribute("comment"); commentAttr != nil {
		commentBytes := commentAttr.Expr().BuildTokens(nil).Bytes()
		commentExpr, d := hclsyntax.ParseExpression(commentBytes, "comment", hcl.InitialPos)
		if !d.HasErrors() {
			// Handle iterator references appropriately
			commentExpr = stripIteratorValueSuffix(commentExpr, iteratorName)
			items = append(items, hclsyntax.ObjectConsItem{
				KeyExpr:   ast.NewKeyExpr("comment"),
				ValueExpr: commentExpr,
			})
		}
		diags.HclDiagnostics.Extend(d)
	}
	
	if len(items) == 0 {
		return nil
	}
	
	return &hclsyntax.ObjectConsExpr{Items: items}
}

// extractValueBlockItems extracts items from a value block based on list kind
func extractValueBlockItems(vBlock *hclwrite.Block, kind string, iteratorName string, diags ast.Diagnostics) []hclsyntax.ObjectConsItem {
	vBody := vBlock.Body()
	items := []hclsyntax.ObjectConsItem{}
	
	switch kind {
	case "ip":
		if ipAttr := vBody.GetAttribute("ip"); ipAttr != nil {
			ipBytes := ipAttr.Expr().BuildTokens(nil).Bytes()
			ipExpr, d := hclsyntax.ParseExpression(ipBytes, "ip", hcl.InitialPos)
			if !d.HasErrors() {
				// Handle iterator references appropriately
				ipExpr = stripIteratorValueSuffix(ipExpr, iteratorName)
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr:   ast.NewKeyExpr("ip"),
					ValueExpr: ipExpr,
				})
			}
			diags.HclDiagnostics.Extend(d)
		}
		
	case "asn":
		if asnAttr := vBody.GetAttribute("asn"); asnAttr != nil {
			asnBytes := asnAttr.Expr().BuildTokens(nil).Bytes()
			asnExpr, d := hclsyntax.ParseExpression(asnBytes, "asn", hcl.InitialPos)
			if !d.HasErrors() {
				// Handle iterator references appropriately
				asnExpr = stripIteratorValueSuffix(asnExpr, iteratorName)
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr:   ast.NewKeyExpr("asn"),
					ValueExpr: asnExpr,
				})
			}
			diags.HclDiagnostics.Extend(d)
		}
		
	case "hostname":
		// Handle hostname nested structure
		for _, hBlock := range vBody.Blocks() {
			if hBlock.Type() == "hostname" {
				hostnameObj := buildHostnameObjectWithIterator(hBlock, iteratorName, diags)
				if hostnameObj != nil {
					items = append(items, hclsyntax.ObjectConsItem{
						KeyExpr:   ast.NewKeyExpr("hostname"),
						ValueExpr: hostnameObj,
					})
				}
			}
		}
		
	case "redirect":
		// Handle redirect nested structure with boolean conversions
		for _, rBlock := range vBody.Blocks() {
			if rBlock.Type() == "redirect" {
				redirectObj := buildRedirectObjectWithIterator(rBlock, iteratorName, diags)
				if redirectObj != nil {
					items = append(items, hclsyntax.ObjectConsItem{
						KeyExpr:   ast.NewKeyExpr("redirect"),
						ValueExpr: redirectObj,
					})
				}
			}
		}
	}
	
	return items
}

// buildHostnameObject creates a hostname object expression (for static blocks)
func buildHostnameObject(hBlock *hclwrite.Block, diags ast.Diagnostics) hclsyntax.Expression {
	return buildHostnameObjectWithIterator(hBlock, "", diags)
}

// buildHostnameObjectWithIterator creates a hostname object expression with iterator handling
func buildHostnameObjectWithIterator(hBlock *hclwrite.Block, iteratorName string, diags ast.Diagnostics) hclsyntax.Expression {
	hBody := hBlock.Body()
	items := []hclsyntax.ObjectConsItem{}
	
	if urlAttr := hBody.GetAttribute("url_hostname"); urlAttr != nil {
		urlBytes := urlAttr.Expr().BuildTokens(nil).Bytes()
		urlExpr, d := hclsyntax.ParseExpression(urlBytes, "url_hostname", hcl.InitialPos)
		if !d.HasErrors() {
			// Handle iterator references if present
			if iteratorName != "" {
				urlExpr = stripIteratorValueSuffix(urlExpr, iteratorName)
			}
			items = append(items, hclsyntax.ObjectConsItem{
				KeyExpr:   ast.NewKeyExpr("url_hostname"),
				ValueExpr: urlExpr,
			})
		}
		diags.HclDiagnostics.Extend(d)
	}
	
	if len(items) == 0 {
		return nil
	}
	
	return &hclsyntax.ObjectConsExpr{Items: items}
}

// buildRedirectObject creates a redirect object expression with boolean conversions (for static blocks)
func buildRedirectObject(rBlock *hclwrite.Block, diags ast.Diagnostics) hclsyntax.Expression {
	return buildRedirectObjectWithIterator(rBlock, "", diags)
}

// buildRedirectObjectWithIterator creates a redirect object expression with iterator handling
func buildRedirectObjectWithIterator(rBlock *hclwrite.Block, iteratorName string, diags ast.Diagnostics) hclsyntax.Expression {
	rBody := rBlock.Body()
	items := []hclsyntax.ObjectConsItem{}
	
	// Required fields - need iterator handling
	addStringAttributeWithIterator(rBody, "source_url", &items, iteratorName, diags)
	addStringAttributeWithIterator(rBody, "target_url", &items, iteratorName, diags)
	
	// Boolean fields that need conversion from "enabled"/"disabled"
	boolFields := []string{
		"include_subdomains",
		"subpath_matching",
		"preserve_query_string",
		"preserve_path_suffix",
	}
	
	for _, field := range boolFields {
		if attr := rBody.GetAttribute(field); attr != nil {
			value := extractStringValue(*attr.Expr())
			var boolExpr hclsyntax.Expression
			if value == "enabled" {
				boolExpr = &hclsyntax.LiteralValueExpr{Val: cty.BoolVal(true)}
			} else if value == "disabled" {
				boolExpr = &hclsyntax.LiteralValueExpr{Val: cty.BoolVal(false)}
			} else {
				// Keep original expression if not a simple string
				bytes := attr.Expr().BuildTokens(nil).Bytes()
				var d hcl.Diagnostics
				boolExpr, d = hclsyntax.ParseExpression(bytes, field, hcl.InitialPos)
				if !d.HasErrors() && iteratorName != "" {
					// Handle iterator references in complex expressions
					boolExpr = stripIteratorValueSuffix(boolExpr, iteratorName)
				}
				diags.HclDiagnostics.Extend(d)
			}
			
			if boolExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr:   ast.NewKeyExpr(field),
					ValueExpr: boolExpr,
				})
			}
		}
	}
	
	// Optional status_code - need iterator handling
	addNumberAttributeWithIterator(rBody, "status_code", &items, iteratorName, diags)
	
	if len(items) == 0 {
		return nil
	}
	
	return &hclsyntax.ObjectConsExpr{Items: items}
}

// Helper functions

func addStringAttribute(body *hclwrite.Body, name string, items *[]hclsyntax.ObjectConsItem, diags ast.Diagnostics) {
	addStringAttributeWithIterator(body, name, items, "", diags)
}

func addStringAttributeWithIterator(body *hclwrite.Body, name string, items *[]hclsyntax.ObjectConsItem, iteratorName string, diags ast.Diagnostics) {
	if attr := body.GetAttribute(name); attr != nil {
		bytes := attr.Expr().BuildTokens(nil).Bytes()
		expr, d := hclsyntax.ParseExpression(bytes, name, hcl.InitialPos)
		if !d.HasErrors() {
			// Handle iterator references if present
			if iteratorName != "" {
				expr = stripIteratorValueSuffix(expr, iteratorName)
			}
			*items = append(*items, hclsyntax.ObjectConsItem{
				KeyExpr:   ast.NewKeyExpr(name),
				ValueExpr: expr,
			})
		}
		diags.HclDiagnostics.Extend(d)
	}
}

func addNumberAttribute(body *hclwrite.Body, name string, items *[]hclsyntax.ObjectConsItem, diags ast.Diagnostics) {
	addNumberAttributeWithIterator(body, name, items, "", diags)
}

func addNumberAttributeWithIterator(body *hclwrite.Body, name string, items *[]hclsyntax.ObjectConsItem, iteratorName string, diags ast.Diagnostics) {
	if attr := body.GetAttribute(name); attr != nil {
		bytes := attr.Expr().BuildTokens(nil).Bytes()
		expr, d := hclsyntax.ParseExpression(bytes, name, hcl.InitialPos)
		if !d.HasErrors() {
			// Handle iterator references if present
			if iteratorName != "" {
				expr = stripIteratorValueSuffix(expr, iteratorName)
			}
			*items = append(*items, hclsyntax.ObjectConsItem{
				KeyExpr:   ast.NewKeyExpr(name),
				ValueExpr: expr,
			})
		}
		diags.HclDiagnostics.Extend(d)
	}
}

// buildConcatExpression builds a concat expression for mixed static/dynamic items
func buildConcatExpression(staticBlocks, dynamicBlocks []*hclwrite.Block, kind string, syntaxBody *hclsyntax.Body, diags ast.Diagnostics) hclsyntax.Expression {
	var exprs []hclsyntax.Expression
	
	// Add static items if any
	if len(staticBlocks) > 0 {
		staticExpr := buildStaticItemsExpression(staticBlocks, kind, diags)
		if staticExpr != nil {
			exprs = append(exprs, staticExpr)
		}
	}
	
	// Add dynamic expressions
	for _, dynBlock := range dynamicBlocks {
		forExpr := buildForExpressionFromDynamic(dynBlock, kind, syntaxBody, diags)
		if forExpr != nil {
			exprs = append(exprs, forExpr)
		}
	}
	
	if len(exprs) == 0 {
		return nil
	}
	
	if len(exprs) == 1 {
		return exprs[0]
	}
	
	// Create concat function call
	return &hclsyntax.FunctionCallExpr{
		Name: "concat",
		Args: exprs,
	}
}

// buildStaticItemsExpression creates a tuple expression from static item blocks
func buildStaticItemsExpression(blocks []*hclwrite.Block, kind string, diags ast.Diagnostics) hclsyntax.Expression {
	var exprs []hclsyntax.Expression
	
	for _, block := range blocks {
		objExpr := buildObjectFromItemBlock(block, kind, diags)
		if objExpr != nil {
			exprs = append(exprs, objExpr)
		}
	}
	
	if len(exprs) == 0 {
		return nil
	}
	
	return &hclsyntax.TupleConsExpr{Exprs: exprs}
}

// buildObjectFromItemBlock creates an object expression from a static item block
func buildObjectFromItemBlock(block *hclwrite.Block, kind string, diags ast.Diagnostics) hclsyntax.Expression {
	body := block.Body()
	items := []hclsyntax.ObjectConsItem{}
	
	// Process value block first (to get ip, asn, hostname, redirect before comment)
	for _, vBlock := range body.Blocks() {
		if vBlock.Type() == "value" {
			valueItems := extractValueBlockItems(vBlock, kind, "", diags)
			items = append(items, valueItems...)
		}
	}
	
	// Then handle comment
	if commentAttr := body.GetAttribute("comment"); commentAttr != nil {
		commentBytes := commentAttr.Expr().BuildTokens(nil).Bytes()
		commentExpr, d := hclsyntax.ParseExpression(commentBytes, "comment", hcl.InitialPos)
		if !d.HasErrors() {
			items = append(items, hclsyntax.ObjectConsItem{
				KeyExpr:   ast.NewKeyExpr("comment"),
				ValueExpr: commentExpr,
			})
		}
		diags.HclDiagnostics.Extend(d)
	}
	
	if len(items) == 0 {
		return nil
	}
	
	return &hclsyntax.ObjectConsExpr{Items: items}
}

// transformStaticItemBlocks handles pure static item blocks using AST for proper ordering
func transformStaticItemBlocks(body *hclwrite.Body, itemBlocks []*hclwrite.Block, kind string) {
	if len(itemBlocks) == 0 {
		return
	}
	
	// Use AST-based approach for proper attribute ordering
	diags := ast.Diagnostics{}
	
	// Build AST expression for items
	itemsExpr := buildStaticItemsExpression(itemBlocks, kind, diags)
	
	// Remove all item blocks
	for _, itemBlock := range itemBlocks {
		body.RemoveBlock(itemBlock)
	}
	
	// Add the new items attribute if we have an expression
	if itemsExpr != nil {
		// Convert AST expression to hclwrite expression
		writeExpr := ast.Expr2WriteExpr(itemsExpr, diags)
		body.SetAttributeRaw("items", writeExpr.BuildTokens(nil))
	}
}

// transformItemBlockSimple transforms a single item block to a cty.Value for the items array
func transformItemBlockSimple(itemBody *hclwrite.Body, kind string) cty.Value {
	itemMap := make(map[string]cty.Value)

	// Handle comment if present
	if commentAttr := itemBody.GetAttribute("comment"); commentAttr != nil {
		commentValue := extractStringValue(*commentAttr.Expr())
		if commentValue != "" {
			itemMap["comment"] = cty.StringVal(commentValue)
		}
	}

	// Process value block based on list kind
	for _, valueBlock := range itemBody.Blocks() {
		if valueBlock.Type() == "value" {
			valueBody := valueBlock.Body()
			
			switch kind {
			case "ip":
				if ipAttr := valueBody.GetAttribute("ip"); ipAttr != nil {
					ipValue := extractStringValue(*ipAttr.Expr())
					if ipValue != "" {
						itemMap["ip"] = cty.StringVal(ipValue)
					}
				}
				
			case "asn":
				if asnAttr := valueBody.GetAttribute("asn"); asnAttr != nil {
					// ASN can be a number or string
					asnValue := extractStringValue(*asnAttr.Expr())
					if asnValue != "" {
						itemMap["asn"] = cty.StringVal(asnValue)
					}
				}
				
			case "hostname":
				// Handle nested hostname structure
				for _, hostnameBlock := range valueBody.Blocks() {
					if hostnameBlock.Type() == "hostname" {
						hostnameBody := hostnameBlock.Body()
						hostnameMap := make(map[string]cty.Value)
						
						if urlAttr := hostnameBody.GetAttribute("url_hostname"); urlAttr != nil {
							urlValue := extractStringValue(*urlAttr.Expr())
							if urlValue != "" {
								hostnameMap["url_hostname"] = cty.StringVal(urlValue)
							}
						}
						
						if len(hostnameMap) > 0 {
							itemMap["hostname"] = cty.ObjectVal(hostnameMap)
						}
					}
				}
				
			case "redirect":
				// Handle nested redirect structure with boolean conversions
				for _, redirectBlock := range valueBody.Blocks() {
					if redirectBlock.Type() == "redirect" {
						redirectBody := redirectBlock.Body()
						redirectMap := make(map[string]cty.Value)
						
						// Required fields
						if sourceAttr := redirectBody.GetAttribute("source_url"); sourceAttr != nil {
							sourceValue := extractStringValue(*sourceAttr.Expr())
							if sourceValue != "" {
								redirectMap["source_url"] = cty.StringVal(sourceValue)
							}
						}
						
						if targetAttr := redirectBody.GetAttribute("target_url"); targetAttr != nil {
							targetValue := extractStringValue(*targetAttr.Expr())
							if targetValue != "" {
								redirectMap["target_url"] = cty.StringVal(targetValue)
							}
						}
						
						// Boolean fields that need conversion from "enabled"/"disabled"
						boolFields := []string{
							"include_subdomains",
							"subpath_matching", 
							"preserve_query_string",
							"preserve_path_suffix",
						}
						
						for _, field := range boolFields {
							if attr := redirectBody.GetAttribute(field); attr != nil {
								value := extractStringValue(*attr.Expr())
								if value == "enabled" {
									redirectMap[field] = cty.BoolVal(true)
								} else if value == "disabled" {
									redirectMap[field] = cty.BoolVal(false)
								}
							}
						}
						
						// Optional status_code
						if statusAttr := redirectBody.GetAttribute("status_code"); statusAttr != nil {
							// Try to extract as number
							tokens := statusAttr.Expr().BuildTokens(nil)
							statusStr := string(tokens.Bytes())
							statusStr = strings.TrimSpace(statusStr)
							
							// Try to parse as number
							if statusStr != "" {
								redirectMap["status_code"] = cty.StringVal(statusStr)
							}
						}
						
						if len(redirectMap) > 0 {
							itemMap["redirect"] = cty.ObjectVal(redirectMap)
						}
					}
				}
			}
		}
	}
	
	if len(itemMap) == 0 {
		return cty.NullVal(cty.EmptyObject)
	}
	
	return cty.ObjectVal(itemMap)
}

// addItemsAttributeFromExpression adds an items attribute using an AST expression
func addItemsAttributeFromExpression(body *hclwrite.Body, expr hclsyntax.Expression, diags ast.Diagnostics) {
	// Convert the AST expression to string
	exprStr := ast.Expr2S(expr, diags)
	
	// Create the full attribute HCL
	attrHCL := fmt.Sprintf("items = %s", exprStr)
	
	// Parse it to get proper HCL tokens
	file, parseDiags := hclwrite.ParseConfig([]byte(attrHCL), "items", hcl.InitialPos)
	if parseDiags.HasErrors() {
		diags.HclDiagnostics.Extend(parseDiags)
		// Fallback: add as comment
		comment := fmt.Sprintf("\n# MIGRATION ERROR: Could not create items attribute\n# Attempted: %s", attrHCL)
		tokens := hclwrite.Tokens{
			&hclwrite.Token{
				Type:  hclsyntax.TokenComment,
				Bytes: []byte(comment),
			},
		}
		body.AppendUnstructuredTokens(tokens)
		return
	}
	
	// Extract the items attribute from the parsed file
	if itemsAttr := file.Body().GetAttribute("items"); itemsAttr != nil {
		body.SetAttributeRaw("items", itemsAttr.Expr().BuildTokens(nil))
	}
}

// checkAndWarnProblematicPatterns checks for known problematic patterns
func checkAndWarnProblematicPatterns(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()
	var warnings []string
	
	// Check for_each with toset()
	if forEachAttr := body.GetAttribute("for_each"); forEachAttr != nil {
		forEachStr := string(forEachAttr.Expr().BuildTokens(nil).Bytes())
		if strings.Contains(forEachStr, "toset(") {
			warnings = append(warnings, 
				"toset() in for_each makes keys and values identical. Consider using a map for distinct keys and values.")
		}
	}
	
	// Check for complex conditionals in item blocks
	for _, b := range body.Blocks() {
		if b.Type() == "dynamic" && len(b.Labels()) > 0 && b.Labels()[0] == "item" {
			dynBody := b.Body()
			if forEachAttr := dynBody.GetAttribute("for_each"); forEachAttr != nil {
				forEachStr := string(forEachAttr.Expr().BuildTokens(nil).Bytes())
				// Check for complex ternary operators
				if strings.Count(forEachStr, "?") > 1 || 
				   (strings.Contains(forEachStr, "?") && strings.Contains(forEachStr, "for ")) {
					warnings = append(warnings,
						"Complex conditional logic in dynamic block may require manual review after migration.")
				}
			}
		}
	}
	
	// Add warnings as comments
	for _, warning := range warnings {
		comment := fmt.Sprintf("\n  # MIGRATION WARNING: %s", warning)
		tokens := hclwrite.Tokens{
			&hclwrite.Token{
				Type:  hclsyntax.TokenComment,
				Bytes: []byte(comment),
			},
		}
		body.AppendUnstructuredTokens(tokens)
	}
}

// addDiagnosticsAsComments adds diagnostic warnings as comments to the body
func addDiagnosticsAsComments(body *hclwrite.Body, diags ast.Diagnostics) {
	// Add comment about complicated HCL expressions if any
	if len(diags.ComplicatedHCL) > 0 {
		comment := fmt.Sprintf("\n  # MIGRATION WARNING: Some expressions could not be automatically migrated. Manual review required.")
		tokens := hclwrite.Tokens{
			&hclwrite.Token{
				Type:  hclsyntax.TokenComment,
				Bytes: []byte(comment),
			},
		}
		body.AppendUnstructuredTokens(tokens)
	}
}

// stripIteratorValueSuffix intelligently handles iterator references when converting 
// from dynamic blocks to for expressions.
// 
// In dynamic blocks, the iterator has .key and .value properties.
// In for expressions:
// - If iterating over simple values: item.value -> item
// - If iterating over objects with properties: item.value.prop -> item.prop  
// - The .key accessor is preserved as-is
func stripIteratorValueSuffix(expr hclsyntax.Expression, iteratorName string) hclsyntax.Expression {
	if iteratorName == "" {
		return expr
	}
	
	// Handle template expressions (string interpolations)
	if template, ok := expr.(*hclsyntax.TemplateExpr); ok {
		newParts := make([]hclsyntax.Expression, len(template.Parts))
		changed := false
		for i, part := range template.Parts {
			newPart := stripIteratorValueSuffix(part, iteratorName)
			if newPart != part {
				changed = true
			}
			newParts[i] = newPart
		}
		if changed {
			return &hclsyntax.TemplateExpr{
				Parts: newParts,
			}
		}
		return expr
	}
	
	// Handle wrapped template expressions
	if wrapped, ok := expr.(*hclsyntax.TemplateWrapExpr); ok {
		inner := stripIteratorValueSuffix(wrapped.Wrapped, iteratorName)
		if inner != wrapped.Wrapped {
			return &hclsyntax.TemplateWrapExpr{
				Wrapped: inner,
			}
		}
		return expr
	}
	
	// Handle conditional expressions (ternary)
	if cond, ok := expr.(*hclsyntax.ConditionalExpr); ok {
		newCond := stripIteratorValueSuffix(cond.Condition, iteratorName)
		newTrue := stripIteratorValueSuffix(cond.TrueResult, iteratorName)
		newFalse := stripIteratorValueSuffix(cond.FalseResult, iteratorName)
		
		// Special handling for "enabled"/"disabled" string literals in conditionals
		// Convert them to boolean values for the redirect fields
		if litTrue, ok := newTrue.(*hclsyntax.TemplateExpr); ok && len(litTrue.Parts) == 1 {
			if litVal, ok := litTrue.Parts[0].(*hclsyntax.LiteralValueExpr); ok {
				if litVal.Val.Type() == cty.String {
					strVal := litVal.Val.AsString()
					if strVal == "enabled" {
						newTrue = &hclsyntax.LiteralValueExpr{Val: cty.BoolVal(true)}
					} else if strVal == "disabled" {
						newTrue = &hclsyntax.LiteralValueExpr{Val: cty.BoolVal(false)}
					}
				}
			}
		}
		
		if litFalse, ok := newFalse.(*hclsyntax.TemplateExpr); ok && len(litFalse.Parts) == 1 {
			if litVal, ok := litFalse.Parts[0].(*hclsyntax.LiteralValueExpr); ok {
				if litVal.Val.Type() == cty.String {
					strVal := litVal.Val.AsString()
					if strVal == "enabled" {
						newFalse = &hclsyntax.LiteralValueExpr{Val: cty.BoolVal(true)}
					} else if strVal == "disabled" {
						newFalse = &hclsyntax.LiteralValueExpr{Val: cty.BoolVal(false)}
					}
				}
			}
		}
		
		if newCond != cond.Condition || newTrue != cond.TrueResult || newFalse != cond.FalseResult {
			return &hclsyntax.ConditionalExpr{
				Condition:   newCond,
				TrueResult:  newTrue,
				FalseResult: newFalse,
			}
		}
		return expr
	}
	
	// Check if this is a traversal expression
	if traversal, ok := expr.(*hclsyntax.ScopeTraversalExpr); ok {
		if len(traversal.Traversal) >= 2 {
			// Check if it starts with iteratorName
			if root, ok := traversal.Traversal[0].(hcl.TraverseRoot); ok && root.Name == iteratorName {
				if attr, ok := traversal.Traversal[1].(hcl.TraverseAttr); ok {
					// Handle iterator.key - preserve as-is (not supported for lists)
					if attr.Name == "key" {
						// For lists, we'd need to handle this differently (index)
						// For now, preserve as-is
						return expr
					}
					
					// Handle iterator.value
					if attr.Name == "value" {
						// Check if there are more parts after .value
						if len(traversal.Traversal) > 2 {
							// iterator.value.something -> iterator.something
							// This handles cases where for_each iterates over objects
							newTraversal := make(hcl.Traversal, 0, len(traversal.Traversal)-1)
							newTraversal = append(newTraversal, hcl.TraverseRoot{Name: iteratorName})
							newTraversal = append(newTraversal, traversal.Traversal[2:]...)
							
							return &hclsyntax.ScopeTraversalExpr{
								Traversal: newTraversal,
							}
						} else {
							// Just iterator.value -> iterator
							// This handles cases where for_each iterates over simple values
							return &hclsyntax.ScopeTraversalExpr{
								Traversal: hcl.Traversal{
									hcl.TraverseRoot{Name: iteratorName},
								},
							}
						}
					}
				}
			}
		}
	}
	
	// Return unchanged if not an iterator pattern we recognize
	return expr
}