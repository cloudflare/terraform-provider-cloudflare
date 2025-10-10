// Package interfaces defines the core abstractions for the migration tool refactoring
package interfaces

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
)

// TransformResult represents the result of a resource transformation
type TransformResult struct {
	// Blocks to add (includes the modified original for in-place changes)
	Blocks []*hclwrite.Block

	// Whether to remove the original block
	RemoveOriginal bool
}

// ResourceTransformer defines the strategy pattern interface for resource-specific transformations
// Each resource type will implement this interface to handle its specific migration logic
type ResourceTransformer interface {
	// CanHandle determines if this strategy can transform the given resource type
	CanHandle(resourceType string) bool

	// TransformConfig handles any type of transformation:
	// - In-place: return TransformResult{Blocks: []*hclwrite.Block{modifiedBlock}, RemoveOriginal: true}
	// - Split: return TransformResult{Blocks: newBlocks, RemoveOriginal: true}
	// - Remove: return TransformResult{Blocks: nil, RemoveOriginal: true}
	TransformConfig(block *hclwrite.Block) (*TransformResult, error)

	// TransformState handles transformation of Terraform state files
	// Returns the modified JSON string for the resource
	TransformState(json gjson.Result, resourcePath string) (string, error)

	// GetResourceType returns the primary resource type this transformer handles
	GetResourceType() string

	// Preprocess handles string-level transformations before HCL parsing
	// This is optional - resources that don't need preprocessing can return the content unchanged
	// Common use cases: resource type renames, regex-based transformations
	Preprocess(content string) string
}
