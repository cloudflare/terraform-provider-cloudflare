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

type ZeroTrustTunnelCloudflaredsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &ZeroTrustTunnelCloudflaredsDataSource{}

func NewZeroTrustTunnelCloudflaredsDataSource() datasource.DataSource {
	return &ZeroTrustTunnelCloudflaredsDataSource{}
}

func (d *ZeroTrustTunnelCloudflaredsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_tunnel_cloudflareds"
}

func (d *ZeroTrustTunnelCloudflaredsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ZeroTrustTunnelCloudflaredsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ZeroTrustTunnelCloudflaredsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataExistedAt, errs := data.ExistedAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	dataWasActiveAt, errs := data.WasActiveAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	dataWasInactiveAt, errs := data.WasInactiveAt.ValueRFC3339Time()
	resp.Diagnostics.Append(errs...)
	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*ZeroTrustTunnelCloudflaredsResultDataSourceModel{}
	env := ZeroTrustTunnelCloudflaredsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*ZeroTrustTunnelCloudflaredsResultDataSourceModel{}

	page, err := d.client.ZeroTrust.Tunnels.List(ctx, zero_trust.TunnelListParams{
		AccountID:     cloudflare.F(data.AccountID.ValueString()),
		ExcludePrefix: cloudflare.F(data.ExcludePrefix.ValueString()),
		ExistedAt:     cloudflare.F(dataExistedAt),
		IncludePrefix: cloudflare.F(data.IncludePrefix.ValueString()),
		IsDeleted:     cloudflare.F(data.IsDeleted.ValueBool()),
		Name:          cloudflare.F(data.Name.ValueString()),
		Status:        cloudflare.F(zero_trust.TunnelListParamsStatus(data.Status.ValueString())),
		UUID:          cloudflare.F(data.UUID.ValueString()),
		WasActiveAt:   cloudflare.F(dataWasActiveAt),
		WasInactiveAt: cloudflare.F(dataWasInactiveAt),
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
