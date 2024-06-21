// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomResultEnvelope struct {
	Result WaitingRoomModel `json:"result,computed"`
}

type WaitingRoomModel struct {
	ID                      types.String                         `tfsdk:"id" json:"id,computed"`
	ZoneID                  types.String                         `tfsdk:"zone_id" path:"zone_id"`
	Host                    types.String                         `tfsdk:"host" json:"host"`
	Name                    types.String                         `tfsdk:"name" json:"name"`
	NewUsersPerMinute       types.Int64                          `tfsdk:"new_users_per_minute" json:"new_users_per_minute"`
	TotalActiveUsers        types.Int64                          `tfsdk:"total_active_users" json:"total_active_users"`
	AdditionalRoutes        *[]*WaitingRoomAdditionalRoutesModel `tfsdk:"additional_routes" json:"additional_routes"`
	CookieAttributes        *WaitingRoomCookieAttributesModel    `tfsdk:"cookie_attributes" json:"cookie_attributes"`
	CookieSuffix            types.String                         `tfsdk:"cookie_suffix" json:"cookie_suffix"`
	CustomPageHTML          types.String                         `tfsdk:"custom_page_html" json:"custom_page_html"`
	DefaultTemplateLanguage types.String                         `tfsdk:"default_template_language" json:"default_template_language"`
	Description             types.String                         `tfsdk:"description" json:"description"`
	DisableSessionRenewal   types.Bool                           `tfsdk:"disable_session_renewal" json:"disable_session_renewal"`
	JsonResponseEnabled     types.Bool                           `tfsdk:"json_response_enabled" json:"json_response_enabled"`
	Path                    types.String                         `tfsdk:"path" json:"path"`
	QueueAll                types.Bool                           `tfsdk:"queue_all" json:"queue_all"`
	QueueingMethod          types.String                         `tfsdk:"queueing_method" json:"queueing_method"`
	QueueingStatusCode      types.Int64                          `tfsdk:"queueing_status_code" json:"queueing_status_code"`
	SessionDuration         types.Int64                          `tfsdk:"session_duration" json:"session_duration"`
	Suspended               types.Bool                           `tfsdk:"suspended" json:"suspended"`
}

type WaitingRoomAdditionalRoutesModel struct {
	Host types.String `tfsdk:"host" json:"host"`
	Path types.String `tfsdk:"path" json:"path"`
}

type WaitingRoomCookieAttributesModel struct {
	Samesite types.String `tfsdk:"samesite" json:"samesite"`
	Secure   types.String `tfsdk:"secure" json:"secure"`
}
