// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterResultEnvelope struct {
	Result FilterModel `json:"result,computed"`
}

type FilterResultDataSourceEnvelope struct {
	Result FilterDataSourceModel `json:"result,computed"`
}

type FiltersResultDataSourceEnvelope struct {
	Result FiltersDataSourceModel `json:"result,computed"`
}

type FilterModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String `tfsdk:"id" path:"id"`
	Description    types.String `tfsdk:"description" json:"description,computed"`
	Expression     types.String `tfsdk:"expression" json:"expression,computed"`
	Paused         types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Ref            types.String `tfsdk:"ref" json:"ref,computed"`
}

type FilterDataSourceModel struct {
}

type FiltersDataSourceModel struct {
}
