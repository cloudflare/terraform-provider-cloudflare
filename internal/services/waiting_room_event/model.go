// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomEventResultEnvelope struct {
	Result WaitingRoomEventModel `json:"result"`
}

type WaitingRoomEventModel struct {
	ID                    types.String      `tfsdk:"id" json:"id,computed"`
	WaitingRoomID         types.String      `tfsdk:"waiting_room_id" path:"waiting_room_id,required"`
	ZoneID                types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	EventEndTime          types.String      `tfsdk:"event_end_time" json:"event_end_time,required"`
	EventStartTime        types.String      `tfsdk:"event_start_time" json:"event_start_time,required"`
	Name                  types.String      `tfsdk:"name" json:"name,required"`
	CustomPageHTML        types.String      `tfsdk:"custom_page_html" json:"custom_page_html,optional"`
	DisableSessionRenewal types.Bool        `tfsdk:"disable_session_renewal" json:"disable_session_renewal,optional"`
	NewUsersPerMinute     types.Int64       `tfsdk:"new_users_per_minute" json:"new_users_per_minute,optional"`
	PrequeueStartTime     types.String      `tfsdk:"prequeue_start_time" json:"prequeue_start_time,optional"`
	QueueingMethod        types.String      `tfsdk:"queueing_method" json:"queueing_method,optional"`
	SessionDuration       types.Int64       `tfsdk:"session_duration" json:"session_duration,optional"`
	TotalActiveUsers      types.Int64       `tfsdk:"total_active_users" json:"total_active_users,optional"`
	Description           types.String      `tfsdk:"description" json:"description,computed_optional"`
	ShuffleAtEventStart   types.Bool        `tfsdk:"shuffle_at_event_start" json:"shuffle_at_event_start,computed_optional"`
	Suspended             types.Bool        `tfsdk:"suspended" json:"suspended,computed_optional"`
	CreatedOn             timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn            timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}
