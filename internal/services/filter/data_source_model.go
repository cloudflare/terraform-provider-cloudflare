// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/filters"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterResultDataSourceEnvelope struct {
	Result FilterDataSourceModel `json:"result,computed"`
}

type FilterResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[FilterDataSourceModel] `json:"result,computed"`
}

type FilterDataSourceModel struct {
	ZoneIdentifier types.String                    `tfsdk:"zone_identifier" path:"zone_identifier,optional"`
	ID             types.String                    `tfsdk:"id" path:"id,computed_optional"`
	Description    types.String                    `tfsdk:"description" json:"description,computed"`
	Expression     types.String                    `tfsdk:"expression" json:"expression,computed"`
	Paused         types.Bool                      `tfsdk:"paused" json:"paused,computed"`
	Ref            types.String                    `tfsdk:"ref" json:"ref,computed"`
	Filter         *FilterFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *FilterDataSourceModel) toListParams(_ context.Context) (params filters.FilterListParams, diags diag.Diagnostics) {
	params = filters.FilterListParams{}

	if !m.Filter.ID.IsNull() {
		params.ID = cloudflare.F(m.Filter.ID.ValueString())
	}
	if !m.Filter.Description.IsNull() {
		params.Description = cloudflare.F(m.Filter.Description.ValueString())
	}
	if !m.Filter.Expression.IsNull() {
		params.Expression = cloudflare.F(m.Filter.Expression.ValueString())
	}
	if !m.Filter.Paused.IsNull() {
		params.Paused = cloudflare.F(m.Filter.Paused.ValueBool())
	}
	if !m.Filter.Ref.IsNull() {
		params.Ref = cloudflare.F(m.Filter.Ref.ValueString())
	}

	return
}

type FilterFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier,required"`
	ID             types.String `tfsdk:"id" query:"id,optional"`
	Description    types.String `tfsdk:"description" query:"description,optional"`
	Expression     types.String `tfsdk:"expression" query:"expression,optional"`
	Paused         types.Bool   `tfsdk:"paused" query:"paused,optional"`
	Ref            types.String `tfsdk:"ref" query:"ref,optional"`
}
