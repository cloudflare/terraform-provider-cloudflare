// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelConfigurationResultEnvelope struct {
	Result TunnelConfigurationModel `json:"result,computed"`
}

type TunnelConfigurationModel struct {
	AccountID types.String                    `tfsdk:"account_id" path:"account_id"`
	TunnelID  types.String                    `tfsdk:"tunnel_id" path:"tunnel_id"`
	Config    *TunnelConfigurationConfigModel `tfsdk:"config" json:"config"`
}

type TunnelConfigurationConfigModel struct {
	Ingress       *[]*TunnelConfigurationConfigIngressModel    `tfsdk:"ingress" json:"ingress"`
	OriginRequest *TunnelConfigurationConfigOriginRequestModel `tfsdk:"originrequest" json:"originRequest"`
	WARPRouting   *TunnelConfigurationConfigWARPRoutingModel   `tfsdk:"warp_routing" json:"warp-routing"`
}

type TunnelConfigurationConfigIngressModel struct {
	Hostname      types.String                                        `tfsdk:"hostname" json:"hostname"`
	Service       types.String                                        `tfsdk:"service" json:"service"`
	OriginRequest *TunnelConfigurationConfigIngressOriginRequestModel `tfsdk:"originrequest" json:"originRequest"`
	Path          types.String                                        `tfsdk:"path" json:"path"`
}

type TunnelConfigurationConfigIngressOriginRequestModel struct {
	Access                 *TunnelConfigurationConfigIngressOriginRequestAccessModel `tfsdk:"access" json:"access"`
	CAPool                 types.String                                              `tfsdk:"capool" json:"caPool"`
	ConnectTimeout         types.Int64                                               `tfsdk:"connecttimeout" json:"connectTimeout"`
	DisableChunkedEncoding types.Bool                                                `tfsdk:"disablechunkedencoding" json:"disableChunkedEncoding"`
	HTTP2Origin            types.Bool                                                `tfsdk:"http2origin" json:"http2Origin"`
	HTTPHostHeader         types.String                                              `tfsdk:"httphostheader" json:"httpHostHeader"`
	KeepAliveConnections   types.Int64                                               `tfsdk:"keepaliveconnections" json:"keepAliveConnections"`
	KeepAliveTimeout       types.Int64                                               `tfsdk:"keepalivetimeout" json:"keepAliveTimeout"`
	NoHappyEyeballs        types.Bool                                                `tfsdk:"nohappyeyeballs" json:"noHappyEyeballs"`
	NoTLSVerify            types.Bool                                                `tfsdk:"notlsverify" json:"noTLSVerify"`
	OriginServerName       types.String                                              `tfsdk:"originservername" json:"originServerName"`
	ProxyType              types.String                                              `tfsdk:"proxytype" json:"proxyType"`
	TCPKeepAlive           types.Int64                                               `tfsdk:"tcpkeepalive" json:"tcpKeepAlive"`
	TLSTimeout             types.Int64                                               `tfsdk:"tlstimeout" json:"tlsTimeout"`
}

type TunnelConfigurationConfigIngressOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"audtag" json:"audTag"`
	TeamName types.String    `tfsdk:"teamname" json:"teamName"`
	Required types.Bool      `tfsdk:"required" json:"required"`
}

type TunnelConfigurationConfigOriginRequestModel struct {
	Access                 *TunnelConfigurationConfigOriginRequestAccessModel `tfsdk:"access" json:"access"`
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

type TunnelConfigurationConfigOriginRequestAccessModel struct {
	AUDTag   *[]types.String `tfsdk:"audtag" json:"audTag"`
	TeamName types.String    `tfsdk:"teamname" json:"teamName"`
	Required types.Bool      `tfsdk:"required" json:"required"`
}

type TunnelConfigurationConfigWARPRoutingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
