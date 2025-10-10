package cloudflare_zone_settings_override

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/hcl"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/zclconf/go-cty/cty"
)

var _ interfaces.ResourceTransformer = (*ZoneSettingsOverride)(nil)

type ZoneSettingsOverride struct {
	skipImports bool
}

func NewZoneSettingsOverride() interfaces.ResourceTransformer {
	return &ZoneSettingsOverride{skipImports: false}
}

func (z *ZoneSettingsOverride) CanHandle(resourceType string) bool {
	return resourceType == "cloudflare_zone_settings_override"
}

func (z *ZoneSettingsOverride) GetResourceType() string {
	return "cloudflare_zone_settings_override"
}

// Preprocess handles string-level transformations before HCL parsing
func (z *ZoneSettingsOverride) Preprocess(content string) string {
	// Zone settings override doesn't need any string-level preprocessing
	// All transformations are handled at the AST level in TransformConfig()
	// This method returns the content unchanged
	return content
}

// SetSkipImports allows configuring whether to generate import blocks
func (z *ZoneSettingsOverride) SetSkipImports(skip bool) {
	z.skipImports = skip
}

// Transform implements ResourceTransformer interface
func (z *ZoneSettingsOverride) TransformConfig(block *hclwrite.Block) (*interfaces.TransformResult, error) {
	// Call the wrapper function that contains all logic from cmd/migrate/zone_settings.go
	newBlocks := transformZoneSettingsBlock(block, z.skipImports)

	// Zone settings override always removes the original and returns new blocks
	return &interfaces.TransformResult{
		Blocks:         newBlocks,
		RemoveOriginal: true,
	}, nil
}

func (z *ZoneSettingsOverride) TransformState(json gjson.Result, resourcePath string) (string, error) {
	// Call the wrapper function that replicates logic from cmd/migrate
	jsonStr := json.String()
	result := transformZoneSettingsStateJSON(jsonStr, resourcePath)
	return result, nil
}

// isDeprecatedSetting checks if a setting is deprecated and should be skipped
func isDeprecatedSetting(settingName string) bool {
	deprecatedSettings := map[string]bool{
		"universal_ssl": true, // No longer exists in zone settings API
	}
	return deprecatedSettings[settingName]
}

// mapSettingName translates v4 setting names to v5 setting names
func mapSettingName(v4Name string) string {
	settingNameMap := map[string]string{
		"zero_rtt": "0rtt", // v4 used "zero_rtt" but API expects "0rtt"
	}

	if v5Name, exists := settingNameMap[v4Name]; exists {
		return v5Name
	}
	return v4Name
}

// transformZoneSettingsBlock transforms a zone settings override block into individual zone setting resources
func transformZoneSettingsBlock(oldBlock *hclwrite.Block, skipImports bool) []*hclwrite.Block {
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
			// Process regular attributes in deterministic order
			for _, attrInfo := range hcl.AttributesOrdered(settingsBlock.Body()) {
				// Skip deprecated settings that no longer exist in the API
				if isDeprecatedSetting(attrInfo.Name) {
					continue
				}

				// Map the v4 setting name to the correct v5 setting name
				mappedSettingName := mapSettingName(attrInfo.Name)
				resourceFullName := resourceName + "_" + attrInfo.Name
				newBlock := createZoneSettingResource(
					resourceFullName,
					mappedSettingName, // Use the mapped setting name as the setting_id
					zoneIDAttr,
					attrInfo.Attribute,
				)
				newBlocks = append(newBlocks, newBlock)

				// Create import block for this resource if not skipping imports
				if !skipImports {
					importBlock := createImportBlock(resourceFullName, mappedSettingName, zoneIDAttr)
					newBlocks = append(newBlocks, importBlock)
				}
			}

			// Process nested blocks (security_header, nel)
			for _, nestedBlock := range settingsBlock.Body().Blocks() {
				if nestedBlock.Type() == "security_header" {
					resourceFullName := resourceName + "_security_header"
					newBlocks = append(newBlocks, transformSecurityHeaderBlock(resourceName, zoneIDAttr, nestedBlock))
					// Create import block for security_header if not skipping imports
					if !skipImports {
						importBlock := createImportBlock(resourceFullName, "security_header", zoneIDAttr)
						newBlocks = append(newBlocks, importBlock)
					}
				} else if nestedBlock.Type() == "nel" {
					resourceFullName := resourceName + "_nel"
					newBlocks = append(newBlocks, transformNELBlock(resourceName, zoneIDAttr, nestedBlock))
					// Create import block for nel if not skipping imports
					if !skipImports {
						importBlock := createImportBlock(resourceFullName, "nel", zoneIDAttr)
						newBlocks = append(newBlocks, importBlock)
					}
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
	// Security header needs to be wrapped in strict_transport_security for v5 API
	innerObjectTokens := hcl.BuildObjectFromBlock(securityHeaderBlock)

	// Create the wrapper object with strict_transport_security key
	wrapperTokens := []*hclwrite.Token{
		{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
		{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
		{Type: hclsyntax.TokenIdent, Bytes: []byte("strict_transport_security")},
		{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")},
	}
	wrapperTokens = append(wrapperTokens, innerObjectTokens...)
	wrapperTokens = append(wrapperTokens,
		[]*hclwrite.Token{
			{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
		}...)

	body.SetAttributeRaw("value", wrapperTokens)

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
	objectTokens := hcl.BuildObjectFromBlock(nelBlock)
	body.SetAttributeRaw("value", objectTokens)

	return block
}

// createImportBlock creates an import block for a zone setting resource
func createImportBlock(resourceName, settingID string, zoneIDAttr *hclwrite.Attribute) *hclwrite.Block {
	block := hclwrite.NewBlock("import", nil)
	body := block.Body()

	// Build the "to" value: cloudflare_zone_setting.resource_name
	toTokens := hcl.BuildResourceReference("cloudflare_zone_setting", resourceName)
	body.SetAttributeRaw("to", toTokens)

	// Build the "id" value: "${zone_id}/setting_id"
	if zoneIDAttr != nil {
		zoneIDTokens := zoneIDAttr.Expr().BuildTokens(nil)
		idTokens := hcl.BuildTemplateStringTokens(zoneIDTokens, "/"+settingID)
		body.SetAttributeRaw("id", idTokens)
	}

	return block
}

// transformZoneSettingsStateJSON handles zone settings state transformation
func transformZoneSettingsStateJSON(jsonData, resourcePath string) string {
	// Delete the zone_settings_override resource - import blocks will handle state creation
	result, _ := sjson.Delete(jsonData, resourcePath)
	return result
}
