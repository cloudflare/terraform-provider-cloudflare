package main

import (
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// mapSettingName translates v4 setting names to v5 setting names
// This mirrors the logic from cmd/migrate/zone_settings.go
func mapSettingName(v4Name string) string {
	settingNameMap := map[string]string{
		"zero_rtt": "0rtt", // v4 used "zero_rtt" but API expects "0rtt"
	}
	
	if v5Name, exists := settingNameMap[v4Name]; exists {
		return v5Name
	}
	return v4Name
}

// createZoneSettingResourceWithSubstitution creates a cloudflare_zone_setting resource with variable substitution
func createZoneSettingResourceWithSubstitution(name, settingID string, zoneIDAttr, valueAttr *hclwrite.Attribute, moduleCall *ModuleCall) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_zone_setting", name})
	body := block.Body()

	// Set zone_id with variable substitution
	if zoneIDAttr != nil {
		substitutedZoneID := substituteVariables(zoneIDAttr, moduleCall)
		body.SetAttributeRaw("zone_id", substitutedZoneID)
	}

	// Set setting_id
	body.SetAttributeValue("setting_id", cty.StringVal(settingID))

	// Set value with variable substitution
	if valueAttr != nil {
		substitutedValue := substituteVariables(valueAttr, moduleCall)
		body.SetAttributeRaw("value", substitutedValue)
	}

	return block
}

// createImportBlockWithSubstitution creates an import block with variable substitution
func createImportBlockWithSubstitution(resourceName, settingID string, zoneIDAttr *hclwrite.Attribute, moduleCall *ModuleCall) *hclwrite.Block {
	block := hclwrite.NewBlock("import", nil)
	body := block.Body()

	// Build the "to" value: cloudflare_zone_setting.resource_name
	toTokens := buildResourceReference("cloudflare_zone_setting", resourceName)
	body.SetAttributeRaw("to", toTokens)

	// Build the "id" value with variable substitution: "${zone_id}/setting_id"
	if zoneIDAttr != nil {
		substitutedZoneID := substituteVariables(zoneIDAttr, moduleCall)
		idTokens := buildTemplateStringTokens(substitutedZoneID, "/"+settingID)
		body.SetAttributeRaw("id", idTokens)
	}

	return block
}

// transformSecurityHeaderBlockWithSubstitution transforms a security_header block with variable substitution
func transformSecurityHeaderBlockWithSubstitution(resourceName string, zoneIDAttr *hclwrite.Attribute, securityHeaderBlock *hclwrite.Block, moduleCall *ModuleCall) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_zone_setting", resourceName})
	body := block.Body()

	// Set zone_id with variable substitution
	if zoneIDAttr != nil {
		substitutedZoneID := substituteVariables(zoneIDAttr, moduleCall)
		body.SetAttributeRaw("zone_id", substitutedZoneID)
	}

	// Set setting_id
	body.SetAttributeValue("setting_id", cty.StringVal("security_header"))

	// Build the object tokens with variable substitution
	objectTokens := buildObjectFromBlockWithSubstitution(securityHeaderBlock, moduleCall)
	body.SetAttributeRaw("value", objectTokens)

	return block
}

// transformNELBlockWithSubstitution transforms a nel block with variable substitution
func transformNELBlockWithSubstitution(resourceName string, zoneIDAttr *hclwrite.Attribute, nelBlock *hclwrite.Block, moduleCall *ModuleCall) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_zone_setting", resourceName})
	body := block.Body()

	// Set zone_id with variable substitution
	if zoneIDAttr != nil {
		substitutedZoneID := substituteVariables(zoneIDAttr, moduleCall)
		body.SetAttributeRaw("zone_id", substitutedZoneID)
	}

	// Set setting_id
	body.SetAttributeValue("setting_id", cty.StringVal("nel"))

	// Build the object tokens with variable substitution
	objectTokens := buildObjectFromBlockWithSubstitution(nelBlock, moduleCall)
	body.SetAttributeRaw("value", objectTokens)

	return block
}

// buildObjectFromBlockWithSubstitution creates object tokens from a block's attributes with variable substitution
func buildObjectFromBlockWithSubstitution(block *hclwrite.Block, moduleCall *ModuleCall) hclwrite.Tokens {
	// Get attributes in their original order (this function would need to be implemented or imported)
	var attrs []hclwrite.ObjectAttrTokens
	
	for name, attr := range block.Body().Attributes() {
		// Create tokens for the attribute name
		nameTokens := hclwrite.TokensForIdentifier(name)
		
		// Get the value tokens with variable substitution
		valueTokens := substituteVariables(attr, moduleCall)
		
		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  nameTokens,
			Value: valueTokens,
		})
	}
	
	// Use the built-in TokensForObject function to create properly formatted object tokens
	return hclwrite.TokensForObject(attrs)
}

// substituteVariables replaces module variables with their actual values from the module call
func substituteVariables(attr *hclwrite.Attribute, moduleCall *ModuleCall) hclwrite.Tokens {
	if attr == nil {
		return nil
	}

	// Get the original tokens
	tokens := attr.Expr().BuildTokens(nil)
	originalStr := string(tokens.Bytes())

	// Special case: if it's exactly a variable reference, return the substitution directly
	varPattern := regexp.MustCompile(`^\s*var\.([a-zA-Z_][a-zA-Z0-9_]*)\s*$`)
	if match := varPattern.FindStringSubmatch(originalStr); match != nil {
		varName := match[1]
		if argAttr, exists := moduleCall.Arguments[varName]; exists {
			return argAttr.Expr().BuildTokens(nil)
		}
		// If no module argument exists, this means we shouldn't be generating this resource
		// but if we get here, it's an error in our filtering logic
		return tokens
	}

	// For more complex expressions, try pattern replacement
	generalVarPattern := regexp.MustCompile(`var\.([a-zA-Z_][a-zA-Z0-9_]*)`)
	result := generalVarPattern.ReplaceAllStringFunc(originalStr, func(match string) string {
		varName := match[4:] // Remove "var."
		
		if argAttr, exists := moduleCall.Arguments[varName]; exists {
			argTokens := argAttr.Expr().BuildTokens(nil)
			return string(argTokens.Bytes())
		}
		
		return match
	})

	// If no changes were made, return original tokens
	if result == originalStr {
		return tokens
	}

	// For changed expressions, try to parse them correctly
	// This is a simplified approach - we could improve this by properly parsing HCL
	if strings.HasPrefix(result, `"`) && strings.HasSuffix(result, `"`) {
		// It's a string literal
		return hclwrite.TokensForValue(cty.StringVal(strings.Trim(result, `"`)))
	}

	// Fall back to original tokens for complex cases
	return tokens
}

// Helper functions that mirror cmd/migrate functionality

// buildResourceReference builds tokens for a resource reference like "cloudflare_zone_setting.resource_name"
func buildResourceReference(resourceType, resourceName string) hclwrite.Tokens {
	return hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte(resourceType)},
		{Type: hclsyntax.TokenDot, Bytes: []byte{'.'}},
		{Type: hclsyntax.TokenIdent, Bytes: []byte(resourceName)},
	}
}

// buildTemplateStringTokens builds tokens for a template string like "${zone_id}/setting_id"
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