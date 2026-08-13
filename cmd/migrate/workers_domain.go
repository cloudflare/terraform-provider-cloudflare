package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isWorkersDomainResource checks if a block is a workers_custom_domain resource
func isWorkersDomainResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_custom_domain" || block.Labels()[0] == "cloudflare_worker_domain")
}

// transformWorkersDomainBlock transforms cloudflare_workers_custom_domain resources
// Handles resource rename: cloudflare_worker_domain -> cloudflare_workers_custom_domain
func transformWorkersDomainBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource rename: cloudflare_worker_domain -> cloudflare_workers_custom_domain
	if len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_worker_domain" {
		block.SetLabels([]string{"cloudflare_workers_custom_domain", block.Labels()[1]})
	}
}

// transformWorkersDomainStateJSON transforms the state JSON for workers_custom_domain
func transformWorkersDomainStateJSON(jsonStr string, path string) string {
	// No attribute transformations needed for worker domain, just resource type handled by parent
	return jsonStr
}