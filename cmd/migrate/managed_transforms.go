package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isManagedTransformsResource checks if a block is a cloudflare_managed_transforms resource
// (after rename from cloudflare_managed_headers)
func isManagedTransformsResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_managed_transforms"
}

// transformManagedTransformsBlock ensures both managed_request_headers and managed_response_headers
// are present as required attributes in v5
func transformManagedTransformsBlock(block *hclwrite.Block) {
	body := block.Body()
	
	// Check if managed_request_headers exists
	requestHeaders := body.GetAttribute("managed_request_headers")
	if requestHeaders == nil {
		// Add empty list for managed_request_headers
		body.SetAttributeValue("managed_request_headers", cty.ListValEmpty(cty.Object(map[string]cty.Type{
			"id":      cty.String,
			"enabled": cty.Bool,
		})))
	}
	
	// Check if managed_response_headers exists
	responseHeaders := body.GetAttribute("managed_response_headers")
	if responseHeaders == nil {
		// Add empty list for managed_response_headers
		body.SetAttributeValue("managed_response_headers", cty.ListValEmpty(cty.Object(map[string]cty.Type{
			"id":      cty.String,
			"enabled": cty.Bool,
		})))
	}
}