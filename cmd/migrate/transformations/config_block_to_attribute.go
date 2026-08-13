package transformations

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
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

	// Special handling for cloudflare_page_rule to remove unsupported blocks
	if resourceType == "cloudflare_page_rule" {
		// Remove minify blocks from actions block as they're not supported in v5
		removeUnsupportedPageRuleBlocks(body)
	}

	// Transform blocks to maps
	for _, blockName := range transform.ToMap {
		if err := transformBlocksToMapWithResource(config, body, blockName, resourceType); err != nil {
			return fmt.Errorf("failed to transform %s to map: %w", blockName, err)
		}
	}

	// Transform blocks to lists
	for _, blockName := range transform.ToList {
		if err := transformBlocksToListWithResource(config, body, blockName, resourceType); err != nil {
			return fmt.Errorf("failed to transform %s to list: %w", blockName, err)
		}
	}

	return nil
}

// transformBlocksToList transforms multiple blocks with the same name into a list attribute
func transformBlocksToList(config *TransformationConfig, body *hclwrite.Body, blockName string) error {
	return transformBlocksToListWithResource(config, body, blockName, "")
}

// transformBlocksToListWithResource transforms multiple blocks with resource context
func transformBlocksToListWithResource(config *TransformationConfig, body *hclwrite.Body, blockName string, resourceType string) error {
	blocks := findBlocksByType(body, blockName)
	if len(blocks) == 0 {
		return nil // No blocks to transform
	}

	// Transform blocks to list using raw tokens to preserve complex expressions
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
	}

	for i, block := range blocks {
		// Before converting block, transform nested blocks if needed
		if resourceType != "" && config != nil {
			// Apply transformations to nested blocks within this block
			// For example, 'rules' block in 'cloudflare_load_balancer' has 'overrides' that needs to be a map
			if transform, exists := config.Transformations[resourceType]; exists {
				blockBody := block.Body()

				// Transform nested blocks to maps
				for _, nestedBlockName := range transform.ToMap {
					if nestedBlocks := findBlocksByType(blockBody, nestedBlockName); len(nestedBlocks) > 0 {
						transformBlocksToMapWithResource(config, blockBody, nestedBlockName, resourceType)
					}
				}

				// Transform nested blocks to lists
				for _, nestedBlockName := range transform.ToList {
					if nestedBlocks := findBlocksByType(blockBody, nestedBlockName); len(nestedBlocks) > 0 {
						transformBlocksToListWithResource(config, blockBody, nestedBlockName, resourceType)
					}
				}
			}
		}

		if i > 0 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(", ")})
		}

		// Convert block to tokens (preserves the transformed attributes)
		blockTokens := blockToTokensWithConfig(config, block, resourceType)
		tokens = append(tokens, blockTokens...)
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})

	// Set the attribute using raw tokens
	body.SetAttributeRaw(blockName, tokens)

	// Remove original blocks
	for _, block := range blocks {
		body.RemoveBlock(block)
	}

	return nil
}

// transformBlocksToMap transforms a single block into a map attribute
func transformBlocksToMap(config *TransformationConfig, body *hclwrite.Body, blockName string) error {
	return transformBlocksToMapWithResource(config, body, blockName, "")
}

// transformBlocksToMapWithResource transforms a single block into a map attribute with resource context
func transformBlocksToMapWithResource(config *TransformationConfig, body *hclwrite.Body, blockName string, resourceType string) error {
	blocks := findBlocksByType(body, blockName)
	if len(blocks) == 0 {
		return nil // No blocks to transform
	}
	if len(blocks) > 1 {
		// Multiple blocks found, use first one
		// Note: This is expected in some cases where duplicate blocks are intentional
	}

	block := blocks[0]

	// Always use raw token approach to preserve complex expressions
	// This ensures resource references and other complex expressions are preserved correctly
	err := transformBlockToMapRawWithConfig(config, body, block, blockName, resourceType)
	if err != nil {
		return err
	}
	// Remove original block
	body.RemoveBlock(block)

	return nil
}

// transformBlockToMapRaw converts a block to a map attribute using raw tokens to preserve complex expressions
func transformBlockToMapRaw(body *hclwrite.Body, block *hclwrite.Block, blockName string) error {
	return transformBlockToMapRawWithConfig(nil, body, block, blockName, "")
}

// transformBlockToMapRawWithConfig converts a block to a map attribute using raw tokens with config awareness
func transformBlockToMapRawWithConfig(config *TransformationConfig, body *hclwrite.Body, block *hclwrite.Block, blockName string, parentResourceType string) error {
	blockBody := block.Body()
	attrs := blockBody.Attributes()
	blocks := blockBody.Blocks()

	if len(attrs) == 0 && len(blocks) == 0 {
		// Empty block - create empty object using raw tokens
		body.SetAttributeRaw(blockName, hclwrite.Tokens{
			{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
			{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
		})
		return nil
	}

	// Build the map manually using tokens to preserve complex expressions
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
	}

	// Sort attribute names for consistent output
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	first := true
	for _, name := range names {
		attr := attrs[name]
		if !first {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		first = false

		// Add indentation
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})

		// Add the attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Add the attribute expression tokens (preserves complex expressions)
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)
	}

	// Process nested blocks without grouping - preserve all blocks including duplicates
	for _, nestedBlock := range blocks {
		if !first {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		first = false

		// Add indentation
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})

		// Add the nested block name
		nestedName := nestedBlock.Type()
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedName)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Convert the nested block to tokens
		nestedTokens := blockToTokensWithConfig(config, nestedBlock, parentResourceType)
		tokens = append(tokens, nestedTokens...)
	}

	tokens = append(tokens,
		&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")},
		&hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
	)

	// Set the attribute using raw tokens
	body.SetAttributeRaw(blockName, tokens)

	return nil
}

// blockToTokens converts a block to HCL tokens representing an object
func blockToTokens(block *hclwrite.Block) hclwrite.Tokens {
	return blockToTokensWithConfig(nil, block, "")
}

// blockToTokensWithConfig converts a block to HCL tokens with config awareness
func blockToTokensWithConfig(config *TransformationConfig, block *hclwrite.Block, parentResourceType string) hclwrite.Tokens {
	blockBody := block.Body()
	attrs := blockBody.Attributes()
	blocks := blockBody.Blocks()

	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
	}

	if len(attrs) == 0 && len(blocks) == 0 {
		// Empty block
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
		return tokens
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

	// Sort attribute names for consistent output
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	first := true
	// Process attributes
	for _, name := range names {
		attr := attrs[name]
		if !first {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		first = false

		// Add indentation (6 spaces for nested content)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})

		// Add the attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Add the attribute expression tokens
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)
	}

	// Process nested blocks without grouping - preserve all blocks including duplicates
	for _, nestedBlock := range blocks {
		if !first {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		first = false

		// Add indentation (6 spaces for nested content)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})

		// Add the nested block name
		nestedName := nestedBlock.Type()
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedName)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Convert the nested block to tokens
		nestedTokens := blockToTokensWithConfig(config, nestedBlock, parentResourceType)
		tokens = append(tokens, nestedTokens...)
	}

	tokens = append(tokens,
		&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")},
		&hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
	)

	return tokens
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

// removeUnsupportedPageRuleBlocks removes blocks that are not supported in v5
func removeUnsupportedPageRuleBlocks(body *hclwrite.Body) {
	// Find all actions blocks
	actionsBlocks := findBlocksByType(body, "actions")
	for _, actionsBlock := range actionsBlocks {
		actionsBody := actionsBlock.Body()

		// Find and remove minify blocks
		minifyBlocks := findBlocksByType(actionsBody, "minify")
		for _, minifyBlock := range minifyBlocks {
			actionsBody.RemoveBlock(minifyBlock)
		}

		// Also remove minify attribute if it exists
		if actionsBody.GetAttribute("minify") != nil {
			actionsBody.RemoveAttribute("minify")
		}

		// Remove disable_railgun attribute if it exists
		if actionsBody.GetAttribute("disable_railgun") != nil {
			actionsBody.RemoveAttribute("disable_railgun")
		}
	}
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
