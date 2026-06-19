// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/turnstile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*TurnstileWidgetResource)(nil)
var _ resource.ResourceWithModifyPlan = (*TurnstileWidgetResource)(nil)
var _ resource.ResourceWithImportState = (*TurnstileWidgetResource)(nil)

func NewResource() resource.Resource {
	return &TurnstileWidgetResource{}
}

// TurnstileWidgetResource defines the resource implementation.
type TurnstileWidgetResource struct {
	client *cloudflare.Client
}

func (r *TurnstileWidgetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_turnstile_widget"
}

func (r *TurnstileWidgetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TurnstileWidgetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TurnstileWidgetModel

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
	env := TurnstileWidgetResultEnvelope{*data}
	_, err = r.client.Turnstile.Widgets.New(
		ctx,
		turnstile.WidgetNewParams{
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
	data.ID = data.Sitekey

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *TurnstileWidgetModel

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
	env := TurnstileWidgetResultEnvelope{*data}
	_, err = r.client.Turnstile.Widgets.Update(
		ctx,
		data.Sitekey.ValueString(),
		turnstile.WidgetUpdateParams{
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
	data.ID = data.Sitekey

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := TurnstileWidgetResultEnvelope{*data}
	_, err := r.client.Turnstile.Widgets.Get(
		ctx,
		data.Sitekey.ValueString(),
		turnstile.WidgetGetParams{
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
	data.ID = data.Sitekey

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Turnstile.Widgets.Delete(
		ctx,
		data.Sitekey.ValueString(),
		turnstile.WidgetDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.Sitekey

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(TurnstileWidgetModel)

	path_account_id := ""
	path_sitekey := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<sitekey>",
		&path_account_id,
		&path_sitekey,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.Sitekey = types.StringValue(path_sitekey)

	res := new(http.Response)
	env := TurnstileWidgetResultEnvelope{*data}
	_, err := r.client.Turnstile.Widgets.Get(
		ctx,
		path_sitekey,
		turnstile.WidgetGetParams{
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
	data.ID = data.Sitekey

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// ModifyPlan suppresses no-op churn the Cloudflare API causes on this resource:
// it returns domains sorted (so any other order replans as a diff), and several
// computed fields replan as "(known after apply)".
//
// It reuses the prior domain order when the set is unchanged and carries the
// computed values from state, then checks whether the plan now matches state
// exactly. If it does the run was pure churn and modified_on is pinned too,
// yielding no diff. If it does not there is a real change, so the original domain
// order and modified_on are restored and the change plans normally.
func (r *TurnstileWidgetResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var plan, state *TurnstileWidgetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || plan == nil || state == nil {
		return
	}

	origDomains := plan.Domains
	origModifiedOn := plan.ModifiedOn

	// Reuse the prior domain order when only the order differs.
	if plan.Domains != nil && state.Domains != nil && domainsEqualAsSet(*plan.Domains, *state.Domains) {
		plan.Domains = state.Domains
	}

	// Carry computed values from state when the plan leaves them unknown.
	if plan.Sitekey.IsUnknown() {
		plan.Sitekey = state.Sitekey
	}
	if plan.Secret.IsUnknown() {
		plan.Secret = state.Secret
	}
	if plan.CreatedOn.IsUnknown() {
		plan.CreatedOn = state.CreatedOn
	}
	if plan.BotFightMode.IsUnknown() {
		plan.BotFightMode = state.BotFightMode
	}
	if plan.ClearanceLevel.IsUnknown() {
		plan.ClearanceLevel = state.ClearanceLevel
	}
	if plan.EphemeralID.IsUnknown() {
		plan.EphemeralID = state.EphemeralID
	}
	if plan.Offlabel.IsUnknown() {
		plan.Offlabel = state.Offlabel
	}
	if plan.Region.IsUnknown() {
		plan.Region = state.Region
	}

	// Pin modified_on only on a true no-op; tftypes equality is reliable where a
	// Go-level comparison of framework types is not.
	plan.ModifiedOn = state.ModifiedOn
	resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !resp.Plan.Raw.Equal(req.State.Raw) {
		plan.Domains = origDomains
		plan.ModifiedOn = origModifiedOn
		resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
	}
}

// domainsEqualAsSet reports whether a and b contain the same domains ignoring
// order. It returns false on any null/unknown element so an unstable plan is
// left untouched.
func domainsEqualAsSet(a, b []types.String) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[string]int, len(a))
	for _, s := range a {
		if s.IsNull() || s.IsUnknown() {
			return false
		}
		counts[s.ValueString()]++
	}
	for _, s := range b {
		if s.IsNull() || s.IsUnknown() {
			return false
		}
		counts[s.ValueString()]--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
