// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

func (r *ManagedTransformsResource) checkAllDisabledBeforeCreation(ctx context.Context, zoneId string) diag.Diagnostics {
	var diagnostics diag.Diagnostics

	res, err := r.client.ManagedTransforms.List(
		ctx,
		managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.F(zoneId),
		},
	)

	if err != nil {
		diagnostics.AddError("failed to get managed transforms", err.Error())
		return diagnostics
	}

	if res == nil {
		return diagnostics
	}

	for _, t := range res.ManagedRequestHeaders {
		if t.Enabled {
			diagnostics.AddError("cannot create resource", fmt.Sprintf("managed request header transform %s cannot be enabled before creation", t.ID))
		}
	}

	for _, t := range res.ManagedResponseHeaders {
		if t.Enabled {
			diagnostics.AddError("cannot create resource", fmt.Sprintf("managed response header transform %s cannot be enabled before creation", t.ID))
		}
	}

	return diagnostics
}

func (r *ManagedTransformsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ManagedTransformsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

<<<<<<< Updated upstream
	// We need to check that all transformations are disabled, as we don't want them to be silently overwritten.
	// This is also needed for the correctness of `Create()`, because if there were enabled transformations, we
	// would need to disable them if they are not part of the plan (like we do in `Update()`).
	resp.Diagnostics.Append(r.checkAllDisabledBeforeCreation(ctx, data.ZoneID.ValueString())...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
=======
	// Prepare the data with disabled headers for those not in config
	preparedData, err := r.prepareDataWithDisabledHeaders(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to prepare data", err.Error())
		return
	}

	dataBytes, err := preparedData.MarshalJSON()
>>>>>>> Stashed changes
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{*preparedData}
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

<<<<<<< Updated upstream
	var plan *ManagedTransformsModel
	req.Plan.Get(ctx, &plan)

	normalizeResponse(data, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
=======
	// Filter the response to only include configured headers
	filteredData := r.filterResponseToConfigured(data, &env.Result)
	filteredData.ID = data.ZoneID

	resp.Diagnostics.Append(resp.State.Set(ctx, filteredData)...)
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream
	// Because we are patching the resource we need to explicitly add `enable = false` in order to remove
	// transformations. Otherwise, we will simply leave them with their int previous state.
	err := r.disableMissingTransformations(ctx, data)

	if err != nil {
		resp.Diagnostics.AddError("failed to disable missing transformations", err.Error())
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
=======
	// Prepare the data with disabled headers for those not in config
	preparedData, err := r.prepareDataWithDisabledHeaders(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to prepare data", err.Error())
		return
	}

	dataBytes, err := preparedData.MarshalJSONForUpdate(*state)
>>>>>>> Stashed changes
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{*preparedData}
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

<<<<<<< Updated upstream
	var plan *ManagedTransformsModel
	req.Plan.Get(ctx, &plan)

	normalizeResponse(data, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
=======
	// Filter the response to only include configured headers
	filteredData := r.filterResponseToConfigured(data, &env.Result)
	filteredData.ID = data.ZoneID

	resp.Diagnostics.Append(resp.State.Set(ctx, filteredData)...)
>>>>>>> Stashed changes
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
		return err
	}

	if res == nil {
		return nil
	}

	if plan.ManagedRequestHeaders != nil {
		var existingTransformations = []*ManagedTransformsManagedRequestHeadersModel{}

		for _, t := range res.ManagedRequestHeaders {
			existingTransformations = append(existingTransformations, &ManagedTransformsManagedRequestHeadersModel{
				ID: types.StringValue(t.ID),
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
				ID: types.StringValue(t.ID),
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

<<<<<<< Updated upstream
	var state *ManagedTransformsModel
	req.State.Get(ctx, &state)

	normalizeResponse(data, state)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
=======
	// Filter to only include enabled headers, preserving order from current state
	filteredData := &ManagedTransformsModel{
		ID:                     data.ZoneID,
		ZoneID:                 data.ZoneID,
		ManagedRequestHeaders:  &[]*ManagedTransformsManagedRequestHeadersModel{},
		ManagedResponseHeaders: &[]*ManagedTransformsManagedResponseHeadersModel{},
	}

	// Create maps of enabled headers from API
	enabledRequestHeaders := make(map[string]*ManagedTransformsManagedRequestHeadersModel)
	enabledResponseHeaders := make(map[string]*ManagedTransformsManagedResponseHeadersModel)

	if env.Result.ManagedRequestHeaders != nil {
		for _, h := range *env.Result.ManagedRequestHeaders {
			if h.Enabled.ValueBool() {
				enabledRequestHeaders[h.ID.ValueString()] = h
			}
		}
	}

	if env.Result.ManagedResponseHeaders != nil {
		for _, h := range *env.Result.ManagedResponseHeaders {
			if h.Enabled.ValueBool() {
				enabledResponseHeaders[h.ID.ValueString()] = h
			}
		}
	}

	// Preserve order from current state if available
	if data.ManagedRequestHeaders != nil && len(*data.ManagedRequestHeaders) > 0 {
		// Use order from state
		for _, stateHeader := range *data.ManagedRequestHeaders {
			if apiHeader, exists := enabledRequestHeaders[stateHeader.ID.ValueString()]; exists {
				*filteredData.ManagedRequestHeaders = append(*filteredData.ManagedRequestHeaders, apiHeader)
			}
		}
	} else {
		// No state, just add all enabled headers
		for _, h := range enabledRequestHeaders {
			*filteredData.ManagedRequestHeaders = append(*filteredData.ManagedRequestHeaders, h)
		}
	}

	if data.ManagedResponseHeaders != nil && len(*data.ManagedResponseHeaders) > 0 {
		// Use order from state
		for _, stateHeader := range *data.ManagedResponseHeaders {
			if apiHeader, exists := enabledResponseHeaders[stateHeader.ID.ValueString()]; exists {
				*filteredData.ManagedResponseHeaders = append(*filteredData.ManagedResponseHeaders, apiHeader)
			}
		}
	} else {
		// No state, just add all enabled headers
		for _, h := range enabledResponseHeaders {
			*filteredData.ManagedResponseHeaders = append(*filteredData.ManagedResponseHeaders, h)
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, filteredData)...)
>>>>>>> Stashed changes
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

func normalizeResponse(response *ManagedTransformsModel, state *ManagedTransformsModel) {
	if response.ManagedRequestHeaders != nil {
		stateManagedRequestHeaders := []*ManagedTransformsManagedRequestHeadersModel{}
		if state != nil && state.ManagedRequestHeaders != nil {
			stateManagedRequestHeaders = *state.ManagedRequestHeaders
		}

		t := transformationsView(*response.ManagedRequestHeaders, stateManagedRequestHeaders)
		response.ManagedRequestHeaders = &t
	}
	if response.ManagedResponseHeaders != nil {
		stateManagedResponseHeaders := []*ManagedTransformsManagedResponseHeadersModel{}
		if state != nil && state.ManagedRequestHeaders != nil {
			stateManagedResponseHeaders = *state.ManagedResponseHeaders
		}

		t := transformationsView(*response.ManagedResponseHeaders, stateManagedResponseHeaders)
		response.ManagedResponseHeaders = &t
	}
}

func transformationsView[T transformation](transformations []T, stateTransformations []T) []T {
	inState := make(map[string]bool)

	for _, transformation := range stateTransformations {
		inState[transformation.id()] = true
	}

	newTransformations := []T{}

	// We ignore transformations that are not in the default state (unless we have them explicitly in our state file).
	// This way terraform won't see a diffs where it doesn't exist: without this, it would see that transformations
	// in the default state (disabled) needed to be removed.
	for _, transformation := range transformations {
		isDefaultTransformation := !transformation.enabled()

		if inState[transformation.id()] || !isDefaultTransformation {
			newTransformations = append(newTransformations, transformation)
		}
	}

	return newTransformations
}

func (r *ManagedTransformsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ManagedTransformsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.ManagedTransforms.Delete(
		ctx,
		managed_transforms.ManagedTransformDeleteParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.ZoneID

	normalizeResponse(data, nil)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ManagedTransformsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ManagedTransformsModel = new(ManagedTransformsModel)

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

	normalizeResponse(data, nil)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ManagedTransformsResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}

// prepareDataWithDisabledHeaders fetches currently enabled headers and adds them as disabled if not in config
func (r *ManagedTransformsResource) prepareDataWithDisabledHeaders(ctx context.Context, data *ManagedTransformsModel) (*ManagedTransformsModel, error) {
	// First, get the list of all headers from the API to see what's currently enabled
	res := new(http.Response)
	env := ManagedTransformsResultEnvelope{}
	_, err := r.client.ManagedTransforms.List(
		ctx,
		managed_transforms.ManagedTransformListParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list managed transforms: %w", err)
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize response: %w", err)
	}

	// Create a copy of the data to modify
	result := &ManagedTransformsModel{
		ID:                     data.ID,
		ZoneID:                 data.ZoneID,
		ManagedRequestHeaders:  &[]*ManagedTransformsManagedRequestHeadersModel{},
		ManagedResponseHeaders: &[]*ManagedTransformsManagedResponseHeadersModel{},
	}

	// Copy configured headers
	if data.ManagedRequestHeaders != nil {
		headers := make([]*ManagedTransformsManagedRequestHeadersModel, len(*data.ManagedRequestHeaders))
		copy(headers, *data.ManagedRequestHeaders)
		result.ManagedRequestHeaders = &headers
	}
	if data.ManagedResponseHeaders != nil {
		headers := make([]*ManagedTransformsManagedResponseHeadersModel, len(*data.ManagedResponseHeaders))
		copy(headers, *data.ManagedResponseHeaders)
		result.ManagedResponseHeaders = &headers
	}

	// Track which headers are in the config
	configuredRequestHeaders := make(map[string]bool)
	configuredResponseHeaders := make(map[string]bool)

	if data.ManagedRequestHeaders != nil {
		for _, h := range *data.ManagedRequestHeaders {
			configuredRequestHeaders[h.ID.ValueString()] = true
		}
	}
	if data.ManagedResponseHeaders != nil {
		for _, h := range *data.ManagedResponseHeaders {
			configuredResponseHeaders[h.ID.ValueString()] = true
		}
	}

	// Add any enabled headers from API that aren't in config as disabled
	if env.Result.ManagedRequestHeaders != nil {
		for _, apiHeader := range *env.Result.ManagedRequestHeaders {
			if apiHeader.Enabled.ValueBool() && !configuredRequestHeaders[apiHeader.ID.ValueString()] {
				disabledHeader := &ManagedTransformsManagedRequestHeadersModel{
					ID:      apiHeader.ID,
					Enabled: types.BoolValue(false),
				}
				*result.ManagedRequestHeaders = append(*result.ManagedRequestHeaders, disabledHeader)
			}
		}
	}

	if env.Result.ManagedResponseHeaders != nil {
		for _, apiHeader := range *env.Result.ManagedResponseHeaders {
			if apiHeader.Enabled.ValueBool() && !configuredResponseHeaders[apiHeader.ID.ValueString()] {
				disabledHeader := &ManagedTransformsManagedResponseHeadersModel{
					ID:      apiHeader.ID,
					Enabled: types.BoolValue(false),
				}
				*result.ManagedResponseHeaders = append(*result.ManagedResponseHeaders, disabledHeader)
			}
		}
	}

	return result, nil
}

// filterResponseToConfigured filters the API response to only include headers that are in the config
// and preserves the order from the config
func (r *ManagedTransformsResource) filterResponseToConfigured(config *ManagedTransformsModel, apiResponse *ManagedTransformsModel) *ManagedTransformsModel {
	result := &ManagedTransformsModel{
		ID:                     apiResponse.ID,
		ZoneID:                 apiResponse.ZoneID,
		ManagedRequestHeaders:  &[]*ManagedTransformsManagedRequestHeadersModel{},
		ManagedResponseHeaders: &[]*ManagedTransformsManagedResponseHeadersModel{},
	}

	// Create maps of API response headers for quick lookup
	apiRequestHeaders := make(map[string]*ManagedTransformsManagedRequestHeadersModel)
	apiResponseHeaders := make(map[string]*ManagedTransformsManagedResponseHeadersModel)

	if apiResponse.ManagedRequestHeaders != nil {
		for _, h := range *apiResponse.ManagedRequestHeaders {
			apiRequestHeaders[h.ID.ValueString()] = h
		}
	}
	if apiResponse.ManagedResponseHeaders != nil {
		for _, h := range *apiResponse.ManagedResponseHeaders {
			apiResponseHeaders[h.ID.ValueString()] = h
		}
	}

	// Preserve order from config and use values from API response
	if config.ManagedRequestHeaders != nil {
		for _, configHeader := range *config.ManagedRequestHeaders {
			if apiHeader, exists := apiRequestHeaders[configHeader.ID.ValueString()]; exists {
				*result.ManagedRequestHeaders = append(*result.ManagedRequestHeaders, apiHeader)
			}
		}
	}

	if config.ManagedResponseHeaders != nil {
		for _, configHeader := range *config.ManagedResponseHeaders {
			if apiHeader, exists := apiResponseHeaders[configHeader.ID.ValueString()]; exists {
				*result.ManagedResponseHeaders = append(*result.ManagedResponseHeaders, apiHeader)
			}
		}
	}

	return result
}
