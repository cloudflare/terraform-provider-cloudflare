package customfield_test

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func P[T any](v T) *T { return &v }

type NestedStruct struct {
	A           types.String                                `tfsdk:"a" json:"a"`
	B           types.String                                `tfsdk:"b" json:"b"`
	InnerStruct customfield.NestedObject[NestedInnerStruct] `tfsdk:"inner_struct" json:"inner_struct,computed"`
}

type NestedInnerStruct struct {
	C types.String `tfsdk:"c" json:"c"`
}

func TestPlanReadAndWrite(t *testing.T) {
	t.Run("nested struct", func(t *testing.T) {

		innerType := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"c": tftypes.String,
			},
		}

		plan := tfsdk.Plan{
			Raw: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a":            tftypes.String,
					"b":            tftypes.String,
					"inner_struct": innerType,
				},
			}, map[string]tftypes.Value{
				"a":            tftypes.NewValue(tftypes.String, "foo"),
				"b":            tftypes.NewValue(tftypes.String, nil),
				"inner_struct": tftypes.NewValue(innerType, tftypes.UnknownValue),
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
						CustomType: customfield.NewNestedObjectType[NestedInnerStruct](context.TODO()),
						Computed:   true,
						Attributes: map[string]schema.Attribute{
							"c": schema.StringAttribute{
								Required: true,
							},
						},
					},
				},
			},
		}

		expected := NestedStruct{
			A:           types.StringValue("foo"),
			B:           types.StringNull(),
			InnerStruct: customfield.UnknownObject[NestedInnerStruct](context.TODO()),
		}

		var initial *NestedStruct

		// get the struct from the plan
		diags := plan.Get(context.TODO(), &initial)
		EnsurePlanEquals(t, diags, &expected, initial)

		initial.A = types.StringValue("bar")
		o, _ := customfield.NewObject(context.TODO(), P(NestedInnerStruct{C: types.StringValue("baz")}))
		initial.InnerStruct = o

		// expect that we can save back into the plan
		diags = plan.Set(context.TODO(), initial)
		ExpectNoDiagnostics(t, diags)

		// expect that we can get it back out correctly
		var updated *NestedStruct
		diags = plan.Get(context.TODO(), &updated)
		EnsurePlanEquals(t, diags, initial, updated)
	})
}

func EnsurePlanEquals(t *testing.T, diags diag.Diagnostics, expected interface{}, actual interface{}) {
	ExpectNoDiagnostics(t, diags)

	if !reflect.DeepEqual(actual, expected) {
		spew.Dump(actual, expected)
		t.Fatalf("expected \n%#v\nbut got\n%#v", actual, expected)
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
