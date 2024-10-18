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
	FilterID    types.String                    `tfsdk:"filter_id" path:"filter_id,optional"`
	ZoneID      types.String                    `tfsdk:"zone_id" path:"zone_id,optional"`
	Description types.String                    `tfsdk:"description" json:"description,computed"`
	Expression  types.String                    `tfsdk:"expression" json:"expression,computed"`
	ID          types.String                    `tfsdk:"id" json:"id,computed"`
	Paused      types.Bool                      `tfsdk:"paused" json:"paused,computed"`
	Ref         types.String                    `tfsdk:"ref" json:"ref,computed"`
	Filter      *FilterFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *FilterDataSourceModel) toReadParams(_ context.Context) (params filters.FilterGetParams, diags diag.Diagnostics) {
	params = filters.FilterGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *FilterDataSourceModel) toListParams(_ context.Context) (params filters.FilterListParams, diags diag.Diagnostics) {
	params = filters.FilterListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

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
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	ID          types.String `tfsdk:"id" query:"id,optional"`
	Description types.String `tfsdk:"description" query:"description,optional"`
	Expression  types.String `tfsdk:"expression" query:"expression,optional"`
	Paused      types.Bool   `tfsdk:"paused" query:"paused,optional"`
	Ref         types.String `tfsdk:"ref" query:"ref,optional"`
}
