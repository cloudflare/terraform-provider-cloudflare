package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isZeroTrustAccessIdentityProviderResource checks if a block is an access identity provider resource
// Handles both old (cloudflare_access_identity_provider) and new (cloudflare_zero_trust_access_identity_provider) names
func isZeroTrustAccessIdentityProviderResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		(block.Labels()[0] == "cloudflare_zero_trust_access_identity_provider" ||
			block.Labels()[0] == "cloudflare_access_identity_provider")
}

// transformZeroTrustAccessIdentityProviderBlock transforms the identity provider configuration
// Handles:
// 1. Resource name: access_identity_provider -> zero_trust_access_identity_provider
// 2. config block -> config object conversion (done by grit)
// 3. scim_config block -> scim_config object conversion (done by grit) 
// 4. idp_public_cert -> idp_public_certs field rename and type conversion (string -> list)
// 5. Remove deprecated fields: api_token, group_member_deprovision
// 6. Remove invalid attributes based on provider type validation rules
// 7. Ensure config object exists even for OneTimePin providers
func transformZeroTrustAccessIdentityProviderBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource renaming from v4 to v5
	if block.Labels()[0] == "cloudflare_access_identity_provider" {
		block.SetLabels([]string{"cloudflare_zero_trust_access_identity_provider", block.Labels()[1]})
	}

	// Convert config and scim_config blocks to objects (if they exist as blocks)
	convertConfigBlocksToObjects(block, diags)

	// Ensure config object exists (required in v5, even if empty for OneTimePin)
	ensureConfigObjectExists(block)

	// Get the provider type for validation rules
	providerType := getProviderType(block)
	// Debug: temporary logging to see what provider type is being detected
	// fmt.Printf("DEBUG: transformZeroTrustAccessIdentityProviderBlock called with provider type: '%s'\n", providerType)
	
	// Apply config-specific transformations
	transforms := map[string]ast.ExprTransformer{
		"config":      func(expr *hclsyntax.Expression, diags ast.Diagnostics) { transformConfigObject(expr, diags, providerType) },
		"scim_config": transformScimConfigObject,
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

// ensureConfigObjectExists ensures that the config attribute exists as an empty object
// This is needed because v5 requires config to be present, even for OneTimePin providers
func ensureConfigObjectExists(block *hclwrite.Block) {
	configAttr := block.Body().GetAttribute("config")
	configBlocks := block.Body().Blocks()
	
	// Check if config exists as a block (v4 style)
	hasConfigBlock := false
	for _, b := range configBlocks {
		if b.Type() == "config" {
			hasConfigBlock = true
			break
		}
	}
	
	// If config doesn't exist as either attribute or block, create empty config object
	if configAttr == nil && !hasConfigBlock {
		tokens := hclwrite.Tokens{
			{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
			{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
		}
		block.Body().SetAttributeRaw("config", tokens)
	}
}

// getProviderType extracts the provider type from the block
func getProviderType(block *hclwrite.Block) string {
	if block == nil || block.Body() == nil {
		return ""
	}
	
	if typeAttr := block.Body().GetAttribute("type"); typeAttr != nil && typeAttr.Expr() != nil {
		tokens := typeAttr.Expr().BuildTokens(nil)
		if len(tokens) >= 3 {
			// For quoted strings, tokens are: [quote, content, quote]
			// Take the middle token which contains the actual content
			typeValue := string(tokens[1].Bytes)
			return typeValue
		} else if len(tokens) == 1 {
			// For unquoted identifiers
			typeValue := string(tokens[0].Bytes)
			// Remove quotes if they're part of the value
			if len(typeValue) >= 2 && typeValue[0] == '"' && typeValue[len(typeValue)-1] == '"' {
				return typeValue[1 : len(typeValue)-1]
			}
			return typeValue
		}
	}
	return ""
}

// transformConfigObject handles transformations within the config object
// 1. idp_public_cert -> idp_public_certs (string -> list[string])
// 2. Remove deprecated api_token field
// 3. Remove invalid attributes based on provider type validation rules
func transformConfigObject(expr *hclsyntax.Expression, diags ast.Diagnostics, providerType string) {
	if *expr == nil {
		return
	}

	obj, ok := (*expr).(*hclsyntax.ObjectConsExpr)
	if !ok {
		*expr = *ast.WarnManualMigration4Expr("resources/zero_trust_access_identity_provider#config", expr, diags)
		return
	}

	objWrapper := ast.NewObject(obj, diags)
	
	// Apply positive transforms first (idp_public_cert transformation)
	configTransforms := map[string]ast.ExprTransformer{
		"idp_public_cert": transformIdpPublicCertToList,
	}
	ast.ApplyTransformToAttributes(objWrapper, configTransforms, diags)
	
	// Remove deprecated fields directly to avoid nil expression issues
	deprecatedFields := []string{"api_token"}
	
	// Add type-specific validation rules for fields to remove
	// sign_request is only valid for type saml
	if providerType != "saml" {
		deprecatedFields = append(deprecatedFields, "sign_request")
	}
	
	// conditional_access_enabled, directory_id, support_groups are only valid for azureAD
	if providerType != "azureAD" {
		deprecatedFields = append(deprecatedFields, "conditional_access_enabled", "directory_id", "support_groups")
	}
	
	// Remove deprecated fields directly
	for _, field := range deprecatedFields {
		objWrapper.RemoveAttribute(field, diags)
	}

	// Handle field rename: idp_public_cert -> idp_public_certs
	// This needs to be done after transformation since we're changing the key name
	renameConfigField(obj, "idp_public_cert", "idp_public_certs")
}

// transformScimConfigObject handles transformations within the scim_config object
// 1. Remove deprecated group_member_deprovision field
func transformScimConfigObject(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		return
	}

	obj, ok := (*expr).(*hclsyntax.ObjectConsExpr)
	if !ok {
		*expr = *ast.WarnManualMigration4Expr("resources/zero_trust_access_identity_provider#scim_config", expr, diags)
		return
	}

	// Remove deprecated fields directly
	objWrapper := ast.NewObject(obj, diags)
	objWrapper.RemoveAttribute("group_member_deprovision", diags)
}

// transformIdpPublicCertToList converts idp_public_cert (string) to idp_public_certs (list[string])
// Before: idp_public_cert = "CERT_STRING"
// After:  idp_public_certs = ["CERT_STRING"]
func transformIdpPublicCertToList(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		return
	}

	// Convert the single certificate string to a list with one item
	*expr = &hclsyntax.TupleConsExpr{
		Exprs: []hclsyntax.Expression{*expr},
	}
}


// convertConfigBlocksToObjects converts config and scim_config blocks to objects
// This handles the v4 -> v5 block-to-object conversion
func convertConfigBlocksToObjects(block *hclwrite.Block, diags ast.Diagnostics) {
	// Convert config block to config object
	convertBlockToObject(block, "config")
	
	// Convert scim_config block to scim_config object  
	convertBlockToObject(block, "scim_config")
}

// convertBlockToObject converts a named block to an object attribute
func convertBlockToObject(parentBlock *hclwrite.Block, blockName string) {
	body := parentBlock.Body()
	blocks := body.Blocks()
	
	for _, b := range blocks {
		if b.Type() == blockName {
			// Convert block attributes to object items
			attrs := b.Body().Attributes()
			if len(attrs) == 0 {
				// Empty block -> empty object
				tokens := hclwrite.Tokens{
					{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
					{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				}
				body.SetAttributeRaw(blockName, tokens)
			} else {
				// Create object expression from block attributes
				var objTokens hclwrite.Tokens
				objTokens = append(objTokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")})
				objTokens = append(objTokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
				
				for name, attr := range attrs {
					// Add attribute name
					objTokens = append(objTokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  " + name)})
					objTokens = append(objTokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
					
					// Add attribute value tokens
					objTokens = append(objTokens, attr.Expr().BuildTokens(hclwrite.Tokens{})...)
					objTokens = append(objTokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")})
				}
				
				objTokens = append(objTokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")})
				body.SetAttributeRaw(blockName, objTokens)
			}
			
			// Remove the original block
			body.RemoveBlock(b)
			break
		}
	}
}

// renameConfigField renames a field within an object by finding the old key and updating it
func renameConfigField(obj *hclsyntax.ObjectConsExpr, oldKey, newKey string) {
	for i, item := range obj.Items {
		if ast.Expr2S(item.KeyExpr, ast.NewDiagnostics()) == oldKey {
			// Update the key expression
			obj.Items[i].KeyExpr = ast.NewKeyExpr(newKey)
			break
		}
	}
}