package interfaces

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// TransformContext carries data through the transformation pipeline.
type TransformContext struct {
	// Original file content
	Content []byte

	// Filename for diagnostics
	Filename string

	// Parsed HCL AST (set by ParseHandler)
	AST *hclwrite.File

	// State data for state file transformations
	StateJSON string

	// Diagnostics accumulator
	Diagnostics hcl.Diagnostics

	// Metadata for handlers to communicate
	Metadata map[string]interface{}

	// Flag indicating if this is a dry run
	DryRun bool
}

// TransformationHandler defines the chain of responsibility pattern
// Each handler processes the context and passes it to the next handler
type TransformationHandler interface {
	// Handle processes the transformation context
	Handle(ctx *TransformContext) (*TransformContext, error)

	// SetNext sets the next handler in the chain
	SetNext(handler TransformationHandler)
}

// BaseHandler provides common functionality for all handlers
type BaseHandler struct {
	next TransformationHandler
}

// SetNext sets the next handler in the chain
func (h *BaseHandler) SetNext(handler TransformationHandler) {
	h.next = handler
}

// CallNext passes the context to the next handler if one exists
func (h *BaseHandler) CallNext(ctx *TransformContext) (*TransformContext, error) {
	if h.next != nil {
		return h.next.Handle(ctx)
	}
	return ctx, nil
}
