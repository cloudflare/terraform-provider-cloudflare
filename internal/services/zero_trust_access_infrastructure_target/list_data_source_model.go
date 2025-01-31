// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
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
	CreatedBefore    timetypes.RFC3339                                                                       `tfsdk:"created_before" query:"created_before,optional" format:"date-time"`
	Direction        types.String                                                                            `tfsdk:"direction" query:"direction,optional"`
	Hostname         types.String                                                                            `tfsdk:"hostname" query:"hostname,optional"`
	HostnameContains types.String                                                                            `tfsdk:"hostname_contains" query:"hostname_contains,optional"`
	IPLike           types.String                                                                            `tfsdk:"ip_like" query:"ip_like,optional"`
	IPV4             types.String                                                                            `tfsdk:"ip_v4" query:"ip_v4,optional"`
	IPV6             types.String                                                                            `tfsdk:"ip_v6" query:"ip_v6,optional"`
	IPV4End          types.String                                                                            `tfsdk:"ipv4_end" query:"ipv4_end,optional"`
	IPV4Start        types.String                                                                            `tfsdk:"ipv4_start" query:"ipv4_start,optional"`
	IPV6End          types.String                                                                            `tfsdk:"ipv6_end" query:"ipv6_end,optional"`
	IPV6Start        types.String                                                                            `tfsdk:"ipv6_start" query:"ipv6_start,optional"`
	ModifiedAfter    timetypes.RFC3339                                                                       `tfsdk:"modified_after" query:"modified_after,optional" format:"date-time"`
	ModifiedBefore   timetypes.RFC3339                                                                       `tfsdk:"modified_before" query:"modified_before,optional" format:"date-time"`
	Order            types.String                                                                            `tfsdk:"order" query:"order,optional"`
	VirtualNetworkID types.String                                                                            `tfsdk:"virtual_network_id" query:"virtual_network_id,optional"`
	IPs              *[]types.String                                                                         `tfsdk:"ips" query:"ips,optional"`
	TargetIDs        *[]types.String                                                                         `tfsdk:"target_ids" query:"target_ids,optional"`
	MaxItems         types.Int64                                                                             `tfsdk:"max_items"`
	Result           customfield.NestedObjectList[ZeroTrustAccessInfrastructureTargetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessInfrastructureTargetsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessInfrastructureTargetListParams, diags diag.Diagnostics) {
	mIPs := []string{}
	for _, item := range *m.IPs {
		mIPs = append(mIPs, item.ValueString())
	}
	mTargetIDs := []string{}
	for _, item := range *m.TargetIDs {
		mTargetIDs = append(mTargetIDs, item.ValueString())
	}
	mCreatedAfter, errs := m.CreatedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mCreatedBefore, errs := m.CreatedBefore.ValueRFC3339Time()
	diags.Append(errs...)
	mModifiedAfter, errs := m.ModifiedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mModifiedBefore, errs := m.ModifiedBefore.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.AccessInfrastructureTargetListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		IPs:       cloudflare.F(mIPs),
		TargetIDs: cloudflare.F(mTargetIDs),
	}

	if !m.CreatedAfter.IsNull() {
		params.CreatedAfter = cloudflare.F(mCreatedAfter)
	}
	if !m.CreatedBefore.IsNull() {
		params.CreatedBefore = cloudflare.F(mCreatedBefore)
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.AccessInfrastructureTargetListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Hostname.ValueString())
	}
	if !m.HostnameContains.IsNull() {
		params.HostnameContains = cloudflare.F(m.HostnameContains.ValueString())
	}
	if !m.IPLike.IsNull() {
		params.IPLike = cloudflare.F(m.IPLike.ValueString())
	}
	if !m.IPV4.IsNull() {
		params.IPV4 = cloudflare.F(m.IPV4.ValueString())
	}
	if !m.IPV6.IsNull() {
		params.IPV6 = cloudflare.F(m.IPV6.ValueString())
	}
	if !m.IPV4End.IsNull() {
		params.IPV4End = cloudflare.F(m.IPV4End.ValueString())
	}
	if !m.IPV4Start.IsNull() {
		params.IPV4Start = cloudflare.F(m.IPV4Start.ValueString())
	}
	if !m.IPV6End.IsNull() {
		params.IPV6End = cloudflare.F(m.IPV6End.ValueString())
	}
	if !m.IPV6Start.IsNull() {
		params.IPV6Start = cloudflare.F(m.IPV6Start.ValueString())
	}
	if !m.ModifiedAfter.IsNull() {
		params.ModifiedAfter = cloudflare.F(mModifiedAfter)
	}
	if !m.ModifiedBefore.IsNull() {
		params.ModifiedBefore = cloudflare.F(mModifiedBefore)
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.AccessInfrastructureTargetListParamsOrder(m.Order.ValueString()))
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
