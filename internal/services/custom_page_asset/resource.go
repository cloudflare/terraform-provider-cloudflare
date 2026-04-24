// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_page_asset

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_pages"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*CustomPageAssetResource)(nil)
var _ resource.ResourceWithModifyPlan = (*CustomPageAssetResource)(nil)
var _ resource.ResourceWithImportState = (*CustomPageAssetResource)(nil)

func NewResource() resource.Resource {
	return &CustomPageAssetResource{}
}

// CustomPageAssetResource defines the resource implementation.
type CustomPageAssetResource struct {
	client *cloudflare.Client
}

func (r *CustomPageAssetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_page_asset"
}

func (r *CustomPageAssetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomPageAssetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CustomPageAssetModel

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
	env := CustomPageAssetResultEnvelope{*data}
	params := custom_pages.AssetNewParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.CustomPages.Assets.New(
		ctx,
		params,
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPageAssetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CustomPageAssetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *CustomPageAssetModel

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
	env := CustomPageAssetResultEnvelope{*data}
	params := custom_pages.AssetUpdateParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.CustomPages.Assets.Update(
		ctx,
		data.Name.ValueString(),
		params,
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPageAssetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CustomPageAssetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := CustomPageAssetResultEnvelope{*data}
	params := custom_pages.AssetGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.CustomPages.Assets.Get(
		ctx,
		data.Name.ValueString(),
		params,
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPageAssetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CustomPageAssetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := custom_pages.AssetDeleteParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	err := r.client.CustomPages.Assets.Delete(
		ctx,
		data.Name.ValueString(),
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPageAssetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(CustomPageAssetModel)
	params := custom_pages.AssetGetParams{}

	path_accounts_or_zones, path_account_id_or_zone_id := "", ""
	path_asset_name := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<{accounts|zones}/{account_id|zone_id}>/<asset_name>",
		&path_accounts_or_zones,
		&path_account_id_or_zone_id,
		&path_asset_name,
	)
	resp.Diagnostics.Append(diags...)
	switch path_accounts_or_zones {
	case "accounts":
		params.AccountID = cloudflare.F(path_account_id_or_zone_id)
		data.AccountID = types.StringValue(path_account_id_or_zone_id)
	case "zones":
		params.ZoneID = cloudflare.F(path_account_id_or_zone_id)
		data.ZoneID = types.StringValue(path_account_id_or_zone_id)
	default:
		resp.Diagnostics.AddError("invalid discriminator segment - <{accounts|zones}/{account_id|zone_id}>", "expected discriminator to be one of {accounts|zones}")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	data.Name = types.StringValue(path_asset_name)

	res := new(http.Response)
	env := CustomPageAssetResultEnvelope{*data}
	_, err := r.client.CustomPages.Assets.Get(
		ctx,
		path_asset_name,
		params,
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
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPageAssetResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
