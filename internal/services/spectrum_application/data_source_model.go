// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/spectrum"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationResultDataSourceEnvelope struct {
	Result SpectrumApplicationDataSourceModel `json:"result,computed"`
}

type SpectrumApplicationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SpectrumApplicationDataSourceModel] `json:"result,computed"`
}

type SpectrumApplicationDataSourceModel struct {
	AppID            types.String                                                          `tfsdk:"app_id" path:"app_id,optional"`
	ZoneID           types.String                                                          `tfsdk:"zone_id" path:"zone_id,optional"`
	ArgoSmartRouting types.Bool                                                            `tfsdk:"argo_smart_routing" json:"argo_smart_routing,computed"`
	CreatedOn        timetypes.RFC3339                                                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ID               types.String                                                          `tfsdk:"id" json:"id,computed"`
	IPFirewall       types.Bool                                                            `tfsdk:"ip_firewall" json:"ip_firewall,computed"`
	ModifiedOn       timetypes.RFC3339                                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Protocol         types.String                                                          `tfsdk:"protocol" json:"protocol,computed"`
	ProxyProtocol    types.String                                                          `tfsdk:"proxy_protocol" json:"proxy_protocol,computed"`
	TLS              types.String                                                          `tfsdk:"tls" json:"tls,computed"`
	TrafficType      types.String                                                          `tfsdk:"traffic_type" json:"traffic_type,computed"`
	OriginDirect     customfield.List[types.String]                                        `tfsdk:"origin_direct" json:"origin_direct,computed"`
	DNS              customfield.NestedObject[SpectrumApplicationDNSDataSourceModel]       `tfsdk:"dns" json:"dns,computed"`
	EdgeIPs          customfield.NestedObject[SpectrumApplicationEdgeIPsDataSourceModel]   `tfsdk:"edge_ips" json:"edge_ips,computed"`
	OriginDNS        customfield.NestedObject[SpectrumApplicationOriginDNSDataSourceModel] `tfsdk:"origin_dns" json:"origin_dns,computed"`
	OriginPort       types.Dynamic                                                         `tfsdk:"origin_port" json:"origin_port,computed"`
	Filter           *SpectrumApplicationFindOneByDataSourceModel                          `tfsdk:"filter"`
}

func (m *SpectrumApplicationDataSourceModel) toReadParams(_ context.Context) (params spectrum.AppGetParams, diags diag.Diagnostics) {
	params = spectrum.AppGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *SpectrumApplicationDataSourceModel) toListParams(_ context.Context) (params spectrum.AppListParams, diags diag.Diagnostics) {
	params = spectrum.AppListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(spectrum.AppListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(spectrum.AppListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type SpectrumApplicationDNSDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type SpectrumApplicationEdgeIPsDataSourceModel struct {
	Connectivity types.String                   `tfsdk:"connectivity" json:"connectivity,computed"`
	Type         types.String                   `tfsdk:"type" json:"type,computed"`
	IPs          customfield.List[types.String] `tfsdk:"ips" json:"ips,computed"`
}

type SpectrumApplicationOriginDNSDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	TTL  types.Int64  `tfsdk:"ttl" json:"ttl,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type SpectrumApplicationFindOneByDataSourceModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String `tfsdk:"order" query:"order,computed_optional"`
}
