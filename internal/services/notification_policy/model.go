// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyResultEnvelope struct {
	Result NotificationPolicyModel `json:"result"`
}

type NotificationPolicyModel struct {
	ID            types.String                                                  `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	AlertType     types.String                                                  `tfsdk:"alert_type" json:"alert_type,required"`
	Name          types.String                                                  `tfsdk:"name" json:"name,required"`
	Mechanisms    map[string]*[]*NotificationPolicyMechanismsModel              `tfsdk:"mechanisms" json:"mechanisms,required"`
	AlertInterval types.String                                                  `tfsdk:"alert_interval" json:"alert_interval,optional"`
	Description   types.String                                                  `tfsdk:"description" json:"description,optional"`
	Enabled       types.Bool                                                    `tfsdk:"enabled" json:"enabled,computed_optional"`
	Filters       customfield.NestedObject[NotificationPolicyFiltersModel]      `tfsdk:"filters" json:"filters,computed_optional"`
	Created       timetypes.RFC3339                                             `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified      timetypes.RFC3339                                             `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Success       types.Bool                                                    `tfsdk:"success" json:"success,computed"`
	Errors        customfield.NestedObjectList[NotificationPolicyErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages      customfield.NestedObjectList[NotificationPolicyMessagesModel] `tfsdk:"messages" json:"messages,computed"`
	ResultInfo    customfield.NestedObject[NotificationPolicyResultInfoModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

func (m NotificationPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NotificationPolicyModel) MarshalJSONForUpdate(state NotificationPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NotificationPolicyMechanismsModel struct {
	ID types.String `tfsdk:"id" json:"id,optional"`
}

type NotificationPolicyFiltersModel struct {
	Actions                      *[]types.String `tfsdk:"actions" json:"actions,optional"`
	AffectedASNs                 *[]types.String `tfsdk:"affected_asns" json:"affected_asns,optional"`
	AffectedComponents           *[]types.String `tfsdk:"affected_components" json:"affected_components,optional"`
	AffectedLocations            *[]types.String `tfsdk:"affected_locations" json:"affected_locations,optional"`
	AirportCode                  *[]types.String `tfsdk:"airport_code" json:"airport_code,optional"`
	AlertTriggerPreferences      *[]types.String `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,optional"`
	AlertTriggerPreferencesValue *[]types.String `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,optional"`
	Enabled                      *[]types.String `tfsdk:"enabled" json:"enabled,optional"`
	Environment                  *[]types.String `tfsdk:"environment" json:"environment,optional"`
	Event                        *[]types.String `tfsdk:"event" json:"event,optional"`
	EventSource                  *[]types.String `tfsdk:"event_source" json:"event_source,optional"`
	EventType                    *[]types.String `tfsdk:"event_type" json:"event_type,optional"`
	GroupBy                      *[]types.String `tfsdk:"group_by" json:"group_by,optional"`
	HealthCheckID                *[]types.String `tfsdk:"health_check_id" json:"health_check_id,optional"`
	IncidentImpact               *[]types.String `tfsdk:"incident_impact" json:"incident_impact,optional"`
	InputID                      *[]types.String `tfsdk:"input_id" json:"input_id,optional"`
	Limit                        *[]types.String `tfsdk:"limit" json:"limit,optional"`
	LogoTag                      *[]types.String `tfsdk:"logo_tag" json:"logo_tag,optional"`
	MegabitsPerSecond            *[]types.String `tfsdk:"megabits_per_second" json:"megabits_per_second,optional"`
	NewHealth                    *[]types.String `tfsdk:"new_health" json:"new_health,optional"`
	NewStatus                    *[]types.String `tfsdk:"new_status" json:"new_status,optional"`
	PacketsPerSecond             *[]types.String `tfsdk:"packets_per_second" json:"packets_per_second,optional"`
	PoolID                       *[]types.String `tfsdk:"pool_id" json:"pool_id,optional"`
	PopName                      *[]types.String `tfsdk:"pop_name" json:"pop_name,optional"`
	Product                      *[]types.String `tfsdk:"product" json:"product,optional"`
	ProjectID                    *[]types.String `tfsdk:"project_id" json:"project_id,optional"`
	Protocol                     *[]types.String `tfsdk:"protocol" json:"protocol,optional"`
	QueryTag                     *[]types.String `tfsdk:"query_tag" json:"query_tag,optional"`
	RequestsPerSecond            *[]types.String `tfsdk:"requests_per_second" json:"requests_per_second,optional"`
	Selectors                    *[]types.String `tfsdk:"selectors" json:"selectors,optional"`
	Services                     *[]types.String `tfsdk:"services" json:"services,optional"`
	Slo                          *[]types.String `tfsdk:"slo" json:"slo,optional"`
	Status                       *[]types.String `tfsdk:"status" json:"status,optional"`
	TargetHostname               *[]types.String `tfsdk:"target_hostname" json:"target_hostname,optional"`
	TargetIP                     *[]types.String `tfsdk:"target_ip" json:"target_ip,optional"`
	TargetZoneName               *[]types.String `tfsdk:"target_zone_name" json:"target_zone_name,optional"`
	TrafficExclusions            *[]types.String `tfsdk:"traffic_exclusions" json:"traffic_exclusions,optional"`
	TunnelID                     *[]types.String `tfsdk:"tunnel_id" json:"tunnel_id,optional"`
	TunnelName                   *[]types.String `tfsdk:"tunnel_name" json:"tunnel_name,optional"`
	Where                        *[]types.String `tfsdk:"where" json:"where,optional"`
	Zones                        *[]types.String `tfsdk:"zones" json:"zones,optional"`
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
