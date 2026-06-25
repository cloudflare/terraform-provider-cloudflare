// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/load_balancers"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*LoadBalancerPoolResource)(nil)
var _ resource.ResourceWithModifyPlan = (*LoadBalancerPoolResource)(nil)
var _ resource.ResourceWithImportState = (*LoadBalancerPoolResource)(nil)

func NewResource() resource.Resource {
	return &LoadBalancerPoolResource{}
}

// LoadBalancerPoolResource defines the resource implementation.
type LoadBalancerPoolResource struct {
	client *cloudflare.Client
}

func (r *LoadBalancerPoolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancer_pool"
}

func (r *LoadBalancerPoolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LoadBalancerPoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *LoadBalancerPoolModel

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
	env := LoadBalancerPoolResultEnvelope{*data}
	_, err = r.client.LoadBalancers.Pools.New(
		ctx,
		load_balancers.PoolNewParams{
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

func (r *LoadBalancerPoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *LoadBalancerPoolModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *LoadBalancerPoolModel

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
	env := LoadBalancerPoolResultEnvelope{*data}
	_, err = r.client.LoadBalancers.Pools.Update(
		ctx,
		data.ID.ValueString(),
		load_balancers.PoolUpdateParams{
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

	// Reset created_on to the value from state
	// The API seems to return 0001-01-01T00:00:00Z on update
	data.CreatedOn = state.CreatedOn

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerPoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *LoadBalancerPoolModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := LoadBalancerPoolResultEnvelope{*data}
	_, err := r.client.LoadBalancers.Pools.Get(
		ctx,
		data.ID.ValueString(),
		load_balancers.PoolGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerPoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *LoadBalancerPoolModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.LoadBalancers.Pools.Delete(
		ctx,
		data.ID.ValueString(),
		load_balancers.PoolDeleteParams{
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

func (r *LoadBalancerPoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(LoadBalancerPoolModel)

	path_account_id := ""
	path_pool_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<pool_id>",
		&path_account_id,
		&path_pool_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_pool_id)

	res := new(http.Response)
	env := LoadBalancerPoolResultEnvelope{*data}
	_, err := r.client.LoadBalancers.Pools.Get(
		ctx,
		path_pool_id,
		load_balancers.PoolGetParams{
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

// ModifyPlan suppresses no-op churn the Cloudflare API causes on this resource:
// it returns origins in a different order than config, and omits several computed
// fields on read so they replan as "(known after apply)".
//
// It rebuilds origins in state order and carries the unsynced computed values
// from state, then checks whether the plan now matches state exactly. If it does
// the run was pure churn and modified_on is pinned too, yielding no diff. If it
// does not there is a real change, so the original origin order and modified_on
// are restored and the change plans normally (reordering a real change would put
// a value at a position matching neither config nor prior, which Terraform
// rejects as an invalid plan).
func (r *LoadBalancerPoolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan, state *LoadBalancerPoolModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || plan == nil || state == nil {
		return
	}

	origOrigins := plan.Origins
	origModifiedOn := plan.ModifiedOn

	// Rebuild origins in state order so the API's reordering is a no-op.
	if plan.Origins != nil && state.Origins != nil && originAddressSetsEqual(*plan.Origins, *state.Origins) {
		stateOrigins := *state.Origins
		planByAddr := make(map[string]*LoadBalancerPoolOriginsModel, len(*plan.Origins))
		for _, po := range *plan.Origins {
			planByAddr[po.Address.ValueString()] = po
		}
		reordered := make([]*LoadBalancerPoolOriginsModel, 0, len(stateOrigins))
		for _, so := range stateOrigins {
			cp := *planByAddr[so.Address.ValueString()] // shallow copy; Header stays a shared alias, never mutated here
			useOriginStateForUnknown(&cp, so)
			reordered = append(reordered, &cp)
		}
		plan.Origins = &reordered
	}

	// Carry the computed fields the API omits from state when unknown. Stock
	// UseStateForUnknown can't: it skips a null prior state, which is exactly
	// these fields on a pool that never set them.
	if plan.DisabledAt.IsUnknown() {
		plan.DisabledAt = state.DisabledAt
	}
	if plan.Networks.IsUnknown() {
		plan.Networks = state.Networks
	}
	if plan.NotificationFilter.IsUnknown() {
		plan.NotificationFilter = state.NotificationFilter
	}
	if plan.OriginSteering.IsUnknown() {
		plan.OriginSteering = state.OriginSteering
	}
	if plan.LoadShedding.IsUnknown() {
		plan.LoadShedding = state.LoadShedding
	}

	// Tentatively treat the run as a no-op by pinning modified_on, then compare
	// the planned value to state. tftypes equality is reliable where a Go-level
	// comparison of framework types is not.
	plan.ModifiedOn = state.ModifiedOn
	resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !resp.Plan.Raw.Equal(req.State.Raw) {
		plan.Origins = origOrigins
		plan.ModifiedOn = origModifiedOn
		resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
	}
}

// useOriginStateForUnknown carries the computed origin fields the API omits from
// state into the planned origin when the planned value is unknown.
func useOriginStateForUnknown(plan, state *LoadBalancerPoolOriginsModel) {
	if plan.DisabledAt.IsUnknown() {
		plan.DisabledAt = state.DisabledAt
	}
	if plan.Port.IsUnknown() {
		plan.Port = state.Port
	}
	if plan.Enabled.IsUnknown() {
		plan.Enabled = state.Enabled
	}
	if plan.Weight.IsUnknown() {
		plan.Weight = state.Weight
	}
	if plan.FlattenCNAME.IsUnknown() {
		plan.FlattenCNAME = state.FlattenCNAME
	}
}

// originAddressSetsEqual reports whether a and b hold the same set of non-empty
// addresses with no duplicates, i.e. address is a usable identity key for matching.
func originAddressSetsEqual(a, b []*LoadBalancerPoolOriginsModel) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[string]int, len(a))
	for _, o := range a {
		if o == nil || o.Address.IsNull() || o.Address.IsUnknown() {
			return false
		}
		addr := o.Address.ValueString()
		counts[addr]++
		if counts[addr] > 1 {
			return false // duplicate address: ambiguous identity
		}
	}
	for _, o := range b {
		if o == nil || o.Address.IsNull() || o.Address.IsUnknown() {
			return false
		}
		if _, ok := counts[o.Address.ValueString()]; !ok {
			return false // address in state not present in plan
		}
		counts[o.Address.ValueString()]--
	}
	return true
}
