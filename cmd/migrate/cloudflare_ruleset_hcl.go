package main

import (
	"sort"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// transformCloudflareRulesetBlock handles HCL transformations for cloudflare_ruleset resources
func transformCloudflareRulesetBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle dynamic "rules" blocks by converting them to for expressions
	transformDynamicRulesBlocks(block, diags)

	// Fix Grit's incorrect list conversions for nested objects
	fixGritNestedObjectConversions(block, diags)

	// Transform headers from list to map format (v4 to v5 migration)
	transformHeadersListToMap(block, diags)

	// Transform cache_key query_string include from list to object (v4 to v5 migration)
	transformCacheKeyQueryString(block, diags)

	// Remove deprecated disable_railgun attribute (removed in v5)
	removeDisableRailgun(block, diags)
	
	// Fix quoted resource references in action_parameters.id
	fixQuotedResourceReferences(block, diags)
	
	// Fix overrides.rules from single object to list of objects
	fixOverridesRulesObjectToList(block, diags)
	
	// Fix overrides.categories from single object to list of objects
	fixOverridesCategoriesObjectToList(block, diags)
	
	// Fix status_code_ttl from single object to list of objects  
	fixStatusCodeTtlObjectToList(block, diags)
	
	// Fix headers from flat object to map format
	fixHeadersObjectToMap(block, diags)
}

// transformDynamicRulesBlocks converts dynamic "rules" blocks to for expressions
// V4: dynamic "rules" { for_each = local.rule_configs; content { ... } }
// V5: rules = [for rule in local.rule_configs : { ... }]
func transformDynamicRulesBlocks(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()
	dynamicRulesBlocks := []*hclwrite.Block{}

	// Find all dynamic blocks with label "rules"
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "dynamic" && len(childBlock.Labels()) > 0 && childBlock.Labels()[0] == "rules" {
			dynamicRulesBlocks = append(dynamicRulesBlocks, childBlock)
		}
	}

	// Process each dynamic rules block
	for _, dynBlock := range dynamicRulesBlocks {
		// Extract the for_each expression
		forEachAttr := dynBlock.Body().GetAttribute("for_each")
		if forEachAttr == nil {
			continue
		}

		// Get the iterator name (defaults to the block label if not specified)
		iteratorName := "rules"
		if iteratorAttr := dynBlock.Body().GetAttribute("iterator"); iteratorAttr != nil {
			// Extract iterator name from the expression
			tokens := iteratorAttr.Expr().BuildTokens(nil)
			if len(tokens) > 0 {
				iteratorName = string(tokens[0].Bytes)
				// Remove quotes if present
				if len(iteratorName) >= 2 && iteratorName[0] == '"' && iteratorName[len(iteratorName)-1] == '"' {
					iteratorName = iteratorName[1 : len(iteratorName)-1]
				}
			}
		}

		// Extract content block
		var contentBlock *hclwrite.Block
		for _, cb := range dynBlock.Body().Blocks() {
			if cb.Type() == "content" {
				contentBlock = cb
				break
			}
		}

		if contentBlock == nil {
			continue
		}

		// Build the for expression as raw tokens
		// Use the same iterator name for consistency
		// We want: [for <iterator> in <for_each_expr> : { ... }]
		tokens := hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("for")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(iteratorName)},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("in")},
		}

		// Add the for_each expression tokens
		tokens = append(tokens, forEachAttr.Expr().BuildTokens(nil)...)

		// Add the colon
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenColon, Bytes: []byte(" : ")})

		// Start object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

		// Process attributes from the content block in sorted order for deterministic output
		attrs := contentBlock.Body().Attributes()
		attrNames := make([]string, 0, len(attrs))
		for name := range attrs {
			attrNames = append(attrNames, name)
		}
		sort.Strings(attrNames)

		for _, name := range attrNames {
			attr := attrs[name]
			// Add attribute name with proper indentation
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Get the attribute expression tokens
			attrTokens := attr.Expr().BuildTokens(nil)

			// Replace iterator references - remove .value suffix
			processedTokens := replaceRulesIteratorReferences(attrTokens, iteratorName)
			tokens = append(tokens, processedTokens...)

			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Handle nested blocks within content (e.g., action_parameters, ratelimit, etc.)
		// These need to be converted to attribute format as well
		for _, nestedBlock := range contentBlock.Body().Blocks() {
			blockType := nestedBlock.Type()
			
			// Add the block as an attribute with object value
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + blockType)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

			// Process nested block attributes
			nestedAttrs := nestedBlock.Body().Attributes()
			nestedAttrNames := make([]string, 0, len(nestedAttrs))
			for name := range nestedAttrs {
				nestedAttrNames = append(nestedAttrNames, name)
			}
			sort.Strings(nestedAttrNames)

			for _, name := range nestedAttrNames {
				attr := nestedAttrs[name]
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      " + name)})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
				
				attrTokens := attr.Expr().BuildTokens(nil)
				processedTokens := replaceRulesIteratorReferences(attrTokens, iteratorName)
				tokens = append(tokens, processedTokens...)
				
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			}

			// Close the nested object
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    }")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Close object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  }")})

		// Close array
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})

		// Remove the dynamic block
		body.RemoveBlock(dynBlock)

		// Add the new attribute with the for expression
		body.SetAttributeRaw("rules", tokens)
	}
}

// replaceRulesIteratorReferences replaces iterator.value references with just iterator
// e.g., "rules.value.expression" becomes "rules.expression"
func replaceRulesIteratorReferences(tokens hclwrite.Tokens, iteratorName string) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		if token.Type == hclsyntax.TokenIdent {
			tokenStr := string(token.Bytes)
			
			// Check if this is the iterator name
			if tokenStr == iteratorName {
				// Look ahead for ".value"
				if i+3 < len(tokens) &&
					tokens[i+1].Type == hclsyntax.TokenDot &&
					string(tokens[i+2].Bytes) == "value" &&
					tokens[i+3].Type == hclsyntax.TokenDot {
					// Found "iterator.value." pattern
					// Keep the iterator name
					result = append(result, token)
					// Skip ".value" but keep the next dot
					result = append(result, tokens[i+3])
					i += 4 // Skip past iterator, dot, value, dot
					continue
				} else if i+2 < len(tokens) &&
					tokens[i+1].Type == hclsyntax.TokenDot &&
					string(tokens[i+2].Bytes) == "value" {
					// Found "iterator.value" at end of expression
					// Just keep the iterator name
					result = append(result, token)
					i += 3 // Skip past iterator, dot, value
					continue
				}
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// fixGritNestedObjectConversions fixes Grit's incorrect conversion of nested objects to lists
// Grit converts: cache_key { custom_key { ... } } to cache_key = [{ custom_key = [{ ... }] }]
// Should be: cache_key = { custom_key = { ... } }
func fixGritNestedObjectConversions(block *hclwrite.Block, diags ast.Diagnostics) {
	rulesAttr := block.Body().GetAttribute("rules")
	if rulesAttr == nil {
		return
	}
	
	// Get the rules attribute tokens and fix nested objects
	tokens := rulesAttr.Expr().BuildTokens(nil)
	fixedTokens := fixNestedListToObject(tokens)
	
	// Update the rules attribute
	block.Body().SetAttributeRaw("rules", fixedTokens)
}

// fixNestedListToObject recursively fixes incorrectly converted nested objects
func fixNestedListToObject(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	objectFields := []string{"cache_key", "custom_key", "query_string", "user", "edge_ttl", "serve_stale", "browser_ttl", "from_value", "target_url", "uri", "path", "query"}
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Check if this is an identifier that should be an object
		if token.Type == hclsyntax.TokenIdent {
			tokenStr := string(token.Bytes)
			isObjectField := false
			for _, field := range objectFields {
				if tokenStr == field {
					isObjectField = true
					break
				}
			}
			
			if isObjectField {
				// Look ahead for = [{
				if i+3 < len(tokens) &&
					tokens[i+1].Type == hclsyntax.TokenEqual &&
					tokens[i+2].Type == hclsyntax.TokenOBrack &&
					tokens[i+3].Type == hclsyntax.TokenOBrace {
					// Found field = [{ pattern - convert to field = {
					result = append(result, token)          // field name
					result = append(result, tokens[i+1])    // =
					// Skip the [ and add {
					result = append(result, tokens[i+3])    // {
					i += 4
					
					// Process the content inside, looking for the matching }]
					depth := 1
					for i < len(tokens) && depth > 0 {
						if tokens[i].Type == hclsyntax.TokenOBrace {
							depth++
						} else if tokens[i].Type == hclsyntax.TokenCBrace {
							depth--
							if depth == 0 {
								// This is the closing brace
								result = append(result, tokens[i]) // }
								i++
								
								// Skip whitespace and the closing ]
								for i < len(tokens) && (tokens[i].Type == hclsyntax.TokenNewline || 
									tokens[i].Type == hclsyntax.TokenComment ||
									(tokens[i].Type == hclsyntax.TokenIdent && len(tokens[i].Bytes) == 0)) {
									i++
								}
								if i < len(tokens) && tokens[i].Type == hclsyntax.TokenCBrack {
									i++ // Skip the ]
								}
								break
							}
						}
						// For any other token, recursively check if it needs fixing
						if tokens[i].Type == hclsyntax.TokenIdent {
							// Check if this might be another field that needs fixing
							needsRecursion := false
							for _, field := range objectFields {
								if string(tokens[i].Bytes) == field {
									needsRecursion = true
									break
								}
							}
							if needsRecursion && i+3 < len(tokens) &&
								tokens[i+1].Type == hclsyntax.TokenEqual &&
								tokens[i+2].Type == hclsyntax.TokenOBrack &&
								tokens[i+3].Type == hclsyntax.TokenOBrace {
								// Recursively handle this nested field
								subTokens := hclwrite.Tokens{tokens[i]}
								j := i + 1
								subDepth := 0
								for j < len(tokens) {
									subTokens = append(subTokens, tokens[j])
									if tokens[j].Type == hclsyntax.TokenOBrace {
										subDepth++
									} else if tokens[j].Type == hclsyntax.TokenCBrace {
										subDepth--
										if subDepth == 0 && j+1 < len(tokens) && tokens[j+1].Type == hclsyntax.TokenCBrack {
											subTokens = append(subTokens, tokens[j+1])
											j++
											break
										}
									}
									j++
								}
								fixed := fixNestedListToObject(subTokens)
								result = append(result, fixed...)
								i = j + 1
								continue
							}
						}
						result = append(result, tokens[i])
						i++
					}
					continue
				}
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// transformHeadersListToMap transforms headers from v4 list format to v5 map format
// V4: headers { name = "X-Header" operation = "set" value = "value" }
// Grit: headers = [{ name = "X-Header" operation = "set" value = "value" }]
// V5: headers = { "X-Header" = { operation = "set", value = "value" } }
func transformHeadersListToMap(block *hclwrite.Block, diags ast.Diagnostics) {
	// Check if we have a rules attribute (list format in v5)
	rulesAttr := block.Body().GetAttribute("rules")
	if rulesAttr == nil {
		// Try processing v4 block format
		for _, rulesBlock := range block.Body().Blocks() {
			if rulesBlock.Type() == "rules" {
				// Look for action_parameters block
				for _, apBlock := range rulesBlock.Body().Blocks() {
					if apBlock.Type() == "action_parameters" {
						transformHeadersInActionParameters(apBlock, diags)
					}
				}
			}
		}
		return
	}

	// Process the rules list and transform headers from list to map
	tokens := rulesAttr.Expr().BuildTokens(nil)
	fixedTokens := transformHeadersInTokens(tokens)
	
	// Update the rules attribute
	block.Body().SetAttributeRaw("rules", fixedTokens)
}

// transformHeadersInTokens transforms headers from list to map in a token stream
// Transforms: headers = [{ name = "X-Header", operation = "set", value = "value" }]
// To: headers = { "X-Header" = { operation = "set", value = "value" } }
func transformHeadersInTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for "headers" = [
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "headers" {
			if i+2 < len(tokens) &&
				tokens[i+1].Type == hclsyntax.TokenEqual &&
				tokens[i+2].Type == hclsyntax.TokenOBrack {
				
				// Found headers = [ pattern
				result = append(result, token)       // headers
				result = append(result, tokens[i+1]) // =
				
				// Convert list of objects to map
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
				
				// Find all header objects in the list
				j := i + 3
				for j < len(tokens) {
					// Skip whitespace
					for j < len(tokens) && (tokens[j].Type == hclsyntax.TokenNewline || 
						(tokens[j].Type == hclsyntax.TokenIdent && string(tokens[j].Bytes) == "")) {
						j++
					}
					
					if j >= len(tokens) || tokens[j].Type == hclsyntax.TokenCBrack {
						break // End of list
					}
					
					// We should have a { for the header object
					if tokens[j].Type == hclsyntax.TokenOBrace {
						// Parse the header object to extract name and other attributes
						headerName := ""
						otherAttrs := hclwrite.Tokens{}
						
						k := j + 1
						depth := 1
						for k < len(tokens) && depth > 0 {
							if tokens[k].Type == hclsyntax.TokenOBrace {
								depth++
							} else if tokens[k].Type == hclsyntax.TokenCBrace {
								depth--
								if depth == 0 {
									break
								}
							} else if tokens[k].Type == hclsyntax.TokenIdent && string(tokens[k].Bytes) == "name" {
								// Found name attribute
								if k+2 < len(tokens) && tokens[k+1].Type == hclsyntax.TokenEqual {
									// Skip "name ="
									k += 2
									// Collect the value
									nameTokens := hclwrite.Tokens{}
									for k < len(tokens) && tokens[k].Type != hclsyntax.TokenComma && 
										tokens[k].Type != hclsyntax.TokenNewline &&
										tokens[k].Type != hclsyntax.TokenCBrace {
										nameTokens = append(nameTokens, tokens[k])
										k++
									}
									// Build the header name
									for _, t := range nameTokens {
										headerName += string(t.Bytes)
									}
									continue
								}
							} else if tokens[k].Type == hclsyntax.TokenIdent && 
								(string(tokens[k].Bytes) == "operation" || string(tokens[k].Bytes) == "value" || string(tokens[k].Bytes) == "expression") {
								// Collect other attributes
								attrStart := k
								// Skip to end of attribute
								for k < len(tokens) && tokens[k].Type != hclsyntax.TokenComma && 
									tokens[k].Type != hclsyntax.TokenNewline &&
									tokens[k].Type != hclsyntax.TokenCBrace {
									k++
								}
								otherAttrs = append(otherAttrs, tokens[attrStart:k]...)
								if k < len(tokens) && tokens[k].Type == hclsyntax.TokenComma {
									otherAttrs = append(otherAttrs, tokens[k])
								}
								otherAttrs = append(otherAttrs, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
								continue
							}
							k++
						}
						
						// Add the header as a map entry
						if headerName != "" {
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        ")})
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte(headerName)})
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
							
							// Add other attributes with proper indentation
							for _, t := range otherAttrs {
								if t.Type == hclsyntax.TokenIdent && 
									(string(t.Bytes) == "operation" || string(t.Bytes) == "value" || string(t.Bytes) == "expression") {
									result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("          ")})
								}
								result = append(result, t)
							}
							
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        }")})
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
						}
						
						// Move past this header object
						j = k + 1
						// Skip comma if present
						if j < len(tokens) && tokens[j].Type == hclsyntax.TokenComma {
							j++
						}
					} else {
						j++
					}
				}
				
				// Close the map
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      }")})
				
				// Skip past the closing ]
				for j < len(tokens) && tokens[j].Type != hclsyntax.TokenCBrack {
					j++
				}
				if j < len(tokens) {
					j++ // Skip ]
				}
				
				i = j
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// transformHeadersInActionParameters converts headers blocks to map format
func transformHeadersInActionParameters(apBlock *hclwrite.Block, diags ast.Diagnostics) {
	// Collect all headers blocks
	headersBlocks := []*hclwrite.Block{}
	for _, block := range apBlock.Body().Blocks() {
		if block.Type() == "headers" {
			headersBlocks = append(headersBlocks, block)
		}
	}
	
	if len(headersBlocks) == 0 {
		return
	}
	
	// Build the headers map
	// headers = { "header-name" = { operation = "set", value = "value" } }
	tokens := hclwrite.Tokens{
		&hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
		&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
	}
	
	for i, headerBlock := range headersBlocks {
		// Get the name attribute
		nameAttr := headerBlock.Body().GetAttribute("name")
		if nameAttr == nil {
			continue
		}
		
		// Add the header name as key
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
		tokens = append(tokens, nameAttr.Expr().BuildTokens(nil)...)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		
		// Add other attributes (operation, value, expression)
		attrs := headerBlock.Body().Attributes()
		attrNames := make([]string, 0, len(attrs))
		for name := range attrs {
			if name != "name" { // Skip name as it's now the key
				attrNames = append(attrNames, name)
			}
		}
		sort.Strings(attrNames)
		
		for j, attrName := range attrNames {
			attr := attrs[attrName]
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      " + attrName)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			tokens = append(tokens, attr.Expr().BuildTokens(nil)...)
			
			if j < len(attrNames)-1 {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    }")})
		if i < len(headersBlocks)-1 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		
		// Remove the headers block
		apBlock.Body().RemoveBlock(headerBlock)
	}
	
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  }")})
	
	// Set the headers attribute with the map
	apBlock.Body().SetAttributeRaw("headers", tokens)
}

// transformCacheKeyQueryString transforms cache_key query_string include from list to object
// V4: query_string { include = ["param1", "param2"] }
// V5: query_string { include = { "param1" = true, "param2" = true } }
// Also handles: ["*"] -> "*"
func transformCacheKeyQueryString(block *hclwrite.Block, diags ast.Diagnostics) {
	// Get the rules attribute to process the list format
	rulesAttr := block.Body().GetAttribute("rules")
	if rulesAttr == nil {
		// Try processing v4 block format
		for _, rulesBlock := range block.Body().Blocks() {
			if rulesBlock.Type() == "rules" {
				// Look for action_parameters block
				for _, apBlock := range rulesBlock.Body().Blocks() {
					if apBlock.Type() == "action_parameters" {
						// Look for cache_key block
						for _, ckBlock := range apBlock.Body().Blocks() {
							if ckBlock.Type() == "cache_key" {
								// Look for custom_key block
								for _, customKeyBlock := range ckBlock.Body().Blocks() {
									if customKeyBlock.Type() == "custom_key" {
										// Look for query_string block
										for _, qsBlock := range customKeyBlock.Body().Blocks() {
											if qsBlock.Type() == "query_string" {
												transformQueryStringInclude(qsBlock, diags)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		return
	}
	
	// Process the rules list and transform query_string include
	tokens := rulesAttr.Expr().BuildTokens(nil)
	fixedTokens := transformQueryStringIncludeInTokens(tokens)
	
	// Update the rules attribute
	block.Body().SetAttributeRaw("rules", fixedTokens)
}

// transformQueryStringIncludeInTokens transforms query_string include in a token stream
func transformQueryStringIncludeInTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for query_string context, then find include within it
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "query_string" {
			// Add query_string token
			result = append(result, token)
			i++
			
			// Look for = {
			if i+1 < len(tokens) && tokens[i].Type == hclsyntax.TokenEqual && tokens[i+1].Type == hclsyntax.TokenOBrace {
				result = append(result, tokens[i])   // =
				result = append(result, tokens[i+1]) // {
				i += 2
				
				// Now look for include = [ within this query_string block
				depth := 1
				for i < len(tokens) && depth > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						depth++
						result = append(result, tokens[i])
						i++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						depth--
						if depth == 0 {
							result = append(result, tokens[i])
							i++
							break
						}
						result = append(result, tokens[i])
						i++
					} else if tokens[i].Type == hclsyntax.TokenIdent && string(tokens[i].Bytes) == "include" && depth == 1 {
						// Found include within query_string
						if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual && tokens[i+2].Type == hclsyntax.TokenOBrack {
							// Transform include = [...]
							result = append(result, tokens[i])   // include
							result = append(result, tokens[i+1]) // =
							i += 3 // Skip include, =, [
							
							// Collect and transform the list
							listItems := []string{}
							bracketDepth := 1
							
							for i < len(tokens) && bracketDepth > 0 {
								if tokens[i].Type == hclsyntax.TokenOBrack {
									bracketDepth++
									i++
								} else if tokens[i].Type == hclsyntax.TokenCBrack {
									bracketDepth--
									if bracketDepth == 0 {
										i++
										break
									}
									i++
								} else if tokens[i].Type == hclsyntax.TokenOQuote {
									// Start of string - collect the full string token sequence
									stringTokens := []*hclwrite.Token{tokens[i]} // Opening quote
									i++
									// Collect until closing quote
									for i < len(tokens) && tokens[i].Type != hclsyntax.TokenCQuote {
										stringTokens = append(stringTokens, tokens[i])
										i++
									}
									if i < len(tokens) && tokens[i].Type == hclsyntax.TokenCQuote {
										stringTokens = append(stringTokens, tokens[i]) // Closing quote
										i++
									}
									// Build the string from tokens
									var paramName string
									for _, t := range stringTokens {
										paramName += string(t.Bytes)
									}
									listItems = append(listItems, paramName)
								} else {
									i++
								}
							}
							
							// Output the transformed include
							// V5 structure: include = { list = [...] } or include = { all = true }
							if len(listItems) == 1 && listItems[0] == "\"*\"" {
								// Special case: ["*"] becomes { all = true }
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("              ")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("all")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("true")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("            ")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
							} else if len(listItems) > 0 {
								// Convert to object with list attribute
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("              ")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("list")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")})
								
								for idx, item := range listItems {
									// Add the string (already has quotes from collection)
									if len(item) >= 2 && item[0] == '"' && item[len(item)-1] == '"' {
										// Output the quoted string
										result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte("\"")})
										result = append(result, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(item[1:len(item)-1])})
										result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte("\"")})
									} else {
										// Shouldn't happen, but handle as identifier
										result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(item)})
									}
									if idx < len(listItems)-1 {
										result = append(result, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(", ")})
									}
								}
								
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("            ")})
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
							}
						} else {
							// Not the pattern we're looking for
							result = append(result, tokens[i])
							i++
						}
					} else {
						result = append(result, tokens[i])
						i++
					}
				}
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// removeDisableRailgun removes the deprecated disable_railgun attribute from action_parameters
func removeDisableRailgun(block *hclwrite.Block, diags ast.Diagnostics) {
	// Process rules attribute if it exists (v5 list format)
	if rulesAttr := block.Body().GetAttribute("rules"); rulesAttr != nil {
		// For list format, we need to process tokens to remove disable_railgun
		tokens := rulesAttr.Expr().BuildTokens(nil)
		fixedTokens := removeDisableRailgunFromTokens(tokens)
		block.Body().SetAttributeRaw("rules", fixedTokens)
		return
	}

	// Process rules blocks (v4 block format)
	for _, rulesBlock := range block.Body().Blocks() {
		if rulesBlock.Type() == "rules" {
			// Look for action_parameters block
			for _, apBlock := range rulesBlock.Body().Blocks() {
				if apBlock.Type() == "action_parameters" {
					// Remove disable_railgun attribute if it exists
					if apBlock.Body().GetAttribute("disable_railgun") != nil {
						apBlock.Body().RemoveAttribute("disable_railgun")
					}
				}
			}
		}
	}
}

// removeDisableRailgunFromTokens removes disable_railgun from a token stream
func removeDisableRailgunFromTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		// Look for disable_railgun attribute
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "disable_railgun" {
			// Skip disable_railgun and its value
			// Look for pattern: disable_railgun = value
			if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual {
				// Skip disable_railgun, =, and the value
				i += 2 // Skip disable_railgun and =

				// Skip the value (could be true, false, or an expression)
				if i < len(tokens) {
					if tokens[i].Type == hclsyntax.TokenIdent ||
					   tokens[i].Type == hclsyntax.TokenOQuote {
						// Simple value or string, skip it
						if tokens[i].Type == hclsyntax.TokenOQuote {
							// Skip string tokens until closing quote
							i++
							for i < len(tokens) && tokens[i].Type != hclsyntax.TokenCQuote {
								i++
							}
							if i < len(tokens) && tokens[i].Type == hclsyntax.TokenCQuote {
								i++
							}
						} else {
							i++ // Skip simple identifier
						}
					}
				}

				// Skip any trailing comma or newline
				for i < len(tokens) && (tokens[i].Type == hclsyntax.TokenComma ||
				                        tokens[i].Type == hclsyntax.TokenNewline) {
					i++
				}
				continue
			}
		}

		result = append(result, token)
		i++
	}

	return result
}

// transformQueryStringInclude converts include from list to object
func transformQueryStringInclude(qsBlock *hclwrite.Block, diags ast.Diagnostics) {
	includeAttr := qsBlock.Body().GetAttribute("include")
	if includeAttr == nil {
		return
	}
	
	// Get the include attribute tokens
	tokens := includeAttr.Expr().BuildTokens(nil)
	
	// Check if it's a list
	if len(tokens) > 0 && tokens[0].Type == hclsyntax.TokenOBrack {
		// Convert list to object
		// ["param1", "param2"] -> { "param1" = true, "param2" = true }
		// ["*"] -> "*"
		
		// Special case: ["*"] becomes just "*"
		if len(tokens) == 3 && tokens[1].Type == hclsyntax.TokenOQuote {
			tokenStr := string(tokens[1].Bytes)
			if tokenStr == "\"*\"" {
				// Just set it to "*"
				qsBlock.Body().SetAttributeRaw("include", hclwrite.Tokens{
					&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte("\"*\"")},
				})
				return
			}
		}
		
		// Build object from list elements
		newTokens := hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
			&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		}
		
		// Parse the list elements
		inString := false
		currentString := ""
		params := []string{}
		
		for _, token := range tokens[1:] { // Skip opening bracket
			if token.Type == hclsyntax.TokenCBrack {
				break // End of list
			}
			
			if token.Type == hclsyntax.TokenOQuote {
				if !inString {
					inString = true
					currentString = "\""
				} else {
					currentString += "\""
					params = append(params, currentString)
					currentString = ""
					inString = false
				}
			} else if inString {
				currentString += string(token.Bytes)
			}
		}
		
		// Add each param as a key with value true
		for i, param := range params {
			newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
			newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte(param)})
			newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("true")})
			
			if i < len(params)-1 {
				newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			}
			newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		
		newTokens = append(newTokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    }")})
		
		// Update the attribute
		qsBlock.Body().SetAttributeRaw("include", newTokens)
	}
}

// fixQuotedResourceReferences removes incorrect quotes from resource references in action_parameters.id
func fixQuotedResourceReferences(block *hclwrite.Block, diags ast.Diagnostics) {
	// Process rules attribute if it exists (v5 list format)
	if rulesAttr := block.Body().GetAttribute("rules"); rulesAttr != nil {
		tokens := rulesAttr.Expr().BuildTokens(nil)
		fixedTokens := unquoteResourceReferencesInTokens(tokens)
		block.Body().SetAttributeRaw("rules", fixedTokens)
	}
}

// unquoteResourceReferencesInTokens removes quotes from resource references in action_parameters.id
func unquoteResourceReferencesInTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for action_parameters context
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "action_parameters" {
			result = append(result, token)
			i++
			
			// Look for = {
			if i+1 < len(tokens) && tokens[i].Type == hclsyntax.TokenEqual && tokens[i+1].Type == hclsyntax.TokenOBrace {
				result = append(result, tokens[i])   // =
				result = append(result, tokens[i+1]) // {
				i += 2
				
				// Process inside action_parameters block
				depth := 1
				for i < len(tokens) && depth > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						depth++
						result = append(result, tokens[i])
						i++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						depth--
						result = append(result, tokens[i])
						i++
						if depth == 0 {
							break
						}
					} else if tokens[i].Type == hclsyntax.TokenIdent && string(tokens[i].Bytes) == "id" && depth == 1 {
						// Found id attribute in action_parameters
						result = append(result, tokens[i]) // id
						i++
						
						// Look for = "cloudflare_ruleset..."
						if i+2 < len(tokens) && tokens[i].Type == hclsyntax.TokenEqual && tokens[i+1].Type == hclsyntax.TokenOQuote {
							result = append(result, tokens[i]) // =
							i += 2 // Skip = and opening quote
							
							// Check if this looks like a resource reference
							if i < len(tokens) && tokens[i].Type == hclsyntax.TokenQuotedLit {
								refStr := string(tokens[i].Bytes)
								if strings.HasPrefix(refStr, "cloudflare_") && strings.Contains(refStr, ".") && strings.HasSuffix(refStr, ".id") {
									// This is a quoted resource reference - unquote it
									result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(refStr)})
									i++
									// Skip closing quote
									if i < len(tokens) && tokens[i].Type == hclsyntax.TokenCQuote {
										i++
									}
								} else {
									// Keep as quoted string
									result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte("\"")})
									result = append(result, tokens[i])
									i++
									if i < len(tokens) && tokens[i].Type == hclsyntax.TokenCQuote {
										result = append(result, tokens[i])
										i++
									}
								}
							} else {
								// Not what we expected, keep original
								result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte("\"")})
								result = append(result, tokens[i])
								i++
							}
						} else {
							// Not the pattern we're looking for
							result = append(result, tokens[i])
							i++
						}
					} else {
						result = append(result, tokens[i])
						i++
					}
				}
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}
// fixOverridesRulesObjectToList converts overrides.rules from single object to list of objects
// V4/incorrect: overrides = { rules = { id = "...", action = "..." } }
// V5/correct:   overrides = { rules = [{ id = "...", action = "..." }] }
func fixOverridesRulesObjectToList(block *hclwrite.Block, diags ast.Diagnostics) {
	// Process rules attribute if it exists (v5 list format)
	if rulesAttr := block.Body().GetAttribute("rules"); rulesAttr != nil {
		tokens := rulesAttr.Expr().BuildTokens(nil)
		fixedTokens := convertOverridesRulesToList(tokens)
		block.Body().SetAttributeRaw("rules", fixedTokens)
	}
}

// convertOverridesRulesToList finds overrides.rules objects and converts them to lists
func convertOverridesRulesToList(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for overrides context
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "overrides" {
			result = append(result, token)
			i++
			
			// Look for = {
			if i+1 < len(tokens) && tokens[i].Type == hclsyntax.TokenEqual && tokens[i+1].Type == hclsyntax.TokenOBrace {
				result = append(result, tokens[i])   // =
				result = append(result, tokens[i+1]) // {
				i += 2
				
				// Process inside overrides block
				depth := 1
				for i < len(tokens) && depth > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						depth++
						result = append(result, tokens[i])
						i++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						depth--
						result = append(result, tokens[i])
						i++
						if depth == 0 {
							break
						}
					} else if tokens[i].Type == hclsyntax.TokenIdent && string(tokens[i].Bytes) == "rules" && depth == 1 {
						// Found rules attribute in overrides
						result = append(result, tokens[i]) // rules
						i++
						
						// Look for = {
						if i+1 < len(tokens) && tokens[i].Type == hclsyntax.TokenEqual && tokens[i+1].Type == hclsyntax.TokenOBrace {
							result = append(result, tokens[i]) // =
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")}) // Add [
							result = append(result, tokens[i+1]) // {
							i += 2
							
							// Process the object content and add closing bracket
							objDepth := 1
							for i < len(tokens) && objDepth > 0 {
								if tokens[i].Type == hclsyntax.TokenOBrace {
									objDepth++
									result = append(result, tokens[i])
									i++
								} else if tokens[i].Type == hclsyntax.TokenCBrace {
									objDepth--
									result = append(result, tokens[i])
									if objDepth == 0 {
										// Add closing bracket after the object
										result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})
									}
									i++
								} else {
									result = append(result, tokens[i])
									i++
								}
							}
						} else {
							// Not the pattern we're looking for
							result = append(result, tokens[i])
							i++
						}
					} else {
						result = append(result, tokens[i])
						i++
					}
				}
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// fixOverridesCategoriesObjectToList converts overrides.categories from single object to list of objects
func fixOverridesCategoriesObjectToList(block *hclwrite.Block, diags ast.Diagnostics) {
	if rulesAttr := block.Body().GetAttribute("rules"); rulesAttr != nil {
		tokens := rulesAttr.Expr().BuildTokens(nil)
		fixedTokens := convertOverridesCategoriesToList(tokens)
		block.Body().SetAttributeRaw("rules", fixedTokens)
	}
}

// convertOverridesCategoriesToList finds overrides.categories objects and converts them to lists
func convertOverridesCategoriesToList(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for "categories" within overrides context
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "categories" {
			// Look ahead to see if this is followed by = {
			if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual && tokens[i+2].Type == hclsyntax.TokenOBrace {
				result = append(result, token)       // categories
				result = append(result, tokens[i+1]) // =
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")}) // Add [
				result = append(result, tokens[i+2]) // {
				i += 3
				
				// Process the object content and add closing bracket
				depth := 1
				for i < len(tokens) && depth > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						depth++
						result = append(result, tokens[i])
						i++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						depth--
						result = append(result, tokens[i])
						if depth == 0 {
							// Add closing bracket after the object
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})
						}
						i++
					} else {
						result = append(result, tokens[i])
						i++
					}
				}
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// fixStatusCodeTtlObjectToList converts status_code_ttl from single object to list of objects
func fixStatusCodeTtlObjectToList(block *hclwrite.Block, diags ast.Diagnostics) {
	if rulesAttr := block.Body().GetAttribute("rules"); rulesAttr != nil {
		tokens := rulesAttr.Expr().BuildTokens(nil)
		fixedTokens := convertStatusCodeTtlToList(tokens)
		block.Body().SetAttributeRaw("rules", fixedTokens)
	}
}

// convertStatusCodeTtlToList finds status_code_ttl objects and converts them to lists
func convertStatusCodeTtlToList(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for "status_code_ttl"
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "status_code_ttl" {
			// Look ahead to see if this is followed by = {
			if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual && tokens[i+2].Type == hclsyntax.TokenOBrace {
				result = append(result, token)       // status_code_ttl
				result = append(result, tokens[i+1]) // =
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")}) // Add [
				result = append(result, tokens[i+2]) // {
				i += 3
				
				// Process the object content and add closing bracket
				depth := 1
				for i < len(tokens) && depth > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						depth++
						result = append(result, tokens[i])
						i++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						depth--
						result = append(result, tokens[i])
						if depth == 0 {
							// Add closing bracket after the object
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})
						}
						i++
					} else {
						result = append(result, tokens[i])
						i++
					}
				}
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}

// fixHeadersObjectToMap converts flat headers object to map format
// V4/incorrect: headers = { name = "Header-Name", operation = "set", value = "value" }
// V5/correct:   headers = { "Header-Name" = { operation = "set", value = "value" } }
func fixHeadersObjectToMap(block *hclwrite.Block, diags ast.Diagnostics) {
	if rulesAttr := block.Body().GetAttribute("rules"); rulesAttr != nil {
		tokens := rulesAttr.Expr().BuildTokens(nil)
		fixedTokens := convertHeadersObjectToMap(tokens)
		block.Body().SetAttributeRaw("rules", fixedTokens)
	}
}

// convertHeadersObjectToMap finds headers objects and converts them to proper map format
func convertHeadersObjectToMap(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := make(hclwrite.Tokens, 0, len(tokens))
	
	i := 0
	for i < len(tokens) {
		token := tokens[i]
		
		// Look for "headers = {"
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "headers" {
			if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual && tokens[i+2].Type == hclsyntax.TokenOBrace {
				// Found headers = { pattern
				result = append(result, token)       // headers
				result = append(result, tokens[i+1]) // =
				result = append(result, tokens[i+2]) // {
				i += 3
				
				// Parse the content to extract name and other attributes
				headerName := ""
				otherAttrs := make(map[string]hclwrite.Tokens)
				
				depth := 1
				for i < len(tokens) && depth > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						depth++
						i++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						depth--
						if depth == 0 {
							break
						}
						i++
					} else if tokens[i].Type == hclsyntax.TokenIdent && depth == 1 {
						attrName := string(tokens[i].Bytes)
						if attrName == "name" {
							// Extract the name value
							if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual {
								i += 2 // Skip name =
								nameTokens := hclwrite.Tokens{}
								// Collect tokens until comma, newline, or closing brace
								for i < len(tokens) && 
									tokens[i].Type != hclsyntax.TokenComma &&
									tokens[i].Type != hclsyntax.TokenNewline &&
									tokens[i].Type != hclsyntax.TokenCBrace {
									nameTokens = append(nameTokens, tokens[i])
									i++
								}
								// Convert tokens to string and clean up
								nameStr := string(nameTokens.Bytes())
								nameStr = strings.TrimSpace(nameStr)
								headerName = nameStr
								// Skip comma if present
								if i < len(tokens) && tokens[i].Type == hclsyntax.TokenComma {
									i++
								}
							}
						} else if attrName == "operation" || attrName == "value" || attrName == "expression" {
							// Extract other attributes
							if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenEqual {
								attrTokens := hclwrite.Tokens{tokens[i], tokens[i+1]} // attr =
								i += 2
								// Collect tokens until comma, newline, or closing brace
								for i < len(tokens) && 
									tokens[i].Type != hclsyntax.TokenComma &&
									tokens[i].Type != hclsyntax.TokenNewline &&
									tokens[i].Type != hclsyntax.TokenCBrace {
									attrTokens = append(attrTokens, tokens[i])
									i++
								}
								otherAttrs[attrName] = attrTokens
								// Skip comma if present
								if i < len(tokens) && tokens[i].Type == hclsyntax.TokenComma {
									i++
								}
							}
						} else {
							i++
						}
					} else {
						i++
					}
				}
				
				// Rebuild as map format
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        ")})
				
				// Add header name as key (preserve quotes if present)
				if headerName != "" {
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte(headerName)})
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					
					// Add other attributes
					attrNames := make([]string, 0, len(otherAttrs))
					for name := range otherAttrs {
						attrNames = append(attrNames, name)
					}
					sort.Strings(attrNames)
					
					for j, attrName := range attrNames {
						result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("          ")})
						result = append(result, otherAttrs[attrName]...)
						if j < len(attrNames)-1 {
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
						}
						result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					}
					
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        }")})
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      }")})
				}
				
				i++ // Skip the closing }
				continue
			}
		}
		
		result = append(result, token)
		i++
	}
	
	return result
}
