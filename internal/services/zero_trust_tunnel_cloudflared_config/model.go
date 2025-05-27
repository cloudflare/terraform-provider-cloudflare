// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredConfigResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredConfigModel `json:"result"`
}

type ZeroTrustTunnelCloudflaredConfigModel struct {
	ID        types.String                                                          `tfsdk:"id" json:"-,computed"`
	TunnelID  types.String                                                          `tfsdk:"tunnel_id" path:"tunnel_id,required"`
	AccountID types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	Config    customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	CreatedAt timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Source    types.String                                                          `tfsdk:"source" json:"source,computed"`
	Version   types.Int64                                                           `tfsdk:"version" json:"version,computed"`
}

func (m ZeroTrustTunnelCloudflaredConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustTunnelCloudflaredConfigModel) MarshalJSONForUpdate(state ZeroTrustTunnelCloudflaredConfigModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustTunnelCloudflaredConfigConfigModel struct {
	Ingress       *[]*ZeroTrustTunnelCloudflaredConfigConfigIngressModel                           `tfsdk:"ingress" json:"ingress,optional"`
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel                        `tfsdk:"origin_request" json:"originRequest,optional"`
	WARPRouting   customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel] `tfsdk:"warp_routing" json:"warp-routing,computed"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressModel struct {
	Hostname      types.String                                                     `tfsdk:"hostname" json:"hostname,required"`
	Service       types.String                                                     `tfsdk:"service" json:"service,required"`
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel `tfsdk:"origin_request" json:"originRequest,optional"`
	Path          types.String                                                     `tfsdk:"path" json:"path,optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel struct {
	Access                 *ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel `tfsdk:"access" json:"access,optional"`
	CAPool                 types.String                                                           `tfsdk:"ca_pool" json:"caPool,optional"`
	ConnectTimeout         types.Int64                                                            `tfsdk:"connect_timeout" json:"connectTimeout,optional"`
	DisableChunkedEncoding types.Bool                                                             `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding,optional"`
	HTTP2Origin            types.Bool                                                             `tfsdk:"http2_origin" json:"http2Origin,optional"`
	HTTPHostHeader         types.String                                                           `tfsdk:"http_host_header" json:"httpHostHeader,optional"`
	KeepAliveConnections   types.Int64                                                            `tfsdk:"keep_alive_connections" json:"keepAliveConnections,optional"`
	KeepAliveTimeout       types.Int64                                                            `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout,optional"`
	NoHappyEyeballs        types.Bool                                                             `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs,optional"`
	NoTLSVerify            types.Bool                                                             `tfsdk:"no_tls_verify" json:"noTLSVerify,optional"`
	OriginServerName       types.String                                                           `tfsdk:"origin_server_name" json:"originServerName,optional"`
	ProxyType              types.String                                                           `tfsdk:"proxy_type" json:"proxyType,optional"`
	TCPKeepAlive           types.Int64                                                            `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive,optional"`
	TLSTimeout             types.Int64                                                            `tfsdk:"tls_timeout" json:"tlsTimeout,optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag" json:"audTag,required"`
	TeamName types.String    `tfsdk:"team_name" json:"teamName,required"`
	Required types.Bool      `tfsdk:"required" json:"required,optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel struct {
	Access                 *ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel `tfsdk:"access" json:"access,optional"`
	CAPool                 types.String                                                    `tfsdk:"ca_pool" json:"caPool,optional"`
	ConnectTimeout         types.Int64                                                     `tfsdk:"connect_timeout" json:"connectTimeout,optional"`
	DisableChunkedEncoding types.Bool                                                      `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding,optional"`
	HTTP2Origin            types.Bool                                                      `tfsdk:"http2_origin" json:"http2Origin,optional"`
	HTTPHostHeader         types.String                                                    `tfsdk:"http_host_header" json:"httpHostHeader,optional"`
	KeepAliveConnections   types.Int64                                                     `tfsdk:"keep_alive_connections" json:"keepAliveConnections,optional"`
	KeepAliveTimeout       types.Int64                                                     `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout,optional"`
	NoHappyEyeballs        types.Bool                                                      `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs,optional"`
	NoTLSVerify            types.Bool                                                      `tfsdk:"no_tls_verify" json:"noTLSVerify,optional"`
	OriginServerName       types.String                                                    `tfsdk:"origin_server_name" json:"originServerName,optional"`
	ProxyType              types.String                                                    `tfsdk:"proxy_type" json:"proxyType,optional"`
	TCPKeepAlive           types.Int64                                                     `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive,optional"`
	TLSTimeout             types.Int64                                                     `tfsdk:"tls_timeout" json:"tlsTimeout,optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag" json:"audTag,required"`
	TeamName types.String    `tfsdk:"team_name" json:"teamName,required"`
	Required types.Bool      `tfsdk:"required" json:"required,optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}
