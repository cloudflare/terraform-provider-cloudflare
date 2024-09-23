// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/waiting_rooms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultDataSourceEnvelope struct {
	Result WaitingRoomRulesDataSourceModel `json:"result,computed"`
}

type WaitingRoomRulesDataSourceModel struct {
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *WaitingRoomRulesDataSourceModel) toReadParams(_ context.Context) (params waiting_rooms.RuleGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
