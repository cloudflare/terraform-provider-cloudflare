// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/alerting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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
	ID            types.String                                `tfsdk:"id" json:"id,computed"`
	AlertInterval types.String                                `tfsdk:"alert_interval" json:"alert_interval,computed_optional"`
	AlertType     types.String                                `tfsdk:"alert_type" json:"alert_type,computed_optional"`
	Created       timetypes.RFC3339                           `tfsdk:"created" json:"created,computed" format:"date-time"`
	Description   types.String                                `tfsdk:"description" json:"description,computed_optional"`
	Enabled       types.Bool                                  `tfsdk:"enabled" json:"enabled,computed"`
	Filters       *NotificationPoliciesFiltersDataSourceModel `tfsdk:"filters" json:"filters,computed_optional"`
	Mechanisms    map[string]*[]jsontypes.Normalized          `tfsdk:"mechanisms" json:"mechanisms,computed_optional"`
	Modified      timetypes.RFC3339                           `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name          types.String                                `tfsdk:"name" json:"name,computed_optional"`
}

type NotificationPoliciesFiltersDataSourceModel struct {
	Actions                      *[]types.String `tfsdk:"actions" json:"actions,computed_optional"`
	AffectedASNs                 *[]types.String `tfsdk:"affected_asns" json:"affected_asns,computed_optional"`
	AffectedComponents           *[]types.String `tfsdk:"affected_components" json:"affected_components,computed_optional"`
	AffectedLocations            *[]types.String `tfsdk:"affected_locations" json:"affected_locations,computed_optional"`
	AirportCode                  *[]types.String `tfsdk:"airport_code" json:"airport_code,computed_optional"`
	AlertTriggerPreferences      *[]types.String `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,computed_optional"`
	AlertTriggerPreferencesValue *[]types.String `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,computed_optional"`
	Enabled                      *[]types.String `tfsdk:"enabled" json:"enabled,computed_optional"`
	Environment                  *[]types.String `tfsdk:"environment" json:"environment,computed_optional"`
	Event                        *[]types.String `tfsdk:"event" json:"event,computed_optional"`
	EventSource                  *[]types.String `tfsdk:"event_source" json:"event_source,computed_optional"`
	EventType                    *[]types.String `tfsdk:"event_type" json:"event_type,computed_optional"`
	GroupBy                      *[]types.String `tfsdk:"group_by" json:"group_by,computed_optional"`
	HealthCheckID                *[]types.String `tfsdk:"health_check_id" json:"health_check_id,computed_optional"`
	IncidentImpact               *[]types.String `tfsdk:"incident_impact" json:"incident_impact,computed_optional"`
	InputID                      *[]types.String `tfsdk:"input_id" json:"input_id,computed_optional"`
	Limit                        *[]types.String `tfsdk:"limit" json:"limit,computed_optional"`
	LogoTag                      *[]types.String `tfsdk:"logo_tag" json:"logo_tag,computed_optional"`
	MegabitsPerSecond            *[]types.String `tfsdk:"megabits_per_second" json:"megabits_per_second,computed_optional"`
	NewHealth                    *[]types.String `tfsdk:"new_health" json:"new_health,computed_optional"`
	NewStatus                    *[]types.String `tfsdk:"new_status" json:"new_status,computed_optional"`
	PacketsPerSecond             *[]types.String `tfsdk:"packets_per_second" json:"packets_per_second,computed_optional"`
	PoolID                       *[]types.String `tfsdk:"pool_id" json:"pool_id,computed_optional"`
	Product                      *[]types.String `tfsdk:"product" json:"product,computed_optional"`
	ProjectID                    *[]types.String `tfsdk:"project_id" json:"project_id,computed_optional"`
	Protocol                     *[]types.String `tfsdk:"protocol" json:"protocol,computed_optional"`
	QueryTag                     *[]types.String `tfsdk:"query_tag" json:"query_tag,computed_optional"`
	RequestsPerSecond            *[]types.String `tfsdk:"requests_per_second" json:"requests_per_second,computed_optional"`
	Selectors                    *[]types.String `tfsdk:"selectors" json:"selectors,computed_optional"`
	Services                     *[]types.String `tfsdk:"services" json:"services,computed_optional"`
	Slo                          *[]types.String `tfsdk:"slo" json:"slo,computed_optional"`
	Status                       *[]types.String `tfsdk:"status" json:"status,computed_optional"`
	TargetHostname               *[]types.String `tfsdk:"target_hostname" json:"target_hostname,computed_optional"`
	TargetIP                     *[]types.String `tfsdk:"target_ip" json:"target_ip,computed_optional"`
	TargetZoneName               *[]types.String `tfsdk:"target_zone_name" json:"target_zone_name,computed_optional"`
	TrafficExclusions            *[]types.String `tfsdk:"traffic_exclusions" json:"traffic_exclusions,computed_optional"`
	TunnelID                     *[]types.String `tfsdk:"tunnel_id" json:"tunnel_id,computed_optional"`
	TunnelName                   *[]types.String `tfsdk:"tunnel_name" json:"tunnel_name,computed_optional"`
	Where                        *[]types.String `tfsdk:"where" json:"where,computed_optional"`
	Zones                        *[]types.String `tfsdk:"zones" json:"zones,computed_optional"`
}
