package secrets_store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/secrets_store"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigure = (*SecretsStoreResource)(nil)
var _ resource.ResourceWithImportState = (*SecretsStoreResource)(nil)

func NewResource() resource.Resource {
	return &SecretsStoreResource{}
}

type SecretsStoreResource struct {
	client *cloudflare.Client
}

func (r *SecretsStoreResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secrets_store"
}

func (r *SecretsStoreResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SecretsStoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SecretsStoreModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Workaround for SDK bug: StoreNewParams.MarshalJSON() sends array instead of object
	// See: https://github.com/cloudflare/cloudflare-go/issues/4244
	result, err := r.createStoreWithCorrectJSON(ctx, data.AccountID.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to create secrets store", err.Error())
		return
	}

	data.ID = types.StringValue(result.ID)
	data.Name = types.StringValue(result.Name)
	data.CreatedAt = types.StringValue(result.Created)
	data.ModifiedAt = types.StringValue(result.Modified)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// storeCreateResponse is the API response for creating a store
type storeCreateResponse struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	Created   string `json:"created"`
	Modified  string `json:"modified"`
}

type storeCreateResponseEnvelope struct {
	Result   storeCreateResponse `json:"result"`
	Success  bool                `json:"success"`
	Errors   []interface{}       `json:"errors"`
	Messages []interface{}       `json:"messages"`
}

// createStoreWithCorrectJSON makes a direct API call with correctly formatted JSON
// to work around the SDK's StoreNewParams.MarshalJSON() bug
//
// We use client.Post with a []byte body to bypass the broken MarshalJSON on StoreNewParams
func (r *SecretsStoreResource) createStoreWithCorrectJSON(ctx context.Context, accountID, name string) (*storeCreateResponse, error) {
	path := fmt.Sprintf("accounts/%s/secrets_store/stores", accountID)

	// Create body with correct JSON structure (not wrapped in array)
	body := map[string]string{"name": name}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	var result storeCreateResponseEnvelope
	// Use client.Post which internally uses Execute -> requestconfig.ExecuteNewRequest
	// But since we pass []byte, it will be used as-is without MarshalJSON
	err = r.client.Post(ctx, path, bodyBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to create secrets store: %w", err)
	}
	if !result.Success {
		return nil, fmt.Errorf("API error: %v", result.Errors)
	}
	return &result.Result, nil
}

func (r *SecretsStoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SecretsStoreModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError("updates not supported", "Secrets Store does not support updates. Please update the resource by using `terraform taint` and applying changes.")
}

func (r *SecretsStoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SecretsStoreModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	stores, err := r.client.SecretsStore.Stores.List(
		ctx,
		secrets_store.StoreListParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to list secrets stores", err.Error())
		return
	}

	found := false
	for _, store := range stores.Result {
		if store.ID == data.ID.ValueString() {
			data.ID = types.StringValue(store.ID)
			data.Name = types.StringValue(store.Name)
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddWarning("Resource not found", "The secrets store was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretsStoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SecretsStoreModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SecretsStore.Stores.Delete(
		ctx,
		data.ID.ValueString(),
		secrets_store.StoreDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete secrets store", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretsStoreResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(SecretsStoreModel)

	path_account_id := ""
	path_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<id>",
		&path_account_id,
		&path_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_id)

	stores, err := r.client.SecretsStore.Stores.List(
		ctx,
		secrets_store.StoreListParams{
			AccountID: cloudflare.F(path_account_id),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to list secrets stores", err.Error())
		return
	}

	found := false
	for _, store := range stores.Result {
		if store.ID == path_id {
			data.ID = types.StringValue(store.ID)
			data.Name = types.StringValue(store.Name)
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddWarning("Resource not found", "The secrets store was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
