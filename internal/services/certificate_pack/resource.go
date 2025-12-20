// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*CertificatePackResource)(nil)
var _ resource.ResourceWithModifyPlan = (*CertificatePackResource)(nil)
var _ resource.ResourceWithImportState = (*CertificatePackResource)(nil)

func NewResource() resource.Resource {
	return &CertificatePackResource{}
}

// CertificatePackResource defines the resource implementation.
type CertificatePackResource struct {
	client *cloudflare.Client
}

func (r *CertificatePackResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_pack"
}

func (r *CertificatePackResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CertificatePackResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CertificatePackModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve plan values that the API may not return
	planData := *data

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := CertificatePackResultEnvelope{*data}
	_, err = r.client.SSL.CertificatePacks.New(
		ctx,
		ssl.CertificatePackNewParams{
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

	// Restore plan values if API didn't return them, or initialize computed fields to empty
	if data.CloudflareBranding.IsUnknown() || data.CloudflareBranding.IsNull() {
		if !planData.CloudflareBranding.IsNull() && !planData.CloudflareBranding.IsUnknown() {
			// Use plan value if it was explicitly set
			data.CloudflareBranding = planData.CloudflareBranding
		} else {
			// Keep as null if not set in plan and not returned by API
			data.CloudflareBranding = types.BoolNull()
		}
	}
	// Initialize computed list fields to empty lists if API didn't return them
	if data.Certificates.IsNull() {
		data.Certificates, _ = customfield.NewObjectList[CertificatePackCertificatesModel](ctx, []CertificatePackCertificatesModel{})
	}
	if data.ValidationErrors.IsNull() {
		data.ValidationErrors, _ = customfield.NewObjectList[CertificatePackValidationErrorsModel](ctx, []CertificatePackValidationErrorsModel{})
	}
	if data.ValidationRecords.IsNull() {
		data.ValidationRecords, _ = customfield.NewObjectList[CertificatePackValidationRecordsModel](ctx, []CertificatePackValidationRecordsModel{})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificatePackResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *CertificatePackResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CertificatePackModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve state values that the API may not return
	stateData := *data

	res := new(http.Response)
	env := CertificatePackResultEnvelope{*data}
	_, err := r.client.SSL.CertificatePacks.Get(
		ctx,
		data.ID.ValueString(),
		ssl.CertificatePackGetParams{
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

	// Preserve state values for optional+computed fields
	// If cloudflare_branding was null in state (user didn't set it), keep it null even if API returns a value
	if !data.CloudflareBranding.IsNull() && stateData.CloudflareBranding.IsNull() {
		data.CloudflareBranding = types.BoolNull()
	} else if data.CloudflareBranding.IsNull() && !stateData.CloudflareBranding.IsNull() {
		// If API didn't return it but state had it, preserve state
		data.CloudflareBranding = stateData.CloudflareBranding
	}
	// For computed list fields, preserve state if API didn't return them
	if data.Certificates.IsNull() && !stateData.Certificates.IsNull() {
		data.Certificates = stateData.Certificates
	}
	if data.ValidationErrors.IsNull() && !stateData.ValidationErrors.IsNull() {
		data.ValidationErrors = stateData.ValidationErrors
	}
	if data.ValidationRecords.IsNull() && !stateData.ValidationRecords.IsNull() {
		data.ValidationRecords = stateData.ValidationRecords
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificatePackResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CertificatePackModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SSL.CertificatePacks.Delete(
		ctx,
		data.ID.ValueString(),
		ssl.CertificatePackDeleteParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificatePackResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(CertificatePackModel)

	path_zone_id := ""
	path_certificate_pack_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<certificate_pack_id>",
		&path_zone_id,
		&path_certificate_pack_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.ID = types.StringValue(path_certificate_pack_id)

	res := new(http.Response)
	env := CertificatePackResultEnvelope{*data}
	_, err := r.client.SSL.CertificatePacks.Get(
		ctx,
		path_certificate_pack_id,
		ssl.CertificatePackGetParams{
			ZoneID: cloudflare.F(path_zone_id),
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

func (r *CertificatePackResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If the resource is being destroyed or there's no state, nothing to do
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan, state, config *CertificatePackModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Handle computed fields that may not be returned immediately by the API
	// but are populated on subsequent reads. Preserve state values to prevent drift.
	modified := false

	// Only preserve cloudflare_branding from state if it was explicitly set in config
	if plan.CloudflareBranding.IsNull() && !state.CloudflareBranding.IsNull() && !config.CloudflareBranding.IsNull() {
		plan.CloudflareBranding = state.CloudflareBranding
		modified = true
	}

	if plan.Certificates.IsNull() && !state.Certificates.IsNull() {
		plan.Certificates = state.Certificates
		modified = true
	}

	if plan.ValidationErrors.IsNull() && !state.ValidationErrors.IsNull() {
		plan.ValidationErrors = state.ValidationErrors
		modified = true
	}

	if plan.ValidationRecords.IsNull() && !state.ValidationRecords.IsNull() {
		plan.ValidationRecords = state.ValidationRecords
		modified = true
	}

	// Handle hosts list to avoid unnecessary replacements due to API behavior
	// The API may:
	// 1. Return hosts in a different order than submitted
	// 2. Add additional hosts (e.g., cloudflaressl.com subdomain when cloudflare_branding=true)
	if !plan.Hosts.IsNull() && !state.Hosts.IsNull() {
		planSet := make(map[string]bool)
		stateSet := make(map[string]bool)

		for _, h := range plan.Hosts.Elements() {
			if str, ok := h.(types.String); ok && !str.IsNull() {
				planSet[str.ValueString()] = true
			}
		}
		for _, h := range state.Hosts.Elements() {
			if str, ok := h.(types.String); ok && !str.IsNull() {
				stateSet[str.ValueString()] = true
			}
		}

		// Check if plan hosts are a subset of state hosts (API may have added extra hosts)
		// If all plan hosts exist in state, use state's hosts to prevent replacement
		allPlanHostsInState := true
		for host := range planSet {
			if !stateSet[host] {
				allPlanHostsInState = false
				break
			}
		}

		if allPlanHostsInState {
			// Use state's hosts list to prevent unnecessary replacement
			plan.Hosts = state.Hosts
			modified = true
		}
	}

	if modified {
		resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
	}
}

func mapsEqual(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}
