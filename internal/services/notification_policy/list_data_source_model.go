// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPoliciesResultListDataSourceEnvelope struct {
	Result *[]*NotificationPoliciesItemsDataSourceModel `json:"result,computed"`
}

type NotificationPoliciesDataSourceModel struct {
	AccountID types.String                                 `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                  `tfsdk:"max_items"`
	Items     *[]*NotificationPoliciesItemsDataSourceModel `tfsdk:"items"`
}

type NotificationPoliciesItemsDataSourceModel struct {
	ID          types.String               `tfsdk:"id" json:"id,computed"`
	AlertType   types.String               `tfsdk:"alert_type" json:"alert_type,computed"`
	Created     types.String               `tfsdk:"created" json:"created,computed"`
	Description types.String               `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                 `tfsdk:"enabled" json:"enabled,computed"`
	Mechanisms  map[string]*[]types.String `tfsdk:"mechanisms" json:"mechanisms,computed"`
	Modified    types.String               `tfsdk:"modified" json:"modified,computed"`
	Name        types.String               `tfsdk:"name" json:"name,computed"`
}

type NotificationPoliciesItemsFiltersDataSourceModel struct {
	Actions                      *[]types.String `tfsdk:"actions" json:"actions,computed"`
	AffectedASNs                 *[]types.String `tfsdk:"affected_asns" json:"affected_asns,computed"`
	AffectedComponents           *[]types.String `tfsdk:"affected_components" json:"affected_components,computed"`
	AffectedLocations            *[]types.String `tfsdk:"affected_locations" json:"affected_locations,computed"`
	AirportCode                  *[]types.String `tfsdk:"airport_code" json:"airport_code,computed"`
	AlertTriggerPreferences      *[]types.String `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,computed"`
	AlertTriggerPreferencesValue *[]types.String `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,computed"`
	Enabled                      *[]types.String `tfsdk:"enabled" json:"enabled,computed"`
	Environment                  *[]types.String `tfsdk:"environment" json:"environment,computed"`
	Event                        *[]types.String `tfsdk:"event" json:"event,computed"`
	EventSource                  *[]types.String `tfsdk:"event_source" json:"event_source,computed"`
	EventType                    *[]types.String `tfsdk:"event_type" json:"event_type,computed"`
	GroupBy                      *[]types.String `tfsdk:"group_by" json:"group_by,computed"`
	HealthCheckID                *[]types.String `tfsdk:"health_check_id" json:"health_check_id,computed"`
	IncidentImpact               *[]types.String `tfsdk:"incident_impact" json:"incident_impact,computed"`
	InputID                      *[]types.String `tfsdk:"input_id" json:"input_id,computed"`
	Limit                        *[]types.String `tfsdk:"limit" json:"limit,computed"`
	LogoTag                      *[]types.String `tfsdk:"logo_tag" json:"logo_tag,computed"`
	MegabitsPerSecond            *[]types.String `tfsdk:"megabits_per_second" json:"megabits_per_second,computed"`
	NewHealth                    *[]types.String `tfsdk:"new_health" json:"new_health,computed"`
	NewStatus                    *[]types.String `tfsdk:"new_status" json:"new_status,computed"`
	PacketsPerSecond             *[]types.String `tfsdk:"packets_per_second" json:"packets_per_second,computed"`
	PoolID                       *[]types.String `tfsdk:"pool_id" json:"pool_id,computed"`
	Product                      *[]types.String `tfsdk:"product" json:"product,computed"`
	ProjectID                    *[]types.String `tfsdk:"project_id" json:"project_id,computed"`
	Protocol                     *[]types.String `tfsdk:"protocol" json:"protocol,computed"`
	QueryTag                     *[]types.String `tfsdk:"query_tag" json:"query_tag,computed"`
	RequestsPerSecond            *[]types.String `tfsdk:"requests_per_second" json:"requests_per_second,computed"`
	Selectors                    *[]types.String `tfsdk:"selectors" json:"selectors,computed"`
	Services                     *[]types.String `tfsdk:"services" json:"services,computed"`
	Slo                          *[]types.String `tfsdk:"slo" json:"slo,computed"`
	Status                       *[]types.String `tfsdk:"status" json:"status,computed"`
	TargetHostname               *[]types.String `tfsdk:"target_hostname" json:"target_hostname,computed"`
	TargetIP                     *[]types.String `tfsdk:"target_ip" json:"target_ip,computed"`
	TargetZoneName               *[]types.String `tfsdk:"target_zone_name" json:"target_zone_name,computed"`
	TrafficExclusions            *[]types.String `tfsdk:"traffic_exclusions" json:"traffic_exclusions,computed"`
	TunnelID                     *[]types.String `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName                   *[]types.String `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
	Where                        *[]types.String `tfsdk:"where" json:"where,computed"`
	Zones                        *[]types.String `tfsdk:"zones" json:"zones,computed"`
}
