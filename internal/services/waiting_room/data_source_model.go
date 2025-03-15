// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/waiting_rooms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomResultDataSourceEnvelope struct {
	Result WaitingRoomDataSourceModel `json:"result,computed"`
}

type WaitingRoomDataSourceModel struct {
	ID                         types.String                                                             `tfsdk:"id" json:"-,computed"`
	WaitingRoomID              types.String                                                             `tfsdk:"waiting_room_id" path:"waiting_room_id,optional"`
	ZoneID                     types.String                                                             `tfsdk:"zone_id" path:"zone_id,required"`
	CookieSuffix               types.String                                                             `tfsdk:"cookie_suffix" json:"cookie_suffix,computed"`
	CreatedOn                  timetypes.RFC3339                                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomPageHTML             types.String                                                             `tfsdk:"custom_page_html" json:"custom_page_html,computed"`
	DefaultTemplateLanguage    types.String                                                             `tfsdk:"default_template_language" json:"default_template_language,computed"`
	Description                types.String                                                             `tfsdk:"description" json:"description,computed"`
	DisableSessionRenewal      types.Bool                                                               `tfsdk:"disable_session_renewal" json:"disable_session_renewal,computed"`
	Host                       types.String                                                             `tfsdk:"host" json:"host,computed"`
	JsonResponseEnabled        types.Bool                                                               `tfsdk:"json_response_enabled" json:"json_response_enabled,computed"`
	ModifiedOn                 timetypes.RFC3339                                                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name                       types.String                                                             `tfsdk:"name" json:"name,computed"`
	NewUsersPerMinute          types.Int64                                                              `tfsdk:"new_users_per_minute" json:"new_users_per_minute,computed"`
	NextEventPrequeueStartTime types.String                                                             `tfsdk:"next_event_prequeue_start_time" json:"next_event_prequeue_start_time,computed"`
	NextEventStartTime         types.String                                                             `tfsdk:"next_event_start_time" json:"next_event_start_time,computed"`
	Path                       types.String                                                             `tfsdk:"path" json:"path,computed"`
	QueueAll                   types.Bool                                                               `tfsdk:"queue_all" json:"queue_all,computed"`
	QueueingMethod             types.String                                                             `tfsdk:"queueing_method" json:"queueing_method,computed"`
	QueueingStatusCode         types.Int64                                                              `tfsdk:"queueing_status_code" json:"queueing_status_code,computed"`
	SessionDuration            types.Int64                                                              `tfsdk:"session_duration" json:"session_duration,computed"`
	Suspended                  types.Bool                                                               `tfsdk:"suspended" json:"suspended,computed"`
	TotalActiveUsers           types.Int64                                                              `tfsdk:"total_active_users" json:"total_active_users,computed"`
	TurnstileAction            types.String                                                             `tfsdk:"turnstile_action" json:"turnstile_action,computed"`
	TurnstileMode              types.String                                                             `tfsdk:"turnstile_mode" json:"turnstile_mode,computed"`
	EnabledOriginCommands      customfield.List[types.String]                                           `tfsdk:"enabled_origin_commands" json:"enabled_origin_commands,computed"`
	AdditionalRoutes           customfield.NestedObjectList[WaitingRoomAdditionalRoutesDataSourceModel] `tfsdk:"additional_routes" json:"additional_routes,computed"`
	CookieAttributes           customfield.NestedObject[WaitingRoomCookieAttributesDataSourceModel]     `tfsdk:"cookie_attributes" json:"cookie_attributes,computed"`
}

func (m *WaitingRoomDataSourceModel) toReadParams(_ context.Context) (params waiting_rooms.WaitingRoomGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.WaitingRoomGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *WaitingRoomDataSourceModel) toListParams(_ context.Context) (params waiting_rooms.WaitingRoomListParams, diags diag.Diagnostics) {
	params = waiting_rooms.WaitingRoomListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type WaitingRoomAdditionalRoutesDataSourceModel struct {
	Host types.String `tfsdk:"host" json:"host,computed"`
	Path types.String `tfsdk:"path" json:"path,computed"`
}

type WaitingRoomCookieAttributesDataSourceModel struct {
	Samesite types.String `tfsdk:"samesite" json:"samesite,computed"`
	Secure   types.String `tfsdk:"secure" json:"secure,computed"`
}
