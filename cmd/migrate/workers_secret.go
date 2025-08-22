package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isWorkersSecretResource checks if a block is a workers_secret resource
func isWorkersSecretResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 1 &&
		(block.Labels()[0] == "cloudflare_workers_secret" || block.Labels()[0] == "cloudflare_worker_secret")
}

// transformWorkersSecretBlock transforms cloudflare_workers_secret resources
// Handles resource rename: cloudflare_worker_secret -> cloudflare_workers_secret
func transformWorkersSecretBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// Handle resource rename: cloudflare_worker_secret -> cloudflare_workers_secret
	if len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_worker_secret" {
		block.SetLabels([]string{"cloudflare_workers_secret", block.Labels()[1]})
	}
}

// transformWorkersSecretStateJSON transforms the state JSON for workers_secret
func transformWorkersSecretStateJSON(jsonStr string, path string) string {
	// No attribute transformations needed for worker secret, just resource type handled by parent
	return jsonStr
}
