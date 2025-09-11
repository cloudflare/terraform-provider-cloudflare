package main

import (
	"regexp"
	"sort"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// transformLoadBalancerPoolBlock transforms a load_balancer_pool resource block
func transformLoadBalancerPoolBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle dynamic origins blocks by converting them to for expressions
	transformDynamicOriginsBlocks(block, diags)
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
		// We want: [for value in <for_each_expr> : { ... }]
		tokens := hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("for")},
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")},
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
			// Add attribute name
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Get the attribute expression tokens
			attrTokens := attr.Expr().BuildTokens(nil)

			// Replace iterator references with "value"
			processedTokens := replaceIteratorReferences(attrTokens, iteratorName)
			tokens = append(tokens, processedTokens...)

			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Handle nested blocks within content (e.g., header blocks)
		for _, nestedBlock := range contentBlock.Body().Blocks() {
			if nestedBlock.Type() == "header" {
				// Convert header block to header attribute
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    header")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

				// Process header block attributes
				for hName, hAttr := range nestedBlock.Body().Attributes() {
					// Transform "header" = "Host" to "host" = values
					if hName == "header" {
						// Check if the value is "Host"
						headerTokens := hAttr.Expr().BuildTokens(nil)
						if isHostHeader(headerTokens) {
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      host")})
						} else {
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      " + hName)})
						}
					} else if hName == "values" {
						// Skip the "values" key if we're dealing with Host header
						continue
					} else {
						tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      " + hName)})
					}
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

					// For "Host" header, use the values directly
					if hName == "header" && isHostHeader(hAttr.Expr().BuildTokens(nil)) {
						// Find the values attribute
						if valuesAttr := nestedBlock.Body().GetAttribute("values"); valuesAttr != nil {
							valuesTokens := valuesAttr.Expr().BuildTokens(nil)
							processedTokens := replaceIteratorReferences(valuesTokens, iteratorName)
							tokens = append(tokens, processedTokens...)
						}
					} else {
						attrTokens := hAttr.Expr().BuildTokens(nil)
						processedTokens := replaceIteratorReferences(attrTokens, iteratorName)
						tokens = append(tokens, processedTokens...)
					}
					tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
				}

				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
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

// replaceIteratorReferences replaces references to the iterator variable with "value"
func replaceIteratorReferences(tokens hclwrite.Tokens, iteratorName string) hclwrite.Tokens {
	result := hclwrite.Tokens{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == iteratorName {
			// Check if it's followed by .value to handle the special case
			if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenDot &&
				i+2 < len(tokens) && tokens[i+2].Type == hclsyntax.TokenIdent && string(tokens[i+2].Bytes) == "value" {
				// Replace "iterator.value" with just "value"
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
				// Skip the next two tokens (.value)
				i += 2
			} else if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenDot {
				// Replace iterator.something with value.something
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
				// Skip the iterator token, keep the dot
			} else {
				// Just replace the iterator with value
				result = append(result, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
			}
		} else {
			result = append(result, token)
		}
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
	// This handles the case where origins have been converted to a list but header remains as a block
	// The pattern needs to be flexible with whitespace and indentation
	headerBlockPattern := regexp.MustCompile(`(?m)([ \t]*)header\s*\{\s*\n[ \t]*header\s*=\s*"Host"\s*\n[ \t]*values\s*=\s*(\[[^\]]+\])\s*\n[ \t]*\}`)

	// Replace header blocks with header attributes
	result := headerBlockPattern.ReplaceAllString(content, `${1}header = { host = ${2} }`)

	return result
}

