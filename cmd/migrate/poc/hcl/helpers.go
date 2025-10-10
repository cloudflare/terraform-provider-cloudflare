// Package hcl provides common helper functions for HCL manipulation
// Mostly all from the previous cmd/migrate implementation
package hcl

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// AttributeInfo holds an attribute name and its corresponding Attribute object
type AttributeInfo struct {
	Name      string
	Attribute *hclwrite.Attribute
}

// AttributesOrdered returns attributes from a body in their original order
// This is important when generating HCL that needs to maintain specific field ordering
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

// BuildTemplateStringTokens creates tokens for a template string like "${expr}/literal"
// This is useful for creating import block IDs and other template expressions
func BuildTemplateStringTokens(exprTokens hclwrite.Tokens, suffix string) hclwrite.Tokens {
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

// BuildResourceReference creates tokens for a resource reference like "type.name"
// Used for creating references to resources in moved blocks and import blocks
func BuildResourceReference(resourceType, resourceName string) hclwrite.Tokens {
	return hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte(resourceType)},
		{Type: hclsyntax.TokenDot, Bytes: []byte{'.'}},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(resourceName)},
	}
}

// CreateMovedBlock creates a moved block for resource migration
// This is used when resources are renamed or restructured between provider versions
func CreateMovedBlock(from, to string) *hclwrite.Block {
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

// CreateImportBlock creates an import block for a resource
// Used for generating import blocks when transforming resources
func CreateImportBlock(resourceType, resourceName, importID string) *hclwrite.Block {
	block := hclwrite.NewBlock("import", nil)
	body := block.Body()

	// Build the "to" value: resource_type.resource_name
	toTokens := BuildResourceReference(resourceType, resourceName)
	body.SetAttributeRaw("to", toTokens)

	// Set the import ID
	body.SetAttributeValue("id", cty.StringVal(importID))

	return block
}

// CreateImportBlockWithTokens creates an import block using raw tokens for the ID
// This variant is useful when the import ID needs to be a template expression
func CreateImportBlockWithTokens(resourceType, resourceName string, idTokens hclwrite.Tokens) *hclwrite.Block {
	block := hclwrite.NewBlock("import", nil)
	body := block.Body()

	// Build the "to" value: resource_type.resource_name
	toTokens := BuildResourceReference(resourceType, resourceName)
	body.SetAttributeRaw("to", toTokens)

	// Set the ID using raw tokens
	body.SetAttributeRaw("id", idTokens)

	return block
}

// BuildObjectFromBlock creates object tokens from a block's attributes
// Useful for converting block syntax to object syntax
func BuildObjectFromBlock(block *hclwrite.Block) hclwrite.Tokens {
	// Get attributes in their original order
	orderedAttrs := AttributesOrdered(block.Body())

	// Build a list of attribute tokens preserving the original order
	var attrs []hclwrite.ObjectAttrTokens

	for _, attrInfo := range orderedAttrs {
		// Create tokens for the attribute name (as a simple identifier)
		nameTokens := hclwrite.TokensForIdentifier(attrInfo.Name)

		// Get the value tokens from the attribute's expression
		valueTokens := attrInfo.Attribute.Expr().BuildTokens(nil)

		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  nameTokens,
			Value: valueTokens,
		})
	}

	// Use the built-in TokensForObject function to create properly formatted object tokens
	return hclwrite.TokensForObject(attrs)
}

// SetAttributeValue is a helper that sets an attribute value based on its Go type
// It automatically converts common Go types to their cty equivalents
func SetAttributeValue(body *hclwrite.Body, name string, val interface{}) {
	switch v := val.(type) {
	case string:
		body.SetAttributeValue(name, cty.StringVal(v))
	case int:
		body.SetAttributeValue(name, cty.NumberIntVal(int64(v)))
	case int64:
		body.SetAttributeValue(name, cty.NumberIntVal(v))
	case float64:
		body.SetAttributeValue(name, cty.NumberFloatVal(v))
	case bool:
		body.SetAttributeValue(name, cty.BoolVal(v))
	case []string:
		values := make([]cty.Value, len(v))
		for i, s := range v {
			values[i] = cty.StringVal(s)
		}
		body.SetAttributeValue(name, cty.ListVal(values))
	case map[string]string:
		values := make(map[string]cty.Value)
		for k, v := range v {
			values[k] = cty.StringVal(v)
		}
		body.SetAttributeValue(name, cty.ObjectVal(values))
	default:
		// For complex types, caller should use SetAttributeRaw with tokens
		// or SetAttributeValue with a properly constructed cty.Value
	}
}

// CopyAttribute copies an attribute from one body to another, preserving its expression
func CopyAttribute(from, to *hclwrite.Body, attrName string) {
	if attr := from.GetAttribute(attrName); attr != nil {
		tokens := attr.Expr().BuildTokens(nil)
		to.SetAttributeRaw(attrName, tokens)
	}
}

// RemoveEmptyBlocks removes blocks with no attributes or nested blocks
func RemoveEmptyBlocks(body *hclwrite.Body, blockType string) {
	var blocksToRemove []*hclwrite.Block

	for _, block := range body.Blocks() {
		if block.Type() == blockType {
			blockBody := block.Body()
			if len(blockBody.Attributes()) == 0 && len(blockBody.Blocks()) == 0 {
				blocksToRemove = append(blocksToRemove, block)
			}
		}
	}

	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}
}

// TokensForSimpleValue creates tokens for a simple value (string, number, bool)
func TokensForSimpleValue(val interface{}) hclwrite.Tokens {
	switch v := val.(type) {
	case string:
		return hclwrite.TokensForValue(cty.StringVal(v))
	case int:
		return hclwrite.TokensForValue(cty.NumberIntVal(int64(v)))
	case int64:
		return hclwrite.TokensForValue(cty.NumberIntVal(v))
	case float64:
		return hclwrite.TokensForValue(cty.NumberFloatVal(v))
	case bool:
		return hclwrite.TokensForValue(cty.BoolVal(v))
	default:
		return nil
	}
}
