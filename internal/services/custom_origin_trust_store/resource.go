// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/acm"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*CustomOriginTrustStoreResource)(nil)
var _ resource.ResourceWithModifyPlan = (*CustomOriginTrustStoreResource)(nil)
var _ resource.ResourceWithImportState = (*CustomOriginTrustStoreResource)(nil)

func NewResource() resource.Resource {
	return &CustomOriginTrustStoreResource{}
}

// CustomOriginTrustStoreResource defines the resource implementation.
type CustomOriginTrustStoreResource struct {
	client *cloudflare.Client
}

func (r *CustomOriginTrustStoreResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_origin_trust_store"
}

func (r *CustomOriginTrustStoreResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomOriginTrustStoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CustomOriginTrustStoreModel

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
	env := CustomOriginTrustStoreResultEnvelope{*data}
	_, err = r.client.ACM.CustomTrustStore.New(
		ctx,
		acm.CustomTrustStoreNewParams{
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

	// Get the original config value to preserve its format
	var configData CustomOriginTrustStoreModel
	req.Config.Get(ctx, &configData)

	apiNormalized := strings.TrimRight(data.Certificate.ValueString(), "\n")
	configNormalized := strings.TrimRight(configData.Certificate.ValueString(), "\n")

	// If they match, use the config format
	if apiNormalized == configNormalized {
		data.Certificate = configData.Certificate
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomOriginTrustStoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *CustomOriginTrustStoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CustomOriginTrustStoreModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := CustomOriginTrustStoreResultEnvelope{*data}
	_, err := r.client.ACM.CustomTrustStore.Get(
		ctx,
		data.ID.ValueString(),
		acm.CustomTrustStoreGetParams{
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

	// Keep the original state format if they're semantically equal
	var stateData CustomOriginTrustStoreModel
	req.State.Get(ctx, &stateData)

	apiNormalized := strings.TrimRight(data.Certificate.ValueString(), "\n")
	stateNormalized := strings.TrimRight(stateData.Certificate.ValueString(), "\n")

	if apiNormalized == stateNormalized {
		data.Certificate = stateData.Certificate
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomOriginTrustStoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CustomOriginTrustStoreModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ACM.CustomTrustStore.Delete(
		ctx,
		data.ID.ValueString(),
		acm.CustomTrustStoreDeleteParams{
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

func (r *CustomOriginTrustStoreResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(CustomOriginTrustStoreModel)

	path_zone_id := ""
	path_custom_origin_trust_store_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<custom_origin_trust_store_id>",
		&path_zone_id,
		&path_custom_origin_trust_store_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.ID = types.StringValue(path_custom_origin_trust_store_id)

	res := new(http.Response)
	env := CustomOriginTrustStoreResultEnvelope{*data}
	_, err := r.client.ACM.CustomTrustStore.Get(
		ctx,
		path_custom_origin_trust_store_id,
		acm.CustomTrustStoreGetParams{
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

func (r *CustomOriginTrustStoreResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
