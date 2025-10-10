package handlers

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
)

// PreprocessHandler handles string-level transformations before HCL parsing. Delegated to resource specific impl
type PreprocessHandler struct {
	interfaces.BaseHandler
	registry *registry.StrategyRegistry
}

// NewPreprocessHandler creates a new preprocessing handler
func NewPreprocessHandler(reg *registry.StrategyRegistry) interfaces.TransformationHandler {
	return &PreprocessHandler{
		registry: reg,
	}
}

// Handle applies string-level transformations
func (h *PreprocessHandler) Handle(ctx *interfaces.TransformContext) (*interfaces.TransformContext, error) {
	contentStr := string(ctx.Content)
	contentStr = h.applyAllPreprocessors(contentStr)
	ctx.Content = []byte(contentStr)
	return h.CallNext(ctx)
}

// applyAllPreprocessors applies all registered preprocessors to the content
func (h *PreprocessHandler) applyAllPreprocessors(content string) string {
	strategies := h.registry.GetAll()

	// Apply resource specific preprocessors
	for _, strategy := range strategies {
		content = strategy.Preprocess(content)
	}

	return content
}
