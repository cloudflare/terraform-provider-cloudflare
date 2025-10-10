package base

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
)

// ModelParser defines the interface for parsing HCL to v4 models
type ModelParser interface {
	Parse(block *hclwrite.Block) (interface{}, error)
}

// ModelTransformer defines the interface for v4 to v5 model transformation
type ModelTransformer interface {
	Transform(v4Model interface{}) (interface{}, error)
}

// ModelGenerator defines the interface for generating HCL from v5 models
type ModelGenerator interface {
	Generate(v5Model interface{}, labels []string) *hclwrite.Block
}

// BaseStructTransformer provides a generic implementation of ResourceTransformer
// using struct-based transformation approach
type BaseStructTransformer struct {
	ResourceType string
	Parser       ModelParser
	Transformer  ModelTransformer
	Generator    ModelGenerator
}

// NewBaseStructTransformer creates a new generic struct transformer
func NewBaseStructTransformer(
	resourceType string,
	parser ModelParser,
	transformer ModelTransformer,
	generator ModelGenerator,
) *BaseStructTransformer {
	return &BaseStructTransformer{
		ResourceType: resourceType,
		Parser:       parser,
		Transformer:  transformer,
		Generator:    generator,
	}
}

// CanHandle checks if this transformer can handle the given resource type
func (t *BaseStructTransformer) CanHandle(resourceType string) bool {
	return resourceType == t.ResourceType
}

// GetResourceType returns the resource type this transformer handles
func (t *BaseStructTransformer) GetResourceType() string {
	return t.ResourceType
}

// TransformConfig transforms the HCL configuration block
func (t *BaseStructTransformer) TransformConfig(block *hclwrite.Block) (*interfaces.TransformResult, error) {
	// Create v4 model
	v4Model, err := t.Parser.Parse(block)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", t.ResourceType, err)
	}

	// Migrate to v5
	v5Model, err := t.Transformer.Transform(v4Model)
	if err != nil {
		return nil, fmt.Errorf("failed to transform %s: %w", t.ResourceType, err)
	}

	// Create v5 config
	newBlock := t.Generator.Generate(v5Model, block.Labels())

	return &interfaces.TransformResult{
		Blocks:         []*hclwrite.Block{newBlock},
		RemoveOriginal: true,
	}, nil
}

// TODO:: implement transformation
func (t *BaseStructTransformer) TransformState(json gjson.Result, resourcePath string) (string, error) {
	return json.String(), nil
}

// Preprocess handles string-level transformations before HCL parsing
func (t *BaseStructTransformer) Preprocess(content string) string {
	// No preprocessing needed for struct-based approach
	// The resource type rename happens during generation
	return content
}
