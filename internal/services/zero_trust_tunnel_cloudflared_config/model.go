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
	Source    types.String                                                          `tfsdk:"source" json:"source,computed_optional"`
	Version   types.Int64                                                           `tfsdk:"version" json:"version,computed"`
}

func (m ZeroTrustTunnelCloudflaredConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustTunnelCloudflaredConfigModel) MarshalJSONForUpdate(state ZeroTrustTunnelCloudflaredConfigModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustTunnelCloudflaredConfigConfigModel struct {
	Ingress       customfield.NestedObjectList[ZeroTrustTunnelCloudflaredConfigConfigIngressModel]   `tfsdk:"ingress" json:"ingress,computed_optional"`
	OriginRequest customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel] `tfsdk:"origin_request" json:"originRequest,computed_optional"`
	WARPRouting   customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel]   `tfsdk:"warp_routing" json:"warp-routing,computed_optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressModel struct {
	Hostname      types.String                                                                              `tfsdk:"hostname" json:"hostname,optional"`
	Service       types.String                                                                              `tfsdk:"service" json:"service,required"`
	OriginRequest customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel] `tfsdk:"origin_request" json:"originRequest,computed_optional"`
	Path          types.String                                                                              `tfsdk:"path" json:"path,computed_optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel struct {
	Access                 customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel] `tfsdk:"access" json:"access,computed_optional"`
	CAPool                 types.String                                                                                    `tfsdk:"ca_pool" json:"caPool,computed_optional"`
	ConnectTimeout         types.Int64                                                                                     `tfsdk:"connect_timeout" json:"connectTimeout,computed_optional"`
	DisableChunkedEncoding types.Bool                                                                                      `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding,optional"`
	HTTP2Origin            types.Bool                                                                                      `tfsdk:"http2_origin" json:"http2Origin,optional"`
	HTTPHostHeader         types.String                                                                                    `tfsdk:"http_host_header" json:"httpHostHeader,optional"`
	KeepAliveConnections   types.Int64                                                                                     `tfsdk:"keep_alive_connections" json:"keepAliveConnections,computed_optional"`
	KeepAliveTimeout       types.Int64                                                                                     `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout,computed_optional"`
	NoHappyEyeballs        types.Bool                                                                                      `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs,computed_optional"`
	NoTLSVerify            types.Bool                                                                                      `tfsdk:"no_tls_verify" json:"noTLSVerify,computed_optional"`
	OriginServerName       types.String                                                                                    `tfsdk:"origin_server_name" json:"originServerName,computed_optional"`
	ProxyType              types.String                                                                                    `tfsdk:"proxy_type" json:"proxyType,computed_optional"`
	TCPKeepAlive           types.Int64                                                                                     `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive,computed_optional"`
	TLSTimeout             types.Int64                                                                                     `tfsdk:"tls_timeout" json:"tlsTimeout,computed_optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag" json:"audTag,required"`
	TeamName types.String    `tfsdk:"team_name" json:"teamName,computed_optional"`
	Required types.Bool      `tfsdk:"required" json:"required,computed_optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel struct {
	Access                 customfield.NestedObject[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel] `tfsdk:"access" json:"access,computed_optional"`
	CAPool                 types.String                                                                             `tfsdk:"ca_pool" json:"caPool,computed_optional"`
	ConnectTimeout         types.Int64                                                                              `tfsdk:"connect_timeout" json:"connectTimeout,computed_optional"`
	DisableChunkedEncoding types.Bool                                                                               `tfsdk:"disable_chunked_encoding" json:"disableChunkedEncoding,optional"`
	HTTP2Origin            types.Bool                                                                               `tfsdk:"http2_origin" json:"http2Origin,optional"`
	HTTPHostHeader         types.String                                                                             `tfsdk:"http_host_header" json:"httpHostHeader,optional"`
	KeepAliveConnections   types.Int64                                                                              `tfsdk:"keep_alive_connections" json:"keepAliveConnections,computed_optional"`
	KeepAliveTimeout       types.Int64                                                                              `tfsdk:"keep_alive_timeout" json:"keepAliveTimeout,computed_optional"`
	NoHappyEyeballs        types.Bool                                                                               `tfsdk:"no_happy_eyeballs" json:"noHappyEyeballs,computed_optional"`
	NoTLSVerify            types.Bool                                                                               `tfsdk:"no_tls_verify" json:"noTLSVerify,computed_optional"`
	OriginServerName       types.String                                                                             `tfsdk:"origin_server_name" json:"originServerName,computed_optional"`
	ProxyType              types.String                                                                             `tfsdk:"proxy_type" json:"proxyType,computed_optional"`
	TCPKeepAlive           types.Int64                                                                              `tfsdk:"tcp_keep_alive" json:"tcpKeepAlive,computed_optional"`
	TLSTimeout             types.Int64                                                                              `tfsdk:"tls_timeout" json:"tlsTimeout,computed_optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag" json:"audTag,required"`
	TeamName types.String    `tfsdk:"team_name" json:"teamName,computed_optional"`
	Required types.Bool      `tfsdk:"required" json:"required,computed_optional"`
}

type ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}
