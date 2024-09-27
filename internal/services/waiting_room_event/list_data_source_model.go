// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomEventsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WaitingRoomEventsResultDataSourceModel] `json:"result,computed"`
}

type WaitingRoomEventsDataSourceModel struct {
	WaitingRoomID types.String                                                         `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
	ZoneID        types.String                                                         `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems      types.Int64                                                          `tfsdk:"max_items"`
	Result        customfield.NestedObjectList[WaitingRoomEventsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WaitingRoomEventsDataSourceModel) toListParams(_ context.Context) (params waiting_rooms.EventListParams, diags diag.Diagnostics) {
	params = waiting_rooms.EventListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type WaitingRoomEventsResultDataSourceModel struct {
	ID                    types.String      `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomPageHTML        types.String      `tfsdk:"custom_page_html" json:"custom_page_html,computed"`
	Description           types.String      `tfsdk:"description" json:"description,computed"`
	DisableSessionRenewal types.Bool        `tfsdk:"disable_session_renewal" json:"disable_session_renewal,computed"`
	EventEndTime          types.String      `tfsdk:"event_end_time" json:"event_end_time,computed"`
	EventStartTime        types.String      `tfsdk:"event_start_time" json:"event_start_time,computed"`
	ModifiedOn            timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name                  types.String      `tfsdk:"name" json:"name,computed"`
	NewUsersPerMinute     types.Int64       `tfsdk:"new_users_per_minute" json:"new_users_per_minute,computed"`
	PrequeueStartTime     types.String      `tfsdk:"prequeue_start_time" json:"prequeue_start_time,computed"`
	QueueingMethod        types.String      `tfsdk:"queueing_method" json:"queueing_method,computed"`
	SessionDuration       types.Int64       `tfsdk:"session_duration" json:"session_duration,computed"`
	ShuffleAtEventStart   types.Bool        `tfsdk:"shuffle_at_event_start" json:"shuffle_at_event_start,computed"`
	Suspended             types.Bool        `tfsdk:"suspended" json:"suspended,computed"`
	TotalActiveUsers      types.Int64       `tfsdk:"total_active_users" json:"total_active_users,computed"`
}
