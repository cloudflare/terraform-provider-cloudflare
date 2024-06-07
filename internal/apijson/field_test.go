package apijson

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Struct struct {
	A types.String `json:"a"`
	B types.Int64  `json:"b"`
	C types.Bool   `json:"c"`
}

type FieldStruct struct {
	A types.String `json:"a"`
	B types.Int64  `json:"b"`
	C *Struct      `json:"c"`
	F types.Int64  `json:"f"`
}

func TestFieldMarshal(t *testing.T) {
	tests := map[string]struct {
		value    interface{}
		expected string
	}{
		"null_string":  {types.StringNull(), ""},
		"null_int64":   {types.Int64Null(), ""},
		"null_float64": {types.Float64Null(), ""},
		"null_bool":    {types.BoolNull(), ""},
		"null_struct":  {nil, ""},

		"string": {types.StringValue("mystring"), `"mystring"`},
		"int64":  {types.Int64Value(123), "123"},
		"bool":   {types.BoolValue(true), "true"},
		"struct": {Struct{A: types.StringValue("yo"), B: types.Int64Value(123), C: types.BoolValue(true)}, `{"a":"yo","b":123,"c":true}`},

		"param_struct": {
			FieldStruct{
				A: types.StringValue("hello"),
				B: types.Int64Value(12),
				C: &Struct{A: types.StringValue("nice"), C: types.BoolValue(false)},
			},
			`{"a":"hello","b":12,"c":{"a":"nice","c":false}}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := Marshal(test.value)
			if err != nil {
				t.Fatalf("didn't expect error %v", err)
			}
			if string(b) != test.expected {
				t.Fatalf("expected %s, received %s", test.expected, string(b))
			}
		})
	}
}
