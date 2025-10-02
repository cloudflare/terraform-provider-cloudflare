package common

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// TransformerFunc is a function that transforms HCL blocks
type TransformerFunc func(*hclwrite.Block, *TransformContext) error

// TransformContext provides context for transformations
type TransformContext struct {
	Diagnostics []string
	Metadata    map[string]interface{}
}