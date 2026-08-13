package main

import (
	"strings"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isZeroTrustAccessMTLSHostnameSettingsResource checks if a block is a cloudflare_zero_trust_access_mtls_hostname_settings resource
func isZeroTrustAccessMTLSHostnameSettingsResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_mtls_hostname_settings"
}

// isAccessMutualTLSHostnameSettingsResource checks if a block is a cloudflare_access_mutual_tls_hostname_settings resource (v4 old name)
func isAccessMutualTLSHostnameSettingsResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_access_mutual_tls_hostname_settings"
}

// transformZeroTrustAccessMTLSHostnameSettingsBlock transforms settings blocks to attributes
// Handles the v4 to v5 migration:
// V4: settings { hostname = "example.com"; china_network = false; client_certificate_forwarding = true }
// V5: settings = [{ hostname = "example.com", china_network = false, client_certificate_forwarding = true }]
// Dynamic blocks are converted to for expressions:
// V4: dynamic "settings" { for_each = local.domains; content { hostname = settings.value } }
// V5: settings = [for value in local.domains : { hostname = value, china_network = false, client_certificate_forwarding = false }]
func transformZeroTrustAccessMTLSHostnameSettingsBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Look for settings blocks and convert them to a settings attribute with list value
	settingsBlocks := []*hclwrite.Block{}
	dynamicSettingsBlocks := []*hclwrite.Block{}
	body := block.Body()

	// Collect all settings blocks (both static and dynamic)
	for _, childBlock := range body.Blocks() {
		if childBlock.Type() == "settings" {
			settingsBlocks = append(settingsBlocks, childBlock)
		} else if childBlock.Type() == "dynamic" && len(childBlock.Labels()) > 0 && childBlock.Labels()[0] == "settings" {
			dynamicSettingsBlocks = append(dynamicSettingsBlocks, childBlock)
		}
	}

	// Handle dynamic blocks - convert them to for expressions
	if len(dynamicSettingsBlocks) > 0 {
		for _, dynBlock := range dynamicSettingsBlocks {
			// Extract the for_each expression
			forEachAttr := dynBlock.Body().GetAttribute("for_each")
			if forEachAttr == nil {
				continue
			}
			
			// Get the iterator name (defaults to the block label if not specified)
			iteratorName := "settings"
			if iteratorAttr := dynBlock.Body().GetAttribute("iterator"); iteratorAttr != nil {
				// Extract iterator name from the expression
				tokens := iteratorAttr.Expr().BuildTokens(nil)
				if len(tokens) > 0 {
					iteratorName = string(tokens[0].Bytes)
					// Remove quotes if present
					if len(iteratorName) >= 2 && iteratorName[0] == '"' && iteratorName[len(iteratorName)-1] == '"' {
						iteratorName = iteratorName[1 : len(iteratorName)-1]
					}
				}
			}
			
			// Extract content block
			var contentBlock *hclwrite.Block
			for _, cb := range dynBlock.Body().Blocks() {
				if cb.Type() == "content" {
					contentBlock = cb
					break
				}
			}
			
			if contentBlock == nil {
				continue
			}
			
			// Extract field values from content block
			hostnameTokens := hclwrite.Tokens{}
			chinaNetworkVal := "false"
			clientCertForwardingVal := "false"
			
			if attr := contentBlock.Body().GetAttribute("hostname"); attr != nil {
				hostnameTokens = attr.Expr().BuildTokens(nil)
			}
			
			if attr := contentBlock.Body().GetAttribute("china_network"); attr != nil {
				if val, ok := extractBoolValue(*attr.Expr()); ok {
					if val {
						chinaNetworkVal = "true"
					}
				}
			}
			
			if attr := contentBlock.Body().GetAttribute("client_certificate_forwarding"); attr != nil {
				if val, ok := extractBoolValue(*attr.Expr()); ok {
					if val {
						clientCertForwardingVal = "true"
					}
				}
			}
			
			// Build the for expression as raw tokens
			// We want: [for value in <for_each_expr> : { hostname = <hostname_expr>, china_network = <val>, client_certificate_forwarding = <val> }]
			tokens := hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("for")},
				&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")},
				&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("in")},
			}
			
			// Add the for_each expression tokens
			tokens = append(tokens, forEachAttr.Expr().BuildTokens(nil)...)
			
			// Add the colon
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenColon, Bytes: []byte(":")})
			
			// Start object
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			
			// Add hostname field
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    hostname")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte("=")})
			
			// If hostname references the iterator, replace it with "value"
			if len(hostnameTokens) > 0 {
				// Check if it's using the iterator variable (e.g., "settings.value")
				// We need to replace "settings.value" with just "value"
				isIteratorRef := false
				for i, token := range hostnameTokens {
					if token.Type == hclsyntax.TokenIdent && string(token.Bytes) == iteratorName {
						// Found the iterator reference, check if it's followed by .value
						if i+1 < len(hostnameTokens) && hostnameTokens[i+1].Type == hclsyntax.TokenDot &&
							i+2 < len(hostnameTokens) && string(hostnameTokens[i+2].Bytes) == "value" {
							// It's "iterator.value", replace with just "value"
							tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
							isIteratorRef = true
							break
						}
					}
				}
				
				if !isIteratorRef {
					// Use the original hostname expression
					tokens = append(tokens, hostnameTokens...)
				}
			} else {
				tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("value")})
			}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			
			// Add china_network field
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    china_network")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte("=")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(chinaNetworkVal)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			
			// Add client_certificate_forwarding field
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    client_certificate_forwarding")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte("=")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(clientCertForwardingVal)})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
			
			// Close object
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")})
			
			// Set the settings attribute with the for expression
			body.SetAttributeRaw("settings", tokens)
			
			// Remove the dynamic block
			body.RemoveBlock(dynBlock)
		}
		return
	}

	if len(settingsBlocks) == 0 {
		return // No settings blocks to convert
	}

	// Convert static settings blocks to a list attribute
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
	
	// Remove surrounding quotes if present and unescape
	if len(result) >= 2 && result[0] == '"' && result[len(result)-1] == '"' {
		unquoted := result[1 : len(result)-1]
		// Unescape common escape sequences
		unquoted = strings.ReplaceAll(unquoted, `\"`, `"`)
		unquoted = strings.ReplaceAll(unquoted, `\\`, `\`)
		unquoted = strings.ReplaceAll(unquoted, `\n`, "\n")
		unquoted = strings.ReplaceAll(unquoted, `\r`, "\r")
		unquoted = strings.ReplaceAll(unquoted, `\t`, "\t")
		return unquoted
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