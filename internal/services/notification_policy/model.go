// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyResultEnvelope struct {
	Result NotificationPolicyModel `json:"result"`
}

type NotificationPolicyModel struct {
	ID            types.String                                                  `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                                  `tfsdk:"account_id" path:"account_id"`
	AlertType     types.String                                                  `tfsdk:"alert_type" json:"alert_type"`
	Name          types.String                                                  `tfsdk:"name" json:"name"`
	Mechanisms    map[string]*[]jsontypes.Normalized                            `tfsdk:"mechanisms" json:"mechanisms"`
	AlertInterval types.String                                                  `tfsdk:"alert_interval" json:"alert_interval"`
	Description   types.String                                                  `tfsdk:"description" json:"description"`
	Filters       *NotificationPolicyFiltersModel                               `tfsdk:"filters" json:"filters"`
	Enabled       types.Bool                                                    `tfsdk:"enabled" json:"enabled,computed_optional"`
	Created       timetypes.RFC3339                                             `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified      timetypes.RFC3339                                             `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Success       types.Bool                                                    `tfsdk:"success" json:"success,computed"`
	Errors        customfield.NestedObjectList[NotificationPolicyErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages      customfield.NestedObjectList[NotificationPolicyMessagesModel] `tfsdk:"messages" json:"messages,computed"`
	ResultInfo    customfield.NestedObject[NotificationPolicyResultInfoModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

type NotificationPolicyFiltersModel struct {
	Actions                      types.List `tfsdk:"actions" json:"actions,computed_optional"`
	AffectedASNs                 types.List `tfsdk:"affected_asns" json:"affected_asns,computed_optional"`
	AffectedComponents           types.List `tfsdk:"affected_components" json:"affected_components,computed_optional"`
	AffectedLocations            types.List `tfsdk:"affected_locations" json:"affected_locations,computed_optional"`
	AirportCode                  types.List `tfsdk:"airport_code" json:"airport_code,computed_optional"`
	AlertTriggerPreferences      types.List `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,computed_optional"`
	AlertTriggerPreferencesValue types.List `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,computed_optional"`
	Enabled                      types.List `tfsdk:"enabled" json:"enabled,computed_optional"`
	Environment                  types.List `tfsdk:"environment" json:"environment,computed_optional"`
	Event                        types.List `tfsdk:"event" json:"event,computed_optional"`
	EventSource                  types.List `tfsdk:"event_source" json:"event_source,computed_optional"`
	EventType                    types.List `tfsdk:"event_type" json:"event_type,computed_optional"`
	GroupBy                      types.List `tfsdk:"group_by" json:"group_by,computed_optional"`
	HealthCheckID                types.List `tfsdk:"health_check_id" json:"health_check_id,computed_optional"`
	IncidentImpact               types.List `tfsdk:"incident_impact" json:"incident_impact,computed_optional"`
	InputID                      types.List `tfsdk:"input_id" json:"input_id,computed_optional"`
	Limit                        types.List `tfsdk:"limit" json:"limit,computed_optional"`
	LogoTag                      types.List `tfsdk:"logo_tag" json:"logo_tag,computed_optional"`
	MegabitsPerSecond            types.List `tfsdk:"megabits_per_second" json:"megabits_per_second,computed_optional"`
	NewHealth                    types.List `tfsdk:"new_health" json:"new_health,computed_optional"`
	NewStatus                    types.List `tfsdk:"new_status" json:"new_status,computed_optional"`
	PacketsPerSecond             types.List `tfsdk:"packets_per_second" json:"packets_per_second,computed_optional"`
	PoolID                       types.List `tfsdk:"pool_id" json:"pool_id,computed_optional"`
	Product                      types.List `tfsdk:"product" json:"product,computed_optional"`
	ProjectID                    types.List `tfsdk:"project_id" json:"project_id,computed_optional"`
	Protocol                     types.List `tfsdk:"protocol" json:"protocol,computed_optional"`
	QueryTag                     types.List `tfsdk:"query_tag" json:"query_tag,computed_optional"`
	RequestsPerSecond            types.List `tfsdk:"requests_per_second" json:"requests_per_second,computed_optional"`
	Selectors                    types.List `tfsdk:"selectors" json:"selectors,computed_optional"`
	Services                     types.List `tfsdk:"services" json:"services,computed_optional"`
	Slo                          types.List `tfsdk:"slo" json:"slo,computed_optional"`
	Status                       types.List `tfsdk:"status" json:"status,computed_optional"`
	TargetHostname               types.List `tfsdk:"target_hostname" json:"target_hostname,computed_optional"`
	TargetIP                     types.List `tfsdk:"target_ip" json:"target_ip,computed_optional"`
	TargetZoneName               types.List `tfsdk:"target_zone_name" json:"target_zone_name,computed_optional"`
	TrafficExclusions            types.List `tfsdk:"traffic_exclusions" json:"traffic_exclusions,computed_optional"`
	TunnelID                     types.List `tfsdk:"tunnel_id" json:"tunnel_id,computed_optional"`
	TunnelName                   types.List `tfsdk:"tunnel_name" json:"tunnel_name,computed_optional"`
	Where                        types.List `tfsdk:"where" json:"where,computed_optional"`
	Zones                        types.List `tfsdk:"zones" json:"zones,computed_optional"`
}

type NotificationPolicyErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type NotificationPolicyMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type NotificationPolicyResultInfoModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed_optional"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed_optional"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed_optional"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed_optional"`
}
