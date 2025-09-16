package transformations

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/hcl/v2/hclsyntax"
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
		if err := transformBlocksToMap(config, body, blockName); err != nil {
			return fmt.Errorf("failed to transform %s to map: %w", blockName, err)
		}
	}

	// Transform blocks to lists
	for _, blockName := range transform.ToList {
		if err := transformBlocksToList(config, body, blockName, resourceType); err != nil {
			return fmt.Errorf("failed to transform %s to list: %w", blockName, err)
		}
	}

	return nil
}

// transformBlocksToList transforms multiple blocks with the same name into a list attribute
func transformBlocksToList(config *TransformationConfig, body *hclwrite.Body, blockName string, resourceType string) error {
	blocks := findBlocksByType(body, blockName)
	if len(blocks) == 0 {
		return nil // No blocks to transform
	}

	// Special handling for origins blocks in cloudflare_load_balancer_pool which may contain template expressions
	if resourceType == "cloudflare_load_balancer_pool" && blockName == "origins" {
		return transformOriginsBlocksToList(body, blocks)
	}

	// Special handling for rules blocks in cloudflare_ruleset which may contain heredoc expressions
	// Use standard transformation with enhanced post-processing for now
	// The token-based approach has indentation issues that need more work
	// if resourceType == "cloudflare_ruleset" && blockName == "rules" {
	//     return transformRulesBlocksToList(body, blocks)
	// }

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

// transformOriginsBlocksToList handles origins blocks specifically to preserve template expressions
func transformOriginsBlocksToList(body *hclwrite.Body, blocks []*hclwrite.Block) error {
	// Build tokens for the list directly to preserve template expressions
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
	}

	for i, block := range blocks {
		if i > 0 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

		// Start object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

		// Add attributes from block, preserving original expressions
		attrs := block.Body().Attributes()
		attrNames := make([]string, 0, len(attrs))
		for name := range attrs {
			attrNames = append(attrNames, name)
		}
		sort.Strings(attrNames)

		for j, name := range attrNames {
			attr := attrs[name]
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Preserve the original expression tokens
			exprTokens := attr.Expr().BuildTokens(nil)
			tokens = append(tokens, exprTokens...)

			if j < len(attrNames)-1 {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Close object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})

	// Set the attribute using raw tokens to preserve expressions
	body.SetAttributeRaw("origins", tokens)

	// Remove original blocks
	for _, block := range blocks {
		body.RemoveBlock(block)
	}

	return nil
}

// transformRulesBlocksToList handles rules blocks specifically to preserve heredoc expressions
func transformRulesBlocksToList(body *hclwrite.Body, blocks []*hclwrite.Block) error {
	// Build tokens for the list directly to preserve expressions
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
	}

	for i, block := range blocks {
		if i > 0 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

		// Start object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

		// Add attributes from block, preserving original expressions
		attrs := block.Body().Attributes()
		attrNames := make([]string, 0, len(attrs))
		for name := range attrs {
			attrNames = append(attrNames, name)
		}
		sort.Strings(attrNames)

		for j, name := range attrNames {
			attr := attrs[name]
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Preserve the original expression tokens to avoid double dollar signs
			exprTokens := attr.Expr().BuildTokens(nil)
			tokens = append(tokens, exprTokens...)

			// Add comma if there are more attributes or nested blocks
			nestedBlocks := block.Body().Blocks()
			if j < len(attrNames)-1 || len(nestedBlocks) > 0 {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Process nested blocks - convert them to attribute format
		nestedBlocks := block.Body().Blocks()
		for k, nestedBlock := range nestedBlocks {
			blockType := nestedBlock.Type()
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + blockType)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Determine how to transform this nested block based on the configuration
			// For now, convert blocks to objects using a simplified approach
			nestedTokens := convertNestedBlockToTokens(nestedBlock)
			tokens = append(tokens, nestedTokens...)

			if k < len(nestedBlocks)-1 {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Close object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})

	// Set the attribute using raw tokens to preserve expressions
	body.SetAttributeRaw("rules", tokens)

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

	// Always use raw token approach to preserve complex expressions
	// This ensures resource references and other complex expressions are preserved correctly
	transformBlockToMapRaw(body, block, blockName)

	// Remove original block
	body.RemoveBlock(block)

	return nil
}

// transformBlockToMapRaw converts a block to a map attribute using raw tokens to preserve complex expressions
func transformBlockToMapRaw(body *hclwrite.Body, block *hclwrite.Block, blockName string) error {
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

	// Process nested blocks (convert them to nested maps/objects)
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

		// Recursively convert the nested block to tokens
		nestedTokens := blockToTokens(nestedBlock)
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

	// Process nested blocks
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

		// Recursively convert the nested block
		nestedTokens := blockToTokens(nestedBlock)
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
		// String value - remove quotes
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

// blockToObjectTokens converts a block to object format using raw tokens to preserve expressions
func blockToObjectTokens(block *hclwrite.Block) hclwrite.Tokens {
	return blockToObjectTokensWithIndent(block, "      ")
}

// blockToObjectTokensWithIndent converts a block to object format with specified indentation
func blockToObjectTokensWithIndent(block *hclwrite.Block, indent string) hclwrite.Tokens {
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
	}

	// Get attributes and nested blocks
	attrs := block.Body().Attributes()
	nestedBlocks := block.Body().Blocks()
	
	// Process attributes first
	attrNames := make([]string, 0, len(attrs))
	for name := range attrs {
		attrNames = append(attrNames, name)
	}
	sort.Strings(attrNames)

	for j, name := range attrNames {
		attr := attrs[name]
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(indent + name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Preserve original expression tokens
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)

		// Add comma if not the last attribute, or if there are nested blocks
		if j < len(attrNames)-1 || len(nestedBlocks) > 0 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	// Process nested blocks recursively with increased indentation
	for k, nestedBlock := range nestedBlocks {
		blockType := nestedBlock.Type()
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(indent + blockType)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		
		// Recursively convert nested block with increased indentation
		nestedIndent := indent + "  "
		nestedTokens := blockToObjectTokensWithIndent(nestedBlock, nestedIndent)
		tokens = append(tokens, nestedTokens...)

		// Add comma if not the last block
		if k < len(nestedBlocks)-1 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	// Close with proper indentation (one level less than content)
	closeIndent := indent[:len(indent)-2] // Remove 2 spaces
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(closeIndent + "}")})
	return tokens
}
// convertNestedBlockToTokens converts nested blocks to appropriate format (map or list) based on cloudflare_ruleset rules
func convertNestedBlockToTokens(block *hclwrite.Block) hclwrite.Tokens {
	blockType := block.Type()
	
	// Determine transformation type based on cloudflare_ruleset configuration
	// Rules within overrides should be lists, most other blocks should be maps
	if blockType == "rules" {
		// This is a nested rules block (like in overrides.rules) - should be a list
		return convertNestedBlockToListTokens(block)
	}
	
	// Default: convert to map/object
	return convertNestedBlockToMapTokens(block)
}

// convertNestedBlockToMapTokens converts a block to object format
func convertNestedBlockToMapTokens(block *hclwrite.Block) hclwrite.Tokens {
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
	}

	// Handle attributes only - keep it simple to avoid structural issues
	attrs := block.Body().Attributes()
	attrNames := make([]string, 0, len(attrs))
	for name := range attrs {
		attrNames = append(attrNames, name)
	}
	sort.Strings(attrNames)

	for j, name := range attrNames {
		attr := attrs[name]
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      " + name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Preserve original expression tokens
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)

		if j < len(attrNames)-1 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    }")})
	return tokens
}

// convertNestedBlockToListTokens converts a block to list format (for cases like overrides.rules)
func convertNestedBlockToListTokens(block *hclwrite.Block) hclwrite.Tokens {
	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
	}

	// Handle attributes only for simplicity
	attrs := block.Body().Attributes()
	attrNames := make([]string, 0, len(attrs))
	for name := range attrs {
		attrNames = append(attrNames, name)
	}
	sort.Strings(attrNames)

	for j, name := range attrNames {
		attr := attrs[name]
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        " + name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Preserve original expression tokens
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)

		if j < len(attrNames)-1 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      }")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})
	return tokens
}
