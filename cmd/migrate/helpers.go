package main

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// AttributeInfo holds an attribute name and its corresponding Attribute object
type AttributeInfo struct {
	Name      string
	Attribute *hclwrite.Attribute
}

// AttributesOrdered returns attributes from a body in their original order
func AttributesOrdered(body *hclwrite.Body) []AttributeInfo {
	// Get all attributes as a map for lookup
	attrMap := body.Attributes()

	// Get tokens to find the original order
	tokens := body.BuildTokens(nil)

	var orderedAttrs []AttributeInfo
	seenAttrs := make(map[string]bool)

	// Scan through tokens to find attribute names in order
	for i := range tokens {
		token := tokens[i]

		// Look for identifier tokens that could be attribute names
		if token.Type == hclsyntax.TokenIdent && i+1 < len(tokens) {
			// Check if the next token is an equals sign
			nextToken := tokens[i+1]
			if nextToken.Type == hclsyntax.TokenEqual {
				attrName := string(token.Bytes)

				// Check if this is actually an attribute and we haven't seen it yet
				if attr, exists := attrMap[attrName]; exists && !seenAttrs[attrName] {
					orderedAttrs = append(orderedAttrs, AttributeInfo{
						Name:      attrName,
						Attribute: attr,
					})
					seenAttrs[attrName] = true
				}
			}
		}
	}

	return orderedAttrs
}

// buildTemplateStringTokens creates tokens for a template string like "${expr}/literal"
func buildTemplateStringTokens(exprTokens hclwrite.Tokens, suffix string) hclwrite.Tokens {
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
		{Type: hclsyntax.TokenTemplateInterp, Bytes: []byte("${")},
	}

	tokens = append(tokens, exprTokens...)
	tokens = append(tokens,
		&hclwrite.Token{Type: hclsyntax.TokenTemplateSeqEnd, Bytes: []byte{'}'}},
	)

	if suffix != "" {
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenTemplateControl, Bytes: []byte(suffix)})
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}})

	return tokens
}

// buildResourceReference creates tokens for a resource reference like "type.name"
func buildResourceReference(resourceType, resourceName string) hclwrite.Tokens {
	return hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte(resourceType)},
		{Type: hclsyntax.TokenDot, Bytes: []byte{'.'}},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(resourceName)},
	}
}

// createMovedBlock creates a moved block for resource migration
func createMovedBlock(from, to string) *hclwrite.Block {
	block := hclwrite.NewBlock("moved", nil)
	body := block.Body()

	// Create traversals for from and to
	fromParts := strings.Split(from, ".")
	toParts := strings.Split(to, ".")

	// Build from traversal
	fromTraversal := hcl.Traversal{}
	for i, part := range fromParts {
		if i == 0 {
			fromTraversal = append(fromTraversal, hcl.TraverseRoot{Name: part})
		} else {
			fromTraversal = append(fromTraversal, hcl.TraverseAttr{Name: part})
		}
	}

	// Build to traversal
	toTraversal := hcl.Traversal{}
	for i, part := range toParts {
		if i == 0 {
			toTraversal = append(toTraversal, hcl.TraverseRoot{Name: part})
		} else {
			toTraversal = append(toTraversal, hcl.TraverseAttr{Name: part})
		}
	}

	body.SetAttributeTraversal("from", fromTraversal)
	body.SetAttributeTraversal("to", toTraversal)

	return block
}
