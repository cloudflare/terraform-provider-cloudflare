package resources

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/tidwall/gjson"
)

// TransformationMode represents the transformation approach
type TransformationMode int

const (
	// ASTMode uses direct HCL AST manipulation
	ASTMode TransformationMode = iota
	// StructMode uses struct-based transformation with parsing to Go structs
	StructMode
)

// BaseResourceTransformer provides common functionality for resource transformers
// that can operate in either AST or Struct mode
type BaseResourceTransformer struct {
	resourceType string
	mode         TransformationMode

	// AST-based transformation function
	astTransformer func(block *hclwrite.Block) (*interfaces.TransformResult, error)

	// Struct-based transformation function
	structTransformer func(block *hclwrite.Block) (*interfaces.TransformResult, error)

	// State transformation function (common to both modes)
	stateTransformer func(json gjson.Result, resourcePath string) (string, error)

	// Preprocessing function (common to both modes)
	preprocessor func(content string) string
}

// NewBaseResourceTransformer creates a new base transformer
func NewBaseResourceTransformer(resourceType string, mode TransformationMode) *BaseResourceTransformer {
	return &BaseResourceTransformer{
		resourceType: resourceType,
		mode:         mode,
	}
}

// SetASTTransformer sets the AST-based transformation function
func (b *BaseResourceTransformer) SetASTTransformer(fn func(block *hclwrite.Block) (*interfaces.TransformResult, error)) {
	b.astTransformer = fn
}

// SetStructTransformer sets the struct-based transformation function
func (b *BaseResourceTransformer) SetStructTransformer(fn func(block *hclwrite.Block) (*interfaces.TransformResult, error)) {
	b.structTransformer = fn
}

// SetStateTransformer sets the state transformation function
func (b *BaseResourceTransformer) SetStateTransformer(fn func(json gjson.Result, resourcePath string) (string, error)) {
	b.stateTransformer = fn
}

// SetPreprocessor sets the preprocessing function
func (b *BaseResourceTransformer) SetPreprocessor(fn func(content string) string) {
	b.preprocessor = fn
}

// CanHandle determines if this strategy can transform the given resource type
func (b *BaseResourceTransformer) CanHandle(resourceType string) bool {
	return resourceType == b.resourceType
}

// GetResourceType returns the primary resource type this transformer handles
func (b *BaseResourceTransformer) GetResourceType() string {
	return b.resourceType
}

// TransformConfig handles the transformation based on the configured mode
func (b *BaseResourceTransformer) TransformConfig(block *hclwrite.Block) (*interfaces.TransformResult, error) {
	switch b.mode {
	case StructMode:
		if b.structTransformer != nil {
			return b.structTransformer(block)
		}
		// Fall back to AST mode if no struct transformer is defined
		fallthrough
	case ASTMode:
		if b.astTransformer != nil {
			return b.astTransformer(block)
		}
		// If no transformer is defined, return the block unchanged
		return &interfaces.TransformResult{
			Blocks:         []*hclwrite.Block{block},
			RemoveOriginal: false,
		}, nil
	default:
		// Default to AST mode
		if b.astTransformer != nil {
			return b.astTransformer(block)
		}
		return &interfaces.TransformResult{
			Blocks:         []*hclwrite.Block{block},
			RemoveOriginal: false,
		}, nil
	}
}

// TransformState handles transformation of Terraform state files
func (b *BaseResourceTransformer) TransformState(json gjson.Result, resourcePath string) (string, error) {
	if b.stateTransformer != nil {
		return b.stateTransformer(json, resourcePath)
	}
	// Return unchanged if no state transformer is defined
	return json.String(), nil
}

// Preprocess handles string-level transformations before HCL parsing
func (b *BaseResourceTransformer) Preprocess(content string) string {
	if b.preprocessor != nil {
		return b.preprocessor(content)
	}
	// Return unchanged if no preprocessor is defined
	return content
}

// GetMode returns the current transformation mode
func (b *BaseResourceTransformer) GetMode() TransformationMode {
	return b.mode
}

// SetMode changes the transformation mode
func (b *BaseResourceTransformer) SetMode(mode TransformationMode) {
	b.mode = mode
}

// Ensure BaseResourceTransformer implements ResourceTransformer
var _ interfaces.ResourceTransformer = (*BaseResourceTransformer)(nil)