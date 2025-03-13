// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_token

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/option"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ZeroTrustTunnelCloudflaredTokenDataSource struct {
  client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*ZeroTrustTunnelCloudflaredTokenDataSource)(nil)

func NewZeroTrustTunnelCloudflaredTokenDataSource() datasource.DataSource {
  return &ZeroTrustTunnelCloudflaredTokenDataSource{}
}

func (d *ZeroTrustTunnelCloudflaredTokenDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_zero_trust_tunnel_cloudflared_token"
}

func (d *ZeroTrustTunnelCloudflaredTokenDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ZeroTrustTunnelCloudflaredTokenDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  var data *ZeroTrustTunnelCloudflaredTokenDataSourceModel

  resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params, diags := data.toReadParams(ctx)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

	res := new(http.Response)
	env := ZeroTrustTunnelCloudflaredTokenResultDataSourceEnvelope{}
	_, err := d.client.ZeroTrust.Tunnels.Cloudflared.Token.Get(
		ctx,
		data.TunnelID.ValueString(),
		params,
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
	data.Token = env.Result

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
