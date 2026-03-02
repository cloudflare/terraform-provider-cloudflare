package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceSpectrumApplicationModel represents the legacy resource state from v4.x provider.
// Schema version: 1 (v4 provider schema version)
// Resource type: cloudflare_spectrum_application
//
// Key differences from v5:
// - dns, origin_dns, edge_ips stored as arrays (MaxItems:1 blocks)
// - origin_port_range was a separate block (merged into origin_port in v5)
// - origin_port was a plain integer (now DynamicAttribute in v5)
type SourceSpectrumApplicationModel struct {
	ID               types.String                  `tfsdk:"id"`
	ZoneID           types.String                  `tfsdk:"zone_id"`
	Protocol         types.String                  `tfsdk:"protocol"`
	DNS              []SourceDNSModel              `tfsdk:"dns"`
	OriginDirect     types.List                    `tfsdk:"origin_direct"`
	OriginDNS        []SourceOriginDNSModel        `tfsdk:"origin_dns"`
	OriginPort       types.Int64                   `tfsdk:"origin_port"`
	OriginPortRange  []SourceOriginPortRangeModel  `tfsdk:"origin_port_range"`
	ArgoSmartRouting types.Bool                    `tfsdk:"argo_smart_routing"`
	IPFirewall       types.Bool                    `tfsdk:"ip_firewall"`
	ProxyProtocol    types.String                  `tfsdk:"proxy_protocol"`
	TLS              types.String                  `tfsdk:"tls"`
	TrafficType      types.String                  `tfsdk:"traffic_type"`
	EdgeIPs          []SourceEdgeIPsModel          `tfsdk:"edge_ips"`
	CreatedOn        types.String                  `tfsdk:"created_on"`
	ModifiedOn       types.String                  `tfsdk:"modified_on"`
}

type SourceDNSModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type SourceOriginDNSModel struct {
	Name types.String `tfsdk:"name"`
	TTL  types.Int64  `tfsdk:"ttl"`
	Type types.String `tfsdk:"type"`
}

type SourceEdgeIPsModel struct {
	Connectivity types.String `tfsdk:"connectivity"`
	Type         types.String `tfsdk:"type"`
	IPs          types.List   `tfsdk:"ips"`
}

type SourceOriginPortRangeModel struct {
	Start types.Int64 `tfsdk:"start"`
	End   types.Int64 `tfsdk:"end"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetSpectrumApplicationModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_spectrum_application
type TargetSpectrumApplicationModel struct {
	ID               types.String                                       `tfsdk:"id"`
	ZoneID           types.String                                       `tfsdk:"zone_id"`
	Protocol         types.String                                       `tfsdk:"protocol"`
	DNS              *TargetDNSModel                                    `tfsdk:"dns"`
	OriginDirect     *[]types.String                                    `tfsdk:"origin_direct"`
	OriginDNS        *TargetOriginDNSModel                              `tfsdk:"origin_dns"`
	OriginPort       customfield.NormalizedDynamicValue                 `tfsdk:"origin_port"`
	ArgoSmartRouting types.Bool                                         `tfsdk:"argo_smart_routing"`
	IPFirewall       types.Bool                                         `tfsdk:"ip_firewall"`
	ProxyProtocol    types.String                                       `tfsdk:"proxy_protocol"`
	TLS              types.String                                       `tfsdk:"tls"`
	TrafficType      types.String                                       `tfsdk:"traffic_type"`
	EdgeIPs          customfield.NestedObject[TargetEdgeIPsModel]       `tfsdk:"edge_ips"`
	CreatedOn        timetypes.RFC3339                                  `tfsdk:"created_on"`
	ModifiedOn       timetypes.RFC3339                                  `tfsdk:"modified_on"`
}

type TargetDNSModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type TargetOriginDNSModel struct {
	Name types.String `tfsdk:"name"`
	TTL  types.Int64  `tfsdk:"ttl"`
	Type types.String `tfsdk:"type"`
}

type TargetEdgeIPsModel struct {
	Connectivity types.String    `tfsdk:"connectivity"`
	Type         types.String    `tfsdk:"type"`
	IPs          *[]types.String `tfsdk:"ips"`
}
