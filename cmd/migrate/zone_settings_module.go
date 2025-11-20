package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Zone setting mapping from the zone_settings module
var zoneSettingMappings = map[string]string{
	"always_use_https":             "always_use_https",
	"automatic_https_rewrites":     "automatic_https_rewrites",
	"browser_check":                "browser_check",
	"challenge_ttl":                "challenge_ttl",
	"tls_client_auth":              "tls_client_auth",
	"ciphers":                      "ciphers",
	"http2":                        "http2",
	"http3":                        "http3",
	"origin_max_http_version":      "origin_max_http_version",
	"h2_prioritization":            "h2_prioritization",
	"zero_rtt":                     "0rtt", // Note: v4 uses zero_rtt but API expects 0rtt
	"always_online":                "always_online",
	"ipv6":                         "ipv6",
	"max_upload":                   "max_upload",
	"websockets":                   "websockets",
	"proxy_read_timeout":           "proxy_read_timeout",
	"security_level":               "security_level",
	"ip_geolocation":               "ip_geolocation",
	"email_obfuscation":            "email_obfuscation",
	"polish":                       "polish",
	"mirage":                       "mirage",
	"image_resizing":               "image_resizing",
	"fonts":                        "fonts",
	"binary_ast":                   "binary_ast",
	"min_tls_version":              "min_tls_version",
	"tls_1_3":                      "tls_1_3",
	"ssl":                          "ssl",
	"origin_error_page_pass_thru":  "origin_error_page_pass_thru",
}

// Special settings that need complex transformation
var complexSettings = map[string]bool{
	"security_header_enabled":              true,
	"security_header_include_subdomains":   true,
	"security_header_max_age":              true,
	"security_header_nosniff":              true,
	"security_header_preload":              true,
	"enable_network_error_logging":         true,
}

// expandZoneSettingsModules finds zone_settings module calls and expands them to individual resources
func expandZoneSettingsModules(content string, skipImports bool) string {
	// Pattern to match module "zone_settings" blocks
	modulePattern := `module\s+"zone_settings"\s*\{[^}]*\}`
	re := regexp.MustCompile(modulePattern)

	// Find all zone_settings module blocks
	matches := re.FindAllString(content, -1)

	for _, match := range matches {
		expanded := expandSingleZoneSettingsModule(match, skipImports)
		if expanded != "" {
			content = strings.Replace(content, match, expanded, 1)
		}
	}

	return content
}

// expandSingleZoneSettingsModule expands a single zone_settings module call
func expandSingleZoneSettingsModule(moduleBlock string, skipImports bool) string {
	// Parse the module block to extract attributes
	file, diags := hclwrite.ParseConfig([]byte(moduleBlock), "module.tf", hcl.InitialPos)
	if diags.HasErrors() {
		fmt.Printf("Warning: Could not parse zone_settings module block: %v\n", diags)
		return moduleBlock // Return original if can't parse
	}

	// Find the module block
	var moduleBlockParsed *hclwrite.Block
	for _, block := range file.Body().Blocks() {
		if block.Type() == "module" && len(block.Labels()) > 0 && block.Labels()[0] == "zone_settings" {
			moduleBlockParsed = block
			break
		}
	}

	if moduleBlockParsed == nil {
		return moduleBlock // Return original if can't find module
	}

	var results []string

	// Extract zone_id attribute
	var zoneIdAttr *hclwrite.Attribute
	if attr := moduleBlockParsed.Body().GetAttribute("zone_id"); attr != nil {
		zoneIdAttr = attr
	}

	// Generate unique resource name prefix
	resourcePrefix := "zone_settings"

	// Process each attribute in the module
	for _, attr := range AttributesOrdered(moduleBlockParsed.Body()) {
		attrName := attr.Name

		// Skip source attribute
		if attrName == "source" {
			continue
		}

		// Skip zone_id as it's used for all resources
		if attrName == "zone_id" {
			continue
		}

		// Handle special complex settings
		if strings.HasPrefix(attrName, "security_header_") {
			// Group all security_header_* attributes
			continue // Will be handled separately
		}

		if attrName == "enable_network_error_logging" {
			// Handle NEL setting
			settingId := "nel"
			resourceName := fmt.Sprintf("%s_%s", resourcePrefix, "nel")

			resource := createZoneSettingResourceFromModule(resourceName, settingId, zoneIdAttr, attr.Attribute, true)
			if resource != "" { // Only add if resource was created (not skipped for null values)
				results = append(results, resource)

				if !skipImports {
					importBlock := createImportBlockFromModule(resourceName, settingId, zoneIdAttr)
					results = append(results, importBlock)
				}
			}
			continue
		}

		// Handle standard settings
		if settingId, exists := zoneSettingMappings[attrName]; exists {
			resourceName := fmt.Sprintf("%s_%s", resourcePrefix, attrName)

			resource := createZoneSettingResourceFromModule(resourceName, settingId, zoneIdAttr, attr.Attribute, false)
			if resource != "" { // Only add if resource was created (not skipped for null values)
				results = append(results, resource)

				if !skipImports {
					importBlock := createImportBlockFromModule(resourceName, settingId, zoneIdAttr)
					results = append(results, importBlock)
				}
			}
		}
	}

	// Handle security_header settings as a group
	securityHeaderAttrs := extractSecurityHeaderAttributes(moduleBlockParsed.Body())
	if len(securityHeaderAttrs) > 0 {
		resourceName := fmt.Sprintf("%s_security_header", resourcePrefix)
		settingId := "security_header"

		resource := createSecurityHeaderResourceFromModule(resourceName, settingId, zoneIdAttr, securityHeaderAttrs)
		results = append(results, resource)

		if !skipImports {
			importBlock := createImportBlockFromModule(resourceName, settingId, zoneIdAttr)
			results = append(results, importBlock)
		}
	}

	return strings.Join(results, "\n\n")
}

// extractSecurityHeaderAttributes extracts all security_header_* attributes
func extractSecurityHeaderAttributes(body *hclwrite.Body) map[string]*hclwrite.Attribute {
	attrs := make(map[string]*hclwrite.Attribute)

	for _, attr := range AttributesOrdered(body) {
		if strings.HasPrefix(attr.Name, "security_header_") {
			// Map attribute names to security header field names
			fieldName := strings.TrimPrefix(attr.Name, "security_header_")
			attrs[fieldName] = attr.Attribute
		}
	}

	return attrs
}

// createZoneSettingResourceFromModule creates a zone_setting resource from module attributes
func createZoneSettingResourceFromModule(resourceName, settingId string, zoneIdAttr, valueAttr *hclwrite.Attribute, isNEL bool) string {
	// Check if the value is null - if so, don't create the resource at all
	if valueAttr != nil {
		valueStr := tokensToString(valueAttr.Expr().BuildTokens(nil))
		if strings.TrimSpace(valueStr) == "null" {
			fmt.Printf("SKIPPED: cloudflare_zone_setting.%s (reason: value = null)\n", resourceName)
			fmt.Printf("         (setting_id: %s has no meaningful value)\n", settingId)
			return "" // Return empty string to indicate no resource should be created
		}
	}

	var lines []string
	lines = append(lines, fmt.Sprintf(`resource "cloudflare_zone_setting" "%s" {`, resourceName))

	// Add zone_id
	if zoneIdAttr != nil {
		zoneIdStr := tokensToString(zoneIdAttr.Expr().BuildTokens(nil))
		lines = append(lines, fmt.Sprintf("  zone_id    = %s", zoneIdStr))
	}

	// Add setting_id
	lines = append(lines, fmt.Sprintf(`  setting_id = "%s"`, settingId))

	// Add value
	if valueAttr != nil {
		if isNEL {
			// NEL requires wrapping the boolean value in an object with "enabled" field
			valueStr := tokensToString(valueAttr.Expr().BuildTokens(nil))
			lines = append(lines, fmt.Sprintf("  value      = { enabled = %s }", valueStr))
		} else {
			valueStr := tokensToString(valueAttr.Expr().BuildTokens(nil))
			lines = append(lines, fmt.Sprintf("  value      = %s", valueStr))
		}
	}

	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}

// createSecurityHeaderResourceFromModule creates a security_header zone_setting resource
func createSecurityHeaderResourceFromModule(resourceName, settingId string, zoneIdAttr *hclwrite.Attribute, securityAttrs map[string]*hclwrite.Attribute) string {
	var lines []string
	lines = append(lines, fmt.Sprintf(`resource "cloudflare_zone_setting" "%s" {`, resourceName))

	// Add zone_id
	if zoneIdAttr != nil {
		zoneIdStr := tokensToString(zoneIdAttr.Expr().BuildTokens(nil))
		lines = append(lines, fmt.Sprintf("  zone_id    = %s", zoneIdStr))
	}

	// Add setting_id
	lines = append(lines, fmt.Sprintf(`  setting_id = "%s"`, settingId))

	// Build security header object
	var valueLines []string
	valueLines = append(valueLines, "  value = {")
	valueLines = append(valueLines, "    strict_transport_security = {")

	// Map the security header attributes in deterministic order
	var fieldNames []string
	for fieldName := range securityAttrs {
		fieldNames = append(fieldNames, fieldName)
	}
	sort.Strings(fieldNames)

	for _, fieldName := range fieldNames {
		attr := securityAttrs[fieldName]
		valueStr := tokensToString(attr.Expr().BuildTokens(nil))
		valueLines = append(valueLines, fmt.Sprintf("      %s = %s", fieldName, valueStr))
	}

	valueLines = append(valueLines, "    }")
	valueLines = append(valueLines, "  }")

	lines = append(lines, strings.Join(valueLines, "\n"))
	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}

// createImportBlockFromModule creates an import block for zone_setting resources
func createImportBlockFromModule(resourceName, settingId string, zoneIdAttr *hclwrite.Attribute) string {
	var lines []string
	lines = append(lines, "import {")
	lines = append(lines, fmt.Sprintf(`  to = cloudflare_zone_setting.%s`, resourceName))

	if zoneIdAttr != nil {
		zoneIdStr := tokensToString(zoneIdAttr.Expr().BuildTokens(nil))
		lines = append(lines, fmt.Sprintf(`  id = "${%s}/%s"`, zoneIdStr, settingId))
	}

	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}

// tokensToString converts HCL tokens to a string representation
func tokensToString(tokens hclwrite.Tokens) string {
	var parts []string
	for _, token := range tokens {
		parts = append(parts, string(token.Bytes))
	}
	return strings.Join(parts, "")
}
