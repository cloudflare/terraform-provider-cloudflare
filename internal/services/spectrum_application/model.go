// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultEnvelope struct {
	Result SpectrumApplicationModel `json:"result"`
}

type SpectrumApplicationModel struct {
	ID               types.String                                                `tfsdk:"id" json:"id,required"`
	ZoneID           types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	Protocol         types.String                                                `tfsdk:"protocol" json:"protocol,required"`
	DNS              *SpectrumApplicationDNSModel                                `tfsdk:"dns" json:"dns,required"`
	IPFirewall       types.Bool                                                  `tfsdk:"ip_firewall" json:"ip_firewall,optional"`
	TLS              types.String                                                `tfsdk:"tls" json:"tls,optional"`
	OriginDirect     *[]types.String                                             `tfsdk:"origin_direct" json:"origin_direct,optional"`
	OriginPort       types.Dynamic                                               `tfsdk:"origin_port" json:"origin_port,optional"`
	ArgoSmartRouting types.Bool                                                  `tfsdk:"argo_smart_routing" json:"argo_smart_routing,computed_optional"`
	ProxyProtocol    types.String                                                `tfsdk:"proxy_protocol" json:"proxy_protocol,computed_optional"`
	TrafficType      types.String                                                `tfsdk:"traffic_type" json:"traffic_type,computed_optional"`
	EdgeIPs          customfield.NestedObject[SpectrumApplicationEdgeIPsModel]   `tfsdk:"edge_ips" json:"edge_ips,computed_optional"`
	OriginDNS        customfield.NestedObject[SpectrumApplicationOriginDNSModel] `tfsdk:"origin_dns" json:"origin_dns,computed_optional"`
	CreatedOn        timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn       timetypes.RFC3339                                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m SpectrumApplicationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SpectrumApplicationModel) MarshalJSONForUpdate(state SpectrumApplicationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type SpectrumApplicationDNSModel struct {
	Name types.String `tfsdk:"name" json:"name,optional"`
	Type types.String `tfsdk:"type" json:"type,optional"`
}

type SpectrumApplicationEdgeIPsModel struct {
	Connectivity types.String    `tfsdk:"connectivity" json:"connectivity,optional"`
	Type         types.String    `tfsdk:"type" json:"type,optional"`
	IPs          *[]types.String `tfsdk:"ips" json:"ips,optional"`
}

type SpectrumApplicationOriginDNSModel struct {
	Name types.String `tfsdk:"name" json:"name,optional"`
	TTL  types.Int64  `tfsdk:"ttl" json:"ttl,optional"`
	Type types.String `tfsdk:"type" json:"type,optional"`
}
