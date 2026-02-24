// File generated for v4 to v5 state migration

package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function handles the complete v4 → v5 transformation.
func Transform(ctx context.Context, source SourceCloudflareNotificationPolicyModel) (*TargetNotificationPolicyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for notification_policy migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for notification_policy migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.AlertType.IsNull() || source.AlertType.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"alert_type is required for notification_policy migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetNotificationPolicyModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Name:        source.Name,
		AlertType:   source.AlertType,
		Description: source.Description,
		Enabled:     source.Enabled,
	}

	// Step 3: Handle new v5 fields
	// alert_interval is new in v5, set to Null
	target.AlertInterval = types.StringNull()

	// Step 4: Handle timestamp conversions (String → RFC3339)
	if !source.Created.IsNull() && !source.Created.IsUnknown() {
		target.Created = timetypes.NewRFC3339ValueMust(source.Created.ValueString())
	} else {
		target.Created = timetypes.NewRFC3339Null()
	}

	if !source.Modified.IsNull() && !source.Modified.IsUnknown() {
		target.Modified = timetypes.NewRFC3339ValueMust(source.Modified.ValueString())
	} else {
		target.Modified = timetypes.NewRFC3339Null()
	}

	// Step 5: Transform filters (List MaxItems:1 → SingleNestedAttribute)
	target.Filters = transformFilters(ctx, source.Filters, &diags)

	// Step 6: Build mechanisms from integration Sets
	target.Mechanisms = buildMechanisms(ctx, source.EmailIntegration, source.WebhooksIntegration, source.PagerdutyIntegration, &diags)

	if diags.HasError() {
		return nil, diags
	}

	return target, diags
}

// transformFilters converts v4 filters list (MaxItems:1) to v5 filters object.
// In v4, filters was TypeList MaxItems:1 (stored as array).
// In v5, filters is SingleNestedAttribute (stored as object).
func transformFilters(ctx context.Context, sourceFilters types.List, diags *diag.Diagnostics) *TargetNotificationPolicyFiltersModel {
	// Check if filters is null or has no elements
	if sourceFilters.IsNull() || sourceFilters.IsUnknown() {
		return nil
	}

	// Extract filters array elements
	var sourceFiltersList []SourceFiltersModel
	diags.Append(sourceFilters.ElementsAs(ctx, &sourceFiltersList, false)...)
	if diags.HasError() || len(sourceFiltersList) == 0 {
		return nil
	}

	// Get the first (and only) element (MaxItems:1)
	sourceFilter := sourceFiltersList[0]

	// Initialize target filters
	targetFilters := &TargetNotificationPolicyFiltersModel{}

	// Convert each Set field to *[]types.String
	// This is repetitive but straightforward
	targetFilters.Actions = convertSetToStringSlice(ctx, sourceFilter.Actions, diags)
	targetFilters.AirportCode = convertSetToStringSlice(ctx, sourceFilter.AirportCode, diags)
	targetFilters.AffectedComponents = convertSetToStringSlice(ctx, sourceFilter.AffectedComponents, diags)
	targetFilters.Status = convertSetToStringSlice(ctx, sourceFilter.Status, diags)
	targetFilters.HealthCheckID = convertSetToStringSlice(ctx, sourceFilter.HealthCheckID, diags)
	targetFilters.Zones = convertSetToStringSlice(ctx, sourceFilter.Zones, diags)
	targetFilters.Services = convertSetToStringSlice(ctx, sourceFilter.Services, diags)
	targetFilters.Product = convertSetToStringSlice(ctx, sourceFilter.Product, diags)
	targetFilters.Limit = convertSetToStringSlice(ctx, sourceFilter.Limit, diags)
	targetFilters.Enabled = convertSetToStringSlice(ctx, sourceFilter.Enabled, diags)
	targetFilters.PoolID = convertSetToStringSlice(ctx, sourceFilter.PoolID, diags)
	targetFilters.Slo = convertSetToStringSlice(ctx, sourceFilter.Slo, diags)
	targetFilters.Where = convertSetToStringSlice(ctx, sourceFilter.Where, diags)
	targetFilters.GroupBy = convertSetToStringSlice(ctx, sourceFilter.GroupBy, diags)
	targetFilters.AlertTriggerPreferences = convertSetToStringSlice(ctx, sourceFilter.AlertTriggerPreferences, diags)
	targetFilters.RequestsPerSecond = convertSetToStringSlice(ctx, sourceFilter.RequestsPerSecond, diags)
	targetFilters.TargetZoneName = convertSetToStringSlice(ctx, sourceFilter.TargetZoneName, diags)
	targetFilters.TargetHostname = convertSetToStringSlice(ctx, sourceFilter.TargetHostname, diags)
	targetFilters.TargetIP = convertSetToStringSlice(ctx, sourceFilter.TargetIP, diags)
	targetFilters.PacketsPerSecond = convertSetToStringSlice(ctx, sourceFilter.PacketsPerSecond, diags)
	targetFilters.Protocol = convertSetToStringSlice(ctx, sourceFilter.Protocol, diags)
	targetFilters.ProjectID = convertSetToStringSlice(ctx, sourceFilter.ProjectID, diags)
	targetFilters.Environment = convertSetToStringSlice(ctx, sourceFilter.Environment, diags)
	targetFilters.Event = convertSetToStringSlice(ctx, sourceFilter.Event, diags)
	targetFilters.EventSource = convertSetToStringSlice(ctx, sourceFilter.EventSource, diags)
	targetFilters.NewHealth = convertSetToStringSlice(ctx, sourceFilter.NewHealth, diags)
	targetFilters.InputID = convertSetToStringSlice(ctx, sourceFilter.InputID, diags)
	targetFilters.EventType = convertSetToStringSlice(ctx, sourceFilter.EventType, diags)
	targetFilters.MegabitsPerSecond = convertSetToStringSlice(ctx, sourceFilter.MegabitsPerSecond, diags)
	targetFilters.IncidentImpact = convertSetToStringSlice(ctx, sourceFilter.IncidentImpact, diags)
	targetFilters.NewStatus = convertSetToStringSlice(ctx, sourceFilter.NewStatus, diags)
	targetFilters.Selectors = convertSetToStringSlice(ctx, sourceFilter.Selectors, diags)
	targetFilters.TunnelID = convertSetToStringSlice(ctx, sourceFilter.TunnelID, diags)
	targetFilters.TunnelName = convertSetToStringSlice(ctx, sourceFilter.TunnelName, diags)

	// New v5 filter fields default to nil (not present in v4)
	targetFilters.AffectedASNs = nil
	targetFilters.AffectedLocations = nil
	targetFilters.AlertTriggerPreferencesValue = nil
	targetFilters.InsightClass = nil
	targetFilters.LogoTag = nil
	targetFilters.POPNames = nil
	targetFilters.QueryTag = nil
	targetFilters.TrafficExclusions = nil
	targetFilters.Type = nil

	return targetFilters
}

// convertSetToStringSlice converts a types.Set to *[]types.String.
// Returns nil if the set is null or unknown.
func convertSetToStringSlice(ctx context.Context, set types.Set, diags *diag.Diagnostics) *[]types.String {
	if set.IsNull() || set.IsUnknown() {
		return nil
	}

	// Extract to native Go []string first
	var rawStrings []string
	diags.Append(set.ElementsAs(ctx, &rawStrings, false)...)
	if diags.HasError() || len(rawStrings) == 0 {
		return nil
	}

	// Convert to []types.String
	result := make([]types.String, len(rawStrings))
	for i, str := range rawStrings {
		result[i] = types.StringValue(str)
	}

	return &result
}

// buildMechanisms constructs the v5 mechanisms object from v4 integration Sets.
// In v4: email_integration, webhooks_integration, pagerduty_integration (separate Sets)
// In v5: mechanisms object with email, webhooks, pagerduty (nested Sets)
// Note: The "name" field is dropped - only "id" is preserved.
func buildMechanisms(
	ctx context.Context,
	emailIntegration types.Set,
	webhooksIntegration types.Set,
	pagerdutyIntegration types.Set,
	diags *diag.Diagnostics,
) *TargetNotificationPolicyMechanismsModel {
	mechanisms := &TargetNotificationPolicyMechanismsModel{}

	// Convert email integration
	if !emailIntegration.IsNull() && !emailIntegration.IsUnknown() {
		var sourceEmailItems []SourceIntegrationModel
		diags.Append(emailIntegration.ElementsAs(ctx, &sourceEmailItems, false)...)
		if !diags.HasError() && len(sourceEmailItems) > 0 {
			emailMechanisms := make([]TargetNotificationPolicyMechanismsEmailModel, len(sourceEmailItems))
			for i, item := range sourceEmailItems {
				emailMechanisms[i] = TargetNotificationPolicyMechanismsEmailModel{
					ID: item.ID, // Drop item.Name
				}
			}
			emailSet, emailDiags := customfield.NewObjectSet(ctx, emailMechanisms)
			diags.Append(emailDiags...)
			mechanisms.Email = emailSet
		} else {
			mechanisms.Email = customfield.NullObjectSet[TargetNotificationPolicyMechanismsEmailModel](ctx)
		}
	} else {
		mechanisms.Email = customfield.NullObjectSet[TargetNotificationPolicyMechanismsEmailModel](ctx)
	}

	// Convert webhooks integration
	if !webhooksIntegration.IsNull() && !webhooksIntegration.IsUnknown() {
		var sourceWebhooksItems []SourceIntegrationModel
		diags.Append(webhooksIntegration.ElementsAs(ctx, &sourceWebhooksItems, false)...)
		if !diags.HasError() && len(sourceWebhooksItems) > 0 {
			webhooksMechanisms := make([]TargetNotificationPolicyMechanismsWebhooksModel, len(sourceWebhooksItems))
			for i, item := range sourceWebhooksItems {
				webhooksMechanisms[i] = TargetNotificationPolicyMechanismsWebhooksModel{
					ID: item.ID, // Drop item.Name
				}
			}
			webhooksSet, webhooksDiags := customfield.NewObjectSet(ctx, webhooksMechanisms)
			diags.Append(webhooksDiags...)
			mechanisms.Webhooks = webhooksSet
		} else {
			mechanisms.Webhooks = customfield.NullObjectSet[TargetNotificationPolicyMechanismsWebhooksModel](ctx)
		}
	} else {
		mechanisms.Webhooks = customfield.NullObjectSet[TargetNotificationPolicyMechanismsWebhooksModel](ctx)
	}

	// Convert pagerduty integration
	if !pagerdutyIntegration.IsNull() && !pagerdutyIntegration.IsUnknown() {
		var sourcePagerdutyItems []SourceIntegrationModel
		diags.Append(pagerdutyIntegration.ElementsAs(ctx, &sourcePagerdutyItems, false)...)
		if !diags.HasError() && len(sourcePagerdutyItems) > 0 {
			pagerdutyMechanisms := make([]TargetNotificationPolicyMechanismsPagerdutyModel, len(sourcePagerdutyItems))
			for i, item := range sourcePagerdutyItems {
				pagerdutyMechanisms[i] = TargetNotificationPolicyMechanismsPagerdutyModel{
					ID: item.ID, // Drop item.Name
				}
			}
			pagerdutySet, pagerdutyDiags := customfield.NewObjectSet(ctx, pagerdutyMechanisms)
			diags.Append(pagerdutyDiags...)
			mechanisms.Pagerduty = pagerdutySet
		} else {
			mechanisms.Pagerduty = customfield.NullObjectSet[TargetNotificationPolicyMechanismsPagerdutyModel](ctx)
		}
	} else {
		mechanisms.Pagerduty = customfield.NullObjectSet[TargetNotificationPolicyMechanismsPagerdutyModel](ctx)
	}

	return mechanisms
}
