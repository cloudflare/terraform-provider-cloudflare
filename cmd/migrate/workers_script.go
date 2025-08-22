package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// isWorkersScriptResource checks if a block is a workers_script resource
func isWorkersScriptResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_script" || block.Labels()[0] == "cloudflare_worker_script")
}

// transformWorkersScriptBlock transforms cloudflare_workers_script resources
// V4: name -> V5: script_name
// Also handles resource rename: cloudflare_worker_script -> cloudflare_workers_script
func transformWorkersScriptBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource rename: cloudflare_worker_script -> cloudflare_workers_script
	if len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_worker_script" {
		block.SetLabels([]string{"cloudflare_workers_script", block.Labels()[1]})
	}

	// Rename name attribute to script_name
	block.Body().RenameAttribute("name", "script_name")
}

// transformWorkersScriptStateJSON transforms the state JSON for workers_script
// V4: name -> V5: script_name
func transformWorkersScriptStateJSON(jsonStr string, path string) string {
	result := jsonStr

	// Get the current name value
	namePath := path + ".attributes.name"
	scriptNamePath := path + ".attributes.script_name"

	// Get the value and move it
	nameValue := gjson.Get(jsonStr, namePath)
	if nameValue.Exists() {
		// Set the new attribute
		result, _ = sjson.Set(result, scriptNamePath, nameValue.Value())
		// Delete the old attribute
		result, _ = sjson.Delete(result, namePath)
	}

	return result
}
