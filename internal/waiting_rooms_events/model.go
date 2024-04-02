// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_events

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomsEventsResultEnvelope struct {
	Result WaitingRoomsEventsModel `json:"result,computed"`
}

type WaitingRoomsEventsModel struct {
	ZoneID                types.String `tfsdk:"zone_id" path:"zone_id"`
	WaitingRoomID         types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	EventID               types.String `tfsdk:"event_id" path:"event_id"`
	EventEndTime          types.String `tfsdk:"event_end_time" json:"event_end_time"`
	EventStartTime        types.String `tfsdk:"event_start_time" json:"event_start_time"`
	Name                  types.String `tfsdk:"name" json:"name"`
	CustomPageHTML        types.String `tfsdk:"custom_page_html" json:"custom_page_html"`
	Description           types.String `tfsdk:"description" json:"description"`
	DisableSessionRenewal types.Bool   `tfsdk:"disable_session_renewal" json:"disable_session_renewal"`
	NewUsersPerMinute     types.Int64  `tfsdk:"new_users_per_minute" json:"new_users_per_minute"`
	PrequeueStartTime     types.String `tfsdk:"prequeue_start_time" json:"prequeue_start_time"`
	QueueingMethod        types.String `tfsdk:"queueing_method" json:"queueing_method"`
	SessionDuration       types.Int64  `tfsdk:"session_duration" json:"session_duration"`
	ShuffleAtEventStart   types.Bool   `tfsdk:"shuffle_at_event_start" json:"shuffle_at_event_start"`
	Suspended             types.Bool   `tfsdk:"suspended" json:"suspended"`
	TotalActiveUsers      types.Int64  `tfsdk:"total_active_users" json:"total_active_users"`
	ID                    types.String `tfsdk:"id" json:"id"`
	CreatedOn             types.String `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn            types.String `tfsdk:"modified_on" json:"modified_on,computed"`
}
