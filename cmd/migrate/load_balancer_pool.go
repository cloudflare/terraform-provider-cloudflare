package main

import (
	"regexp"
	"sort"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// transformLoadBalancerPoolBlock transforms a load_balancer_pool resource block
func transformLoadBalancerPoolBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle dynamic origins blocks by converting them to for expressions
	transformDynamicOriginsBlocks(block, diags)

	// Handle static origins blocks by transforming nested header blocks
	transformStaticOriginsBlocks(block, diags)
}

// transformStaticOriginsBlocks handles static origins blocks that have been partially migrated
// It fixes the nested header blocks within origins arrays
func transformStaticOriginsBlocks(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()

	// Check for origins blocks (not yet converted to attribute)
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "origins" {
			// Transform the header block within this origins block
			transformHeaderBlockInOrigins(childBlock, diags)
		}
	}

	// Also check if there's an origins attribute that needs transformation
	originsAttr := body.GetAttribute("origins")
	if originsAttr != nil {
		// The origins attribute exists - it might have malformed header syntax
		// that was fixed by fixMalformedHeaderBlocks
		// Now we need to transform header = { header = "Host", values = [...] }
		// to header = { host = [...] }
		transformOriginsAttribute(body, originsAttr, diags)
	}
}

// transformOriginsAttribute transforms the origins attribute to fix header syntax
func transformOriginsAttribute(body *hclwrite.Body, originsAttr *hclwrite.Attribute, diags ast.Diagnostics) {
	// Get the current expression tokens
	tokens := originsAttr.Expr().BuildTokens(nil)

	// Transform header = { header = "Host", values = [...] } to header = { host = [...] }
	newTokens := transformHeaderTokens(tokens)

	// Set the transformed expression back
	body.SetAttributeRaw("origins", newTokens)
}

// transformHeaderTokens transforms header object syntax in tokens
func transformHeaderTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	result := hclwrite.Tokens{}
	i := 0

	for i < len(tokens) {
		// Look for pattern: header = { ... header = "Host" ... values = [...] ... }
		// Need to be more flexible with whitespace
		if isHeaderHostPattern(tokens, i) {

			// Found the pattern, now transform it
			// Copy "header = {"
			result = append(result, tokens[i]) // header
			i++

			// Skip whitespace and copy "="
			for i < len(tokens) && isWhitespaceToken(tokens[i]) {
				result = append(result, tokens[i])
				i++
			}
			result = append(result, tokens[i]) // =
			i++

			// Skip whitespace and copy "{"
			for i < len(tokens) && isWhitespaceToken(tokens[i]) {
				result = append(result, tokens[i])
				i++
			}
			result = append(result, tokens[i]) // {
			i++

			// Add newline if there was one
			if i < len(tokens) && tokens[i].Type == hclsyntax.TokenNewline {
				result = append(result, tokens[i])
				i++
			}

			// Skip whitespace
			for i < len(tokens) && isWhitespaceToken(tokens[i]) {
				result = append(result, tokens[i])
				i++
			}

			// Now we should be at "header" - skip past "header = \"Host\""
			// Skip "header"
			i++
			// Skip whitespace
			for i < len(tokens) && isWhitespaceToken(tokens[i]) {
				i++
			}
			// Skip "="
			i++
			// Skip whitespace
			for i < len(tokens) && isWhitespaceToken(tokens[i]) {
				i++
			}
			// Skip "\"Host\""
			i += 3

			// Add "host = "
			result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("host")})
			result = append(result, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Skip whitespace and newlines to find "values"
			for i < len(tokens) && (tokens[i].Type == hclsyntax.TokenNewline || isWhitespaceToken(tokens[i])) {
				i++
			}

			if i < len(tokens) && tokens[i].Type == hclsyntax.TokenIdent && string(tokens[i].Bytes) == "values" {
				i++ // skip "values"

				// Skip whitespace
				for i < len(tokens) && isWhitespaceToken(tokens[i]) {
					i++
				}

				// Skip the = sign
				if i < len(tokens) && tokens[i].Type == hclsyntax.TokenEqual {
					i++
				}

				// Skip whitespace
				for i < len(tokens) && isWhitespaceToken(tokens[i]) {
					i++
				}

				// Now copy the array value and closing brace
				braceCount := 1
				for i < len(tokens) && braceCount > 0 {
					if tokens[i].Type == hclsyntax.TokenOBrace {
						braceCount++
					} else if tokens[i].Type == hclsyntax.TokenCBrace {
						braceCount--
						if braceCount == 0 {
							// This is the closing brace for header object
							result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" ")})
							result = append(result, tokens[i])
							i++
							break
						}
					}
					result = append(result, tokens[i])
					i++
				}
			}
		} else {
			// Not our pattern, copy as-is
			result = append(result, tokens[i])
			i++
		}
	}

	return result
}

// isHeaderHostPattern checks if we have a header = { header = "Host" ... } pattern starting at position i
func isHeaderHostPattern(tokens hclwrite.Tokens, start int) bool {
	i := start

	// Check for "header"
	if i >= len(tokens) || tokens[i].Type != hclsyntax.TokenIdent || string(tokens[i].Bytes) != "header" {
		return false
	}
	i++

	// Skip whitespace
	for i < len(tokens) && isWhitespaceToken(tokens[i]) {
		i++
	}

	// Check for "="
	if i >= len(tokens) || tokens[i].Type != hclsyntax.TokenEqual {
		return false
	}
	i++

	// Skip whitespace
	for i < len(tokens) && isWhitespaceToken(tokens[i]) {
		i++
	}

	// Check for "{"
	if i >= len(tokens) || tokens[i].Type != hclsyntax.TokenOBrace {
		return false
	}
	i++

	// Skip whitespace and newlines
	for i < len(tokens) && (isWhitespaceToken(tokens[i]) || tokens[i].Type == hclsyntax.TokenNewline) {
		i++
	}

	// Look for "header = \"Host\""
	if i >= len(tokens) || tokens[i].Type != hclsyntax.TokenIdent || string(tokens[i].Bytes) != "header" {
		return false
	}
	i++

	// Skip whitespace
	for i < len(tokens) && isWhitespaceToken(tokens[i]) {
		i++
	}

	// Check for "="
	if i >= len(tokens) || tokens[i].Type != hclsyntax.TokenEqual {
		return false
	}
	i++

	// Skip whitespace
	for i < len(tokens) && isWhitespaceToken(tokens[i]) {
		i++
	}

	// Check for "\"Host\""
	if i+2 >= len(tokens) {
		return false
	}
	if tokens[i].Type != hclsyntax.TokenOQuote {
		return false
	}
	if tokens[i+1].Type != hclsyntax.TokenQuotedLit || string(tokens[i+1].Bytes) != "Host" {
		return false
	}
	if tokens[i+2].Type != hclsyntax.TokenCQuote {
		return false
	}

	// Found the pattern
	return true
}

// isWhitespaceToken checks if a token is whitespace
func isWhitespaceToken(token *hclwrite.Token) bool {
	return token.Type == hclsyntax.TokenIdent && (len(token.Bytes) == 0 || string(token.Bytes) == " " || string(token.Bytes) == "\t")
}

// transformHeaderBlockInOrigins transforms header blocks within an origins block
func transformHeaderBlockInOrigins(originsBlock *hclwrite.Block, diags ast.Diagnostics) {
	body := originsBlock.Body()

	// Find header blocks
	for _, headerBlock := range body.Blocks() {
		if headerBlock.Type() == "header" {
			// Check if this is a "Host" header
			headerAttr := headerBlock.Body().GetAttribute("header")
			valuesAttr := headerBlock.Body().GetAttribute("values")

			if headerAttr != nil && valuesAttr != nil {
				// Remove the header block
				body.RemoveBlock(headerBlock)

				// Add header as an attribute instead
				// We need to build the tokens for the header attribute value
				tokens := hclwrite.Tokens{
					&hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
					&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" host")},
					&hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")},
				}

				// Add the values expression tokens
				tokens = append(tokens, valuesAttr.Expr().BuildTokens(nil)...)

				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" ")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})

				// Set the header attribute
				body.SetAttributeRaw("header", tokens)
			}
		}
	}
}

// transformDynamicOriginsBlocks converts dynamic "origins" blocks to for expressions
// V4: dynamic "origins" { for_each = local.origin_list; content { ... } }
// V5: origins = [for item in local.origin_list : { ... }]
func transformDynamicOriginsBlocks(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()
	dynamicOriginsBlocks := []*hclwrite.Block{}

	// Find all dynamic blocks with label "origins"
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "dynamic" && len(childBlock.Labels()) > 0 && childBlock.Labels()[0] == "origins" {
			dynamicOriginsBlocks = append(dynamicOriginsBlocks, childBlock)
		}
	}

	// Process each dynamic origins block
	for _, dynBlock := range dynamicOriginsBlocks {
		// Extract the for_each expression
		forEachAttr := dynBlock.Body().GetAttribute("for_each")
		if forEachAttr == nil {
			continue
		}

		// Get the iterator name (defaults to the block label if not specified)
		iteratorName := "origins"
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
		// We need to detect if we're iterating over a map or list
		// For now, we'll assume maps need "key, value" and generate that pattern
		// This handles the common case where dynamic blocks iterate over maps
		tokens := hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("for")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" key")},
			&hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" value")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" in")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(" ")},
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

			// Special handling for header attribute
			if name == "header" {
				attrTokens := attr.Expr().BuildTokens(nil)
				tokenStr := string(hclwrite.Tokens(attrTokens).Bytes())

				// Check if this is a Host header structure
				if strings.Contains(tokenStr, `header = "Host"`) && strings.Contains(tokenStr, "values =") {
					// Transform header = { header = "Host", values = [...] } to header = { host = [...] }
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    header")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      host")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

					// Extract the values part
					// Find the position of "values = " in the token string
					valuesStart := strings.Index(tokenStr, "values =")
					if valuesStart != -1 {
						// Find the array part after "values = "
						valuesStr := tokenStr[valuesStart+8:] // Skip "values ="
						valuesStr = strings.TrimSpace(valuesStr)
						// Find the end of the array (should end with ])
						arrayEnd := strings.LastIndex(valuesStr, "]")
						if arrayEnd != -1 {
							arrayStr := valuesStr[:arrayEnd+1]
							// Replace iterator references in the array string
							arrayStr = strings.ReplaceAll(arrayStr, iteratorName+".value", "value")
							arrayStr = strings.ReplaceAll(arrayStr, iteratorName+".key", "key")

							// Create tokens for the array
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(arrayStr)})
						}
					}

					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
				} else {
					// Not a Host header, keep as is
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
					processedTokens := replaceIteratorReferencesCarefully(attrTokens, iteratorName)
					tokens = append(tokens, processedTokens...)
				}
			} else {
				// Regular attribute
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

				attrTokens := attr.Expr().BuildTokens(nil)
				processedTokens := replaceIteratorReferencesCarefully(attrTokens, iteratorName)
				tokens = append(tokens, processedTokens...)
			}

			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Handle nested blocks within content (e.g., header blocks)
		for _, nestedBlock := range contentBlock.Body().Blocks() {
			if nestedBlock.Type() == "header" {
				// Check if this is a Host header
				headerAttr := nestedBlock.Body().GetAttribute("header")
				valuesAttr := nestedBlock.Body().GetAttribute("values")

				if headerAttr != nil && valuesAttr != nil {
					headerTokens := headerAttr.Expr().BuildTokens(nil)
					if isHostHeader(headerTokens) {
						// Transform header { header = "Host", values = [...] } to header = { host = [...] }
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    header")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

						// Add host = values
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      host")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

						valuesTokens := valuesAttr.Expr().BuildTokens(nil)
						processedTokens := replaceIteratorReferencesCarefully(valuesTokens, iteratorName)
						tokens = append(tokens, processedTokens...)

						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					} else {
						// Not a Host header, keep the original structure
						// Convert header block to header attribute
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    header")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

						// Process all header block attributes
						for hName, hAttr := range nestedBlock.Body().Attributes() {
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      " + hName)})
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

							attrTokens := hAttr.Expr().BuildTokens(nil)
							processedTokens := replaceIteratorReferencesCarefully(attrTokens, iteratorName)
							tokens = append(tokens, processedTokens...)
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
						}

						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
					}
				}
			}
		}

		// Close object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})

		// Set the origins attribute with the for expression
		body.SetAttributeRaw("origins", tokens)

		// Remove the dynamic block
		body.RemoveBlock(dynBlock)
	}
}

// replaceIteratorReferencesCarefully is a more careful version for problematic cases like local.origins[origins.value]
func replaceIteratorReferencesCarefully(tokens hclwrite.Tokens, iteratorName string) hclwrite.Tokens {
	// For expressions like local.origins[origins.value].weight where iterator is "origins"
	// We need to replace only origins.value -> value, not local.origins

	result := hclwrite.Tokens{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Pass through template markers unchanged
		if token.Type == hclsyntax.TokenTemplateInterp || token.Type == hclsyntax.TokenTemplateSeqEnd {
			result = append(result, token)
			continue
		}

		// Check if this token is the iterator name
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == iteratorName {
			// Check if this is preceded by a dot - if so, it's not the iterator
			isPrecededByDot := false
			if i > 0 && tokens[i-1].Type == hclsyntax.TokenDot {
				isPrecededByDot = true
			}

			if isPrecededByDot {
				// This is something like "local.origins" - not the iterator
				result = append(result, token)
				continue
			}

			// This could be the iterator - check what follows
			if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenDot {
				if i+2 < len(tokens) && tokens[i+2].Type == hclsyntax.TokenIdent {
					next := string(tokens[i+2].Bytes)
					if next == "value" {
						// iterator.value -> value
						result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
						i += 2 // Skip .value
						continue
					} else if next == "key" {
						// iterator.key -> key
						result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("key")})
						i += 2 // Skip .key
						continue
					} else {
						// iterator.something -> value.something
						result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
						// Keep the dot and the rest
						continue
					}
				}
			} else {
				// Just the iterator by itself -> value
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
				continue
			}
		}

		// Keep the token as-is
		result = append(result, token)
	}

	return result
}

// isHostHeader checks if the token represents "Host" string
func isHostHeader(tokens hclwrite.Tokens) bool {
	if len(tokens) == 0 {
		return false
	}
	tokenStr := string(tokens[0].Bytes)
	return tokenStr == `"Host"` || tokenStr == "Host"
}

// isLoadBalancerPoolResource checks if a block is a load_balancer_pool resource
func isLoadBalancerPoolResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_load_balancer_pool"
}

// transformLoadBalancerPoolHeaders fixes the syntax of partially migrated resources
// This is done at the string level before HCL parsing to fix syntax errors introduced by Grit
func transformLoadBalancerPoolHeaders(content string) string {
	// Fix malformed blocks that Grit has partially converted
	// Grit converts blocks to attributes with objects/lists but leaves nested blocks
	// which creates invalid syntax like: header { ... } instead of header = { ... }

	// Fix the missing = signs to make the file parseable
	result := fixMalformedBlocks(content)

	return result
}

// fixMalformedBlocks adds missing = signs to blocks that should be attributes
func fixMalformedBlocks(content string) string {
	// Fix various blocks that Grit leaves in invalid state
	// These blocks appear inside objects/lists and need = signs

	// Fix header blocks inside origins
	result := regexp.MustCompile(`(\s+)header\s*\{`).ReplaceAllString(content, `${1}header = {`)

	// Fix region_pools, pop_pools, country_pools blocks inside overrides or rules
	// Match when they appear without = sign
	result = regexp.MustCompile(`(\s+)(region_pools|pop_pools|country_pools)(\s+)\{`).ReplaceAllString(result, `${1}${2} = {`)

	// Fix overrides blocks if they're missing =
	result = regexp.MustCompile(`(\s+)overrides\s+\{`).ReplaceAllString(result, `${1}overrides = {`)

	return result
}
