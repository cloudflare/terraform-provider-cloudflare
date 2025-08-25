package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// IsCloudflareListResource checks if a block is a cloudflare_list resource
func IsCloudflareListResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_list"
}

// ListTransformResult holds the result of transforming a cloudflare_list block
type ListTransformResult struct {
	NewBlocks       []*hclwrite.Block
	HasDynamicItems bool
	ListName        string
	StaticItemCount int
}

// TransformCloudflareListBlock transforms a v4 cloudflare_list resource with items into:
// 1. A v5 cloudflare_list resource without items
// 2. Multiple cloudflare_list_item resources (one per item)
//
// Naming Convention:
// - Static items: Creates resources named "{list_name}_item_{index}" (e.g., "mylist_item_0", "mylist_item_1")
// - Dynamic items: Creates a single resource named "{list_name}_items" with for_each (e.g., "mylist_items")
//
// State Migration Note:
// The state migration (in list_state.go) DOES create cloudflare_list_item resources.
// It fetches real item IDs from the Cloudflare API and creates separate state entries
// for each item. The migration:
// 1. Removes items from cloudflare_list resources
// 2. Creates new cloudflare_list_item resources with actual IDs from the API
// 3. Preserves timestamps and metadata from the API responses
//
// After migration, users should use moved blocks (see list_moved.go) for zero-downtime
// migration to avoid recreation of existing resources.
func TransformCloudflareListBlock(oldBlock *hclwrite.Block) []*hclwrite.Block {
	var newBlocks []*hclwrite.Block
	diags := ast.Diagnostics{}

	// Get the resource name from the old block
	resourceName := oldBlock.Labels()[1]

	// Create the new cloudflare_list resource without items
	newListBlock := hclwrite.NewBlock("resource", []string{"cloudflare_list", resourceName})

	// Copy over non-item attributes
	copyNonItemAttributes(oldBlock, newListBlock)

	// Get account_id and kind for use in list_item resources
	accountIDAttr := oldBlock.Body().GetAttribute("account_id")
	listKind := extractListKind(oldBlock)

	newBlocks = append(newBlocks, newListBlock)

	// Process both static item blocks and dynamic blocks
	// First collect all items with their sort keys
	type itemWithKey struct {
		block     *hclwrite.Block
		sortKey   string
		isDynamic bool
	}
	var items []itemWithKey

	// Collect static item blocks with their sort keys
	for _, block := range oldBlock.Body().Blocks() {
		if block.Type() == "item" {
			sortKey := extractItemSortKey(block, listKind)
			items = append(items, itemWithKey{
				block:     block,
				sortKey:   sortKey,
				isDynamic: false,
			})
		}
	}

	// Sort items by their sort keys for deterministic ordering
	sort.Slice(items, func(i, j int) bool {
		return items[i].sortKey < items[j].sortKey
	})

	// Now process sorted static items
	itemIndex := 0
	for _, item := range items {
		if !item.isDynamic {
			itemResourceName := fmt.Sprintf("%s_item_%d", resourceName, itemIndex)

			// Create the cloudflare_list_item resource using AST helpers
			newItemBlock := createListItemResourceWithAST(
				itemResourceName,
				resourceName,
				accountIDAttr,
				item.block,
				listKind,
				&diags,
			)

			if newItemBlock != nil {
				newBlocks = append(newBlocks, newItemBlock)
			}
			itemIndex++
		}
	}

	// Second, process dynamic blocks
	for _, block := range oldBlock.Body().Blocks() {
		if block.Type() == "dynamic" && len(block.Labels()) > 0 && block.Labels()[0] == "item" {
			// Handle dynamic item blocks
			dynamicItems := processDynamicItemBlock(block, resourceName, accountIDAttr, listKind, &itemIndex, &diags)
			newBlocks = append(newBlocks, dynamicItems...)
		}
	}

	// Log any diagnostics
	if len(diags.ComplicatedHCL) > 0 {
		// In production, we'd handle these diagnostics properly
		// For now, we'll just continue as the test expects
	}

	return newBlocks
}

// copyNonItemAttributes copies all non-item attributes from old block to new block
func copyNonItemAttributes(oldBlock, newBlock *hclwrite.Block) {
	for name, attr := range oldBlock.Body().Attributes() {
		// Skip any item-related attributes (shouldn't be any at this level)
		if name != "item" {
			tokens := attr.Expr().BuildTokens(nil)
			newBlock.Body().SetAttributeRaw(name, tokens)
		}
	}
}

// extractListKind extracts the kind value from the list block
func extractListKind(block *hclwrite.Block) string {
	if kindAttr := block.Body().GetAttribute("kind"); kindAttr != nil {
		// Try to extract the kind value if it's a literal string
		kindTokens := kindAttr.Expr().BuildTokens(nil)
		for _, token := range kindTokens {
			if token.Type == hclsyntax.TokenQuotedLit {
				return strings.Trim(string(token.Bytes), "\"")
			}
		}
	}
	return ""
}

// createListItemResourceWithAST creates a new cloudflare_list_item resource using AST helpers
func createListItemResourceWithAST(itemName, listName string, accountIDAttr *hclwrite.Attribute, itemBlock *hclwrite.Block, listKind string, diags *ast.Diagnostics) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_list_item", itemName})
	body := block.Body()

	// Set account_id
	if accountIDAttr != nil {
		tokens := accountIDAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("account_id", tokens)
	}

	// Set list_id reference using helper function
	listIDTokens := buildResourceAttributeReference("cloudflare_list", listName, "id")
	body.SetAttributeRaw("list_id", listIDTokens)

	// Set comment if present
	if commentAttr := itemBlock.Body().GetAttribute("comment"); commentAttr != nil {
		tokens := commentAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("comment", tokens)
	}

	// Process value block
	for _, valueBlock := range itemBlock.Body().Blocks() {
		if valueBlock.Type() == "value" {
			processValueBlock(body, valueBlock, listKind, diags)
		}
	}

	return block
}

// processValueBlock processes the value block within an item
func processValueBlock(body *hclwrite.Body, valueBlock *hclwrite.Block, listKind string, diags *ast.Diagnostics) {
	switch listKind {
	case "ip":
		if ipAttr := valueBlock.Body().GetAttribute("ip"); ipAttr != nil {
			tokens := ipAttr.Expr().BuildTokens(nil)
			body.SetAttributeRaw("ip", tokens)
		}
	case "asn":
		if asnAttr := valueBlock.Body().GetAttribute("asn"); asnAttr != nil {
			// ASN values should be numbers, not strings
			// Parse the value and create number tokens
			tokens := asnAttr.Expr().BuildTokens(nil)
			asnTokens := parseASNAsNumber(tokens)
			body.SetAttributeRaw("asn", asnTokens)
		}
	case "hostname":
		// hostname block becomes an object attribute
		for _, hostnameBlock := range valueBlock.Body().Blocks() {
			if hostnameBlock.Type() == "hostname" {
				hostnameTokens := buildObjectFromBlockForDynamics(hostnameBlock, "", diags)
				body.SetAttributeRaw("hostname", hostnameTokens)
			}
		}
	case "redirect":
		// redirect block becomes an object attribute with boolean conversions
		for _, redirectBlock := range valueBlock.Body().Blocks() {
			if redirectBlock.Type() == "redirect" {
				redirectTokens := buildRedirectObject(redirectBlock, "", diags)
				body.SetAttributeRaw("redirect", redirectTokens)
			}
		}
	}
}

// buildObjectFromBlock creates object tokens from a block using AST helpers
// If iteratorName is provided, it replaces dynamic iterator references in the value tokens
func buildObjectFromBlockForDynamics(block *hclwrite.Block, iteratorName string, diags *ast.Diagnostics) hclwrite.Tokens {
	var attrs []hclwrite.ObjectAttrTokens

	// Get attributes in order
	orderedAttrs := AttributesOrdered(block.Body())
	for _, attrInfo := range orderedAttrs {
		nameTokens := hclwrite.TokensForIdentifier(attrInfo.Name)
		valueTokens := attrInfo.Attribute.Expr().BuildTokens(nil)

		// Replace iterator references if iteratorName is provided (for dynamic blocks)
		if iteratorName != "" {
			valueTokens = replaceDynamicIteratorTokens(valueTokens, iteratorName)
		}

		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  nameTokens,
			Value: valueTokens,
		})
	}

	return hclwrite.TokensForObject(attrs)
}

// buildRedirectObject creates redirect object with boolean conversions
// If iteratorName is provided, it replaces dynamic iterator references but skips boolean conversion
func buildRedirectObject(block *hclwrite.Block, iteratorName string, diags *ast.Diagnostics) hclwrite.Tokens {
	var attrs []hclwrite.ObjectAttrTokens

	// Define boolean conversion fields
	booleanFields := map[string]bool{
		"include_subdomains":    true,
		"subpath_matching":      true,
		"preserve_query_string": true,
		"preserve_path_suffix":  true,
	}

	// Get attributes in order
	orderedAttrs := AttributesOrdered(block.Body())
	for _, attrInfo := range orderedAttrs {
		nameTokens := hclwrite.TokensForIdentifier(attrInfo.Name)
		valueTokens := attrInfo.Attribute.Expr().BuildTokens(nil)

		// Replace iterator references if iteratorName is provided (for dynamic blocks)
		if iteratorName != "" {
			valueTokens = replaceDynamicIteratorTokens(valueTokens, iteratorName)
		}

		// Check if this field needs boolean conversion (only for static blocks)
		if _, needsConversion := booleanFields[attrInfo.Name]; needsConversion && iteratorName == "" {
			// For static blocks, convert strings to booleans
			valueTokens = convertStringToBooleanTokens(attrInfo.Attribute.Expr())
		}

		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  nameTokens,
			Value: valueTokens,
		})
	}

	return hclwrite.TokensForObject(attrs)
}

// convertStringToBooleanTokens converts "enabled"/"disabled" strings to boolean tokens
func convertStringToBooleanTokens(expr *hclwrite.Expression) hclwrite.Tokens {
	exprTokens := expr.BuildTokens(nil)

	// Check for a simple quoted string literal (3 tokens: open quote, literal, close quote)
	if len(exprTokens) == 3 &&
		exprTokens[0].Type == hclsyntax.TokenOQuote &&
		exprTokens[1].Type == hclsyntax.TokenQuotedLit &&
		exprTokens[2].Type == hclsyntax.TokenCQuote {

		strVal := string(exprTokens[1].Bytes)
		if strVal == "enabled" {
			return hclwrite.TokensForValue(cty.BoolVal(true))
		} else if strVal == "disabled" {
			return hclwrite.TokensForValue(cty.BoolVal(false))
		}
	}

	// Not a simple literal string or not "enabled"/"disabled", keep original expression
	return exprTokens
}

// processDynamicItemBlock handles dynamic "item" blocks and creates for_each list_item resources
func processDynamicItemBlock(dynamicBlock *hclwrite.Block, listName string, accountIDAttr *hclwrite.Attribute, listKind string, itemIndex *int, diags *ast.Diagnostics) []*hclwrite.Block {
	var blocks []*hclwrite.Block

	// Extract for_each expression
	forEachAttr := dynamicBlock.Body().GetAttribute("for_each")
	if forEachAttr == nil {
		// Dynamic block without for_each is invalid, add warning
		return blocks
	}

	// Extract iterator name (defaults to label if not specified)
	// The default iterator name is the label of the dynamic block ("item" in this case)
	iteratorName := "item"
	if iteratorAttr := dynamicBlock.Body().GetAttribute("iterator"); iteratorAttr != nil {
		// Parse iterator name from the attribute
		// The iterator attribute contains just the name as an identifier
		tokens := iteratorAttr.Expr().BuildTokens(nil)
		for _, token := range tokens {
			if token.Type == hclsyntax.TokenIdent {
				iteratorName = string(token.Bytes)
				break
			}
		}
	}

	// Find the content block
	var contentBlock *hclwrite.Block
	for _, block := range dynamicBlock.Body().Blocks() {
		if block.Type() == "content" {
			contentBlock = block
			break
		}
	}

	if contentBlock == nil {
		return blocks
	}

	// Check if the content block suggests we're dealing with objects
	// by looking for iterator.value.field patterns
	isObjectIteration := checkIfObjectIteration(contentBlock, iteratorName)

	// Create a for_each cloudflare_list_item resource
	itemResourceName := fmt.Sprintf("%s_items", listName)
	newItemBlock := hclwrite.NewBlock("resource", []string{"cloudflare_list_item", itemResourceName})
	body := newItemBlock.Body()

	// Add for_each attribute - handle differently based on whether it's objects or simple values
	var forEachTokens hclwrite.Tokens
	if isObjectIteration {
		// For objects, create a for expression with the appropriate key field based on list kind
		forEachTokens = wrapForEachForObjects(forEachAttr.Expr(), listKind)
	} else {
		// For simple values, wrap with toset()
		forEachTokens = wrapForEachWithToSet(forEachAttr.Expr())
	}
	body.SetAttributeRaw("for_each", forEachTokens)

	// Set account_id
	if accountIDAttr != nil {
		tokens := accountIDAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("account_id", tokens)
	}

	// Set list_id reference
	listIDTraversal := hcl.Traversal{
		hcl.TraverseRoot{Name: "cloudflare_list"},
		hcl.TraverseAttr{Name: listName},
		hcl.TraverseAttr{Name: "id"},
	}
	body.SetAttributeTraversal("list_id", listIDTraversal)

	// Process the content block to extract item values
	processDynamicContentBlock(body, contentBlock, iteratorName, listKind, diags)

	blocks = append(blocks, newItemBlock)
	return blocks
}

// processDynamicContentBlock processes the content block of a dynamic item
func processDynamicContentBlock(body *hclwrite.Body, contentBlock *hclwrite.Block, iteratorName string, listKind string, diags *ast.Diagnostics) {
	// Handle comment if present
	if commentAttr := contentBlock.Body().GetAttribute("comment"); commentAttr != nil {
		tokens := commentAttr.Expr().BuildTokens(nil)
		// Replace iterator references with each.value or each.key
		tokens = replaceDynamicIteratorTokens(tokens, iteratorName)
		body.SetAttributeRaw("comment", tokens)
	}

	// Process value block
	for _, valueBlock := range contentBlock.Body().Blocks() {
		if valueBlock.Type() == "value" {
			processDynamicValueBlock(body, valueBlock, iteratorName, listKind, diags)
		}
	}
}

// processDynamicValueBlock processes the value block within a dynamic item's content
func processDynamicValueBlock(body *hclwrite.Body, valueBlock *hclwrite.Block, iteratorName string, listKind string, diags *ast.Diagnostics) {
	switch listKind {
	case "ip":
		if ipAttr := valueBlock.Body().GetAttribute("ip"); ipAttr != nil {
			tokens := ipAttr.Expr().BuildTokens(nil)
			tokens = replaceDynamicIteratorTokens(tokens, iteratorName)
			body.SetAttributeRaw("ip", tokens)
		}
	case "asn":
		if asnAttr := valueBlock.Body().GetAttribute("asn"); asnAttr != nil {
			tokens := asnAttr.Expr().BuildTokens(nil)
			tokens = replaceDynamicIteratorTokens(tokens, iteratorName)
			// ASN values should be numbers, not strings
			asnTokens := parseASNAsNumber(tokens)
			body.SetAttributeRaw("asn", asnTokens)
		}
	case "hostname":
		// hostname block becomes an object attribute
		for _, hostnameBlock := range valueBlock.Body().Blocks() {
			if hostnameBlock.Type() == "hostname" {
				hostnameTokens := buildObjectFromBlockForDynamics(hostnameBlock, iteratorName, diags)
				body.SetAttributeRaw("hostname", hostnameTokens)
			}
		}
	case "redirect":
		// redirect block becomes an object attribute with boolean conversions
		for _, redirectBlock := range valueBlock.Body().Blocks() {
			if redirectBlock.Type() == "redirect" {
				redirectTokens := buildRedirectObject(redirectBlock, iteratorName, diags)
				body.SetAttributeRaw("redirect", redirectTokens)
			}
		}
	}
}

// replaceDynamicIteratorTokens replaces iterator references with each.value or each.key
// Handles patterns like:
// - iterator.value -> each.value
// - iterator.key -> each.key
// - iterator.value.property -> each.value.property (for objects)
func replaceDynamicIteratorTokens(tokens hclwrite.Tokens, iteratorName string) hclwrite.Tokens {
	var result hclwrite.Tokens

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Look for the iterator name
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == iteratorName {
			// Check if next token is a dot
			if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenDot {
				// Check what comes after the dot
				if i+2 < len(tokens) && tokens[i+2].Type == hclsyntax.TokenIdent {
					// Replace iterator with "each"
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("each"),
					})

					// If it's .value or .key, keep them as is
					// This handles both simple cases (iterator.value) and
					// complex cases (iterator.value.property)
					continue
				}
			}
		}

		// Keep the original token
		result = append(result, token)
	}

	return result
}

// buildResourceAttributeReference builds a reference like cloudflare_list.name.id
func buildResourceAttributeReference(resourceType, resourceName, attribute string) hclwrite.Tokens {
	traversal := hcl.Traversal{
		hcl.TraverseRoot{Name: resourceType},
		hcl.TraverseAttr{Name: resourceName},
		hcl.TraverseAttr{Name: attribute},
	}
	return hclwrite.TokensForTraversal(traversal)
}

// checkIfObjectIteration checks if the content block suggests we're iterating over objects
// by looking for patterns like iterator.value.field
func checkIfObjectIteration(contentBlock *hclwrite.Block, iteratorName string) bool {
	// Check all attributes in the content block
	for _, attr := range contentBlock.Body().Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		// Look for pattern: iteratorName.value.something
		for i := 0; i < len(tokens)-4; i++ {
			if tokens[i].Type == hclsyntax.TokenIdent && string(tokens[i].Bytes) == iteratorName &&
				i+4 < len(tokens) &&
				tokens[i+1].Type == hclsyntax.TokenDot &&
				tokens[i+2].Type == hclsyntax.TokenIdent && string(tokens[i+2].Bytes) == "value" &&
				tokens[i+3].Type == hclsyntax.TokenDot &&
				tokens[i+4].Type == hclsyntax.TokenIdent {
				// Found pattern like iterator.value.field
				return true
			}
		}
	}

	// Also check value blocks for nested attributes
	for _, block := range contentBlock.Body().Blocks() {
		if block.Type() == "value" {
			for _, attr := range block.Body().Attributes() {
				tokens := attr.Expr().BuildTokens(nil)
				// Look for pattern: iteratorName.value.something
				for i := 0; i < len(tokens)-4; i++ {
					if tokens[i].Type == hclsyntax.TokenIdent && string(tokens[i].Bytes) == iteratorName &&
						i+4 < len(tokens) &&
						tokens[i+1].Type == hclsyntax.TokenDot &&
						tokens[i+2].Type == hclsyntax.TokenIdent && string(tokens[i+2].Bytes) == "value" &&
						tokens[i+3].Type == hclsyntax.TokenDot &&
						tokens[i+4].Type == hclsyntax.TokenIdent {
						// Found pattern like iterator.value.field
						return true
					}
				}
			}
		}
	}

	return false
}

// extractItemSortKey extracts a sort key from an item block for deterministic ordering
func extractItemSortKey(block *hclwrite.Block, listKind string) string {
	// Try to extract the primary value from the item block
	// This ensures consistent ordering between config and state migrations

	// Look for value block
	for _, valueBlock := range block.Body().Blocks() {
		if valueBlock.Type() == "value" {
			switch listKind {
			case "ip":
				if ipAttr := valueBlock.Body().GetAttribute("ip"); ipAttr != nil {
					// Extract the IP value as a string for sorting
					tokens := ipAttr.Expr().BuildTokens(nil)
					return tokensToSortString(tokens)
				}
			case "asn":
				if asnAttr := valueBlock.Body().GetAttribute("asn"); asnAttr != nil {
					// Extract the ASN value as a string for sorting
					tokens := asnAttr.Expr().BuildTokens(nil)
					// Pad with zeros for numeric sorting
					asnStr := tokensToSortString(tokens)
					// Remove quotes if present
					asnStr = strings.Trim(asnStr, "\"")
					// Pad to ensure numeric sorting works correctly
					return fmt.Sprintf("%020s", asnStr)
				}
			case "hostname":
				// For hostname, look for the url_hostname attribute
				for _, hostnameBlock := range valueBlock.Body().Blocks() {
					if hostnameBlock.Type() == "hostname" {
						if urlAttr := hostnameBlock.Body().GetAttribute("url_hostname"); urlAttr != nil {
							tokens := urlAttr.Expr().BuildTokens(nil)
							return tokensToSortString(tokens)
						}
					}
				}
			case "redirect":
				// For redirect, use source_url as the sort key
				for _, redirectBlock := range valueBlock.Body().Blocks() {
					if redirectBlock.Type() == "redirect" {
						if srcAttr := redirectBlock.Body().GetAttribute("source_url"); srcAttr != nil {
							tokens := srcAttr.Expr().BuildTokens(nil)
							return tokensToSortString(tokens)
						}
					}
				}
			}
		}
	}

	// If we can't extract a value, use comment as fallback
	if commentAttr := block.Body().GetAttribute("comment"); commentAttr != nil {
		tokens := commentAttr.Expr().BuildTokens(nil)
		return "zzz_" + tokensToSortString(tokens) // Prefix to sort comments last
	}

	return "zzz_unknown" // Default for items without identifiable values
}

// tokensToSortString converts HCL tokens to a string for sorting
func tokensToSortString(tokens hclwrite.Tokens) string {
	var result strings.Builder
	for _, tok := range tokens {
		// Skip quotes and whitespace for cleaner comparison
		if tok.Type != hclsyntax.TokenOQuote &&
			tok.Type != hclsyntax.TokenCQuote &&
			tok.Type != hclsyntax.TokenNewline {
			result.Write(tok.Bytes)
		}
	}
	return strings.TrimSpace(result.String())
}

// parseASNAsNumber converts ASN string tokens to number tokens
// This ensures ASN values are output as numbers (9999) instead of strings ("9999")
func parseASNAsNumber(tokens hclwrite.Tokens) hclwrite.Tokens {
	// Check if it's a quoted string
	if len(tokens) >= 3 {
		hasOpenQuote := false
		hasCloseQuote := false
		var numberValue string

		for _, tok := range tokens {
			if tok.Type == hclsyntax.TokenOQuote {
				hasOpenQuote = true
			} else if tok.Type == hclsyntax.TokenCQuote {
				hasCloseQuote = true
			} else if tok.Type == hclsyntax.TokenQuotedLit || tok.Type == hclsyntax.TokenIdent {
				numberValue = string(tok.Bytes)
			}
		}

		// If it's a quoted string, return just the number
		if hasOpenQuote && hasCloseQuote && numberValue != "" {
			return hclwrite.Tokens{
				&hclwrite.Token{
					Type:  hclsyntax.TokenNumberLit,
					Bytes: []byte(numberValue),
				},
			}
		}
	}

	// Return original tokens if not a quoted string
	return tokens
}

// wrapForEachForObjects creates a for expression for lists of objects
// The key field is determined by the list kind
func wrapForEachForObjects(expr *hclwrite.Expression, listKind string) hclwrite.Tokens {
	tokens := expr.BuildTokens(nil)

	// Determine the key field based on list kind
	var keyField string
	switch listKind {
	case "ip":
		keyField = "ip"
	case "asn":
		keyField = "asn"
	case "hostname":
		keyField = "url_hostname"
	case "redirect":
		keyField = "source_url"
	default:
		// Fallback to index-based approach if unknown kind
		return buildForExpressionWithIndex(tokens)
	}

	// Build: { for item in <expr> : item.<keyField> => item }
	var result hclwrite.Tokens

	// Opening brace
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenOBrace,
		Bytes: []byte("{"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// "for item in"
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("for"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("item"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("in"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// Add the original list expression
	result = append(result, tokens...)

	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenColon,
		Bytes: []byte(":"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// "item.<keyField> => item"
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("item"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenDot,
		Bytes: []byte("."),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte(keyField),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenFatArrow,
		Bytes: []byte("=>"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("item"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// Closing brace
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrace,
		Bytes: []byte("}"),
	})

	return result
}

// buildForExpressionWithIndex creates a for expression that uses index as the key
// Used as a fallback when we can't determine the appropriate key field
func buildForExpressionWithIndex(listTokens hclwrite.Tokens) hclwrite.Tokens {
	var result hclwrite.Tokens

	// Opening brace
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenOBrace,
		Bytes: []byte("{"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// "for idx, item in"
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("for"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("idx"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenComma,
		Bytes: []byte(","),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("item"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("in"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// Add the original list expression
	result = append(result, listTokens...)

	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenColon,
		Bytes: []byte(":"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// "tostring(idx) => item"
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("tostring"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenOParen,
		Bytes: []byte("("),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("idx"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenCParen,
		Bytes: []byte(")"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenFatArrow,
		Bytes: []byte("=>"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("item"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte(" "),
	})

	// Closing brace
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrace,
		Bytes: []byte("}"),
	})

	return result
}

// wrapForEachWithToSet wraps the for_each expression with toset() if needed
// For lists of objects, it creates a proper for expression instead
func wrapForEachWithToSet(expr *hclwrite.Expression) hclwrite.Tokens {
	tokens := expr.BuildTokens(nil)

	// Check if the expression already starts with toset(
	hasToSet := false
	hasForExpr := false
	for i, token := range tokens {
		if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == "toset" {
			// Check if it's followed by an opening parenthesis
			if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenOParen {
				hasToSet = true
				break
			}
		}
		// Check if it's already a for expression
		if token.Type == hclsyntax.TokenOBrace {
			hasForExpr = true
			break
		}
	}

	// If it already has toset() or is a for expression, return as-is
	if hasToSet || hasForExpr {
		return tokens
	}

	// For now, wrap with toset() - this works for lists of strings
	// For lists of objects, users will need to manually adjust
	// TODO: Detect list of objects and handle differently
	var result hclwrite.Tokens
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenIdent,
		Bytes: []byte("toset"),
	})
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenOParen,
		Bytes: []byte("("),
	})
	result = append(result, tokens...)
	result = append(result, &hclwrite.Token{
		Type:  hclsyntax.TokenCParen,
		Bytes: []byte(")"),
	})

	return result
}
