package main

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func isManagedTransformsResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_managed_transforms"
}

func transformManagedTransformsBlock(block *hclwrite.Block) {
	body := block.Body()

	// Check if managed_request_headers exists
	requestHeaders := body.GetAttribute("managed_request_headers")
	if requestHeaders == nil {
		// In v5, this is a required attribute, so we need to add it
		// We add an empty list
		body.SetAttributeValue("managed_request_headers", cty.ListValEmpty(cty.EmptyObject))
	}

	// Check if managed_response_headers exists
	responseHeaders := body.GetAttribute("managed_response_headers")
	if responseHeaders == nil {
		// In v5, this is a required attribute, so we need to add it
		// We add an empty list
		body.SetAttributeValue("managed_response_headers", cty.ListValEmpty(cty.EmptyObject))
	}
}
