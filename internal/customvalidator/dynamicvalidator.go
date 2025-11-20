package customvalidator

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	t "github.com/cloudflare/terraform-provider-cloudflare/internal/types"
)

var _ validator.Dynamic = subtypesValidator{}

type subtypesValidator struct {
	allowedTypes []attr.Type
}

func compatible(ctx context.Context, ty attr.Type, val attr.Value) bool {
	switch tty := ty.(type) {
	case basetypes.Int32Typable:
		if ok, _ := t.IntValue(val); ok {
			return true
		}
	case basetypes.Int64Typable:
		if ok, _ := t.IntValue(val); ok {
			return true
		}
	case basetypes.Float32Typable:
		if ok, _ := t.FloatValue(val); ok {
			return true
		}
	case basetypes.Float64Typable:
		if ok, _ := t.FloatValue(val); ok {
			return true
		}
	case basetypes.NumberTypable:
		if ok, _ := t.FloatValue(val); ok {
			return true
		}
	case basetypes.TupleType:
		if ok, items := t.ChildItems(val); ok {
			return slices.CompareFunc(tty.ElemTypes, items, func(lhs attr.Type, rhs attr.Value) int {
				if compatible(ctx, lhs, rhs) {
					return 0
				}
				return 1
			}) == 0
		}
	case basetypes.ListType:
		if ok, items := t.ChildItems(val); ok {
			return !slices.ContainsFunc(items, func(value attr.Value) bool {
				return !compatible(ctx, tty.ElemType, value)
			})
		}
	case basetypes.SetType:
		if ok, items := t.ChildItems(val); ok {
			return !slices.ContainsFunc(items, func(value attr.Value) bool {
				return !compatible(ctx, tty.ElemType, value)
			})
		}
	case basetypes.MapType:
		if ok, attrs := t.ChildAttributes(val); ok {
			for _, v := range attrs {
				if !compatible(ctx, tty.ElemType, v) {
					return false
				}
			}
			return true
		}
	case basetypes.ObjectType:
		if ok, attrs := t.ChildAttributes(val); ok {
			if len(attrs) != len(tty.AttrTypes) {
				return false
			}
			for name, attrType := range tty.AttrTypes {
				actualValue, exists := attrs[name]
				if !exists || !compatible(ctx, attrType, actualValue) {
					return false
				}
			}
			return true
		}
	default:
	}

	return ty.Equal(val.Type(ctx))
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
