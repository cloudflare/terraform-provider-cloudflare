package customvalidator

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ validator.Dynamic = subtypesValidator{}

type subtypesValidator struct {
	allowedTypes []attr.Type
}

func compatibleTypes(ty attr.Type) (types []attr.Type) {
	types = append(types, ty)
	switch ty.(type) {
	case basetypes.Int64Type:
		types = append(types, basetypes.Int32Type{})
	case basetypes.Float64Type:
		types = append(types, basetypes.Int32Type{}, basetypes.Int64Type{}, basetypes.Float32Type{}, basetypes.NumberType{})
	case basetypes.NumberType:
		types = append(types, basetypes.Int32Type{}, basetypes.Int64Type{}, basetypes.Float32Type{}, basetypes.Float64Type{})
	}
	return
}

func compatible(ctx context.Context, ty attr.Type, val attr.Value) bool {
	valTy := val.Type(ctx)
	if slices.ContainsFunc(compatibleTypes(ty), valTy.Equal) {
		return true
	}

	if v, ok := val.(basetypes.NumberValue); ok {
		big := v.ValueBigFloat()
		switch ty.(type) {
		case basetypes.Int32Type:
			return big.IsInt()
		case basetypes.Int64Type:
			return big.IsInt()
		}
	}

	return false
}

func (v subtypesValidator) ValidateDynamic(ctx context.Context, req validator.DynamicRequest, resp *validator.DynamicResponse) {
	dynamic := req.ConfigValue
	if dynamic.IsNull() || dynamic.IsUnknown() || dynamic.IsUnderlyingValueNull() || dynamic.IsUnderlyingValueUnknown() {
		return
	}
	value := dynamic.UnderlyingValue()

	if slices.ContainsFunc(v.allowedTypes, func(ty attr.Type) bool { return compatible(ctx, ty, value) }) {
		return
	}

	detail := fmt.Sprintf("%s Received: %T", v.Description(ctx), value.Type(ctx))
	resp.Diagnostics.AddAttributeError(req.Path, "Value is not one of the allowed types", detail)
}

func (v subtypesValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v subtypesValidator) MarkdownDescription(_ context.Context) string {
	var s []string
	for _, t := range v.allowedTypes {
		s = append(s, t.String())
	}
	return fmt.Sprintf("The following types are allowed: %s.", strings.Join(s, ", "))
}

func AllowedSubtypes(types ...attr.Type) validator.Dynamic {
	return subtypesValidator{allowedTypes: types}
}
