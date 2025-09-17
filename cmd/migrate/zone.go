package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// isZoneResource checks if a block is a cloudflare_zone resource
func isZoneResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		block.Labels()[0] == "cloudflare_zone"
}

// transformZoneBlock transforms cloudflare_zone resources
// Handles v4→v5 transformations:
// 1. zone → name (attribute rename)
// 2. account_id → account = { id = "..." } (string to nested object)
// 3. Remove jump_start (no v5 equivalent)
// 4. Remove plan (becomes computed-only nested object)
func transformZoneBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()

	// 1. Rename zone → name
	body.RenameAttribute("zone", "name")

	// 2. Transform account_id to account manually
	if accountIdAttr := body.GetAttribute("account_id"); accountIdAttr != nil {
		// Transform account_id → account = { id = "..." }
		transforms := map[string]ast.ExprTransformer{
			"account_id": transformAccountIdToNestedAccount,
		}
		ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)

		// Rename the attribute after transformation
		body.RenameAttribute("account_id", "account")
	}

	// 3-4. Remove obsolete attributes
	transforms := map[string]ast.ExprTransformer{
		"jump_start": removeAttribute,
		"plan":       removeAttribute,
	}
	ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

// transformAccountIdToNestedAccount converts account_id string to account = { id = "..." }
func transformAccountIdToNestedAccount(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	if *expr == nil {
		return
	}

	// Convert the account_id value to a nested account object
	// account_id = "abc123" → account = { id = "abc123" }
	*expr = &hclsyntax.ObjectConsExpr{
		Items: []hclsyntax.ObjectConsItem{
			{
				KeyExpr:   ast.NewKeyExpr("id"),
				ValueExpr: *expr, // Use the original expression as the id value
			},
		},
	}
}

// removeAttribute removes an attribute by setting it to nil
func removeAttribute(expr *hclsyntax.Expression, diags ast.Diagnostics) {
	// Set to nil to remove the attribute
	*expr = nil
}

// transformZoneInstanceStateJSON handles state file transformations for a single cloudflare_zone instance
func transformZoneInstanceStateJSON(state string, path string) string {
	basePath := path + ".attributes"
	var err error
	result := state

	// Let's see what attributes are actually available in this instance
	attrs := gjson.Get(state, basePath)
	if attrs.Exists() {
		attrs.ForEach(func(key, value gjson.Result) bool {
			return true
		})
	}

	// 1. zone → name
	if gjson.Get(state, basePath+".zone").Exists() {
		zoneValue := gjson.Get(state, basePath+".zone")
		result, err = sjson.Set(result, basePath+".name", zoneValue.Value())
		if err == nil {
			result, _ = sjson.Delete(result, basePath+".zone")
		}
	}

	// 2. account_id → account = { id = "..." }
	if gjson.Get(state, basePath+".account_id").Exists() {
		accountId := gjson.Get(state, basePath+".account_id")
		result, err = sjson.Set(result, basePath+".account", map[string]interface{}{
			"id": accountId.Value(),
		})
		if err == nil {
			result, _ = sjson.Delete(result, basePath+".account_id")
		}
	}

	// 3. Remove jump_start
	if gjson.Get(state, basePath+".jump_start").Exists() {
		result, _ = sjson.Delete(result, basePath+".jump_start")
	}

	// 4. Remove plan
	if gjson.Get(state, basePath+".plan").Exists() {
		result, _ = sjson.Delete(result, basePath+".plan")
	}

	// 5. Transform meta from map[string]bool to structured object
	// For now, let the API handle this since meta is computed
	// Complex meta transformations would go here if needed

	return result
}
