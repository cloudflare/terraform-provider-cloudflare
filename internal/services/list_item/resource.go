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

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	requestTimeout = 10 * time.Second
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
		option.WithRequestTimeout(requestTimeout),
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

	err = list.PollListBulkOperation(ctx, data.AccountID.ValueString(), createEnv.Result.OperationID.ValueString(), r.client)
	if err != nil {
		resp.Diagnostics.AddError("list item bulk operation failed", err.Error())
		return
	}

	searchTerm := getSearchTerm(data)
	listItems := r.client.Rules.Lists.Items.ListAutoPaging(
		ctx,
		data.ListID.ValueString(),
		rules.ListItemListParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Search:    cloudflare.F(searchTerm),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
		option.WithRequestTimeout(requestTimeout),
	)
	if listItems.Err() != nil {
		resp.Diagnostics.AddError("failed to search list items", listItems.Err().Error())
		return
	}
	if listItems == nil {
		resp.Diagnostics.AddWarning("failed to search list items", "list item pagination was nil")
	}

	// find the actual list item, don't rely on the response to have the first entry be the correct one
	var listItemID string
	for listItems.Next() {
		item := listItems.Current()
		if matchedItemID, ok := listItemMatchesOriginal(data, item); ok {
			listItemID = matchedItemID
			break
		}
	}

	if listItemID == "" {
		resp.Diagnostics.AddError("failed to find list item", "list item pagination did not return a matching list item")
		return
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
		option.WithRequestTimeout(requestTimeout),
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

	// If ID is empty, null, or unknown, the resource doesn't exist yet.
	// This occurs when terraform/tofu runs refresh before create for new resources.
	// Return early to signal that the resource needs to be created.
	if data.ID.IsNull() || data.ID.IsUnknown() || data.ID.ValueString() == "" {
		resp.State.RemoveResource(ctx)
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
		option.WithRequestTimeout(requestTimeout),
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
		option.WithRequestTimeout(requestTimeout),
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

func listItemMatchesOriginal(original *ListItemModel, item rules.ListItemListResponse) (string, bool) {
	if original.IP.ValueString() != item.IP {
		return "", false
	}

	if original.ASN.ValueInt64() != item.ASN {
		return "", false
	}

	if !original.Hostname.IsNull() && !hostnameEqual(original.Hostname, item.Hostname) {
		return "", false
	}

	if !original.Redirect.IsNull() && !redirectEqual(original.Redirect, item.Redirect) {
		return "", false
	}

	return item.ID, true
}

func hostnameEqual(original customfield.NestedObject[ListItemHostnameModel], item rules.Hostname) bool {
	originalVal, err := original.Value(context.TODO())
	if err != nil {
		return false
	}

	if originalVal.URLHostname.ValueString() != item.URLHostname {
		return false
	}

	if originalVal.ExcludeExactHostname.ValueBool() != item.ExcludeExactHostname {
		return false
	}

	return true
}

func redirectEqual(original customfield.NestedObject[ListItemRedirectModel], item rules.Redirect) bool {
	originalVal, err := original.Value(context.TODO())
	if err != nil {
		return false
	}

	if originalVal.SourceURL.ValueString() != item.SourceURL {
		return false
	}

	if originalVal.TargetURL.ValueString() != item.TargetURL {
		return false
	}

	if originalVal.IncludeSubdomains.ValueBool() != item.IncludeSubdomains {
		return false
	}

	if originalVal.PreservePathSuffix.ValueBool() != item.PreservePathSuffix {
		return false
	}

	if originalVal.PreserveQueryString.ValueBool() != item.PreserveQueryString {
		return false
	}

	if originalVal.StatusCode.ValueInt64() != int64(item.StatusCode) {
		return false
	}

	if originalVal.SubpathMatching.ValueBool() != item.SubpathMatching {
		return false
	}

	return true
}
