package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isZeroTrustAccessMTLSHostnameSettingsResource checks if a block is a cloudflare_zero_trust_access_mtls_hostname_settings resource
func isZeroTrustAccessMTLSHostnameSettingsResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_mtls_hostname_settings"
}

// transformZeroTrustAccessMTLSHostnameSettingsBlock transforms settings blocks to attributes
// Handles the v4 to v5 migration:
// V4: settings { hostname = "example.com"; china_network = false; client_certificate_forwarding = true }
// V5: settings = [{ hostname = "example.com", china_network = false, client_certificate_forwarding = true }]
func transformZeroTrustAccessMTLSHostnameSettingsBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Look for settings blocks and convert them to a settings attribute with list value
	settingsBlocks := []*hclwrite.Block{}
	body := block.Body()

	// Collect all settings blocks
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "settings" {
			settingsBlocks = append(settingsBlocks, childBlock)
		}
	}

	if len(settingsBlocks) == 0 {
		return // No settings blocks to convert
	}

	// Convert settings blocks to a list attribute
	var settingsObjects []cty.Value

	for _, settingsBlock := range settingsBlocks {
		settingsBody := settingsBlock.Body()
		objectMap := make(map[string]cty.Value)

		// Extract attributes from the settings block
		for name, attr := range settingsBody.Attributes() {
			expr := attr.Expr()
			
			// Convert the expression to cty.Value
			switch name {
			case "hostname":
				if val := extractStringValue(*expr); val != "" {
					objectMap["hostname"] = cty.StringVal(val)
				}
			case "china_network":
				if val, ok := extractBoolValue(*expr); ok {
					objectMap["china_network"] = cty.BoolVal(val)
				} else {
					// Default to false if not present or cannot parse
					objectMap["china_network"] = cty.BoolVal(false)
				}
			case "client_certificate_forwarding":
				if val, ok := extractBoolValue(*expr); ok {
					objectMap["client_certificate_forwarding"] = cty.BoolVal(val)
				} else {
					// Default to false if not present or cannot parse
					objectMap["client_certificate_forwarding"] = cty.BoolVal(false)
				}
			}
		}

		// Ensure required boolean attributes have defaults
		if _, exists := objectMap["china_network"]; !exists {
			objectMap["china_network"] = cty.BoolVal(false)
		}
		if _, exists := objectMap["client_certificate_forwarding"]; !exists {
			objectMap["client_certificate_forwarding"] = cty.BoolVal(false)
		}

		if len(objectMap) > 0 {
			settingsObjects = append(settingsObjects, cty.ObjectVal(objectMap))
		}
	}

	if len(settingsObjects) > 0 {
		// Create the settings list attribute
		listValue := cty.ListVal(settingsObjects)
		body.SetAttributeValue("settings", listValue)

		// Remove all settings blocks
		for _, settingsBlock := range settingsBlocks {
			body.RemoveBlock(settingsBlock)
		}
	}
}

// extractStringValue extracts a string value from an HCL expression
func extractStringValue(expr hclwrite.Expression) string {
	// Try to extract string literal by building tokens and concatenating
	tokens := expr.BuildTokens(nil)
	if len(tokens) == 0 {
		return ""
	}
	
	// Concatenate all token bytes to get the full value
	var result string
	for _, token := range tokens {
		result += string(token.Bytes)
	}
	
	// Remove surrounding quotes if present
	if len(result) >= 2 && result[0] == '"' && result[len(result)-1] == '"' {
		return result[1 : len(result)-1]
	}
	return result
}

// extractBoolValue extracts a boolean value from an HCL expression
func extractBoolValue(expr hclwrite.Expression) (bool, bool) {
	tokens := expr.BuildTokens(nil)
	if len(tokens) >= 1 {
		val := string(tokens[0].Bytes)
		switch val {
		case "true":
			return true, true
		case "false":
			return false, true
		}
	}
	return false, false
}