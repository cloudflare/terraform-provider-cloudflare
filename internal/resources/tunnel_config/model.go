// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelConfigResultEnvelope struct {
	Result TunnelConfigModel `json:"result,computed"`
}

type TunnelConfigModel struct {
	AccountID types.String             `tfsdk:"account_id" path:"account_id"`
	TunnelID  types.String             `tfsdk:"tunnel_id" path:"tunnel_id"`
	Config    *TunnelConfigConfigModel `tfsdk:"config" json:"config"`
}

type TunnelConfigConfigModel struct {
	Ingress       *[]*TunnelConfigConfigIngressModel    `tfsdk:"ingress" json:"ingress"`
	OriginRequest *TunnelConfigConfigOriginRequestModel `tfsdk:"originrequest" json:"originRequest"`
	WARPRouting   *TunnelConfigConfigWARPRoutingModel   `tfsdk:"warp_routing" json:"warp-routing"`
}

type TunnelConfigConfigIngressModel struct {
	Hostname      types.String                                 `tfsdk:"hostname" json:"hostname"`
	Service       types.String                                 `tfsdk:"service" json:"service"`
	OriginRequest *TunnelConfigConfigIngressOriginRequestModel `tfsdk:"originrequest" json:"originRequest"`
	Path          types.String                                 `tfsdk:"path" json:"path"`
}

type TunnelConfigConfigIngressOriginRequestModel struct {
	Access                 *TunnelConfigConfigIngressOriginRequestAccessModel `tfsdk:"access" json:"access"`
	CAPool                 types.String                                       `tfsdk:"capool" json:"caPool"`
	ConnectTimeout         types.Int64                                        `tfsdk:"connecttimeout" json:"connectTimeout"`
	DisableChunkedEncoding types.Bool                                         `tfsdk:"disablechunkedencoding" json:"disableChunkedEncoding"`
	HTTP2Origin            types.Bool                                         `tfsdk:"http2origin" json:"http2Origin"`
	HTTPHostHeader         types.String                                       `tfsdk:"httphostheader" json:"httpHostHeader"`
	KeepAliveConnections   types.Int64                                        `tfsdk:"keepaliveconnections" json:"keepAliveConnections"`
	KeepAliveTimeout       types.Int64                                        `tfsdk:"keepalivetimeout" json:"keepAliveTimeout"`
	NoHappyEyeballs        types.Bool                                         `tfsdk:"nohappyeyeballs" json:"noHappyEyeballs"`
	NoTLSVerify            types.Bool                                         `tfsdk:"notlsverify" json:"noTLSVerify"`
	OriginServerName       types.String                                       `tfsdk:"originservername" json:"originServerName"`
	ProxyType              types.String                                       `tfsdk:"proxytype" json:"proxyType"`
	TCPKeepAlive           types.Int64                                        `tfsdk:"tcpkeepalive" json:"tcpKeepAlive"`
	TLSTimeout             types.Int64                                        `tfsdk:"tlstimeout" json:"tlsTimeout"`
}

type TunnelConfigConfigIngressOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"audtag" json:"audTag"`
	TeamName types.String    `tfsdk:"teamname" json:"teamName"`
	Required types.Bool      `tfsdk:"required" json:"required"`
}

type TunnelConfigConfigOriginRequestModel struct {
	Access                 *TunnelConfigConfigOriginRequestAccessModel `tfsdk:"access" json:"access"`
	CAPool                 types.String                                `tfsdk:"capool" json:"caPool"`
	ConnectTimeout         types.Int64                                 `tfsdk:"connecttimeout" json:"connectTimeout"`
	DisableChunkedEncoding types.Bool                                  `tfsdk:"disablechunkedencoding" json:"disableChunkedEncoding"`
	HTTP2Origin            types.Bool                                  `tfsdk:"http2origin" json:"http2Origin"`
	HTTPHostHeader         types.String                                `tfsdk:"httphostheader" json:"httpHostHeader"`
	KeepAliveConnections   types.Int64                                 `tfsdk:"keepaliveconnections" json:"keepAliveConnections"`
	KeepAliveTimeout       types.Int64                                 `tfsdk:"keepalivetimeout" json:"keepAliveTimeout"`
	NoHappyEyeballs        types.Bool                                  `tfsdk:"nohappyeyeballs" json:"noHappyEyeballs"`
	NoTLSVerify            types.Bool                                  `tfsdk:"notlsverify" json:"noTLSVerify"`
	OriginServerName       types.String                                `tfsdk:"originservername" json:"originServerName"`
	ProxyType              types.String                                `tfsdk:"proxytype" json:"proxyType"`
	TCPKeepAlive           types.Int64                                 `tfsdk:"tcpkeepalive" json:"tcpKeepAlive"`
	TLSTimeout             types.Int64                                 `tfsdk:"tlstimeout" json:"tlsTimeout"`
}

type TunnelConfigConfigOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"audtag" json:"audTag"`
	TeamName types.String    `tfsdk:"teamname" json:"teamName"`
	Required types.Bool      `tfsdk:"required" json:"required"`
}

type TunnelConfigConfigWARPRoutingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
