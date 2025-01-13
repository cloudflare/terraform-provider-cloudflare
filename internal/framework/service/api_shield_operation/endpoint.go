package apishieldoperation

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var varMatch = regexp.MustCompile(`{([\w\d-_]+)}`)

// EndpointType implemented based on https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/custom

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = EndpointType{}
var _ basetypes.StringValuableWithSemanticEquals = EndpointValue{}

type EndpointType struct {
	basetypes.StringType
}

func (t EndpointType) Equal(o attr.Type) bool {
	other, ok := o.(EndpointType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t EndpointType) String() string {
	return "EndpointType"
}

func (t EndpointType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	value := EndpointValue{
		StringValue: in,
	}

	return value, nil
}

func (t EndpointType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t EndpointType) ValueType(ctx context.Context) attr.Value {
	return EndpointValue{}
}

var _ basetypes.StringValuable = EndpointValue{}

type EndpointValue struct {
	basetypes.StringValue
}

func (v EndpointValue) Equal(o attr.Value) bool {
	other, ok := o.(EndpointValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v EndpointValue) Type(ctx context.Context) attr.Type {
	return EndpointType{}
}

func NewEndpointValue(value string) EndpointValue {
	if value == "" {
		return EndpointValue{types.StringNull()}
	}
	return EndpointValue{types.StringValue(value)}
}

func (v EndpointValue) StringSemanticEquals(ctx context.Context, o basetypes.StringValuable) (bool, diag.Diagnostics) {
	oStrVal, diag := o.ToStringValue(ctx)
	result := varMatch.ReplaceAllString(v.StringValue.ValueString(), "{var}") == varMatch.ReplaceAllString(oStrVal.ValueString(), "{var}")
	return result, diag
}
