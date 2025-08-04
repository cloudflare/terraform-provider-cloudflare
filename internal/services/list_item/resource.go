// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ListItemResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ListItemResource)(nil)
var _ resource.ResourceWithImportState = (*ListItemResource)(nil)

func NewResource() resource.Resource {
	return &ListItemResource{}
}

// ListItemResource defines the resource implementation.
type ListItemResource struct {
	client *cloudflare.Client
}

func (r *ListItemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list_item"
}

func (r *ListItemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ListItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ListItemModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}

	wrappedBytes := data.MarshalSingleToCollectionJSON(dataBytes)

	res := new(http.Response)
	createEnv := ListItemResultEnvelope{*data}
	_, err = r.client.Rules.Lists.Items.New(
		ctx,
		data.ListID.ValueString(),
		rules.ListItemNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", wrappedBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
		option.WithRequestTimeout(time.Second*3),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)

	err = apijson.UnmarshalComputed(bytes, &createEnv)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	err = pollBulkOperation(ctx, data.AccountID.ValueString(), createEnv.Result.OperationID.ValueString(), r.client)
	if err != nil {
		resp.Diagnostics.AddError("list item bulk operation failed", err.Error())
		return
	}

	searchTerm := getSearchTerm(data)
	findItemRes := new(http.Response)
	listItems, err := r.client.Rules.Lists.Items.List(
		ctx,
		data.ListID.ValueString(),
		rules.ListItemListParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Search:    cloudflare.F(searchTerm),
			// TODO: when pagination is fixed in the API schema (and go sdk) we should not need to set this (items we are looking for are expected to be sorted near the top of the result list)
			PerPage: cloudflare.Int(500),
		},
		option.WithResponseBodyInto(&findItemRes),
		option.WithMiddleware(logging.Middleware(ctx)),
		option.WithRequestTimeout(time.Second*3),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch individual list item", err.Error())
		return
	}
	if listItems == nil {
		resp.Diagnostics.AddWarning("failed to fetch individual list item", "list item pagination was nil")
	}

	listItemsBytes, _ := io.ReadAll(findItemRes.Body)

	// TODO: when pagination is fixed in the API schema (and go sdk) this should paginate properly
	var apiResult pagination.SinglePage[ListItemModel]
	err = apijson.Unmarshal(listItemsBytes, &apiResult)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch individual list item", err.Error())
	}

	// find the actual list item, don't rely on the response to have the first entry be the correct one
	var listItemID string
	for _, item := range apiResult.Result {
		if matchedItemID, ok := listItemMatchesOriginal(data, item); ok {
			listItemID = matchedItemID
			break
		}
	}

	data.ID = types.StringValue(listItemID)

	env := ListItemResultEnvelope{*data}
	listItemRes := new(http.Response)
	_, err = r.client.Rules.Lists.Items.Get(
		ctx,
		data.ListID.ValueString(),
		data.ID.ValueString(),
		rules.ListItemGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&listItemRes),
		option.WithMiddleware(logging.Middleware(ctx)),
		option.WithRequestTimeout(time.Second*3),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch individual list item", err.Error())
		return
	}
	listItem, _ := io.ReadAll(listItemRes.Body)
	err = apijson.UnmarshalComputed(listItem, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("update is not supported for list items", "")
}

func (r *ListItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ListItemModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ListItemResultEnvelope{*data}
	_, err := r.client.Rules.Lists.Items.Get(
		ctx,
		data.ListID.ValueString(),
		data.ID.ValueString(),
		rules.ListItemGetParams{AccountID: cloudflare.F(data.AccountID.ValueString())},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
		option.WithRequestTimeout(time.Second*3),
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ListItemModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	deletePayload := bodyDeletePayload{
		Items: []bodyDeleteItems{{
			ID: data.ID.ValueString(),
		}},
	}
	deleteBody, _ := json.Marshal(deletePayload)

	_, err := r.client.Rules.Lists.Items.Delete(
		ctx,
		data.ListID.ValueString(),
		rules.ListItemDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
		option.WithRequestBody("application/json", deleteBody),
		option.WithRequestTimeout(time.Second*3),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ListItemModel = new(ListItemModel)

	path_account_id := ""
	path_list_id := ""
	path_item_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<list_id>/<item_id>",
		&path_account_id,
		&path_list_id,
		&path_item_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ListID = types.StringValue(path_list_id)
	data.ID = types.StringValue(path_item_id)

	res := new(http.Response)
	env := ListItemResultEnvelope{*data}
	_, err := r.client.Rules.Lists.Items.Get(
		ctx,
		path_list_id,
		path_item_id,
		rules.ListItemGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListItemResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}

func pollBulkOperation(ctx context.Context, accountID, operationID string, client *cloudflare.Client) error {
	backoff := 1 * time.Second
	maxBackoff := 30 * time.Second

	for {
		bulkOperation, err := client.Rules.Lists.BulkOperations.Get(
			ctx,
			operationID,
			rules.ListBulkOperationGetParams{
				AccountID: cloudflare.F(accountID),
			},
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			return err
		}
		switch bulkOperation.Status {
		case rules.ListBulkOperationGetResponseStatusCompleted:
			return nil
		case rules.ListBulkOperationGetResponseStatusFailed:
			return fmt.Errorf("failed to create list item: %s", bulkOperation.Error)
		default:
			time.Sleep(backoff)
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
		}
	}
}

type bodyDeletePayload struct {
	Items []bodyDeleteItems `json:"items"`
}

type bodyDeleteItems struct {
	ID string `json:"id"`
}

// getSearchTerm takes the schema and works out which "type" we are looking for
// and returns it.
func getSearchTerm(d *ListItemModel) string {
	if d.IP.ValueString() != "" {
		return d.IP.ValueString()
	}

	if d.ASN.ValueInt64() > 0 {
		return strconv.Itoa(int(d.ASN.ValueInt64()))
	}

	if !d.Hostname.IsNull() {
		if h, _ := d.Hostname.Value(context.TODO()); h.URLHostname.ValueString() != "" {
			return h.URLHostname.ValueString()
		}
	}

	if !d.Redirect.IsNull() {
		if r, _ := d.Redirect.Value(context.TODO()); r.SourceURL.ValueString() != "" {
			return r.SourceURL.ValueString()
		}
	}

	return ""
}

func listItemMatchesOriginal(original *ListItemModel, item ListItemModel) (string, bool) {
	if original.IP != item.IP {
		return "", false
	}

	if original.ASN != item.ASN {
		return "", false
	}

	if !original.Hostname.IsNull() && !item.Hostname.IsNull() && !original.Hostname.Equal(item.Hostname) {
		return "", false
	}

	if !original.Redirect.IsNull() && !item.Redirect.IsNull() && !original.Redirect.Equal(item.Redirect) {
		return "", false
	}

	return item.ID.ValueString(), true
}
