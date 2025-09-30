// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*AccountSubscriptionResource)(nil)
var _ resource.ResourceWithModifyPlan = (*AccountSubscriptionResource)(nil)
var _ resource.ResourceWithImportState = (*AccountSubscriptionResource)(nil)

func NewResource() resource.Resource {
	return &AccountSubscriptionResource{}
}

// AccountSubscriptionResource defines the resource implementation.
type AccountSubscriptionResource struct {
	client *cloudflare.Client
}

func (r *AccountSubscriptionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_subscription"
}

func (r *AccountSubscriptionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AccountSubscriptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AccountSubscriptionModel

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
	env := AccountSubscriptionResultEnvelope{*data}
	_, err = r.client.Accounts.Subscriptions.New(
		ctx,
		accounts.SubscriptionNewParams{
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

func (r *AccountSubscriptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AccountSubscriptionModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AccountSubscriptionModel

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
	env := AccountSubscriptionResultEnvelope{*data}
	_, err = r.client.Accounts.Subscriptions.Update(
		ctx,
		data.ID.ValueString(),
		accounts.SubscriptionUpdateParams{
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

func (r *AccountSubscriptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccountSubscriptionModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Store the account ID for later use
	accountID := data.AccountID.ValueString()
	subscriptionID := data.ID.ValueString()

	res := new(http.Response)

	// Define a custom response envelope for list response
	type SubscriptionListResponse struct {
		Result  []AccountSubscriptionModel `json:"result"`
		Success bool                       `json:"success"`
	}

	listResponse := &SubscriptionListResponse{}

	_, err := r.client.Accounts.Subscriptions.Get(
		ctx,
		accounts.SubscriptionGetParams{
			AccountID: cloudflare.F(accountID),
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
	err = apijson.Unmarshal(bytes, &listResponse)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	// Find the subscription with matching ID
	found := false
	for _, subscription := range listResponse.Result {
		if subscription.ID.ValueString() == subscriptionID {
			// Found the subscription, update the data

			// Store the original rate_plan for comparison
			//originalRatePlan := data.RatePlan
			*data = subscription

			// Ensure account ID is set correctly
			data.AccountID = types.StringValue(accountID)

			// Add debug logging to see the difference
			// fmt.Printf("SUI - Read: Original rate_plan: %+v\n", originalRatePlan)
			// fmt.Printf("SUI - Read: API rate_plan: %+v\n", data.RatePlan)

			found = true
			break
		}
	}

	if !found && subscriptionID != "" {
		resp.Diagnostics.AddWarning("Resource not found", "The subscription with ID "+subscriptionID+" was not found and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountSubscriptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AccountSubscriptionModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Accounts.Subscriptions.Delete(
		ctx,
		data.ID.ValueString(),
		accounts.SubscriptionDeleteParams{
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

func (r *AccountSubscriptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *AccountSubscriptionModel = new(AccountSubscriptionModel)

	path_account_id := ""
	path_subscription_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<subscription_id>",
		&path_account_id,
		&path_subscription_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_subscription_id)

	res := new(http.Response)
	// Define a custom response envelope for list response
	type SubscriptionListResponse struct {
		Result  []AccountSubscriptionModel `json:"result"`
		Success bool                       `json:"success"`
	}

	listResponse := &SubscriptionListResponse{}

	_, err := r.client.Accounts.Subscriptions.Get(
		ctx,
		accounts.SubscriptionGetParams{
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
	err = apijson.Unmarshal(bytes, &listResponse)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	// Find the subscription with matching ID
	found := false
	for _, subscription := range listResponse.Result {
		if subscription.ID.ValueString() == path_subscription_id {
			// Found the subscription, update the data
			*data = subscription

			// Ensure account ID is set correctly
			data.AccountID = types.StringValue(path_account_id)

			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddError(
			"Subscription not found",
			fmt.Sprintf("Subscription with ID %s not found in account %s", data.ID.ValueString(), data.AccountID.ValueString()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountSubscriptionResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
