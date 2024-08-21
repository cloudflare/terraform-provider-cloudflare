// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultDataSourceEnvelope struct {
	Result WaitingRoomRulesDataSourceModel `json:"result,computed"`
}

type WaitingRoomRulesDataSourceModel struct {
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
}

func (m *WaitingRoomRulesDataSourceModel) toReadParams() (params waiting_rooms.RuleGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
