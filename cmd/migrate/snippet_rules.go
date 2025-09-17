package main

import (
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isSnippetRulesResource checks if a block is a cloudflare_snippet_rules resource
func isSnippetRulesResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_snippet_rules"
}

// transformSnippetRulesBlock handles migration of cloudflare_snippet_rules resources in config
// Main transformation: Convert "rules" blocks to list attribute
func transformSnippetRulesBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()

	// Find and collect all rules blocks
	rulesBlocks := body.Blocks()
	var blocksToRemove []*hclwrite.Block
	var hasRules bool

	// First pass: check if we have rules blocks
	for _, rulesBlock := range rulesBlocks {
		if rulesBlock.Type() == "rules" {
			hasRules = true
			blocksToRemove = append(blocksToRemove, rulesBlock)
		}
	}

	// Remove all rules blocks
	for _, rulesBlock := range blocksToRemove {
		body.RemoveBlock(rulesBlock)
	}

	// Add rules as a list attribute if we have any rules
	if hasRules {
		// Build the rules list using raw tokens to preserve formatting
		var rulesTokens hclwrite.Tokens
		rulesTokens = append(rulesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte("["),
		})
		rulesTokens = append(rulesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		})

		// Process each rules block
		for i, rulesBlock := range blocksToRemove {
			if rulesBlock.Type() == "rules" {
				// Add indentation
				rulesTokens = append(rulesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("    "),
				})
				rulesTokens = append(rulesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenOBrace,
					Bytes: []byte("{"),
				})
				rulesTokens = append(rulesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenNewline,
					Bytes: []byte("\n"),
				})

				ruleBody := rulesBlock.Body()

				// Process attributes in order: enabled, expression, description, snippet_name
				// This maintains a consistent order in the output
				attributeOrder := []string{"enabled", "expression", "description", "snippet_name"}
				
				for _, attrName := range attributeOrder {
					if attr := ruleBody.GetAttribute(attrName); attr != nil {
						rulesTokens = append(rulesTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenIdent,
							Bytes: []byte("      " + attrName),
						})
						
						// Calculate padding for alignment
						padding := 14 - len(attrName) // Align to "snippet_name" length
						if padding < 1 {
							padding = 1
						}
						paddingBytes := make([]byte, padding)
						for j := range paddingBytes {
							paddingBytes[j] = ' '
						}
						rulesTokens = append(rulesTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenIdent,
							Bytes: paddingBytes,
						})
						
						rulesTokens = append(rulesTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenEqual,
							Bytes: []byte("="),
						})
						rulesTokens = append(rulesTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenIdent,
							Bytes: []byte(" "),
						})
						
						// Special handling for snippet_name attribute
						if attrName == "snippet_name" {
							// Get the expression and check if it's a reference to cloudflare_snippet.*.name
							exprTokens := attr.Expr().BuildTokens(nil)
							exprStr := string(exprTokens.Bytes())
							
							// Check if this references a cloudflare_snippet resource's name attribute
							// Pattern: cloudflare_snippet.<resource_name>.name
							if strings.Contains(exprStr, "cloudflare_snippet.") && strings.HasSuffix(strings.TrimSpace(exprStr), ".name") {
								// Replace .name with .snippet_name
								modifiedExprStr := strings.TrimSuffix(strings.TrimSpace(exprStr), ".name") + ".snippet_name"
								rulesTokens = append(rulesTokens, &hclwrite.Token{
									Type:  hclsyntax.TokenIdent,
									Bytes: []byte(modifiedExprStr),
								})
							} else {
								// Use the original expression tokens
								rulesTokens = append(rulesTokens, exprTokens...)
							}
						} else {
							// For other attributes, use the original tokens
							rulesTokens = append(rulesTokens, attr.Expr().BuildTokens(nil)...)
						}
						
						rulesTokens = append(rulesTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenNewline,
							Bytes: []byte("\n"),
						})
					}
				}

				// Close object
				rulesTokens = append(rulesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("    "),
				})
				rulesTokens = append(rulesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenCBrace,
					Bytes: []byte("}"),
				})

				// Add comma if not the last item
				if i < len(blocksToRemove)-1 {
					rulesTokens = append(rulesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenComma,
						Bytes: []byte(","),
					})
				}
				rulesTokens = append(rulesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenNewline,
					Bytes: []byte("\n"),
				})
			}
		}

		// Close list
		rulesTokens = append(rulesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte("  "),
		})
		rulesTokens = append(rulesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte("]"),
		})

		// Set the rules attribute using raw tokens
		body.SetAttributeRaw("rules", rulesTokens)
	}
}