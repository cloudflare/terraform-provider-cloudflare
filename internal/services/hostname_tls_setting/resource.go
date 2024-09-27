// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostname_tls_setting

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/hostnames"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*HostnameTLSSettingResource)(nil)
var _ resource.ResourceWithModifyPlan = (*HostnameTLSSettingResource)(nil)
var _ resource.ResourceWithImportState = (*HostnameTLSSettingResource)(nil)

func NewResource() resource.Resource {
	return &HostnameTLSSettingResource{}
}

// HostnameTLSSettingResource defines the resource implementation.
type HostnameTLSSettingResource struct {
	client *cloudflare.Client
}

func (r *HostnameTLSSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hostname_tls_setting"
}

func (r *HostnameTLSSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *HostnameTLSSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *HostnameTLSSettingModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := HostnameTLSSettingResultEnvelope{*data}
	_, err = r.client.Hostnames.Settings.TLS.Update(
		ctx,
		hostnames.SettingTLSUpdateParamsSettingID(data.SettingID.ValueString()),
		data.Hostname.ValueString(),
		hostnames.SettingTLSUpdateParams{
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
	data.ID = data.SettingID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostnameTLSSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *HostnameTLSSettingModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *HostnameTLSSettingModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.MarshalForUpdate(data, state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := HostnameTLSSettingResultEnvelope{*data}
	_, err = r.client.Hostnames.Settings.TLS.Update(
		ctx,
		hostnames.SettingTLSUpdateParamsSettingID(data.SettingID.ValueString()),
		data.SettingID.ValueString(),
		hostnames.SettingTLSUpdateParams{
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
	data.ID = data.SettingID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostnameTLSSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *HostnameTLSSettingModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := HostnameTLSSettingResultEnvelope{*data}
	_, err := r.client.Hostnames.Settings.TLS.Get(
		ctx,
		hostnames.SettingTLSGetParamsSettingID(data.SettingID.ValueString()),
		hostnames.SettingTLSGetParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
	data.ID = data.SettingID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostnameTLSSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *HostnameTLSSettingModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Hostnames.Settings.TLS.Delete(
		ctx,
		hostnames.SettingTLSDeleteParamsSettingID(data.SettingID.ValueString()),
		data.Hostname.ValueString(),
		hostnames.SettingTLSDeleteParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.SettingID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostnameTLSSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *HostnameTLSSettingModel

	path_zone_id := ""
	path_setting_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<setting_id>",
		&path_zone_id,
		&path_setting_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := HostnameTLSSettingResultEnvelope{*data}
	_, err := r.client.Hostnames.Settings.TLS.Get(
		ctx,
		hostnames.SettingTLSGetParamsSettingID(path_setting_id),
		hostnames.SettingTLSGetParams{
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
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.SettingID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostnameTLSSettingResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
