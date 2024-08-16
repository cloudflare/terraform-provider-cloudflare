// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ZeroTrustTunnelCloudflaredDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &ZeroTrustTunnelCloudflaredDataSource{}

func NewZeroTrustTunnelCloudflaredDataSource() datasource.DataSource {
	return &ZeroTrustTunnelCloudflaredDataSource{}
}

func (d *ZeroTrustTunnelCloudflaredDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_tunnel_cloudflared"
}

func (d *ZeroTrustTunnelCloudflaredDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ZeroTrustTunnelCloudflaredDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ZeroTrustTunnelCloudflaredDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataFilterExistedAt, errs := data.Filter.ExistedAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	dataFilterWasActiveAt, errs := data.Filter.WasActiveAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	dataFilterWasInactiveAt, errs := data.Filter.WasInactiveAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*ZeroTrustTunnelCloudflaredDataSourceModel{}
	env := ZeroTrustTunnelCloudflaredResultListDataSourceEnvelope{items}

	page, err := d.client.ZeroTrust.Tunnels.List(ctx, zero_trust.TunnelListParams{
		AccountID:     cloudflare.F(data.Filter.AccountID.ValueString()),
		ExcludePrefix: cloudflare.F(data.Filter.ExcludePrefix.ValueString()),
		ExistedAt:     cloudflare.F(dataFilterExistedAt),
		IncludePrefix: cloudflare.F(data.Filter.IncludePrefix.ValueString()),
		IsDeleted:     cloudflare.F(data.Filter.IsDeleted.ValueBool()),
		Name:          cloudflare.F(data.Filter.Name.ValueString()),
		Status:        cloudflare.F(zero_trust.TunnelListParamsStatus(data.Filter.Status.ValueString())),
		TunTypes:      cloudflare.F(data.Filter.TunTypes.ValueString()),
		UUID:          cloudflare.F(data.Filter.UUID.ValueString()),
		WasActiveAt:   cloudflare.F(dataFilterWasActiveAt),
		WasInactiveAt: cloudflare.F(dataFilterWasInactiveAt),
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
