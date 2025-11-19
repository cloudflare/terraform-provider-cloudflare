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
			basetypes.NewInt32Value(-42),
			basetypes.NewInt64Value(9223372036854775807), // max int64
			basetypes.NewFloat32Value(0),
			basetypes.NewFloat64Value(0),
			basetypes.NewFloat32Value(42.0),   // whole number float
			basetypes.NewFloat64Value(-100.0), // whole number float
			basetypes.NewNumberValue(big.NewFloat(0)),
			basetypes.NewNumberValue(big.NewFloat(123)), // whole number
		},
		[]attr.Value{
			basetypes.NewFloat32Value(0.1),
			basetypes.NewFloat64Value(0.1),
			basetypes.NewFloat32Value(-3.14),
			basetypes.NewFloat64Value(2.718),
			basetypes.NewNumberValue(big.NewFloat(1.1)),
			basetypes.NewNumberValue(big.NewFloat(-0.5)),
		},
	},
	"float64": {
		customvalidator.AllowedSubtypes(basetypes.Float64Type{}),
		[]attr.Value{
			basetypes.NewInt32Value(0),
			basetypes.NewInt64Value(0),
			basetypes.NewInt32Value(-123),
			basetypes.NewInt64Value(456789),
			basetypes.NewFloat32Value(0),
			basetypes.NewFloat64Value(0),
			basetypes.NewFloat32Value(3.14159),
			basetypes.NewFloat64Value(-2.71828),
			basetypes.NewFloat32Value(1e6), // scientific notation
			basetypes.NewFloat64Value(-1.23e-4),
			basetypes.NewNumberValue(big.NewFloat(0)),
			basetypes.NewNumberValue(big.NewFloat(999.999)),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
			basetypes.NewBoolValue(true),
		},
	},
	"tuple": {
		customvalidator.AllowedSubtypes(basetypes.TupleType{
			ElemTypes: []attr.Type{basetypes.StringType{}, basetypes.Int64Type{}},
		}),
		[]attr.Value{
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.Int64Type{}},
				[]attr.Value{basetypes.NewStringValue(""), basetypes.NewInt64Value(0)},
			),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.Int32Type{}},
				[]attr.Value{basetypes.NewStringValue("test"), basetypes.NewInt32Value(42)},
			),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.Int64Type{}},
				[]attr.Value{basetypes.NewStringValue("negative"), basetypes.NewInt64Value(-123)},
			),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
			basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("")}),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
				[]attr.Value{basetypes.NewInt64Value(123), basetypes.NewStringValue("wrong order")},
			),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.Float64Type{}},
				[]attr.Value{basetypes.NewStringValue("pi"), basetypes.NewFloat64Value(3.14)},
			),
		},
	},
	"tuple_list": {
		customvalidator.AllowedSubtypes(basetypes.TupleType{
			ElemTypes: []attr.Type{basetypes.StringType{}, basetypes.StringType{}},
		}),
		[]attr.Value{
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.StringType{}},
				[]attr.Value{basetypes.NewStringValue(""), basetypes.NewStringValue("")},
			),
			basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue(""), basetypes.NewStringValue("")}),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
			basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("")}),
			basetypes.NewListValueMust(basetypes.Int64Type{}, []attr.Value{basetypes.NewInt64Value(0)}),
		},
	},
	"nested_tuple": {
		customvalidator.AllowedSubtypes(basetypes.TupleType{
			ElemTypes: []attr.Type{
				basetypes.TupleType{
					ElemTypes: []attr.Type{basetypes.Int64Type{}, basetypes.Float64Type{}},
				},
				basetypes.StringType{},
			},
		}),
		[]attr.Value{
			basetypes.NewTupleValueMust(
				[]attr.Type{
					basetypes.TupleType{
						ElemTypes: []attr.Type{basetypes.Int64Type{}, basetypes.Float64Type{}},
					},
					basetypes.StringType{},
				},
				[]attr.Value{
					basetypes.NewTupleValueMust(
						[]attr.Type{basetypes.Int64Type{}, basetypes.Float64Type{}},
						[]attr.Value{basetypes.NewInt64Value(42), basetypes.NewFloat64Value(3.14)},
					),
					basetypes.NewStringValue("nested"),
				},
			),
			basetypes.NewTupleValueMust(
				[]attr.Type{
					basetypes.TupleType{
						ElemTypes: []attr.Type{basetypes.Int32Type{}, basetypes.Float32Type{}},
					},
					basetypes.StringType{},
				},
				[]attr.Value{
					basetypes.NewTupleValueMust(
						[]attr.Type{basetypes.Int32Type{}, basetypes.Float32Type{}},
						[]attr.Value{basetypes.NewInt32Value(-100), basetypes.NewFloat32Value(2.718)},
					),
					basetypes.NewStringValue("compatible types"),
				},
			),
		},
		[]attr.Value{
			basetypes.NewStringValue("not a tuple"),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.StringType{}},
				[]attr.Value{basetypes.NewStringValue("wrong"), basetypes.NewStringValue("structure")},
			),
		},
	},
	"mixed_numeric_tuple": {
		customvalidator.AllowedSubtypes(basetypes.TupleType{
			ElemTypes: []attr.Type{basetypes.Int64Type{}, basetypes.Float64Type{}, basetypes.NumberType{}},
		}),
		[]attr.Value{
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.Int64Type{}, basetypes.Float64Type{}, basetypes.NumberType{}},
				[]attr.Value{
					basetypes.NewInt64Value(123),
					basetypes.NewFloat64Value(45.6),
					basetypes.NewNumberValue(big.NewFloat(789.012)),
				},
			),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.Int32Type{}, basetypes.Float32Type{}, basetypes.NumberType{}},
				[]attr.Value{
					basetypes.NewInt32Value(-456),
					basetypes.NewFloat32Value(-78.9),
					basetypes.NewNumberValue(big.NewFloat(-0.123)),
				},
			),
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.NumberType{}, basetypes.NumberType{}, basetypes.NumberType{}},
				[]attr.Value{
					basetypes.NewNumberValue(big.NewFloat(1)),
					basetypes.NewNumberValue(big.NewFloat(2.5)),
					basetypes.NewNumberValue(big.NewFloat(-3.14159)),
				},
			),
		},
		[]attr.Value{
			basetypes.NewTupleValueMust(
				[]attr.Type{basetypes.StringType{}, basetypes.Float64Type{}, basetypes.NumberType{}},
				[]attr.Value{
					basetypes.NewStringValue("not numeric"),
					basetypes.NewFloat64Value(1.0),
					basetypes.NewNumberValue(big.NewFloat(2)),
				},
			),
		},
	},
	"list": {
		customvalidator.AllowedSubtypes(basetypes.ListType{
			ElemType: basetypes.StringType{},
		}),
		[]attr.Value{
			basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("")}),
			basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{}),
			basetypes.NewTupleValueMust([]attr.Type{basetypes.StringType{}, basetypes.StringType{}}, []attr.Value{basetypes.NewStringValue(""), basetypes.NewStringValue("")}),
			basetypes.NewSetValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("")}),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
		},
	},
	"map": {
		customvalidator.AllowedSubtypes(basetypes.MapType{
			ElemType: basetypes.StringType{},
		}),
		[]attr.Value{
			basetypes.NewMapValueMust(basetypes.StringType{}, map[string]attr.Value{
				"": basetypes.NewStringValue(""),
			}),
			basetypes.NewMapValueMust(basetypes.StringType{}, map[string]attr.Value{}),
			basetypes.NewMapValueMust(basetypes.StringType{}, map[string]attr.Value{
				"key1": basetypes.NewStringValue("value1"),
				"key2": basetypes.NewStringValue("value2"),
				"key3": basetypes.NewStringValue(""),
			}),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"": basetypes.StringType{}},
				map[string]attr.Value{"": basetypes.NewStringValue("")},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"field1": basetypes.StringType{},
					"field2": basetypes.StringType{},
				},
				map[string]attr.Value{
					"field1": basetypes.NewStringValue("test"),
					"field2": basetypes.NewStringValue("data"),
				},
			),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
			basetypes.NewMapValueMust(basetypes.Int64Type{}, map[string]attr.Value{
				"key": basetypes.NewInt64Value(123),
			}),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"field": basetypes.Int64Type{}},
				map[string]attr.Value{"field": basetypes.NewInt64Value(456)},
			),
		},
	},
	"map_numeric": {
		customvalidator.AllowedSubtypes(basetypes.MapType{
			ElemType: basetypes.Int64Type{},
		}),
		[]attr.Value{
			basetypes.NewMapValueMust(basetypes.Int64Type{}, map[string]attr.Value{
				"count": basetypes.NewInt64Value(42),
				"total": basetypes.NewInt64Value(-100),
			}),
			basetypes.NewMapValueMust(basetypes.Int32Type{}, map[string]attr.Value{
				"items": basetypes.NewInt32Value(789),
			}),
			basetypes.NewMapValueMust(basetypes.Float64Type{}, map[string]attr.Value{
				"whole": basetypes.NewFloat64Value(25.0),
			}),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"score": basetypes.Int64Type{},
					"level": basetypes.Int32Type{},
				},
				map[string]attr.Value{
					"score": basetypes.NewInt64Value(9999),
					"level": basetypes.NewInt32Value(5),
				},
			),
		},
		[]attr.Value{
			basetypes.NewMapValueMust(basetypes.Float64Type{}, map[string]attr.Value{
				"decimal": basetypes.NewFloat64Value(3.14),
			}),
			basetypes.NewMapValueMust(basetypes.StringType{}, map[string]attr.Value{
				"text": basetypes.NewStringValue("not numeric"),
			}),
		},
	},
	"set": {
		customvalidator.AllowedSubtypes(basetypes.SetType{
			ElemType: basetypes.StringType{},
		}),
		[]attr.Value{
			basetypes.NewSetValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("")}),
			basetypes.NewSetValueMust(basetypes.StringType{}, []attr.Value{}),
			basetypes.NewSetValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("")}),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
		},
	},
	"object": {
		customvalidator.AllowedSubtypes(basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"1": basetypes.StringType{},
				"2": basetypes.Int64Type{},
			},
		}),
		[]attr.Value{
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"1": basetypes.StringType{}, "2": basetypes.Int64Type{}},
				map[string]attr.Value{"1": basetypes.NewStringValue(""), "2": basetypes.NewInt64Value(0)},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"1": basetypes.StringType{}, "2": basetypes.Int32Type{}},
				map[string]attr.Value{"1": basetypes.NewStringValue("test"), "2": basetypes.NewInt32Value(42)},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"1": basetypes.StringType{}, "2": basetypes.Int64Type{}},
				map[string]attr.Value{"1": basetypes.NewStringValue("negative"), "2": basetypes.NewInt64Value(-789)},
			),
		},
		[]attr.Value{
			basetypes.NewStringValue(""),
			basetypes.NewMapValueMust(basetypes.StringType{}, map[string]attr.Value{
				"key": basetypes.NewStringValue(""),
			}),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"1": basetypes.StringType{}, "2": basetypes.Float64Type{}},
				map[string]attr.Value{"1": basetypes.NewStringValue("float"), "2": basetypes.NewFloat64Value(3.14)},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{"1": basetypes.Int64Type{}, "2": basetypes.Int64Type{}},
				map[string]attr.Value{"1": basetypes.NewInt64Value(1), "2": basetypes.NewInt64Value(2)},
			),
		},
	},
	"object_mixed_types": {
		customvalidator.AllowedSubtypes(basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    basetypes.StringType{},
				"count":   basetypes.Int64Type{},
				"enabled": basetypes.BoolType{},
				"score":   basetypes.Float64Type{},
			},
		}),
		[]attr.Value{
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"name":    basetypes.StringType{},
					"count":   basetypes.Int64Type{},
					"enabled": basetypes.BoolType{},
					"score":   basetypes.Float64Type{},
				},
				map[string]attr.Value{
					"name":    basetypes.NewStringValue("example"),
					"count":   basetypes.NewInt64Value(100),
					"enabled": basetypes.NewBoolValue(true),
					"score":   basetypes.NewFloat64Value(85.5),
				},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"name":    basetypes.StringType{},
					"count":   basetypes.Int32Type{},
					"enabled": basetypes.BoolType{},
					"score":   basetypes.Float32Type{},
				},
				map[string]attr.Value{
					"name":    basetypes.NewStringValue("test"),
					"count":   basetypes.NewInt32Value(50),
					"enabled": basetypes.NewBoolValue(false),
					"score":   basetypes.NewFloat32Value(72.25),
				},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"name":    basetypes.StringType{},
					"count":   basetypes.NumberType{},
					"enabled": basetypes.BoolType{},
					"score":   basetypes.NumberType{},
				},
				map[string]attr.Value{
					"name":    basetypes.NewStringValue("numbers"),
					"count":   basetypes.NewNumberValue(big.NewFloat(200)),
					"enabled": basetypes.NewBoolValue(true),
					"score":   basetypes.NewNumberValue(big.NewFloat(95.75)),
				},
			),
		},
		[]attr.Value{
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"name":    basetypes.Int64Type{}, // wrong type
					"count":   basetypes.Int64Type{},
					"enabled": basetypes.BoolType{},
					"score":   basetypes.Float64Type{},
				},
				map[string]attr.Value{
					"name":    basetypes.NewInt64Value(123),
					"count":   basetypes.NewInt64Value(100),
					"enabled": basetypes.NewBoolValue(true),
					"score":   basetypes.NewFloat64Value(85.5),
				},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"name":  basetypes.StringType{},
					"count": basetypes.Int64Type{},
					// missing required fields
				},
				map[string]attr.Value{
					"name":  basetypes.NewStringValue("incomplete"),
					"count": basetypes.NewInt64Value(10),
				},
			),
		},
	},
	"object_nested": {
		customvalidator.AllowedSubtypes(basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"meta": basetypes.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":   basetypes.Int64Type{},
						"name": basetypes.StringType{},
					},
				},
				"data": basetypes.ListType{ElemType: basetypes.StringType{}},
			},
		}),
		[]attr.Value{
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"meta": basetypes.ObjectType{
						AttrTypes: map[string]attr.Type{
							"id":   basetypes.Int64Type{},
							"name": basetypes.StringType{},
						},
					},
					"data": basetypes.ListType{ElemType: basetypes.StringType{}},
				},
				map[string]attr.Value{
					"meta": basetypes.NewObjectValueMust(
						map[string]attr.Type{
							"id":   basetypes.Int64Type{},
							"name": basetypes.StringType{},
						},
						map[string]attr.Value{
							"id":   basetypes.NewInt64Value(123),
							"name": basetypes.NewStringValue("nested"),
						},
					),
					"data": basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{
						basetypes.NewStringValue("item1"),
						basetypes.NewStringValue("item2"),
					}),
				},
			),
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"meta": basetypes.ObjectType{
						AttrTypes: map[string]attr.Type{
							"id":   basetypes.Int32Type{}, // compatible numeric type
							"name": basetypes.StringType{},
						},
					},
					"data": basetypes.SetType{ElemType: basetypes.StringType{}}, // compatible collection type
				},
				map[string]attr.Value{
					"meta": basetypes.NewObjectValueMust(
						map[string]attr.Type{
							"id":   basetypes.Int32Type{},
							"name": basetypes.StringType{},
						},
						map[string]attr.Value{
							"id":   basetypes.NewInt32Value(456),
							"name": basetypes.NewStringValue("compatible"),
						},
					),
					"data": basetypes.NewSetValueMust(basetypes.StringType{}, []attr.Value{
						basetypes.NewStringValue("set1"),
						basetypes.NewStringValue("set2"),
					}),
				},
			),
		},
		[]attr.Value{
			basetypes.NewObjectValueMust(
				map[string]attr.Type{
					"meta": basetypes.StringType{}, // wrong nested type
					"data": basetypes.ListType{ElemType: basetypes.StringType{}},
				},
				map[string]attr.Value{
					"meta": basetypes.NewStringValue("should be object"),
					"data": basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{}),
				},
			),
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
