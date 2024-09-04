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
	AccountID     types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	AlertInterval types.String                                                  `tfsdk:"alert_interval" json:"alert_interval,computed_optional"`
	AlertType     types.String                                                  `tfsdk:"alert_type" json:"alert_type,computed_optional"`
	Description   types.String                                                  `tfsdk:"description" json:"description,computed_optional"`
	Enabled       types.Bool                                                    `tfsdk:"enabled" json:"enabled,computed_optional"`
	Name          types.String                                                  `tfsdk:"name" json:"name,computed_optional"`
	Mechanisms    customfield.Map[customfield.List[jsontypes.Normalized]]       `tfsdk:"mechanisms" json:"mechanisms,computed_optional"`
	Filters       customfield.NestedObject[NotificationPolicyFiltersModel]      `tfsdk:"filters" json:"filters,computed_optional"`
	Created       timetypes.RFC3339                                             `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified      timetypes.RFC3339                                             `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Success       types.Bool                                                    `tfsdk:"success" json:"success,computed"`
	Errors        customfield.NestedObjectList[NotificationPolicyErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages      customfield.NestedObjectList[NotificationPolicyMessagesModel] `tfsdk:"messages" json:"messages,computed"`
	ResultInfo    customfield.NestedObject[NotificationPolicyResultInfoModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

type NotificationPolicyFiltersModel struct {
	Actions                      customfield.List[types.String] `tfsdk:"actions" json:"actions,computed_optional"`
	AffectedASNs                 customfield.List[types.String] `tfsdk:"affected_asns" json:"affected_asns,computed_optional"`
	AffectedComponents           customfield.List[types.String] `tfsdk:"affected_components" json:"affected_components,computed_optional"`
	AffectedLocations            customfield.List[types.String] `tfsdk:"affected_locations" json:"affected_locations,computed_optional"`
	AirportCode                  customfield.List[types.String] `tfsdk:"airport_code" json:"airport_code,computed_optional"`
	AlertTriggerPreferences      customfield.List[types.String] `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,computed_optional"`
	AlertTriggerPreferencesValue customfield.List[types.String] `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,computed_optional"`
	Enabled                      customfield.List[types.String] `tfsdk:"enabled" json:"enabled,computed_optional"`
	Environment                  customfield.List[types.String] `tfsdk:"environment" json:"environment,computed_optional"`
	Event                        customfield.List[types.String] `tfsdk:"event" json:"event,computed_optional"`
	EventSource                  customfield.List[types.String] `tfsdk:"event_source" json:"event_source,computed_optional"`
	EventType                    customfield.List[types.String] `tfsdk:"event_type" json:"event_type,computed_optional"`
	GroupBy                      customfield.List[types.String] `tfsdk:"group_by" json:"group_by,computed_optional"`
	HealthCheckID                customfield.List[types.String] `tfsdk:"health_check_id" json:"health_check_id,computed_optional"`
	IncidentImpact               customfield.List[types.String] `tfsdk:"incident_impact" json:"incident_impact,computed_optional"`
	InputID                      customfield.List[types.String] `tfsdk:"input_id" json:"input_id,computed_optional"`
	Limit                        customfield.List[types.String] `tfsdk:"limit" json:"limit,computed_optional"`
	LogoTag                      customfield.List[types.String] `tfsdk:"logo_tag" json:"logo_tag,computed_optional"`
	MegabitsPerSecond            customfield.List[types.String] `tfsdk:"megabits_per_second" json:"megabits_per_second,computed_optional"`
	NewHealth                    customfield.List[types.String] `tfsdk:"new_health" json:"new_health,computed_optional"`
	NewStatus                    customfield.List[types.String] `tfsdk:"new_status" json:"new_status,computed_optional"`
	PacketsPerSecond             customfield.List[types.String] `tfsdk:"packets_per_second" json:"packets_per_second,computed_optional"`
	PoolID                       customfield.List[types.String] `tfsdk:"pool_id" json:"pool_id,computed_optional"`
	Product                      customfield.List[types.String] `tfsdk:"product" json:"product,computed_optional"`
	ProjectID                    customfield.List[types.String] `tfsdk:"project_id" json:"project_id,computed_optional"`
	Protocol                     customfield.List[types.String] `tfsdk:"protocol" json:"protocol,computed_optional"`
	QueryTag                     customfield.List[types.String] `tfsdk:"query_tag" json:"query_tag,computed_optional"`
	RequestsPerSecond            customfield.List[types.String] `tfsdk:"requests_per_second" json:"requests_per_second,computed_optional"`
	Selectors                    customfield.List[types.String] `tfsdk:"selectors" json:"selectors,computed_optional"`
	Services                     customfield.List[types.String] `tfsdk:"services" json:"services,computed_optional"`
	Slo                          customfield.List[types.String] `tfsdk:"slo" json:"slo,computed_optional"`
	Status                       customfield.List[types.String] `tfsdk:"status" json:"status,computed_optional"`
	TargetHostname               customfield.List[types.String] `tfsdk:"target_hostname" json:"target_hostname,computed_optional"`
	TargetIP                     customfield.List[types.String] `tfsdk:"target_ip" json:"target_ip,computed_optional"`
	TargetZoneName               customfield.List[types.String] `tfsdk:"target_zone_name" json:"target_zone_name,computed_optional"`
	TrafficExclusions            customfield.List[types.String] `tfsdk:"traffic_exclusions" json:"traffic_exclusions,computed_optional"`
	TunnelID                     customfield.List[types.String] `tfsdk:"tunnel_id" json:"tunnel_id,computed_optional"`
	TunnelName                   customfield.List[types.String] `tfsdk:"tunnel_name" json:"tunnel_name,computed_optional"`
	Where                        customfield.List[types.String] `tfsdk:"where" json:"where,computed_optional"`
	Zones                        customfield.List[types.String] `tfsdk:"zones" json:"zones,computed_optional"`
}

type NotificationPolicyErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type NotificationPolicyMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type NotificationPolicyResultInfoModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}
