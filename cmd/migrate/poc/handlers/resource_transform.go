package handlers

import (
	"fmt"
	"log"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ResourceTransformHandler applies resource-specific transformations using strategies
type ResourceTransformHandler struct {
	interfaces.BaseHandler
	registry *registry.StrategyRegistry
}

// NewResourceTransformHandler creates a new resource transformation handler
func NewResourceTransformHandler(reg *registry.StrategyRegistry) interfaces.TransformationHandler {
	return &ResourceTransformHandler{
		registry: reg,
	}
}

// Handle applies transformations to each resource block
func (h *ResourceTransformHandler) Handle(ctx *interfaces.TransformContext) (*interfaces.TransformContext, error) {
	if ctx.AST == nil {
		return ctx, fmt.Errorf("AST is nil - ParseHandler must run before ResourceTransformHandler")
	}

	body := ctx.AST.Body()
	blocks := body.Blocks()

	var blocksToRemove []*hclwrite.Block
	var blocksToAdd []*hclwrite.Block

	// Process each block
	for _, block := range blocks {
		if block.Type() != "resource" {
			continue
		}

		labels := block.Labels()
		if len(labels) < 1 {
			continue
		}

		resourceType := labels[0]
		strategy := h.registry.Find(resourceType)

		if strategy == nil {
			log.Printf("No strategy found for resource type: %s", resourceType)
			continue
		}

		// Apply transformation using the unified TransformConfig method
		result, err := strategy.TransformConfig(block)
		if err != nil {
			log.Printf("Error transforming resource %s: %v", resourceType, err)
			ctx.Diagnostics = append(ctx.Diagnostics, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf("Failed to transform %s resource", resourceType),
				Detail:   err.Error(),
			})
			continue
		}

		// Handle the result uniformly
		if result.RemoveOriginal {
			blocksToRemove = append(blocksToRemove, block)
			// Only add new blocks if we're removing the original
			blocksToAdd = append(blocksToAdd, result.Blocks...)
		}

		h.trackTransformation(ctx, resourceType)
	}

	// Apply all changes
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}
	for _, block := range blocksToAdd {
		body.AppendBlock(block)
	}

	return h.CallNext(ctx)
}

func (h *ResourceTransformHandler) trackTransformation(ctx *interfaces.TransformContext, resourceType string) {
	if ctx.Metadata == nil {
		ctx.Metadata = make(map[string]interface{})
	}
	transformedKey := fmt.Sprintf("transformed_%s", resourceType)
	count := 0
	if val, ok := ctx.Metadata[transformedKey]; ok {
		count = val.(int)
	}
	ctx.Metadata[transformedKey] = count + 1
}
