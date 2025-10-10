package struct_transform

import (
	"fmt"
	"log"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/resources"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// StructTransformHandler handles struct-based transformations
// This follows the same pattern as ResourceTransformHandler but uses struct-based transformers
type StructTransformHandler struct {
	interfaces.BaseHandler
	registry *registry.StrategyRegistry
}

// NewStructTransformHandler creates a new struct transformation handler with struct-based transformers
func NewStructTransformHandler(reg *registry.StrategyRegistry) interfaces.TransformationHandler {
	// If no registry provided, create one with all struct-based transformers
	if reg == nil {
		reg = registry.NewStrategyRegistry()
		// Register all struct-based transformers from factory
		resources.RegisterAllStructTransformers(reg)
	}

	return &StructTransformHandler{
		registry: reg,
	}
}

// Handle applies transformations to each resource block using struct-based transformers
// This mirrors the logic in ResourceTransformHandler but uses struct-based strategies
func (h *StructTransformHandler) Handle(ctx *interfaces.TransformContext) (*interfaces.TransformContext, error) {
	if ctx.AST == nil {
		return ctx, fmt.Errorf("AST is nil - ParseHandler must run before StructTransformHandler")
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
			log.Printf("No struct strategy found for resource type: %s", resourceType)
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
		// If not removing original, the block was modified in-place
		// Don't add the blocks from result as they're the same modified block

		// Track successful transformation
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

func (h *StructTransformHandler) trackTransformation(ctx *interfaces.TransformContext, resourceType string) {
	if ctx.Metadata == nil {
		ctx.Metadata = make(map[string]interface{})
	}
	transformedKey := fmt.Sprintf("struct_transformed_%s", resourceType)
	count := 0
	if val, ok := ctx.Metadata[transformedKey]; ok {
		count = val.(int)
	}
	ctx.Metadata[transformedKey] = count + 1

	// Also track that struct mode was used
	ctx.Metadata["struct_mode_used"] = true
}