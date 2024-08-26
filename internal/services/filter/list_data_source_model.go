// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/filters"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FiltersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[FiltersResultDataSourceModel] `json:"result,computed"`
}

type FiltersDataSourceModel struct {
	ZoneIdentifier types.String                                               `tfsdk:"zone_identifier" path:"zone_identifier"`
	Description    types.String                                               `tfsdk:"description" query:"description"`
	Expression     types.String                                               `tfsdk:"expression" query:"expression"`
	ID             types.String                                               `tfsdk:"id" query:"id"`
	Paused         types.Bool                                                 `tfsdk:"paused" query:"paused"`
	Ref            types.String                                               `tfsdk:"ref" query:"ref"`
	MaxItems       types.Int64                                                `tfsdk:"max_items"`
	Result         customfield.NestedObjectList[FiltersResultDataSourceModel] `tfsdk:"result"`
}

func (m *FiltersDataSourceModel) toListParams() (params filters.FilterListParams, diags diag.Diagnostics) {
	params = filters.FilterListParams{}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Description.IsNull() {
		params.Description = cloudflare.F(m.Description.ValueString())
	}
	if !m.Expression.IsNull() {
		params.Expression = cloudflare.F(m.Expression.ValueString())
	}
	if !m.Paused.IsNull() {
		params.Paused = cloudflare.F(m.Paused.ValueBool())
	}
	if !m.Ref.IsNull() {
		params.Ref = cloudflare.F(m.Ref.ValueString())
	}

	return
}

type FiltersResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Description types.String `tfsdk:"description" json:"description,computed_optional"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed_optional"`
}
