package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// transformLoadBalancerFile applies cloudflare_load_balancer specific transformations to an HCL file
func transformLoadBalancerFile(file *hclwrite.File) {
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_load_balancer" {
			transformLoadBalancerBlock(block)
		}
	}
}

// transformLoadBalancerBlock handles block-level transformations for cloudflare_load_balancer
func transformLoadBalancerBlock(block *hclwrite.Block) {
	// Currently no config transformations needed for load_balancer
	// The state transformations are in transformLoadBalancerState
}

