package zero_trust_connectivity_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &ConnectivitySettingsResource{}
	_ resource.ResourceWithImportState = &ConnectivitySettingsResource{}
)

func NewResource() resource.Resource {
	return &ConnectivitySettingsResource{}
}

type ConnectivitySettingsResource struct {
	client *muxclient.Client
}

func (r *ConnectivitySettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_connectivity_settings"
}

func (r *ConnectivitySettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*muxclient.Client)
}

func (r *ConnectivitySettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConnectivitySettingsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	accountID := plan.AccountID.ValueString()
	_, err := r.client.V2.ZeroTrust.ConnectivitySettings.Edit(ctx, zero_trust.ConnectivitySettingEditParams{
		AccountID:          cloudflare.F(accountID),
		IcmpProxyEnabled:   cloudflare.F(plan.IcmpProxyEnabled.ValueBool()),
		OfframpWARPEnabled: cloudflare.F(plan.OfframpWARPEnabled.ValueBool()),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating Zero Trust Connectivity Settings", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ConnectivitySettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConnectivitySettingsModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := state.AccountID.ValueString()
	params := zero_trust.ConnectivitySettingGetParams{
		AccountID: cloudflare.F(accountID),
	}
	settings, err := r.client.V2.ZeroTrust.ConnectivitySettings.Get(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Zero Trust Connectivity Settings", err.Error())
		return
	}

	state.IcmpProxyEnabled = types.BoolValue(settings.IcmpProxyEnabled)
	state.OfframpWARPEnabled = types.BoolValue(settings.OfframpWARPEnabled)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ConnectivitySettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConnectivitySettingsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	accountID := plan.AccountID.ValueString()
	_, err := r.client.V2.ZeroTrust.ConnectivitySettings.Edit(ctx, zero_trust.ConnectivitySettingEditParams{
		AccountID:          cloudflare.F(accountID),
		IcmpProxyEnabled:   cloudflare.F(plan.IcmpProxyEnabled.ValueBool()),
		OfframpWARPEnabled: cloudflare.F(plan.OfframpWARPEnabled.ValueBool()),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating Zero Trust Connectivity Settings", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ConnectivitySettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConnectivitySettingsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := state.AccountID.ValueString()
	_, err := r.client.V2.ZeroTrust.ConnectivitySettings.Edit(ctx, zero_trust.ConnectivitySettingEditParams{
		AccountID:          cloudflare.F(accountID),
		IcmpProxyEnabled:   cloudflare.F(false),
		OfframpWARPEnabled: cloudflare.F(false),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error resetting Zero Trust Connectivity Settings", err.Error())
		return
	}
}

func (r *ConnectivitySettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	accountID := req.ID
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), accountID)...)
}
