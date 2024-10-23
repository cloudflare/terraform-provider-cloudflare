// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/spectrum"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumApplicationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SpectrumApplicationsResultDataSourceModel] `json:"result,computed"`
}

type SpectrumApplicationsDataSourceModel struct {
	ZoneID    types.String                                                            `tfsdk:"zone_id" path:"zone_id,required"`
	Direction types.String                                                            `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String                                                            `tfsdk:"order" query:"order,computed_optional"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[SpectrumApplicationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *SpectrumApplicationsDataSourceModel) toListParams(_ context.Context) (params spectrum.AppListParams, diags diag.Diagnostics) {
	params = spectrum.AppListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(spectrum.AppListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(spectrum.AppListParamsOrder(m.Order.ValueString()))
	}

	return
}

type SpectrumApplicationsResultDataSourceModel struct {
	ID               types.String                                                                 `tfsdk:"id" json:"id,computed"`
	CreatedOn        timetypes.RFC3339                                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DNS              customfield.NestedObject[SpectrumApplicationsResultDNSDataSourceModel]       `tfsdk:"dns" json:"dns,computed"`
	IPFirewall       types.Bool                                                                   `tfsdk:"ip_firewall" json:"ip_firewall,computed"`
	ModifiedOn       timetypes.RFC3339                                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Protocol         types.String                                                                 `tfsdk:"protocol" json:"protocol,computed"`
	ProxyProtocol    types.String                                                                 `tfsdk:"proxy_protocol" json:"proxy_protocol,computed"`
	TLS              types.String                                                                 `tfsdk:"tls" json:"tls,computed"`
	TrafficType      types.String                                                                 `tfsdk:"traffic_type" json:"traffic_type,computed"`
	ArgoSmartRouting types.Bool                                                                   `tfsdk:"argo_smart_routing" json:"argo_smart_routing,computed"`
	EdgeIPs          customfield.NestedObject[SpectrumApplicationsResultEdgeIPsDataSourceModel]   `tfsdk:"edge_ips" json:"edge_ips,computed"`
	OriginDirect     customfield.List[types.String]                                               `tfsdk:"origin_direct" json:"origin_direct,computed"`
	OriginDNS        customfield.NestedObject[SpectrumApplicationsResultOriginDNSDataSourceModel] `tfsdk:"origin_dns" json:"origin_dns,computed"`
	OriginPort       types.Dynamic                                                                `tfsdk:"origin_port" json:"origin_port,computed"`
}

type SpectrumApplicationsErrorsDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type SpectrumApplicationsMessagesDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type SpectrumApplicationsResultDNSDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type SpectrumApplicationsResultEdgeIPsDataSourceModel struct {
	Connectivity types.String                   `tfsdk:"connectivity" json:"connectivity,computed"`
	Type         types.String                   `tfsdk:"type" json:"type,computed"`
	IPs          customfield.List[types.String] `tfsdk:"ips" json:"ips,computed"`
}

type SpectrumApplicationsResultOriginDNSDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	TTL  types.Int64  `tfsdk:"ttl" json:"ttl,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type SpectrumApplicationsResultInfoDataSourceModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}
