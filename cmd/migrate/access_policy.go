package main

import (
	"log"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isAccessPolicyResource checks if a block is a cloudflare_zero_trust_access_policy resource
// (grit has already renamed from cloudflare_access_policy)
func isAccessPolicyResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_policy"
}

// transformAccessPolicyBlock transforms include/exclude/require attributes
// Currently only handles boolean attributes like everyone, certificate, any_valid_service_token
// After grit: include = [{ everyone = true }]
// After this: include = [{ everyone = {} }]
func transformAccessPolicyBlock(block *hclwrite.Block) {
	// Process include, exclude, and require attributes (grit has already converted them to lists)
	conditionAttributes := []string{"include", "exclude", "require"}

	for _, attrName := range conditionAttributes {
		attr := block.Body().GetAttribute(attrName)
		if attr == nil {
			continue // Attribute doesn't exist
		}

		transformedTokens := transformBooleanAttributesInList2(*attr)
		if transformedTokens != nil {
			block.Body().SetAttributeRaw(attrName, transformedTokens)
		}
	}
}

func transformBooleanAttributesInList2(attr hclwrite.Attribute) hclwrite.Tokens {
	expr, d := ast.WriteExpr2Expr(*attr.Expr())
	log.Println(d)

	tup, ok := expr.(*hclsyntax.TupleConsExpr)
	if !ok {
		return nil
	}
	for _, expr := range tup.Exprs {
		obj, ok := expr.(*hclsyntax.ObjectConsExpr)
		if !ok {
			return nil
		}
		acc := []hclsyntax.ObjectConsItem{}
		for _, item := range obj.Items {
			key, _ := ast.Expr2S(item.KeyExpr)
			if isBooleanPolicyAttribute(key) {
				val, _ := ast.Expr2S(item.ValueExpr)
				if val == "false" {
					continue
				}

				acc = append(acc, hclsyntax.ObjectConsItem{
					KeyExpr:   item.KeyExpr,
					ValueExpr: &hclsyntax.ObjectConsExpr{},
				})
			} else {
				acc = append(acc, item)
			}
		}

		obj.Items = acc
	}

	write, d := ast.Expr2WriteExpr(tup)
	log.Println(d)
	return write.BuildTokens(nil)
}

// transformBooleanAttributesInList transforms boolean attributes within a condition list
// Only handles: everyone, certificate, any_valid_service_token
func transformBooleanAttributesInList(tokens hclwrite.Tokens) hclwrite.Tokens {
	var result hclwrite.Tokens
	i := 0

	for i < len(tokens) {
		token := tokens[i]

		// Look for boolean attribute names
		if token.Type == hclsyntax.TokenIdent {
			attrName := string(token.Bytes)
			if isBooleanPolicyAttribute(attrName) && i+1 < len(tokens) {
				nextToken := tokens[i+1]
				if nextToken.Type == hclsyntax.TokenEqual {
					// Found a boolean attribute assignment
					result = append(result, token, nextToken)
					i += 2

					// Skip any whitespace
					for i < len(tokens) && (tokens[i].Type == hclsyntax.TokenNewline || isWhitespaceToken(tokens[i])) {
						result = append(result, tokens[i])
						i++
					}

					// Collect the value (should be true or false)
					if i < len(tokens) {
						valueToken := tokens[i]
						valueStr := strings.TrimSpace(string(valueToken.Bytes))

						if valueStr == "false" {
							// Skip false values - remove the attribute entirely
							// We need to backtrack and remove what we just added
							result = result[:len(result)-2] // Remove attrName and =
							// Also remove any trailing whitespace we added
							for len(result) > 0 && isWhitespaceToken(result[len(result)-1]) {
								result = result[:len(result)-1]
							}
							i++

							// Skip any trailing comma if this was the only/last item
							if i < len(tokens) && tokens[i].Type == hclsyntax.TokenComma {
								i++
							}
							continue
						} else if valueStr == "true" {
							// Replace true with empty object
							result = append(result, &hclwrite.Token{
								Type:  hclsyntax.TokenOBrace,
								Bytes: []byte("{}"),
							})
							i++
							continue
						}
					}
				}
			}
		}

		result = append(result, token)
		i++
	}

	return result
}

// isBooleanPolicyAttribute checks if an attribute is a boolean that should become an empty object
func isBooleanPolicyAttribute(attrName string) bool {
	booleanAttributes := map[string]bool{
		"everyone":                true,
		"certificate":             true,
		"any_valid_service_token": true,
	}
	return booleanAttributes[attrName]
}

// isWhitespaceToken checks if a token is whitespace
func isWhitespaceToken(token *hclwrite.Token) bool {
	return token.Type == hclsyntax.TokenNewline ||
		(len(token.Bytes) > 0 && strings.TrimSpace(string(token.Bytes)) == "")
}
