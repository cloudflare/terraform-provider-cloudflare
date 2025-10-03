package basic

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ResourceTypeChanger creates a transformer that changes resource types
//
// Example YAML configuration:
//   resource_type_changes:
//     from: cloudflare_access_application
//     to: cloudflare_zero_trust_access_application
//     generate_moved_block: true
//
// Transforms:
//   resource "cloudflare_access_application" "example" {
//     zone_id = "abc123"
//     name = "test app"
//     domain = "test.example.com"
//   }
//
// Into:
//   resource "cloudflare_zero_trust_access_application" "example" {
//     zone_id = "abc123"
//     name = "test app"
//     domain = "test.example.com"
//   }
//
//   moved {
//     from = cloudflare_access_application.example
//     to = cloudflare_zero_trust_access_application.example
//   }
func ResourceTypeChanger(change *ResourceTypeChange) TransformFunc {
	if change == nil {
		return func(block *hclwrite.Block, ctx *TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *TransformContext) error {
		if block.Type() != "resource" {
			return nil
		}

		labels := block.Labels()
		if len(labels) < 2 {
			return nil
		}

		// Check if this is the resource type we want to change
		if labels[0] == change.From {
			// Store the original resource reference for moved block
			if change.GenerateMovedBlock && ctx != nil {
				originalRef := change.From + "." + labels[1]
				newRef := change.To + "." + labels[1]
				
				// Add to context for later moved block generation
				if ctx.MovedBlocks == nil {
					ctx.MovedBlocks = make(map[string]string)
				}
				ctx.MovedBlocks[originalRef] = newRef
			}

			// Note: We can't directly change labels on an existing block in hclwrite
			// Instead, we need to mark this for reconstruction
			if ctx.ResourceTypeChanges == nil {
				ctx.ResourceTypeChanges = make(map[*hclwrite.Block]string)
			}
			ctx.ResourceTypeChanges[block] = change.To
		}

		return nil
	}
}

// ApplyResourceTypeChanges reconstructs blocks with new resource types
func ApplyResourceTypeChanges(blocks []*hclwrite.Block, ctx *TransformContext) []*hclwrite.Block {
	if ctx.ResourceTypeChanges == nil || len(ctx.ResourceTypeChanges) == 0 {
		return blocks
	}

	var result []*hclwrite.Block
	
	for _, block := range blocks {
		if newType, ok := ctx.ResourceTypeChanges[block]; ok {
			// Create new block with updated resource type
			labels := block.Labels()
			newLabels := make([]string, len(labels))
			copy(newLabels, labels)
			newLabels[0] = newType
			
			newBlock := hclwrite.NewBlock(block.Type(), newLabels)
			
			// Copy all attributes
			body := block.Body()
			newBody := newBlock.Body()
			
			for name, attr := range body.Attributes() {
				newBody.SetAttributeRaw(name, attr.Expr().BuildTokens(nil))
			}
			
			// Copy all nested blocks
			for _, nestedBlock := range body.Blocks() {
				copyBlockContent(newBody, nestedBlock)
			}
			
			result = append(result, newBlock)
		} else {
			result = append(result, block)
		}
	}
	
	// Add moved blocks if any
	for from, to := range ctx.MovedBlocks {
		movedBlock := createMovedBlock(from, to)
		result = append(result, movedBlock)
	}
	
	return result
}

// Helper function to copy block content
func copyBlockContent(targetBody *hclwrite.Body, sourceBlock *hclwrite.Block) {
	newBlock := hclwrite.NewBlock(sourceBlock.Type(), sourceBlock.Labels())
	sourceBody := sourceBlock.Body()
	newBody := newBlock.Body()
	
	// Copy attributes
	for name, attr := range sourceBody.Attributes() {
		newBody.SetAttributeRaw(name, attr.Expr().BuildTokens(nil))
	}
	
	// Recursively copy nested blocks
	for _, nested := range sourceBody.Blocks() {
		copyBlockContent(newBody, nested)
	}
	
	targetBody.AppendBlock(newBlock)
}

// Helper function to create moved blocks
func createMovedBlock(from, to string) *hclwrite.Block {
	movedBlock := hclwrite.NewBlock("moved", []string{})
	body := movedBlock.Body()
	
	// Set from attribute
	body.SetAttributeTraversal("from", hcl.Traversal{
		hcl.TraverseRoot{Name: from},
	})
	
	// Set to attribute
	body.SetAttributeTraversal("to", hcl.Traversal{
		hcl.TraverseRoot{Name: to},
	})
	
	return movedBlock
}