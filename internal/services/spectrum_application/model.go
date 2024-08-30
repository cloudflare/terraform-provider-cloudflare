// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultEnvelope struct {
	Result SpectrumApplicationModel `json:"result"`
}

type SpectrumApplicationModel struct {
	ID               types.String                       `tfsdk:"id" json:"id,computed"`
	ZoneID           types.String                       `tfsdk:"zone_id" path:"zone_id"`
	Protocol         types.String                       `tfsdk:"protocol" json:"protocol"`
	DNS              *SpectrumApplicationDNSModel       `tfsdk:"dns" json:"dns"`
	IPFirewall       types.Bool                         `tfsdk:"ip_firewall" json:"ip_firewall"`
	TLS              types.String                       `tfsdk:"tls" json:"tls"`
	OriginDirect     *[]types.String                    `tfsdk:"origin_direct" json:"origin_direct"`
	EdgeIPs          *SpectrumApplicationEdgeIPsModel   `tfsdk:"edge_ips" json:"edge_ips"`
	OriginDNS        *SpectrumApplicationOriginDNSModel `tfsdk:"origin_dns" json:"origin_dns"`
	OriginPort       types.Dynamic                      `tfsdk:"origin_port" json:"origin_port"`
	ArgoSmartRouting types.Bool                         `tfsdk:"argo_smart_routing" json:"argo_smart_routing,computed_optional"`
	ProxyProtocol    types.String                       `tfsdk:"proxy_protocol" json:"proxy_protocol,computed_optional"`
	TrafficType      types.String                       `tfsdk:"traffic_type" json:"traffic_type,computed_optional"`
}

type SpectrumApplicationDNSModel struct {
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
	Type types.String `tfsdk:"type" json:"type,computed_optional"`
}

type SpectrumApplicationEdgeIPsModel struct {
	Connectivity types.String    `tfsdk:"connectivity" json:"connectivity"`
	Type         types.String    `tfsdk:"type" json:"type,computed_optional"`
	IPs          *[]types.String `tfsdk:"ips" json:"ips"`
}

type SpectrumApplicationOriginDNSModel struct {
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
	TTL  types.Int64  `tfsdk:"ttl" json:"ttl,computed_optional"`
	Type types.String `tfsdk:"type" json:"type,computed_optional"`
}
