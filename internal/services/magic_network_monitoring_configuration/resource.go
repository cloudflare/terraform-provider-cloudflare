// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_network_monitoring"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*MagicNetworkMonitoringConfigurationResource)(nil)
var _ resource.ResourceWithModifyPlan = (*MagicNetworkMonitoringConfigurationResource)(nil)

func NewResource() resource.Resource {
	return &MagicNetworkMonitoringConfigurationResource{}
}

// MagicNetworkMonitoringConfigurationResource defines the resource implementation.
type MagicNetworkMonitoringConfigurationResource struct {
	client *cloudflare.Client
}

func (r *MagicNetworkMonitoringConfigurationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_magic_network_monitoring_configuration"
}

func (r *MagicNetworkMonitoringConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MagicNetworkMonitoringConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MagicNetworkMonitoringConfigurationModel

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
	env := MagicNetworkMonitoringConfigurationResultEnvelope{*data}
	_, err = r.client.MagicNetworkMonitoring.Configs.New(
		ctx,
		magic_network_monitoring.ConfigNewParams{
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

func (r *MagicNetworkMonitoringConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MagicNetworkMonitoringConfigurationModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *MagicNetworkMonitoringConfigurationModel

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
	env := MagicNetworkMonitoringConfigurationResultEnvelope{*data}
	_, err = r.client.MagicNetworkMonitoring.Configs.Update(
		ctx,
		magic_network_monitoring.ConfigUpdateParams{
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

func (r *MagicNetworkMonitoringConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MagicNetworkMonitoringConfigurationModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := MagicNetworkMonitoringConfigurationResultEnvelope{*data}
	_, err := r.client.MagicNetworkMonitoring.Configs.Get(
		ctx,
		magic_network_monitoring.ConfigGetParams{
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
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MagicNetworkMonitoringConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MagicNetworkMonitoringConfigurationModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.MagicNetworkMonitoring.Configs.Delete(
		ctx,
		magic_network_monitoring.ConfigDeleteParams{
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

func (r *MagicNetworkMonitoringConfigurationResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
