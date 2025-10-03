package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	
	
)

// BlockRemover creates a transformer that removes specified block types
//
// Example YAML configuration:
//   block_removals:
//     - deprecated_block
//     - legacy_settings
//     - obsolete_config
//
// Transforms:
//   resource "example" "test" {
//     name = "test"
//     
//     deprecated_block {
//       old_setting = "value"
//     }
//     
//     legacy_settings {
//       param = "old"
//     }
//     
//     current_block {
//       setting = "keep"
//     }
//   }
//
// Into:
//   resource "example" "test" {
//     name = "test"
//     
//     current_block {
//       setting = "keep"
//     }
//   }
func BlockRemover(blockTypes ...string) TransformFunc {
	if len(blockTypes) == 0 {
		return func(block *hclwrite.Block, ctx *TransformContext) error {
			return nil
		}
	}

	// Create a set for efficient lookup
	blockTypeSet := make(map[string]bool)
	for _, bt := range blockTypes {
		blockTypeSet[bt] = true
	}

	return func(block *hclwrite.Block, ctx *TransformContext) error {
		removeBlocksRecursively(block.Body(), blockTypeSet)
		return nil
	}
}

// removeBlocksRecursively removes blocks from a body and its nested blocks
func removeBlocksRecursively(body *hclwrite.Body, blockTypeSet map[string]bool) {
	// Collect blocks to remove from this body
	var blocksToRemove []*hclwrite.Block
	
	for _, block := range body.Blocks() {
		if blockTypeSet[block.Type()] {
			blocksToRemove = append(blocksToRemove, block)
		} else {
			// Recursively process nested blocks
			removeBlocksRecursively(block.Body(), blockTypeSet)
		}
	}
	
	// Remove blocks from this body
	for _, blockToRemove := range blocksToRemove {
		body.RemoveBlock(blockToRemove)
	}
}

// SelectiveBlockRemover creates a transformer that removes blocks based on more complex criteria
func SelectiveBlockRemover(criteria func(*hclwrite.Block) bool) TransformFunc {
	if criteria == nil {
		return func(block *hclwrite.Block, ctx *TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *TransformContext) error {
		selectiveRemoveBlocksRecursively(block.Body(), criteria)
		return nil
	}
}

// selectiveRemoveBlocksRecursively removes blocks matching criteria from a body and its nested blocks
func selectiveRemoveBlocksRecursively(body *hclwrite.Body, criteria func(*hclwrite.Block) bool) {
	// Collect blocks to remove from this body
	var blocksToRemove []*hclwrite.Block
	
	for _, block := range body.Blocks() {
		if criteria(block) {
			blocksToRemove = append(blocksToRemove, block)
		} else {
			// Recursively process nested blocks
			selectiveRemoveBlocksRecursively(block.Body(), criteria)
		}
	}
	
	// Remove blocks from this body
	for _, blockToRemove := range blocksToRemove {
		body.RemoveBlock(blockToRemove)
	}
}