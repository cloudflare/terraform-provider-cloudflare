// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/alerting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPoliciesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[NotificationPoliciesResultDataSourceModel] `json:"result,computed"`
}

type NotificationPoliciesDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[NotificationPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *NotificationPoliciesDataSourceModel) toListParams() (params alerting.PolicyListParams, diags diag.Diagnostics) {
	params = alerting.PolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type NotificationPoliciesResultDataSourceModel struct {
	ID            types.String                                                         `tfsdk:"id" json:"id,computed"`
	AlertInterval types.String                                                         `tfsdk:"alert_interval" json:"alert_interval,computed"`
	AlertType     types.String                                                         `tfsdk:"alert_type" json:"alert_type,computed"`
	Created       timetypes.RFC3339                                                    `tfsdk:"created" json:"created,computed" format:"date-time"`
	Description   types.String                                                         `tfsdk:"description" json:"description,computed"`
	Enabled       types.Bool                                                           `tfsdk:"enabled" json:"enabled,computed"`
	Filters       customfield.NestedObject[NotificationPoliciesFiltersDataSourceModel] `tfsdk:"filters" json:"filters,computed"`
	Mechanisms    map[string]types.List                                                `tfsdk:"mechanisms" json:"mechanisms,computed"`
	Modified      timetypes.RFC3339                                                    `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name          types.String                                                         `tfsdk:"name" json:"name,computed"`
}

type NotificationPoliciesFiltersDataSourceModel struct {
	Actions                      types.List `tfsdk:"actions" json:"actions,computed"`
	AffectedASNs                 types.List `tfsdk:"affected_asns" json:"affected_asns,computed"`
	AffectedComponents           types.List `tfsdk:"affected_components" json:"affected_components,computed"`
	AffectedLocations            types.List `tfsdk:"affected_locations" json:"affected_locations,computed"`
	AirportCode                  types.List `tfsdk:"airport_code" json:"airport_code,computed"`
	AlertTriggerPreferences      types.List `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,computed"`
	AlertTriggerPreferencesValue types.List `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,computed"`
	Enabled                      types.List `tfsdk:"enabled" json:"enabled,computed"`
	Environment                  types.List `tfsdk:"environment" json:"environment,computed"`
	Event                        types.List `tfsdk:"event" json:"event,computed"`
	EventSource                  types.List `tfsdk:"event_source" json:"event_source,computed"`
	EventType                    types.List `tfsdk:"event_type" json:"event_type,computed"`
	GroupBy                      types.List `tfsdk:"group_by" json:"group_by,computed"`
	HealthCheckID                types.List `tfsdk:"health_check_id" json:"health_check_id,computed"`
	IncidentImpact               types.List `tfsdk:"incident_impact" json:"incident_impact,computed"`
	InputID                      types.List `tfsdk:"input_id" json:"input_id,computed"`
	Limit                        types.List `tfsdk:"limit" json:"limit,computed"`
	LogoTag                      types.List `tfsdk:"logo_tag" json:"logo_tag,computed"`
	MegabitsPerSecond            types.List `tfsdk:"megabits_per_second" json:"megabits_per_second,computed"`
	NewHealth                    types.List `tfsdk:"new_health" json:"new_health,computed"`
	NewStatus                    types.List `tfsdk:"new_status" json:"new_status,computed"`
	PacketsPerSecond             types.List `tfsdk:"packets_per_second" json:"packets_per_second,computed"`
	PoolID                       types.List `tfsdk:"pool_id" json:"pool_id,computed"`
	Product                      types.List `tfsdk:"product" json:"product,computed"`
	ProjectID                    types.List `tfsdk:"project_id" json:"project_id,computed"`
	Protocol                     types.List `tfsdk:"protocol" json:"protocol,computed"`
	QueryTag                     types.List `tfsdk:"query_tag" json:"query_tag,computed"`
	RequestsPerSecond            types.List `tfsdk:"requests_per_second" json:"requests_per_second,computed"`
	Selectors                    types.List `tfsdk:"selectors" json:"selectors,computed"`
	Services                     types.List `tfsdk:"services" json:"services,computed"`
	Slo                          types.List `tfsdk:"slo" json:"slo,computed"`
	Status                       types.List `tfsdk:"status" json:"status,computed"`
	TargetHostname               types.List `tfsdk:"target_hostname" json:"target_hostname,computed"`
	TargetIP                     types.List `tfsdk:"target_ip" json:"target_ip,computed"`
	TargetZoneName               types.List `tfsdk:"target_zone_name" json:"target_zone_name,computed"`
	TrafficExclusions            types.List `tfsdk:"traffic_exclusions" json:"traffic_exclusions,computed"`
	TunnelID                     types.List `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName                   types.List `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
	Where                        types.List `tfsdk:"where" json:"where,computed"`
	Zones                        types.List `tfsdk:"zones" json:"zones,computed"`
}
