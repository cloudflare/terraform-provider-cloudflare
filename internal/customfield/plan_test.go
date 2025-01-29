package customfield_test

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
)

func P[T any](v T) *T { return &v }

type NestedStruct struct {
	A           types.String                                    `tfsdk:"a" json:"a"`
	B           types.String                                    `tfsdk:"b" json:"b"`
	InnerStruct customfield.NestedObject[NestedInnerStruct]     `tfsdk:"inner_struct" json:"inner_struct,computed"`
	PrimList    customfield.List[basetypes.StringValue]         `tfsdk:"prim_list" json:"prim_list,computed"`
	PrimSet     customfield.Set[basetypes.StringValue]          `tfsdk:"prim_set" json:"prim_set,computed"`
	StructList  customfield.NestedObjectList[NestedInnerStruct] `tfsdk:"struct_list" json:"struct_list,computed"`
	StructSet   customfield.NestedObjectSet[NestedIntStruct]    `tfsdk:"struct_set" json:"struct_set,computed"`
}

type NestedInnerStruct struct {
	C types.String `tfsdk:"c" json:"c"`
}
type NestedIntStruct struct {
	D types.Int64 `tfsdk:"d" json:"d"`
}

func TestPlanReadAndWrite(t *testing.T) {
	spew.Config.ContinueOnMethod = true
	ctx := context.TODO()
	t.Run("nested struct", func(t *testing.T) {

		innerType := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"c": tftypes.String,
			},
		}
		structListType := tftypes.List{ElementType: innerType}

		innerIntType := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"d": tftypes.Number,
			},
		}

		structSetType := tftypes.Set{ElementType: innerIntType}

		plan := tfsdk.Plan{
			Raw: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a":            tftypes.String,
					"b":            tftypes.String,
					"inner_struct": innerType,
					"prim_list":    tftypes.List{ElementType: tftypes.String},
					"prim_set":     tftypes.Set{ElementType: tftypes.String},
					"struct_list":  structListType,
					"struct_set":   structSetType,
				},
			}, map[string]tftypes.Value{
				"a":            tftypes.NewValue(tftypes.String, "foo"),
				"b":            tftypes.NewValue(tftypes.String, nil),
				"inner_struct": tftypes.NewValue(innerType, tftypes.UnknownValue),
				"prim_list":    tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, tftypes.UnknownValue),
				"prim_set":     tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, tftypes.UnknownValue),
				"struct_list":  tftypes.NewValue(structListType, tftypes.UnknownValue),
				"struct_set":   tftypes.NewValue(structSetType, tftypes.UnknownValue),
			}),
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"a": schema.StringAttribute{
						Required: true,
					},
					"b": schema.StringAttribute{
						Optional: true,
					},
					"inner_struct": schema.SingleNestedAttribute{
						CustomType: customfield.NewNestedObjectType[NestedInnerStruct](ctx),
						Computed:   true,
						Attributes: map[string]schema.Attribute{
							"c": schema.StringAttribute{
								Required: true,
							},
						},
					},
					"prim_list": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[basetypes.StringValue](ctx),
						ElementType: types.StringType,
					},
					"prim_set": schema.SetAttribute{
						Computed:    true,
						CustomType:  customfield.NewSetType[basetypes.StringValue](ctx),
						ElementType: types.StringType,
					},
					"struct_list": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[NestedInnerStruct](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"c": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
					"struct_set": schema.SetNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectSetType[NestedIntStruct](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"d": schema.Int64Attribute{
									Required: true,
								},
							},
						},
					},
				},
			},
		}

		expected := NestedStruct{
			A:           types.StringValue("foo"),
			B:           types.StringNull(),
			InnerStruct: customfield.UnknownObject[NestedInnerStruct](ctx),
			PrimList:    customfield.UnknownList[basetypes.StringValue](ctx),
			PrimSet:     customfield.UnknownSet[basetypes.StringValue](ctx),
			StructList:  customfield.UnknownObjectList[NestedInnerStruct](ctx),
			StructSet:   customfield.UnknownObjectSet[NestedIntStruct](ctx),
		}

		var initial *NestedStruct

		// get the struct from the plan
		diags := plan.Get(ctx, &initial)
		EnsurePlanEquals(t, diags, &expected, initial)

		initial.A = types.StringValue("bar")
		o, _ := customfield.NewObject(ctx, P(NestedInnerStruct{C: types.StringValue("baz")}))
		initial.InnerStruct = o

		initial.PrimList = customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("quux")})
		initial.PrimSet = customfield.NewSetMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("quux2")})
		initial.StructList = customfield.NewObjectListMust(ctx, []NestedInnerStruct{{C: types.StringValue("qux")}})
		initial.StructSet = customfield.NewObjectSetMust(ctx, []NestedIntStruct{{D: types.Int64Value(42)}})

		// expect that we can save back into the plan
		diags = plan.Set(ctx, initial)
		ExpectNoDiagnostics(t, diags)

		// expect that we can get it back out correctly
		var updated *NestedStruct
		diags = plan.Get(ctx, &updated)
		EnsurePlanEquals(t, diags, initial, updated)
	})
}

func EnsurePlanEquals(t *testing.T, diags diag.Diagnostics, expected interface{}, actual interface{}) {
	ExpectNoDiagnostics(t, diags)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected \n%s\nbut got\n%s", spew.Sdump(actual), spew.Sdump(expected))
	}
}

func ExpectNoDiagnostics(t *testing.T, diags diag.Diagnostics) {
	if diags.HasError() {
		t.Fatalf("unexpected error: %v", errorFromDiagnostics(diags).Error())
	}
}

func errorFromDiagnostics(diags diag.Diagnostics) error {
	if diags == nil {
		return nil
	}
	messages := []string{}
	for _, err := range diags {
		messages = append(messages, err.Summary())
		messages = append(messages, err.Detail())
	}
	return errors.New(strings.Join(messages, " "))
}
