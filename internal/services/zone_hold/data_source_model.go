// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zones"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneHoldResultDataSourceEnvelope struct {
	Result ZoneHoldDataSourceModel `json:"result,computed"`
}

type ZoneHoldDataSourceModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	Hold              types.Bool   `tfsdk:"hold" json:"hold"`
	HoldAfter         types.String `tfsdk:"hold_after" json:"hold_after"`
	IncludeSubdomains types.String `tfsdk:"include_subdomains" json:"include_subdomains"`
}

func (m *ZoneHoldDataSourceModel) toReadParams() (params zones.HoldGetParams, diags diag.Diagnostics) {
	params = zones.HoldGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
