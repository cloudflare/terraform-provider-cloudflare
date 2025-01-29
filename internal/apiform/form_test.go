package apiform

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `json:"a"`
	B int     `json:"b"`
	C uint    `json:"c"`
	D float64 `json:"d"`
	E float32 `json:"e"`
	F []int   `json:"f"`
}

type PrimitivePointers struct {
	A *bool    `json:"a"`
	B *int     `json:"b"`
	C *uint    `json:"c"`
	D *float64 `json:"d"`
	E *float32 `json:"e"`
	F *[]int   `json:"f"`
}

type TerraformTypes struct {
	A types.Bool                                        `tfsdk:"a" json:"a"`
	B types.Int64                                       `tfsdk:"b" json:"b"`
	C types.Float64                                     `tfsdk:"c" json:"c"`
	D types.String                                      `tfsdk:"d" json:"d"`
	E timetypes.RFC3339                                 `tfsdk:"e" json:"e"`
	F customfield.NestedObject[NestedTerraformType]     `tfsdk:"f" json:"f"`
	G types.Object                                      `tfsdk:"g" json:"g"`
	H types.List                                        `tfsdk:"h" json:"h"`
	I types.Map                                         `tfsdk:"i" json:"i"`
	J types.Set                                         `tfsdk:"j" json:"j"`
	K types.Dynamic                                     `tfsdk:"k" json:"k"`
	L customfield.List[types.String]                    `tfsdk:"l" json:"l"`
	M customfield.Map[types.String]                     `tfsdk:"m" json:"m"`
	N customfield.Set[types.String]                     `tfsdk:"n" json:"n"`
	O customfield.NestedObjectList[NestedTerraformType] `tfsdk:"o" json:"o"`
	P customfield.NestedObjectMap[NestedTerraformType]  `tfsdk:"p" json:"p"`
	Q customfield.NestedObjectSet[NestedTerraformType]  `tfsdk:"q" json:"q"`
	R jsontypes.Normalized                              `tfsdk:"r" json:"r"`
}

type NestedTerraformType struct {
	NestedA types.Bool `tfsdk:"nested_a" json:"nested_a"`
}

type Slices struct {
	Slice []Primitives `json:"slices"`
}

type DateTime struct {
	Date     time.Time `json:"date" format:"date"`
	DateTime time.Time `json:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                   `json:"a"`
	Extras map[string]interface{} `json:"-,extras"`
}

type TypedAdditionalProperties struct {
	A      bool           `json:"a"`
	Extras map[string]int `json:"-,extras"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int                   `json:"number2"`
	Extras map[string]interface{} `json:"-,extras"`
}

type Recursive struct {
	Name  string     `json:"name"`
	Child *Recursive `json:"child"`
}

type UnknownStruct struct {
	Unknown interface{} `json:"unknown"`
}

type UnionStruct struct {
	Union Union `json:"union" format:"date"`
}

type Union interface {
	union()
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `json:"type"`
	A    string `json:"a"`
	B    string `json:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `json:"type"`
	A    string `json:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

type ReaderStruct struct {
}

var tests = map[string]struct {
	buf string
	val interface{}
}{
	"map_string": {
		`--xxx
Content-Disposition: form-data; name="foo"

bar
--xxx--
`,
		map[string]string{"foo": "bar"},
	},

	"map_interface": {
		`--xxx
Content-Disposition: form-data; name="a"

1
--xxx
Content-Disposition: form-data; name="b"

str
--xxx
Content-Disposition: form-data; name="c"

false
--xxx--
`,
		map[string]interface{}{"a": float64(1), "b": "str", "c": false},
	},

	"primitive_struct": {
		`--xxx
Content-Disposition: form-data; name="a"

false
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

9999.43
--xxx
Content-Disposition: form-data; name="e"

43.76
--xxx
Content-Disposition: form-data; name="f"

1
--xxx
Content-Disposition: form-data; name="f"

2
--xxx
Content-Disposition: form-data; name="f"

3
--xxx
Content-Disposition: form-data; name="f"

4
--xxx--
`,
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},

	"terraform_types": {
		`--xxx
Content-Disposition: form-data; name="a"

true
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

a string value
--xxx
Content-Disposition: form-data; name="e"

2006-01-02T15:04:05Z
--xxx
Content-Disposition: form-data; name="f"
Content-Type: application/json

{"nested_a":false}
--xxx
Content-Disposition: form-data; name="g"
Content-Type: application/json

{"hello":"world"}
--xxx
Content-Disposition: form-data; name="h"

a
--xxx
Content-Disposition: form-data; name="h"

b
--xxx
Content-Disposition: form-data; name="i"
Content-Type: application/json

{"a":3,"b":8932}
--xxx
Content-Disposition: form-data; name="j"

23.345
--xxx
Content-Disposition: form-data; name="j"

15
--xxx
Content-Disposition: form-data; name="k"
Content-Type: application/json

{"dynamic_hello":"dynamic_world"}
--xxx
Content-Disposition: form-data; name="l"

a
--xxx
Content-Disposition: form-data; name="l"

b
--xxx
Content-Disposition: form-data; name="m"
Content-Type: application/json

{"a":"3","b":"8932"}
--xxx
Content-Disposition: form-data; name="n"

23.345
--xxx
Content-Disposition: form-data; name="n"

15
--xxx
Content-Disposition: form-data; name="o"
Content-Type: application/json

{"nested_a":false}
--xxx
Content-Disposition: form-data; name="o"
Content-Type: application/json

{"nested_a":true}
--xxx
Content-Disposition: form-data; name="p"
Content-Type: application/json

{"a":{"nested_a":false},"b":{"nested_a":true}}
--xxx
Content-Disposition: form-data; name="q"
Content-Type: application/json

{"nested_a":false}
--xxx
Content-Disposition: form-data; name="q"
Content-Type: application/json

{"nested_a":true}
--xxx
Content-Disposition: form-data; name="r"
Content-Type: application/json

{"hello": "world"}
--xxx--
`,
		TerraformTypes{
			A: types.BoolValue(true),
			B: types.Int64Value(237628372683),
			C: types.Float64Value(654),
			D: types.StringValue("a string value"),
			E: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			F: customfield.NewObjectMust(context.TODO(), &NestedTerraformType{
				NestedA: types.BoolValue(false),
			}),
			G: types.ObjectValueMust(map[string]attr.Type{"hello": basetypes.StringType{}}, map[string]attr.Value{"hello": basetypes.NewStringValue("world")}),
			H: types.ListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("a"), basetypes.NewStringValue("b")}),
			I: types.MapValueMust(basetypes.Int64Type{}, map[string]attr.Value{"a": basetypes.NewInt64Value(3), "b": basetypes.NewInt64Value(8932)}),
			J: types.SetValueMust(basetypes.Float64Type{}, []attr.Value{basetypes.NewFloat64Value(23.345), basetypes.NewFloat64Value(15)}),
			K: types.DynamicValue(types.ObjectValueMust(map[string]attr.Type{"dynamic_hello": basetypes.StringType{}}, map[string]attr.Value{"dynamic_hello": basetypes.NewStringValue("dynamic_world")})),
			L: customfield.NewListMust[types.String](context.TODO(), []attr.Value{basetypes.NewStringValue("a"), basetypes.NewStringValue("b")}),
			M: customfield.NewMapMust[types.String](context.TODO(), map[string]types.String{"a": basetypes.NewStringValue("3"), "b": basetypes.NewStringValue("8932")}),
			N: customfield.NewSetMust[types.String](context.TODO(), []attr.Value{basetypes.NewStringValue("23.345"), basetypes.NewStringValue("15")}),
			O: customfield.NewObjectListMust(context.TODO(), []NestedTerraformType{
				{
					NestedA: types.BoolValue(false),
				},
				{
					NestedA: types.BoolValue(true),
				},
			}),
			P: customfield.NewObjectMapMust(context.TODO(), map[string]NestedTerraformType{
				"a": {
					NestedA: types.BoolValue(false),
				},
				"b": {
					NestedA: types.BoolValue(true),
				},
			}),
			Q: customfield.NewObjectSetMust(context.TODO(), []NestedTerraformType{
				{
					NestedA: types.BoolValue(false),
				},
				{
					NestedA: types.BoolValue(true),
				},
			}),
			R: jsontypes.NewNormalizedValue(`{"hello": "world"}`),
		},
	},

	"slices": {
		`--xxx
Content-Disposition: form-data; name="slices"
Content-Type: application/json

{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}
--xxx--
`,
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},

	"primitive_pointer_struct": {
		`--xxx
Content-Disposition: form-data; name="a"

false
--xxx
Content-Disposition: form-data; name="b"

237628372683
--xxx
Content-Disposition: form-data; name="c"

654
--xxx
Content-Disposition: form-data; name="d"

9999.43
--xxx
Content-Disposition: form-data; name="e"

43.76
--xxx
Content-Disposition: form-data; name="f"

1
--xxx
Content-Disposition: form-data; name="f"

2
--xxx
Content-Disposition: form-data; name="f"

3
--xxx
Content-Disposition: form-data; name="f"

4
--xxx
Content-Disposition: form-data; name="f"

5
--xxx--
`,
		PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},

	"datetime_struct": {
		`--xxx
Content-Disposition: form-data; name="date"

2006-01-02
--xxx
Content-Disposition: form-data; name="date-time"

2006-01-02T15:04:05Z
--xxx--
`,
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"additional_properties": {
		`--xxx
Content-Disposition: form-data; name="a"

true
--xxx
Content-Disposition: form-data; name="bar"

value
--xxx
Content-Disposition: form-data; name="foo"

true
--xxx--
`,
		AdditionalProperties{
			A: true,
			Extras: map[string]interface{}{
				"bar": "value",
				"foo": true,
			},
		},
	},

	"recursive_struct": {
		`--xxx
Content-Disposition: form-data; name="child"
Content-Type: application/json

{"name":"Alex"}
--xxx
Content-Disposition: form-data; name="name"

Robert
--xxx--
`,
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"unknown_struct_number": {
		`--xxx
Content-Disposition: form-data; name="unknown"

12
--xxx--
`,
		UnknownStruct{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		`--xxx
Content-Disposition: form-data; name="unknown.foo"

bar
--xxx--
`,
		UnknownStruct{
			Unknown: map[string]interface{}{
				"foo": "bar",
			},
		},
	},

	"union_integer": {
		`--xxx
Content-Disposition: form-data; name="union"

12
--xxx--
`,
		UnionStruct{
			Union: UnionInteger(12),
		},
	},

	"union_struct_discriminated_a": {
		`--xxx
Content-Disposition: form-data; name="union"
Content-Type: application/json

{"a":"foo","b":"bar","type":"typeA"}
--xxx--
`,

		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"union_struct_discriminated_b": {
		`--xxx
Content-Disposition: form-data; name="union"
Content-Type: application/json

{"a":"foo","type":"typeB"}
--xxx--
`,
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		`--xxx
Content-Disposition: form-data; name="union"

2010-05-23
--xxx--
`,
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},
}

func TestEncode(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			writer := multipart.NewWriter(buf)
			writer.SetBoundary("xxx")
			err := MarshalRoot(test.val, writer)
			if err != nil {
				t.Errorf("serialization of %v\nfailed with error:\n%v", test.val, err)
			}
			err = writer.Close()
			if err != nil {
				t.Errorf("serialization of %v\nfailed with error:\n%v", test.val, err)
			}
			raw := buf.Bytes()
			if string(raw) != strings.ReplaceAll(test.buf, "\n", "\r\n") {
				t.Errorf("serialization did not match: %+#v\n\n#### EXPECTED\n%s\n\n#### ACTUAL\n\n%s", test.val, test.buf, string(raw))
			}
		})
	}
}

func DropDiagnostic[resType interface{}](res resType, diags diag.Diagnostics) resType {
	for _, d := range diags {
		panic(fmt.Sprintf("%s: %s", d.Summary(), d.Detail()))
	}
	return res
}
