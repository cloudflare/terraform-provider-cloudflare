// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersRoutesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersRoutesResultDataSourceModel] `json:"result,computed"`
}

type WorkersRoutesDataSourceModel struct {
	ZoneID   types.String                                                     `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[WorkersRoutesResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersRoutesDataSourceModel) toListParams(_ context.Context) (params workers.RouteListParams, diags diag.Diagnostics) {
	params = workers.RouteListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type WorkersRoutesResultDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Pattern types.String `tfsdk:"pattern" json:"pattern,computed"`
	Script  types.String `tfsdk:"script" json:"script,computed"`
}
