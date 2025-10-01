// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustGatewayPolicyResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustGatewayPolicyResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustGatewayPolicyResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustGatewayPolicyResource{}
}

// ZeroTrustGatewayPolicyResource defines the resource implementation.
type ZeroTrustGatewayPolicyResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustGatewayPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_gateway_policy"
}

func (r *ZeroTrustGatewayPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ZeroTrustGatewayPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustGatewayPolicyModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustGatewayPolicyResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Gateway.Rules.New(
		ctx,
		zero_trust.GatewayRuleNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustGatewayPolicyModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustGatewayPolicyModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// First, get the current API state to detect drift before applying changes
	currentAPIState, err := r.getCurrentAPIState(ctx, data.ID.ValueString(), data.AccountID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to retrieve current API state for drift detection", err.Error())
		return
	}

	// Detect and report drift between the current API state and planned configuration
	if currentAPIState != nil {
		driftDiags := r.detectDrift(ctx, currentAPIState, data)
		resp.Diagnostics.Append(driftDiags...)
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustGatewayPolicyResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Gateway.Rules.Update(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayRuleUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustGatewayPolicyModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Store the current Terraform state for drift comparison
	currentTerraformState := *data

	res := new(http.Response)
	env := ZeroTrustGatewayPolicyResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Gateway.Rules.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayRuleGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	// Detect drift between the current API state and Terraform state
	// Only compare user-configurable fields (exclude computed-only fields)
	driftDiags := r.detectConfigurationDrift(ctx, data, &currentTerraformState)
	resp.Diagnostics.Append(driftDiags...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustGatewayPolicyModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Gateway.Rules.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.GatewayRuleDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustGatewayPolicyModel = new(ZeroTrustGatewayPolicyModel)

	path_account_id := ""
	path_rule_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<rule_id>",
		&path_account_id,
		&path_rule_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_rule_id)

	res := new(http.Response)
	env := ZeroTrustGatewayPolicyResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Gateway.Rules.Get(
		ctx,
		path_rule_id,
		zero_trust.GatewayRuleGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustGatewayPolicyResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}

// detectDrift compares the current API state with the planned configuration and returns diagnostic messages
// showing the differences between what's configured vs what exists in the API
func (r *ZeroTrustGatewayPolicyResource) detectDrift(ctx context.Context, apiState, plannedConfig *ZeroTrustGatewayPolicyModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if apiState == nil || plannedConfig == nil {
		return diags
	}

	differences := r.compareModels(apiState, plannedConfig)

	if len(differences) > 0 {
		var driftDetails strings.Builder
		driftDetails.WriteString("Configuration drift detected between API state and Terraform configuration:\n\n")

		for _, diff := range differences {
			driftDetails.WriteString(fmt.Sprintf("Field: %s\n", diff.Field))
			driftDetails.WriteString(fmt.Sprintf("  API State:     %+v\n", diff.APIValue))
			driftDetails.WriteString(fmt.Sprintf("  Configuration: %+v\n", diff.ConfigValue))
			driftDetails.WriteString("\n")
		}

		diags.AddWarning(
			"Configuration Drift Detected",
			driftDetails.String(),
		)
	}

	return diags
}

// DriftDifference represents a single field difference between API state and configuration
type DriftDifference struct {
	Field       string
	APIValue    string
	ConfigValue string
}

// compareModels performs a detailed comparison between API state and planned configuration
func (r *ZeroTrustGatewayPolicyResource) compareModels(apiState, plannedConfig *ZeroTrustGatewayPolicyModel) []DriftDifference {
	var differences []DriftDifference

	// Compare basic string fields
	differences = append(differences, r.compareStringField("name", apiState.Name, plannedConfig.Name)...)
	differences = append(differences, r.compareStringField("description", apiState.Description, plannedConfig.Description)...)
	differences = append(differences, r.compareStringField("action", apiState.Action, plannedConfig.Action)...)
	differences = append(differences, r.compareStringField("device_posture", apiState.DevicePosture, plannedConfig.DevicePosture)...)
	differences = append(differences, r.compareStringField("identity", apiState.Identity, plannedConfig.Identity)...)
	differences = append(differences, r.compareStringField("traffic", apiState.Traffic, plannedConfig.Traffic)...)

	// Compare boolean fields
	differences = append(differences, r.compareBoolField("enabled", apiState.Enabled, plannedConfig.Enabled)...)

	// Compare integer fields
	differences = append(differences, r.compareInt64Field("precedence", apiState.Precedence, plannedConfig.Precedence)...)

	// Compare slice fields
	differences = append(differences, r.compareStringSliceField("filters", apiState.Filters, plannedConfig.Filters)...)

	// Compare nested objects
	differences = append(differences, r.compareNestedObjects(apiState, plannedConfig)...)

	return differences
}

// compareStringField compares two string fields and returns differences if they don't match
func (r *ZeroTrustGatewayPolicyResource) compareStringField(fieldName string, apiValue, configValue types.String) []DriftDifference {
	var differences []DriftDifference

	// Handle null/unknown values
	apiStr := ""
	configStr := ""

	if !apiValue.IsNull() && !apiValue.IsUnknown() {
		apiStr = apiValue.ValueString()
	}

	if !configValue.IsNull() && !configValue.IsUnknown() {
		configStr = configValue.ValueString()
	}

	if apiStr != configStr {
		differences = append(differences, DriftDifference{
			Field:       fieldName,
			APIValue:    fmt.Sprintf("'%s'", apiStr),
			ConfigValue: fmt.Sprintf("'%s'", configStr),
		})
	}

	return differences
}

// compareBoolField compares two boolean fields and returns differences if they don't match
func (r *ZeroTrustGatewayPolicyResource) compareBoolField(fieldName string, apiValue, configValue types.Bool) []DriftDifference {
	var differences []DriftDifference

	// Handle null/unknown values
	apiVal := false
	configVal := false

	if !apiValue.IsNull() && !apiValue.IsUnknown() {
		apiVal = apiValue.ValueBool()
	}

	if !configValue.IsNull() && !configValue.IsUnknown() {
		configVal = configValue.ValueBool()
	}

	if apiVal != configVal {
		differences = append(differences, DriftDifference{
			Field:       fieldName,
			APIValue:    fmt.Sprintf("%t", apiVal),
			ConfigValue: fmt.Sprintf("%t", configVal),
		})
	}

	return differences
}

// compareInt64Field compares two int64 fields and returns differences if they don't match
func (r *ZeroTrustGatewayPolicyResource) compareInt64Field(fieldName string, apiValue, configValue types.Int64) []DriftDifference {
	var differences []DriftDifference

	// Handle null/unknown values
	apiVal := int64(0)
	configVal := int64(0)

	if !apiValue.IsNull() && !apiValue.IsUnknown() {
		apiVal = apiValue.ValueInt64()
	}

	if !configValue.IsNull() && !configValue.IsUnknown() {
		configVal = configValue.ValueInt64()
	}

	if apiVal != configVal {
		differences = append(differences, DriftDifference{
			Field:       fieldName,
			APIValue:    fmt.Sprintf("%d", apiVal),
			ConfigValue: fmt.Sprintf("%d", configVal),
		})
	}

	return differences
}

// compareStringSliceField compares two string slice fields and returns differences if they don't match
func (r *ZeroTrustGatewayPolicyResource) compareStringSliceField(fieldName string, apiValue, configValue *[]types.String) []DriftDifference {
	var differences []DriftDifference

	apiSlice := []string{}
	configSlice := []string{}

	if apiValue != nil {
		for _, v := range *apiValue {
			if !v.IsNull() && !v.IsUnknown() {
				apiSlice = append(apiSlice, v.ValueString())
			}
		}
	}

	if configValue != nil {
		for _, v := range *configValue {
			if !v.IsNull() && !v.IsUnknown() {
				configSlice = append(configSlice, v.ValueString())
			}
		}
	}

	if !reflect.DeepEqual(apiSlice, configSlice) {
		differences = append(differences, DriftDifference{
			Field:       fieldName,
			APIValue:    fmt.Sprintf("%v", apiSlice),
			ConfigValue: fmt.Sprintf("%v", configSlice),
		})
	}

	return differences
}

// compareNestedObjects compares nested objects for differences
func (r *ZeroTrustGatewayPolicyResource) compareNestedObjects(apiState, plannedConfig *ZeroTrustGatewayPolicyModel) []DriftDifference {
	var differences []DriftDifference

	// Compare expiration settings
	if !r.isExpirationEqual(apiState.Expiration, plannedConfig.Expiration) {
		differences = append(differences, DriftDifference{
			Field:       "expiration",
			APIValue:    r.formatExpiration(apiState.Expiration),
			ConfigValue: r.formatExpiration(plannedConfig.Expiration),
		})
	}

	// Compare schedule settings
	if !r.isScheduleEqual(apiState.Schedule, plannedConfig.Schedule) {
		differences = append(differences, DriftDifference{
			Field:       "schedule",
			APIValue:    r.formatSchedule(apiState.Schedule),
			ConfigValue: r.formatSchedule(plannedConfig.Schedule),
		})
	}

	// Compare rule settings (simplified comparison for complex nested structure)
	if !r.isRuleSettingsEqual(apiState.RuleSettings, plannedConfig.RuleSettings) {
		differences = append(differences, DriftDifference{
			Field:       "rule_settings",
			APIValue:    "Complex nested object - see detailed logs",
			ConfigValue: "Complex nested object - see detailed logs",
		})
	}

	return differences
}

// Helper functions for nested object comparisons
func (r *ZeroTrustGatewayPolicyResource) isExpirationEqual(api, config customfield.NestedObject[ZeroTrustGatewayPolicyExpirationModel]) bool {
	// Simplified comparison - in a real implementation, you'd want to compare all fields
	return reflect.DeepEqual(api, config)
}

func (r *ZeroTrustGatewayPolicyResource) isScheduleEqual(api, config customfield.NestedObject[ZeroTrustGatewayPolicyScheduleModel]) bool {
	// Simplified comparison - in a real implementation, you'd want to compare all fields
	return reflect.DeepEqual(api, config)
}

func (r *ZeroTrustGatewayPolicyResource) isRuleSettingsEqual(api, config customfield.NestedObject[ZeroTrustGatewayPolicyRuleSettingsModel]) bool {
	// Handle null cases
	if api.IsNull() && config.IsNull() {
		return true
	}
	if api.IsNull() || config.IsNull() {
		return false
	}
	if api.IsUnknown() || config.IsUnknown() {
		return false
	}

	// Get the actual values
	apiValue := api.ValueRuleSettingsPointer()
	configValue := config.ValueRuleSettingsPointer()

	if apiValue == nil && configValue == nil {
		return true
	}
	if apiValue == nil || configValue == nil {
		return false
	}

	// Custom comparison that treats null and false as equivalent for boolean fields
	return r.compareRuleSettingsValues(*apiValue, *configValue)
}

// compareRuleSettingsValues performs a deep comparison of rule settings values,
// treating null and false as equivalent for boolean fields to avoid false drift detection
func (r *ZeroTrustGatewayPolicyResource) compareRuleSettingsValues(api, config ZeroTrustGatewayPolicyRuleSettingsModel) bool {
	// Helper function to compare boolean values where null == false
	compareBooleanField := func(apiVal, configVal types.Bool) bool {
		// If both are null, they're equal
		if apiVal.IsNull() && configVal.IsNull() {
			return true
		}
		// If both are not null, compare their values
		if !apiVal.IsNull() && !configVal.IsNull() {
			return apiVal.Equal(configVal)
		}
		// If one is null and the other is false, they're equivalent
		if apiVal.IsNull() && !configVal.IsNull() {
			return !configVal.ValueBool() // null == false
		}
		if !apiVal.IsNull() && configVal.IsNull() {
			return !apiVal.ValueBool() // null == false
		}
		return false
	}

	// Compare all non-boolean fields using standard equality
	if !api.AddHeaders.Equal(config.AddHeaders) ||
		!api.AllowChildBypass.Equal(config.AllowChildBypass) ||
		!api.AuditSSH.Equal(config.AuditSSH) ||
		!api.BlockPage.Equal(config.BlockPage) ||
		!api.BlockPageEnabled.Equal(config.BlockPageEnabled) ||
		!api.BlockReason.Equal(config.BlockReason) ||
		!api.BypassParentRule.Equal(config.BypassParentRule) ||
		!api.CheckSession.Equal(config.CheckSession) ||
		!api.DNSResolvers.Equal(config.DNSResolvers) ||
		!api.Egress.Equal(config.Egress) ||
		!api.IgnoreCnameCategoryMatches.Equal(config.IgnoreCnameCategoryMatches) ||
		!api.InsecureDisableDNSSECValidation.Equal(config.InsecureDisableDNSSECValidation) ||
		!api.IPCategories.Equal(config.IPCategories) ||
		!api.IPIndicatorFeeds.Equal(config.IPIndicatorFeeds) ||
		!api.L4Override.Equal(config.L4Override) ||
		!api.NotificationSettings.Equal(config.NotificationSettings) ||
		!api.OverrideHost.Equal(config.OverrideHost) ||
		!api.OverrideIPs.Equal(config.OverrideIPs) ||
		!api.PayloadLog.Equal(config.PayloadLog) ||
		!api.Quarantine.Equal(config.Quarantine) ||
		!api.Redirect.Equal(config.Redirect) ||
		!api.ResolveDNSInternally.Equal(config.ResolveDNSInternally) ||
		!api.ResolveDNSThroughCloudflare.Equal(config.ResolveDNSThroughCloudflare) ||
		!api.UntrustedCert.Equal(config.UntrustedCert) {
		return false
	}

	// Special handling for BisoAdminControls which contains boolean fields that may be null
	if !api.BisoAdminControls.IsNull() || !config.BisoAdminControls.IsNull() {
		if api.BisoAdminControls.IsNull() && !config.BisoAdminControls.IsNull() {
			return false
		}
		if !api.BisoAdminControls.IsNull() && config.BisoAdminControls.IsNull() {
			return false
		}
		if !api.BisoAdminControls.IsNull() && !config.BisoAdminControls.IsNull() {
			apiBAC := api.BisoAdminControls.ValueRuleSettingsBisoAdminControlsPointer()
			configBAC := config.BisoAdminControls.ValueRuleSettingsBisoAdminControlsPointer()

			if apiBAC == nil && configBAC == nil {
				return true
			}
			if apiBAC == nil || configBAC == nil {
				return false
			}

			// Compare non-boolean fields normally
			if !apiBAC.Copy.Equal(configBAC.Copy) ||
				!apiBAC.Download.Equal(configBAC.Download) ||
				!apiBAC.Keyboard.Equal(configBAC.Keyboard) ||
				!apiBAC.Paste.Equal(configBAC.Paste) ||
				!apiBAC.Printing.Equal(configBAC.Printing) ||
				!apiBAC.Upload.Equal(configBAC.Upload) ||
				!apiBAC.Version.Equal(configBAC.Version) {
				return false
			}

			// Compare boolean fields with null == false logic
			if !compareBooleanField(apiBAC.DCP, configBAC.DCP) ||
				!compareBooleanField(apiBAC.DD, configBAC.DD) ||
				!compareBooleanField(apiBAC.DK, configBAC.DK) ||
				!compareBooleanField(apiBAC.DP, configBAC.DP) ||
				!compareBooleanField(apiBAC.DU, configBAC.DU) {
				return false
			}
		}
	}

	return true
}

func (r *ZeroTrustGatewayPolicyResource) formatExpiration(exp customfield.NestedObject[ZeroTrustGatewayPolicyExpirationModel]) string {
	if exp.IsNull() || exp.IsUnknown() {
		return "null"
	}
	return "configured"
}

func (r *ZeroTrustGatewayPolicyResource) formatSchedule(sched customfield.NestedObject[ZeroTrustGatewayPolicyScheduleModel]) string {
	if sched.IsNull() || sched.IsUnknown() {
		return "null"
	}
	return "configured"
}

// getCurrentAPIState retrieves the current state of the resource from the API
func (r *ZeroTrustGatewayPolicyResource) getCurrentAPIState(ctx context.Context, ruleID, accountID string) (*ZeroTrustGatewayPolicyModel, error) {
	res := new(http.Response)
	var data ZeroTrustGatewayPolicyModel
	env := ZeroTrustGatewayPolicyResultEnvelope{data}

	_, err := r.client.ZeroTrust.Gateway.Rules.Get(
		ctx,
		ruleID,
		zero_trust.GatewayRuleGetParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current API state: %w", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to retrieve current API state: %w", err)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize API response: %w", err)
	}

	return &env.Result, nil
}

// detectConfigurationDrift compares API state with Terraform state, focusing only on user-configurable fields
// This is used during Read operations to detect drift in the configuration
func (r *ZeroTrustGatewayPolicyResource) detectConfigurationDrift(ctx context.Context, apiState, terraformState *ZeroTrustGatewayPolicyModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if apiState == nil || terraformState == nil {
		return diags
	}

	// Only compare user-configurable fields (exclude computed-only fields like ID, timestamps, etc.)
	differences := r.compareUserConfigurableFields(apiState, terraformState)

	if len(differences) > 0 {
		var driftDetails strings.Builder
		driftDetails.WriteString("Configuration drift detected! The actual API state differs from your Terraform configuration:\n\n")
		driftDetails.WriteString("This may indicate that changes were made outside of Terraform.\n")
		driftDetails.WriteString("Consider updating your configuration or running 'terraform apply' to reconcile.\n\n")

		for _, diff := range differences {
			driftDetails.WriteString(fmt.Sprintf("Field: %s\n", diff.Field))
			driftDetails.WriteString(fmt.Sprintf("  Current API State:        %#v\n", diff.APIValue))
			driftDetails.WriteString(fmt.Sprintf("  Your Terraform Config:    %#v\n", diff.ConfigValue))
			driftDetails.WriteString("\n")
		}

		// Show detailed RuleSettings comparison
		driftDetails.WriteString("=== RuleSettings Detailed Comparison ===\n")
		driftDetails.WriteString(fmt.Sprintf("API State RuleSettings:        %s\n", apiState.RuleSettings.String()))
		driftDetails.WriteString(fmt.Sprintf("Terraform Config RuleSettings: %s\n", terraformState.RuleSettings.String()))

		// Show the diff if they're different
		if apiState.RuleSettings.String() != terraformState.RuleSettings.String() {
			driftDetails.WriteString("\nDifference (- API State, + Terraform Config):\n")
			driftDetails.WriteString(cmp.Diff(apiState.RuleSettings.String(), terraformState.RuleSettings.String()))
		} else {
			driftDetails.WriteString("\nRuleSettings are identical.\n")
		}
		driftDetails.WriteString("==========================================\n")

		diags.AddWarning(
			"Configuration Drift Detected",
			driftDetails.String(),
		)
	}

	return diags
}

// compareUserConfigurableFields compares only the fields that users can configure (excludes computed-only fields)
func (r *ZeroTrustGatewayPolicyResource) compareUserConfigurableFields(apiState, terraformState *ZeroTrustGatewayPolicyModel) []DriftDifference {
	var differences []DriftDifference

	// Compare user-configurable string fields
	differences = append(differences, r.compareStringField("name", apiState.Name, terraformState.Name)...)
	differences = append(differences, r.compareStringField("description", apiState.Description, terraformState.Description)...)
	differences = append(differences, r.compareStringField("action", apiState.Action, terraformState.Action)...)
	differences = append(differences, r.compareStringField("device_posture", apiState.DevicePosture, terraformState.DevicePosture)...)
	differences = append(differences, r.compareStringField("identity", apiState.Identity, terraformState.Identity)...)
	differences = append(differences, r.compareStringField("traffic", apiState.Traffic, terraformState.Traffic)...)

	// Compare user-configurable boolean fields
	differences = append(differences, r.compareBoolField("enabled", apiState.Enabled, terraformState.Enabled)...)

	// Compare user-configurable integer fields
	differences = append(differences, r.compareInt64Field("precedence", apiState.Precedence, terraformState.Precedence)...)

	// Compare user-configurable slice fields
	differences = append(differences, r.compareStringSliceField("filters", apiState.Filters, terraformState.Filters)...)

	// Compare user-configurable nested objects
	differences = append(differences, r.compareUserConfigurableNestedObjects(apiState, terraformState)...)

	return differences
}

// compareUserConfigurableNestedObjects compares nested objects that users can configure
func (r *ZeroTrustGatewayPolicyResource) compareUserConfigurableNestedObjects(apiState, terraformState *ZeroTrustGatewayPolicyModel) []DriftDifference {
	var differences []DriftDifference

	// Compare expiration settings (user-configurable)
	if !r.isExpirationEqual(apiState.Expiration, terraformState.Expiration) {
		differences = append(differences, DriftDifference{
			Field:       "expiration",
			APIValue:    r.formatExpiration(apiState.Expiration),
			ConfigValue: r.formatExpiration(terraformState.Expiration),
		})
	}

	// Compare schedule settings (user-configurable)
	if !r.isScheduleEqual(apiState.Schedule, terraformState.Schedule) {
		differences = append(differences, DriftDifference{
			Field:       "schedule",
			APIValue:    r.formatSchedule(apiState.Schedule),
			ConfigValue: r.formatSchedule(terraformState.Schedule),
		})
	}

	// Compare rule settings (user-configurable)
	if !r.isRuleSettingsEqual(apiState.RuleSettings, terraformState.RuleSettings) {
		differences = append(differences, DriftDifference{
			Field:       "rule_settings",
			APIValue:    PrettyPrint(apiState.RuleSettings),
			ConfigValue: PrettyPrint(terraformState.RuleSettings),
		})
	}

	return differences
}

func PrettyPrint(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {

		return "error marshalling json"
	}
	return string(bytes)
}
