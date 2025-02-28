// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ListItemResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ListItemResource)(nil)

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

	searchTerm := getSearchTerm(data)
	findItemRes := new(http.Response)
	_, err = r.client.Rules.Lists.Items.List(
		ctx,
		data.ListID.ValueString(),
		rules.ListItemListParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
			Search:    cloudflare.F(searchTerm),
		},
		option.WithResponseBodyInto(&findItemRes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to fetch individual list item", err.Error())
		return
	}
	findListItem, _ := io.ReadAll(findItemRes.Body)
	itemID := gjson.Get(string(findListItem), "result.0.id")
	data.ID = types.StringValue(itemID.String())

	env := ListItemResultEnvelope{*data}
	listItemRes := new(http.Response)
	_, err = r.client.Rules.Lists.Items.Get(
		ctx,
		data.AccountID.ValueString(),
		data.ListID.ValueString(),
		data.ID.ValueString(),
		option.WithResponseBodyInto(&listItemRes),
		option.WithMiddleware(logging.Middleware(ctx)),
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
	var data *ListItemModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ListItemModel

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
	env := ListItemResultEnvelope{*data}
	_, err = r.client.Rules.Lists.Items.Update(
		ctx,
		data.ListID.ValueString(),
		rules.ListItemUpdateParams{
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
		data.AccountID.ValueString(),
		data.ListID.ValueString(),
		data.ID.ValueString(),
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
	err = apijson.UnmarshalComputed(bytes, &env)
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
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

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
