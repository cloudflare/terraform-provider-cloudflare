package internal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"gopkg.in/yaml.v3"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/structural"
)

// MigrationConfig defines the YAML configuration for a migration
type MigrationConfig struct {
	ResourceType               string                `yaml:"resource_type"`
	SourceVersion              string                `yaml:"source_version"`
	TargetVersion              string                `yaml:"target_version"`
	Description                string                `yaml:"description"`
	RequiresFileTransformation bool                  `yaml:"requires_file_transformation"`
	Config                     ConfigTransformations `yaml:"config"`
	State                      StateTransformations  `yaml:"state"`
}

// ConfigTransformations defines configuration transformations
type ConfigTransformations struct {
	AttributeRenames    map[string]string      `yaml:"attribute_renames"`
	AttributeRemovals   []string               `yaml:"attribute_removals"`
	ConditionalRemovals []ConditionalRemoval   `yaml:"conditional_removals"`
	TypeConversions     []TypeConversion       `yaml:"type_conversions"`
	BlocksToLists       []string               `yaml:"blocks_to_lists"`
	ListsToBlocks       []string               `yaml:"lists_to_blocks"`
	DefaultValues       map[string]interface{} `yaml:"default_values"`
	StructuralChanges   []StructuralChange     `yaml:"structural_changes"`
}

// StateTransformations defines state transformations
type StateTransformations struct {
	AttributeRenames map[string]string `yaml:"attribute_renames"`
	TypeConversions  []TypeConversion  `yaml:"type_conversions"`
	SchemaVersion    int               `yaml:"schema_version"`
}

// ConditionalRemoval defines when to remove an attribute based on conditions
type ConditionalRemoval struct {
	Attribute string                 `yaml:"attribute"`
	Condition map[string]interface{} `yaml:"condition"`
}

// TypeConversion defines how to convert attribute types
type TypeConversion struct {
	Attribute string `yaml:"attribute"`
	FromType  string `yaml:"from_type"`
	ToType    string `yaml:"to_type"`
	Pattern   string `yaml:"pattern,omitempty"`
}

// StructuralChange defines complex structural transformations
type StructuralChange struct {
	Type       string                 `yaml:"type"`
	Source     string                 `yaml:"source"`
	Target     string                 `yaml:"target"`
	Transform  string                 `yaml:"transform"`
	Parameters map[string]interface{} `yaml:"parameters,omitempty"`
}


// Migration provides base functionality for migrations
type Migration struct {
	config       MigrationConfig
	resourceType string
	sourceVer    string
	targetVer    string
}

// NewMigration creates a new base migration from YAML configuration
func NewMigration(yamlContent []byte) (*Migration, error) {
	var config MigrationConfig
	if err := yaml.Unmarshal(yamlContent, &config); err != nil {
		return nil, fmt.Errorf("failed to parse migration config: %w", err)
	}

	return &Migration{
		config:       config,
		resourceType: config.ResourceType,
		sourceVer:    config.SourceVersion,
		targetVer:    config.TargetVersion,
	}, nil
}

// ResourceType returns the resource type
func (m *Migration) ResourceType() string {
	return m.resourceType
}

// SourceVersion returns the source version
func (m *Migration) SourceVersion() string {
	return m.sourceVer
}

// TargetVersion returns the target version
func (m *Migration) TargetVersion() string {
	return m.targetVer
}

// MigrateConfig automatically applies configuration transformations based on YAML
func (m *Migration) MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error {
	// Create a TransformContext for the transformation functions
	transformCtx := &basic.TransformContext{}

	// Apply attribute renames using the common transformer
	if len(m.config.Config.AttributeRenames) > 0 {
		renamer := basic.AttributeRenamer(m.config.Config.AttributeRenames)
		if err := renamer(block, transformCtx); err != nil {
			ctx.AddError("Failed to rename attributes", err.Error(), m.resourceType)
			return err
		}
	}

	// Apply attribute removals using the common transformer
	if len(m.config.Config.AttributeRemovals) > 0 {
		remover := basic.AttributeRemover(m.config.Config.AttributeRemovals...)
		if err := remover(block, transformCtx); err != nil {
			ctx.AddError("Failed to remove attributes", err.Error(), m.resourceType)
			return err
		}
	}

	// Apply conditional removals using the common transformer
	for _, removal := range m.config.Config.ConditionalRemovals {
		condition := m.createConditionFunc(removal)
		condRemover := basic.ConditionalRemover(removal.Attribute, condition)
		if err := condRemover(block, transformCtx); err != nil {
			ctx.AddError("Failed to apply conditional removal", err.Error(), m.resourceType)
			return err
		}
	}

	// Apply blocks to lists conversions using the config transformer
	for _, blockType := range m.config.Config.BlocksToLists {
		converter := structural.BlocksToListConverter(blockType)
		if err := converter(block, transformCtx); err != nil {
			ctx.AddError("Failed to convert blocks to list", err.Error(), m.resourceType)
			return err
		}
	}

	// Apply lists to blocks conversions using the config transformer
	for _, attrName := range m.config.Config.ListsToBlocks {
		converter := structural.ListToBlocksConverter(attrName)
		if err := converter(block, transformCtx); err != nil {
			ctx.AddError("Failed to convert list to blocks", err.Error(), m.resourceType)
			return err
		}
	}

	// Apply default values using the common transformer
	if len(m.config.Config.DefaultValues) > 0 {
		setter := basic.DefaultValueSetter(m.config.Config.DefaultValues)
		if err := setter(block, transformCtx); err != nil {
			ctx.AddError("Failed to set default values", err.Error(), m.resourceType)
			return err
		}
	}

	// Apply type conversions
	for _, conversion := range m.config.Config.TypeConversions {
		if conversion.FromType == "set" && conversion.ToType == "list" {
			// Use the SetToListConverter for set to list conversions
			converter := basic.SetToListConverter(conversion.Attribute)
			if err := converter(block, transformCtx); err != nil {
				ctx.AddError("Failed to convert type", err.Error(), m.resourceType)
				return err
			}
		} else {
			// Handle other type conversions
			m.applyTypeConversion(block, conversion, ctx)
		}
	}

	// Apply structural changes
	for _, change := range m.config.Config.StructuralChanges {
		if err := m.applyStructuralChange(block, change, ctx); err != nil {
			ctx.AddWarning(
				fmt.Sprintf("Failed to apply structural change: %s", change.Type),
				err.Error(),
				m.resourceType,
			)
		}
	}


	return nil
}

// MigrateState automatically applies state transformations based on YAML
func (m *Migration) MigrateState(state map[string]interface{}, ctx *MigrationContext) error {
	// First apply config-level transformations that also affect state

	// Apply attribute removals from config section
	// Check both top-level and nested attributes
	if attrs, ok := state["attributes"].(map[string]interface{}); ok {
		for _, attrName := range m.config.Config.AttributeRemovals {
			delete(attrs, attrName)
		}
	} else {
		// Fallback to top-level if no nested attributes
		for _, attrName := range m.config.Config.AttributeRemovals {
			delete(state, attrName)
		}
	}

	// Apply conditional removals from config section
	for _, removal := range m.config.Config.ConditionalRemovals {
		// Check condition based on state values
		shouldRemove := true
		for condAttr, condValue := range removal.Condition {
			condValueStr := fmt.Sprintf("%v", condValue)

			// Check if this is a negative condition (value starts with !)
			if strings.HasPrefix(condValueStr, "!") {
				// Negative condition - remove if NOT equal
				expectedValue := strings.TrimPrefix(condValueStr, "!")
				if val, exists := state[condAttr]; exists {
					actualValue := fmt.Sprintf("%v", val)
					if actualValue == expectedValue {
						// Value matches what we DON'T want, so don't remove
						shouldRemove = false
						break
					}
				}
			} else {
				// Positive condition - remove if equal
				if val, exists := state[condAttr]; exists {
					actualValue := fmt.Sprintf("%v", val)
					if actualValue != condValueStr {
						// Value doesn't match condition, so don't remove
						shouldRemove = false
						break
					}
				} else {
					// Attribute doesn't exist, condition not met
					shouldRemove = false
					break
				}
			}
		}
		if shouldRemove {
			delete(state, removal.Attribute)
		}
	}

	// Apply default values from config section
	for attrName, defaultValue := range m.config.Config.DefaultValues {
		if _, exists := state[attrName]; !exists {
			state[attrName] = defaultValue
		}
	}

	// Now apply state-specific transformations

	// Apply attribute renames in state
	// State attributes are nested under "attributes" key
	if attrs, ok := state["attributes"].(map[string]interface{}); ok {
		for oldName, newName := range m.config.State.AttributeRenames {
			if val, exists := attrs[oldName]; exists {
				attrs[newName] = val
				delete(attrs, oldName)
			}
		}
	}

	// Apply type conversions in state
	for _, conversion := range m.config.State.TypeConversions {
		m.applyStateTypeConversion(state, conversion, ctx)
	}

	// Update schema version if specified
	if m.config.State.SchemaVersion > 0 {
		state["schema_version"] = m.config.State.SchemaVersion
	}

	return nil
}

// Validate checks if the migration can be applied
func (m *Migration) Validate(block *hclwrite.Block) error {
	// Basic validation - check if resource type matches
	return nil
}

// RequiresFileTransformation returns whether this migration needs file-level processing
func (m *Migration) RequiresFileTransformation() bool {
	return m.config.RequiresFileTransformation
}

// createConditionFunc creates a condition function for ConditionalRemover
func (m *Migration) createConditionFunc(removal ConditionalRemoval) func(*hclwrite.Block) bool {
	return func(block *hclwrite.Block) bool {
		body := block.Body()

		// Check each condition
		for attrName, expectedValue := range removal.Condition {
			expected := fmt.Sprintf("%v", expectedValue)

			// Check if this is a negated condition (!value means NOT equal)
			negated := strings.HasPrefix(expected, "!")
			if negated {
				expected = strings.TrimPrefix(expected, "!")
			}

			attr := body.GetAttribute(attrName)
			if attr == nil {
				// If condition attribute doesn't exist:
				// - For negated conditions, we should remove (attribute missing != expected value)
				// - For normal conditions, we shouldn't remove (can't match if missing)
				return negated
			}

			// Get the actual attribute value
			attrValue := strings.TrimSpace(string(attr.Expr().BuildTokens(nil).Bytes()))
			// Remove quotes if present
			attrValue = strings.Trim(attrValue, `"`)

			// Check if values match
			matches := attrValue == expected

			// For negated conditions, remove if NOT equal
			// For normal conditions, remove if equal
			if negated {
				if matches {
					return false // Don't remove if type IS app_launcher
				}
			} else {
				if !matches {
					return false // Don't remove if condition doesn't match
				}
			}
		}

		return true // All conditions met, remove the attribute
	}
}

// applyTypeConversion applies a type conversion to an attribute for non-standard conversions
func (m *Migration) applyTypeConversion(block *hclwrite.Block, conversion TypeConversion, ctx *MigrationContext) {
	body := block.Body()
	attr := body.GetAttribute(conversion.Attribute)

	if attr == nil {
		return
	}

	// Handle non-standard type conversions here
	// Standard set->list is handled by SetToListConverter
	switch conversion.ToType {
	case "string":
		if conversion.FromType == "list" {
			// Convert single-element list to string
			// This would need proper HCL parsing implementation
			ctx.AddWarning(
				"Type conversion not fully implemented",
				fmt.Sprintf("List to string conversion for %s needs manual review", conversion.Attribute),
				m.resourceType,
			)
		}
	default:
		ctx.AddWarning(
			"Unknown type conversion",
			fmt.Sprintf("Conversion from %s to %s for %s is not supported", conversion.FromType, conversion.ToType, conversion.Attribute),
			m.resourceType,
		)
	}
}

// applyStateTypeConversion applies type conversion in state
func (m *Migration) applyStateTypeConversion(state map[string]interface{}, conversion TypeConversion, ctx *MigrationContext) {
	// Check both top-level and nested attributes
	attrs, hasAttrs := state["attributes"].(map[string]interface{})
	
	// Try to find the attribute in nested attributes first
	var val interface{}
	var exists bool
	var targetMap map[string]interface{}
	
	if hasAttrs {
		val, exists = attrs[conversion.Attribute]
		targetMap = attrs
	}
	
	// Fall back to top-level if not found in attributes
	if !exists {
		val, exists = state[conversion.Attribute]
		targetMap = state
	}
	
	if exists {
		switch conversion.ToType {
		case "list":
			if conversion.FromType == "set" {
				// Sets and lists are often the same in JSON state
				// Just ensure it's an array
				if _, ok := val.([]interface{}); !ok {
					targetMap[conversion.Attribute] = []interface{}{val}
				}
			}
		case "string":
			if conversion.FromType == "list" {
				// Convert single-element list to string
				if list, ok := val.([]interface{}); ok && len(list) == 1 {
					targetMap[conversion.Attribute] = list[0]
				}
			}
		}
	}
}

// applyStructuralChange applies complex structural changes
func (m *Migration) applyStructuralChange(block *hclwrite.Block, change StructuralChange, ctx *MigrationContext) error {
	// Handle transform-based structural changes
	if change.Transform != "" {
		return m.applyTransform(block, change, ctx)
	}

	// Handle type-based structural changes
	switch change.Type {
	case "split_object":
		return m.applySplitObject(block, change, ctx)
	case "merge_attributes":
		return m.applyMergeAttributes(block, change, ctx)
	case "nested_restructure":
		return m.applyNestedRestructure(block, change, ctx)
	case "policies_transformation":
		// Handle legacy naming for backward compatibility
		return m.applyTransform(block, change, ctx)
	default:
		return fmt.Errorf("unknown structural change type: %s", change.Type)
	}
}

// applySplitObject splits an object attribute into multiple attributes
func (m *Migration) applySplitObject(block *hclwrite.Block, change StructuralChange, ctx *MigrationContext) error {
	transformCtx := &basic.TransformContext{}

	// Extract parameters
	attributes := []string{}
	if attrs, ok := change.Parameters["attributes"].([]interface{}); ok {
		for _, attr := range attrs {
			if s, ok := attr.(string); ok {
				attributes = append(attributes, s)
			}
		}
	}

	// Check for attribute mapping
	if attrMap, ok := change.Parameters["attribute_map"].(map[string]interface{}); ok {
		// Use mapping version
		mappedAttrs := make(map[string]string)
		for k, v := range attrMap {
			if s, ok := v.(string); ok {
				mappedAttrs[k] = s
			}
		}
		transformer := structural.SplitObjectWithMapping(change.Source, mappedAttrs)
		return transformer(block, transformCtx)
	}

	// Use simple split with optional prefix
	prefix := ""
	if p, ok := change.Parameters["prefix"].(string); ok {
		prefix = p
	}

	transformer := structural.SplitObjectTransformer(change.Source, attributes, prefix)
	return transformer(block, transformCtx)
}

// applyMergeAttributes merges multiple attributes into one
func (m *Migration) applyMergeAttributes(block *hclwrite.Block, change StructuralChange, ctx *MigrationContext) error {
	transformCtx := &basic.TransformContext{}

	// Check for attribute mapping
	if attrMap, ok := change.Parameters["attribute_map"].(map[string]interface{}); ok {
		// Use mapping version
		mappedAttrs := make(map[string]string)
		for k, v := range attrMap {
			if s, ok := v.(string); ok {
				mappedAttrs[k] = s
			}
		}
		transformer := structural.MergeAttributesWithMapping(change.Target, mappedAttrs)
		return transformer(block, transformCtx)
	}

	// Extract source attributes list
	sourceAttrs := []string{}
	if attrs, ok := change.Parameters["source_attributes"].([]interface{}); ok {
		for _, attr := range attrs {
			if s, ok := attr.(string); ok {
				sourceAttrs = append(sourceAttrs, s)
			}
		}
	}

	// Get format (default to "object")
	format := "object"
	if f, ok := change.Parameters["format"].(string); ok {
		format = f
	}

	transformer := structural.MergeAttributesTransformer(change.Target, sourceAttrs, format)
	return transformer(block, transformCtx)
}

// applyNestedRestructure restructures nested attributes
func (m *Migration) applyNestedRestructure(block *hclwrite.Block, change StructuralChange, ctx *MigrationContext) error {
	transformCtx := &basic.TransformContext{}

	// Check if this is a path-based restructure
	if paths, ok := change.Parameters["paths"].(map[string]interface{}); ok {
		// Convert to string map
		pathMap := make(map[string]string)
		for k, v := range paths {
			if s, ok := v.(string); ok {
				pathMap[k] = s
			}
		}
		transformer := structural.NestedRestructureTransformer(change.Source, pathMap)
		return transformer(block, transformCtx)
	}

	// Otherwise, this might be a flatten operation
	// Check for flatten parameters
	separator := "_"
	if sep, ok := change.Parameters["separator"].(string); ok {
		separator = sep
	}

	maxDepth := 2
	if depth, ok := change.Parameters["depth"].(int); ok {
		maxDepth = depth
	}

	// Check for prefix
	if prefix, ok := change.Parameters["prefix"].(string); ok && prefix != "" {
		transformer := structural.FlattenWithPrefix(change.Source, prefix, separator)
		return transformer(block, transformCtx)
	}

	// Use regular flatten
	transformer := structural.FlattenNestedTransformer(change.Source, separator, maxDepth)
	return transformer(block, transformCtx)
}

// applyTransform applies transformation-based structural changes
func (m *Migration) applyTransform(block *hclwrite.Block, change StructuralChange, ctx *MigrationContext) error {
	transformCtx := &basic.TransformContext{}

	switch change.Transform {
	case "string_list_to_object_list":
		// Handle string list to object list transformation
		objectKey := "id" // default
		if key, ok := change.Parameters["object_key"].(string); ok {
			objectKey = key
		}

		// Check if we need to do resource renaming
		if updateRefs, ok := change.Parameters["update_references"].([]interface{}); ok && len(updateRefs) > 0 {
			if ref, ok := updateRefs[0].(map[string]interface{}); ok {
				from := ""
				to := ""
				if fromVal, ok := ref["from"].(string); ok {
					from = fromVal
				}
				if toVal, ok := ref["to"].(string); ok {
					to = toVal
				}

				if from != "" && to != "" {
					// Use the StringArrayToObjectArrayWithRename transformer
					transformer := structural.StringArrayToObjectArrayWithRename(change.Source, objectKey, from, to)
					return transformer(block, transformCtx)
				}
			}
		}

		// No renaming needed, use simple transformer
		transformer := structural.ArrayToObjectArrayConverter(change.Source, objectKey)
		return transformer(block, transformCtx)

	case "object_to_list":
		// Convert nested blocks to a list attribute
		converter := structural.BlocksToListConverter(change.Source)
		return converter(block, transformCtx)

	case "list_to_blocks":
		// Convert list attribute to blocks
		// Check for attribute mapping
		if attrMap, ok := change.Parameters["attribute_map"].(map[string]interface{}); ok {
			// Use mapping version
			mappedAttrs := make(map[string]string)
			for k, v := range attrMap {
				if s, ok := v.(string); ok {
					mappedAttrs[k] = s
				}
			}
			transformer := structural.ListToBlocksWithMapping(change.Source, mappedAttrs)
			return transformer(block, transformCtx)
		}

		// Use simple converter
		transformer := structural.ListToBlocksConverter(change.Source)
		return transformer(block, transformCtx)

	case "flatten_nested":
		// Use the nested restructure handler which includes flatten functionality
		return m.applyNestedRestructure(block, change, ctx)

	default:
		return fmt.Errorf("unknown transform type: %s", change.Transform)
	}
}
