package customvalidator_test

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var ctx = context.TODO()

var testcases = map[string](struct {
	validator validator.Dynamic
	valid     []attr.Value
	invalid   []attr.Value
}){
	"int64": {
		customvalidator.AllowedSubtypes(basetypes.Int64Type{}),
		[]attr.Value{
			basetypes.NewInt32Value(0),
			basetypes.NewInt64Value(0),
			basetypes.NewNumberValue(big.NewFloat(0)),
		},
		[]attr.Value{
			basetypes.NewFloat32Value(0),
			basetypes.NewFloat64Value(0),
			basetypes.NewNumberValue(big.NewFloat(1.1)),
		},
	},
	"float64": {
		customvalidator.AllowedSubtypes(basetypes.Float64Type{}),
		[]attr.Value{
			basetypes.NewInt32Value(0),
			basetypes.NewInt64Value(0),
			basetypes.NewFloat32Value(0),
			basetypes.NewFloat64Value(0),
			basetypes.NewNumberValue(big.NewFloat(0)),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
		},
	},
}

func runRequest(v validator.Dynamic, val attr.Value) diag.Diagnostics {
	req := validator.DynamicRequest{
		ConfigValue: basetypes.NewDynamicValue(val),
	}
	rsp := validator.DynamicResponse{}
	v.ValidateDynamic(ctx, req, &rsp)
	return rsp.Diagnostics
}

func TestNumberValidators(t *testing.T) {
	for name, testcase := range testcases {
		t.Run(name, func(t *testing.T) {
			for _, val := range testcase.valid {
				d := runRequest(testcase.validator, val)
				if d.HasError() {
					t.Fatalf("Failed to validate: %+v", val.Type(ctx))
				}
			}
			for _, val := range testcase.invalid {
				d := runRequest(testcase.validator, val)
				if !d.HasError() {
					t.Fatalf("Failed to invalidate: %+v", val.Type(ctx))
				}
			}
		})
	}
}
