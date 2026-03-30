// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
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

	// Normalize duration strings to prevent drift from API returning verbose formats
	normalizeRuleSettingsDurations(ctx, data)

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
	var data = new(ZeroTrustGatewayPolicyModel)

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

// normalizeDuration converts duration strings to their verbose canonical form
// to match the API response format. For example: "24h" -> "24h0m0s"
// This ensures state matches the API response format.
func normalizeDuration(s string) string {
	if s == "" {
		return s
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return s // Return original if parsing fails
	}
	return d.String()
}

// normalizeRuleSettingsDurations normalizes duration fields in rule_settings
// to prevent drift from API returning verbose formats like "24h0m0s" vs config "24h"
func normalizeRuleSettingsDurations(ctx context.Context, data *ZeroTrustGatewayPolicyModel) {
	if data.RuleSettings.IsNull() || data.RuleSettings.IsUnknown() {
		return
	}

	rs, diags := data.RuleSettings.Value(ctx)
	if diags.HasError() || rs == nil {
		return
	}

	// Normalize check_session.duration
	if rs.CheckSession != nil && !rs.CheckSession.Duration.IsNull() && !rs.CheckSession.Duration.IsUnknown() {
		normalized := normalizeDuration(rs.CheckSession.Duration.ValueString())
		rs.CheckSession.Duration = types.StringValue(normalized)
	}

	// Update the rule settings back - create a new NestedObject with the modified value
	newRuleSettings, _ := customfield.NewObject(ctx, rs)
	data.RuleSettings = newRuleSettings
}
