package main

import (
	"slices"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/zclconf/go-cty/cty"
)

// v4BindingToV5Type maps v4 binding attribute/block names to v5 binding types
var v4BindingToV5Type = map[string]string{
	"plain_text_binding":        "plain_text",
	"kv_namespace_binding":      "kv_namespace",
	"secret_text_binding":       "secret_text",
	"r2_bucket_binding":         "r2_bucket",
	"queue_binding":             "queue",
	"d1_database_binding":       "d1",
	"analytics_engine_binding":  "analytics_engine",
	"service_binding":           "service",
	"hyperdrive_config_binding": "hyperdrive",
}

// isWorkersScriptResource checks if a block is a workers_script resource
func isWorkersScriptResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_script" || block.Labels()[0] == "cloudflare_worker_script")
}

// transformWorkersScriptBlock transforms cloudflare_workers_script resources
// V4: name -> V5: script_name
// Also handles resource rename: cloudflare_worker_script -> cloudflare_workers_script
// Also transforms v4 binding blocks to v5 bindings list
func transformWorkersScriptBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource rename: cloudflare_worker_script -> cloudflare_workers_script
	if len(block.Labels()) >= 1 &&
		block.Labels()[0] == "cloudflare_worker_script" {
		block.SetLabels([]string{"cloudflare_workers_script", block.Labels()[1]})
	}

	// Rename name attribute to script_name
	block.Body().RenameAttribute("name", "script_name")

	// Transform v4 binding blocks to v5 bindings list
	transformBindings(block, diags)

	// Transform v4 dispatch_namespace attribute to v5 dispatch_namespace binding
	transformDispatchNamespace(block, diags)

	// Transform v4 module attribute to v5 main_module/body_part attributes
	transformModule(block, diags)
}

// transformWorkersScriptStateJSON transforms the state JSON for workers_script
// V4: name -> V5: script_name
// Also transforms v4 binding attributes to v5 bindings list
func transformWorkersScriptStateJSON(jsonStr string, path string) string {
	result := jsonStr

	// Get the current name value
	namePath := path + ".attributes.name"
	scriptNamePath := path + ".attributes.script_name"

	// Get the value and move it
	nameValue := gjson.Get(jsonStr, namePath)
	if nameValue.Exists() {
		// Set the new attribute
		result, _ = sjson.Set(result, scriptNamePath, nameValue.Value())
		// Delete the old attribute
		result, _ = sjson.Delete(result, namePath)
	}

	// Transform v4 binding attributes to v5 bindings list
	result = transformBindingsInState(result, path)

	// Transform v4 dispatch_namespace attribute to v5 dispatch_namespace binding
	result = transformDispatchNamespaceInState(result, path)

	// Transform v4 module attribute to v5 main_module/body_part attributes
	result = transformModuleInState(result, path)

	// Fix placement attribute format: v4 uses array [], v5 expects object {}
	placementPath := path + ".attributes.placement"
	placementValue := gjson.Get(result, placementPath)
	if placementValue.Exists() && placementValue.IsArray() {
		// If placement is an empty array, set it to empty object
		if len(placementValue.Array()) == 0 {
			result, _ = sjson.Set(result, placementPath, map[string]interface{}{})
		}
	}

	return result
}

// transformBindingsInState converts v4 binding attributes to v5 bindings format in state JSON
func transformBindingsInState(jsonStr string, path string) string {
	result := jsonStr

	// List of v4 binding attribute names to look for
	// Get list of v4 binding attribute names from the map
	var v4BindingAttributes []string
	for bindingAttr := range v4BindingToV5Type {
		v4BindingAttributes = append(v4BindingAttributes, bindingAttr)
	}

	var bindings []interface{}

	// Check each v4 binding type and convert to v5 bindings format
	for _, bindingAttr := range v4BindingAttributes {
		bindingPath := path + ".attributes." + bindingAttr
		bindingValue := gjson.Get(jsonStr, bindingPath)

		if bindingValue.Exists() {
			// Map v4 binding attribute to v5 binding type
			bindingType, supported := v4BindingToV5Type[bindingAttr]
			if !supported {
				// Skip unknown binding types
				continue
			}

			// Parse the binding data and add type field
			bindingData := bindingValue.Value()
			if bindingArray, ok := bindingData.([]interface{}); ok {
				// Handle array of bindings
				for _, binding := range bindingArray {
					if bindingMap, ok := binding.(map[string]interface{}); ok {
						bindingMap["type"] = bindingType
						// Apply attribute renames for this binding type
						renameBindingAttributes(bindingMap, bindingAttr)
						bindings = append(bindings, bindingMap)
					}
				}
			} else if bindingMap, ok := bindingData.(map[string]interface{}); ok {
				// Handle single binding
				bindingMap["type"] = bindingType
				// Apply attribute renames for this binding type
				renameBindingAttributes(bindingMap, bindingAttr)
				bindings = append(bindings, bindingMap)
			}

			// Remove the old binding attribute
			result, _ = sjson.Delete(result, bindingPath)
		}
	}

	// Handle webassembly_binding separately - not supported in v5, just remove
	webassemblyPath := path + ".attributes.webassembly_binding"
	webassemblyValue := gjson.Get(jsonStr, webassemblyPath)
	if webassemblyValue.Exists() {
		// Remove the webassembly_binding attribute without migration
		result, _ = sjson.Delete(result, webassemblyPath)
		// Note: Warning will be generated by config transformation, not state transformation
	}

	// If we found any bindings, add them to the state (preserve original ordering)
	if len(bindings) > 0 {
		bindingsPath := path + ".attributes.bindings"
		result, _ = sjson.Set(result, bindingsPath, bindings)
	}

	return result
}

// renameBindingAttributes renames attributes within binding objects for v4→v5 migration
func renameBindingAttributes(bindingMap map[string]interface{}, bindingType string) {
	switch bindingType {
	case "d1_database_binding":
		// d1_database_binding: database_id → id
		if databaseID, exists := bindingMap["database_id"]; exists {
			bindingMap["id"] = databaseID
			delete(bindingMap, "database_id")
		}
	}
}

// transformDispatchNamespaceInState transforms v4 dispatch_namespace attribute to v5 dispatch_namespace binding in state
func transformDispatchNamespaceInState(jsonStr string, path string) string {
	result := jsonStr

	// Check if dispatch_namespace attribute exists
	dispatchNamespacePath := path + ".attributes.dispatch_namespace"
	dispatchValue := gjson.Get(jsonStr, dispatchNamespacePath)

	if !dispatchValue.Exists() {
		return result // No dispatch_namespace to transform
	}

	// Extract the dispatch namespace value
	namespaceID := dispatchValue.String()

	// Create dispatch_namespace binding
	dispatchBinding := map[string]interface{}{
		"type":         "dispatch_namespace",
		"namespace_id": namespaceID,
	}

	// Get existing bindings
	bindingsPath := path + ".attributes.bindings"
	existingBindings := gjson.Get(result, bindingsPath)

	var allBindings []interface{}

	// Parse existing bindings if they exist
	if existingBindings.Exists() && existingBindings.IsArray() {
		for _, binding := range existingBindings.Array() {
			allBindings = append(allBindings, binding.Value())
		}
	}

	// Add the dispatch_namespace binding
	allBindings = append(allBindings, dispatchBinding)

	// Update state with new bindings
	result, _ = sjson.Set(result, bindingsPath, allBindings)

	// Remove the original dispatch_namespace attribute
	result, _ = sjson.Delete(result, dispatchNamespacePath)

	return result
}

// transformModuleInState transforms v4 module attribute to v5 main_module/body_part attributes in state
func transformModuleInState(jsonStr string, path string) string {
	result := jsonStr

	// Check if module attribute exists
	modulePath := path + ".attributes.module"
	moduleValue := gjson.Get(jsonStr, modulePath)

	if !moduleValue.Exists() {
		return result // No module attribute to transform
	}

	// Extract the boolean value
	isModule := moduleValue.Bool()

	if isModule {
		// module = true → main_module = "worker.js" (ES Module syntax)
		mainModulePath := path + ".attributes.main_module"
		result, _ = sjson.Set(result, mainModulePath, "worker.js")
	} else {
		// module = false → body_part = "worker.js" (Service Worker syntax)
		bodyPartPath := path + ".attributes.body_part"
		result, _ = sjson.Set(result, bodyPartPath, "worker.js")
	}

	// Remove the original module attribute
	result, _ = sjson.Delete(result, modulePath)

	return result
}

// transformBindings transforms v4 binding blocks to v5 bindings list
// V4 had separate blocks: plain_text_binding, kv_namespace_binding, etc.
// V5 has unified bindings list with type field
func transformBindings(block *hclwrite.Block, diags ast.Diagnostics) {
	var bindings []hclsyntax.Expression
	var blocksToRemove []*hclwrite.Block

	// Scan all blocks to find binding blocks
	for _, childBlock := range block.Body().Blocks() {
		if bindingType, isBindingBlock := v4BindingToV5Type[childBlock.Type()]; isBindingBlock {
			// Convert this block to a binding object
			binding := transformBindingBlock(childBlock, bindingType, diags)
			if binding != nil {
				bindings = append(bindings, binding)
			}
			// Mark this block for removal
			blocksToRemove = append(blocksToRemove, childBlock)
		} else if childBlock.Type() == "webassembly_binding" {
			// webassembly_binding is not supported in v5 - replace with warning comment
			warningTokens := []*hclwrite.Token{
				{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
				{Type: hclsyntax.TokenComment, Bytes: []byte(`  # MIGRATION WARNING: webassembly_binding is not supported in v5.`)},
				{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
				{Type: hclsyntax.TokenComment, Bytes: []byte(`  # WebAssembly modules must be bundled into the script content instead.`)},
				{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
				{Type: hclsyntax.TokenComment, Bytes: []byte(`  # Please update your build process and remove this warning.`)},
				{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			}
			// Add warning comment before removing the block
			block.Body().AppendUnstructuredTokens(warningTokens)
			// Mark for removal
			blocksToRemove = append(blocksToRemove, childBlock)
		}
	}

	// Remove the old binding blocks
	for _, blockToRemove := range blocksToRemove {
		block.Body().RemoveBlock(blockToRemove)
	}

	// If we found any bindings, create the bindings attribute
	if len(bindings) > 0 {
		bindingsExpr := &hclsyntax.TupleConsExpr{
			Exprs: bindings,
		}

		// Use AST to set the bindings expression directly (it will create the attribute if it doesn't exist)
		transforms := map[string]ast.ExprTransformer{
			"bindings": func(expr *hclsyntax.Expression, diags ast.Diagnostics) {
				*expr = bindingsExpr
			},
		}
		ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
	}
}

// transformBindingBlock converts a v4 binding block to a v5 binding object
func transformBindingBlock(block *hclwrite.Block, bindingType string, diags ast.Diagnostics) hclsyntax.Expression {
	var items []hclsyntax.ObjectConsItem

	// Always add the type field
	items = append(items, hclsyntax.ObjectConsItem{
		KeyExpr:   ast.NewKeyExpr("type"),
		ValueExpr: &hclsyntax.LiteralValueExpr{Val: cty.StringVal(bindingType)},
	})

	// Copy all attributes from the binding block to the binding object
	// Sort attribute names for deterministic output
	var attrNames []string
	attributes := block.Body().Attributes()
	for attrName := range attributes {
		attrNames = append(attrNames, attrName)
	}
	slices.Sort(attrNames)

	for _, attrName := range attrNames {
		attr := attributes[attrName]
		// Convert the hclwrite.Expression to hclsyntax.Expression
		syntaxExpr := ast.WriteExpr2Expr(attr.Expr(), diags)

		// Apply attribute renaming for this binding type
		finalAttrName := renameBindingAttribute(attrName, bindingType)

		items = append(items, hclsyntax.ObjectConsItem{
			KeyExpr:   ast.NewKeyExpr(finalAttrName),
			ValueExpr: syntaxExpr,
		})
	}

	return &hclsyntax.ObjectConsExpr{
		Items: items,
	}
}

// renameBindingAttribute renames a single attribute name for v4→v5 binding migration
func renameBindingAttribute(attrName, bindingType string) string {
	switch bindingType {
	case "d1":
		// d1 binding: database_id → id
		if attrName == "database_id" {
			return "id"
		}
	}
	return attrName
}

// transformDispatchNamespace transforms v4 dispatch_namespace attribute to v5 dispatch_namespace binding
func transformDispatchNamespace(block *hclwrite.Block, diags ast.Diagnostics) {
	// Check if block has dispatch_namespace attribute
	dispatchAttr := block.Body().GetAttribute("dispatch_namespace")
	if dispatchAttr == nil {
		return // No dispatch_namespace to transform
	}

	// Extract the dispatch namespace value
	dispatchExpr := ast.WriteExpr2Expr(dispatchAttr.Expr(), diags)
	dispatchValue := ast.Expr2S(dispatchExpr, diags)

	// Handle quoted strings by removing quotes
	if strings.HasPrefix(dispatchValue, `"`) && strings.HasSuffix(dispatchValue, `"`) {
		dispatchValue = strings.Trim(dispatchValue, `"`)
	}

	// If we can't parse the value, add a manual migration warning
	if dispatchValue == "" || strings.Contains(dispatchValue, "var.") || strings.Contains(dispatchValue, "local.") {
		warningTokens := []*hclwrite.Token{
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenComment, Bytes: []byte(`  # MIGRATION WARNING: dispatch_namespace attribute needs manual migration to dispatch_namespace binding`)},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenComment, Bytes: []byte(`  # Convert: dispatch_namespace = "value" → bindings = [{type = "dispatch_namespace", namespace_id = "value"}]`)},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		}
		block.Body().AppendUnstructuredTokens(warningTokens)
		return
	}

	// Create dispatch_namespace binding
	dispatchBinding := &hclsyntax.ObjectConsExpr{
		Items: []hclsyntax.ObjectConsItem{
			{
				KeyExpr:   ast.NewKeyExpr("type"),
				ValueExpr: &hclsyntax.LiteralValueExpr{Val: cty.StringVal("dispatch_namespace")},
			},
			{
				KeyExpr:   ast.NewKeyExpr("namespace_id"),
				ValueExpr: &hclsyntax.LiteralValueExpr{Val: cty.StringVal(dispatchValue)},
			},
		},
	}

	// Add this binding to the existing bindings or create new bindings attribute
	existingBindingsAttr := block.Body().GetAttribute("bindings")
	var allBindings []hclsyntax.Expression

	if existingBindingsAttr != nil {
		// Parse existing bindings
		existingExpr := ast.WriteExpr2Expr(existingBindingsAttr.Expr(), diags)
		if tuple, ok := existingExpr.(*hclsyntax.TupleConsExpr); ok {
			allBindings = append(allBindings, tuple.Exprs...)
		} else {
			// Single binding or other format - preserve it
			allBindings = append(allBindings, existingExpr)
		}
	}

	// Add the dispatch_namespace binding
	allBindings = append(allBindings, dispatchBinding)

	// Create the new bindings expression
	bindingsExpr := &hclsyntax.TupleConsExpr{
		Exprs: allBindings,
	}

	// Set the bindings attribute
	transforms := map[string]ast.ExprTransformer{
		"bindings": func(expr *hclsyntax.Expression, diags ast.Diagnostics) {
			*expr = bindingsExpr
		},
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)

	// Remove the original dispatch_namespace attribute
	block.Body().RemoveAttribute("dispatch_namespace")
}

// transformModule transforms v4 module attribute to v5 main_module/body_part attributes
func transformModule(block *hclwrite.Block, diags ast.Diagnostics) {
	// Check if block has module attribute
	moduleAttr := block.Body().GetAttribute("module")
	if moduleAttr == nil {
		return // No module attribute to transform
	}

	// Extract the module value
	moduleExpr := ast.WriteExpr2Expr(moduleAttr.Expr(), diags)
	moduleValue := ast.Expr2S(moduleExpr, diags)

	// Handle different boolean representations
	var isModule bool
	var parseError bool

	switch strings.ToLower(strings.TrimSpace(moduleValue)) {
	case "true":
		isModule = true
	case "false":
		isModule = false
	default:
		// If we can't parse the value, add a manual migration warning
		warningTokens := []*hclwrite.Token{
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenComment, Bytes: []byte(`  # MIGRATION WARNING: module attribute needs manual migration`)},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenComment, Bytes: []byte(`  # Convert: module = true → main_module = "worker.js"`)},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenComment, Bytes: []byte(`  # Convert: module = false → body_part = "worker.js"`)},
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		}
		block.Body().AppendUnstructuredTokens(warningTokens)
		parseError = true
	}

	if !parseError {
		// Create the appropriate attribute based on module value
		if isModule {
			// module = true → main_module = "worker.js" (ES Module syntax)
			block.Body().SetAttributeValue("main_module", cty.StringVal("worker.js"))
		} else {
			// module = false → body_part = "worker.js" (Service Worker syntax)
			block.Body().SetAttributeValue("body_part", cty.StringVal("worker.js"))
		}
	}

	// Remove the original module attribute
	block.Body().RemoveAttribute("module")
}
