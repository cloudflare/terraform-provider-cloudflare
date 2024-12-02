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

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

func P[T any](v T) *T { return &v }

type Primitives struct {
	A bool    `form:"a"`
	B int     `form:"b"`
	C uint    `form:"c"`
	D float64 `form:"d"`
	E float32 `form:"e"`
	F []int   `form:"f"`
}

type PrimitivePointers struct {
	A *bool    `form:"a"`
	B *int     `form:"b"`
	C *uint    `form:"c"`
	D *float64 `form:"d"`
	E *float32 `form:"e"`
	F *[]int   `form:"f"`
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
	Slice []Primitives `form:"slices"`
}

type DateTime struct {
	Date     time.Time `form:"date" format:"date"`
	DateTime time.Time `form:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                   `form:"a"`
	Extras map[string]interface{} `form:"-,extras"`
}

type TypedAdditionalProperties struct {
	A      bool           `form:"a"`
	Extras map[string]int `form:"-,extras"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int                   `form:"number2"`
	Extras map[string]interface{} `form:"-,extras"`
}

type Recursive struct {
	Name  string     `form:"name"`
	Child *Recursive `form:"child"`
}

type UnknownStruct struct {
	Unknown interface{} `form:"unknown"`
}

type UnionStruct struct {
	Union Union `form:"union" format:"date"`
}

type Union interface {
	union()
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `form:"type"`
	A    string `form:"a"`
	B    string `form:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `form:"type"`
	A    string `form:"a"`
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
Content-Disposition: form-data; name="f.0"

1
--xxx
Content-Disposition: form-data; name="f.1"

2
--xxx
Content-Disposition: form-data; name="f.2"

3
--xxx
Content-Disposition: form-data; name="f.3"

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
Content-Disposition: form-data; name="f.nested_a"

false
--xxx
Content-Disposition: form-data; name="g.hello"

world
--xxx
Content-Disposition: form-data; name="h.0"

a
--xxx
Content-Disposition: form-data; name="h.1"

b
--xxx
Content-Disposition: form-data; name="i.a"

3
--xxx
Content-Disposition: form-data; name="i.b"

8932
--xxx
Content-Disposition: form-data; name="j.0"

23.345
--xxx
Content-Disposition: form-data; name="j.1"

15
--xxx
Content-Disposition: form-data; name="k.dynamic_hello"

dynamic_world
--xxx
Content-Disposition: form-data; name="l.0"

a
--xxx
Content-Disposition: form-data; name="l.1"

b
--xxx
Content-Disposition: form-data; name="m.a"

3
--xxx
Content-Disposition: form-data; name="m.b"

8932
--xxx
Content-Disposition: form-data; name="n.0"

23.345
--xxx
Content-Disposition: form-data; name="n.1"

15
--xxx
Content-Disposition: form-data; name="o.0.nested_a"

false
--xxx
Content-Disposition: form-data; name="o.1.nested_a"

true
--xxx
Content-Disposition: form-data; name="p.a.nested_a"

false
--xxx
Content-Disposition: form-data; name="p.b.nested_a"

true
--xxx
Content-Disposition: form-data; name="q.0.nested_a"

false
--xxx
Content-Disposition: form-data; name="q.1.nested_a"

true
--xxx
Content-Disposition: form-data; name="r"

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
Content-Disposition: form-data; name="slices.0.a"

false
--xxx
Content-Disposition: form-data; name="slices.0.b"

237628372683
--xxx
Content-Disposition: form-data; name="slices.0.c"

654
--xxx
Content-Disposition: form-data; name="slices.0.d"

9999.43
--xxx
Content-Disposition: form-data; name="slices.0.e"

43.76
--xxx
Content-Disposition: form-data; name="slices.0.f.0"

1
--xxx
Content-Disposition: form-data; name="slices.0.f.1"

2
--xxx
Content-Disposition: form-data; name="slices.0.f.2"

3
--xxx
Content-Disposition: form-data; name="slices.0.f.3"

4
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
Content-Disposition: form-data; name="f.0"

1
--xxx
Content-Disposition: form-data; name="f.1"

2
--xxx
Content-Disposition: form-data; name="f.2"

3
--xxx
Content-Disposition: form-data; name="f.3"

4
--xxx
Content-Disposition: form-data; name="f.4"

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
Content-Disposition: form-data; name="child.name"

Alex
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
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.b"

bar
--xxx
Content-Disposition: form-data; name="union.type"

typeA
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
Content-Disposition: form-data; name="union.a"

foo
--xxx
Content-Disposition: form-data; name="union.type"

typeB
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
		if name != "terraform_types" {
			continue
		}
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			writer := multipart.NewWriter(buf)
			writer.SetBoundary("xxx")
			err := Marshal(test.val, writer)
			if err != nil {
				t.Errorf("serialization of %v\nfailed with error:\n%v", test.val, err)
			}
			err = writer.Close()
			if err != nil {
				t.Errorf("serialization of %v\nfailed with error:\n%v", test.val, err)
			}
			raw := buf.Bytes()
			if string(raw) != strings.ReplaceAll(test.buf, "\n", "\r\n") {
				t.Errorf("expected %+#v to serialize to '%s' but got '%s'", test.val, test.buf, string(raw))
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
