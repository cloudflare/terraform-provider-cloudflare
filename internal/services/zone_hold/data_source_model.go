// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zones"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneHoldResultDataSourceEnvelope struct {
	Result ZoneHoldDataSourceModel `json:"result,computed"`
}

type ZoneHoldDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *ZoneHoldDataSourceModel) toReadParams(_ context.Context) (params zones.HoldGetParams, diags diag.Diagnostics) {
	params = zones.HoldGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
