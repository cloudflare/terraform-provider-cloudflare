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
		if err := transformBlocksToMap(config, body, blockName, resourceType); err != nil {
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

	// Use token-based approach for ALL resources to preserve expressions properly
	// This ensures variable references, function calls, and interpolations are preserved
	return transformBlocksToListPreservingTokens(body, blocks, blockName)
}

// transformBlocksToListPreservingTokens transforms blocks to a list while preserving expressions/heredocs
func transformBlocksToListPreservingTokens(body *hclwrite.Body, blocks []*hclwrite.Block, blockName string) error {
	// Build tokens for the list directly to preserve all expressions
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

		// Process attributes
		attrs := block.Body().Attributes()
		attrNames := make([]string, 0, len(attrs))
		for name := range attrs {
			attrNames = append(attrNames, name)
		}
		sort.Strings(attrNames)

		// Process nested blocks to determine if we need commas after attributes
		nestedBlocks := block.Body().Blocks()
		
		for j, name := range attrNames {
			attr := attrs[name]
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + name)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Preserve the original expression tokens
			exprTokens := attr.Expr().BuildTokens(nil)
			tokens = append(tokens, exprTokens...)

			// Add comma after each attribute except the last one, OR if there are nested blocks coming
			if j < len(attrNames)-1 || len(nestedBlocks) > 0 {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Process nested blocks
		for k, nestedBlock := range nestedBlocks {
			// Add comma after previous nested block (but not before the first one)
			if k > 0 {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			}

			// Add nested block name as attribute
			nestedName := nestedBlock.Type()
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    " + nestedName)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

			// Convert nested block to object tokens - use a cleaner format
			nestedTokens := blockToMapTokens(nestedBlock)
			tokens = append(tokens, nestedTokens...)
		}

		// Close object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})

	// Set the attribute using raw tokens
	body.SetAttributeRaw(blockName, tokens)

	// Remove original blocks
	for _, block := range blocks {
		body.RemoveBlock(block)
	}

	return nil
}

// blockToMapTokens converts a block to a map format suitable for nested objects
func blockToMapTokens(block *hclwrite.Block) hclwrite.Tokens {
	blockBody := block.Body()
	attrs := blockBody.Attributes()

	tokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
	}

	if len(attrs) == 0 {
		// Empty block
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
		return tokens
	}

	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

	// Sort attribute names
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	// Process attributes
	for i, name := range names {
		attr := attrs[name]

		// Add indentation (6 spaces for nested content in list items)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})

		// Add the attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Add the attribute expression tokens
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)

		// Add comma after each attribute except the last
		if i < len(names)-1 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	// Close the object
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})

	return tokens
}

// blockToTokensPreserving converts a block to tokens while preserving all expressions
func blockToTokensPreserving(block *hclwrite.Block) hclwrite.Tokens {
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

	// Sort attribute names
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
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}
		first = false

		// Add indentation (6 spaces for nested content)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})

		// Add the attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Add the attribute expression tokens (preserves heredocs, interpolations, etc)
		exprTokens := attr.Expr().BuildTokens(nil)
		tokens = append(tokens, exprTokens...)
	}

	// Process nested blocks
	for i, nestedBlock := range blocks {
		// Only add comma if there were attributes before OR this is not the first block
		if len(names) > 0 || i > 0 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Add indentation
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})

		// Add the nested block name
		nestedName := nestedBlock.Type()
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedName)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})

		// Recursively convert the nested block (use blockToMapTokens for proper formatting)
		nestedTokens := blockToMapTokens(nestedBlock)
		tokens = append(tokens, nestedTokens...)
	}

	tokens = append(tokens,
		&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")},
		&hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
	)

	return tokens
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

// transformBlocksToMap transforms a single block into a map attribute
func transformBlocksToMap(config *TransformationConfig, body *hclwrite.Body, blockName string, resourceType string) error {
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
		// String value - remove only the outer quotes and unescape inner quotes
		if len(str) >= 2 {
			inner := str[1 : len(str)-1]
			// Unescape quotes to prevent double-escaping when HCL formats them
			inner = strings.ReplaceAll(inner, `\"`, `"`)
			return cty.StringVal(inner)
		}
		return cty.StringVal("")
	}
	// Try as number
	if num, err := strconv.ParseInt(str, 10, 64); err == nil {
		return cty.NumberIntVal(num)
	}
	if num, err := strconv.ParseFloat(str, 64); err == nil {
		return cty.NumberFloatVal(num)
	}
	
	// Check if this is a resource reference (e.g., resource_type.name.attribute)
	// Resource references should not be converted to strings as they need to remain as references
	if isResourceReference(str) {
		// Return as a dynamic value that will be preserved as-is
		return cty.DynamicVal
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
// isResourceReference detects if a string represents a Terraform resource reference
func isResourceReference(str string) bool {
	// Resource references have the pattern: resource_type.resource_name.attribute
	// or data.data_type.data_name.attribute
	// They contain dots and don't start with quotes
	if strings.HasPrefix(str, "\"") || strings.HasSuffix(str, "\"") {
		return false // Quoted strings are not resource references
	}
	
	// Must contain at least two dots (resource_type.name.attribute)
	parts := strings.Split(str, ".")
	if len(parts) < 3 {
		return false
	}
	
	// First part should be a resource type (starts with letter, contains only alphanumeric and underscores)
	firstPart := parts[0]
	if len(firstPart) == 0 {
		return false
	}
	
	// Check if it looks like a resource type or data source
	if firstPart == "data" || 
	   (firstPart[0] >= 'a' && firstPart[0] <= 'z' && 
	    strings.Contains(firstPart, "_")) {
		return true
	}
	
	return false
}
