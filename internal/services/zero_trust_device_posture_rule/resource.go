// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_rule

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustDevicePostureRuleResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustDevicePostureRuleResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustDevicePostureRuleResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustDevicePostureRuleResource{}
}

// ZeroTrustDevicePostureRuleResource defines the resource implementation.
type ZeroTrustDevicePostureRuleResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustDevicePostureRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_device_posture_rule"
}

func (r *ZeroTrustDevicePostureRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustDevicePostureRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustDevicePostureRuleModel

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
	env := ZeroTrustDevicePostureRuleResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Devices.Posture.New(
		ctx,
		zero_trust.DevicePostureNewParams{
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

func (r *ZeroTrustDevicePostureRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustDevicePostureRuleModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustDevicePostureRuleModel

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
	env := ZeroTrustDevicePostureRuleResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Devices.Posture.Update(
		ctx,
		data.ID.ValueString(),
		zero_trust.DevicePostureUpdateParams{
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

func (r *ZeroTrustDevicePostureRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustDevicePostureRuleModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ZeroTrustDevicePostureRuleResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Devices.Posture.Get(
		ctx,
		data.ID.ValueString(),
		zero_trust.DevicePostureGetParams{
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

	// Normalize API response to match configuration expectations  
	r.normalizeReadData(ctx, data, req.State)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDevicePostureRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustDevicePostureRuleModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.Devices.Posture.Delete(
		ctx,
		data.ID.ValueString(),
		zero_trust.DevicePostureDeleteParams{
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

func (r *ZeroTrustDevicePostureRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZeroTrustDevicePostureRuleModel = new(ZeroTrustDevicePostureRuleModel)

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
	env := ZeroTrustDevicePostureRuleResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Devices.Posture.Get(
		ctx,
		path_rule_id,
		zero_trust.DevicePostureGetParams{
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

	// For import, we want to clean up any API-added defaults to match typical configs
	r.normalizeImportData(data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDevicePostureRuleResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// ModifyPlan is disabled since we can't modify non-computed attributes
	// We'll handle the differences in the Read method instead
}

func (r *ZeroTrustDevicePostureRuleResource) normalizeReadData(ctx context.Context, data *ZeroTrustDevicePostureRuleModel, state tfsdk.State) {
	// Get the current state to compare with API response
	var currentState *ZeroTrustDevicePostureRuleModel
	state.Get(ctx, &currentState)

	if currentState == nil {
		return
	}

	// Handle schedule field: if it was null in current state and API added "5m" default, keep it null
	if currentState.Schedule.IsNull() && !data.Schedule.IsNull() && data.Schedule.ValueString() == "5m" {
		// For certain rule types, API sets default schedule="5m" when none was configured
		// Keep it null to match the original configuration intent
		data.Schedule = types.StringNull()
	}

	// Handle input fields that API might not return
	if data.Input != nil && currentState.Input != nil {
		// operating_system field: if it was null in current state and API added it, keep it null
		if currentState.Input.OperatingSystem.IsNull() && !data.Input.OperatingSystem.IsNull() {
			// API sometimes adds operating_system automatically, keep it null to match config
			data.Input.OperatingSystem = types.StringNull()
		}

		// check_disks field: if it was null in current state but API returns empty array, keep it null
		if currentState.Input.CheckDisks == nil && data.Input.CheckDisks != nil && len(*data.Input.CheckDisks) == 0 {
			// API returns empty array when config had none, keep it null to match config
			data.Input.CheckDisks = nil
		}

		// SentinelOne fields: API may not return these fields even if config specifies them
		if !currentState.Input.ActiveThreats.IsNull() && data.Input.ActiveThreats.IsNull() {
			data.Input.ActiveThreats = currentState.Input.ActiveThreats
		}
		if !currentState.Input.Operator.IsNull() && data.Input.Operator.IsNull() {
			data.Input.Operator = currentState.Input.Operator
		}
		if !currentState.Input.Infected.IsNull() && data.Input.Infected.IsNull() {
			data.Input.Infected = currentState.Input.Infected
		}
		if !currentState.Input.IsActive.IsNull() && data.Input.IsActive.IsNull() {
			data.Input.IsActive = currentState.Input.IsActive
		}
		if !currentState.Input.NetworkStatus.IsNull() && data.Input.NetworkStatus.IsNull() {
			data.Input.NetworkStatus = currentState.Input.NetworkStatus
		}
		if !currentState.Input.OperationalState.IsNull() && data.Input.OperationalState.IsNull() {
			data.Input.OperationalState = currentState.Input.OperationalState
		}
	}
}

func (r *ZeroTrustDevicePostureRuleResource) normalizeImportData(data *ZeroTrustDevicePostureRuleModel) {
	// For imports, remove commonly added defaults that users typically don't configure
	// Only remove schedule if it's "5m" for rule types that commonly don't need it (like serial_number, application)
	// but keep it for rule types that typically specify it (like file rules)
	
	if !data.Schedule.IsNull() && data.Schedule.ValueString() == "5m" {
		ruleType := data.Type.ValueString()
		// Only remove default schedule for rule types that typically don't specify it
		if ruleType == "serial_number" || ruleType == "application" {
			data.Schedule = types.StringNull()
		}
	}

	// Remove operating_system from input if it matches the platform in match
	if data.Input != nil && data.Match != nil && len(*data.Match) > 0 {
		match := (*data.Match)[0]
		if !data.Input.OperatingSystem.IsNull() && !match.Platform.IsNull() {
			// Map platform to operating system equivalents
			platformToOS := map[string]string{
				"mac":     "mac",
				"windows": "windows", 
				"linux":   "linux",
			}
			
			platform := match.Platform.ValueString()
			os := data.Input.OperatingSystem.ValueString()
			
			if expectedOS, exists := platformToOS[platform]; exists && os == expectedOS {
				// If operating_system matches the expected value based on platform, remove it
				data.Input.OperatingSystem = types.StringNull()
			}
		}
	}
}
