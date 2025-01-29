package test_helpers_test

import (
	"context"
	"testing"

	ds "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rs "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

type empty struct{}

type builtin struct {
	A bool                   `tfsdk:"A" path:"required"`
	B *int                   `tfsdk:"B" query:"computed_optional"`
	C **int64                `tfsdk:"C" json:"computed"`
	D ***float64             `tfsdk:"D" path:"optional"`
	E ****float64            `tfsdk:"E" query:"required"`
	F *****string            `tfsdk:"F" json:"computed_optional"`
	G []bool                 `tfsdk:"G" query:"computed"`
	H *[]*bool               `tfsdk:"H" json:"optional"`
	I map[**string]**bool    `tfsdk:"I" path:"required"`
	J ***builtin             `tfsdk:"J" query:"computed_optional"`
	K []builtin              `tfsdk:"K" json:"computed"`
	L *[]*builtin            `tfsdk:"L" path:"optional"`
	M **map[string]**builtin `tfsdk:"M"`
}

type reallysimple struct {
	A bool `tfsdk:"A"`
	B int  `tfsdk:"B"`
}

type simple struct {
	A bool                 `tfsdk:"A"`
	B *basetypes.BoolValue `tfsdk:"B"`
}

type bstype struct {
	A *basetypes.DynamicValue  `tfsdk:"A"`
	B **basetypes.BoolValue    `tfsdk:"B"`
	C ***basetypes.Int64Value  `tfsdk:"C"`
	D **basetypes.Float64Value `tfsdk:"D"`
	E *basetypes.NumberValue   `tfsdk:"E"`
	F basetypes.StringValue    `tfsdk:"F"`
	G basetypes.ListValue      `tfsdk:"G"`
	H basetypes.SetValue       `tfsdk:"H"`
	I basetypes.MapValue       `tfsdk:"I"`
	J basetypes.ObjectValue    `tfsdk:"J"`
	K simple                   `tfsdk:"K"`
	L *[]simple                `tfsdk:"L"`
	M []*simple                `tfsdk:"M"`
	N *map[string]*simple      `tfsdk:"N"`
}

type custype struct {
	A customfield.List[basetypes.BoolValue]                 `tfsdk:"A"`
	B customfield.Set[basetypes.BoolValue]                  `tfsdk:"B"`
	C customfield.Map[basetypes.BoolValue]                  `tfsdk:"C"`
	D customfield.Map[customfield.Map[basetypes.BoolValue]] `tfsdk:"D"`
	E customfield.NestedObject[reallysimple]                `tfsdk:"E"`
	F customfield.NestedObjectList[reallysimple]            `tfsdk:"F"`
	G customfield.NestedObjectSet[reallysimple]             `tfsdk:"G"`
	H customfield.NestedObjectMap[reallysimple]             `tfsdk:"H"`
}

var ctx = context.TODO()

var datasourceTests = map[string]struct {
	model  any
	schema ds.Schema
	errors []string
}{
	"empty": {
		(*empty)(nil),
		ds.Schema{},
		[]string{},
	},
	"builtin": {
		(*builtin)(nil),
		ds.Schema{
			Attributes: map[string]ds.Attribute{
				"A": rs.BoolAttribute{
					Required: true,
				},
				"B": rs.Int64Attribute{
					Computed: true,
					Optional: true,
				},
				"C": rs.Int64Attribute{
					Computed: true,
				},
				"D": rs.Float64Attribute{
					Optional: true,
				},
				"E": rs.NumberAttribute{
					Required: true,
				},
				"F": rs.StringAttribute{
					Computed: true,
					Optional: true,
				},
				"G": rs.ListAttribute{
					Computed:    true,
					ElementType: basetypes.BoolType{},
				},
				"H": rs.SetAttribute{
					Optional:    true,
					ElementType: basetypes.BoolType{},
				},
				"I": rs.MapAttribute{
					Required:    true,
					ElementType: basetypes.BoolType{},
				},
				"J": rs.SingleNestedAttribute{
					Computed: true,
					Optional: true,
					Attributes: map[string]rs.Attribute{
						"A": ds.BoolAttribute{
							Required: true,
						},
						"B": ds.Int64Attribute{
							Computed: true,
							Optional: true,
						},
						"C": ds.Int64Attribute{
							Computed: true,
						},
					},
				},
				"K": rs.ListNestedAttribute{
					Computed: true,
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{},
					},
				},
				"L": rs.SetNestedAttribute{
					Optional: true,
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"D": ds.Float64Attribute{
								Optional: true,
							},
							"E": ds.NumberAttribute{
								Required: true,
							},
							"F": ds.StringAttribute{
								Computed: true,
								Optional: true,
							},
						},
					},
				},
				"M": rs.MapNestedAttribute{
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{},
					},
				},
			},
		},
		[]string{},
	},
	"basetypes": {
		(*bstype)(nil),
		ds.Schema{
			Attributes: map[string]ds.Attribute{
				"A": rs.DynamicAttribute{},
				"B": rs.BoolAttribute{},
				"C": rs.Int64Attribute{},
				"D": rs.Float64Attribute{},
				"E": rs.NumberAttribute{},
				"F": rs.StringAttribute{},
				"G": rs.ListAttribute{ElementType: basetypes.BoolType{}},
				"H": rs.SetAttribute{ElementType: basetypes.BoolType{}},
				"I": rs.MapAttribute{ElementType: basetypes.BoolType{}},
				"K": rs.SingleNestedAttribute{
					Attributes: map[string]rs.Attribute{
						"A": ds.BoolAttribute{},
						"B": ds.BoolAttribute{},
					},
				},
				"L": rs.ListNestedAttribute{
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.BoolAttribute{},
						},
					},
				},
				"M": rs.SetNestedAttribute{
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.BoolAttribute{},
						},
					},
				},
				"N": rs.MapNestedAttribute{
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.BoolAttribute{},
						},
					},
				},
			},
		},
		[]string{
			".@bstype.G",
			".@bstype.H",
			".@bstype.I",
		},
	},
	"customfield": {
		(*custype)(nil),
		ds.Schema{
			Attributes: map[string]ds.Attribute{
				"A": ds.ListAttribute{
					ElementType: basetypes.BoolType{},
				},
				"B": ds.SetAttribute{
					ElementType: basetypes.BoolType{},
				},
				"C": ds.MapAttribute{
					ElementType: basetypes.BoolType{},
				},
				"E": ds.SingleNestedAttribute{
					CustomType: customfield.NewNestedObjectType[reallysimple](ctx),
					Attributes: map[string]ds.Attribute{
						"A": ds.BoolAttribute{},
						"B": ds.Int64Attribute{},
					},
				},
				"F": ds.ListNestedAttribute{
					CustomType: customfield.NewNestedObjectListType[reallysimple](ctx),
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.Int64Attribute{},
						},
					},
				},
				"G": ds.SetNestedAttribute{
					CustomType: customfield.NewNestedObjectSetType[reallysimple](ctx),
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.Int64Attribute{},
						},
					},
				},
				"H": ds.MapNestedAttribute{
					CustomType: customfield.NewNestedObjectMapType[reallysimple](ctx),
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.Int64Attribute{},
						},
					},
				},
			},
		},
		[]string{},
	},
}

var resourceTests = map[string]struct {
	model  any
	schema rs.Schema
	errors []string
}{
	"empty": {
		(*empty)(nil),
		rs.Schema{},
		[]string{},
	},
	"builtin": {
		(*builtin)(nil),
		rs.Schema{
			Attributes: map[string]rs.Attribute{
				"A": ds.BoolAttribute{
					Required: true,
				},
				"B": ds.Int64Attribute{
					Computed: true,
					Optional: true,
				},
				"C": ds.Int64Attribute{
					Computed: true,
				},
				"D": ds.Float64Attribute{
					Optional: true,
				},
				"E": ds.NumberAttribute{
					Required: true,
				},
				"F": ds.StringAttribute{
					Computed: true,
					Optional: true,
				},
				"G": ds.ListAttribute{
					Computed:    true,
					ElementType: basetypes.BoolType{},
				},
				"H": ds.SetAttribute{
					Optional:    true,
					ElementType: basetypes.BoolType{},
				},
				"I": ds.MapAttribute{
					Required:    true,
					ElementType: basetypes.BoolType{},
				},
				"J": ds.SingleNestedAttribute{
					Computed: true,
					Optional: true,
					Attributes: map[string]ds.Attribute{
						"A": rs.BoolAttribute{
							Required: true,
						},
						"B": rs.Int64Attribute{
							Computed: true,
							Optional: true,
						},
						"C": rs.Int64Attribute{
							Computed: true,
						},
					},
				},
				"K": ds.ListNestedAttribute{
					Computed: true,
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{},
					},
				},
				"L": ds.SetNestedAttribute{
					Optional: true,
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"D": rs.Float64Attribute{
								Optional: true,
							},
							"E": rs.NumberAttribute{
								Required: true,
							},
							"F": rs.StringAttribute{
								Computed: true,
								Optional: true,
							},
						},
					},
				},
				"M": ds.MapNestedAttribute{
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{},
					},
				},
			},
		},
		[]string{},
	},
	"basetypes": {
		(*bstype)(nil),
		rs.Schema{
			Attributes: map[string]rs.Attribute{
				"A": ds.DynamicAttribute{},
				"B": ds.BoolAttribute{},
				"C": ds.Int64Attribute{},
				"D": ds.Float64Attribute{},
				"E": ds.NumberAttribute{},
				"F": ds.StringAttribute{},
				"G": ds.ListAttribute{ElementType: basetypes.BoolType{}},
				"H": ds.SetAttribute{ElementType: basetypes.BoolType{}},
				"I": ds.MapAttribute{ElementType: basetypes.BoolType{}},
				"K": ds.SingleNestedAttribute{
					Attributes: map[string]ds.Attribute{
						"A": rs.BoolAttribute{},
						"B": rs.BoolAttribute{},
					},
				},
				"L": ds.ListNestedAttribute{
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"A": rs.BoolAttribute{},
							"B": rs.BoolAttribute{},
						},
					},
				},
				"M": ds.SetNestedAttribute{
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"A": rs.BoolAttribute{},
							"B": rs.BoolAttribute{},
						},
					},
				},
				"N": ds.MapNestedAttribute{
					NestedObject: ds.NestedAttributeObject{
						Attributes: map[string]ds.Attribute{
							"A": rs.BoolAttribute{},
							"B": rs.BoolAttribute{},
						},
					},
				},
			},
		},
		[]string{
			".@bstype.G",
			".@bstype.H",
			".@bstype.I",
		},
	},
	"customfield": {
		(*custype)(nil),
		rs.Schema{
			Attributes: map[string]rs.Attribute{
				"A": rs.ListAttribute{
					ElementType: basetypes.BoolType{},
				},
				"B": rs.SetAttribute{
					ElementType: basetypes.BoolType{},
				},
				"C": rs.MapAttribute{
					ElementType: basetypes.BoolType{},
				},
				"E": rs.SingleNestedAttribute{
					CustomType: customfield.NewNestedObjectType[reallysimple](ctx),
					Attributes: map[string]rs.Attribute{
						"A": ds.BoolAttribute{},
						"B": ds.Int64Attribute{},
					},
				},
				"F": rs.ListNestedAttribute{
					CustomType: customfield.NewNestedObjectListType[reallysimple](ctx),
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.Int64Attribute{},
						},
					},
				},
				"G": rs.SetNestedAttribute{
					CustomType: customfield.NewNestedObjectSetType[reallysimple](ctx),
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.Int64Attribute{},
						},
					},
				},
				"H": rs.MapNestedAttribute{
					CustomType: customfield.NewNestedObjectMapType[reallysimple](ctx),
					NestedObject: rs.NestedAttributeObject{
						Attributes: map[string]rs.Attribute{
							"A": ds.BoolAttribute{},
							"B": ds.Int64Attribute{},
						},
					},
				},
			},
		},
		[]string{},
	},
}

func TestParity(t *testing.T) {
	t.Parallel()

	for name, testcase := range datasourceTests {
		t.Run("datasource_"+name, func(t *testing.T) {
			errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(testcase.model, testcase.schema)
			errs.Ignore(t, testcase.errors...)
			errs.Report(t)
		})
	}

	for name, testcase := range resourceTests {
		t.Run("resource_"+name, func(t *testing.T) {
			errs := test_helpers.ValidateResourceModelSchemaIntegrity(testcase.model, testcase.schema)
			errs.Ignore(t, testcase.errors...)
			errs.Report(t)
		})
	}
}
