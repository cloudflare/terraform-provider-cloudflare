// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ListResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ListResource)(nil)
var _ resource.ResourceWithImportState = (*ListResource)(nil)

func NewResource() resource.Resource {
	return &ListResource{}
}

// ListResource defines the resource implementation.
type ListResource struct {
	client *cloudflare.Client
}

func (r *ListResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list"
}

func (r *ListResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ListModel
	var plan *ListModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ListResultEnvelope{*data}
	_, err := r.client.Rules.Lists.New(
		ctx,
		rules.ListNewParams{
			AccountID:   cloudflare.F(data.AccountID.ValueString()),
			Kind:        cloudflare.F(rules.ListNewParamsKind(data.Kind.ValueString())),
			Name:        cloudflare.F(data.Name.ValueString()),
			Description: cloudflare.F(data.Description.ValueString()),
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

	if !plan.Items.IsNull() {
		diags := bulkUpdateList(ctx, r.client, data.AccountID.ValueString(), data.ID.ValueString(), plan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		itemsSet, diags := getAllListItems(ctx, r.client, data.AccountID.ValueString(), data.ID.ValueString())
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		var items customfield.NestedObjectSet[ListItemModel]

		items, diags = customfield.NewObjectSet[ListItemModel](ctx, itemsSet)
		resp.Diagnostics.Append(diags...)

		data.Items = items
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state *ListModel
	var data *ListModel
	var plan *ListModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ListResultEnvelope{*data}
	_, err := r.client.Rules.Lists.Update(
		ctx,
		data.ID.ValueString(),
		rules.ListUpdateParams{
			AccountID:   cloudflare.F(data.AccountID.ValueString()),
			Description: cloudflare.F(data.Description.ValueString()),
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

	if !plan.Items.Equal(state.Items) {
		diags := bulkUpdateList(ctx, r.client, data.AccountID.ValueString(), data.ID.ValueString(), plan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		itemsSet, diags := getAllListItems(ctx, r.client, data.AccountID.ValueString(), data.ID.ValueString())
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		var items customfield.NestedObjectSet[ListItemModel]

		if plan.Items.IsNull() && len(itemsSet) == 0 {
			items = customfield.NullObjectSet[ListItemModel](ctx)
		} else {
			items, diags = customfield.NewObjectSet[ListItemModel](ctx, itemsSet)
			resp.Diagnostics.Append(diags...)
		}
		data.Items = items
	} else {
		data.Items = state.Items
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ListModel
	var prev *ListModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &prev)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ListResultEnvelope{*data}
	_, err := r.client.Rules.Lists.Get(
		ctx,
		data.ID.ValueString(),
		rules.ListGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
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

	if !prev.Items.IsNull() {

		itemsSet, diags := getAllListItems(ctx, r.client, data.AccountID.ValueString(), data.ID.ValueString())
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		var items customfield.NestedObjectSet[ListItemModel]

		items, diags = customfield.NewObjectSet[ListItemModel](ctx, itemsSet)
		resp.Diagnostics.Append(diags...)

		data.Items = items
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ListModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Rules.Lists.Delete(
		ctx,
		data.ID.ValueString(),
		rules.ListDeleteParams{
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

func (r *ListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ListModel = new(ListModel)

	path_account_id := ""
	path_list_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<list_id>",
		&path_account_id,
		&path_list_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_list_id)

	res := new(http.Response)
	env := ListResultEnvelope{*data}
	_, err := r.client.Rules.Lists.Get(
		ctx,
		path_list_id,
		rules.ListGetParams{
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

	itemsSet, diags := getAllListItems(ctx, r.client, data.AccountID.ValueString(), data.ID.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var items customfield.NestedObjectSet[ListItemModel]

	if len(itemsSet) == 0 {
		items = customfield.NullObjectSet[ListItemModel](ctx)
	} else {
		items, diags = customfield.NewObjectSet[ListItemModel](ctx, itemsSet)
		resp.Diagnostics.Append(diags...)
	}
	data.Items = items

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ListResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}

func PollListBulkOperation(ctx context.Context, accountID, operationID string, client *cloudflare.Client) error {
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

func bulkUpdateList(ctx context.Context, client *cloudflare.Client, accountID, listID string, plan *ListModel) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	items, diags := plan.Items.AsStructSliceT(ctx)
	diagnostics.Append(diags...)
	if diagnostics.HasError() {
		return diagnostics
	}

	itemsBytes, err := apijson.MarshalRoot(items)
	if err != nil {
		diagnostics.AddError("failed to serialize http request", err.Error())
		return diagnostics
	}

	result, err := client.Rules.Lists.Items.Update(ctx, listID, rules.ListItemUpdateParams{
		AccountID: cloudflare.F(accountID),
	},
		option.WithRequestBody("application/json", itemsBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		diagnostics.AddError("failed to update list items", err.Error())
		return diagnostics
	}

	err = PollListBulkOperation(ctx, accountID, result.OperationID, client)
	if err != nil {
		diagnostics.AddError("list item bulk operation failed", err.Error())
		return diagnostics
	}

	return diagnostics
}

func getAllListItems(ctx context.Context, client *cloudflare.Client, accountID, listID string) ([]ListItemModel, diag.Diagnostics) {
	var diagnostics diag.Diagnostics
	paginatedItems := make([]ListItemModel, 0)

	listItems := client.Rules.Lists.Items.ListAutoPaging(
		ctx,
		listID,
		rules.ListItemListParams{
			AccountID: cloudflare.F(accountID),
			PerPage:   cloudflare.F(int64(500)),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if listItems.Err() != nil {
		diagnostics.AddError("failed to search list items", listItems.Err().Error())
		return paginatedItems, diagnostics
	}
	if listItems == nil {
		diagnostics.AddError("failed to search list items", "list item pagination was nil")
		return paginatedItems, diagnostics
	}

	for listItems.Next() {
		current := listItems.Current()

		var item ListItemModel
		err := apijson.UnmarshalRoot([]byte(current.JSON.RawJSON()), &item)
		if err != nil {
			diagnostics.AddError("failed to unmarshal list item", err.Error())
			return paginatedItems, diagnostics
		}

		paginatedItems = append(paginatedItems, item)
	}
	if listItems.Err() != nil {
		diagnostics.AddError("failed to paginate list items", listItems.Err().Error())
		return paginatedItems, diagnostics
	}
	return paginatedItems, diagnostics
}
