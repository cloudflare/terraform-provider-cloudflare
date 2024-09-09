// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_datasets

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*DLPDatasetsResource)(nil)
var _ resource.ResourceWithModifyPlan = (*DLPDatasetsResource)(nil)

func NewResource() resource.Resource {
	return &DLPDatasetsResource{}
}

// DLPDatasetsResource defines the resource implementation.
type DLPDatasetsResource struct {
	client *cloudflare.Client
}

func (r *DLPDatasetsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dlp_datasets"
}

func (r *DLPDatasetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DLPDatasetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DLPDatasetsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := DLPDatasetsResultEnvelope{*data}
	_, err = r.client.ZeroTrust.DLP.Datasets.New(
		ctx,
		zero_trust.DLPDatasetNewParams{
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

func (r *DLPDatasetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DLPDatasetsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *DLPDatasetsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.MarshalForUpdate(data, state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := DLPDatasetsResultEnvelope{*data}
	_, err = r.client.ZeroTrust.DLP.Datasets.Update(
		ctx,
		data.DatasetID.ValueString(),
		zero_trust.DLPDatasetUpdateParams{
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

func (r *DLPDatasetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DLPDatasetsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := DLPDatasetsResultEnvelope{*data}
	_, err := r.client.ZeroTrust.DLP.Datasets.Get(
		ctx,
		data.DatasetID.ValueString(),
		zero_trust.DLPDatasetGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DLPDatasetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DLPDatasetsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.ZeroTrust.DLP.Datasets.Delete(
		ctx,
		data.DatasetID.ValueString(),
		zero_trust.DLPDatasetDeleteParams{
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

func (r *DLPDatasetsResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
