// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_route

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type TunnelRouteDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &TunnelRouteDataSource{}

func NewTunnelRouteDataSource() datasource.DataSource {
	return &TunnelRouteDataSource{}
}

func (d *TunnelRouteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnel_route"
}

func (d *TunnelRouteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *TunnelRouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TunnelRouteDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataFilterExistedAt, errs := data.Filter.ExistedAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*TunnelRouteDataSourceModel{}
	env := TunnelRouteResultListDataSourceEnvelope{items}

	page, err := d.client.ZeroTrust.Networks.Routes.List(ctx, zero_trust.NetworkRouteListParams{
		AccountID:        cloudflare.F(data.Filter.AccountID.ValueString()),
		Comment:          cloudflare.F(data.Filter.Comment.ValueString()),
		ExistedAt:        cloudflare.F(dataFilterExistedAt),
		IsDeleted:        cloudflare.F(data.Filter.IsDeleted.ValueBool()),
		NetworkSubset:    cloudflare.F(data.Filter.NetworkSubset.ValueString()),
		NetworkSuperset:  cloudflare.F(data.Filter.NetworkSuperset.ValueString()),
		Page:             cloudflare.F(data.Filter.Page.ValueFloat64()),
		PerPage:          cloudflare.F(data.Filter.PerPage.ValueFloat64()),
		RouteID:          cloudflare.F(data.Filter.RouteID.ValueString()),
		TunTypes:         cloudflare.F(data.Filter.TunTypes.ValueString()),
		TunnelID:         cloudflare.F(data.Filter.TunnelID.ValueString()),
		VirtualNetworkID: cloudflare.F(data.Filter.VirtualNetworkID.ValueString()),
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
