// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomEventResultDataSourceEnvelope struct {
	Result WaitingRoomEventDataSourceModel `json:"result,computed"`
}

type WaitingRoomEventResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WaitingRoomEventDataSourceModel] `json:"result,computed"`
}

type WaitingRoomEventDataSourceModel struct {
	EventID               types.String                              `tfsdk:"event_id" path:"event_id"`
	WaitingRoomID         types.String                              `tfsdk:"waiting_room_id" path:"waiting_room_id"`
	ZoneID                types.String                              `tfsdk:"zone_id" path:"zone_id"`
	CreatedOn             timetypes.RFC3339                         `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                              `tfsdk:"description" json:"description,computed"`
	ModifiedOn            timetypes.RFC3339                         `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ShuffleAtEventStart   types.Bool                                `tfsdk:"shuffle_at_event_start" json:"shuffle_at_event_start,computed"`
	Suspended             types.Bool                                `tfsdk:"suspended" json:"suspended,computed"`
	CustomPageHTML        types.String                              `tfsdk:"custom_page_html" json:"custom_page_html,computed_optional"`
	DisableSessionRenewal types.Bool                                `tfsdk:"disable_session_renewal" json:"disable_session_renewal,computed_optional"`
	EventEndTime          types.String                              `tfsdk:"event_end_time" json:"event_end_time,computed_optional"`
	EventStartTime        types.String                              `tfsdk:"event_start_time" json:"event_start_time,computed_optional"`
	ID                    types.String                              `tfsdk:"id" json:"id,computed_optional"`
	Name                  types.String                              `tfsdk:"name" json:"name,computed_optional"`
	NewUsersPerMinute     types.Int64                               `tfsdk:"new_users_per_minute" json:"new_users_per_minute,computed_optional"`
	PrequeueStartTime     types.String                              `tfsdk:"prequeue_start_time" json:"prequeue_start_time,computed_optional"`
	QueueingMethod        types.String                              `tfsdk:"queueing_method" json:"queueing_method,computed_optional"`
	SessionDuration       types.Int64                               `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	TotalActiveUsers      types.Int64                               `tfsdk:"total_active_users" json:"total_active_users,computed_optional"`
	Filter                *WaitingRoomEventFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WaitingRoomEventDataSourceModel) toReadParams() (params waiting_rooms.EventGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.EventGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *WaitingRoomEventDataSourceModel) toListParams() (params waiting_rooms.EventListParams, diags diag.Diagnostics) {
	params = waiting_rooms.EventListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type WaitingRoomEventFindOneByDataSourceModel struct {
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
	WaitingRoomID types.String `tfsdk:"waiting_room_id" path:"waiting_room_id"`
}
