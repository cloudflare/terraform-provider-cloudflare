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
// 6. Ensure config object exists even for OneTimePin providers
func transformZeroTrustAccessIdentityProviderBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource renaming from v4 to v5
	if block.Labels()[0] == "cloudflare_access_identity_provider" {
		block.SetLabels([]string{"cloudflare_zero_trust_access_identity_provider", block.Labels()[1]})
	}

	// Convert config and scim_config blocks to objects (if they exist as blocks)
	convertConfigBlocksToObjects(block, diags)

	// Ensure config object exists (required in v5, even if empty for OneTimePin)
	ensureConfigObjectExists(block)

	// Apply config-specific transformations
	transforms := map[string]ast.ExprTransformer{
		"config":      transformConfigObject,
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

// transformConfigObject handles transformations within the config object
// 1. idp_public_cert -> idp_public_certs (string -> list[string])
// 2. Remove deprecated api_token field
func transformConfigObject(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		return
	}

	obj, ok := (*expr).(*hclsyntax.ObjectConsExpr)
	if !ok {
		*expr = *ast.WarnManualMigration4Expr("resources/zero_trust_access_identity_provider#config", expr, diags)
		return
	}

	// Apply config-specific transforms
	configTransforms := map[string]ast.ExprTransformer{
		"idp_public_cert": transformIdpPublicCertToList,
		"api_token":       removeDeprecatedField,
	}
	ast.ApplyTransformToAttributes(ast.NewObject(obj, diags), configTransforms, diags)

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

	// Apply scim_config-specific transforms
	scimTransforms := map[string]ast.ExprTransformer{
		"group_member_deprovision": removeDeprecatedField,
	}
	ast.ApplyTransformToAttributes(ast.NewObject(obj, diags), scimTransforms, diags)
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

// removeDeprecatedField removes deprecated fields by setting them to nil
func removeDeprecatedField(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	*expr = nil
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