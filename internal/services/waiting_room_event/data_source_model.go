// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomEventResultDataSourceEnvelope struct {
	Result WaitingRoomEventDataSourceModel `json:"result,computed"`
}

type WaitingRoomEventResultListDataSourceEnvelope struct {
	Result *[]*WaitingRoomEventDataSourceModel `json:"result,computed"`
}

type WaitingRoomEventDataSourceModel struct {
	ZoneID                types.String                              `tfsdk:"zone_id" path:"zone_id"`
	WaitingRoomID         types.String                              `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	EventID               types.String                              `tfsdk:"event_id" path:"event_id"`
	ID                    types.String                              `tfsdk:"id" json:"id"`
	CreatedOn             types.String                              `tfsdk:"created_on" json:"created_on,computed"`
	CustomPageHTML        types.String                              `tfsdk:"custom_page_html" json:"custom_page_html"`
	Description           types.String                              `tfsdk:"description" json:"description,computed"`
	DisableSessionRenewal types.Bool                                `tfsdk:"disable_session_renewal" json:"disable_session_renewal"`
	EventEndTime          types.String                              `tfsdk:"event_end_time" json:"event_end_time"`
	EventStartTime        types.String                              `tfsdk:"event_start_time" json:"event_start_time"`
	ModifiedOn            types.String                              `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                  types.String                              `tfsdk:"name" json:"name"`
	NewUsersPerMinute     types.Int64                               `tfsdk:"new_users_per_minute" json:"new_users_per_minute"`
	PrequeueStartTime     types.String                              `tfsdk:"prequeue_start_time" json:"prequeue_start_time"`
	QueueingMethod        types.String                              `tfsdk:"queueing_method" json:"queueing_method"`
	SessionDuration       types.Int64                               `tfsdk:"session_duration" json:"session_duration"`
	ShuffleAtEventStart   types.Bool                                `tfsdk:"shuffle_at_event_start" json:"shuffle_at_event_start,computed"`
	Suspended             types.Bool                                `tfsdk:"suspended" json:"suspended,computed"`
	TotalActiveUsers      types.Int64                               `tfsdk:"total_active_users" json:"total_active_users"`
	FindOneBy             *WaitingRoomEventFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type WaitingRoomEventFindOneByDataSourceModel struct {
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	Page          types.String `tfsdk:"page" query:"page"`
	PerPage       types.String `tfsdk:"per_page" query:"per_page"`
}
