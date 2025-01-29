// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredConfigResultDataSourceEnvelope struct {
	Result ZeroTrustTunnelCloudflaredConfigDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredConfigDataSourceModel struct {
	AccountID types.String                                                                    `tfsdk:"account_id" path:"account_id,computed"`
	TunnelID  types.String                                                                    `tfsdk:"tunnel_id" path:"tunnel_id,computed"`
	CreatedAt timetypes.RFC3339                                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Source    types.String                                                                    `tfsdk:"source" json:"source,computed"`
	Version   types.Int64                                                                     `tfsdk:"version" json:"version,computed"`
	Config    customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
}

func (m *ZeroTrustTunnelCloudflaredConfigDataSourceModel) toReadParams(_ context.Context) (params zero_trust.TunnelConfigurationGetParams, diags diag.Diagnostics) {
	params = zero_trust.TunnelConfigurationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustTunnelCloudflaredConfigConfigDataSourceModel struct {
	Ingress       customfield.NestedObjectList[ZeroTrustTunnelCloudflaredConfigConfigIngressDataSourceModel]   `tfsdk:"ingress" json:"ingress,computed"`
	OriginRequest customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestDataSourceModel] `tfsdk:"origin_request" json:"originRequest,computed"`
	WARPRouting   customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingDataSourceModel]   `tfsdk:"warp_routing" json:"warp-routing,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressDataSourceModel struct {
	Hostname      types.String                                                                                        `tfsdk:"hostname" json:"hostname,computed"`
	Service       types.String                                                                                        `tfsdk:"service" json:"service,computed"`
	OriginRequest customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestDataSourceModel] `tfsdk:"origin_request" json:"originRequest,computed"`
	Path          types.String                                                                                        `tfsdk:"path" json:"path,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestDataSourceModel struct {
	Access                 customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessDataSourceModel] `tfsdk:"access" json:"access,computed"`
	CAPool                 types.String                                                                                              `tfsdk:"ca_pool" json:"caPool,computed"`
	ConnectTimeout         types.Int64                                                                                               `tfsdk:"connect_timeout" json:"connectTimeout,computed"`
	DisableChunkedEncoding types.Bool                                                                                                `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding,computed"`
	HTTP2Origin            types.Bool                                                                                                `tfsdk:"http2_origin" json:"http2Origin,computed"`
	HTTPHostHeader         types.String                                                                                              `tfsdk:"http_host_header" json:"httpHostHeader,computed"`
	KeepAliveConnections   types.Int64                                                                                               `tfsdk:"keep_alive_connections" json:"keepAliveConnections,computed"`
	KeepAliveTimeout       types.Int64                                                                                               `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout,computed"`
	NoHappyEyeballs        types.Bool                                                                                                `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs,computed"`
	NoTLSVerify            types.Bool                                                                                                `tfsdk:"no_tls_verify" json:"noTLSVerify,computed"`
	OriginServerName       types.String                                                                                              `tfsdk:"origin_server_name" json:"originServerName,computed"`
	ProxyType              types.String                                                                                              `tfsdk:"proxy_type" json:"proxyType,computed"`
	TCPKeepAlive           types.Int64                                                                                               `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive,computed"`
	TLSTimeout             types.Int64                                                                                               `tfsdk:"tls_timeout" json:"tlsTimeout,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessDataSourceModel struct {
	AUDTag   customfield.List[types.String] `tfsdk:"aud_tag" json:"audTag,computed"`
	TeamName types.String                   `tfsdk:"team_name" json:"teamName,computed"`
	Required types.Bool                     `tfsdk:"required" json:"required,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestDataSourceModel struct {
	Access                 customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessDataSourceModel] `tfsdk:"access" json:"access,computed"`
	CAPool                 types.String                                                                                       `tfsdk:"ca_pool" json:"caPool,computed"`
	ConnectTimeout         types.Int64                                                                                        `tfsdk:"connect_timeout" json:"connectTimeout,computed"`
	DisableChunkedEncoding types.Bool                                                                                         `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding,computed"`
	HTTP2Origin            types.Bool                                                                                         `tfsdk:"http2_origin" json:"http2Origin,computed"`
	HTTPHostHeader         types.String                                                                                       `tfsdk:"http_host_header" json:"httpHostHeader,computed"`
	KeepAliveConnections   types.Int64                                                                                        `tfsdk:"keep_alive_connections" json:"keepAliveConnections,computed"`
	KeepAliveTimeout       types.Int64                                                                                        `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout,computed"`
	NoHappyEyeballs        types.Bool                                                                                         `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs,computed"`
	NoTLSVerify            types.Bool                                                                                         `tfsdk:"no_tls_verify" json:"noTLSVerify,computed"`
	OriginServerName       types.String                                                                                       `tfsdk:"origin_server_name" json:"originServerName,computed"`
	ProxyType              types.String                                                                                       `tfsdk:"proxy_type" json:"proxyType,computed"`
	TCPKeepAlive           types.Int64                                                                                        `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive,computed"`
	TLSTimeout             types.Int64                                                                                        `tfsdk:"tls_timeout" json:"tlsTimeout,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessDataSourceModel struct {
	AUDTag   customfield.List[types.String] `tfsdk:"aud_tag" json:"audTag,computed"`
	TeamName types.String                   `tfsdk:"team_name" json:"teamName,computed"`
	Required types.Bool                     `tfsdk:"required" json:"required,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}
