package main

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isAccessApplicationResource checks if a block is a cloudflare_zero_trust_access_application resource
// (grit has already renamed from cloudflare_access_application)
func isAccessApplicationResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_application"
}

// transformAccessApplicationBlock transforms policies attribute from list of strings to list of objects
// Before: policies = ["id1", "id2"]
// After: policies = [{id = "id1"}, {id = "id2"}]
func transformAccessApplicationBlock(block *hclwrite.Block) {
	policiesAttr := block.Body().GetAttribute("policies")
	if policiesAttr == nil {
		return // No policies attribute
	}

	// Get the expression tokens
	tokens := policiesAttr.Expr().BuildTokens(nil)
	transformedTokens := transformPoliciesAttribute(tokens)
	if transformedTokens != nil {
		block.Body().SetAttributeRaw("policies", transformedTokens)
	}
}

// transformPoliciesAttribute transforms a policies list from strings to objects
func transformPoliciesAttribute(tokens hclwrite.Tokens) hclwrite.Tokens {
	var result hclwrite.Tokens
	i := 0
	inList := false

	for i < len(tokens) {
		token := tokens[i]

		// Track when we enter the list
		if token.Type == hclsyntax.TokenOBrack {
			inList = true
			result = append(result, token)
			i++
			continue
		}

		// Track when we exit the list
		if token.Type == hclsyntax.TokenCBrack {
			inList = false
			result = append(result, token)
			i++
			continue
		}

		// When inside the list, transform values to objects
		if inList {
			// Look for identifiers that could be policy references (e.g., cloudflare_zero_trust_access_policy.example.id)
			if token.Type == hclsyntax.TokenIdent {
				// Collect the full reference (could be multiple tokens for a.b.c)
				refTokens := collectReference(tokens[i:])
				if len(refTokens) > 0 && isPolicyReference(refTokens) {
					// Transform to {id = reference}
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenOBrace,
						Bytes: []byte("{"),
					})
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("id"),
					})
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenEqual,
						Bytes: []byte(" = "),
					})
					result = append(result, refTokens...)
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenCBrace,
						Bytes: []byte("}"),
					})
					i += len(refTokens)
					continue
				}
			}

			// Look for opening quotes (literal policy IDs)
			if token.Type == hclsyntax.TokenOQuote {
				// Check if next token is a quoted literal
				if i+2 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenQuotedLit && tokens[i+2].Type == hclsyntax.TokenCQuote {
					// Transform to {id = "string"}
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenOBrace,
						Bytes: []byte("{"),
					})
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("id"),
					})
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenEqual,
						Bytes: []byte(" = "),
					})
					// Include all three tokens: OQuote, QuotedLit, CQuote
					result = append(result, token)           // TokenOQuote
					result = append(result, tokens[i+1])     // TokenQuotedLit
					result = append(result, tokens[i+2])     // TokenCQuote
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenCBrace,
						Bytes: []byte("}"),
					})
					i += 3
					continue
				}
			}
		}

		result = append(result, token)
		i++
	}

	return result
}

// collectReference collects tokens that form a reference (e.g., cloudflare_zero_trust_access_policy.example.id)
func collectReference(tokens hclwrite.Tokens) hclwrite.Tokens {
	var result hclwrite.Tokens
	i := 0

	for i < len(tokens) {
		token := tokens[i]

		// References are made of identifiers and dots
		if token.Type == hclsyntax.TokenIdent {
			result = append(result, token)
			i++
			// Check for dot and continue
			if i < len(tokens) && tokens[i].Type == hclsyntax.TokenDot {
				result = append(result, tokens[i])
				i++
				continue
			}
			break
		} else {
			break
		}
	}

	return result
}

// isPolicyReference checks if tokens look like a policy reference
func isPolicyReference(tokens hclwrite.Tokens) bool {
	if len(tokens) < 5 {
		// Need at least: resource_type.name.id
		return false
	}

	// Check if it ends with .id
	if len(tokens) >= 2 {
		lastToken := tokens[len(tokens)-1]
		secondLastToken := tokens[len(tokens)-2]
		if lastToken.Type == hclsyntax.TokenIdent && string(lastToken.Bytes) == "id" &&
			secondLastToken.Type == hclsyntax.TokenDot {
			// Check if it starts with a known policy resource type
			firstToken := tokens[0]
			if firstToken.Type == hclsyntax.TokenIdent {
				resourceType := string(firstToken.Bytes)
				return resourceType == "cloudflare_zero_trust_access_policy" ||
					resourceType == "cloudflare_access_policy" // Handle both old and new names
			}
		}
	}

	return false
}