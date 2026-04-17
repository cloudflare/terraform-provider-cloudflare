// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustDeviceDefaultProfileResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustDeviceDefaultProfileResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustDeviceDefaultProfileResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustDeviceDefaultProfileResource{}
}

// ZeroTrustDeviceDefaultProfileResource defines the resource implementation.
type ZeroTrustDeviceDefaultProfileResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustDeviceDefaultProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_device_default_profile"
}

func (r *ZeroTrustDeviceDefaultProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// syncSplitTunnelInclude writes the include list to PUT /devices/policy/include.
// The main PATCH /devices/policy endpoint silently ignores the include field;
// the WARP client only honors entries written via this dedicated endpoint.
// See: https://github.com/cloudflare/terraform-provider-cloudflare/issues/6608
func (r *ZeroTrustDeviceDefaultProfileResource) syncSplitTunnelInclude(ctx context.Context, data *ZeroTrustDeviceDefaultProfileModel) error {
	if data.Include.IsNull() || data.Include.IsUnknown() {
		return nil
	}

	elements, diags := data.Include.AsStructSliceT(ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to read include list from plan: %s", diagsString(diags))
	}

	body := make([]zero_trust.SplitTunnelIncludeUnionParam, 0, len(elements))
	for _, elem := range elements {
		entry := zero_trust.SplitTunnelIncludeParam{
			Description: cloudflare.F(elem.Description.ValueString()),
		}
		if !elem.Address.IsNull() && elem.Address.ValueString() != "" {
			entry.Address = cloudflare.F(elem.Address.ValueString())
		}
		if !elem.Host.IsNull() && elem.Host.ValueString() != "" {
			entry.Host = cloudflare.F(elem.Host.ValueString())
		}
		body = append(body, entry)
	}

	if len(body) == 0 {
		return nil
	}

	_, err := r.client.ZeroTrust.Devices.Policies.Default.Includes.Update(
		ctx,
		zero_trust.DevicePolicyDefaultIncludeUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Body:      body,
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	return err
}

// syncSplitTunnelExclude writes the exclude list to PUT /devices/policy/exclude.
// See syncSplitTunnelInclude for context.
func (r *ZeroTrustDeviceDefaultProfileResource) syncSplitTunnelExclude(ctx context.Context, data *ZeroTrustDeviceDefaultProfileModel) error {
	if data.Exclude.IsNull() || data.Exclude.IsUnknown() {
		return nil
	}

	elements, diags := data.Exclude.AsStructSliceT(ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to read exclude list from plan: %s", diagsString(diags))
	}

	body := make([]zero_trust.SplitTunnelExcludeUnionParam, 0, len(elements))
	for _, elem := range elements {
		entry := zero_trust.SplitTunnelExcludeParam{
			Description: cloudflare.F(elem.Description.ValueString()),
		}
		if !elem.Address.IsNull() && elem.Address.ValueString() != "" {
			entry.Address = cloudflare.F(elem.Address.ValueString())
		}
		if !elem.Host.IsNull() && elem.Host.ValueString() != "" {
			entry.Host = cloudflare.F(elem.Host.ValueString())
		}
		body = append(body, entry)
	}

	if len(body) == 0 {
		return nil
	}

	_, err := r.client.ZeroTrust.Devices.Policies.Default.Excludes.Update(
		ctx,
		zero_trust.DevicePolicyDefaultExcludeUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Body:      body,
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	return err
}

// readSplitTunnelInclude fetches the include list from GET /devices/policy/include.
func (r *ZeroTrustDeviceDefaultProfileResource) readSplitTunnelInclude(ctx context.Context, accountID string) (customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileIncludeModel], error) {
	null := customfield.NullObjectList[ZeroTrustDeviceDefaultProfileIncludeModel](ctx)
	iter := r.client.ZeroTrust.Devices.Policies.Default.Includes.GetAutoPaging(
		ctx,
		zero_trust.DevicePolicyDefaultIncludeGetParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	var items []ZeroTrustDeviceDefaultProfileIncludeModel
	for iter.Next() {
		item := iter.Current()
		items = append(items, ZeroTrustDeviceDefaultProfileIncludeModel{
			Address:     types.StringValue(item.Address),
			Description: types.StringValue(item.Description),
			Host:        types.StringValue(item.Host),
		})
	}
	if err := iter.Err(); err != nil {
		return null, err
	}
	list, diags := customfield.NewObjectList(ctx, items)
	if diags.HasError() {
		return null, fmt.Errorf("failed to build include list: %s", diagsString(diags))
	}
	return list, nil
}

// readSplitTunnelExclude fetches the exclude list from GET /devices/policy/exclude.
func (r *ZeroTrustDeviceDefaultProfileResource) readSplitTunnelExclude(ctx context.Context, accountID string) (customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileExcludeModel], error) {
	null := customfield.NullObjectList[ZeroTrustDeviceDefaultProfileExcludeModel](ctx)
	iter := r.client.ZeroTrust.Devices.Policies.Default.Excludes.GetAutoPaging(
		ctx,
		zero_trust.DevicePolicyDefaultExcludeGetParams{
			AccountID: cloudflare.F(accountID),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	var items []ZeroTrustDeviceDefaultProfileExcludeModel
	for iter.Next() {
		item := iter.Current()
		items = append(items, ZeroTrustDeviceDefaultProfileExcludeModel{
			Address:     types.StringValue(item.Address),
			Description: types.StringValue(item.Description),
			Host:        types.StringValue(item.Host),
		})
	}
	if err := iter.Err(); err != nil {
		return null, err
	}
	list, diags := customfield.NewObjectList(ctx, items)
	if diags.HasError() {
		return null, fmt.Errorf("failed to build exclude list: %s", diagsString(diags))
	}
	return list, nil
}

// applyPlanSplitTunnel restores plan include/exclude on data after the main
// PATCH overwrites them with empty values from the API response, then syncs
// both lists to their dedicated endpoints. Unknown plan values are normalised
// to null — unknown lists occur during a tainted replace, and unknown nested
// string fields occur when the user configures only one of address/host (the
// other two fields are computed_optional and plan as unknown). Leaving
// unknowns in state triggers "provider returned unknown value after apply".
func (r *ZeroTrustDeviceDefaultProfileResource) applyPlanSplitTunnel(ctx context.Context, data *ZeroTrustDeviceDefaultProfileModel, planInclude customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileIncludeModel], planExclude customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileExcludeModel]) error {
	if planInclude.IsUnknown() {
		planInclude = customfield.NullObjectList[ZeroTrustDeviceDefaultProfileIncludeModel](ctx)
	} else {
		resolved, diags := resolveDefaultIncludeUnknowns(ctx, planInclude)
		if diags.HasError() {
			return fmt.Errorf("resolve include unknowns: %s", diagsString(diags))
		}
		planInclude = resolved
	}
	if planExclude.IsUnknown() {
		planExclude = customfield.NullObjectList[ZeroTrustDeviceDefaultProfileExcludeModel](ctx)
	} else {
		resolved, diags := resolveDefaultExcludeUnknowns(ctx, planExclude)
		if diags.HasError() {
			return fmt.Errorf("resolve exclude unknowns: %s", diagsString(diags))
		}
		planExclude = resolved
	}
	data.Include = planInclude
	data.Exclude = planExclude

	if err := r.syncSplitTunnelInclude(ctx, data); err != nil {
		return fmt.Errorf("sync include list: %w", err)
	}
	if err := r.syncSplitTunnelExclude(ctx, data); err != nil {
		return fmt.Errorf("sync exclude list: %w", err)
	}
	return nil
}

// resolveString returns the input if known, otherwise types.StringNull().
// Used to turn computed_optional "known after apply" into concrete null values.
func resolveString(s types.String) types.String {
	if s.IsUnknown() {
		return types.StringNull()
	}
	return s
}

// resolveDefaultIncludeUnknowns walks a non-null/non-unknown NestedObjectList and
// replaces any Unknown string fields on its elements with Null.
func resolveDefaultIncludeUnknowns(ctx context.Context, in customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileIncludeModel]) (customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileIncludeModel], diag.Diagnostics) {
	if in.IsNull() || in.IsUnknown() {
		return in, nil
	}
	elements, diags := in.AsStructSliceT(ctx)
	if diags.HasError() {
		return in, diags
	}
	for i := range elements {
		elements[i].Address = resolveString(elements[i].Address)
		elements[i].Description = resolveString(elements[i].Description)
		elements[i].Host = resolveString(elements[i].Host)
	}
	return customfield.NewObjectList(ctx, elements)
}

// resolveDefaultExcludeUnknowns mirrors resolveDefaultIncludeUnknowns for the exclude list.
func resolveDefaultExcludeUnknowns(ctx context.Context, in customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileExcludeModel]) (customfield.NestedObjectList[ZeroTrustDeviceDefaultProfileExcludeModel], diag.Diagnostics) {
	if in.IsNull() || in.IsUnknown() {
		return in, nil
	}
	elements, diags := in.AsStructSliceT(ctx)
	if diags.HasError() {
		return in, diags
	}
	for i := range elements {
		elements[i].Address = resolveString(elements[i].Address)
		elements[i].Description = resolveString(elements[i].Description)
		elements[i].Host = resolveString(elements[i].Host)
	}
	return customfield.NewObjectList(ctx, elements)
}

func (r *ZeroTrustDeviceDefaultProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustDeviceDefaultProfileModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planInclude := data.Include
	planExclude := data.Exclude

	// Strip include/exclude from the main PATCH body. The endpoint silently
	// ignores them on success and can return a 500 when a previously populated
	// exclude (e.g. the account default) is swapped for an include in the same
	// request — see #6608. Both fields are written via dedicated endpoints below.
	data.Include = customfield.NullObjectList[ZeroTrustDeviceDefaultProfileIncludeModel](ctx)
	data.Exclude = customfield.NullObjectList[ZeroTrustDeviceDefaultProfileExcludeModel](ctx)

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustDeviceDefaultProfileResultEnvelope{*data}
	_, err = r.client.ZeroTrust.Devices.Policies.Default.Edit(
		ctx,
		zero_trust.DevicePolicyDefaultEditParams{
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
	data.ID = data.AccountID

	if err := r.applyPlanSplitTunnel(ctx, data, planInclude, planExclude); err != nil {
		resp.Diagnostics.AddError("failed to sync split tunnel", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceDefaultProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustDeviceDefaultProfileModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustDeviceDefaultProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planInclude := data.Include
	planExclude := data.Exclude

	// Strip include/exclude from both plan and state before diffing — the main
	// PATCH body must not contain these fields (see Create for rationale and #6608).
	nullInclude := customfield.NullObjectList[ZeroTrustDeviceDefaultProfileIncludeModel](ctx)
	nullExclude := customfield.NullObjectList[ZeroTrustDeviceDefaultProfileExcludeModel](ctx)
	data.Include = nullInclude
	data.Exclude = nullExclude
	state.Include = nullInclude
	state.Exclude = nullExclude

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}

	// MarshalJSONForUpdate (MarshalForPatch) returns nil when every serialisable
	// field is either unchanged or computed_optional unknown. Sending a nil body
	// produces an empty PATCH request which the API rejects with 400. In that
	// case skip the PATCH and fall through to a GET to refresh computed fields.
	if dataBytes != nil {
		res := new(http.Response)
		env := ZeroTrustDeviceDefaultProfileResultEnvelope{*data}
		_, err = r.client.ZeroTrust.Devices.Policies.Default.Edit(
			ctx,
			zero_trust.DevicePolicyDefaultEditParams{
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
		data.ID = data.AccountID
	} else {
		// No-op update: MarshalForPatch produced nil (no serialisable changes).
		// Do a GET to populate computed fields into state.
		res := new(http.Response)
		env := ZeroTrustDeviceDefaultProfileResultEnvelope{*data}
		_, err = r.client.ZeroTrust.Devices.Policies.Default.Get(
			ctx,
			zero_trust.DevicePolicyDefaultGetParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
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
		data.ID = data.AccountID
	}

	if err := r.applyPlanSplitTunnel(ctx, data, planInclude, planExclude); err != nil {
		resp.Diagnostics.AddError("failed to sync split tunnel", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceDefaultProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustDeviceDefaultProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Capture include/exclude from prior state. The main GET /devices/policy
	// does not return these fields, so we cannot let the API response overwrite
	// them. The WARP API preserves the order we wrote, but a re-read in plan
	// would compare against config order (HCL map iteration sorts by key) and
	// produce perpetual diffs. Preserving prior state values keeps Terraform
	// in sync with what it last wrote.
	priorInclude := data.Include
	priorExclude := data.Exclude

	res := new(http.Response)
	env := ZeroTrustDeviceDefaultProfileResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Devices.Policies.Default.Get(
		ctx,
		zero_trust.DevicePolicyDefaultGetParams{
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
	data.ID = data.AccountID
	data.Include = priorInclude
	data.Exclude = priorExclude

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceDefaultProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *ZeroTrustDeviceDefaultProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ZeroTrustDeviceDefaultProfileModel)

	path := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>",
		&path,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path)

	res := new(http.Response)
	env := ZeroTrustDeviceDefaultProfileResultEnvelope{*data}
	_, err := r.client.ZeroTrust.Devices.Policies.Default.Get(
		ctx,
		zero_trust.DevicePolicyDefaultGetParams{
			AccountID: cloudflare.F(path),
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
	data.ID = data.AccountID

	// On import, the include/exclude lists must be fetched from their
	// dedicated endpoints since GET /devices/policy doesn't return them.
	includeList, err := r.readSplitTunnelInclude(ctx, path)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch split tunnel include list", err.Error())
		return
	}
	data.Include = includeList

	excludeList, err := r.readSplitTunnelExclude(ctx, path)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch split tunnel exclude list", err.Error())
		return
	}
	data.Exclude = excludeList

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDeviceDefaultProfileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() {
		resp.Diagnostics.AddWarning(
			"Resource Destruction Considerations",
			"This resource cannot be destroyed from Terraform. If you create this resource, it will be "+
				"present in the API until manually deleted.",
		)
	}
	if req.Plan.Raw.IsNull() {
		resp.Diagnostics.AddWarning(
			"Resource Destruction Considerations",
			"Applying this resource destruction will remove the resource from the Terraform state "+
				"but will not change it in the API. If you would like to destroy or reset this resource "+
				"in the API, refer to the documentation for how to do it manually.",
		)
	}

	// include and exclude are mutually exclusive at the API level (and validated
	// via ConflictsWith on config). The schema's listplanmodifier.UseStateForUnknown()
	// runs after validators and copies the stale prior-state value onto the side
	// the user just removed from config — so when a user swaps include<->exclude
	// the plan ends up with BOTH populated, and apply tries to sync both which
	// race and clobber each other on the API.
	//
	// When one side is explicitly in config, force the other to null so apply
	// clears it from the API. When neither side is in config and the plan is
	// Unknown (fresh create, or null prior state), copy state to avoid a
	// permanent "(known after apply)" diff for unconfigured lists.
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var state, plan, config ZeroTrustDeviceDefaultProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		return
	}
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		return
	}
	if diags := req.Config.Get(ctx, &config); diags.HasError() {
		return
	}
	includeInConfig := !config.Include.IsNull() && !config.Include.IsUnknown()
	excludeInConfig := !config.Exclude.IsNull() && !config.Exclude.IsUnknown()
	changed := false
	if includeInConfig && !plan.Exclude.IsNull() {
		plan.Exclude = customfield.NullObjectList[ZeroTrustDeviceDefaultProfileExcludeModel](ctx)
		changed = true
	}
	if excludeInConfig && !plan.Include.IsNull() {
		plan.Include = customfield.NullObjectList[ZeroTrustDeviceDefaultProfileIncludeModel](ctx)
		changed = true
	}
	if !includeInConfig && plan.Include.IsUnknown() {
		plan.Include = state.Include
		changed = true
	}
	if !excludeInConfig && plan.Exclude.IsUnknown() {
		plan.Exclude = state.Exclude
		changed = true
	}
	if changed {
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

// diagsString joins error diagnostics into a single string for error wrapping.
func diagsString(diags diag.Diagnostics) string {
	errs := diags.Errors()
	parts := make([]string, 0, len(errs))
	for _, d := range errs {
		parts = append(parts, fmt.Sprintf("%s: %s", d.Summary(), d.Detail()))
	}
	if len(parts) == 0 {
		return "(no error details)"
	}
	out := parts[0]
	for _, p := range parts[1:] {
		out += "; " + p
	}
	return out
}
