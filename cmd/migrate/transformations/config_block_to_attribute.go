package transformations

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// TransformResourceBlock transforms blocks within a resource according to the configuration
func TransformResourceBlock(config *TransformationConfig, resourceBlock *hclwrite.Block, resourceType string) error {
	transform, exists := config.Transformations[resourceType]
	if !exists {
		// No transformations defined for this resource type
		return nil
	}

	body := resourceBlock.Body()

	// Transform blocks to maps
	for _, blockName := range transform.ToMap {
		if err := transformBlocksToMap(config, body, blockName); err != nil {
			return fmt.Errorf("failed to transform %s to map: %w", blockName, err)
		}
	}

	// Transform blocks to lists
	for _, blockName := range transform.ToList {
		if err := transformBlocksToList(config, body, blockName); err != nil {
			return fmt.Errorf("failed to transform %s to list: %w", blockName, err)
		}
	}

	return nil
}

// transformBlocksToList transforms multiple blocks with the same name into a list attribute
func transformBlocksToList(config *TransformationConfig, body *hclwrite.Body, blockName string) error {
	blocks := findBlocksByType(body, blockName)
	if len(blocks) == 0 {
		return nil // No blocks to transform
	}

	// Build list of objects from blocks
	var objects []cty.Value
	for _, block := range blocks {
		obj := blockToObject(block)
		objects = append(objects, obj)
	}

	// Create the list value
	listVal := cty.TupleVal(objects)

	// Set the attribute
	body.SetAttributeValue(blockName, listVal)

	// Remove original blocks
	for _, block := range blocks {
		body.RemoveBlock(block)
	}

	return nil
}

// transformBlocksToMap transforms a single block into a map attribute
func transformBlocksToMap(config *TransformationConfig, body *hclwrite.Body, blockName string) error {
	blocks := findBlocksByType(body, blockName)
	if len(blocks) == 0 {
		return nil // No blocks to transform
	}

	if len(blocks) > 1 {
		// Multiple blocks found, use first and warn
		fmt.Printf("Warning: Multiple %s blocks found, using first one\n", blockName)
	}

	block := blocks[0]

	// Convert block to object
	obj := blockToObject(block)

	// Set the attribute
	body.SetAttributeValue(blockName, obj)

	// Remove original block
	body.RemoveBlock(block)

	return nil
}

// blockToObject converts a block's content to a cty.Value object
func blockToObject(block *hclwrite.Block) cty.Value {
	attrs := make(map[string]cty.Value)
	body := block.Body()

	// Process attributes
	for name, attr := range body.Attributes() {
		// Parse the expression to get its value
		expr := attr.Expr()
		tokens := expr.BuildTokens(nil)

		// Simple heuristic to determine value type from tokens
		val := tokensToValue(tokens)
		attrs[name] = val
	}

	// Process nested blocks
	for _, nestedBlock := range body.Blocks() {
		nestedName := nestedBlock.Type()
		nestedObj := blockToObject(nestedBlock)
		attrs[nestedName] = nestedObj
	}

	return cty.ObjectVal(attrs)
}

// tokensToValue attempts to convert tokens to a cty.Value
func tokensToValue(tokens hclwrite.Tokens) cty.Value {
	// Build the string representation
	var sb strings.Builder
	for _, token := range tokens {
		sb.Write(token.Bytes)
	}
	str := strings.TrimSpace(sb.String())

	// Try to determine the type
	if str == "true" {
		return cty.BoolVal(true)
	}
	if str == "false" {
		return cty.BoolVal(false)
	}
	if strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]") {
		// It's a list - for simplicity, treat as string list
		inner := strings.Trim(str, "[]")
		if inner == "" {
			return cty.ListValEmpty(cty.String)
		}
		// Parse list items (simplified)
		items := strings.Split(inner, ",")
		var values []cty.Value
		for _, item := range items {
			item = strings.TrimSpace(item)
			item = strings.Trim(item, "\"")
			values = append(values, cty.StringVal(item))
		}
		if len(values) > 0 {
			return cty.ListVal(values)
		}
		return cty.ListValEmpty(cty.String)
	}
	if strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"") {
		// String value
		return cty.StringVal(strings.Trim(str, "\""))
	}
	// Try as number
	if num, err := strconv.ParseInt(str, 10, 64); err == nil {
		return cty.NumberIntVal(num)
	}
	if num, err := strconv.ParseFloat(str, 64); err == nil {
		return cty.NumberFloatVal(num)
	}

	// Default to string
	return cty.StringVal(str)
}

// findBlocksByType finds all blocks with a specific type in the body
func findBlocksByType(body *hclwrite.Body, blockType string) []*hclwrite.Block {
	var result []*hclwrite.Block
	for _, block := range body.Blocks() {
		if block.Type() == blockType {
			result = append(result, block)
		}
	}
	return result
}

// GetTransformationType returns the transformation type for a given resource and block
func GetTransformationType(config *TransformationConfig, resourceType, blockName string) string {
	transform, exists := config.Transformations[resourceType]
	if !exists {
		return ""
	}

	for _, name := range transform.ToMap {
		if name == blockName {
			return "map"
		}
	}

	for _, name := range transform.ToList {
		if name == blockName {
			return "list"
		}
	}

	return ""
}