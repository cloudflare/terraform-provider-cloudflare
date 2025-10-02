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

// detectDriftOnUpdate compares Terraform planned config vs API without filtering computed fields.
func (r *ZeroTrustGatewayPolicyResource) detectDriftOnUpdate(ctx context.Context, apiState, plannedConfig *ZeroTrustGatewayPolicyModel) diag.Diagnostics {
	header := "detectDriftOnUpdate@UPDATE"
	return r.detectDrift(ctx, plannedConfig, apiState, false, header, "Terraform", "API")
}

// detectConfigurationDriftOnRead compares Terraform state vs API and filters computed-only fields.
func (r *ZeroTrustGatewayPolicyResource) detectConfigurationDriftOnRead(ctx context.Context, apiState, terraformState *ZeroTrustGatewayPolicyModel) diag.Diagnostics {
	header := "detectConfigurationDriftOnRead@READ"
	return r.detectDrift(ctx, terraformState, apiState, true, header, "Terraform", "API")
}

// detectDrift normalizes two models to JSON, optionally filters computed-only keys,
// pretty-prints them, and emits a side-by-side diff if they differ.
func (r *ZeroTrustGatewayPolicyResource) detectDrift(
	ctx context.Context,
	left, right *ZeroTrustGatewayPolicyModel,
	filterComputed bool,
	header string,
	leftLabel string,
	rightLabel string,
) diag.Diagnostics {
	var diags diag.Diagnostics

	if left == nil || right == nil {
		return diags
	}

	// Serialize both objects using model-aware JSON marshaling
	leftBytes, errL := left.MarshalJSON()
	rightBytes, errR := right.MarshalJSON()
	if errL != nil || errR != nil {
		diags.AddWarning(
			"Configuration Drift Detected",
			fmt.Sprintf("error marshalling for drift detection (left=%v, right=%v)", errL, errR),
		)
		return diags
	}

	// Normalize into maps for optional filtering
	var leftMap map[string]any
	var rightMap map[string]any
	if err := json.Unmarshal(leftBytes, &leftMap); err != nil {
		diags.AddWarning("Configuration Drift Detected", "failed to parse left object for drift detection")
		return diags
	}
	if err := json.Unmarshal(rightBytes, &rightMap); err != nil {
		diags.AddWarning("Configuration Drift Detected", "failed to parse right object for drift detection")
		return diags
	}

	if filterComputed {
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
		remove(leftMap)
		remove(rightMap)
	}

	normLeft, _ := json.MarshalIndent(leftMap, "", "  ")
	normRight, _ := json.MarshalIndent(rightMap, "", "  ")

	if string(normLeft) == string(normRight) {
		return diags
	}

	var msg strings.Builder
	if header != "" {
		msg.WriteString(header)
		msg.WriteString("\nConfiguration drift detected between API state and Terraform configuration.")
		msg.WriteString("\n\n")
	}

	msg.WriteString(fmt.Sprintf("Side-by-side (%s | %s):\n", leftLabel, rightLabel))
	msg.WriteString(sideBySideDiff(string(normLeft), string(normRight), 60))
	msg.WriteString("\n\nTo fix the drift, update your terraform declaration to match the current API state.")

	diags.AddWarning("Configuration Drift Detected", msg.String())
	return diags
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
			// Red color for difference marker
			marker = "\x1b[31m≠\x1b[0m"
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
