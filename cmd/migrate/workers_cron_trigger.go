package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isWorkersCronTriggerResource checks if a block is a workers_cron_trigger resource
func isWorkersCronTriggerResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_cron_trigger" || block.Labels()[0] == "cloudflare_worker_cron_trigger")
}

// transformWorkersCronTriggerBlock transforms cloudflare_workers_cron_trigger resources
// Handles resource rename: cloudflare_worker_cron_trigger -> cloudflare_workers_cron_trigger
func transformWorkersCronTriggerBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource rename: cloudflare_worker_cron_trigger -> cloudflare_workers_cron_trigger
	if len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_worker_cron_trigger" {
		block.SetLabels([]string{"cloudflare_workers_cron_trigger", block.Labels()[1]})
	}
}

// transformWorkersCronTriggerStateJSON transforms the state JSON for workers_cron_trigger
func transformWorkersCronTriggerStateJSON(jsonStr string, path string) string {
	// No attribute transformations needed for cron trigger, just resource type handled by parent
	return jsonStr
}