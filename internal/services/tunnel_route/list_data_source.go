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

type TunnelRoutesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &TunnelRoutesDataSource{}

func NewTunnelRoutesDataSource() datasource.DataSource {
	return &TunnelRoutesDataSource{}
}

func (d *TunnelRoutesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnel_routes"
}

func (d *TunnelRoutesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TunnelRoutesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TunnelRoutesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataExistedAt, errs := data.ExistedAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*TunnelRoutesResultDataSourceModel{}
	env := TunnelRoutesResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*TunnelRoutesResultDataSourceModel{}

	page, err := d.client.ZeroTrust.Networks.Routes.List(ctx, zero_trust.NetworkRouteListParams{
		AccountID:        cloudflare.F(data.AccountID.ValueString()),
		Comment:          cloudflare.F(data.Comment.ValueString()),
		ExistedAt:        cloudflare.F(dataExistedAt),
		IsDeleted:        cloudflare.F(data.IsDeleted.ValueBool()),
		NetworkSubset:    cloudflare.F(data.NetworkSubset.ValueString()),
		NetworkSuperset:  cloudflare.F(data.NetworkSuperset.ValueString()),
		Page:             cloudflare.F(data.Page.ValueFloat64()),
		PerPage:          cloudflare.F(data.PerPage.ValueFloat64()),
		RouteID:          cloudflare.F(data.RouteID.ValueString()),
		TunTypes:         cloudflare.F(data.TunTypes.ValueString()),
		TunnelID:         cloudflare.F(data.TunnelID.ValueString()),
		VirtualNetworkID: cloudflare.F(data.VirtualNetworkID.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Result) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, *items...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:maxItems]
	data.Result = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
