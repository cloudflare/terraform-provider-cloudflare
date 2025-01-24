// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/filters"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterResultDataSourceEnvelope struct {
	Result FilterDataSourceModel `json:"result,computed"`
}

type FilterDataSourceModel struct {
	FilterID    types.String `tfsdk:"filter_id" path:"filter_id,required"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Expression  types.String `tfsdk:"expression" json:"expression,computed"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed"`
}

func (m *FilterDataSourceModel) toReadParams(_ context.Context) (params filters.FilterGetParams, diags diag.Diagnostics) {
	params = filters.FilterGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
