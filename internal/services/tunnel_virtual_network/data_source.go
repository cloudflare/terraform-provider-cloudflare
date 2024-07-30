// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_virtual_network

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type TunnelVirtualNetworkDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &TunnelVirtualNetworkDataSource{}

func NewTunnelVirtualNetworkDataSource() datasource.DataSource {
	return &TunnelVirtualNetworkDataSource{}
}

func (d *TunnelVirtualNetworkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnel_virtual_network"
}

func (r *TunnelVirtualNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TunnelVirtualNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TunnelVirtualNetworkDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*TunnelVirtualNetworkDataSourceModel{}
	env := TunnelVirtualNetworkResultListDataSourceEnvelope{items}

	page, err := r.client.ZeroTrust.Networks.VirtualNetworks.List(ctx, zero_trust.NetworkVirtualNetworkListParams{
		AccountID: cloudflare.F(data.Filter.AccountID.ValueString()),
		ID:        cloudflare.F(data.Filter.ID.ValueString()),
		IsDefault: cloudflare.F(data.Filter.IsDefault.ValueBool()),
		IsDeleted: cloudflare.F(data.Filter.IsDeleted.ValueBool()),
		Name:      cloudflare.F(data.Filter.Name.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes := []byte(page.JSON.RawJSON())
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
		return
	}

	if count := len(*items); count != 1 {
		resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
		return
	}
	data = (*items)[0]

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
