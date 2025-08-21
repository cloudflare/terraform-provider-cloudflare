// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustDeviceCustomProfileResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustDeviceCustomProfileResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustDeviceCustomProfileResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustDeviceCustomProfileResource{}
}

// ZeroTrustDeviceCustomProfileResource defines the resource implementation.
type ZeroTrustDeviceCustomProfileResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustDeviceCustomProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_device_custom_profile"
}

func (r *ZeroTrustDeviceCustomProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustDeviceCustomProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustDeviceCustomProfileModel

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
	env := ZeroTrustDeviceCustomProfileResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Devices.Policies.Custom.New(
		ctx,
		zero_trust.DevicePolicyCustomNewParams{
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
	data.ID = data.PolicyID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceCustomProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustDeviceCustomProfileModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustDeviceCustomProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}

	// If there are no changes (empty JSON), skip the API call and use plan with computed fields from state
	if len(dataBytes) == 0 || string(dataBytes) == "" {
		// No API changes needed - but preserve computed fields from state
		// Keep the plan as-is for user-specified fields (including nulls)
		// Only update computed fields from state to prevent drift
		if !state.Default.IsNull() {
			data.Default = state.Default
		}
		if !state.GatewayUniqueID.IsNull() {
			data.GatewayUniqueID = state.GatewayUniqueID
		}
		if !state.FallbackDomains.IsNull() {
			data.FallbackDomains = state.FallbackDomains
		}
		if !state.TargetTests.IsNull() {
			data.TargetTests = state.TargetTests
		}
		data.ID = data.PolicyID
	} else {
		// There are changes - proceed with the API call
		res := new(http.Response)
		env := ZeroTrustDeviceCustomProfileResultEnvelope{*data}
		_, err = r.client.ZeroTrust.Devices.Policies.Custom.Edit(
			ctx,
			data.PolicyID.ValueString(),
			zero_trust.DevicePolicyCustomEditParams{
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
		data.ID = data.PolicyID

		// After API call, adjust the result to match user's plan for optional fields
		// If user wanted to remove description (set to null), ensure the final state reflects that
		// even if API returned empty string
		var originalPlan *ZeroTrustDeviceCustomProfileModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &originalPlan)...)
		if !resp.Diagnostics.HasError() {
			if originalPlan.Description.IsNull() {
				data.Description = types.StringNull()
			}
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceCustomProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustDeviceCustomProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustDeviceCustomProfileResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Devices.Policies.Custom.Get(
		ctx,
		data.PolicyID.ValueString(),
		zero_trust.DevicePolicyCustomGetParams{
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
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.PolicyID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceCustomProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustDeviceCustomProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Devices.Policies.Custom.Delete(
		ctx,
		data.PolicyID.ValueString(),
		zero_trust.DevicePolicyCustomDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.PolicyID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceCustomProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustDeviceCustomProfileModel = new(ZeroTrustDeviceCustomProfileModel)

	path_account_id := ""
	path_policy_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<policy_id>",
		&path_account_id,
		&path_policy_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.PolicyID = types.StringValue(path_policy_id)

	res := new(http.Response)
	env := ZeroTrustDeviceCustomProfileResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Devices.Policies.Custom.Get(
		ctx,
		path_policy_id,
		zero_trust.DevicePolicyCustomGetParams{
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
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.PolicyID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceCustomProfileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If this is a destroy operation, we don't need to modify the plan
	if req.Plan.Raw.IsNull() {
		return
	}

	// Get the planned and current state
	var plan, state ZeroTrustDeviceCustomProfileModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If there's no state (i.e., create operation), we don't need to modify the plan
	if req.State.Raw.IsNull() {
		return
	}

	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// ONLY modify computed and computed_optional fields in ModifyPlan
	// Do NOT modify optional fields - they should be handled in MarshalJSONForUpdate
	planModified := false

	// Computed fields should always be preserved from state to prevent drift
	if !state.Default.IsNull() {
		plan.Default = state.Default
		planModified = true
	}
	if !state.GatewayUniqueID.IsNull() {
		plan.GatewayUniqueID = state.GatewayUniqueID
		planModified = true
	}
	if !state.FallbackDomains.IsNull() {
		plan.FallbackDomains = state.FallbackDomains
		planModified = true
	}
	if !state.TargetTests.IsNull() {
		plan.TargetTests = state.TargetTests
		planModified = true
	}

	// For computed_optional fields, preserve state values when plan is null/unknown
	if (plan.AllowModeSwitch.IsNull() || plan.AllowModeSwitch.IsUnknown()) && !state.AllowModeSwitch.IsNull() {
		plan.AllowModeSwitch = state.AllowModeSwitch
		planModified = true
	}
	if (plan.AllowUpdates.IsNull() || plan.AllowUpdates.IsUnknown()) && !state.AllowUpdates.IsNull() {
		plan.AllowUpdates = state.AllowUpdates
		planModified = true
	}
	if (plan.AllowedToLeave.IsNull() || plan.AllowedToLeave.IsUnknown()) && !state.AllowedToLeave.IsNull() {
		plan.AllowedToLeave = state.AllowedToLeave
		planModified = true
	}
	if (plan.AutoConnect.IsNull() || plan.AutoConnect.IsUnknown()) && !state.AutoConnect.IsNull() {
		plan.AutoConnect = state.AutoConnect
		planModified = true
	}
	if (plan.CaptivePortal.IsNull() || plan.CaptivePortal.IsUnknown()) && !state.CaptivePortal.IsNull() {
		plan.CaptivePortal = state.CaptivePortal
		planModified = true
	}
	if (plan.DisableAutoFallback.IsNull() || plan.DisableAutoFallback.IsUnknown()) && !state.DisableAutoFallback.IsNull() {
		plan.DisableAutoFallback = state.DisableAutoFallback
		planModified = true
	}
	if (plan.ExcludeOfficeIPs.IsNull() || plan.ExcludeOfficeIPs.IsUnknown()) && !state.ExcludeOfficeIPs.IsNull() {
		plan.ExcludeOfficeIPs = state.ExcludeOfficeIPs
		planModified = true
	}
	if (plan.RegisterInterfaceIPWithDNS.IsNull() || plan.RegisterInterfaceIPWithDNS.IsUnknown()) && !state.RegisterInterfaceIPWithDNS.IsNull() {
		plan.RegisterInterfaceIPWithDNS = state.RegisterInterfaceIPWithDNS
		planModified = true
	}
	if (plan.SccmVpnBoundarySupport.IsNull() || plan.SccmVpnBoundarySupport.IsUnknown()) && !state.SccmVpnBoundarySupport.IsNull() {
		plan.SccmVpnBoundarySupport = state.SccmVpnBoundarySupport
		planModified = true
	}
	if (plan.SupportURL.IsNull() || plan.SupportURL.IsUnknown()) && !state.SupportURL.IsNull() {
		plan.SupportURL = state.SupportURL
		planModified = true
	}
	if (plan.SwitchLocked.IsNull() || plan.SwitchLocked.IsUnknown()) && !state.SwitchLocked.IsNull() {
		plan.SwitchLocked = state.SwitchLocked
		planModified = true
	}
	if (plan.TunnelProtocol.IsNull() || plan.TunnelProtocol.IsUnknown()) && !state.TunnelProtocol.IsNull() {
		plan.TunnelProtocol = state.TunnelProtocol
		planModified = true
	}

	// If we modified the plan, update it
	if planModified {
		diags = resp.Plan.Set(ctx, plan)
		resp.Diagnostics.Append(diags...)
	}
}
