package internal

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

// ConditionalResource defines conditional resource creation
type ConditionalResource struct {
	Condition ResourceCondition `yaml:"condition"`
	Actions   []ResourceAction  `yaml:"actions"`
}

// ResourceCondition defines a condition for resource operations
type ResourceCondition struct {
	Attribute string `yaml:"attribute"`
	Operator  string `yaml:"operator"`
	Value     string `yaml:"value,omitempty"`
}

// ResourceAction defines an action to take on a resource
type ResourceAction struct {
	Type               string            `yaml:"type"`
	ResourceType       string            `yaml:"resource_type,omitempty"`
	To                 string            `yaml:"to,omitempty"`
	CopyAttributes     []string          `yaml:"copy_attributes,omitempty"`
	SetAttributes      map[string]interface{} `yaml:"set_attributes,omitempty"`
	Attribute          string            `yaml:"attribute,omitempty"`
	Value              interface{}      `yaml:"value,omitempty"`
	GenerateMovedBlock bool              `yaml:"generate_moved_block,omitempty"`
	RemoveOriginal     bool              `yaml:"remove_original,omitempty"`
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

// StructuralTransform defines complex structural transformations
type StructuralTransform struct {
	Type         string            `yaml:"type"`
	Attributes   []string          `yaml:"attributes,omitempty"`
	Except       []string          `yaml:"except,omitempty"`
	WrapperName  string            `yaml:"wrapper_name,omitempty"`
	RenameMap    map[string]string `yaml:"rename_map,omitempty"`
	RemoveIfNull bool              `yaml:"remove_if_null,omitempty"`
	Source       string            `yaml:"source,omitempty"`
	Target       string            `yaml:"target,omitempty"`
	Prefix       string            `yaml:"prefix,omitempty"`
}

// ExtendedConfigTransformations includes all new transformation types
type ExtendedConfigTransformations struct {
	ConfigTransformations `yaml:",inline"`
	
	// New fields
	ResourceTypeChange    *ResourceTypeChange    `yaml:"resource_type_change,omitempty"`
	BlockRemovals         []string               `yaml:"block_removals,omitempty"`
	ValueMappings         []ValueMapping         `yaml:"value_mappings,omitempty"`
	ListTransforms        []ListTransform        `yaml:"list_transforms,omitempty"`
	ConditionalTransforms []ConditionalTransform `yaml:"conditional_transforms,omitempty"`
	ConditionalResources  []ConditionalResource  `yaml:"conditional_resources,omitempty"`
	ResourceSplits        []ResourceSplit        `yaml:"resource_splits,omitempty"`
	StructuralTransforms  []StructuralTransform  `yaml:"structural_transforms,omitempty"`
}

// ExtendedStateTransformations includes all new state transformation types
type ExtendedStateTransformations struct {
	StateTransformations `yaml:",inline"`
	
	// New fields
	ResourceTypeChange   *ResourceTypeChange   `yaml:"resource_type_change,omitempty"`
	ValueMappings        []ValueMapping        `yaml:"value_mappings,omitempty"`
	ListTransforms       []ListTransform       `yaml:"list_transforms,omitempty"`
	ConditionalResources []ConditionalResource `yaml:"conditional_resources,omitempty"`
	ResourceSplits       []ResourceSplit       `yaml:"resource_splits,omitempty"`
}

// ExtendedMigrationConfig includes all extended transformations
type ExtendedMigrationConfig struct {
	ResourceType  string                        `yaml:"resource_type"`
	SourceVersion string                        `yaml:"source_version"`
	TargetVersion string                        `yaml:"target_version"`
	Description   string                        `yaml:"description"`
	Config        ExtendedConfigTransformations `yaml:"config"`
	State         ExtendedStateTransformations  `yaml:"state"`
}