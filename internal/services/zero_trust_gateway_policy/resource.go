// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
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
		driftDiags := r.detectDriftOnUpdate(ctx, currentAPIState, data)
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
	driftDiags := r.detectConfigurationDriftOnRead(ctx, data, &currentTerraformState)
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

// detectDriftOnUpdate compares the current API state with the planned configuration and returns diagnostic messages
// showing the differences between what's configured vs what exists in the API
func (r *ZeroTrustGatewayPolicyResource) detectDriftOnUpdate(ctx context.Context, apiState, plannedConfig *ZeroTrustGatewayPolicyModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if apiState == nil || plannedConfig == nil {
		return diags
	}

	differences := r.compareModels(apiState, plannedConfig)

	if len(differences) > 0 {
		var driftDetails strings.Builder
		driftDetails.WriteString("Configuration drift detected between API state and Terraform configuration.\n\n")
		// We expect a single consolidated difference (policy JSON)
		diff := differences[0]
		driftDetails.WriteString("detectDriftOnUpdate@UPDATE")
		driftDetails.WriteString("Side-by-side (Terraform | API):\n")
		driftDetails.WriteString(sideBySideDiff(diff.ConfigValue, diff.APIValue, 60))

		driftDetails.WriteString("\n\nTo fix the drift, update your terraform declaration to match the current API state.")

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
	// Marshal using model's JSON to respect tags/omissions
	apiBytes, errA := apiState.MarshalJSON()
	cfgBytes, errC := plannedConfig.MarshalJSON()

	if errA != nil || errC != nil {
		return []DriftDifference{
			{
				Field:       "policy",
				APIValue:    fmt.Sprintf("error marshalling api state: %v", errA),
				ConfigValue: fmt.Sprintf("error marshalling config: %v", errC),
			},
		}
	}

	// Normalize JSON for stable comparison
	var apiObj any
	var cfgObj any
	if err := json.Unmarshal(apiBytes, &apiObj); err != nil {
		return []DriftDifference{{Field: "policy", APIValue: string(apiBytes), ConfigValue: string(cfgBytes)}}
	}
	if err := json.Unmarshal(cfgBytes, &cfgObj); err != nil {
		return []DriftDifference{{Field: "policy", APIValue: string(apiBytes), ConfigValue: string(cfgBytes)}}
	}
	normAPI, _ := json.MarshalIndent(apiObj, "", "  ")
	normCfg, _ := json.MarshalIndent(cfgObj, "", "  ")

	if string(normAPI) == string(normCfg) {
		return nil
	}

	return []DriftDifference{{Field: "policy", APIValue: string(normAPI), ConfigValue: string(normCfg)}}
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

// This is used during Read operations to detect drift in the configuration
func (r *ZeroTrustGatewayPolicyResource) detectConfigurationDriftOnRead(ctx context.Context, apiState, terraformState *ZeroTrustGatewayPolicyModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if apiState == nil || terraformState == nil {
		return diags
	}

	// Serialize both objects using model-aware JSON marshaling
	apiBytes, errA := apiState.MarshalJSON()
	cfgBytes, errC := terraformState.MarshalJSON()

	if errA != nil || errC != nil {
		diags.AddWarning(
			"Configuration Drift Detected",
			fmt.Sprintf("error marshalling for drift detection (api=%v, config=%v)", errA, errC),
		)
		return diags
	}

	// Normalize: drop computed-only fields, then pretty-print
	var apiMap map[string]any
	var cfgMap map[string]any
	if err := json.Unmarshal(apiBytes, &apiMap); err != nil {
		diags.AddWarning("Configuration Drift Detected", "failed to parse API state for drift detection")
		return diags
	}
	if err := json.Unmarshal(cfgBytes, &cfgMap); err != nil {
		diags.AddWarning("Configuration Drift Detected", "failed to parse Terraform state for drift detection")
		return diags
	}

	// Remove computed-only keys to focus on user-configurable fields
	remove := func(m map[string]any) {
		delete(m, "id")
		delete(m, "account_id")
		delete(m, "created_at")
		delete(m, "deleted_at")
		delete(m, "read_only")
		delete(m, "sharable")
		delete(m, "source_account")
		delete(m, "updated_at")
		delete(m, "version")
		delete(m, "warning_status")
	}
	remove(apiMap)
	remove(cfgMap)

	normAPI, _ := json.MarshalIndent(apiMap, "", "  ")
	normCfg, _ := json.MarshalIndent(cfgMap, "", "  ")

	if string(normAPI) == string(normCfg) {
		return diags
	}

	var msg strings.Builder

	msg.WriteString("detectConfigurationDriftOnRead@READ")

	msg.WriteString("Configuration drift detected! The actual API state differs from your Terraform configuration.\n\n")
	msg.WriteString("Side-by-side (Terraform | API) for user-configurable fields:\n")
	msg.WriteString(sideBySideDiff(string(normCfg), string(normAPI), 60))

	msg.WriteString("\n\nTo fix the drift, update your terraform declaration to match the current API state.")

	diags.AddWarning("Configuration Drift Detected", msg.String())
	return diags
}

// sideBySideDiff renders a simple side-by-side view of two multi-line strings.
// leftWidth controls the width of the left column (API). Differences are marked with '≠'.
func sideBySideDiff(right, left string, leftWidth int) string {
	// Intentionally interpret first arg as Terraform (left), second as API (right)
	lnsL := strings.Split(right, "\n") // Terraform
	lnsR := strings.Split(left, "\n")  // API
	n := len(lnsL)
	if len(lnsR) > n {
		n = len(lnsR)
	}
	rightWidth := leftWidth
	var b strings.Builder

	// Header (Terraform left, API right)
	b.WriteString(fmt.Sprintf("%4s %-*s   %4s %s\n", "#", leftWidth, "Terraform", "#", "API"))
	b.WriteString(fmt.Sprintf("%4s %-*s   %4s %s\n", strings.Repeat("-", 4), leftWidth, strings.Repeat("-", leftWidth), strings.Repeat("-", 4), strings.Repeat("-", rightWidth)))

	// Rows
	for i := 0; i < n; i++ {
		var L, R string
		if i < len(lnsL) {
			L = lnsL[i]
		}
		if i < len(lnsR) {
			R = lnsR[i]
		}
		marker := " "
		if L != R {
			marker = "≠"
		}
		// Truncate to fit columns
		if leftWidth > 1 && len(L) > leftWidth {
			L = L[:leftWidth-1] + "…"
		}
		if rightWidth > 1 && len(R) > rightWidth {
			R = R[:rightWidth-1] + "…"
		}
		b.WriteString(fmt.Sprintf("%4d %-*s %s %4d %s\n", i+1, leftWidth, L, marker, i+1, R))
	}
	return b.String()
}
