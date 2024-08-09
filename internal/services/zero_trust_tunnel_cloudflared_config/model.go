// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredConfigResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredConfigModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredConfigModel struct {
	ID        types.String                                 `tfsdk:"id" json:"-,computed"`
	TunnelID  types.String                                 `tfsdk:"tunnel_id" path:"tunnel_id"`
	AccountID types.String                                 `tfsdk:"account_id" path:"account_id"`
	Config    *ZeroTrustTunnelCloudflaredConfigConfigModel `tfsdk:"config" json:"config"`
}

type ZeroTrustTunnelCloudflaredConfigConfigModel struct {
	Ingress       *[]*ZeroTrustTunnelCloudflaredConfigConfigIngressModel    `tfsdk:"ingress" json:"ingress"`
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel `tfsdk:"origin_request" json:"originRequest"`
	WARPRouting   *ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel   `tfsdk:"warp_routing" json:"warp-routing"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressModel struct {
	Hostname      types.String                                                     `tfsdk:"hostname" json:"hostname"`
	Service       types.String                                                     `tfsdk:"service" json:"service"`
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel `tfsdk:"origin_request" json:"originRequest"`
	Path          types.String                                                     `tfsdk:"path" json:"path"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel struct {
	Access                 *ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel `tfsdk:"access" json:"access"`
	CAPool                 types.String                                                           `tfsdk:"ca_pool" json:"caPool"`
	ConnectTimeout         types.Int64                                                            `tfsdk:"connect_timeout" json:"connectTimeout"`
	DisableChunkedEncoding types.Bool                                                             `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding"`
	HTTP2Origin            types.Bool                                                             `tfsdk:"http2_origin" json:"http2Origin"`
	HTTPHostHeader         types.String                                                           `tfsdk:"http_host_header" json:"httpHostHeader"`
	KeepAliveConnections   types.Int64                                                            `tfsdk:"keep_alive_connections" json:"keepAliveConnections"`
	KeepAliveTimeout       types.Int64                                                            `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout"`
	NoHappyEyeballs        types.Bool                                                             `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs"`
	NoTLSVerify            types.Bool                                                             `tfsdk:"no_tls_verify" json:"noTLSVerify"`
	OriginServerName       types.String                                                           `tfsdk:"origin_server_name" json:"originServerName"`
	ProxyType              types.String                                                           `tfsdk:"proxy_type" json:"proxyType"`
	TCPKeepAlive           types.Int64                                                            `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive"`
	TLSTimeout             types.Int64                                                            `tfsdk:"tls_timeout" json:"tlsTimeout"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag" json:"audTag"`
	TeamName types.String    `tfsdk:"team_name" json:"teamName"`
	Required types.Bool      `tfsdk:"required" json:"required"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel struct {
	Access                 *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel `tfsdk:"access" json:"access"`
	CAPool                 types.String                                                    `tfsdk:"ca_pool" json:"caPool"`
	ConnectTimeout         types.Int64                                                     `tfsdk:"connect_timeout" json:"connectTimeout"`
	DisableChunkedEncoding types.Bool                                                      `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding"`
	HTTP2Origin            types.Bool                                                      `tfsdk:"http2_origin" json:"http2Origin"`
	HTTPHostHeader         types.String                                                    `tfsdk:"http_host_header" json:"httpHostHeader"`
	KeepAliveConnections   types.Int64                                                     `tfsdk:"keep_alive_connections" json:"keepAliveConnections"`
	KeepAliveTimeout       types.Int64                                                     `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout"`
	NoHappyEyeballs        types.Bool                                                      `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs"`
	NoTLSVerify            types.Bool                                                      `tfsdk:"no_tls_verify" json:"noTLSVerify"`
	OriginServerName       types.String                                                    `tfsdk:"origin_server_name" json:"originServerName"`
	ProxyType              types.String                                                    `tfsdk:"proxy_type" json:"proxyType"`
	TCPKeepAlive           types.Int64                                                     `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive"`
	TLSTimeout             types.Int64                                                     `tfsdk:"tls_timeout" json:"tlsTimeout"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag" json:"audTag"`
	TeamName types.String    `tfsdk:"team_name" json:"teamName"`
	Required types.Bool      `tfsdk:"required" json:"required"`
}

type ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
