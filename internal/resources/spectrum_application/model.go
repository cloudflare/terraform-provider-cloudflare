// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultEnvelope struct {
	Result SpectrumApplicationModel `json:"result,computed"`
}

type SpectrumApplicationModel struct {
	ID               types.String                       `tfsdk:"id" json:"id,computed"`
	Zone             types.String                       `tfsdk:"zone" path:"zone"`
	DNS              *SpectrumApplicationDNSModel       `tfsdk:"dns" json:"dns"`
	OriginDNS        *SpectrumApplicationOriginDNSModel `tfsdk:"origin_dns" json:"origin_dns"`
	OriginPort       types.String                       `tfsdk:"origin_port" json:"origin_port"`
	Protocol         types.String                       `tfsdk:"protocol" json:"protocol"`
	ArgoSmartRouting types.Bool                         `tfsdk:"argo_smart_routing" json:"argo_smart_routing"`
	EdgeIPs          *SpectrumApplicationEdgeIPsModel   `tfsdk:"edge_ips" json:"edge_ips"`
	IPFirewall       types.Bool                         `tfsdk:"ip_firewall" json:"ip_firewall"`
	ProxyProtocol    types.String                       `tfsdk:"proxy_protocol" json:"proxy_protocol"`
	TLS              types.String                       `tfsdk:"tls" json:"tls"`
	TrafficType      types.String                       `tfsdk:"traffic_type" json:"traffic_type"`
}

type SpectrumApplicationDNSModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}

type SpectrumApplicationOriginDNSModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	TTL  types.Int64  `tfsdk:"ttl" json:"ttl"`
	Type types.String `tfsdk:"type" json:"type"`
}

type SpectrumApplicationEdgeIPsModel struct {
	Connectivity types.String `tfsdk:"connectivity" json:"connectivity"`
	Type         types.String `tfsdk:"type" json:"type"`
	IPs          types.String `tfsdk:"ips" json:"ips"`
}
