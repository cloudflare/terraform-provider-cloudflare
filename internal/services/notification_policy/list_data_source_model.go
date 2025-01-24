// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/alerting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPoliciesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[NotificationPoliciesResultDataSourceModel] `json:"result,computed"`
}

type NotificationPoliciesDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[NotificationPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *NotificationPoliciesDataSourceModel) toListParams(_ context.Context) (params alerting.PolicyListParams, diags diag.Diagnostics) {
	params = alerting.PolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type NotificationPoliciesResultDataSourceModel struct {
	ID            types.String                                                            `tfsdk:"id" json:"id,computed"`
	AlertInterval types.String                                                            `tfsdk:"alert_interval" json:"alert_interval,computed"`
	AlertType     types.String                                                            `tfsdk:"alert_type" json:"alert_type,computed"`
	Created       timetypes.RFC3339                                                       `tfsdk:"created" json:"created,computed" format:"date-time"`
	Description   types.String                                                            `tfsdk:"description" json:"description,computed"`
	Enabled       types.Bool                                                              `tfsdk:"enabled" json:"enabled,computed"`
	Filters       customfield.NestedObject[NotificationPoliciesFiltersDataSourceModel]    `tfsdk:"filters" json:"filters,computed"`
	Mechanisms    customfield.NestedObject[NotificationPoliciesMechanismsDataSourceModel] `tfsdk:"mechanisms" json:"mechanisms,computed"`
	Modified      timetypes.RFC3339                                                       `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name          types.String                                                            `tfsdk:"name" json:"name,computed"`
}

type NotificationPoliciesFiltersDataSourceModel struct {
	Actions                      customfield.List[types.String] `tfsdk:"actions" json:"actions,computed"`
	AffectedASNs                 customfield.List[types.String] `tfsdk:"affected_asns" json:"affected_asns,computed"`
	AffectedComponents           customfield.List[types.String] `tfsdk:"affected_components" json:"affected_components,computed"`
	AffectedLocations            customfield.List[types.String] `tfsdk:"affected_locations" json:"affected_locations,computed"`
	AirportCode                  customfield.List[types.String] `tfsdk:"airport_code" json:"airport_code,computed"`
	AlertTriggerPreferences      customfield.List[types.String] `tfsdk:"alert_trigger_preferences" json:"alert_trigger_preferences,computed"`
	AlertTriggerPreferencesValue customfield.List[types.String] `tfsdk:"alert_trigger_preferences_value" json:"alert_trigger_preferences_value,computed"`
	Enabled                      customfield.List[types.String] `tfsdk:"enabled" json:"enabled,computed"`
	Environment                  customfield.List[types.String] `tfsdk:"environment" json:"environment,computed"`
	Event                        customfield.List[types.String] `tfsdk:"event" json:"event,computed"`
	EventSource                  customfield.List[types.String] `tfsdk:"event_source" json:"event_source,computed"`
	EventType                    customfield.List[types.String] `tfsdk:"event_type" json:"event_type,computed"`
	GroupBy                      customfield.List[types.String] `tfsdk:"group_by" json:"group_by,computed"`
	HealthCheckID                customfield.List[types.String] `tfsdk:"health_check_id" json:"health_check_id,computed"`
	IncidentImpact               customfield.List[types.String] `tfsdk:"incident_impact" json:"incident_impact,computed"`
	InputID                      customfield.List[types.String] `tfsdk:"input_id" json:"input_id,computed"`
	Limit                        customfield.List[types.String] `tfsdk:"limit" json:"limit,computed"`
	LogoTag                      customfield.List[types.String] `tfsdk:"logo_tag" json:"logo_tag,computed"`
	MegabitsPerSecond            customfield.List[types.String] `tfsdk:"megabits_per_second" json:"megabits_per_second,computed"`
	NewHealth                    customfield.List[types.String] `tfsdk:"new_health" json:"new_health,computed"`
	NewStatus                    customfield.List[types.String] `tfsdk:"new_status" json:"new_status,computed"`
	PacketsPerSecond             customfield.List[types.String] `tfsdk:"packets_per_second" json:"packets_per_second,computed"`
	PoolID                       customfield.List[types.String] `tfsdk:"pool_id" json:"pool_id,computed"`
	POPName                      customfield.List[types.String] `tfsdk:"pop_name" json:"pop_name,computed"`
	Product                      customfield.List[types.String] `tfsdk:"product" json:"product,computed"`
	ProjectID                    customfield.List[types.String] `tfsdk:"project_id" json:"project_id,computed"`
	Protocol                     customfield.List[types.String] `tfsdk:"protocol" json:"protocol,computed"`
	QueryTag                     customfield.List[types.String] `tfsdk:"query_tag" json:"query_tag,computed"`
	RequestsPerSecond            customfield.List[types.String] `tfsdk:"requests_per_second" json:"requests_per_second,computed"`
	Selectors                    customfield.List[types.String] `tfsdk:"selectors" json:"selectors,computed"`
	Services                     customfield.List[types.String] `tfsdk:"services" json:"services,computed"`
	Slo                          customfield.List[types.String] `tfsdk:"slo" json:"slo,computed"`
	Status                       customfield.List[types.String] `tfsdk:"status" json:"status,computed"`
	TargetHostname               customfield.List[types.String] `tfsdk:"target_hostname" json:"target_hostname,computed"`
	TargetIP                     customfield.List[types.String] `tfsdk:"target_ip" json:"target_ip,computed"`
	TargetZoneName               customfield.List[types.String] `tfsdk:"target_zone_name" json:"target_zone_name,computed"`
	TrafficExclusions            customfield.List[types.String] `tfsdk:"traffic_exclusions" json:"traffic_exclusions,computed"`
	TunnelID                     customfield.List[types.String] `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName                   customfield.List[types.String] `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
	Where                        customfield.List[types.String] `tfsdk:"where" json:"where,computed"`
	Zones                        customfield.List[types.String] `tfsdk:"zones" json:"zones,computed"`
}

type NotificationPoliciesMechanismsDataSourceModel struct {
	Email     customfield.NestedObjectList[NotificationPoliciesMechanismsEmailDataSourceModel]     `tfsdk:"email" json:"email,computed"`
	Pagerduty customfield.NestedObjectList[NotificationPoliciesMechanismsPagerdutyDataSourceModel] `tfsdk:"pagerduty" json:"pagerduty,computed"`
	Webhooks  customfield.NestedObjectList[NotificationPoliciesMechanismsWebhooksDataSourceModel]  `tfsdk:"webhooks" json:"webhooks,computed"`
}

type NotificationPoliciesMechanismsEmailDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type NotificationPoliciesMechanismsPagerdutyDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type NotificationPoliciesMechanismsWebhooksDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
