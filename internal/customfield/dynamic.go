package customfield

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	t "github.com/cloudflare/terraform-provider-cloudflare/internal/types"
)

var (
	_ basetypes.DynamicTypable                    = (*NormalizedDynamicType)(nil)
	_ basetypes.DynamicValuableWithSemanticEquals = (*NormalizedDynamicValue)(nil)
	_ planmodifier.Dynamic                        = (*NormalizingDynamicPlanModifier)(nil)
)

type NormalizedDynamicType struct {
	basetypes.DynamicType
}

func (t NormalizedDynamicType) ValueFromDynamic(ctx context.Context, in types.Dynamic) (basetypes.DynamicValuable, diag.Diagnostics) {
	return RawNormalizedDynamicValue(in), nil
}

func (t NormalizedDynamicType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.DynamicType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	dynValue, ok := attrValue.(types.Dynamic)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	dynValuable, diags := t.ValueFromDynamic(ctx, dynValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting DynamicValue to DynamicValuableWithSemanticEquals: %v", diags)
	}

	return dynValuable, nil

}

func (t NormalizedDynamicType) ValueType(context.Context) attr.Value {
	return NormalizedDynamicValue{}
}

func (t NormalizedDynamicType) Equal(o attr.Type) bool {
	other, ok := o.(NormalizedDynamicType)
	if !ok {
		return false
	}

	return t.DynamicType.Equal(other.DynamicType)
}

func (t NormalizedDynamicType) String() string {
	return "Normalized" + t.DynamicType.String()
}

type NormalizedDynamicValue struct {
	types.Dynamic
}

func (v NormalizedDynamicValue) Type(context.Context) attr.Type {
	return NormalizedDynamicType{}
}

func RawNormalizedDynamicValue(in types.Dynamic) NormalizedDynamicValue {
	return NormalizedDynamicValue{in}
}

func RawNormalizedDynamicValueFrom(in attr.Value) NormalizedDynamicValue {
	return NormalizedDynamicValue{basetypes.NewDynamicValue(in)}
}

func (v NormalizedDynamicValue) ToDynamicValue(ctx context.Context) (types.Dynamic, diag.Diagnostics) {
	return v.Dynamic, nil
}

func semanticEquals(ctx context.Context, lhs attr.Value, rhs attr.Value) (eq bool, diag diag.Diagnostics) {
	if lhs == nil || rhs == nil {
		return lhs == rhs, nil
	}

	if (lhs.Equal(rhs)) || (lhs.IsNull() && rhs.IsNull()) || (lhs.IsUnknown() && rhs.IsUnknown()) {
		return true, nil
	}

	if l, ok := lhs.(basetypes.DynamicValuable); ok {
		if r, ok := rhs.(basetypes.DynamicValuable); ok {
			ld, d := l.ToDynamicValue(ctx)
			diag.Append(d...)
			rd, d := r.ToDynamicValue(ctx)
			diag.Append(d...)
			lv, rv := ld.UnderlyingValue(), rd.UnderlyingValue()

			return semanticEquals(ctx, lv, rv)
		}
	}

	if ok, lvalue := t.IntValue(lhs); ok {
		if ok, rvalue := t.IntValue(rhs); ok {
			if lvalue.Cmp(rvalue) == 0 {
				return true, diag
			}
		}
	}

	if ok, lvalue := t.FloatValue(lhs); ok {
		if ok, rvalue := t.FloatValue(rhs); ok {
			if lvalue.Cmp(rvalue) == 0 {
				return true, diag
			}
		}
	}

	// in terraform a list of primitives below a certain length is considered a tuple
	// tuple: `[1, 2]`, list `tolist([1, 2])`, set `toset([1, 2])`
	if ok, lvalues := t.ChildItems(lhs); ok {
		if ok, rvalues := t.ChildItems(rhs); ok {
			eq := slices.EqualFunc(lvalues, rvalues,
				func(l attr.Value, r attr.Value) bool {
					e, d := semanticEquals(ctx, l, r)
					diag.Append(d...)
					return e
				})

			if eq {
				return true, diag
			}
		}
	}

	// object value `{a = 2}` and map value `tomap({ a = 2 })` should be similar to how tuple and lists behave
	if ok, lvalues := t.ChildAttributes(lhs); ok {
		if ok, rvalues := t.ChildAttributes(rhs); ok {
			eq := maps.EqualFunc(lvalues, rvalues,
				func(l attr.Value, r attr.Value) bool {
					e, d := semanticEquals(ctx, l, r)
					diag.Append(d...)
					return e
				})

			if eq {
				return true, diag
			}
		}
	}

	return false, diag
}

func (v NormalizedDynamicValue) DynamicSemanticEquals(ctx context.Context, other basetypes.DynamicValuable) (eq bool, diag diag.Diagnostics) {
	return semanticEquals(ctx, v, other)
}

type NormalizingDynamicPlanModifier struct{}

func (m NormalizingDynamicPlanModifier) Description(ctx context.Context) string {
	return ""
}

func (m NormalizingDynamicPlanModifier) MarkdownDescription(ctx context.Context) string {
	return ""
}

func validate(ctx context.Context, value attr.Value) (diags diag.Diagnostics) {
	if val, ok := value.(basetypes.DynamicValuable); ok {
		v, d := val.ToDynamicValue(ctx)
		diags.Append(d...)
		return validate(ctx, v.UnderlyingValue())
	}

	if _, ok := value.(basetypes.MapValue); ok {
		diags.AddError("invalid dynamic type", "due to Terraform limitations map types are not currently supported in dynamic values, you can work around this using `jsonencode(jsondecode(...))`")
		return
	}

	if _, ok := value.(basetypes.SetValue); ok {
		diags.AddError("invalid dynamic type", "due to Terraform limitations set types are not currently supported in dynamic values, you can work around this using `tolist(...)`")
		return
	}

	if ok, values := t.ChildItems(value); ok {
		for _, val := range values {
			diags.Append(validate(ctx, val)...)
		}
	}

	if ok, values := t.ChildAttributes(value); ok {
		for _, val := range values {
			diags.Append(validate(ctx, val)...)
		}
	}

	return
}

func (m NormalizingDynamicPlanModifier) PlanModifyDynamic(ctx context.Context, req planmodifier.DynamicRequest, resp *planmodifier.DynamicResponse) {
	resp.Diagnostics.Append(validate(ctx, req.PlanValue)...)
	if resp.Diagnostics.HasError() {
		return
	}

	eq, d := semanticEquals(ctx, req.PlanValue, req.StateValue)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	if eq {
		resp.PlanValue = req.StateValue
	}
}

func NormalizeDynamicPlanModifier() planmodifier.Dynamic {
	return NormalizingDynamicPlanModifier{}
}
