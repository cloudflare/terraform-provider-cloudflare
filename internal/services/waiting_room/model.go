// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomResultEnvelope struct {
	Result WaitingRoomModel `json:"result"`
}

type WaitingRoomModel struct {
	ID                         types.String                         `tfsdk:"id" json:"id,computed"`
	ZoneID                     types.String                         `tfsdk:"zone_id" path:"zone_id"`
	Host                       types.String                         `tfsdk:"host" json:"host"`
	Name                       types.String                         `tfsdk:"name" json:"name"`
	NewUsersPerMinute          types.Int64                          `tfsdk:"new_users_per_minute" json:"new_users_per_minute"`
	TotalActiveUsers           types.Int64                          `tfsdk:"total_active_users" json:"total_active_users"`
	CookieSuffix               types.String                         `tfsdk:"cookie_suffix" json:"cookie_suffix"`
	AdditionalRoutes           *[]*WaitingRoomAdditionalRoutesModel `tfsdk:"additional_routes" json:"additional_routes"`
	CookieAttributes           *WaitingRoomCookieAttributesModel    `tfsdk:"cookie_attributes" json:"cookie_attributes"`
	CustomPageHTML             types.String                         `tfsdk:"custom_page_html" json:"custom_page_html,computed_optional"`
	DefaultTemplateLanguage    types.String                         `tfsdk:"default_template_language" json:"default_template_language,computed_optional"`
	Description                types.String                         `tfsdk:"description" json:"description,computed_optional"`
	DisableSessionRenewal      types.Bool                           `tfsdk:"disable_session_renewal" json:"disable_session_renewal,computed_optional"`
	JsonResponseEnabled        types.Bool                           `tfsdk:"json_response_enabled" json:"json_response_enabled,computed_optional"`
	Path                       types.String                         `tfsdk:"path" json:"path,computed_optional"`
	QueueAll                   types.Bool                           `tfsdk:"queue_all" json:"queue_all,computed_optional"`
	QueueingMethod             types.String                         `tfsdk:"queueing_method" json:"queueing_method,computed_optional"`
	QueueingStatusCode         types.Int64                          `tfsdk:"queueing_status_code" json:"queueing_status_code,computed_optional"`
	SessionDuration            types.Int64                          `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	Suspended                  types.Bool                           `tfsdk:"suspended" json:"suspended,computed_optional"`
	EnabledOriginCommands      customfield.List[types.String]       `tfsdk:"enabled_origin_commands" json:"enabled_origin_commands,computed_optional"`
	CreatedOn                  timetypes.RFC3339                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn                 timetypes.RFC3339                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	NextEventPrequeueStartTime types.String                         `tfsdk:"next_event_prequeue_start_time" json:"next_event_prequeue_start_time,computed"`
	NextEventStartTime         types.String                         `tfsdk:"next_event_start_time" json:"next_event_start_time,computed"`
}

type WaitingRoomAdditionalRoutesModel struct {
	Host types.String `tfsdk:"host" json:"host,computed_optional"`
	Path types.String `tfsdk:"path" json:"path,computed_optional"`
}

type WaitingRoomCookieAttributesModel struct {
	Samesite types.String `tfsdk:"samesite" json:"samesite,computed_optional"`
	Secure   types.String `tfsdk:"secure" json:"secure,computed_optional"`
}
