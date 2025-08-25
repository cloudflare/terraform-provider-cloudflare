package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// mapSettingName translates v4 setting names to v5 setting names
// This handles cases where the v4 provider used different names internally
func mapSettingName(v4Name string) string {
	settingNameMap := map[string]string{
		"zero_rtt": "0rtt", // v4 used "zero_rtt" but API expects "0rtt"
	}
	
	if v5Name, exists := settingNameMap[v4Name]; exists {
		return v5Name
	}
	return v4Name
}

// isZoneSettingsOverrideResource checks if a block is a cloudflare_zone_settings_override resource
func isZoneSettingsOverrideResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zone_settings_override"
}

// transformZoneSettingsBlock transforms a zone settings override block into individual zone setting resources
func transformZoneSettingsBlock(oldBlock *hclwrite.Block) []*hclwrite.Block {
	var newBlocks []*hclwrite.Block

	// Get the resource name from the old block
	resourceName := oldBlock.Labels()[1]

	// Get zone_id attribute from the old block
	var zoneIDAttr *hclwrite.Attribute
	if attr := oldBlock.Body().GetAttribute("zone_id"); attr != nil {
		zoneIDAttr = attr
	}

	// Find the settings block
	for _, settingsBlock := range oldBlock.Body().Blocks() {
		if settingsBlock.Type() == "settings" {
			// Process regular attributes - migrate ALL attributes dynamically
			for name, attr := range settingsBlock.Body().Attributes() {
				// Map the v4 setting name to the correct v5 setting name
				mappedSettingName := mapSettingName(name)
				resourceFullName := resourceName + "_" + name
				newBlock := createZoneSettingResource(
					resourceFullName,
					mappedSettingName, // Use the mapped setting name as the setting_id
					zoneIDAttr,
					attr,
				)
				newBlocks = append(newBlocks, newBlock)
				
				// Create import block for this resource
				importBlock := createImportBlock(resourceFullName, mappedSettingName, zoneIDAttr)
				newBlocks = append(newBlocks, importBlock)
			}

			// Process nested blocks (security_header, nel)
			for _, nestedBlock := range settingsBlock.Body().Blocks() {
				if nestedBlock.Type() == "security_header" {
					resourceFullName := resourceName + "_security_header"
					newBlocks = append(newBlocks, transformSecurityHeaderBlock(resourceName, zoneIDAttr, nestedBlock))
					// Create import block for security_header
					importBlock := createImportBlock(resourceFullName, "security_header", zoneIDAttr)
					newBlocks = append(newBlocks, importBlock)
				} else if nestedBlock.Type() == "nel" {
					resourceFullName := resourceName + "_nel"
					newBlocks = append(newBlocks, transformNELBlock(resourceName, zoneIDAttr, nestedBlock))
					// Create import block for nel
					importBlock := createImportBlock(resourceFullName, "nel", zoneIDAttr)
					newBlocks = append(newBlocks, importBlock)
				}
			}
		}
	}

	return newBlocks
}

// createZoneSettingResource creates a new cloudflare_zone_setting resource block
func createZoneSettingResource(name, settingID string, zoneIDAttr, valueAttr *hclwrite.Attribute) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_zone_setting", name})
	body := block.Body()

	// Set zone_id with the expression from the original attribute
	if zoneIDAttr != nil {
		tokens := zoneIDAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("zone_id", tokens)
	}

	// Set setting_id
	body.SetAttributeValue("setting_id", cty.StringVal(settingID))

	// Set value with the expression from the original attribute
	if valueAttr != nil {
		tokens := valueAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("value", tokens)
	}

	return block
}

// transformSecurityHeaderBlock transforms a security_header block into a zone setting resource
func transformSecurityHeaderBlock(baseName string, zoneIDAttr *hclwrite.Attribute, securityHeaderBlock *hclwrite.Block) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_zone_setting", baseName + "_security_header"})
	body := block.Body()

	// Set zone_id with the expression from the original attribute
	if zoneIDAttr != nil {
		tokens := zoneIDAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("zone_id", tokens)
	}

	// Set setting_id
	body.SetAttributeValue("setting_id", cty.StringVal("security_header"))

	// Build the object tokens manually to preserve variable references
	objectTokens := buildObjectFromBlock(securityHeaderBlock)
	body.SetAttributeRaw("value", objectTokens)

	return block
}

// transformNELBlock transforms a nel block into a zone setting resource
func transformNELBlock(baseName string, zoneIDAttr *hclwrite.Attribute, nelBlock *hclwrite.Block) *hclwrite.Block {
	block := hclwrite.NewBlock("resource", []string{"cloudflare_zone_setting", baseName + "_nel"})
	body := block.Body()

	// Set zone_id with the expression from the original attribute
	if zoneIDAttr != nil {
		tokens := zoneIDAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("zone_id", tokens)
	}

	// Set setting_id
	body.SetAttributeValue("setting_id", cty.StringVal("nel"))

	// Build the object tokens manually to preserve variable references
	objectTokens := buildObjectFromBlock(nelBlock)
	body.SetAttributeRaw("value", objectTokens)

	return block
}

// buildObjectFromBlock creates object tokens from a block's attributes
func buildObjectFromBlock(block *hclwrite.Block) hclwrite.Tokens {
	// Get attributes in their original order
	orderedAttrs := AttributesOrdered(block.Body())
	
	// Build a list of attribute tokens preserving the original order
	var attrs []hclwrite.ObjectAttrTokens
	
	for _, attrInfo := range orderedAttrs {
		// Create tokens for the attribute name (as a simple identifier)
		nameTokens := hclwrite.TokensForIdentifier(attrInfo.Name)
		
		// Get the value tokens from the attribute's expression
		valueTokens := attrInfo.Attribute.Expr().BuildTokens(nil)
		
		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  nameTokens,
			Value: valueTokens,
		})
	}
	
	// Use the built-in TokensForObject function to create properly formatted object tokens
	return hclwrite.TokensForObject(attrs)
}

// createImportBlock creates an import block for a zone setting resource
func createImportBlock(resourceName, settingID string, zoneIDAttr *hclwrite.Attribute) *hclwrite.Block {
	block := hclwrite.NewBlock("import", nil)
	body := block.Body()

	// Build the "to" value: cloudflare_zone_setting.resource_name
	toTokens := buildResourceReference("cloudflare_zone_setting", resourceName)
	body.SetAttributeRaw("to", toTokens)

	// Build the "id" value: "${zone_id}/setting_id"
	if zoneIDAttr != nil {
		zoneIDTokens := zoneIDAttr.Expr().BuildTokens(nil)
		idTokens := buildTemplateStringTokens(zoneIDTokens, "/" + settingID)
		body.SetAttributeRaw("id", idTokens)
	}

	return block
}
