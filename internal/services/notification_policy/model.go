// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyResultEnvelope struct {
	Result NotificationPolicyModel `json:"result,computed"`
}

type NotificationPolicyModel struct {
	ID          types.String                    `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                    `tfsdk:"account_id" path:"account_id"`
	AlertType   types.String                    `tfsdk:"alert_type" json:"alert_type"`
	Enabled     types.Bool                      `tfsdk:"enabled" json:"enabled"`
	Mechanisms  map[string]*[]types.String      `tfsdk:"mechanisms" json:"mechanisms"`
	Name        types.String                    `tfsdk:"name" json:"name"`
	Description types.String                    `tfsdk:"description" json:"description"`
	Filters     *NotificationPolicyFiltersModel `tfsdk:"filters" json:"filters"`
}

type NotificationPolicyFiltersModel struct {
	Actions                      *[]types.String `tfsdk:"actions" json:"actions"`
	AffectedASNs                 *[]types.String `tfsdk:"affected_asns" json:"affected_asns"`
	AffectedComponents           *[]types.String `tfsdk:"affected_components" json:"affected_components"`
	AffectedLocations            *[]types.String `tfsdk:"affected_locations" json:"affected_locations"`
	AirportCode                  *[]types.String `tfsdk:"airport_code" json:"airport_code"`
	AlertTriggerPreferences      *[]types.String `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences"`
	AlertTriggerPreferencesValue *[]types.String `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value"`
	Enabled                      *[]types.String `tfsdk:"enabled" json:"enabled"`
	Environment                  *[]types.String `tfsdk:"environment" json:"environment"`
	Event                        *[]types.String `tfsdk:"event" json:"event"`
	EventSource                  *[]types.String `tfsdk:"event_source" json:"event_source"`
	EventType                    *[]types.String `tfsdk:"event_type" json:"event_type"`
	GroupBy                      *[]types.String `tfsdk:"group_by" json:"group_by"`
	HealthCheckID                *[]types.String `tfsdk:"health_check_id" json:"health_check_id"`
	IncidentImpact               *[]types.String `tfsdk:"incident_impact" json:"incident_impact"`
	InputID                      *[]types.String `tfsdk:"input_id" json:"input_id"`
	Limit                        *[]types.String `tfsdk:"limit" json:"limit"`
	LogoTag                      *[]types.String `tfsdk:"logo_tag" json:"logo_tag"`
	MegabitsPerSecond            *[]types.String `tfsdk:"megabits_per_second" json:"megabits_per_second"`
	NewHealth                    *[]types.String `tfsdk:"new_health" json:"new_health"`
	NewStatus                    *[]types.String `tfsdk:"new_status" json:"new_status"`
	PacketsPerSecond             *[]types.String `tfsdk:"packets_per_second" json:"packets_per_second"`
	PoolID                       *[]types.String `tfsdk:"pool_id" json:"pool_id"`
	Product                      *[]types.String `tfsdk:"product" json:"product"`
	ProjectID                    *[]types.String `tfsdk:"project_id" json:"project_id"`
	Protocol                     *[]types.String `tfsdk:"protocol" json:"protocol"`
	QueryTag                     *[]types.String `tfsdk:"query_tag" json:"query_tag"`
	RequestsPerSecond            *[]types.String `tfsdk:"requests_per_second" json:"requests_per_second"`
	Selectors                    *[]types.String `tfsdk:"selectors" json:"selectors"`
	Services                     *[]types.String `tfsdk:"services" json:"services"`
	Slo                          *[]types.String `tfsdk:"slo" json:"slo"`
	Status                       *[]types.String `tfsdk:"status" json:"status"`
	TargetHostname               *[]types.String `tfsdk:"target_hostname" json:"target_hostname"`
	TargetIP                     *[]types.String `tfsdk:"target_ip" json:"target_ip"`
	TargetZoneName               *[]types.String `tfsdk:"target_zone_name" json:"target_zone_name"`
	TrafficExclusions            *[]types.String `tfsdk:"traffic_exclusions" json:"traffic_exclusions"`
	TunnelID                     *[]types.String `tfsdk:"tunnel_id" json:"tunnel_id"`
	TunnelName                   *[]types.String `tfsdk:"tunnel_name" json:"tunnel_name"`
	Where                        *[]types.String `tfsdk:"where" json:"where"`
	Zones                        *[]types.String `tfsdk:"zones" json:"zones"`
}
