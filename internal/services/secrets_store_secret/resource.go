package secrets_store_secret

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigure = (*SecretsStoreSecretResource)(nil)
var _ resource.ResourceWithImportState = (*SecretsStoreSecretResource)(nil)

func NewResource() resource.Resource {
	return &SecretsStoreSecretResource{}
}

type SecretsStoreSecretResource struct {
	client *cloudflare.Client
}

func (r *SecretsStoreSecretResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secrets_store_secret"
}

func (r *SecretsStoreSecretResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SecretsStoreSecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SecretsStoreSecretModel

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
	env := SecretsStoreSecretResultEnvelope{*data}

	_, err = r.client.SecretsStore.Stores.Secrets.New(
		ctx,
		data.StoreID.ValueString(),
		secrets_store.StoreSecretNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Name:      cloudflare.F(data.Name.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create secrets store secret", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize response", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretsStoreSecretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SecretsStoreSecretModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *SecretsStoreSecretModel

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
	env := SecretsStoreSecretResultEnvelope{*data}

	_, err = r.client.SecretsStore.Stores.Secrets.Update(
		ctx,
		state.StoreID.ValueString(),
		state.ID.ValueString(),
		secrets_store.StoreSecretUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to update secrets store secret", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize response", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretsStoreSecretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SecretsStoreSecretModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := SecretsStoreSecretResultEnvelope{*data}

	secrets, err := r.client.SecretsStore.Stores.Secrets.List(
		ctx,
		data.StoreID.ValueString(),
		secrets_store.StoreSecretListParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to list secrets store secrets", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize response", err.Error())
		return
	}

	found := false
	for _, secret := range secrets.Result {
		if secret.ID == data.ID.ValueString() {
			data.ID = types.StringValue(secret.ID)
			data.Name = types.StringValue(secret.Name)
			data.Status = types.StringValue(secret.Status)
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddWarning("Resource not found", "The secrets store secret was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretsStoreSecretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SecretsStoreSecretModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SecretsStore.Stores.Secrets.Delete(
		ctx,
		data.StoreID.ValueString(),
		data.ID.ValueString(),
		secrets_store.StoreSecretDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete secrets store secret", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretsStoreSecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(SecretsStoreSecretModel)

	path_account_id := ""
	path_store_id := ""
	path_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<store_id>/<id>",
		&path_account_id,
		&path_store_id,
		&path_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.StoreID = types.StringValue(path_store_id)
	data.ID = types.StringValue(path_id)

	res := new(http.Response)
	env := SecretsStoreSecretResultEnvelope{*data}

	secrets, err := r.client.SecretsStore.Stores.Secrets.List(
		ctx,
		path_store_id,
		secrets_store.StoreSecretListParams{
			AccountID: cloudflare.F(path_account_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to list secrets store secrets", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize response", err.Error())
		return
	}

	found := false
	for _, secret := range secrets.Result {
		if secret.ID == path_id {
			data.ID = types.StringValue(secret.ID)
			data.Name = types.StringValue(secret.Name)
			data.Status = types.StringValue(secret.Status)
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddWarning("Resource not found", "The secrets store secret was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
