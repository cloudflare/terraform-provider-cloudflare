// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FiltersResultListDataSourceEnvelope struct {
	Result *[]*FiltersItemsDataSourceModel `json:"result,computed"`
}

type FiltersDataSourceModel struct {
	ZoneIdentifier types.String                    `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                    `tfsdk:"id" query:"id"`
	Description    types.String                    `tfsdk:"description" query:"description"`
	Expression     types.String                    `tfsdk:"expression" query:"expression"`
	Page           types.Float64                   `tfsdk:"page" query:"page"`
	Paused         types.Bool                      `tfsdk:"paused" query:"paused"`
	PerPage        types.Float64                   `tfsdk:"per_page" query:"per_page"`
	Ref            types.String                    `tfsdk:"ref" query:"ref"`
	MaxItems       types.Int64                     `tfsdk:"max_items"`
	Items          *[]*FiltersItemsDataSourceModel `tfsdk:"items"`
}

type FiltersItemsDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Description types.String `tfsdk:"description" json:"description"`
	Ref         types.String `tfsdk:"ref" json:"ref"`
}
