// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomsResultListDataSourceEnvelope struct {
	Result *[]*WaitingRoomsResultDataSourceModel `json:"result,computed"`
}

type WaitingRoomsDataSourceModel struct {
	ZoneID   types.String                          `tfsdk:"zone_id" path:"zone_id"`
	Page     jsontypes.Normalized                  `tfsdk:"page" query:"page"`
	PerPage  jsontypes.Normalized                  `tfsdk:"per_page" query:"per_page"`
	MaxItems types.Int64                           `tfsdk:"max_items"`
	Result   *[]*WaitingRoomsResultDataSourceModel `tfsdk:"result"`
}

type WaitingRoomsResultDataSourceModel struct {
	ID                         types.String                                    `tfsdk:"id" json:"id"`
	AdditionalRoutes           *[]*WaitingRoomsAdditionalRoutesDataSourceModel `tfsdk:"additional_routes" json:"additional_routes"`
	CookieAttributes           *WaitingRoomsCookieAttributesDataSourceModel    `tfsdk:"cookie_attributes" json:"cookie_attributes"`
	CookieSuffix               types.String                                    `tfsdk:"cookie_suffix" json:"cookie_suffix"`
	CreatedOn                  timetypes.RFC3339                               `tfsdk:"created_on" json:"created_on,computed"`
	CustomPageHTML             types.String                                    `tfsdk:"custom_page_html" json:"custom_page_html,computed"`
	DefaultTemplateLanguage    types.String                                    `tfsdk:"default_template_language" json:"default_template_language,computed"`
	Description                types.String                                    `tfsdk:"description" json:"description,computed"`
	DisableSessionRenewal      types.Bool                                      `tfsdk:"disable_session_renewal" json:"disable_session_renewal,computed"`
	Host                       types.String                                    `tfsdk:"host" json:"host"`
	JsonResponseEnabled        types.Bool                                      `tfsdk:"json_response_enabled" json:"json_response_enabled,computed"`
	ModifiedOn                 timetypes.RFC3339                               `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                       types.String                                    `tfsdk:"name" json:"name"`
	NewUsersPerMinute          types.Int64                                     `tfsdk:"new_users_per_minute" json:"new_users_per_minute"`
	NextEventPrequeueStartTime types.String                                    `tfsdk:"next_event_prequeue_start_time" json:"next_event_prequeue_start_time"`
	NextEventStartTime         types.String                                    `tfsdk:"next_event_start_time" json:"next_event_start_time"`
	Path                       types.String                                    `tfsdk:"path" json:"path,computed"`
	QueueAll                   types.Bool                                      `tfsdk:"queue_all" json:"queue_all,computed"`
	QueueingMethod             types.String                                    `tfsdk:"queueing_method" json:"queueing_method,computed"`
	QueueingStatusCode         types.Int64                                     `tfsdk:"queueing_status_code" json:"queueing_status_code,computed"`
	SessionDuration            types.Int64                                     `tfsdk:"session_duration" json:"session_duration,computed"`
	Suspended                  types.Bool                                      `tfsdk:"suspended" json:"suspended,computed"`
	TotalActiveUsers           types.Int64                                     `tfsdk:"total_active_users" json:"total_active_users"`
}

type WaitingRoomsAdditionalRoutesDataSourceModel struct {
	Host types.String `tfsdk:"host" json:"host"`
	Path types.String `tfsdk:"path" json:"path,computed"`
}

type WaitingRoomsCookieAttributesDataSourceModel struct {
	Samesite types.String `tfsdk:"samesite" json:"samesite,computed"`
	Secure   types.String `tfsdk:"secure" json:"secure,computed"`
}
