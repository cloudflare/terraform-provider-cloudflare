package main

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// mergeListItemResources finds all cloudflare_list_item resources and merges them
// into their parent cloudflare_list resources as items arrays
func mergeListItemResources(blocks []*hclwrite.Block) (blocksToRemove []*hclwrite.Block) {
	diags := ast.NewDiagnostics()
	
	// Step 1: Find all cloudflare_list and cloudflare_list_item resources
	listResources := make(map[string]*hclwrite.Block)
	listItemsByParent := make(map[string][]*hclwrite.Block)
	
	for _, block := range blocks {
		if isCloudflareListResource(block) {
			listName := block.Labels()[1]
			listResources[listName] = block
		} else if isCloudflareListItemResource(block) {
			// Extract the parent list reference
			parentList := extractParentListName(block)
			if parentList != "" {
				listItemsByParent[parentList] = append(listItemsByParent[parentList], block)
			}
		}
	}
	
	// Step 2: For each list that has associated items, merge them
	for listName, listBlock := range listResources {
		items := listItemsByParent[listName]
		if len(items) == 0 {
			continue // No items to merge for this list
		}
		
		// Get the list kind to properly structure items
		kind := extractListKind(listBlock)
		if kind == "" {
			addMigrationWarning(listBlock.Body(), 
				"Cannot determine list kind for merging list_item resources")
			continue
		}
		
		// Check if there's a single for_each or count pattern
		if isSingleDynamicPattern(items) {
			// Handle dynamic pattern (for_each or count)
			itemsExpr := createDynamicItemsExpression(items[0], kind, diags)
			if itemsExpr != nil {
				setItemsAttribute(listBlock.Body(), itemsExpr)
				blocksToRemove = append(blocksToRemove, items...)
			}
		} else {
			// Handle multiple static items
			itemsExpr := createStaticItemsExpression(items, kind, diags)
			if itemsExpr != nil {
				setItemsAttribute(listBlock.Body(), itemsExpr)
				blocksToRemove = append(blocksToRemove, items...)
			}
		}
	}
	
	return blocksToRemove
}

// isCloudflareListItemResource checks if a block is a cloudflare_list_item resource
func isCloudflareListItemResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_list_item"
}

// extractParentListName extracts the parent list name from a list_id reference
func extractParentListName(itemBlock *hclwrite.Block) string {
	body := itemBlock.Body()
	listIdAttr := body.GetAttribute("list_id")
	if listIdAttr == nil {
		return ""
	}
	
	// Parse the list_id expression to find the parent list reference
	tokens := listIdAttr.Expr().BuildTokens(nil)
	tokenStr := string(tokens.Bytes())
	
	// Handle common patterns:
	// - cloudflare_list.example.id
	// - cloudflare_list["example"].id
	if strings.Contains(tokenStr, "cloudflare_list.") {
		// Extract list name from reference
		parts := strings.Split(tokenStr, ".")
		if len(parts) >= 3 {
			listName := parts[1]
			// Remove .id suffix if present
			if strings.HasSuffix(listName, ".id") {
				listName = strings.TrimSuffix(listName, ".id")
			}
			return listName
		}
	} else if strings.Contains(tokenStr, `cloudflare_list["`) {
		// Handle bracket notation
		start := strings.Index(tokenStr, `["`) + 2
		end := strings.Index(tokenStr[start:], `"`)
		if end > 0 {
			return tokenStr[start : start+end]
		}
	}
	
	// If we can't determine the parent, return empty
	return ""
}

// extractListKind extracts the kind attribute from a cloudflare_list resource
func extractListKind(listBlock *hclwrite.Block) string {
	kindAttr := listBlock.Body().GetAttribute("kind")
	if kindAttr == nil {
		return ""
	}
	return extractStringValue(*kindAttr.Expr())
}

// isSingleDynamicPattern checks if all items use the same for_each or count
func isSingleDynamicPattern(items []*hclwrite.Block) bool {
	if len(items) != 1 {
		return false
	}
	
	body := items[0].Body()
	return body.GetAttribute("for_each") != nil || body.GetAttribute("count") != nil
}

// createDynamicItemsExpression creates a for expression from a dynamic list_item
func createDynamicItemsExpression(itemBlock *hclwrite.Block, kind string, diags ast.Diagnostics) hclsyntax.Expression {
	body := itemBlock.Body()
	
	// Check for for_each
	if forEachAttr := body.GetAttribute("for_each"); forEachAttr != nil {
		return createForEachItemsExpression(itemBlock, kind, forEachAttr, diags)
	}
	
	// Check for count
	if countAttr := body.GetAttribute("count"); countAttr != nil {
		return createCountItemsExpression(itemBlock, kind, countAttr, diags)
	}
	
	return nil
}

// createForEachItemsExpression creates items array from for_each pattern
func createForEachItemsExpression(itemBlock *hclwrite.Block, kind string, forEachAttr *hclwrite.Attribute, diags ast.Diagnostics) hclsyntax.Expression {
	body := itemBlock.Body()
	
	// Parse the for_each expression
	forEachBytes := forEachAttr.Expr().BuildTokens(nil).Bytes()
	forEachExpr, d := hclsyntax.ParseExpression(forEachBytes, "for_each", hcl.InitialPos)
	if d.HasErrors() {
		diags.HclDiagnostics.Extend(d)
		return nil
	}
	
	// Build the item object based on kind
	itemObj := buildItemObject(body, kind, "each", diags)
	if itemObj == nil {
		return nil
	}
	
	// Create for expression: [ for k, v in <for_each> : <item> ]
	// We use a special marker to indicate this should be a list comprehension with two variables
	forExpr := &hclsyntax.ForExpr{
		KeyVar: "k",
		ValVar: "v", 
		CollExpr: forEachExpr,
		ValExpr: itemObj,
	}
	
	// Return as-is; the setItemsAttribute will handle the special formatting
	return forExpr
}

// createCountItemsExpression creates items array from count pattern
func createCountItemsExpression(itemBlock *hclwrite.Block, kind string, countAttr *hclwrite.Attribute, diags ast.Diagnostics) hclsyntax.Expression {
	body := itemBlock.Body()
	
	// Parse the count expression
	countBytes := countAttr.Expr().BuildTokens(nil).Bytes()
	countExpr, d := hclsyntax.ParseExpression(countBytes, "count", hcl.InitialPos)
	if d.HasErrors() {
		diags.HclDiagnostics.Extend(d)
		return nil
	}
	
	// Build the item object, replacing count.index references
	itemObj := buildItemObjectForCount(body, kind, diags)
	if itemObj == nil {
		return nil
	}
	
	// Create range expression for count
	// [ for i in range(<count>) : <item with i replacing count.index> ]
	rangeCall := &hclsyntax.FunctionCallExpr{
		Name: "range",
		Args: []hclsyntax.Expression{countExpr},
	}
	
	return &hclsyntax.ForExpr{
		ValVar: "i",
		CollExpr: rangeCall,
		ValExpr: itemObj,
	}
}

// createStaticItemsExpression creates items array from multiple static list_item resources
func createStaticItemsExpression(items []*hclwrite.Block, kind string, diags ast.Diagnostics) hclsyntax.Expression {
	var itemExprs []hclsyntax.Expression
	
	for _, itemBlock := range items {
		body := itemBlock.Body()
		itemObj := buildItemObject(body, kind, "", diags)
		if itemObj != nil {
			itemExprs = append(itemExprs, itemObj)
		}
	}
	
	if len(itemExprs) == 0 {
		return nil
	}
	
	// Create tuple expression for array
	return &hclsyntax.TupleConsExpr{
		Exprs: itemExprs,
	}
}

// buildItemObject builds an object expression for a list item based on its kind
func buildItemObject(body *hclwrite.Body, kind string, iteratorPrefix string, diags ast.Diagnostics) hclsyntax.Expression {
	items := []hclsyntax.ObjectConsItem{}
	
	// Add comment if present
	if commentAttr := body.GetAttribute("comment"); commentAttr != nil {
		commentExpr := parseAndReplaceReferences(commentAttr, iteratorPrefix, diags)
		if commentExpr != nil {
			items = append(items, hclsyntax.ObjectConsItem{
				KeyExpr: ast.NewKeyExpr("comment"),
				ValueExpr: commentExpr,
			})
		}
	}
	
	// Add kind-specific fields
	switch kind {
	case "ip":
		if ipAttr := body.GetAttribute("ip"); ipAttr != nil {
			ipExpr := parseAndReplaceReferences(ipAttr, iteratorPrefix, diags)
			if ipExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("ip"),
					ValueExpr: ipExpr,
				})
			}
		}
		
	case "asn":
		if asnAttr := body.GetAttribute("asn"); asnAttr != nil {
			asnExpr := parseAndReplaceReferences(asnAttr, iteratorPrefix, diags)
			if asnExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("asn"),
					ValueExpr: asnExpr,
				})
			}
		}
		
	case "hostname":
		if hostnameAttr := body.GetAttribute("hostname"); hostnameAttr != nil {
			// hostname is already an object in list_item
			hostnameExpr := parseAndReplaceReferences(hostnameAttr, iteratorPrefix, diags)
			if hostnameExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("hostname"),
					ValueExpr: hostnameExpr,
				})
			}
		}
		
	case "redirect":
		if redirectAttr := body.GetAttribute("redirect"); redirectAttr != nil {
			// redirect is already an object in list_item
			redirectExpr := parseAndReplaceReferences(redirectAttr, iteratorPrefix, diags)
			if redirectExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("redirect"),
					ValueExpr: redirectExpr,
				})
			}
		}
	}
	
	if len(items) == 0 {
		return nil
	}
	
	return &hclsyntax.ObjectConsExpr{
		Items: items,
	}
}

// buildItemObjectForCount is similar to buildItemObject but replaces count.index with i
func buildItemObjectForCount(body *hclwrite.Body, kind string, diags ast.Diagnostics) hclsyntax.Expression {
	items := []hclsyntax.ObjectConsItem{}
	
	// Add comment if present
	if commentAttr := body.GetAttribute("comment"); commentAttr != nil {
		commentExpr := parseAndReplaceCountIndex(commentAttr, diags)
		if commentExpr != nil {
			items = append(items, hclsyntax.ObjectConsItem{
				KeyExpr: ast.NewKeyExpr("comment"),
				ValueExpr: commentExpr,
			})
		}
	}
	
	// Add kind-specific fields with count.index replacement
	switch kind {
	case "ip":
		if ipAttr := body.GetAttribute("ip"); ipAttr != nil {
			ipExpr := parseAndReplaceCountIndex(ipAttr, diags)
			if ipExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("ip"),
					ValueExpr: ipExpr,
				})
			}
		}
		
	case "asn":
		if asnAttr := body.GetAttribute("asn"); asnAttr != nil {
			asnExpr := parseAndReplaceCountIndex(asnAttr, diags)
			if asnExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("asn"),
					ValueExpr: asnExpr,
				})
			}
		}
		
	case "hostname":
		if hostnameAttr := body.GetAttribute("hostname"); hostnameAttr != nil {
			hostnameExpr := parseAndReplaceCountIndex(hostnameAttr, diags)
			if hostnameExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("hostname"),
					ValueExpr: hostnameExpr,
				})
			}
		}
		
	case "redirect":
		if redirectAttr := body.GetAttribute("redirect"); redirectAttr != nil {
			redirectExpr := parseAndReplaceCountIndex(redirectAttr, diags)
			if redirectExpr != nil {
				items = append(items, hclsyntax.ObjectConsItem{
					KeyExpr: ast.NewKeyExpr("redirect"),
					ValueExpr: redirectExpr,
				})
			}
		}
	}
	
	if len(items) == 0 {
		return nil
	}
	
	return &hclsyntax.ObjectConsExpr{
		Items: items,
	}
}

// parseAndReplaceReferences parses an attribute and replaces each.key/each.value references
func parseAndReplaceReferences(attr *hclwrite.Attribute, iteratorPrefix string, diags ast.Diagnostics) hclsyntax.Expression {
	bytes := attr.Expr().BuildTokens(nil).Bytes()
	exprStr := string(bytes)
	
	// If we have an iterator prefix (for_each case), replace references
	if iteratorPrefix == "each" {
		// Replace each.value with v and each.key with k
		exprStr = strings.ReplaceAll(exprStr, "each.value", "v")
		exprStr = strings.ReplaceAll(exprStr, "each.key", "k")
	}
	
	expr, d := hclsyntax.ParseExpression([]byte(exprStr), "attribute", hcl.InitialPos)
	if d.HasErrors() {
		diags.HclDiagnostics.Extend(d)
		return nil
	}
	return expr
}

// parseAndReplaceCountIndex parses an attribute and replaces count.index with i
func parseAndReplaceCountIndex(attr *hclwrite.Attribute, diags ast.Diagnostics) hclsyntax.Expression {
	bytes := attr.Expr().BuildTokens(nil).Bytes()
	exprStr := string(bytes)
	
	// Replace count.index with i
	exprStr = strings.ReplaceAll(exprStr, "count.index", "i")
	
	expr, d := hclsyntax.ParseExpression([]byte(exprStr), "attribute", hcl.InitialPos)
	if d.HasErrors() {
		diags.HclDiagnostics.Extend(d)
		return nil
	}
	return expr
}

// setItemsAttribute sets the items attribute on a list resource
func setItemsAttribute(body *hclwrite.Body, itemsExpr hclsyntax.Expression) {
	if itemsExpr == nil {
		return
	}
	
	// For simple cases, try to directly generate the HCL
	if forExpr, ok := itemsExpr.(*hclsyntax.ForExpr); ok {
		// Generate for expression manually
		var hclStr string
		if forExpr.KeyVar != "" && forExpr.KeyExpr == nil {
			// List comprehension with two iteration variables (for_each pattern)
			hclStr = fmt.Sprintf("[\n    for %s, %s in %s : %s\n  ]", 
				forExpr.KeyVar, forExpr.ValVar, 
				exprToSimpleString(forExpr.CollExpr),
				exprToSimpleString(forExpr.ValExpr))
		} else if forExpr.KeyVar != "" && forExpr.KeyExpr != nil {
			// Dict comprehension (with key => value)
			keyStr := exprToSimpleString(forExpr.KeyExpr)
			valStr := exprToSimpleString(forExpr.ValExpr)
			hclStr = fmt.Sprintf("{\n    for %s, %s in %s : %s => %s\n  }", 
				forExpr.KeyVar, forExpr.ValVar, 
				exprToSimpleString(forExpr.CollExpr),
				keyStr, valStr)
		} else {
			// List comprehension with single variable
			hclStr = fmt.Sprintf("[\n    for %s in %s : %s\n  ]", 
				forExpr.ValVar,
				exprToSimpleString(forExpr.CollExpr),
				exprToSimpleString(forExpr.ValExpr))
		}
		
		// Parse the generated HCL to get tokens
		file, diags := hclwrite.ParseConfig([]byte("attr = "+hclStr), "", hcl.InitialPos)
		if !diags.HasErrors() && file != nil && file.Body() != nil {
			if attr := file.Body().GetAttribute("attr"); attr != nil {
				body.SetAttributeRaw("items", attr.Expr().BuildTokens(nil))
				return
			}
		}
	}
	
	// Fall back to using Expr2WriteExpr
	diags := ast.NewDiagnostics()
	defer func() {
		if r := recover(); r != nil {
			// If conversion fails, add a warning
			addMigrationWarning(body, "Could not automatically merge list_item resources")
		}
	}()
	
	writeExpr := ast.Expr2WriteExpr(itemsExpr, diags)
	tokens := writeExpr.BuildTokens(nil)
	if len(tokens) > 0 {
		body.SetAttributeRaw("items", tokens)
	}
}

// exprToSimpleString converts a simple expression to a string representation
func exprToSimpleString(expr hclsyntax.Expression) string {
	switch e := expr.(type) {
	case *hclsyntax.ScopeTraversalExpr:
		// Variable reference like var.ip_addresses
		var parts []string
		for _, t := range e.Traversal {
			switch tt := t.(type) {
			case hcl.TraverseRoot:
				parts = append(parts, tt.Name)
			case hcl.TraverseAttr:
				parts = append(parts, tt.Name)
			case hcl.TraverseIndex:
				// Handle index access
				key, _ := tt.Key.AsBigFloat().Int64()
				parts = append(parts, fmt.Sprintf("[%d]", key))
			}
		}
		return strings.Join(parts, ".")
	case *hclsyntax.FunctionCallExpr:
		// Function call like range(length(var.ip_list))
		args := []string{}
		for _, arg := range e.Args {
			args = append(args, exprToSimpleString(arg))
		}
		return fmt.Sprintf("%s(%s)", e.Name, strings.Join(args, ", "))
	case *hclsyntax.ObjectConsExpr:
		// Object like { ip = v.ip, comment = v.comment }
		items := []string{}
		for _, item := range e.Items {
			key := exprToSimpleString(item.KeyExpr)
			val := exprToSimpleString(item.ValueExpr)
			items = append(items, fmt.Sprintf("%s = %s", key, val))
		}
		return fmt.Sprintf("{\n      %s\n    }", strings.Join(items, ",\n      "))
	case *hclsyntax.LiteralValueExpr:
		// String literal
		if e.Val.Type() == cty.String {
			return fmt.Sprintf(`"%s"`, e.Val.AsString())
		}
		return e.Val.GoString()
	case *hclsyntax.TemplateExpr:
		// Template string
		if e.IsStringLiteral() {
			val, _ := e.Value(nil)
			return fmt.Sprintf(`"%s"`, val.AsString())
		}
		return "\"${...}\""
	case *hclsyntax.ObjectConsKeyExpr:
		// Unwrap the key expression
		return exprToSimpleString(e.Wrapped)
	default:
		// For other types, try to get a basic representation
		return "..."
	}
}

// addMigrationWarning adds a warning comment to the body
func addMigrationWarning(body *hclwrite.Body, message string) {
	comment := fmt.Sprintf("\n  # MIGRATION WARNING: %s", message)
	tokens := hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenComment,
			Bytes: []byte(comment),
		},
	}
	body.AppendUnstructuredTokens(tokens)
}