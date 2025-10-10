package generator

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/zclconf/go-cty/cty"
)

// GenericHCLGenerator provides helper functions for generating HCL
// This is a utility class that resources can use, but each resource
// should implement its own generation logic
type GenericHCLGenerator struct {
	// Options
	PreserveOrder   bool // Try to preserve attribute order
	IncludeComments bool // Add migration comments
	CompactFormat   bool // Use compact formatting
}

// NewGenericHCLGenerator creates a new generic generator with helpers
func NewGenericHCLGenerator() *GenericHCLGenerator {
	return &GenericHCLGenerator{
		PreserveOrder:   true,
		IncludeComments: false,
		CompactFormat:   false,
	}
}

// AddStringAttribute adds a string attribute if it has a value
func (g *GenericHCLGenerator) AddStringAttribute(body *hclwrite.Body, name string, value types.String) {
	if !value.IsNull() && !value.IsUnknown() {
		body.SetAttributeValue(name, cty.StringVal(value.ValueString()))
	}
}

// AddFloat64Attribute adds a float64 attribute if it has a value
func (g *GenericHCLGenerator) AddFloat64Attribute(body *hclwrite.Body, name string, value types.Float64) {
	if !value.IsNull() && !value.IsUnknown() {
		body.SetAttributeValue(name, cty.NumberFloatVal(value.ValueFloat64()))
	}
}

// AddBoolAttribute adds a boolean attribute if it has a value
func (g *GenericHCLGenerator) AddBoolAttribute(body *hclwrite.Body, name string, value types.Bool) {
	if !value.IsNull() && !value.IsUnknown() {
		body.SetAttributeValue(name, cty.BoolVal(value.ValueBool()))
	}
}

// AddStringArrayAttribute adds a string array attribute if it has values
func (g *GenericHCLGenerator) AddStringArrayAttribute(body *hclwrite.Body, name string, values []types.String) {
	if len(values) == 0 {
		return
	}

	ctyValues := make([]cty.Value, len(values))
	for i, v := range values {
		if !v.IsNull() && !v.IsUnknown() {
			ctyValues[i] = cty.StringVal(v.ValueString())
		} else {
			ctyValues[i] = cty.StringVal("")
		}
	}

	body.SetAttributeValue(name, cty.ListVal(ctyValues))
}

// CreateBlock creates a new HCL block with the given type and labels
func (g *GenericHCLGenerator) CreateBlock(blockType string, labels []string) *hclwrite.Block {
	return hclwrite.NewBlock(blockType, labels)
}

// AddBlock appends a new block to the body
func (g *GenericHCLGenerator) AddBlock(body *hclwrite.Body, blockType string, labels []string) *hclwrite.Block {
	return body.AppendNewBlock(blockType, labels)
}
