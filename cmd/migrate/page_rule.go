package main

import (
	"strconv"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/zclconf/go-cty/cty"
)

// isPageRuleResource checks if a block is a cloudflare_page_rule resource
func isPageRuleResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		block.Labels()[0] == "cloudflare_page_rule"
}

// transformPageRuleBlock handles v4 to v5 config migration for cloudflare_page_rule
func transformPageRuleBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()

	// Transform the contents of the actions block BEFORE converting it to an attribute
	// Find actions block first
	var actionsBlock *hclwrite.Block
	for _, b := range body.Blocks() {
		if b.Type() == "actions" {
			actionsBlock = b
			break
		}
	}

	if actionsBlock != nil {
		// Transform cache_key_fields inside the actions block
		transformCacheKeyFieldsInActionsBlock(actionsBlock, diags)

		// Clean up empty attributes in the actions block
		cleanupEmptyAttributes(actionsBlock.Body())
	}

	// Finally, convert actions block to attribute (v5 requirement)
	convertActionsBlockToAttribute(body, diags)
}

// convertActionsBlockToAttribute converts actions from block to attribute syntax
// v4: actions { ... }
// v5: actions = { ... }
func convertActionsBlockToAttribute(body *hclwrite.Body, diags ast.Diagnostics) {
	// Find actions block
	var actionsBlock *hclwrite.Block
	var blocksToRemove []*hclwrite.Block

	for _, block := range body.Blocks() {
		if block.Type() == "actions" {
			actionsBlock = block
			blocksToRemove = append(blocksToRemove, block)
		}
	}

	if actionsBlock == nil {
		return
	}

	// Remove the actions block first
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}

	// Build the actions object from the block contents
	actionsBody := actionsBlock.Body()
	attrs := make(map[string]cty.Value)

	// Process all attributes from the block
	for name, attr := range actionsBody.Attributes() {
		// Get the raw tokens for this attribute to preserve expressions
		tokens := attr.Expr().BuildTokens(nil)
		// Create a temporary body to parse the expression
		tempBody := hclwrite.NewEmptyFile().Body()
		tempBody.SetAttributeRaw("temp", tokens)
		if tempAttr := tempBody.GetAttribute("temp"); tempAttr != nil {
			// Store the expression for reconstruction
			attrs[name] = cty.DynamicVal // Placeholder, we'll use raw tokens
		}
	}

	// Process nested blocks within actions
	nestedObjs := make(map[string]cty.Value)
	for _, block := range actionsBody.Blocks() {
		blockType := block.Type()
		if blockType == "cache_key_fields" || blockType == "forwarding_url" {
			// Convert nested block to object
			nestedObjs[blockType] = cty.DynamicVal // Placeholder
		}
	}

	// Now we need to reconstruct the entire actions as an attribute
	// We'll use a manual token construction approach to preserve exact formatting
	var tokens hclwrite.Tokens
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

	// Add all attributes from the original block
	for name, attr := range actionsBody.Attributes() {
		// Add indentation
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
		// Add attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		// Add attribute value
		tokens = append(tokens, attr.Expr().BuildTokens(nil)...)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	// Add nested blocks as nested objects
	for _, block := range actionsBody.Blocks() {
		blockType := block.Type()
		// Add indentation
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
		// Add block name as attribute
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(blockType)})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		// Convert block body to object
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

		// Add block's attributes
		blockBody := block.Body()
		for attrName, attr := range blockBody.Attributes() {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(attrName)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			tokens = append(tokens, attr.Expr().BuildTokens(nil)...)
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		// Handle nested blocks within this block (like query_string within cache_key_fields)
		for _, nestedBlock := range blockBody.Blocks() {
			nestedType := nestedBlock.Type()
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedType)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})

			nestedBody := nestedBlock.Body()
			for nestedAttrName, nestedAttr := range nestedBody.Attributes() {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        ")})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedAttrName)})
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
				tokens = append(tokens, nestedAttr.Expr().BuildTokens(nil)...)
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			}

			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
		}

		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
	}

	// Close the actions object
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})

	// Set the actions attribute with the constructed tokens
	body.SetAttributeRaw("actions", tokens)
}

// transformCacheKeyFieldsInActionsBlock transforms cache_key_fields inside an actions block
// In v4: cache_key_fields { query_string { ignore = true/false } }
// In v5: cache_key_fields { query_string { include = ["*"] } } or { exclude = ["*"] }
func transformCacheKeyFieldsInActionsBlock(actionsBlock *hclwrite.Block, diags ast.Diagnostics) {
	actionsBody := actionsBlock.Body()

	// Find cache_key_fields block
	var cacheKeyFieldsBlock *hclwrite.Block
	for _, block := range actionsBody.Blocks() {
		if block.Type() == "cache_key_fields" {
			cacheKeyFieldsBlock = block
			break
		}
	}

	if cacheKeyFieldsBlock == nil {
		return
	}

	cacheKeyFieldsBody := cacheKeyFieldsBlock.Body()

	// Find query_string block
	var queryStringBlock *hclwrite.Block
	for _, block := range cacheKeyFieldsBody.Blocks() {
		if block.Type() == "query_string" {
			queryStringBlock = block
			break
		}
	}

	if queryStringBlock == nil {
		return
	}

	queryStringBody := queryStringBlock.Body()

	// Check for 'ignore' attribute
	ignoreAttr := queryStringBody.GetAttribute("ignore")
	if ignoreAttr == nil {
		return
	}

	// Get the value of ignore attribute
	ignoreExpr := ignoreAttr.Expr()
	if ignoreExpr == nil {
		return
	}

	// Convert the expression to determine if it's true or false
	tokens := ignoreExpr.BuildTokens(nil)
	ignoreValue := strings.TrimSpace(string(tokens.Bytes()))

	// Remove the old 'ignore' attribute
	queryStringBody.RemoveAttribute("ignore")

	// Add new include/exclude attribute based on ignore value
	if ignoreValue == "false" {
		// ignore = false -> include all query strings
		queryStringBody.SetAttributeValue("include", cty.ListVal([]cty.Value{
			cty.StringVal("*"),
		}))
	} else if ignoreValue == "true" {
		// ignore = true -> exclude all query strings
		queryStringBody.SetAttributeValue("exclude", cty.ListVal([]cty.Value{
			cty.StringVal("*"),
		}))
	}
}

// transformPageRuleActions handles transformations within the actions block
func transformPageRuleActions(body *hclwrite.Body, diags ast.Diagnostics) {
	// Find actions block
	var actionsBlock *hclwrite.Block
	for _, block := range body.Blocks() {
		if block.Type() == "actions" {
			actionsBlock = block
			break
		}
	}

	if actionsBlock == nil {
		return
	}

	actionsBody := actionsBlock.Body()

	// Transform cache_ttl_by_status blocks to map attribute
	transformCacheTTLByStatusBlocks(actionsBody)

	// List of attributes that should be removed if they have empty/zero values
	// This helps with the v5 schema validation
	cleanupEmptyAttributes(actionsBody)
}

// transformCacheTTLByStatusBlocks converts multiple cache_ttl_by_status blocks to a single map attribute
// v4: multiple blocks like cache_ttl_by_status { codes = "200" ttl = 86400 }
// v5: single map attribute cache_ttl_by_status = { "200" = 86400 }
func transformCacheTTLByStatusBlocks(body *hclwrite.Body) {
	ttlMap := make(map[string]cty.Value)
	var blocksToRemove []*hclwrite.Block

	// Collect all cache_ttl_by_status blocks
	for _, block := range body.Blocks() {
		if block.Type() == "cache_ttl_by_status" {
			blockBody := block.Body()

			// Get codes and ttl attributes
			codesAttr := blockBody.GetAttribute("codes")
			ttlAttr := blockBody.GetAttribute("ttl")

			if codesAttr != nil && ttlAttr != nil {
				// Get the string value for codes
				codesTokens := codesAttr.Expr().BuildTokens(nil)
				codesValue := strings.Trim(strings.TrimSpace(string(codesTokens.Bytes())), `"`)

				// Get the ttl value
				ttlTokens := ttlAttr.Expr().BuildTokens(nil)
				ttlValue := strings.TrimSpace(string(ttlTokens.Bytes()))

				// Try to parse as number
				if ttlNum, err := strconv.ParseInt(ttlValue, 10, 64); err == nil {
					ttlMap[codesValue] = cty.NumberIntVal(ttlNum)
				} else {
					// Keep as string if not a number
					ttlMap[codesValue] = cty.StringVal(ttlValue)
				}
			}

			blocksToRemove = append(blocksToRemove, block)
		}
	}

	// Remove the blocks
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}

	// Add as map attribute if we collected any values
	if len(ttlMap) > 0 {
		body.SetAttributeValue("cache_ttl_by_status", cty.MapVal(ttlMap))
	}
}

// cleanupEmptyAttributes removes attributes with empty or zero values that cause issues
func cleanupEmptyAttributes(body *hclwrite.Body) {
	// First remove deprecated fields that don't exist in v5
	deprecatedFields := []string{
		"disable_railgun",
		"minify",
		"server_side_exclude",
	}

	for _, attrName := range deprecatedFields {
		if body.GetAttribute(attrName) != nil {
			body.RemoveAttribute(attrName)
		}
	}

	// Then remove attributes with empty or zero values
	attributesToCheck := []string{
		"browser_cache_ttl",
		"edge_cache_ttl",
		"origin_error_page_pass_thru",
		"polish",
		"sort_query_string_for_cache",
	}

	for _, attrName := range attributesToCheck {
		attr := body.GetAttribute(attrName)
		if attr != nil {
			tokens := attr.Expr().BuildTokens(nil)
			value := strings.TrimSpace(string(tokens.Bytes()))

			// Remove if empty string or zero
			// Check for both quoted and unquoted empty strings
			if value == `""` || value == "" || value == "0" {
				body.RemoveAttribute(attrName)
			}
		}
	}
}

// transformPageRuleStateJSON handles complete v4 to v5 state migration for cloudflare_page_rule
// This replaces the partial transformation and handles all aspects of page_rule state migration
func transformPageRuleStateJSON(json string, instancePath string) string {
	attrPath := instancePath + ".attributes"
	result := json

	// Set schema_version to 0 for v5
	result, _ = sjson.Set(result, instancePath+".schema_version", 0)

	// Check if actions attribute exists
	actions := gjson.Get(json, attrPath+".actions")
	if !actions.Exists() {
		return result
	}

	// IMPORTANT: Handle the case where actions is an array instead of an object
	// In v4, actions might be stored as [{}] but v5 expects {}
	if actions.IsArray() {
		actionsArray := actions.Array()
		if len(actionsArray) == 1 {
			// Unwrap single-element array to object
			result, _ = sjson.Set(result, attrPath+".actions", actionsArray[0].Value())
		} else if len(actionsArray) == 0 {
			// Empty array becomes empty object
			result, _ = sjson.Set(result, attrPath+".actions", map[string]interface{}{})
			return result
		}
	}

	// Handle actions transformations
	result = transformPageRuleActionsState(result, attrPath+".actions")

	// Handle cache_key_fields transformations
	result = transformPageRuleCacheKeyFieldsState(result, attrPath+".actions.cache_key_fields")

	// Handle cache_ttl_by_status transformations
	result = transformPageRuleCacheTTLByStatusState(result, attrPath+".actions")

	// Handle forwarding_url transformations (from YAML spec)
	result = transformPageRuleForwardingURLState(result, attrPath+".actions")

	return result
}

// transformPageRuleActionsState handles actions-level state transformations
func transformPageRuleActionsState(json string, actionsPath string) string {
	result := json

	// FIRST: Check ALL fields in actions and unwrap any single-element arrays
	// In v4, many fields might be stored as arrays but v5 expects them as direct values
	actions := gjson.Get(result, actionsPath)
	if actions.Exists() && actions.IsObject() {
		// Iterate through all fields in actions
		actions.ForEach(func(key, value gjson.Result) bool {
			fieldPath := actionsPath + "." + key.String()

			// If this field is an array, handle it
			if value.IsArray() {
				valueArray := value.Array()
				if len(valueArray) == 1 {
					// Unwrap single-element array
					result, _ = sjson.Set(result, fieldPath, valueArray[0].Value())
				} else if len(valueArray) == 0 {
					// Empty array - remove it entirely
					result, _ = sjson.Delete(result, fieldPath)
				}
			}
			return true // continue iteration
		})
	}

	// Remove fields that exist in v4 but are not supported in v5
	deprecatedFields := []string{
		"disable_railgun",
		"minify",
		"server_side_exclude",
	}

	for _, field := range deprecatedFields {
		fieldPath := actionsPath + "." + field
		if gjson.Get(result, fieldPath).Exists() {
			result, _ = sjson.Delete(result, fieldPath)
		}
	}

	// List of numeric fields that should not be empty strings or null
	numericFields := []string{
		"browser_cache_ttl",
		"edge_cache_ttl",
		"origin_error_page_pass_thru",
		"polish",
		"sort_query_string_for_cache",
	}

	// Fix any empty string or null numeric fields
	for _, field := range numericFields {
		fieldPath := actionsPath + "." + field
		fieldValue := gjson.Get(result, fieldPath)

		// If the field exists and is null or empty string, remove it completely
		if fieldValue.Exists() {
			if fieldValue.Type == gjson.Null ||
			   (fieldValue.Type == gjson.String && fieldValue.String() == "") ||
			   (fieldValue.Type == gjson.Number && fieldValue.Float() == 0) {
				result, _ = sjson.Delete(result, fieldPath)
			}
		}
	}

	// Handle fields that use "on"/"off" string values
	// Some of these need to be converted to booleans, others stay as strings
	onOffBooleanFields := []string{
		"always_use_https",
		"always_online",
		"email_obfuscation",
		"server_side_exclude",
		"opportunistic_encryption",
		"respect_strong_etag",
		"response_buffering",
		"true_client_ip_header",
		"websockets",
		"mirage",
		"rocket_loader",
	}

	for _, field := range onOffBooleanFields {
		fieldPath := actionsPath + "." + field
		fieldValue := gjson.Get(json, fieldPath)

		if fieldValue.Exists() && fieldValue.Type == gjson.String {
			strVal := fieldValue.String()
			if strVal == "on" || strVal == "true" {
				result, _ = sjson.Set(result, fieldPath, true)
			} else if strVal == "off" || strVal == "false" {
				result, _ = sjson.Set(result, fieldPath, false)
			}
		}
	}

	// These fields keep their "on"/"off" string values
	onOffStringFields := []string{
		"automatic_https_rewrites",
		"browser_check",
		"cache_by_device_type",
		"cache_deception_armor",
	}

	// Just ensure they're not empty strings
	for _, field := range onOffStringFields {
		fieldPath := actionsPath + "." + field
		fieldValue := gjson.Get(json, fieldPath)

		if fieldValue.Exists() && fieldValue.Type == gjson.String && fieldValue.String() == "" {
			result, _ = sjson.Delete(result, fieldPath)
		}
	}

	// Handle string enum fields - remove empty values
	stringEnumFields := []string{
		"automatic_https_rewrites",
		"browser_check",
		"cache_by_device_type",
		"cache_deception_armor",
		"cache_level",
		"security_level",
		"ssl",
	}

	for _, field := range stringEnumFields {
		fieldPath := actionsPath + "." + field
		fieldValue := gjson.Get(json, fieldPath)

		if fieldValue.Exists() && fieldValue.Type == gjson.String {
			strVal := fieldValue.String()
			// Remove if empty
			if strVal == "" {
				result, _ = sjson.Delete(result, fieldPath)
			}
		}
	}

	return result
}

// transformPageRuleCacheKeyFieldsState handles cache_key_fields state transformations
func transformPageRuleCacheKeyFieldsState(json string, cacheKeyFieldsPath string) string {
	result := json

	// Check if cache_key_fields exists
	cacheKeyFields := gjson.Get(json, cacheKeyFieldsPath)
	if !cacheKeyFields.Exists() {
		return result
	}

	// IMPORTANT: Handle the case where cache_key_fields is an array instead of an object
	// In v4, cache_key_fields might be stored as [{}] but v5 expects {}
	if cacheKeyFields.IsArray() {
		cacheKeyFieldsArray := cacheKeyFields.Array()
		if len(cacheKeyFieldsArray) == 1 {
			// Unwrap single-element array to object
			result, _ = sjson.Set(result, cacheKeyFieldsPath, cacheKeyFieldsArray[0].Value())
		} else if len(cacheKeyFieldsArray) == 0 {
			// Empty array - remove it entirely as cache_key_fields is optional
			result, _ = sjson.Delete(result, cacheKeyFieldsPath)
			return result
		}
		// Re-get the cache_key_fields after unwrapping
		cacheKeyFields = gjson.Get(result, cacheKeyFieldsPath)
	}

	// Handle nested objects within cache_key_fields that might also be arrays
	// Check and unwrap query_string if it's an array
	queryStringPath := cacheKeyFieldsPath + ".query_string"
	queryString := gjson.Get(result, queryStringPath)

	if queryString.IsArray() {
		queryStringArray := queryString.Array()
		if len(queryStringArray) == 1 {
			result, _ = sjson.Set(result, queryStringPath, queryStringArray[0].Value())
		} else if len(queryStringArray) == 0 {
			result, _ = sjson.Delete(result, queryStringPath)
		}
		// Re-get after potential unwrapping
		queryString = gjson.Get(result, queryStringPath)
	}

	// Transform query_string.ignore to include/exclude
	if queryString.Exists() && queryString.IsObject() {
		ignorePath := queryStringPath + ".ignore"
		ignoreValue := gjson.Get(result, ignorePath)

		if ignoreValue.Exists() {
			// Remove the ignore field
			result, _ = sjson.Delete(result, ignorePath)

			// Add include or exclude based on ignore value
			if ignoreValue.Type == gjson.False {
				// ignore = false -> include all query strings
				result, _ = sjson.Set(result, queryStringPath+".include", []string{"*"})
			} else if ignoreValue.Type == gjson.True {
				// ignore = true -> exclude all query strings
				result, _ = sjson.Set(result, queryStringPath+".exclude", []string{"*"})
			}
		}
	}

	// Check and unwrap other nested objects if they are arrays
	nestedFields := []string{"cookie", "header", "host", "user"}
	for _, field := range nestedFields {
		fieldPath := cacheKeyFieldsPath + "." + field
		fieldValue := gjson.Get(result, fieldPath)

		if fieldValue.IsArray() {
			fieldArray := fieldValue.Array()
			if len(fieldArray) == 1 {
				result, _ = sjson.Set(result, fieldPath, fieldArray[0].Value())
			} else if len(fieldArray) == 0 {
				result, _ = sjson.Delete(result, fieldPath)
			}
		}
	}

	// Ensure all array fields have proper empty arrays instead of nil
	// This prevents nil vs empty array differences in v5

	// For header fields
	headerPath := cacheKeyFieldsPath + ".header"
	if gjson.Get(result, headerPath).Exists() {
		headerArrayFields := []string{"include", "exclude", "check_presence"}
		for _, field := range headerArrayFields {
			fieldPath := headerPath + "." + field
			fieldValue := gjson.Get(result, fieldPath)
			if !fieldValue.Exists() || fieldValue.Type == gjson.Null {
				result, _ = sjson.Set(result, fieldPath, []string{})
			}
		}
	}

	// For cookie fields
	cookiePath := cacheKeyFieldsPath + ".cookie"
	if gjson.Get(result, cookiePath).Exists() {
		cookieArrayFields := []string{"include", "check_presence"}
		for _, field := range cookieArrayFields {
			fieldPath := cookiePath + "." + field
			fieldValue := gjson.Get(result, fieldPath)
			if !fieldValue.Exists() || fieldValue.Type == gjson.Null {
				result, _ = sjson.Set(result, fieldPath, []string{})
			}
		}
	}

	// Ensure boolean fields have explicit false values instead of nil
	// This prevents plan differences in v5 which has default values for these fields
	userPath := cacheKeyFieldsPath + ".user"
	if gjson.Get(result, userPath).Exists() {
		// Set defaults for user fields if they are nil
		userFields := []string{"device_type", "geo", "lang"}
		for _, field := range userFields {
			fieldPath := userPath + "." + field
			fieldValue := gjson.Get(result, fieldPath)
			if !fieldValue.Exists() || fieldValue.Type == gjson.Null {
				result, _ = sjson.Set(result, fieldPath, false)
			}
		}
	}

	// Set default for host.resolved if it's nil
	hostPath := cacheKeyFieldsPath + ".host"
	if gjson.Get(result, hostPath).Exists() {
		resolvedPath := hostPath + ".resolved"
		resolvedValue := gjson.Get(result, resolvedPath)
		if !resolvedValue.Exists() || resolvedValue.Type == gjson.Null {
			result, _ = sjson.Set(result, resolvedPath, false)
		}
	}

	return result
}

// transformPageRuleCacheTTLByStatusState transforms cache_ttl_by_status from array format to object format
// v4: cache_ttl_by_status: [{"codes": "200", "ttl": 86400}, {"codes": "404", "ttl": 300}]
// v5: cache_ttl_by_status: {"200": "86400", "404": "300"}
func transformPageRuleCacheTTLByStatusState(json string, actionsPath string) string {
	result := json

	cacheTTLPath := actionsPath + ".cache_ttl_by_status"
	cacheTTL := gjson.Get(json, cacheTTLPath)

	if !cacheTTL.Exists() {
		return result
	}

	// Only transform if it's an array (v4 format)
	// If it's already an object (v5 format), leave it alone to avoid corruption
	if cacheTTL.IsArray() {
		ttlMap := make(map[string]string) // v5 expects string values

		// v4 stores as array of objects with "codes" and "ttl" fields
		for _, item := range cacheTTL.Array() {
			if item.IsObject() {
				codes := gjson.Get(item.Raw, "codes")
				ttl := gjson.Get(item.Raw, "ttl")

				if codes.Exists() && ttl.Exists() {
					codesValue := codes.String()

					// Convert ttl to string (v4 has numbers, v5 needs strings)
					var ttlValue string
					if ttl.Type == gjson.Number {
						ttlValue = ttl.String()
					} else if ttl.Type == gjson.String {
						ttlValue = ttl.String()
					} else {
						continue
					}

					ttlMap[codesValue] = ttlValue
				}
			}
		}

		// Replace the array with the map
		if len(ttlMap) > 0 {
			result, _ = sjson.Set(result, cacheTTLPath, ttlMap)
		} else if len(cacheTTL.Array()) == 0 {
			// Empty array - remove the field
			result, _ = sjson.Delete(result, cacheTTLPath)
		}
	}
	// Don't touch it if it's already an object - v5 format should be left as-is

	return result
}

// transformPageRuleForwardingURLState handles forwarding_url transformations
func transformPageRuleForwardingURLState(json string, actionsPath string) string {
	result := json

	forwardingURLPath := actionsPath + ".forwarding_url"
	forwardingURL := gjson.Get(json, forwardingURLPath)

	if !forwardingURL.Exists() {
		return result
	}

	// Handle empty array -> null transformation
	if forwardingURL.IsArray() && len(forwardingURL.Array()) == 0 {
		result, _ = sjson.Delete(result, forwardingURLPath)
		return result
	}

	// Handle single-element array unwrapping
	if forwardingURL.IsArray() && len(forwardingURL.Array()) == 1 {
		// Unwrap the single element
		result, _ = sjson.Set(result, forwardingURLPath, forwardingURL.Array()[0].Value())
	}

	return result
}