// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ManagedTransformsResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ManagedTransformsResource)(nil)
var _ resource.ResourceWithImportState = (*ManagedTransformsResource)(nil)

func NewResource() resource.Resource {
	return &ManagedTransformsResource{}
}

// ManagedTransformsResource defines the resource implementation.
type ManagedTransformsResource struct {
	client *cloudflare.Client
}

func (r *ManagedTransformsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_managed_transforms"
}

func (r *ManagedTransformsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ManagedTransformsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ManagedTransformsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Disable any transforms that are currently enabled in the zone but not in the plan.
	// This ensures Create correctly handles zones with pre-existing enabled transforms,
	// without deleting underlying rulesets (which would break subsequent PATCH calls).
	err := r.disableMissingTransformations(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to disable existing transforms before creation", err.Error())
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{*data}
	_, err = r.client.ManagedTransforms.Edit(
		ctx,
		managed_transforms.ManagedTransformEditParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	data.ID = data.ZoneID

	var plan *ManagedTransformsModel
	req.Plan.Get(ctx, &plan)

	normalizeResponse(data, plan, false)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ManagedTransformsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ManagedTransformsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ManagedTransformsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Because we are patching the resource we need to explicitly add `enable = false` in order to remove
	// transformations. Otherwise, we will simply leave them with their int previous state.
	err := r.disableMissingTransformations(ctx, data)

	if err != nil {
		resp.Diagnostics.AddError("failed to disable missing transformations", err.Error())
		return
	}

	// Always send the full body (not just changed fields) because the managed_headers API
	// treats PATCH as a replace per category: omitting managed_response_headers deletes the
	// response headers zone ruleset, causing subsequent enables to fail with 404.
	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{*data}
	_, err = r.client.ManagedTransforms.Edit(
		ctx,
		managed_transforms.ManagedTransformEditParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	data.ID = data.ZoneID

	var plan *ManagedTransformsModel
	req.Plan.Get(ctx, &plan)

	normalizeResponse(data, plan, false)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ManagedTransformsResource) disableMissingTransformations(
	ctx context.Context,
	plan *ManagedTransformsModel,
) error {
	res, err := r.client.ManagedTransforms.List(
		ctx,
		managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.F(plan.ZoneID.ValueString()),
		},
	)

	if err != nil {
		// If the API returns a 403 due to conflicting transforms being enabled,
		// attempt to clear all transforms by using a fallback disable-all approach.
		// This handles cases where a zone is in an inconsistent state from
		// previous operations or manual changes.
		errStr := err.Error()
		if strings.Contains(errStr, "403") && strings.Contains(errStr, "conflict") {
			return r.clearConflictingTransforms(ctx, plan)
		}
		return err
	}

	if res == nil {
		return nil
	}

	if plan.ManagedRequestHeaders != nil {
		var existingTransformations = []*ManagedTransformsManagedRequestHeadersModel{}

		for _, t := range res.ManagedRequestHeaders {
			existingTransformations = append(existingTransformations, &ManagedTransformsManagedRequestHeadersModel{
				ID:      types.StringValue(t.ID),
				Enabled: types.BoolValue(t.Enabled),
			})
		}

		newTransformations := disableMissingTransformations(
			*plan.ManagedRequestHeaders,
			existingTransformations,
		)
		plan.ManagedRequestHeaders = &newTransformations
	}

	if plan.ManagedResponseHeaders != nil {
		var existingTransformations = []*ManagedTransformsManagedResponseHeadersModel{}

		for _, t := range res.ManagedResponseHeaders {
			existingTransformations = append(existingTransformations, &ManagedTransformsManagedResponseHeadersModel{
				ID:      types.StringValue(t.ID),
				Enabled: types.BoolValue(t.Enabled),
			})
		}

		newTransformations := disableMissingTransformations(
			*plan.ManagedResponseHeaders,
			existingTransformations,
		)
		plan.ManagedResponseHeaders = &newTransformations
	}

	return nil
}

// clearConflictingTransforms attempts to clear all conflicting transforms by
// disabling known conflicting headers. This is used as a fallback when the API
// returns a 403 conflict error, preventing us from listing current transforms.
func (r *ManagedTransformsResource) clearConflictingTransforms(
	ctx context.Context,
	plan *ManagedTransformsModel,
) error {
	zoneID := plan.ZoneID.ValueString()

	// Build a list of all known request headers that might be in conflict.
	// This includes both the ones in the plan (to ensure they're set correctly)
	// and known conflicting headers that might be enabled.
	knownRequestHeaders := []string{
		"add_true_client_ip_headers",
		"remove_visitor_ip_headers",
		"add_visitor_location_headers",
		"remove_x_powered_by_headers",
	}

	// Start with headers from the plan
	reqHeadersMap := make(map[string]bool)
	if plan.ManagedRequestHeaders != nil {
		for _, h := range *plan.ManagedRequestHeaders {
			reqHeadersMap[h.ID.ValueString()] = h.Enabled.ValueBool()
		}
	}

	// Add all known headers as disabled (unless they're in the plan and enabled)
	var reqHeaders []*ManagedTransformsManagedRequestHeadersModel
	for _, id := range knownRequestHeaders {
		enabled, inPlan := reqHeadersMap[id]
		if inPlan {
			reqHeaders = append(reqHeaders, &ManagedTransformsManagedRequestHeadersModel{
				ID:      types.StringValue(id),
				Enabled: types.BoolValue(enabled),
			})
		} else {
			reqHeaders = append(reqHeaders, &ManagedTransformsManagedRequestHeadersModel{
				ID:      types.StringValue(id),
				Enabled: types.BoolValue(false),
			})
		}
	}

	// Also include any other headers from the plan that aren't in our known list
	if plan.ManagedRequestHeaders != nil {
		for _, h := range *plan.ManagedRequestHeaders {
			if _, known := reqHeadersMap[h.ID.ValueString()]; !known {
				reqHeaders = append(reqHeaders, &ManagedTransformsManagedRequestHeadersModel{
					ID:      h.ID,
					Enabled: h.Enabled,
				})
			}
		}
	}

	// Handle response headers similarly
	var respHeaders []*ManagedTransformsManagedResponseHeadersModel
	if plan.ManagedResponseHeaders != nil {
		for _, h := range *plan.ManagedResponseHeaders {
			respHeaders = append(respHeaders, &ManagedTransformsManagedResponseHeadersModel{
				ID:      h.ID,
				Enabled: h.Enabled,
			})
		}
	}

	data := &ManagedTransformsModel{
		ZoneID:                 types.StringValue(zoneID),
		ManagedRequestHeaders:  &reqHeaders,
		ManagedResponseHeaders: &respHeaders,
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = r.client.ManagedTransforms.Edit(
		ctx,
		managed_transforms.ManagedTransformEditParams{
			ZoneID: cloudflare.F(zoneID),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	return err
}

func disableMissingTransformations[T transformation](
	transformations []T,
	existingTransformations []T,
) []T {
	inTransformations := make(map[string]bool)

	for _, transformation := range transformations {
		inTransformations[transformation.id()] = true
	}

	newTransformations := transformations

	for _, transformation := range existingTransformations {
		if transformation.enabled() && !inTransformations[transformation.id()] {
			transformation.disable()
			newTransformations = append(newTransformations, transformation)
		}
	}

	return newTransformations
}

func (r *ManagedTransformsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ManagedTransformsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{*data}
	_, err := r.client.ManagedTransforms.List(
		ctx,
		managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	data.ID = data.ZoneID

	var state *ManagedTransformsModel
	req.State.Get(ctx, &state)

	// stateOnly=true: during Read, only include transforms that are explicitly tracked in
	// state, to prevent unrelated enabled transforms from polluting state.
	normalizeResponse(data, state, true)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type transformation interface {
	id() string
	enabled() bool

	disable()
}

func (m ManagedTransformsManagedRequestHeadersModel) id() string {
	return m.ID.ValueString()
}

func (m ManagedTransformsManagedRequestHeadersModel) enabled() bool {
	return m.Enabled.ValueBool()
}

func (m *ManagedTransformsManagedRequestHeadersModel) disable() {
	m.Enabled = types.BoolValue(false)
}

func (m ManagedTransformsManagedResponseHeadersModel) id() string {
	return m.ID.ValueString()
}

func (m ManagedTransformsManagedResponseHeadersModel) enabled() bool {
	return m.Enabled.ValueBool()
}

func (m *ManagedTransformsManagedResponseHeadersModel) disable() {
	m.Enabled = types.BoolValue(false)
}

// normalizeResponse filters the API response to only include transforms relevant to the
// current state/plan. When stateOnly=true (used in Read), only transforms explicitly in
// state are kept, preventing unrelated enabled transforms from polluting state.
// When stateOnly=false (used in Create/Update/Import), enabled transforms not in state
// are also included so they can be tracked.
func normalizeResponse(response *ManagedTransformsModel, state *ManagedTransformsModel, stateOnly bool) {
	if response.ManagedRequestHeaders != nil {
		stateManagedRequestHeaders := []*ManagedTransformsManagedRequestHeadersModel{}
		if state != nil && state.ManagedRequestHeaders != nil {
			stateManagedRequestHeaders = *state.ManagedRequestHeaders
		}

		t := transformationsView(*response.ManagedRequestHeaders, stateManagedRequestHeaders, stateOnly)
		response.ManagedRequestHeaders = &t
	}
	if response.ManagedResponseHeaders != nil {
		stateManagedResponseHeaders := []*ManagedTransformsManagedResponseHeadersModel{}
		if state != nil && state.ManagedResponseHeaders != nil {
			stateManagedResponseHeaders = *state.ManagedResponseHeaders
		}

		t := transformationsView(*response.ManagedResponseHeaders, stateManagedResponseHeaders, stateOnly)
		response.ManagedResponseHeaders = &t
	}
}

func transformationsView[T transformation](transformations []T, stateTransformations []T, stateOnly bool) []T {
	inState := make(map[string]bool)
	stateByID := make(map[string]T)

	for _, transformation := range stateTransformations {
		inState[transformation.id()] = true
		stateByID[transformation.id()] = transformation
	}

	newTransformations := []T{}

	for _, transformation := range transformations {
		isDefaultTransformation := !transformation.enabled()

		if stateOnly {
			// During Read: only include transforms explicitly tracked in state.
			// This prevents unrelated enabled transforms (left by other tests/operations)
			// from being added to state and causing spurious plan diffs.
			// Use the state value rather than the API value to handle API eventual
			// consistency issues where the API transiently returns stale data after a write.
			if inState[transformation.id()] {
				newTransformations = append(newTransformations, stateByID[transformation.id()])
			}
		} else {
			// During Create/Update/Import: include transforms in state OR any that are
			// non-default (enabled=true), so newly enabled transforms are tracked.
			if inState[transformation.id()] || !isDefaultTransformation {
				newTransformations = append(newTransformations, transformation)
			}
		}
	}

	return newTransformations
}

// disableAllTransforms disables all currently enabled managed transforms by sending a PATCH
// with enabled=false for all transforms. This preserves the underlying zone rulesets,
// unlike the DELETE endpoint which removes them and prevents subsequent re-enablement.
func (r *ManagedTransformsResource) disableAllTransforms(ctx context.Context, zoneID string) error {
	res, err := r.client.ManagedTransforms.List(
		ctx,
		managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.F(zoneID),
		},
	)
	if err != nil {
		return err
	}
	if res == nil {
		return nil
	}

	// Build a model with all transforms disabled.
	var reqHeaders []*ManagedTransformsManagedRequestHeadersModel
	for _, t := range res.ManagedRequestHeaders {
		if t.Enabled {
			reqHeaders = append(reqHeaders, &ManagedTransformsManagedRequestHeadersModel{
				ID:      types.StringValue(t.ID),
				Enabled: types.BoolValue(false),
			})
		}
	}
	var respHeaders []*ManagedTransformsManagedResponseHeadersModel
	for _, t := range res.ManagedResponseHeaders {
		if t.Enabled {
			respHeaders = append(respHeaders, &ManagedTransformsManagedResponseHeadersModel{
				ID:      types.StringValue(t.ID),
				Enabled: types.BoolValue(false),
			})
		}
	}

	// If nothing is enabled, nothing to do.
	if len(reqHeaders) == 0 && len(respHeaders) == 0 {
		return nil
	}

	data := &ManagedTransformsModel{
		ZoneID:                 types.StringValue(zoneID),
		ManagedRequestHeaders:  &reqHeaders,
		ManagedResponseHeaders: &respHeaders,
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = r.client.ManagedTransforms.Edit(
		ctx,
		managed_transforms.ManagedTransformEditParams{
			ZoneID: cloudflare.F(zoneID),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	return err
}

func (r *ManagedTransformsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ManagedTransformsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use PATCH to disable all transforms instead of DELETE, which would remove the
	// underlying zone rulesets and prevent subsequent re-enablement via PATCH.
	// If the PATCH fails because the backing rulesets no longer exist (404), the
	// transforms are effectively disabled, so we treat this as a success.
	err := r.disableAllTransforms(ctx, data.ZoneID.ValueString())
	if err != nil {
		errStr := err.Error()
		// If the rulesets don't exist (404/not found), transforms are effectively disabled.
		if strings.Contains(errStr, "404") || strings.Contains(errStr, "not found") || strings.Contains(errStr, "could not find") {
			return
		}
		resp.Diagnostics.AddError("failed to disable managed transforms", err.Error())
		return
	}
}

func (r *ManagedTransformsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ManagedTransformsModel)

	path := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>",
		&path,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path)

	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{*data}
	_, err := r.client.ManagedTransforms.List(
		ctx,
		managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.F(path),
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
	data.ID = data.ZoneID

	// stateOnly=false: during ImportState, include all enabled transforms since we have
	// no prior state to compare against.
	normalizeResponse(data, nil, false)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ManagedTransformsResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
