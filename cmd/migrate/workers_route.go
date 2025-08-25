package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// isWorkersRouteResource checks if a block is a workers_route resource
func isWorkersRouteResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_route" || block.Labels()[0] == "cloudflare_worker_route")
}

// transformWorkersRouteBlock transforms cloudflare_workers_route resources
// V4: script_name -> V5: script
// Also handles resource rename: cloudflare_worker_route -> cloudflare_workers_route
func transformWorkersRouteBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource rename: cloudflare_worker_route -> cloudflare_workers_route
	if len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_worker_route" {
		block.SetLabels([]string{"cloudflare_workers_route", block.Labels()[1]})
	}

	// Rename script_name attribute to script
	block.Body().RenameAttribute("script_name", "script")
}

// transformWorkersRouteStateJSON transforms the state JSON for workers_route
// V4: script_name -> V5: script
func transformWorkersRouteStateJSON(jsonStr string, path string) string {
	result := jsonStr

	// Get the current script_name value
	scriptNamePath := path + ".attributes.script_name"
	scriptPath := path + ".attributes.script"

	// Get the value and move it
	scriptNameValue := gjson.Get(jsonStr, scriptNamePath)
	if scriptNameValue.Exists() {
		// Set the new attribute
		result, _ = sjson.Set(result, scriptPath, scriptNameValue.Value())
		// Delete the old attribute
		result, _ = sjson.Delete(result, scriptNamePath)
	}

	return result
}
