// File generated for v4 to v5 state migration

package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x, SDK v2)
// ============================================================================

// SourceCloudflareNotificationPolicyModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (SDK v2 implicit version)
// Resource type: cloudflare_notification_policy
//
// This model matches the v4 SDK v2 schema structure, including:
// - filters as TypeList MaxItems:1 (stored as array in state)
// - Three separate integration Set fields
// - Integration items with both id and name fields
// - All filter fields as TypeSet
type SourceCloudflareNotificationPolicyModel struct {
	ID                   types.String `tfsdk:"id"`
	AccountID            types.String `tfsdk:"account_id"`
	Name                 types.String `tfsdk:"name"`
	AlertType            types.String `tfsdk:"alert_type"`
	Description          types.String `tfsdk:"description"`
	Enabled              types.Bool   `tfsdk:"enabled"`
	Created              types.String `tfsdk:"created"`  // String in v4, RFC3339 in v5
	Modified             types.String `tfsdk:"modified"` // String in v4, RFC3339 in v5
	Filters              types.List   `tfsdk:"filters"`  // List of SourceFiltersModel (MaxItems:1, stored as array)
	EmailIntegration     types.Set    `tfsdk:"email_integration"`     // Set of SourceIntegrationModel
	WebhooksIntegration  types.Set    `tfsdk:"webhooks_integration"`  // Set of SourceIntegrationModel
	PagerdutyIntegration types.Set    `tfsdk:"pagerduty_integration"` // Set of SourceIntegrationModel
}

// SourceFiltersModel represents the filters structure from v4.x provider.
// In v4, this was TypeList MaxItems:1, stored as array in state.
// All fields are TypeSet in v4 schema.
type SourceFiltersModel struct {
	Actions                  types.Set `tfsdk:"actions"`
	AirportCode              types.Set `tfsdk:"airport_code"`
	AffectedComponents       types.Set `tfsdk:"affected_components"`
	Status                   types.Set `tfsdk:"status"`
	HealthCheckID            types.Set `tfsdk:"health_check_id"`
	Zones                    types.Set `tfsdk:"zones"`
	Services                 types.Set `tfsdk:"services"`
	Product                  types.Set `tfsdk:"product"`
	Limit                    types.Set `tfsdk:"limit"`
	Enabled                  types.Set `tfsdk:"enabled"`
	PoolID                   types.Set `tfsdk:"pool_id"`
	Slo                      types.Set `tfsdk:"slo"`
	Where                    types.Set `tfsdk:"where"`
	GroupBy                  types.Set `tfsdk:"group_by"`
	AlertTriggerPreferences  types.Set `tfsdk:"alert_trigger_preferences"`
	RequestsPerSecond        types.Set `tfsdk:"requests_per_second"`
	TargetZoneName           types.Set `tfsdk:"target_zone_name"`
	TargetHostname           types.Set `tfsdk:"target_hostname"`
	TargetIP                 types.Set `tfsdk:"target_ip"`
	PacketsPerSecond         types.Set `tfsdk:"packets_per_second"`
	Protocol                 types.Set `tfsdk:"protocol"`
	ProjectID                types.Set `tfsdk:"project_id"`
	Environment              types.Set `tfsdk:"environment"`
	Event                    types.Set `tfsdk:"event"`
	EventSource              types.Set `tfsdk:"event_source"`
	NewHealth                types.Set `tfsdk:"new_health"`
	InputID                  types.Set `tfsdk:"input_id"`
	EventType                types.Set `tfsdk:"event_type"`
	MegabitsPerSecond        types.Set `tfsdk:"megabits_per_second"`
	IncidentImpact           types.Set `tfsdk:"incident_impact"`
	NewStatus                types.Set `tfsdk:"new_status"`
	Selectors                types.Set `tfsdk:"selectors"`
	TunnelID                 types.Set `tfsdk:"tunnel_id"`
	TunnelName               types.Set `tfsdk:"tunnel_name"`
}

// SourceIntegrationModel represents the integration structure from v4.x provider.
// Used for email_integration, webhooks_integration, and pagerduty_integration.
// Note: The "name" field is present in v4 but dropped in v5.
type SourceIntegrationModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Present in v4, dropped in v5
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================
//
// Note: We import types from the parent package to ensure consistency.
// These are defined here to keep the migration package self-contained,
// but they should match the structures in the parent model.go file.

// TargetNotificationPolicyModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_notification_policy
type TargetNotificationPolicyModel struct {
	ID            types.String                                `tfsdk:"id"`
	AccountID     types.String                                `tfsdk:"account_id"`
	AlertType     types.String                                `tfsdk:"alert_type"`
	Name          types.String                                `tfsdk:"name"`
	Mechanisms    *TargetNotificationPolicyMechanismsModel    `tfsdk:"mechanisms"` // New nested structure
	AlertInterval types.String                                `tfsdk:"alert_interval"` // New in v5
	Description   types.String                                `tfsdk:"description"`
	Filters       *TargetNotificationPolicyFiltersModel       `tfsdk:"filters"` // SingleNestedAttribute in v5
	Enabled       types.Bool                                  `tfsdk:"enabled"`
	Created       timetypes.RFC3339                           `tfsdk:"created"`  // RFC3339 type in v5
	Modified      timetypes.RFC3339                           `tfsdk:"modified"` // RFC3339 type in v5
}

// TargetNotificationPolicyMechanismsModel represents the mechanisms structure from v5.x+ provider.
// This is a new structure in v5 that consolidates the three separate integration fields from v4.
type TargetNotificationPolicyMechanismsModel struct {
	Email     customfield.NestedObjectSet[TargetNotificationPolicyMechanismsEmailModel]     `tfsdk:"email"`
	Pagerduty customfield.NestedObjectSet[TargetNotificationPolicyMechanismsPagerdutyModel] `tfsdk:"pagerduty"`
	Webhooks  customfield.NestedObjectSet[TargetNotificationPolicyMechanismsWebhooksModel]  `tfsdk:"webhooks"`
}

// TargetNotificationPolicyMechanismsEmailModel represents email mechanism item.
// Note: Only has "id" field - "name" field was dropped from v4.
type TargetNotificationPolicyMechanismsEmailModel struct {
	ID types.String `tfsdk:"id"`
}

// TargetNotificationPolicyMechanismsPagerdutyModel represents pagerduty mechanism item.
// Note: Only has "id" field - "name" field was dropped from v4.
type TargetNotificationPolicyMechanismsPagerdutyModel struct {
	ID types.String `tfsdk:"id"`
}

// TargetNotificationPolicyMechanismsWebhooksModel represents webhooks mechanism item.
// Note: Only has "id" field - "name" field was dropped from v4.
type TargetNotificationPolicyMechanismsWebhooksModel struct {
	ID types.String `tfsdk:"id"`
}

// TargetNotificationPolicyFiltersModel represents the filters structure from v5.x+ provider.
// In v5, this is SingleNestedAttribute (object) instead of MaxItems:1 list.
// All fields are List (slices) in v5 instead of Set.
type TargetNotificationPolicyFiltersModel struct {
	Actions                      *[]types.String `tfsdk:"actions"`
	AffectedASNs                 *[]types.String `tfsdk:"affected_asns"`                   // New in v5
	AffectedComponents           *[]types.String `tfsdk:"affected_components"`
	AffectedLocations            *[]types.String `tfsdk:"affected_locations"`              // New in v5
	AirportCode                  *[]types.String `tfsdk:"airport_code"`
	AlertTriggerPreferences      *[]types.String `tfsdk:"alert_trigger_preferences"`
	AlertTriggerPreferencesValue *[]types.String `tfsdk:"alert_trigger_preferences_value"` // New in v5
	Enabled                      *[]types.String `tfsdk:"enabled"`
	Environment                  *[]types.String `tfsdk:"environment"`
	Event                        *[]types.String `tfsdk:"event"`
	EventSource                  *[]types.String `tfsdk:"event_source"`
	EventType                    *[]types.String `tfsdk:"event_type"`
	GroupBy                      *[]types.String `tfsdk:"group_by"`
	HealthCheckID                *[]types.String `tfsdk:"health_check_id"`
	IncidentImpact               *[]types.String `tfsdk:"incident_impact"`
	InputID                      *[]types.String `tfsdk:"input_id"`
	InsightClass                 *[]types.String `tfsdk:"insight_class"` // New in v5
	Limit                        *[]types.String `tfsdk:"limit"`
	LogoTag                      *[]types.String `tfsdk:"logo_tag"`              // New in v5
	MegabitsPerSecond            *[]types.String `tfsdk:"megabits_per_second"`
	NewHealth                    *[]types.String `tfsdk:"new_health"`
	NewStatus                    *[]types.String `tfsdk:"new_status"`
	PacketsPerSecond             *[]types.String `tfsdk:"packets_per_second"`
	PoolID                       *[]types.String `tfsdk:"pool_id"`
	POPNames                     *[]types.String `tfsdk:"pop_names"`     // New in v5
	Product                      *[]types.String `tfsdk:"product"`
	ProjectID                    *[]types.String `tfsdk:"project_id"`
	Protocol                     *[]types.String `tfsdk:"protocol"`
	QueryTag                     *[]types.String `tfsdk:"query_tag"`     // New in v5
	RequestsPerSecond            *[]types.String `tfsdk:"requests_per_second"`
	Selectors                    *[]types.String `tfsdk:"selectors"`
	Services                     *[]types.String `tfsdk:"services"`
	Slo                          *[]types.String `tfsdk:"slo"`
	Status                       *[]types.String `tfsdk:"status"`
	TargetHostname               *[]types.String `tfsdk:"target_hostname"`
	TargetIP                     *[]types.String `tfsdk:"target_ip"`
	TargetZoneName               *[]types.String `tfsdk:"target_zone_name"`
	TrafficExclusions            *[]types.String `tfsdk:"traffic_exclusions"` // New in v5
	TunnelID                     *[]types.String `tfsdk:"tunnel_id"`
	TunnelName                   *[]types.String `tfsdk:"tunnel_name"`
	Type                         *[]types.String `tfsdk:"type"`  // New in v5
	Where                        *[]types.String `tfsdk:"where"`
	Zones                        *[]types.String `tfsdk:"zones"`
}
