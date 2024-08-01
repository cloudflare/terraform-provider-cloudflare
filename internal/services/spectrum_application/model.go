// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultEnvelope struct {
	Result SpectrumApplicationModel `json:"result,computed"`
}

type SpectrumApplicationModel struct {
	ID               types.String                       `tfsdk:"id" json:"id,computed"`
	Zone             types.String                       `tfsdk:"zone" path:"zone"`
	Protocol         types.String                       `tfsdk:"protocol" json:"protocol"`
	DNS              *SpectrumApplicationDNSModel       `tfsdk:"dns" json:"dns"`
	OriginDNS        *SpectrumApplicationOriginDNSModel `tfsdk:"origin_dns" json:"origin_dns"`
	OriginPort       types.Dynamic                      `tfsdk:"origin_port" json:"origin_port"`
	IPFirewall       types.Bool                         `tfsdk:"ip_firewall" json:"ip_firewall"`
	TLS              types.String                       `tfsdk:"tls" json:"tls"`
	ArgoSmartRouting types.Bool                         `tfsdk:"argo_smart_routing" json:"argo_smart_routing"`
	ProxyProtocol    types.String                       `tfsdk:"proxy_protocol" json:"proxy_protocol"`
	TrafficType      types.String                       `tfsdk:"traffic_type" json:"traffic_type"`
	EdgeIPs          *SpectrumApplicationEdgeIPsModel   `tfsdk:"edge_ips" json:"edge_ips"`
	CreatedOn        timetypes.RFC3339                  `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn       timetypes.RFC3339                  `tfsdk:"modified_on" json:"modified_on,computed"`
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
	Connectivity types.String    `tfsdk:"connectivity" json:"connectivity"`
	Type         types.String    `tfsdk:"type" json:"type"`
	IPs          *[]types.String `tfsdk:"ips" json:"ips"`
}
