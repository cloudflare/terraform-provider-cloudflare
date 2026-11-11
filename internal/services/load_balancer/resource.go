// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

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
var _ resource.ResourceWithConfigure = (*LoadBalancerResource)(nil)
var _ resource.ResourceWithImportState = (*LoadBalancerResource)(nil)

func NewResource() resource.Resource {
	return &LoadBalancerResource{}
}

// LoadBalancerResource defines the resource implementation.
type LoadBalancerResource struct {
	client *cloudflare.Client
}

// isZeroOrEmptyTime reports whether s is empty or parses to the Go zero time
// (the Load Balancing API's PUT response omits created_on, which unmarshals to
// "0001-01-01T00:00:00Z").
func isZeroOrEmptyTime(s string) bool {
	if s == "" {
		return true
	}
	t, err := time.Parse(time.RFC3339, s)
	return err == nil && t.IsZero()
}

// preserveCreatedOnFromState keeps the prior created_on when the Update response
// omits it (the Load Balancing API's PUT response does not echo created_on, so
// apijson.UnmarshalComputed leaves it at the Go zero time). Without this the
// framework rejects the apply: the plan carried the known prior created_on but
// the result came back "0001-01-01T00:00:00Z".
func preserveCreatedOnFromState(data *LoadBalancerModel, state *LoadBalancerModel) {
	if data == nil || state == nil {
		return
	}

	if data.CreatedOn.IsUnknown() || data.CreatedOn.IsNull() || isZeroOrEmptyTime(data.CreatedOn.ValueString()) {
		if !state.CreatedOn.IsUnknown() && !state.CreatedOn.IsNull() {
			data.CreatedOn = state.CreatedOn
		}
	}
}

// preserveTTLFromState keeps the prior ttl when an API response omits it. ttl
// applies only to unproxied load balancers; for a proxied LB the GET omits ttl
// (so a refresh would null a server-managed value) while the create/update PUT
// echoes the server default. Restoring the prior known value keeps the plan and
// the stored state stable, avoiding inconsistent-result-after-apply and refresh
// churn.
func preserveTTLFromState(data *LoadBalancerModel, state *LoadBalancerModel) {
	if data == nil || state == nil {
		return
	}

	if data.TTL.IsUnknown() || data.TTL.IsNull() {
		if !state.TTL.IsUnknown() && !state.TTL.IsNull() {
			data.TTL = state.TTL
		}
	}
}

func (r *LoadBalancerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancer"
}

func (r *LoadBalancerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LoadBalancerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *LoadBalancerModel

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
	env := LoadBalancerResultEnvelope{*data}
	_, err = r.client.LoadBalancers.New(
		ctx,
		load_balancers.LoadBalancerNewParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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

func (r *LoadBalancerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *LoadBalancerModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *LoadBalancerModel

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
	env := LoadBalancerResultEnvelope{*data}
	_, err = r.client.LoadBalancers.Update(
		ctx,
		data.ID.ValueString(),
		load_balancers.LoadBalancerUpdateParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	preserveCreatedOnFromState(data, state)
	preserveTTLFromState(data, state)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *LoadBalancerModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// These attributes are not RESTful, and are conditionally modified
	// on the server-side based on other attribute values.
	steeringPolicy := data.SteeringPolicy.ValueString()
	rules := data.Rules
	// The GET omits ttl for proxied LBs; keep the prior value so a refresh does
	// not null a server-managed default (see preserveTTLFromState).
	priorState := *data

	res := new(http.Response)
	env := LoadBalancerResultEnvelope{*data}
	_, err := r.client.LoadBalancers.Get(
		ctx,
		data.ID.ValueString(),
		load_balancers.LoadBalancerGetParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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

	// Set the non-RESTful attributes back to state value
	responseSteeringPolicy := data.SteeringPolicy.ValueString()
	if steeringPolicy == "" && (responseSteeringPolicy == "geo" || responseSteeringPolicy == "off") {
		data.SteeringPolicy = types.StringValue("")
	}
	data.Rules = rules
	preserveTTLFromState(data, &priorState)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *LoadBalancerModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.LoadBalancers.Delete(
		ctx,
		data.ID.ValueString(),
		load_balancers.LoadBalancerDeleteParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(LoadBalancerModel)

	path_zone_id := ""
	path_load_balancer_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<load_balancer_id>",
		&path_zone_id,
		&path_load_balancer_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.ID = types.StringValue(path_load_balancer_id)

	res := new(http.Response)
	env := LoadBalancerResultEnvelope{*data}
	_, err := r.client.LoadBalancers.Get(
		ctx,
		path_load_balancer_id,
		load_balancers.LoadBalancerGetParams{
			ZoneID: cloudflare.F(path_zone_id),
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
