package apijson

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/gjson"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

func P[T any](v T) *T { return &v }

type TfsdkStructs struct {
	BoolValue        types.Bool                                        `tfsdk:"tfsdk_bool_value" json:"bool_value"`
	StringValue      types.String                                      `tfsdk:"tfsdk_string_value" json:"string_value"`
	Data             *EmbeddedTfsdkStruct                              `tfsdk:"tfsdk_data" json:"data"`
	DataObject       customfield.NestedObject[EmbeddedTfsdkStruct]     `tfsdk:"tfsdk_data_object" json:"data_object"`
	ListObject       customfield.List[types.String]                    `tfsdk:"tfsdk_list_object" json:"list_object"`
	NestedObjectList customfield.NestedObjectList[EmbeddedTfsdkStruct] `tfsdk:"tfsdk_nested_object_list" json:"nested_object_list"`
	SetObject        customfield.Set[types.String]                     `tfsdk:"tfsdk_set_object" json:"set_object"`
	NestedObjectSet  customfield.NestedObjectSet[EmbeddedTfsdkStruct]  `tfsdk:"tfsdk_nested_object_set" json:"nested_object_set"`
	MapObject        customfield.Map[types.String]                     `tfsdk:"tfsdk_map_object" json:"map_object"`
	NestedObjectMap  customfield.NestedObjectMap[EmbeddedTfsdkStruct]  `tfsdk:"tfsdk_nested_object_map" json:"nested_object_map"`
	FloatValue       types.Float64                                     `tfsdk:"tfsdk_float_value" json:"float_value"`
	OptionalArray    *[]types.String                                   `tfsdk:"tfsdk_optional_array" json:"optional_array"`
}

type EmbeddedTfsdkStruct struct {
	EmbeddedString types.String                                 `tfsdk:"tfsdk_embedded_string" json:"embedded_string,required"`
	EmbeddedInt    types.Int64                                  `tfsdk:"tfsdk_embedded_int" json:"embedded_int,optional"`
	DataObject     customfield.NestedObject[DoubleNestedStruct] `tfsdk:"tfsdk_data_object" json:"data_object,optional"`
}

type DoubleNestedStruct struct {
	NestedInt types.Int64 `tfsdk:"tfsdk_nested_int" json:"nested_int"`
}

type DoubleNestedStructZero struct {
	NestedInt types.Int64 `tfsdk:"tfsdk_nested_int" json:"nested_int,decode_null_to_zero"`
}

type TfsdkStructsZero struct {
	BoolValue        types.Bool                                            `tfsdk:"tfsdk_bool_value" json:"bool_value,decode_null_to_zero"`
	StringValue      types.String                                          `tfsdk:"tfsdk_string_value" json:"string_value,decode_null_to_zero"`
	Data             *EmbeddedTfsdkStructZero                              `tfsdk:"tfsdk_data" json:"data,decode_null_to_zero"`
	DataObject       customfield.NestedObject[EmbeddedTfsdkStructZero]     `tfsdk:"tfsdk_data_object" json:"data_object,decode_null_to_zero"`
	ListObject       customfield.List[types.String]                        `tfsdk:"tfsdk_list_object" json:"list_object,decode_null_to_zero"`
	NestedObjectList customfield.NestedObjectList[EmbeddedTfsdkStructZero] `tfsdk:"tfsdk_nested_object_list" json:"nested_object_list,decode_null_to_zero"`
	SetObject        customfield.Set[types.String]                         `tfsdk:"tfsdk_set_object" json:"set_object,decode_null_to_zero"`
	NestedObjectSet  customfield.NestedObjectSet[EmbeddedTfsdkStructZero]  `tfsdk:"tfsdk_nested_object_set" json:"nested_object_set,decode_null_to_zero"`
	MapObject        customfield.Map[types.String]                         `tfsdk:"tfsdk_map_object" json:"map_object,decode_null_to_zero"`
	NestedObjectMap  customfield.NestedObjectMap[EmbeddedTfsdkStructZero]  `tfsdk:"tfsdk_nested_object_map" json:"nested_object_map,decode_null_to_zero"`
	FloatValue       types.Float64                                         `tfsdk:"tfsdk_float_value" json:"float_value,decode_null_to_zero"`
	OptionalArray    *[]types.String                                       `tfsdk:"tfsdk_optional_array" json:"optional_array,decode_null_to_zero"`
}

type EmbeddedTfsdkStructZero struct {
	EmbeddedString types.String                                 `tfsdk:"tfsdk_embedded_string" json:"embedded_string,required,decode_null_to_zero"`
	EmbeddedInt    types.Int64                                  `tfsdk:"tfsdk_embedded_int" json:"embedded_int,optional,decode_null_to_zero"`
	DataObject     customfield.NestedObject[DoubleNestedStruct] `tfsdk:"tfsdk_data_object" json:"data_object,optional,decode_null_to_zero"`
}

type Primitives struct {
	A bool    `json:"a"`
	B int     `json:"b"`
	C uint    `json:"c"`
	D float64 `json:"d"`
	E float32 `json:"e"`
	F []int   `json:"f"`
}

type PrimitivesZero struct {
	A bool    `json:"a,decode_null_to_zero"`
	B int     `json:"b,decode_null_to_zero"`
	C uint    `json:"c,decode_null_to_zero"`
	D float64 `json:"d,decode_null_to_zero"`
	E float32 `json:"e,decode_null_to_zero"`
	F []int   `json:"f,decode_null_to_zero"`
}

type PrimitivePointers struct {
	A *bool    `json:"a"`
	B *int     `json:"b"`
	C *uint    `json:"c"`
	D *float64 `json:"d"`
	E *float32 `json:"e"`
	F *[]int   `json:"f"`
}

type PrimitivePointersZero struct {
	A *bool    `json:"a,decode_null_to_zero"`
	B *int     `json:"b,decode_null_to_zero"`
	C *uint    `json:"c,decode_null_to_zero"`
	D *float64 `json:"d,decode_null_to_zero"`
	E *float32 `json:"e,decode_null_to_zero"`
	F *[]int   `json:"f,decode_null_to_zero"`
}

type Slices struct {
	Slice []Primitives `json:"slices"`
}

type SlicesZero struct {
	Slice []PrimitivesZero `json:"slices,decode_null_to_zero"`
}

type DateTime struct {
	Date     time.Time `json:"date" format:"date"`
	DateTime time.Time `json:"date-time" format:"date-time"`
}

type DateTimeZero struct {
	Date     time.Time `json:"date" format:"date,decode_null_to_zero"`
	DateTime time.Time `json:"date-time" format:"date-time,decode_null_to_zero"`
}

type DateTimeCustom struct {
	DateCustom     timetypes.RFC3339 `json:"date" format:"date"`
	DateTimeCustom timetypes.RFC3339 `json:"date-time" format:"date-time"`
}

type DateTimeCustomZero struct {
	DateCustom     timetypes.RFC3339 `json:"date,decode_null_to_zero" format:"date"`
	DateTimeCustom timetypes.RFC3339 `json:"date-time,decode_null_to_zero" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                   `json:"a"`
	Extras map[string]interface{} `json:"-,extras"`
}

type AdditionalPropertiesZero struct {
	A      bool                   `json:"a,decode_null_to_zero"`
	Extras map[string]interface{} `json:"-,extras,decode_null_to_zero"`
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

type RecursiveZero struct {
	Name  string     `json:"name,decode_null_to_zero"`
	Child *Recursive `json:"child,decode_null_to_zero"`
}

type UnknownStruct struct {
	Unknown interface{} `json:"unknown"`
}

type UnknownStructZero struct {
	Unknown interface{} `json:"unknown,decode_null_to_zero"`
}

type UnionStruct struct {
	Union Union `json:"union" format:"date"`
}

type UnionStructZero struct {
	Union Union `json:"union,decode_null_to_zero" format:"date"`
}

type Union interface {
	union()
}

type Inline struct {
	InlineField Primitives `json:"-,inline"`
}

type InlineArray struct {
	InlineField []string `json:"-,inline"`
}

type EncodeStateForUnknownStruct struct {
	NormalField types.String `tfsdk:"normal_field" json:"normal_field"`
	// force_encode flag: don't skip this field even though it's computed
	ComputedWithForceEncode types.String `tfsdk:"computed_force_encode" json:"computed_force_encode,computed,force_encode"`
	// force_encode+encode_state_for_unknown: don't skip this field even though it's computed,
	// AND encode value from state if value from plan is unknown
	ComputedWithStateEncode types.String `tfsdk:"computed_state_encode" json:"computed_state_encode,computed,force_encode,encode_state_for_unknown"`
	// encode_state_for_unknown: encode value from state if value from plan is unknown
	ComputedOptionalWithStateEncode types.String `tfsdk:"computed_optional_state_encode" json:"computed_optional_state_encode,computed_optional,encode_state_for_unknown"`
	ComputedRegular                 types.String `tfsdk:"computed_regular" json:"computed_regular,computed"`
	ComputedOptionalRegular         types.String `tfsdk:"computed_optional_regular" json:"computed_optional_regular,computed_optional"`
}

func init() {
	RegisterUnion(reflect.TypeOf((*Union)(nil)).Elem(), "type",
		UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionTime{}),
		},
		UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionInteger(0)),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeA",
			Type:               reflect.TypeOf(UnionStructA{}),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeB",
			Type:               reflect.TypeOf(UnionStructB{}),
		},
	)
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

type ResultEnvelope struct {
	Result RecordsModel `json:"result"`
}

type RecordsModel struct {
	A types.String `tfsdk:"tfsdk_a" json:"a"`
	B types.String `tfsdk:"tfsdk_b" json:"b"`
	C types.String `tfsdk:"tfsdk_c" json:"c,computed"`
}

type ResultEnvelopeZero struct {
	Result RecordsModelZero `json:"result"`
}

type RecordsModelZero struct {
	A types.String `tfsdk:"tfsdk_a" json:"a,decode_null_to_zero"`
	B types.String `tfsdk:"tfsdk_b" json:"b,decode_null_to_zero"`
	C types.String `tfsdk:"tfsdk_c" json:"c,computed,decode_null_to_zero"`
}

func DropDiagnostic[resType interface{}](res resType, diags diag.Diagnostics) resType {
	for _, d := range diags {
		panic(fmt.Sprintf("%s: %s", d.Summary(), d.Detail()))
	}
	return res
}

type JsonModel struct {
	Arr  jsontypes.Normalized            `tfsdk:"tfsdk_arr" json:"arr"`
	Bol  jsontypes.Normalized            `tfsdk:"tfsdk_bol" json:"bol"`
	Map  jsontypes.Normalized            `tfsdk:"tfsdk_map" json:"map"`
	Nil  jsontypes.Normalized            `tfsdk:"tfsdk_nil" json:"nil"`
	Num  jsontypes.Normalized            `tfsdk:"tfsdk_num" json:"num"`
	Str  jsontypes.Normalized            `tfsdk:"tfsdk_str" json:"str"`
	Arr2 []jsontypes.Normalized          `tfsdk:"tfsdk_arr2" json:"arr2"`
	Map2 map[string]jsontypes.Normalized `tfsdk:"tfsdk_map2" json:"map2"`
}

type JsonModelZero struct {
	Arr  jsontypes.Normalized            `tfsdk:"tfsdk_arr" json:"arr,decode_null_to_zero"`
	Bol  jsontypes.Normalized            `tfsdk:"tfsdk_bol" json:"bol,decode_null_to_zero"`
	Map  jsontypes.Normalized            `tfsdk:"tfsdk_map" json:"map,decode_null_to_zero"`
	Nil  jsontypes.Normalized            `tfsdk:"tfsdk_nil" json:"nil,decode_null_to_zero"`
	Num  jsontypes.Normalized            `tfsdk:"tfsdk_num" json:"num,decode_null_to_zero"`
	Str  jsontypes.Normalized            `tfsdk:"tfsdk_str" json:"str,decode_null_to_zero"`
	Arr2 []jsontypes.Normalized          `tfsdk:"tfsdk_arr2" json:"arr2,decode_null_to_zero"`
	Map2 map[string]jsontypes.Normalized `tfsdk:"tfsdk_map2" json:"map2,decode_null_to_zero"`
}

func time2time(t time.Time) timetypes.RFC3339 {
	return timetypes.NewRFC3339TimePointerValue(&t)
}

var ctx = context.TODO()

var tests = map[string]struct {
	buf string
	val interface{}
}{
	"true":               {"true", true},
	"false":              {"false", false},
	"int":                {"1", 1},
	"int_bigger":         {"12324", 12324},
	"int_string_coerce":  {`"65"`, 65},
	"int_boolean_coerce": {"true", 1},
	"int64":              {"1", int64(1)},
	"int64_huge":         {"123456789123456789", int64(123456789123456789)},
	"uint":               {"1", uint(1)},
	"uint_bigger":        {"12324", uint(12324)},
	"uint_coerce":        {`"65"`, uint(65)},
	"float_1.54":         {"1.54", float32(1.54)},
	"float_1.89":         {"1.89", float64(1.89)},
	"string":             {`"str"`, "str"},
	"string_int_coerce":  {`12`, "12"},
	"array_string":       {`["foo","bar"]`, []string{"foo", "bar"}},
	"array_int":          {`[1,2]`, []int{1, 2}},
	"array_int_coerce":   {`["1",2]`, []int{1, 2}},

	"ptr_true":               {"true", P(true)},
	"ptr_false":              {"false", P(false)},
	"ptr_int":                {"1", P(1)},
	"ptr_int_bigger":         {"12324", P(12324)},
	"ptr_int_string_coerce":  {`"65"`, P(65)},
	"ptr_int_boolean_coerce": {"true", P(1)},
	"ptr_int64":              {"1", P(int64(1))},
	"ptr_int64_huge":         {"123456789123456789", P(int64(123456789123456789))},
	"ptr_uint":               {"1", P(uint(1))},
	"ptr_uint_bigger":        {"12324", P(uint(12324))},
	"ptr_uint_coerce":        {`"65"`, P(uint(65))},
	"ptr_float_1.54":         {"1.54", P(float32(1.54))},
	"ptr_float_1.89":         {"1.89", P(float64(1.89))},

	"date_time":             {`"2007-03-01T13:00:00Z"`, time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC)},
	"date_time_nano_coerce": {`"2007-03-01T13:03:05.123456789Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC)},

	"date_time_missing_t_coerce":              {`"2007-03-01 13:03:05Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)},
	"date_time_missing_timezone_coerce":       {`"2007-03-01T13:03:05"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)},
	"date_time_missing_timezone_colon_coerce": {`"2007-03-01T13:03:05+0100"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", 60*60))},
	"date_time_nano_missing_t_coerce":         {`"2007-03-01 13:03:05.123456789Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC)},

	"date_time_custom":             {`"2007-03-01T13:00:00Z"`, time2time(time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC))},
	"date_time_custom_nano_coerce": {`"2007-03-01T13:03:05.123456789Z"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC))},

	"date_time_custom_missing_t_coerce":              {`"2007-03-01 13:03:05Z"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC))},
	"date_time_custom_missing_timezone_coerce":       {`"2007-03-01T13:03:05"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC))},
	"date_time_custom_missing_timezone_colon_coerce": {`"2007-03-01T13:03:05+0100"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", 60*60)))},
	"date_time_custom_nano_missing_t_coerce":         {`"2007-03-01 13:03:05.123456789Z"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC))},

	"map_string":                       {`{"foo":"bar"}`, map[string]string{"foo": "bar"}},
	"map_string_with_sjson_path_chars": {`{":a.b.c*:d*-1e.f":"bar"}`, map[string]string{":a.b.c*:d*-1e.f": "bar"}},
	"map_interface":                    {`{"a":1,"b":"str","c":false}`, map[string]interface{}{"a": float64(1), "b": "str", "c": false}},

	"primitive_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`,
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},
	"primitive_struct_zero": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`,
		PrimitivesZero{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},

	"slices": {
		`{"slices":[{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}]}`,
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},
	"slices_zero": {
		`{"slices":[{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}]}`,
		SlicesZero{
			Slice: []PrimitivesZero{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},

	"primitive_pointer_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4,5]}`,
		PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},
	"primitive_pointer_struct_zero": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4,5]}`,
		PrimitivePointersZero{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},

	"datetime_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},
	"datetime_struct_zero": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTimeZero{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"datetime_custom_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: time2time(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
		},
	},
	"datetime_custom_struct_zero": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTimeCustomZero{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: time2time(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
		},
	},

	"additional_properties": {
		`{"a":true,"bar":"value","foo":true}`,
		AdditionalProperties{
			A: true,
			Extras: map[string]interface{}{
				"bar": "value",
				"foo": true,
			},
		},
	},
	"additional_properties_zero": {
		`{"a":true,"bar":"value","foo":true}`,
		AdditionalPropertiesZero{
			A: true,
			Extras: map[string]interface{}{
				"bar": "value",
				"foo": true,
			},
		},
	},

	"recursive_struct": {
		`{"child":{"name":"Alex"},"name":"Robert"}`,
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},
	"recursive_struct_zero": {
		`{"child":{"name":"Alex"},"name":"Robert"}`,
		RecursiveZero{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"unknown_struct_number": {
		`{"unknown":12}`,
		UnknownStruct{
			Unknown: 12.,
		},
	},
	"unknown_struct_number_zero": {
		`{"unknown":12}`,
		UnknownStructZero{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		`{"unknown":{"foo":"bar"}}`,
		UnknownStruct{
			Unknown: map[string]interface{}{
				"foo": "bar",
			},
		},
	},
	"unknown_struct_map_zero": {
		`{"unknown":{"foo":"bar"}}`,
		UnknownStructZero{
			Unknown: map[string]interface{}{
				"foo": "bar",
			},
		},
	},

	"union_integer": {
		`{"union":12}`,
		UnionStruct{
			Union: UnionInteger(12),
		},
	},
	"union_integer_zero": {
		`{"union":12}`,
		UnionStructZero{
			Union: UnionInteger(12),
		},
	},

	"union_struct_discriminated_a": {
		`{"union":{"a":"foo","b":"bar","type":"typeA"}}`,
		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},
	"union_struct_discriminated_a_zero": {
		`{"union":{"a":"foo","b":"bar","type":"typeA"}}`,
		UnionStructZero{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"union_struct_discriminated_b": {
		`{"union":{"a":"foo","type":"typeB"}}`,
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},
	"union_struct_discriminated_b_zero": {
		`{"union":{"a":"foo","type":"typeB"}}`,
		UnionStructZero{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		`{"union":"2010-05-23"}`,
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},
	"union_struct_time_zero": {
		`{"union":"2010-05-23"}`,
		UnionStructZero{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},

	"tfsdk_null_string":  {"", types.StringNull()},
	"tfsdk_null_int":     {"", types.Int64Null()},
	"tfsdk_null_float":   {"", types.Float64Null()},
	"tfsdk_null_bool":    {"", types.BoolNull()},
	"tfsdk_null_dynamic": {"", types.DynamicNull()},

	"tfsdk_string":             {`"hey"`, types.StringValue("hey")},
	"tfsdk_true":               {"true", types.BoolValue(true)},
	"tfsdk_false":              {"false", types.BoolValue(false)},
	"tfsdk_int":                {"1", types.Int64Value(1)},
	"tfsdk_int_bigger":         {"12324", types.Int64Value(12324)},
	"tfsdk_int_string_coerce":  {`"65"`, types.Int64Value(65)},
	"tfsdk_int_boolean_coerce": {"true", types.BoolValue(true)},
	"tfsdk_float_1.54":         {"1.54", types.Float64Value(1.54)},
	"tfsdk_float_1.89":         {"1.89", types.Float64Value(1.89)},
	"tfsdk_array_ptr":          {"[\"hi\",null]", &[]types.String{types.StringValue("hi"), types.StringNull()}},
	"tfsdk_dynamic_string":     {`"hey"`, types.DynamicValue(types.StringValue("hey"))},
	"tfsdk_dynamic_int":        {"5", types.DynamicValue(types.Int64Value(5))},

	"tfsdk_list": {
		"[1,2,3]",
		types.ListValueMust(
			basetypes.Int64Type{},
			[]attr.Value{basetypes.NewInt64Value(1), basetypes.NewInt64Value(2), basetypes.NewInt64Value(3)},
		),
	},

	"tfsdk_object": {
		`{"baz":4,"foo":"bar"}`,
		types.ObjectValueMust(
			map[string]attr.Type{"baz": basetypes.Int64Type{}, "foo": basetypes.StringType{}},
			map[string]attr.Value{"baz": basetypes.NewInt64Value(4), "foo": basetypes.NewStringValue("bar")},
		),
	},

	"tfsdk_dynamic_object": {
		`{"baz":4,"foo":"bar"}`,
		types.DynamicValue(
			types.ObjectValueMust(
				map[string]attr.Type{"baz": basetypes.Int64Type{}, "foo": basetypes.StringType{}},
				map[string]attr.Value{"baz": basetypes.NewInt64Value(4), "foo": basetypes.NewStringValue("bar")},
			),
		),
	},

	"embedded_tfsdk_struct": {
		`{"bool_value":true,` +
			`"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},` +
			`"data_object":{"data_object":{"nested_int":19},"embedded_int":18,"embedded_string":"embedded_data_string_value"},` +
			`"float_value":3.14,` +
			`"list_object":["hi_list","there_list"],` +
			`"map_object":{"hi_map":"there_map"},` +
			`"nested_object_list":[{"embedded_int":20,"embedded_string":"nested_object_string"}],` +
			`"nested_object_map":{"nested_object_map_key":{"embedded_int":21,"embedded_string":"nested_object_string_in_map"}},` +
			`"nested_object_set":[{"embedded_int":21,"embedded_string":"nested_object_string_in_set"}],` +
			`"optional_array":["hi","there"],` +
			`"set_object":["hi_set","there_set"],` +
			`"string_value":"string_value"}`,
		TfsdkStructs{
			BoolValue:   types.BoolValue(true),
			StringValue: types.StringValue("string_value"),
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			DataObject: customfield.NewObjectMust(ctx, &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_data_string_value"),
				EmbeddedInt:    types.Int64Value(18),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(19),
				}),
			}),
			ListObject: customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("hi_list"), types.StringValue("there_list")}),
			NestedObjectList: customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{{
				EmbeddedString: types.StringValue("nested_object_string"),
				EmbeddedInt:    types.Int64Value(20),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}}),
			SetObject: customfield.NewSetMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("hi_set"), types.StringValue("there_set")}),
			NestedObjectSet: customfield.NewObjectSetMust(ctx, []EmbeddedTfsdkStruct{{
				EmbeddedString: types.StringValue("nested_object_string_in_set"),
				EmbeddedInt:    types.Int64Value(21),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}}),
			MapObject: customfield.NewMapMust[basetypes.StringValue](ctx, map[string]types.String{"hi_map": types.StringValue("there_map")}),
			NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStruct{"nested_object_map_key": {
				EmbeddedString: types.StringValue("nested_object_string_in_map"),
				EmbeddedInt:    types.Int64Value(21),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			}}),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
		},
	},
	"embedded_tfsdk_struct_zero": {
		`{"bool_value":true,` +
			`"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},` +
			`"data_object":{"data_object":{"nested_int":19},"embedded_int":18,"embedded_string":"embedded_data_string_value"},` +
			`"float_value":3.14,` +
			`"list_object":["hi_list","there_list"],` +
			`"map_object":{"hi_map":"there_map"},` +
			`"nested_object_list":[{"embedded_int":20,"embedded_string":"nested_object_string"}],` +
			`"nested_object_map":{"nested_object_map_key":{"embedded_int":21,"embedded_string":"nested_object_string_in_map"}},` +
			`"nested_object_set":[{"embedded_int":21,"embedded_string":"nested_object_string_in_set"}],` +
			`"optional_array":["hi","there"],` +
			`"set_object":["hi_set","there_set"],` +
			`"string_value":"string_value"}`,
		TfsdkStructsZero{
			BoolValue:   types.BoolValue(true),
			StringValue: types.StringValue("string_value"),
			Data: &EmbeddedTfsdkStructZero{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Null(),
				}),
			},
			DataObject: customfield.NewObjectMust(ctx, &EmbeddedTfsdkStructZero{
				EmbeddedString: types.StringValue("embedded_data_string_value"),
				EmbeddedInt:    types.Int64Value(18),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(19),
				}),
			}),
			ListObject: customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("hi_list"), types.StringValue("there_list")}),
			NestedObjectList: customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStructZero{{
				EmbeddedString: types.StringValue("nested_object_string"),
				EmbeddedInt:    types.Int64Value(20),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Null(),
				}),
			}}),
			SetObject: customfield.NewSetMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("hi_set"), types.StringValue("there_set")}),
			NestedObjectSet: customfield.NewObjectSetMust(ctx, []EmbeddedTfsdkStructZero{{
				EmbeddedString: types.StringValue("nested_object_string_in_set"),
				EmbeddedInt:    types.Int64Value(21),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Null(),
				}),
			}}),
			MapObject: customfield.NewMapMust[basetypes.StringValue](ctx, map[string]types.String{"hi_map": types.StringValue("there_map")}),
			NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStructZero{"nested_object_map_key": {
				EmbeddedString: types.StringValue("nested_object_string_in_map"),
				EmbeddedInt:    types.Int64Value(21),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Null(),
				}),
			}}),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
		},
	},

	"customfield_null_object": {
		"",
		customfield.NullObject[DoubleNestedStruct](ctx),
	},
	"customfield_null_object_zero": {
		"",
		customfield.NullObject[DoubleNestedStructZero](ctx),
	},

	"json_struct_nil1": {`{}`, JsonModel{}},
	"json_struct_nil1_zero": {`{}`, JsonModelZero{
		Arr:  jsontypes.NewNormalizedValue(""),
		Bol:  jsontypes.NewNormalizedValue(""),
		Map:  jsontypes.NewNormalizedValue(""),
		Nil:  jsontypes.NewNormalizedValue(""),
		Num:  jsontypes.NewNormalizedValue(""),
		Str:  jsontypes.NewNormalizedValue(""),
		Arr2: []jsontypes.Normalized{},
		Map2: map[string]jsontypes.Normalized{},
	}},
	"json_struct_nil2": {`{}`, JsonModel{}},
	"json_struct_nil2_zero": {`{}`, JsonModelZero{
		Arr:  jsontypes.NewNormalizedValue(""),
		Bol:  jsontypes.NewNormalizedValue(""),
		Map:  jsontypes.NewNormalizedValue(""),
		Nil:  jsontypes.NewNormalizedValue(""),
		Num:  jsontypes.NewNormalizedValue(""),
		Str:  jsontypes.NewNormalizedValue(""),
		Arr2: []jsontypes.Normalized{},
		Map2: map[string]jsontypes.Normalized{},
	}},
}

type ListWithNestedObj struct {
	A customfield.NestedObjectList[Embedded2] `tfsdk:"tfsdk_a" json:"a"`
}

type Embedded2 struct {
	B types.String      `tfsdk:"tfsdk_b" json:"b"`
	C *Inner            `tfsdk:"tfsdk_c" json:"c"`
	D *[]*Inner         `tfsdk:"tfsdk_d" json:"d"`
	E []string          `tfsdk:"tfsdk_e" json:"e"`
	F *map[string]Inner `tfsdk:"tfsdk_f" json:"f"`
}

type Inner struct {
	D types.String `tfsdk:"tfsdk_d" json:"d"`
}

type ListWithNestedObjZero struct {
	A customfield.NestedObjectList[Embedded2Zero] `tfsdk:"tfsdk_a" json:"a,decode_null_to_zero"`
}

type Embedded2Zero struct {
	B types.String          `tfsdk:"tfsdk_b" json:"b,decode_null_to_zero"`
	C *InnerZero            `tfsdk:"tfsdk_c" json:"c,decode_null_to_zero"`
	D *[]*InnerZero         `tfsdk:"tfsdk_d" json:"d,decode_null_to_zero"`
	E []string              `tfsdk:"tfsdk_e" json:"e,decode_null_to_zero"`
	F *map[string]InnerZero `tfsdk:"tfsdk_f" json:"f,decode_null_to_zero"`
}

type InnerZero struct {
	D types.String `tfsdk:"tfsdk_d" json:"d,decode_null_to_zero"`
}

var decode_only_tests = map[string]struct {
	buf string
	val interface{}
}{
	"tfsdk_struct_decode": {
		`{"result":{"c":"7887590e1967befa70f48ffe9f61ce80","a":"88281d6015751d6172e7313b0c665b5e","extra":"property","another":2,"b":"http://example.com/example.html\t20"}`,
		ResultEnvelope{RecordsModel{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}},
	},
	"tfsdk_struct_decode_zero": {
		`{"result":{"c":"7887590e1967befa70f48ffe9f61ce80","a":"88281d6015751d6172e7313b0c665b5e","extra":"property","another":2,"b":"http://example.com/example.html\t20"}`,
		ResultEnvelopeZero{RecordsModelZero{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}},
	},

	"embedded_tfsdk_struct_nil": {
		`{}`,
		TfsdkStructs{
			BoolValue:        types.BoolNull(),
			StringValue:      types.StringNull(),
			Data:             nil,
			DataObject:       customfield.NullObject[EmbeddedTfsdkStruct](ctx),
			ListObject:       customfield.NullList[basetypes.StringValue](ctx),
			NestedObjectList: customfield.NullObjectList[EmbeddedTfsdkStruct](ctx),
			SetObject:        customfield.NullSet[basetypes.StringValue](ctx),
			NestedObjectSet:  customfield.NullObjectSet[EmbeddedTfsdkStruct](ctx),
			MapObject:        customfield.NullMap[basetypes.StringValue](ctx),
			NestedObjectMap:  customfield.NullObjectMap[EmbeddedTfsdkStruct](ctx),
			FloatValue:       types.Float64Null(),
			OptionalArray:    nil,
		},
	},
	"embedded_tfsdk_struct_nil_zero": {
		`{}`,
		TfsdkStructsZero{
			BoolValue:   types.BoolValue(false),
			StringValue: types.StringValue(""),
			Data: &EmbeddedTfsdkStructZero{
				EmbeddedString: types.StringValue(""),
				EmbeddedInt:    types.Int64Value(0),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Null(),
				}),
			},
			DataObject: customfield.NewObjectMust(ctx, &EmbeddedTfsdkStructZero{
				EmbeddedString: types.StringValue(""),
				EmbeddedInt:    types.Int64Value(0),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Null(),
				}),
			}),
			ListObject:       customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{}),
			NestedObjectList: customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStructZero{}),
			SetObject:        customfield.NewSetMust[basetypes.StringValue](ctx, []attr.Value{}),
			NestedObjectSet:  customfield.NewObjectSetMust(ctx, []EmbeddedTfsdkStructZero{}),
			MapObject:        customfield.NewMapMust(ctx, map[string]basetypes.StringValue{}),
			NestedObjectMap:  customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStructZero{}),
			FloatValue:       types.Float64Value(0),
			OptionalArray:    &[]types.String{},
		},
	},

	"json_struct_decode": {
		`{"arr":[true,1,"one"],"arr2":[true,1,"one"],"bol":false,"map":{"nil":null,"bol":false,"str":"two"},"map2":{"bol":false,"nil":null,"str":"two"},"nil":null,"num":2,"str":"two"}`,
		JsonModel{
			Arr:  jsontypes.NewNormalizedValue(`[true,1,"one"]`),
			Bol:  jsontypes.NewNormalizedValue("false"),
			Map:  jsontypes.NewNormalizedValue(`{"nil":null,"bol":false,"str":"two"}`),
			Nil:  jsontypes.NewNormalizedNull(),
			Num:  jsontypes.NewNormalizedValue("2"),
			Str:  jsontypes.NewNormalizedValue(`"two"`),
			Arr2: []jsontypes.Normalized{jsontypes.NewNormalizedValue("true"), jsontypes.NewNormalizedValue("1"), jsontypes.NewNormalizedValue(`"one"`)},
			Map2: map[string]jsontypes.Normalized{"nil": jsontypes.NewNormalizedNull(), "bol": jsontypes.NewNormalizedValue("false"), "str": jsontypes.NewNormalizedValue(`"two"`)},
		},
	},
	"json_struct_decode_zero": {
		`{"arr":[true,1,"one"],"arr2":[true,1,"one"],"bol":false,"map":{"nil":null,"bol":false,"str":"two"},"map2":{"bol":false,"nil":null,"str":"two"},"nil":null,"num":2,"str":"two"}`,
		JsonModelZero{
			Arr:  jsontypes.NewNormalizedValue(`[true,1,"one"]`),
			Bol:  jsontypes.NewNormalizedValue("false"),
			Map:  jsontypes.NewNormalizedValue(`{"nil":null,"bol":false,"str":"two"}`),
			Nil:  jsontypes.NewNormalizedValue(""),
			Num:  jsontypes.NewNormalizedValue("2"),
			Str:  jsontypes.NewNormalizedValue(`"two"`),
			Arr2: []jsontypes.Normalized{jsontypes.NewNormalizedValue("true"), jsontypes.NewNormalizedValue("1"), jsontypes.NewNormalizedValue(`"one"`)},
			Map2: map[string]jsontypes.Normalized{"nil": jsontypes.NewNormalizedValue(""), "bol": jsontypes.NewNormalizedValue("false"), "str": jsontypes.NewNormalizedValue(`"two"`)},
		},
	},

	"json_struct_nil3": {`{"nil":null}`, JsonModel{}},
	"json_struct_nil3_zero": {`{"nil":null}`, JsonModelZero{
		Arr:  jsontypes.NewNormalizedValue(""),
		Bol:  jsontypes.NewNormalizedValue(""),
		Map:  jsontypes.NewNormalizedValue(""),
		Nil:  jsontypes.NewNormalizedValue(""),
		Num:  jsontypes.NewNormalizedValue(""),
		Str:  jsontypes.NewNormalizedValue(""),
		Arr2: []jsontypes.Normalized{},
		Map2: map[string]jsontypes.Normalized{},
	}},

	"nested_object_list_missing_nested_field": {
		`{"a":[{"b":"foo"}}]}`,
		ListWithNestedObj{
			A: customfield.NewObjectListMust(ctx, []Embedded2{
				{
					B: types.StringValue("foo"),
					C: nil,
					D: nil,
					E: nil,
					F: nil,
				},
			}),
		},
	},
	"nested_object_list_missing_nested_field_zero": {
		`{"a":[{"b":"foo"}}]}`,
		ListWithNestedObjZero{
			A: customfield.NewObjectListMust(ctx, []Embedded2Zero{
				{
					B: types.StringValue("foo"),
					C: &InnerZero{
						D: types.StringValue(""),
					},
					D: &[]*InnerZero{},
					E: []string{},
					F: &map[string]InnerZero{},
				},
			}),
		},
	},
}

var encodeOnlyTests = map[string]struct {
	buf string
	val interface{}
}{
	"tfsdk_struct_encode": {
		`{"result":{"a":"88281d6015751d6172e7313b0c665b5e","b":"http://example.com/example.html\t20"}}`,
		ResultEnvelope{RecordsModel{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}},
	},

	"embedded_tfsdk_struct_nil": {
		`{}`,
		TfsdkStructs{
			BoolValue:   types.BoolNull(),
			StringValue: types.StringNull(),
			FloatValue:  types.Float64Null(),
		},
	},

	"json_struct_encode": {
		`{"arr":[true,1,"one"],"arr2":[true,1,"one"],"bol":false,"map":{"nil":null,"bol":false,"str":"two"},"map2":{"bol":false,"nil":null,"str":"two"},"nil":null,"num":2,"str":"two"}`,
		JsonModel{
			Arr:  jsontypes.NewNormalizedValue(`[true,1,"one"]`),
			Bol:  jsontypes.NewNormalizedValue("false"),
			Map:  jsontypes.NewNormalizedValue(`{"nil":null,"bol":false,"str":"two"}`),
			Nil:  jsontypes.NewNormalizedValue("null"),
			Num:  jsontypes.NewNormalizedValue("2"),
			Str:  jsontypes.NewNormalizedValue(`"two"`),
			Arr2: []jsontypes.Normalized{jsontypes.NewNormalizedValue("true"), jsontypes.NewNormalizedValue("1"), jsontypes.NewNormalizedValue(`"one"`)},
			Map2: map[string]jsontypes.Normalized{"nil": jsontypes.NewNormalizedNull(), "bol": jsontypes.NewNormalizedValue("false"), "str": jsontypes.NewNormalizedValue(`"two"`)},
		},
	},

	"json_struct_nil3": {`{"nil":null}`, JsonModel{Nil: jsontypes.NewNormalizedValue("null")}},

	"tfsdk_dynamic_number": {"5", types.DynamicValue(types.NumberValue(big.NewFloat(5)))},

	"tfsdk_dynamic_tuple": {
		`[5,"hi"]`,
		types.DynamicValue(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		)),
	},

	"tfsdk_tuple": {
		`[5,"hi"]`,
		types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		),
	},

	"tfsdk_nested_tuple": {
		`[10,["hey","there"]]`,
		types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.ListType{ElemType: basetypes.StringType{}}},
			[]attr.Value{basetypes.NewInt64Value(10), types.ListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("hey"), basetypes.NewStringValue("there")})},
		),
	},

	"complex_nested_list_object": {
		`{"a":[{"b":"foo","c":{"d":"pointer_inner"},"d":[{"d":"list_pointer_inner_1"},{"d":"list_pointer_inner_2"}],"e":["a","b"],"f":{"a_key":{"d":"a_value"}}}]}`,
		ListWithNestedObj{
			A: customfield.NewObjectListMust(ctx, []Embedded2{
				{
					B: types.StringValue("foo"),
					C: P(Inner{D: types.StringValue("pointer_inner")}),
					D: P([]*Inner{P(Inner{D: types.StringValue("list_pointer_inner_1")}), P(Inner{D: types.StringValue("list_pointer_inner_2")})}),
					E: []string{"a", "b"},
					F: P(map[string]Inner{
						"a_key": {D: types.StringValue("a_value")},
					}),
				},
			}),
		},
	},

	"nested_map_pointer": {
		`{"outer":[{"a":{"a.b.*":"*"}}]}`,
		struct {
			Outer *[]*structWithMap `json:"outer,required"`
		}{
			Outer: P([]*structWithMap{
				P(structWithMap{A: P(map[string]types.String{"a.b.*": types.StringValue("*")})}),
			}),
		},
	},
}

type structWithMap struct {
	A *map[string]types.String `json:"a,required"`
}

func TestDecode(t *testing.T) {
	spew.Config.SortKeys = true
	for name, test := range merge(tests, decode_only_tests) {
		t.Run(name, func(t *testing.T) {
			resultValue := reflect.New(reflect.TypeOf(test.val))
			if err := Unmarshal([]byte(test.buf), resultValue.Interface()); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", resultValue, err)
			}
			result := resultValue.Elem().Interface()
			if !reflect.DeepEqual(result, test.val) {
				t.Fatalf("incorrect deserialization for '%s':\nexpected:\n%s\nactual:\n%s\n", test.buf, spew.Sdump(test.val), spew.Sdump(result))
			}
		})
	}
}

func TestEncode(t *testing.T) {
	for name, test := range merge(tests, encodeOnlyTests) {
		if strings.HasSuffix(name, "_coerce") ||
			strings.HasSuffix(name, "_zero") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			raw, err := Marshal(test.val)
			if err != nil {
				t.Fatalf("serialization of %v failed with error %v", test.val, err)
			}
			if string(raw) != test.buf {
				var expected, actual string
				errExpected := formatJson(test.buf, &expected)
				if errExpected != nil {
					// invalid json in the expected string is a test error so we panic
					panic(fmt.Sprintf("invalid expected JSON:\n%s\n%v", test.buf, errExpected))
				}
				errActual := formatJson(string(raw), &actual)
				if errActual != nil {
					t.Fatalf("invalid actual JSON:\n%s\n%v", string(raw), errActual)
				}
				t.Fatalf("expected:\n%s\nto serialize to \n%s\n but got \n%s\n", spew.Sdump(test.val), expected, actual)
			}
		})

	}
}

var updateTests = map[string]struct {
	state         interface{}
	plan          interface{}
	expected      string
	expectedPatch string
}{
	"true":           {true, true, "true", ""},
	"terraform_true": {types.BoolValue(true), types.BoolValue(true), "true", ""},

	"null to true":   {types.BoolNull(), types.BoolValue(true), "true", "true"},
	"false to true":  {types.BoolValue(false), types.BoolValue(true), "true", "true"},
	"unset bool":     {types.BoolValue(false), types.BoolNull(), "null", "null"},
	"omit null bool": {types.BoolNull(), types.BoolNull(), "", ""},

	"string set":       {types.StringNull(), types.StringValue("two"), `"two"`, `"two"`},
	"string update":    {types.StringValue("one"), types.StringValue("two"), `"two"`, `"two"`},
	"unset string":     {types.StringValue("hey"), types.StringNull(), "null", "null"},
	"omit null string": {types.StringNull(), types.StringNull(), "", ""},
	"string unchanged": {types.StringValue("one"), types.StringValue("one"), `"one"`, ""},

	"int set":       {types.Int64Null(), types.Int64Value(42), "42", "42"},
	"int update":    {types.Int64Value(42), types.Int64Value(43), "43", "43"},
	"unset int":     {types.Int64Value(42), types.Int64Null(), "null", "null"},
	"omit null int": {types.Int64Null(), types.Int64Null(), "", ""},
	"int unchanged": {types.Int64Value(42), types.Int64Value(42), "42", ""},

	"tuple set": {
		types.TupleNull([]attr.Type{types.Int64Type, types.StringType}),
		types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")}),
		`[1,"two"]`,
		`[1,"two"]`,
	},
	"tuple update": {
		types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")}),
		types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("three")}),
		`[1,"three"]`,
		`[1,"three"]`,
	},
	"tuple unset": {
		types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")}),
		types.TupleNull([]attr.Type{types.Int64Type, types.StringType}),
		`null`,
		`null`,
	},
	"tuple omit null": {
		types.TupleNull([]attr.Type{types.Int64Type, types.StringType}),
		types.TupleNull([]attr.Type{types.Int64Type, types.StringType}),
		``,
		``,
	},
	"tuple unchanged": {
		types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")}),
		types.TupleValueMust([]attr.Type{types.Int64Type, types.StringType}, []attr.Value{types.Int64Value(1), types.StringValue("two")}),
		`[1,"two"]`,
		``,
	},

	"dynamic omit null":                     {types.DynamicNull(), types.DynamicNull(), "", ""},
	"dynamic omit underlying null state":    {types.DynamicValue(types.Int64Null()), types.DynamicNull(), "", ""},
	"dynamic omit underlying null plan":     {types.DynamicNull(), types.DynamicValue(types.Int64Null()), "", ""},
	"dynamic omit unknown":                  {types.DynamicUnknown(), types.DynamicUnknown(), "", ""},
	"dynamic omit underlying unknown state": {types.DynamicValue(types.Int64Unknown()), types.DynamicUnknown(), "", ""},
	"dynamic omit underlying unknown plan":  {types.DynamicUnknown(), types.DynamicValue(types.Int64Unknown()), "", ""},
	"dynamic unset null":                    {types.DynamicValue(types.Int64Value(4)), types.DynamicNull(), "null", "null"},
	"dynamic int set":                       {types.DynamicNull(), types.DynamicValue(types.Int64Value(5)), "5", "5"},
	"dynamic int update":                    {types.DynamicValue(types.Int64Value(4)), types.DynamicValue(types.Int64Value(5)), "5", "5"},
	"dynamic int unchanged":                 {types.DynamicValue(types.Int64Value(4)), types.DynamicValue(types.Int64Value(4)), "4", ""},

	// Test case for dynamic type conversion: state has ListValue, plan has TupleValue
	"dynamic list to tuple conversion": {
		types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		types.DynamicValue(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		`["foo","bar"]`,
		``,
	},

	"normalized list to tuple conversion": {
		customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		`["foo","bar"]`,
		``,
	},

	// Test case for reverse scenario: state has TupleValue, plan has ListValue
	"dynamic tuple to list conversion": {
		types.DynamicValue(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		`["foo","bar"]`,
		``,
	},

	"normalized dynamic tuple to list conversion": {
		customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{types.StringType, types.StringType}, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")})),
		`["foo","bar"]`,
		``,
	},

	// Test case for heterogeneous tuple vs homogeneous list
	"dynamic list to heterogeneous tuple": {
		types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("world")})),
		types.DynamicValue(types.TupleValueMust([]attr.Type{types.StringType, types.Int64Type}, []attr.Value{types.StringValue("hello"), types.Int64Value(42)})),
		`["hello",42]`,
		`["hello",42]`,
	},

	"normalized dynamic list to heterogeneous tuple": {
		customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("world")})),
		customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust([]attr.Type{types.StringType, types.Int64Type}, []attr.Value{types.StringValue("hello"), types.Int64Value(42)})),
		`["hello",42]`,
		`["hello",42]`,
	},

	"set struct fields": {
		TfsdkStructs{},
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
	},

	"update some struct fields": {
		TfsdkStructs{
			BoolValue:   types.BoolValue(true),
			StringValue: types.StringValue("string_value"),
			FloatValue:  types.Float64Value(3.14),
		},
		TfsdkStructs{
			BoolValue:   types.BoolValue(false),
			StringValue: types.StringValue("another_string"),
			FloatValue:  types.Float64Value(1.14),
		},
		`{"bool_value":false,"float_value":1.14,"string_value":"another_string"}`,
		`{"bool_value":false,"float_value":1.14,"string_value":"another_string"}`,
	},

	"unset nested struct fields": {
		TfsdkStructs{
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedInt: types.Int64Value(17),
			},
		},
		TfsdkStructs{
			OptionalArray: &[]types.String{types.StringValue("hi")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedInt: types.Int64Null(),
			},
		},
		`{"data":{"embedded_int":null},"optional_array":["hi"]}`,
		`{"data":{"embedded_int":null},"optional_array":["hi"]}`,
	},

	"unset struct fields": {
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
		},
		TfsdkStructs{},
		`{"bool_value":null,"data":null,"float_value":null,"optional_array":null,"string_value":null}`,
		`{"bool_value":null,"data":null,"float_value":null,"optional_array":null,"string_value":null}`,
	},

	"set empty array": {
		TfsdkStructs{
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
		},
		TfsdkStructs{
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{},
		},
		`{"float_value":3.14,"optional_array":[]}`,
		`{"optional_array":[]}`,
	},

	"set nested map": {
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{}),
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		}),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"unchanged nested struct": {
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		``,
	},

	"nested value changed in nested struct": {
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("changed_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"changed_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		`{"data":{"embedded_string":"changed_string_value"}}`,
	},

	"set array element": {
		TfsdkStructs{
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("one"), types.StringValue("two")},
			ListObject:    customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("three"), types.StringValue("four")}),
		},
		TfsdkStructs{
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("five"), types.StringValue("two")},
			ListObject:    customfield.NewListMust[basetypes.StringValue](ctx, []attr.Value{types.StringValue("six"), types.StringValue("four")}),
		},
		`{"float_value":3.14,"list_object":["six","four"],"optional_array":["five","two"]}`,
		`{"list_object":["six","four"],"optional_array":["five","two"]}`,
	},

	"set nested array value": {
		customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			{
				EmbeddedString: types.StringValue("string value2"),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(19),
				}),
			},
		}),
		customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			{
				EmbeddedString: types.StringValue("string value2"),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(20), // only changed this property
				}),
			},
		}),
		`[{"embedded_string":"string value"},{"data_object":{"nested_int":20},"embedded_string":"string value2"}]`,
		`[{"embedded_string":"string value"},{"data_object":{"nested_int":20},"embedded_string":"string value2"}]`,
	},

	"remove array value encodes": {
		customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
			{
				EmbeddedString: types.StringValue("string value2"),
				DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
					NestedInt: types.Int64Value(20),
				}),
			},
		}),
		customfield.NewObjectListMust(ctx, []EmbeddedTfsdkStruct{
			{
				EmbeddedString: types.StringValue("string value"),
				DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
			},
		}),
		`[{"embedded_string":"string value"}]`,
		`[{"embedded_string":"string value"}]`,
	},

	"set custom map list": {
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{}),
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		}),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"set built-in map list": {
		map[string][]*string{},
		map[string][]*string{
			"Key1": {P("Value1")},
			"Key2": {P("Value2")},
		},
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"remove all keys from a custom map": {
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		}),
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{}),
		`{}`,
		`{}`,
	},

	"update to add a key to a custom map": {
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
		}),
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		}),
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value2"]}`,
	},

	"update a nested array in a custom map": {
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value2")})),
		}),
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
			"Key2": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value3"), basetypes.NewStringValue("Value2")})),
		}),
		`{"Key1":["Value1"],"Key2":["Value3","Value2"]}`,
		`{"Key1":["Value1"],"Key2":["Value3","Value2"]}`,
	},

	"unset custom map": {
		customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
			"Key1": DropDiagnostic(customfield.NewList[types.String](ctx, []types.String{basetypes.NewStringValue("Value1")})),
		}),
		customfield.NullMap[customfield.List[types.String]](ctx),
		`null`,
		`null`,
	},

	"unset key in built-in map": {
		map[string]*string{
			"Key1": P("Value1"),
			"Key2": P("Value2"),
		},
		map[string]*string{
			"Key1": P("Value1"),
		},
		`{"Key1":"Value1"}`,
		`{"Key1":"Value1"}`,
	},

	"set custom object map": {
		customfield.NullObjectMap[TfsdkStructs](ctx),
		customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"Key1": {
				BoolValue:     types.BoolValue(true),
				StringValue:   types.StringValue("string_value"),
				FloatValue:    types.Float64Value(3.14),
				OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
				Data: &EmbeddedTfsdkStruct{
					EmbeddedString: types.StringValue("embedded_string_value"),
					EmbeddedInt:    types.Int64Value(17),
				},
			},
		}),
		`{"Key1":{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}}`,
		`{"Key1":{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}}`,
	},

	"set nested value on custom object map": {
		customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"OuterKey": {
				NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStruct{
					"NestedKey": {
						EmbeddedInt:    types.Int64Value(16),
						EmbeddedString: types.StringValue("nested_string_value"),
					},
				}),
			},
		}),
		customfield.NewObjectMapMust(ctx, map[string]TfsdkStructs{
			"OuterKey": {
				NestedObjectMap: customfield.NewObjectMapMust(ctx, map[string]EmbeddedTfsdkStruct{
					"NestedKey": {
						EmbeddedInt:    types.Int64Value(17),
						EmbeddedString: types.StringValue("nested_string_value"),
					},
				}),
			},
		}),
		`{"OuterKey":{"nested_object_map":{"NestedKey":{"embedded_int":17,"embedded_string":"nested_string_value"}}}}`,
		`{"OuterKey":{"nested_object_map":{"NestedKey":{"embedded_int":17,"embedded_string":"nested_string_value"}}}}`,
	},

	"encode_state_for_unknown with unknown plan": {
		EncodeStateForUnknownStruct{
			NormalField:                     types.StringValue("state_normal"),
			ComputedWithForceEncode:         types.StringValue("computed value from state"),
			ComputedWithStateEncode:         types.StringValue("computed value 2"),
			ComputedOptionalWithStateEncode: types.StringValue("computed optional from state"),
			ComputedOptionalRegular:         types.StringValue("computed optional regular"),
			ComputedRegular:                 types.StringValue("computed regular"),
		},
		EncodeStateForUnknownStruct{
			NormalField:                     types.StringUnknown(),
			ComputedWithForceEncode:         types.StringUnknown(),
			ComputedWithStateEncode:         types.StringUnknown(),
			ComputedOptionalWithStateEncode: types.StringUnknown(),
			ComputedOptionalRegular:         types.StringUnknown(),
			ComputedRegular:                 types.StringUnknown(),
		},
		// Expected result: only values with "encode_state_for_unknown" are encoded
		`{"computed_optional_state_encode":"computed optional from state","computed_state_encode":"computed value 2"}`,
		// NOTE: force_encode should probably override patch behavior, but we don't support that for now
		``,
	},

	"encode_state_for_unknown with known plan": {
		EncodeStateForUnknownStruct{
			NormalField:                     types.StringValue("state_normal"),
			ComputedWithForceEncode:         types.StringValue("computed value from state"),
			ComputedWithStateEncode:         types.StringValue("computed value 2"),
			ComputedOptionalWithStateEncode: types.StringValue("computed optional from state"),
			ComputedOptionalRegular:         types.StringValue("computed optional regular"),
			ComputedRegular:                 types.StringValue("computed regular"),
		},
		EncodeStateForUnknownStruct{
			NormalField:                     types.StringValue("plan normal"),
			ComputedWithForceEncode:         types.StringValue("plan A"),
			ComputedWithStateEncode:         types.StringValue("plan B"),
			ComputedOptionalWithStateEncode: types.StringValue("plan C"),
			ComputedOptionalRegular:         types.StringValue("plan D"),
			ComputedRegular:                 types.StringValue("plan E"),
		},
		// Expected result: we use value from plan for all computed optional fields
		// & for computed fields with force_encode state
		`{"computed_force_encode":"plan A","computed_optional_regular":"plan D","computed_optional_state_encode":"plan C","computed_state_encode":"plan B","normal_field":"plan normal"}`,
		// These show up even w/ patch b/c plan and state values are different; in reality, computed value shouldn't differ b/t plan and state
		`{"computed_force_encode":"plan A","computed_optional_regular":"plan D","computed_optional_state_encode":"plan C","computed_state_encode":"plan B","normal_field":"plan normal"}`},

	"encode_state_for_unknown with null state": {
		EncodeStateForUnknownStruct{
			NormalField:                     types.StringNull(),
			ComputedWithForceEncode:         types.StringNull(),
			ComputedWithStateEncode:         types.StringNull(),
			ComputedOptionalWithStateEncode: types.StringNull(),
			ComputedOptionalRegular:         types.StringNull(),
			ComputedRegular:                 types.StringNull(),
		},
		EncodeStateForUnknownStruct{
			NormalField:                     types.StringUnknown(),
			ComputedWithForceEncode:         types.StringUnknown(),
			ComputedWithStateEncode:         types.StringUnknown(),
			ComputedOptionalWithStateEncode: types.StringUnknown(),
			ComputedOptionalRegular:         types.StringUnknown(),
			ComputedRegular:                 types.StringUnknown(),
		},
		// Don't copy null fields from state
		`{}`,
		``,
	},
}

func TestUpdateEncoding(t *testing.T) {
	for name, test := range updateTests {
		t.Run(name, func(t *testing.T) {
			t.Run("MarshalForUpdate", func(t *testing.T) {
				raw, err := MarshalForUpdate(test.plan, test.state)
				if err != nil {
					t.Fatalf("serialization of %v, %v failed with error %v", test.plan, test.state, err)
				}
				if string(raw) != test.expected {
					t.Fatalf("expected %+#v, %+#v to serialize to \n%s\n but got \n%s\n", test.state, test.plan, test.expected, string(raw))
				}
			})
			t.Run("MarshalForPatch", func(t *testing.T) {
				raw, err := MarshalForPatch(test.plan, test.state)
				if err != nil {
					t.Fatalf("serialization of %v, %v failed with error %v", test.plan, test.state, err)
				}
				if string(raw) != test.expectedPatch {
					t.Fatalf("expected %+#v, %+#v to serialize to \n%s\n but got \n%s\n", test.state, test.plan, test.expectedPatch, string(raw))
				}
			})
		})
	}
}

var decode_from_value_tests = map[string]struct {
	buf      string
	starting interface{}
	expected interface{}
}{

	"tfsdk_dynamic_null": {
		`null`,
		types.DynamicNull(),
		types.DynamicNull(),
	},

	"tfsdk_dynamic_string_from_null": {
		`"hey"`,
		types.DynamicNull(),
		types.DynamicValue(types.StringValue("hey")),
	},

	"tfsdk_dynamic_string_from_unknown": {
		`"hey"`,
		types.DynamicUnknown(),
		types.DynamicValue(types.StringValue("hey")),
	},

	"tfsdk_dynamic_string_from_value": {
		`"hey"`,
		types.DynamicValue(types.StringValue("before_value")),
		types.DynamicValue(types.StringValue("hey")),
	},

	"tfsdk_dynamic_number": {
		"5",
		types.DynamicValue(basetypes.NewNumberNull()),
		types.DynamicValue(types.NumberValue(big.NewFloat(5))),
	},

	"tfsdk_dynamic_int_from_null": {
		`14`,
		types.DynamicNull(),
		types.DynamicValue(types.Int64Value(14)),
	},

	"tfsdk_dynamic_int_from_unknown": {
		`14`,
		types.DynamicUnknown(),
		types.DynamicValue(types.Int64Value(14)),
	},

	"tfsdk_dynamic_int_from_value": {
		`14`,
		types.DynamicValue(types.Int64Value(5)),
		types.DynamicValue(types.Int64Value(14)),
	},

	"tfsdk_dynamic_tuple": {
		`[5,"hi"]`,
		types.DynamicValue(types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		)),
		types.DynamicValue(types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		)),
	},

	"tfsdk_map_value": {
		`{"foo":1,"bar":4}`,
		types.MapNull(types.Int64Type),
		types.MapValueMust(types.Int64Type, map[string]attr.Value{"foo": types.Int64Value(1), "bar": types.Int64Value(4)}),
	},

	"tfsdk_tuple": {
		`[5,"hi"]`,
		types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		),
		types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		),
	},

	"tfsdk_tuple_existing": {
		`[10,"hello there"]`,
		types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(5), basetypes.NewStringValue("hi")},
		),
		types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Value(10), basetypes.NewStringValue("hello there")},
		),
	},

	"tfsdk_tuple_missing_values": {
		`[]`,
		types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		),
		types.TupleValueMust(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
			[]attr.Value{basetypes.NewInt64Null(), basetypes.NewStringNull()},
		),
	},

	"tfsdk_tuple_single_object": {
		`[{"non":"array"}]`,
		types.TupleNull(
			[]attr.Type{basetypes.ObjectType{AttrTypes: map[string]attr.Type{"non": basetypes.StringType{}}}, basetypes.StringType{}},
		),
		types.TupleValueMust(
			[]attr.Type{basetypes.ObjectType{AttrTypes: map[string]attr.Type{"non": basetypes.StringType{}}}, basetypes.StringType{}},
			[]attr.Value{
				basetypes.NewObjectValueMust(
					map[string]attr.Type{"non": basetypes.StringType{}},
					map[string]attr.Value{"non": basetypes.NewStringValue("array")},
				),
				basetypes.NewStringNull(),
			},
		),
	},

	"tfsdk_tuple_non_array_num_error": {
		`5`,
		types.TupleNull(
			[]attr.Type{basetypes.Int64Type{}, basetypes.StringType{}},
		),
		fmt.Errorf("apijson: cannot deserialize unexpected type Number to types.TupleValue"),
	},

	"tfsdk_tuple_non_array_object_error": {
		`{"non":"array"}`,
		types.TupleNull(
			[]attr.Type{basetypes.ObjectType{AttrTypes: map[string]attr.Type{"non": basetypes.StringType{}}}, basetypes.StringType{}},
		),
		fmt.Errorf("apijson: cannot deserialize unexpected type JSON to types.TupleValue"),
	},

	"tfsdk_map_value_existing_data": {
		`{"foo":1,"bar":4}`,
		types.MapValueMust(types.Int64Type, map[string]attr.Value{"baz": types.Int64Value(2)}),
		types.MapValueMust(types.Int64Type, map[string]attr.Value{"foo": types.Int64Value(1), "bar": types.Int64Value(4)}),
	},

	"tfsdk_object_with_attributes": {
		`{"baz":4,"foo":["bar","baz"]}`,
		types.ObjectNull(
			map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
		),
		types.ObjectValueMust(
			map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
			map[string]attr.Value{"baz": types.Int64Value(4), "foo": types.SetValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("baz")})},
		),
	},

	"tfsdk_dynamic_object_with_attributes": {
		`{"baz":4,"foo":["bar","baz"]}`,
		types.DynamicValue(
			types.ObjectNull(
				map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
			),
		),
		types.DynamicValue(
			types.ObjectValueMust(
				map[string]attr.Type{"baz": types.Int64Type, "foo": types.SetType{ElemType: types.StringType}},
				map[string]attr.Value{"baz": types.Int64Value(4), "foo": types.SetValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("baz")})},
			),
		),
	},

	// note it creates a list this time because the dynamic doesn't contain type information
	"tfsdk_dynamic_object_without_attributes": {
		`{"baz":4,"foo":["bar","baz"]}`,
		types.DynamicNull(),
		types.DynamicValue(
			types.ObjectValueMust(
				map[string]attr.Type{"baz": types.Int64Type, "foo": types.ListType{ElemType: types.StringType}},
				map[string]attr.Value{"baz": types.Int64Value(4), "foo": types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("baz")})},
			),
		),
	},

	// Test case for heterogeneous JSON array inference - should create TupleValue, not ListValue
	"tfsdk_dynamic_heterogeneous_array_inference": {
		`["hello",42]`,
		types.DynamicNull(),
		types.DynamicValue(types.TupleValueMust(
			[]attr.Type{types.StringType, types.Int64Type},
			[]attr.Value{types.StringValue("hello"), types.Int64Value(42)},
		)),
	},

	"tfsdk_normalized_dynamic_heterogeneous_array_inference": {
		`["hello",42]`,
		customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull()),
		customfield.RawNormalizedDynamicValueFrom(types.TupleValueMust(
			[]attr.Type{types.StringType, types.Int64Type},
			[]attr.Value{types.StringValue("hello"), types.Int64Value(42)},
		)),
	},

	// Test case for homogeneous JSON array inference - should still create ListValue
	"tfsdk_dynamic_homogeneous_array_inference": {
		`["hello","world"]`,
		types.DynamicNull(),
		types.DynamicValue(types.ListValueMust(
			types.StringType,
			[]attr.Value{types.StringValue("hello"), types.StringValue("world")},
		)),
	},

	"tfsdk_normalized_dynamic_homogeneous_array_inference": {
		`["hello","world"]`,
		customfield.RawNormalizedDynamicValue(basetypes.NewDynamicNull()),
		customfield.RawNormalizedDynamicValueFrom(types.ListValueMust(
			types.StringType,
			[]attr.Value{types.StringValue("hello"), types.StringValue("world")},
		)),
	},

	"tfsdk_struct_populates_unknown_to_null_if_missing": {
		`{"embedded_string":"some_string","data_object":{}}`,
		EmbeddedTfsdkStruct{
			EmbeddedString: types.StringUnknown(),
			EmbeddedInt:    types.Int64Unknown(),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		},
		EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("some_string"),
			EmbeddedInt:    types.Int64Null(),
			DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
				NestedInt: types.Int64Null(),
			}),
		},
	},
	"tfsdk_struct_populates_unknown_to_null_if_missing_zero": {
		`{"embedded_string":"some_string","data_object":{}}`,
		EmbeddedTfsdkStructZero{
			EmbeddedString: types.StringUnknown(),
			EmbeddedInt:    types.Int64Unknown(),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		},
		EmbeddedTfsdkStructZero{
			EmbeddedString: types.StringValue("some_string"),
			EmbeddedInt:    types.Int64Value(0),
			DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
				NestedInt: types.Int64Null(),
			}),
		},
	},

	"tfsdk_struct_overwrites_from_json": {
		`{"embedded_string":"new_value"}`,
		EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		},
		EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("new_value"),
			EmbeddedInt:    types.Int64Null(),
			DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
		},
	},
	"tfsdk_struct_overwrites_from_json_zero": {
		`{"embedded_string":"new_value"}`,
		EmbeddedTfsdkStructZero{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		},
		EmbeddedTfsdkStructZero{
			EmbeddedString: types.StringValue("new_value"),
			EmbeddedInt:    types.Int64Value(0),
			DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
				NestedInt: types.Int64Null(),
			}),
		},
	},

	"tfsdk_date_time_populates_unknown_to_null_if_missing": {
		`{"date":"2006-01-02"}`,
		DateTimeCustom{
			DateCustom:     timetypes.NewRFC3339Unknown(),
			DateTimeCustom: timetypes.NewRFC3339Unknown(),
		},
		DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339Null(),
		},
	},
	"tfsdk_date_time_populates_unknown_to_null_if_missing_zero": {
		`{"date":"2006-01-02"}`,
		DateTimeCustomZero{
			DateCustom:     timetypes.NewRFC3339Unknown(),
			DateTimeCustom: timetypes.NewRFC3339Unknown(),
		},
		DateTimeCustomZero{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: timetypes.NewRFC3339TimeValue(time.Time{}),
		},
	},
}

func TestDecodeFromValue(t *testing.T) {
	spew.Config.ContinueOnMethod = true
	for name, test := range decode_from_value_tests {
		t.Run(name, func(t *testing.T) {
			v := reflect.ValueOf(test.starting)
			starting := reflect.New(v.Type())
			starting.Elem().Set(v)

			expectedErr, errorIsExpected := test.expected.(error)

			err := Unmarshal([]byte(test.buf), starting.Interface())
			if errorIsExpected {
				if err == nil {
					t.Fatalf(`expected error "%s" but did not error`, expectedErr.Error())
				}
				if err.Error() != expectedErr.Error() {
					t.Fatalf(`expected error "%s" but got "%s"`, expectedErr.Error(), err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("deserialization of %v failed with error %v", test.buf, err)
				}
				startingIFace := starting.Elem().Interface()
				if !reflect.DeepEqual(startingIFace, test.expected) {
					t.Fatalf("expected '%s' to deserialize to \n%s\nbut got\n%s", test.buf, spew.Sdump(test.expected), spew.Sdump(startingIFace))
				}
			}
		})
	}
}

var decode_unset_tests = map[string]struct {
	buf string
	val interface{}
}{
	"nested_object_list_is_omitted_null": {
		`{}`,
		ListWithNestedObj{
			A: customfield.NullObjectList[Embedded2](ctx),
		},
	},
	"nested_object_list_is_omitted_null_zero": {
		`{}`,
		ListWithNestedObjZero{
			A: customfield.NewObjectListMust(ctx, []Embedded2Zero{}),
		},
	},
	"nested_object_list_is_explicit_null": {
		`{"a": null}`,
		ListWithNestedObj{
			A: customfield.NullObjectList[Embedded2](ctx),
		},
	},
	"nested_object_list_is_explicit_null_zero": {
		`{"a": null}`,
		ListWithNestedObjZero{
			A: customfield.NewObjectListMust(ctx, []Embedded2Zero{}),
		},
	},
	"nested_object_list_is_empty": {
		`{"a": []}`,
		ListWithNestedObj{
			A: customfield.NewObjectListMust(ctx, []Embedded2{}),
		},
	},
	"nested_object_list_is_empty_zero": {
		`{"a": []}`,
		ListWithNestedObjZero{
			A: customfield.NewObjectListMust(ctx, []Embedded2Zero{}),
		},
	},
}

func TestDecodeUnsetBehaviour(t *testing.T) {
	spew.Config.SortKeys = true
	for name, test := range merge(decode_unset_tests) {
		t.Run(name, func(t *testing.T) {
			resultValue := reflect.New(reflect.TypeOf(test.val))
			d := &decoderBuilder{
				dateFormat:            time.RFC3339,
				unmarshalComputedOnly: false,
				updateBehavior:        IfUnset,
			}
			if err := d.unmarshal([]byte(test.buf), resultValue.Interface()); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", resultValue, err)
			}
			result := resultValue.Elem().Interface()
			if !reflect.DeepEqual(result, test.val) {
				t.Fatalf("incorrect deserialization for '%s':\nexpected:\n%s\nactual:\n%s\n", test.buf, spew.Sdump(test.val), spew.Sdump(result))
			}
		})
	}
}

type StructWithComputedFields struct {
	RegStr            types.String                                             `tfsdk:"str" json:"str,optional"`
	CompStr           types.String                                             `tfsdk:"comp_str" json:"comp_str,computed"`
	CompOptStr        types.String                                             `tfsdk:"opt_str" json:"opt_str,computed_optional"`
	CompTime          timetypes.RFC3339                                        `tfsdk:"time" json:"time,computed"`
	CompOptTime       timetypes.RFC3339                                        `tfsdk:"opt_time" json:"opt_time,computed_optional"`
	Nested            NestedStructWithComputedFields                           `tfsdk:"nested" json:"nested,optional"`
	NestedCust        customfield.NestedObject[NestedStructWithComputedFields] `tfsdk:"nested_obj" json:"nested_obj,optional"`
	CompOptNestedCust customfield.NestedObject[NestedStructWithComputedFields] `tfsdk:"opt_nested_obj" json:"opt_nested_obj,computed_optional"`
	NestedList        *[]*NestedStructWithComputedFields                       `tfsdk:"nested_list" json:"nested_list,optional"`
	MapCust           customfield.Map[customfield.List[types.String]]          `tfsdk:"map_cust" json:"map_cust,optional"`
	MapRegular        *map[string][]*NestedStructWithComputedFields            `tfsdk:"map_regular" json:"map_regular,optional"`
	CompMap           *map[string]*NestedStructWithComputedFields              `tfsdk:"comp_map" json:"comp_map,computed"`
	CompMapList       *map[string][]*NestedStructWithComputedFields            `tfsdk:"comp_map_list" json:"comp_map_list,computed"`
}

type NestedStructWithComputedFields struct {
	RegStr     types.String `tfsdk:"nested_str" json:"nested_str,required"`
	CompStr    types.String `tfsdk:"nested_comp_str" json:"nested_comp_str,computed"`
	CompOptInt types.Int64  `tfsdk:"nested_comp_opt_int" json:"nested_comp_opt_int,computed_optional"`
}

type StructWithComputedFieldsZero struct {
	RegStr            types.String                                                 `tfsdk:"str" json:"str,optional,decode_null_to_zero"`
	CompStr           types.String                                                 `tfsdk:"comp_str" json:"comp_str,computed,decode_null_to_zero"`
	CompOptStr        types.String                                                 `tfsdk:"opt_str" json:"opt_str,computed_optional,decode_null_to_zero"`
	CompTime          timetypes.RFC3339                                            `tfsdk:"time" json:"time,computed,decode_null_to_zero"`
	CompOptTime       timetypes.RFC3339                                            `tfsdk:"opt_time" json:"opt_time,computed_optional,decode_null_to_zero"`
	Nested            NestedStructWithComputedFieldsZero                           `tfsdk:"nested" json:"nested,optional,decode_null_to_zero"`
	NestedCust        customfield.NestedObject[NestedStructWithComputedFieldsZero] `tfsdk:"nested_obj" json:"nested_obj,optional,decode_null_to_zero"`
	CompOptNestedCust customfield.NestedObject[NestedStructWithComputedFieldsZero] `tfsdk:"opt_nested_obj" json:"opt_nested_obj,computed_optional,decode_null_to_zero"`
	NestedList        *[]*NestedStructWithComputedFieldsZero                       `tfsdk:"nested_list" json:"nested_list,optional,decode_null_to_zero"`
	MapCust           customfield.Map[customfield.List[types.String]]              `tfsdk:"map_cust" json:"map_cust,optional,decode_null_to_zero"`
	MapRegular        *map[string][]*NestedStructWithComputedFieldsZero            `tfsdk:"map_regular" json:"map_regular,optional,decode_null_to_zero"`
	CompMap           *map[string]*NestedStructWithComputedFieldsZero              `tfsdk:"comp_map" json:"comp_map,computed,decode_null_to_zero"`
	CompMapList       *map[string][]*NestedStructWithComputedFieldsZero            `tfsdk:"comp_map_list" json:"comp_map_list,computed,decode_null_to_zero"`
}

type NestedStructWithComputedFieldsZero struct {
	RegStr     types.String `tfsdk:"nested_str" json:"nested_str,required,decode_null_to_zero"`
	CompStr    types.String `tfsdk:"nested_comp_str" json:"nested_comp_str,computed,decode_null_to_zero"`
	CompOptInt types.Int64  `tfsdk:"nested_comp_opt_int" json:"nested_comp_opt_int,computed_optional,decode_null_to_zero"`
}

var exampleNestedJson = `{
	"str":"str",
	"comp_str":"comp_str",
	"opt_str":"opt_str",
	"time":"2006-01-02T15:04:05Z",
	"opt_time":"2006-01-02T15:04:05Z",
	"nested":{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42},
	"nested_obj":{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42},
	"opt_nested_obj":{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42},
	"nested_list":[{"nested_str":"nested_str","nested_comp_str":"list_nested_comp_str_1","nested_comp_opt_int":43},{"nested_str":"nested_str","nested_comp_str":"list_nested_comp_str_2","nested_comp_opt_int":44}],
	"map_cust":{"key":["val1","val2"]},
	"map_regular":{"key":[{"nested_str":"nested_str","nested_comp_str":"nested_comp_str","nested_comp_opt_int":42}]},
	"comp_map":{"comp_key":{"nested_comp_str":"nested_comp_str"}},
	"comp_map_list":{"comp_list_key":[{"nested_comp_str":"nested_comp_str"}]}
}`

type nestedMapExample struct {
	SomeStruct customfield.NestedObject[nestedMapStruct] `tfsdk:"some_struct" json:"some_struct,computed_optional"`
}

type nestedMapStruct struct {
	NestedMap map[string]types.Float64 `tfsdk:"nested_map" json:"nested_map,optional"`
}

type nestedMapExampleZero struct {
	SomeStruct customfield.NestedObject[nestedMapStructZero] `tfsdk:"some_struct" json:"some_struct,computed_optional,decode_null_to_zero"`
}

type nestedMapStructZero struct {
	NestedMap map[string]types.Float64 `tfsdk:"nested_map" json:"nested_map,optional,decode_null_to_zero"`
}

type primitiveListExample struct {
	StrList customfield.List[types.String] `tfsdk:"str_list" json:"str_list,computed_optional"`
}

type primitiveListExampleZero struct {
	StrList customfield.List[types.String] `tfsdk:"str_list" json:"str_list,computed_optional,decode_null_to_zero"`
}

var decode_computed_only_tests = map[string]struct {
	buf      string
	starting interface{}
	expected interface{}
}{
	"primitive_list_unchanged": {
		`{}`,
		primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")}),
		},
		primitiveListExample{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")}),
		},
	},
	"primitive_list_unchanged_zero": {
		`{}`,
		primitiveListExampleZero{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")}),
		},
		primitiveListExampleZero{
			StrList: customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c")}),
		},
	},
	"nested_map_unchanged": {
		`{"some_struct": {"nested_map":{"example_key":3.14}}}`,
		nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
		nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
	},
	"nested_map_unchanged_zero": {
		`{"some_struct": {"nested_map":{"example_key":3.14}}}`,
		nestedMapExampleZero{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStructZero{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
		nestedMapExampleZero{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStructZero{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
	},
	"nested_optional_map_avoids_updates": {
		`{"some_struct": {"nested_map":{"example_key":0.123,"new_key":456.7}}}`,
		nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
		nestedMapExample{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStruct{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
	},
	"nested_optional_map_avoids_updates_zero": {
		`{"some_struct": {"nested_map":{"example_key":0.123,"new_key":456.7}}}`,
		nestedMapExampleZero{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStructZero{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
		nestedMapExampleZero{
			SomeStruct: customfield.NewObjectMust(ctx, &nestedMapStructZero{
				NestedMap: map[string]types.Float64{"example_key": types.Float64Value(3.14)},
			}),
		},
	},
	"only_updates_computed_props": {
		exampleNestedJson,
		StructWithComputedFields{
			RegStr:      types.StringNull(),
			CompStr:     types.StringNull(),
			CompOptStr:  types.StringNull(),
			CompTime:    timetypes.NewRFC3339Null(),
			CompOptTime: timetypes.NewRFC3339Null(),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Null(),
			},
			NestedCust:        customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.NullObject[NestedStructWithComputedFields](ctx),
			MapRegular: P(map[string][]*NestedStructWithComputedFields{"key": {
				&NestedStructWithComputedFields{
					RegStr:     types.StringNull(),
					CompStr:    types.StringNull(),
					CompOptInt: types.Int64Null(),
				},
			}}),
		},
		StructWithComputedFields{
			RegStr:      types.StringNull(),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			},
			NestedCust: customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			}),
			MapRegular: P(map[string][]*NestedStructWithComputedFields{"key": {
				&NestedStructWithComputedFields{
					RegStr:     types.StringNull(),
					CompStr:    types.StringValue("nested_comp_str"),
					CompOptInt: types.Int64Value(42),
				},
			}}),
			CompMap: P(map[string]*NestedStructWithComputedFields{"comp_key": {
				CompStr: types.StringValue("nested_comp_str"),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFields{"comp_list_key": {
				&NestedStructWithComputedFields{
					CompStr: types.StringValue("nested_comp_str"),
				},
			}}),
		},
	},
	"only_updates_computed_props_zero": {
		exampleNestedJson,
		StructWithComputedFieldsZero{
			RegStr:      types.StringNull(),
			CompStr:     types.StringNull(),
			CompOptStr:  types.StringNull(),
			CompTime:    timetypes.NewRFC3339Null(),
			CompOptTime: timetypes.NewRFC3339Null(),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringNull(),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Null(),
			},
			NestedCust:        customfield.NullObject[NestedStructWithComputedFieldsZero](ctx),
			CompOptNestedCust: customfield.NullObject[NestedStructWithComputedFieldsZero](ctx),
			MapRegular: P(map[string][]*NestedStructWithComputedFieldsZero{"key": {
				&NestedStructWithComputedFieldsZero{
					RegStr:     types.StringNull(),
					CompStr:    types.StringNull(),
					CompOptInt: types.Int64Null(),
				},
			}}),
		},
		StructWithComputedFieldsZero{
			RegStr:      types.StringNull(),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			},
			NestedCust: customfield.NullObject[NestedStructWithComputedFieldsZero](ctx),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			}),
			MapRegular: P(map[string][]*NestedStructWithComputedFieldsZero{"key": {
				&NestedStructWithComputedFieldsZero{
					RegStr:     types.StringNull(),
					CompStr:    types.StringValue("nested_comp_str"),
					CompOptInt: types.Int64Value(42),
				},
			}}),
			CompMap: P(map[string]*NestedStructWithComputedFieldsZero{"comp_key": {
				RegStr:     types.StringValue(""),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(0),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFieldsZero{"comp_list_key": {
				&NestedStructWithComputedFieldsZero{
					RegStr:     types.StringValue(""),
					CompStr:    types.StringValue("nested_comp_str"),
					CompOptInt: types.Int64Value(0),
				},
			}}),
		},
	},
	"only_updates_computed_props_from_unknown": {
		exampleNestedJson,
		StructWithComputedFields{
			RegStr:      types.StringUnknown(),
			CompStr:     types.StringUnknown(),
			CompOptStr:  types.StringUnknown(),
			CompTime:    timetypes.NewRFC3339Unknown(),
			CompOptTime: timetypes.NewRFC3339Unknown(),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringUnknown(),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Unknown(),
			},
			// when the value is nested and optional/required, we don't currently convert from unknown to null
			// this is because optional/required properties cannot be unknown
			NestedCust:        customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.UnknownObject[NestedStructWithComputedFields](ctx),
		},
		StructWithComputedFields{
			RegStr:      types.StringUnknown(),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringUnknown(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			},
			NestedCust: customfield.NullObject[NestedStructWithComputedFields](ctx),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			}),
			CompMap: P(map[string]*NestedStructWithComputedFields{"comp_key": {
				CompStr: types.StringValue("nested_comp_str"),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFields{"comp_list_key": {
				&NestedStructWithComputedFields{
					CompStr: types.StringValue("nested_comp_str"),
				},
			}}),
		},
	},
	"only_updates_computed_props_from_unknown_zero": {
		exampleNestedJson,
		StructWithComputedFieldsZero{
			RegStr:      types.StringUnknown(),
			CompStr:     types.StringUnknown(),
			CompOptStr:  types.StringUnknown(),
			CompTime:    timetypes.NewRFC3339Unknown(),
			CompOptTime: timetypes.NewRFC3339Unknown(),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringUnknown(),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Unknown(),
			},
			// when the value is nested and optional/required, we don't currently convert from unknown to null
			// this is because optional/required properties cannot be unknown
			NestedCust:        customfield.NullObject[NestedStructWithComputedFieldsZero](ctx),
			CompOptNestedCust: customfield.UnknownObject[NestedStructWithComputedFieldsZero](ctx),
		},
		StructWithComputedFieldsZero{
			RegStr:      types.StringUnknown(),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringUnknown(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			},
			NestedCust: customfield.NullObject[NestedStructWithComputedFieldsZero](ctx),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringNull(),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(42),
			}),
			CompMap: P(map[string]*NestedStructWithComputedFieldsZero{"comp_key": {
				RegStr:     types.StringValue(""),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(0),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFieldsZero{"comp_list_key": {
				&NestedStructWithComputedFieldsZero{
					RegStr:     types.StringValue(""),
					CompStr:    types.StringValue("nested_comp_str"),
					CompOptInt: types.Int64Value(0),
				},
			}}),
		},
	},

	"doesnt_update_computed_optional_if_set": {
		exampleNestedJson,
		StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue("existing_comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_1"),
				CompOptInt: types.Int64Value(11),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
		},
		StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue("list_nested_comp_str_1"),
				CompOptInt: types.Int64Value(11),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
			CompMap: P(map[string]*NestedStructWithComputedFields{"comp_key": {
				CompStr: types.StringValue("nested_comp_str"),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFields{"comp_list_key": {
				&NestedStructWithComputedFields{
					CompStr: types.StringValue("nested_comp_str"),
				},
			}}),
		},
	},
	"doesnt_update_computed_optional_if_set_zero": {
		exampleNestedJson,
		StructWithComputedFieldsZero{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue("existing_comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("existing_nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFieldsZero{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_1"),
				CompOptInt: types.Int64Value(11),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
		},
		StructWithComputedFieldsZero{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue("comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFieldsZero{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue("list_nested_comp_str_1"),
				CompOptInt: types.Int64Value(11),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
			CompMap: P(map[string]*NestedStructWithComputedFieldsZero{"comp_key": {
				RegStr:     types.StringValue(""),
				CompStr:    types.StringValue("nested_comp_str"),
				CompOptInt: types.Int64Value(0),
			}}),
			CompMapList: P(map[string][]*NestedStructWithComputedFieldsZero{"comp_list_key": {
				&NestedStructWithComputedFieldsZero{
					RegStr:     types.StringValue(""),
					CompStr:    types.StringValue("nested_comp_str"),
					CompOptInt: types.Int64Value(0),
				},
			}}),
		},
	},

	"updates_computed_if_JSON_properties_are_missing": {
		`{}`,
		StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue("existing_comp_str"),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringUnknown(),
				CompOptInt: types.Int64Unknown(),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
			MapCust: customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
				"key": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringUnknown(), types.StringValue("val2")}),
			}),
		},
		StructWithComputedFields{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringNull(),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339Null(),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFields{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFields{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Null(),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringNull(),
				CompOptInt: types.Int64Value(12),
			}},
			MapCust: customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
				"key": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringNull(), types.StringValue("val2")}),
			}),
		},
	},
	"updates_computed_if_JSON_properties_are_missing_zero": {
		`{}`,
		StructWithComputedFieldsZero{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue(""),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Time{}),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFieldsZero{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(0),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue("existing_list_nested_comp_str_2"),
				CompOptInt: types.Int64Value(12),
			}},
			MapCust: customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
				"key": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringUnknown(), types.StringValue("val2")}),
			}),
		},
		StructWithComputedFieldsZero{
			RegStr:      types.StringValue("existing_str"),
			CompStr:     types.StringValue(""),
			CompOptStr:  types.StringValue("existing_opt_str"),
			CompTime:    timetypes.NewRFC3339TimeValue(time.Time{}),
			CompOptTime: timetypes.NewRFC3339TimeValue(time.Date(1970, time.January, 2, 15, 4, 5, 0, time.UTC)),
			Nested: NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(10),
			},
			NestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(10),
			}),
			CompOptNestedCust: customfield.NewObjectMust(ctx, &NestedStructWithComputedFieldsZero{
				RegStr:     types.StringValue("existing_nested_str"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(10),
			}),
			NestedList: &[]*NestedStructWithComputedFieldsZero{{
				RegStr:     types.StringValue("existing_list_nested_str_1"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(0),
			}, {
				RegStr:     types.StringValue("existing_list_nested_str_2"),
				CompStr:    types.StringValue(""),
				CompOptInt: types.Int64Value(12),
			}},
			MapCust: customfield.NewMapMust(ctx, map[string]customfield.List[types.String]{
				"key": customfield.NewListMust[types.String](ctx, []attr.Value{types.StringValue(""), types.StringValue("val2")}),
			}),
			CompMap:     P(map[string]*NestedStructWithComputedFieldsZero{}),
			CompMapList: P(map[string][]*NestedStructWithComputedFieldsZero{}),
		},
	},

	"tfsdk_struct_only_overwrites_computed_from_json": {
		`{"embedded_string":"new_value"}`,
		EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		},
		EmbeddedTfsdkStruct{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.NullObject[DoubleNestedStruct](ctx),
		},
	},
	"tfsdk_struct_only_overwrites_computed_from_json_zero": {
		`{"embedded_string":"new_value"}`,
		EmbeddedTfsdkStructZero{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject:     customfield.UnknownObject[DoubleNestedStruct](ctx),
		},
		EmbeddedTfsdkStructZero{
			EmbeddedString: types.StringValue("existing_value"),
			EmbeddedInt:    types.Int64Value(5),
			DataObject: customfield.NewObjectMust(ctx, &DoubleNestedStruct{
				NestedInt: types.Int64Null(),
			}),
		},
	},
}

var test_semantic_equivalence = map[string][]attr.Value{
	"nulls": {
		basetypes.NewBoolNull(),
		basetypes.NewBoolNull(),
		basetypes.NewInt32Null(),
		basetypes.NewMapNull(basetypes.BoolType{}),
		basetypes.NewSetNull(basetypes.StringType{}),
		basetypes.NewListNull(basetypes.NumberType{}),
		basetypes.NewTupleNull([]attr.Type{}),
		basetypes.NewObjectNull(map[string]attr.Type{"hi": basetypes.StringType{}}),
	},
	"unknowns": {
		basetypes.NewBoolUnknown(),
		basetypes.NewBoolUnknown(),
		basetypes.NewInt32Unknown(),
		basetypes.NewMapUnknown(basetypes.BoolType{}),
		basetypes.NewSetUnknown(basetypes.StringType{}),
		basetypes.NewListUnknown(basetypes.NumberType{}),
		basetypes.NewTupleUnknown([]attr.Type{}),
		basetypes.NewObjectUnknown(map[string]attr.Type{"hi": basetypes.StringType{}}),
	},
	"floats": {
		basetypes.NewFloat32Value(12.0),
		basetypes.NewFloat64Value(12.0),
		basetypes.NewNumberValue(big.NewFloat(12.0)),
	},
	"ints": {
		basetypes.NewInt32Value(12),
		basetypes.NewInt64Value(12),
		basetypes.NewNumberValue(big.NewFloat(12)),
	},
	"sequences": {
		basetypes.NewSetValueMust(basetypes.DynamicType{}, []attr.Value{
			basetypes.NewDynamicValue(basetypes.NewInt64Value(12)),
		}),
		basetypes.NewListValueMust(basetypes.DynamicType{}, []attr.Value{
			basetypes.NewDynamicValue(basetypes.NewInt32Value(12)),
		}),
		basetypes.NewTupleValueMust([]attr.Type{customfield.NormalizedDynamicType{}}, []attr.Value{
			customfield.RawNormalizedDynamicValueFrom(basetypes.NewInt64Value(12)),
		}),
	},
	"maps": {
		basetypes.NewMapValueMust(basetypes.DynamicType{}, map[string]attr.Value{
			"12": basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(12.0))),
			"14": basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(14.0))),
		}),
		basetypes.NewObjectValueMust(map[string]attr.Type{"12": basetypes.DynamicType{}, "14": basetypes.DynamicType{}}, map[string]attr.Value{
			"12": basetypes.NewDynamicValue(basetypes.NewInt32Value(12)),
			"14": basetypes.NewDynamicValue(basetypes.NewInt64Value(14)),
		}),
	},
	"nested": {
		basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"inner": basetypes.DynamicType{},
			},
			map[string]attr.Value{
				"inner": basetypes.NewDynamicValue(basetypes.NewListValueMust(basetypes.DynamicType{}, []attr.Value{
					basetypes.NewDynamicValue(basetypes.NewStringValue("hi")),
					basetypes.NewDynamicValue(basetypes.NewStringValue("mom")),
				})),
			},
		),
		basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"inner": basetypes.DynamicType{},
			},
			map[string]attr.Value{
				"inner": basetypes.NewDynamicValue(basetypes.NewListValueMust(customfield.NormalizedDynamicType{}, []attr.Value{
					customfield.RawNormalizedDynamicValueFrom(basetypes.NewStringValue("hi")),
					customfield.RawNormalizedDynamicValueFrom(basetypes.NewStringValue("mom")),
				})),
			},
		),
		basetypes.NewMapValueMust(basetypes.DynamicType{}, map[string]attr.Value{
			"inner": basetypes.NewDynamicValue(basetypes.NewListValueMust(basetypes.DynamicType{}, []attr.Value{
				basetypes.NewDynamicValue(basetypes.NewStringValue("hi")),
				basetypes.NewDynamicValue(basetypes.NewStringValue("mom")),
			})),
		}),
	},
}

func TestDecodeComputedOnly(t *testing.T) {
	spew.Config.ContinueOnMethod = false
	for name, test := range decode_computed_only_tests {
		t.Run(name, func(t *testing.T) {
			v := reflect.ValueOf(test.starting)
			starting := reflect.New(v.Type())
			starting.Elem().Set(v)

			if err := UnmarshalComputed([]byte(test.buf), starting.Interface()); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", test.buf, err)
			}
			startingIFace := starting.Elem().Interface()
			if !reflect.DeepEqual(startingIFace, test.expected) {
				t.Fatalf("expected '%s' to deserialize to \n%s\nbut got\n%s", test.buf, spew.Sdump(test.expected), spew.Sdump(startingIFace))
			}
		})
	}
}

func TestNoStateBetweenDecoders(t *testing.T) {
	// If there is global state between the decoders, these tests will pass individually but fail when run in the same
	// test here. This can happen if our cache key does not capture all the information needed to make these two decoders unique.
	TestDecodeComputedOnly(t)
	TestDecodeFromValue(t)
}

func TestSemanticEquivalence(t *testing.T) {
	ctx := context.TODO()
	for name, values := range test_semantic_equivalence {
		t.Run(name, func(t *testing.T) {
			for i, pair := range pairwise(values) {
				lhs := customfield.RawNormalizedDynamicValueFrom(pair[0])
				rhs := customfield.RawNormalizedDynamicValueFrom(pair[1])

				eq, d := lhs.DynamicSemanticEquals(ctx, rhs)
				if len(d) != 0 {
					t.Fatalf("unexpected Diagnostics: %v", d)
				}
				if !eq {
					t.Fatalf("unexpected inequality index: %d, %v <> %v", i, lhs, rhs)

				}
			}
		})
	}
}

func pairwise[T any](input []T) [][]T {
	pairs := [][]T{}
	if len(input) < 2 {
		return [][]T{input}
	}
	a := input[0]
	for _, b := range input[1:] {
		pairs = append(pairs, []T{a, b})
		a = b
	}
	return pairs
}

func merge[T interface{}](test_array ...map[string]T) map[string]T {
	out := make(map[string]T)
	for _, tests := range test_array {
		for name, t := range tests {
			// panic if there are duplicates because otherwise we'd silently
			// skip some tests
			if _, existing := out[name]; existing {
				// panic(fmt.Sprintf("duplicate test name: %s", name))
				fmt.Printf("duplicate test name: %s", name)
			}
			out[name] = t
		}
	}
	return out
}

func formatJson(jsonString string, out *string) error {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(jsonString), "", "    ")
	if err != nil {
		return err
	}

	*out = prettyJSON.String()
	return nil
}
