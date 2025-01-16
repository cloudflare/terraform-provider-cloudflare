// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessInfrastructureTargetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessInfrastructureTargetsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessInfrastructureTargetsDataSourceModel struct {
	AccountID        types.String                                                                            `tfsdk:"account_id" path:"account_id,required"`
	CreatedAfter     timetypes.RFC3339                                                                       `tfsdk:"created_after" query:"created_after,optional" format:"date-time"`
	Hostname         types.String                                                                            `tfsdk:"hostname" query:"hostname,optional"`
	HostnameContains types.String                                                                            `tfsdk:"hostname_contains" query:"hostname_contains,optional"`
	IPV4             types.String                                                                            `tfsdk:"ip_v4" query:"ip_v4,optional"`
	IPV6             types.String                                                                            `tfsdk:"ip_v6" query:"ip_v6,optional"`
	ModifiedAfter    timetypes.RFC3339                                                                       `tfsdk:"modified_after" query:"modified_after,optional" format:"date-time"`
	VirtualNetworkID types.String                                                                            `tfsdk:"virtual_network_id" query:"virtual_network_id,optional"`
	MaxItems         types.Int64                                                                             `tfsdk:"max_items"`
	Result           customfield.NestedObjectList[ZeroTrustAccessInfrastructureTargetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessInfrastructureTargetsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessInfrastructureTargetListParams, diags diag.Diagnostics) {
	mCreatedAfter, errs := m.CreatedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mModifiedAfter, errs := m.ModifiedAfter.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.AccessInfrastructureTargetListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.CreatedAfter.IsNull() {
		params.CreatedAfter = cloudflare.F(mCreatedAfter)
	}
	if !m.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Hostname.ValueString())
	}
	if !m.HostnameContains.IsNull() {
		params.HostnameContains = cloudflare.F(m.HostnameContains.ValueString())
	}
	if !m.IPV4.IsNull() {
		params.IPV4 = cloudflare.F(m.IPV4.ValueString())
	}
	if !m.IPV6.IsNull() {
		params.IPV6 = cloudflare.F(m.IPV6.ValueString())
	}
	if !m.ModifiedAfter.IsNull() {
		params.ModifiedAfter = cloudflare.F(mModifiedAfter)
	}
	if !m.VirtualNetworkID.IsNull() {
		params.VirtualNetworkID = cloudflare.F(m.VirtualNetworkID.ValueString())
	}

	return
}

type ZeroTrustAccessInfrastructureTargetsResultDataSourceModel struct {
	ID         types.String                                                                    `tfsdk:"id" json:"id,computed"`
	CreatedAt  timetypes.RFC3339                                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Hostname   types.String                                                                    `tfsdk:"hostname" json:"hostname,computed"`
	IP         customfield.NestedObject[ZeroTrustAccessInfrastructureTargetsIPDataSourceModel] `tfsdk:"ip" json:"ip,computed"`
	ModifiedAt timetypes.RFC3339                                                               `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

type ZeroTrustAccessInfrastructureTargetsIPDataSourceModel struct {
	IPV4 customfield.NestedObject[ZeroTrustAccessInfrastructureTargetsIPIPV4DataSourceModel] `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV6 customfield.NestedObject[ZeroTrustAccessInfrastructureTargetsIPIPV6DataSourceModel] `tfsdk:"ipv6" json:"ipv6,computed"`
}

type ZeroTrustAccessInfrastructureTargetsIPIPV4DataSourceModel struct {
	IPAddr           types.String `tfsdk:"ip_addr" json:"ip_addr,computed"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
}

type ZeroTrustAccessInfrastructureTargetsIPIPV6DataSourceModel struct {
	IPAddr           types.String `tfsdk:"ip_addr" json:"ip_addr,computed"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
}
