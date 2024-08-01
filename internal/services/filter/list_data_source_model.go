// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FiltersResultListDataSourceEnvelope struct {
	Result *[]*FiltersResultDataSourceModel `json:"result,computed"`
}

type FiltersDataSourceModel struct {
	ZoneIdentifier types.String                     `tfsdk:"zone_identifier" path:"zone_identifier"`
	Description    types.String                     `tfsdk:"description" query:"description"`
	Expression     types.String                     `tfsdk:"expression" query:"expression"`
	ID             types.String                     `tfsdk:"id" query:"id"`
	Paused         types.Bool                       `tfsdk:"paused" query:"paused"`
	Ref            types.String                     `tfsdk:"ref" query:"ref"`
	Page           types.Float64                    `tfsdk:"page" query:"page"`
	PerPage        types.Float64                    `tfsdk:"per_page" query:"per_page"`
	MaxItems       types.Int64                      `tfsdk:"max_items"`
	Result         *[]*FiltersResultDataSourceModel `tfsdk:"result"`
}

type FiltersResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Description types.String `tfsdk:"description" json:"description"`
	Ref         types.String `tfsdk:"ref" json:"ref"`
}
