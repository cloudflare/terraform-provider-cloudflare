package poc

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/handlers"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
)

func BuildPipeline(reg *registry.StrategyRegistry) *Pipeline {
	return BuildPipelineWithOptions(reg, PipelineOptions{})
}

// PipelineOptions contains configuration options for the pipeline
type PipelineOptions struct {
	// Reserved for future options
}

// BuildPipelineWithOptions creates a pipeline with specified options
func BuildPipelineWithOptions(reg *registry.StrategyRegistry, opts PipelineOptions) *Pipeline {
	// Use a single unified transformation handler
	// Each strategy in the registry decides internally whether to use AST or struct-based approach
	builder := NewPipelineBuilder().
		WithPreprocessing(reg).
		WithParsing().
		WithResourceTransformation(reg).
		// Cross-resource migrations would go here
		// WithHandler(handlers.NewCrossResourceHandler()).
		// Import generation for split resources would go here
		// WithHandler(handlers.NewImportGeneratorHandler()).
		// Validation would go here
		// WithHandler(handlers.NewValidationHandler()).
		WithFormatting().
		Build()

	return builder
}

type Pipeline struct {
	handler  interfaces.TransformationHandler
	registry *registry.StrategyRegistry
}

func NewPipeline(reg *registry.StrategyRegistry) *Pipeline {
	pipeline := BuildPipeline(reg)
	return pipeline
}

func (p *Pipeline) Transform(content []byte, filename string) ([]byte, error) {
	ctx := &interfaces.TransformContext{
		Content:     content,
		Filename:    filename,
		Diagnostics: nil,
		Metadata:    make(map[string]interface{}),
		DryRun:      false,
	}

	result, err := p.handler.Handle(ctx)
	if err != nil {
		return nil, err
	}

	if result.Diagnostics.HasErrors() {
		return nil, result.Diagnostics.Errs()[0]
	}

	return result.Content, nil
}

type PipelineBuilder struct {
	handlers []interfaces.TransformationHandler
	registry *registry.StrategyRegistry
}

func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{
		handlers: make([]interfaces.TransformationHandler, 0),
		registry: nil,
	}
}

func (b *PipelineBuilder) WithPreprocessing(reg *registry.StrategyRegistry) *PipelineBuilder {
	b.handlers = append(b.handlers, handlers.NewPreprocessHandler(reg))
	return b
}

func (b *PipelineBuilder) WithParsing() *PipelineBuilder {
	b.handlers = append(b.handlers, handlers.NewParseHandler())
	return b
}

func (b *PipelineBuilder) WithResourceTransformation(reg *registry.StrategyRegistry) *PipelineBuilder {
	b.registry = reg
	b.handlers = append(b.handlers, handlers.NewResourceTransformHandler(reg))
	return b
}

func (b *PipelineBuilder) WithFormatting() *PipelineBuilder {
	b.handlers = append(b.handlers, handlers.NewFormatterHandler())
	return b
}

func (b *PipelineBuilder) WithHandler(handler interfaces.TransformationHandler) *PipelineBuilder {
	b.handlers = append(b.handlers, handler)
	return b
}

func (b *PipelineBuilder) Build() *Pipeline {
	if len(b.handlers) == 0 {
		return nil
	}
	for i := 0; i < len(b.handlers)-1; i++ {
		b.handlers[i].SetNext(b.handlers[i+1])
	}

	return &Pipeline{
		handler:  b.handlers[0],
		registry: b.registry,
	}
}