// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/bot_management"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultDataSourceEnvelope struct {
	Result BotManagementDataSourceModel `json:"result,computed"`
}

type BotManagementDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}

func (m *BotManagementDataSourceModel) toReadParams() (params bot_management.BotManagementGetParams, diags diag.Diagnostics) {
	params = bot_management.BotManagementGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
