package basic

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// TransformerFunc is a function that transforms HCL blocks
type TransformerFunc func(*hclwrite.Block, *TransformContext) error

// TransformFunc is an alias for TransformerFunc for consistency
type TransformFunc = TransformerFunc

// TransformContext provides context for transformations
type TransformContext struct {
	Diagnostics []string
	Metadata    map[string]interface{}
	
	// Extended fields for new functionality
	ResourceTypeChanges map[*hclwrite.Block]string  // Maps blocks to new resource types
	MovedBlocks         map[string]string           // Maps from->to for moved blocks
	NewBlocks           []*hclwrite.Block           // Additional blocks to add
	BlocksToRemove      map[*hclwrite.Block]bool    // Blocks marked for removal
}

// ResourceTypeChange defines how to change resource types
type ResourceTypeChange struct {
	From               string `yaml:"from"`
	To                 string `yaml:"to"`
	GenerateMovedBlock bool   `yaml:"generate_moved_block"`
}

// ValueMapping defines how to transform attribute values
type ValueMapping struct {
	Attribute  string            `yaml:"attribute"`
	RenameTo   string            `yaml:"rename_to,omitempty"`
	Mappings   map[string]string `yaml:"mappings,omitempty"`
	Type       string            `yaml:"type,omitempty"`
	TrueValue  string            `yaml:"true_value,omitempty"`
	FalseValue string            `yaml:"false_value,omitempty"`
}

// ListTransform defines how to transform list attributes
type ListTransform struct {
	Attribute      string            `yaml:"attribute"`
	Type           string            `yaml:"type"`
	ObjectTemplate map[string]string `yaml:"object_template,omitempty"`
	WrapperKey     string            `yaml:"wrapper_key,omitempty"`
	ObjectKey      string            `yaml:"object_key,omitempty"`
}

// ConditionalTransform defines conditional transformations
type ConditionalTransform struct {
	Condition TransformCondition `yaml:"condition"`
	Then      TransformActions   `yaml:"then"`
	Else      *TransformActions  `yaml:"else,omitempty"`
}

// TransformCondition defines a condition for transformations
type TransformCondition struct {
	Attribute string `yaml:"attribute"`
	Operator  string `yaml:"operator"`
	Value     string `yaml:"value,omitempty"`
}

// TransformActions defines actions to take in conditional transforms
type TransformActions struct {
	RemoveAttributes []string          `yaml:"remove_attributes,omitempty"`
	SetAttributes    map[string]string `yaml:"set_attributes,omitempty"`
	RenameAttributes map[string]string `yaml:"rename_attributes,omitempty"`
}

// ResourceSplit defines how to split resources
type ResourceSplit struct {
	Type                string      `yaml:"type"`
	SourceResource      string      `yaml:"source_resource,omitempty"`
	Splits              []SplitRule `yaml:"splits"`
	Fallback            *SplitRule  `yaml:"fallback,omitempty"`
	GenerateMovedBlocks bool        `yaml:"generate_moved_blocks,omitempty"`
	RemoveOriginal      bool        `yaml:"remove_original,omitempty"`
}

// SplitRule defines a rule for splitting resources
type SplitRule struct {
	WhenAttributeExists string                 `yaml:"when_attribute_exists,omitempty"`
	ChangeResourceType  string                 `yaml:"change_resource_type,omitempty"`
	CreateResource      string                 `yaml:"create_resource,omitempty"`
	AttributeMappings   map[string]string      `yaml:"attribute_mappings,omitempty"`
	SetAttributes       map[string]interface{} `yaml:"set_attributes,omitempty"`
	CopyAttributes      []string               `yaml:"copy_attributes,omitempty"`
	ResourceNameSuffix  string                 `yaml:"resource_name_suffix,omitempty"`
}