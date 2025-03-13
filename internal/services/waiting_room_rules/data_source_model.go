// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/waiting_rooms"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomRulesResultDataSourceEnvelope struct {
	Result WaitingRoomRulesDataSourceModel `json:"result,computed"`
}

type WaitingRoomRulesDataSourceModel struct {
	WaitingRoomID types.String      `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
	ZoneID        types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Action        types.String      `tfsdk:"action" json:"action,computed"`
	Description   types.String      `tfsdk:"description" json:"description,computed"`
	Enabled       types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Expression    types.String      `tfsdk:"expression" json:"expression,computed"`
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	LastUpdated   timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version       types.String      `tfsdk:"version" json:"version,computed"`
}

func (m *WaitingRoomRulesDataSourceModel) toReadParams(_ context.Context) (params waiting_rooms.RuleGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
