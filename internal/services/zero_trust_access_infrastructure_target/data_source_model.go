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

type ZeroTrustAccessInfrastructureTargetResultDataSourceEnvelope struct {
	Result ZeroTrustAccessInfrastructureTargetDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessInfrastructureTargetDataSourceModel struct {
	ID         types.String                                                                   `tfsdk:"id" json:"-,computed"`
	TargetID   types.String                                                                   `tfsdk:"target_id" path:"target_id,optional"`
	AccountID  types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt  timetypes.RFC3339                                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Hostname   types.String                                                                   `tfsdk:"hostname" json:"hostname,computed"`
	ModifiedAt timetypes.RFC3339                                                              `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	IP         customfield.NestedObject[ZeroTrustAccessInfrastructureTargetIPDataSourceModel] `tfsdk:"ip" json:"ip,computed"`
	Filter     *ZeroTrustAccessInfrastructureTargetFindOneByDataSourceModel                   `tfsdk:"filter"`
}

func (m *ZeroTrustAccessInfrastructureTargetDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessInfrastructureTargetGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessInfrastructureTargetGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustAccessInfrastructureTargetDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessInfrastructureTargetListParams, diags diag.Diagnostics) {
	mFilterIPs := []string{}
	for _, item := range *m.Filter.IPs {
		mFilterIPs = append(mFilterIPs, item.ValueString())
	}
	mFilterTargetIDs := []string{}
	for _, item := range *m.Filter.TargetIDs {
		mFilterTargetIDs = append(mFilterTargetIDs, item.ValueString())
	}
	mFilterCreatedAfter, errs := m.Filter.CreatedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterCreatedBefore, errs := m.Filter.CreatedBefore.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterModifiedAfter, errs := m.Filter.ModifiedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterModifiedBefore, errs := m.Filter.ModifiedBefore.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.AccessInfrastructureTargetListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		IPs:       cloudflare.F(mFilterIPs),
		TargetIDs: cloudflare.F(mFilterTargetIDs),
	}

	if !m.Filter.CreatedAfter.IsNull() {
		params.CreatedAfter = cloudflare.F(mFilterCreatedAfter)
	}
	if !m.Filter.CreatedBefore.IsNull() {
		params.CreatedBefore = cloudflare.F(mFilterCreatedBefore)
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.AccessInfrastructureTargetListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Filter.Hostname.ValueString())
	}
	if !m.Filter.HostnameContains.IsNull() {
		params.HostnameContains = cloudflare.F(m.Filter.HostnameContains.ValueString())
	}
	if !m.Filter.IPLike.IsNull() {
		params.IPLike = cloudflare.F(m.Filter.IPLike.ValueString())
	}
	if !m.Filter.IPV4.IsNull() {
		params.IPV4 = cloudflare.F(m.Filter.IPV4.ValueString())
	}
	if !m.Filter.IPV6.IsNull() {
		params.IPV6 = cloudflare.F(m.Filter.IPV6.ValueString())
	}
	if !m.Filter.IPV4End.IsNull() {
		params.IPV4End = cloudflare.F(m.Filter.IPV4End.ValueString())
	}
	if !m.Filter.IPV4Start.IsNull() {
		params.IPV4Start = cloudflare.F(m.Filter.IPV4Start.ValueString())
	}
	if !m.Filter.IPV6End.IsNull() {
		params.IPV6End = cloudflare.F(m.Filter.IPV6End.ValueString())
	}
	if !m.Filter.IPV6Start.IsNull() {
		params.IPV6Start = cloudflare.F(m.Filter.IPV6Start.ValueString())
	}
	if !m.Filter.ModifiedAfter.IsNull() {
		params.ModifiedAfter = cloudflare.F(mFilterModifiedAfter)
	}
	if !m.Filter.ModifiedBefore.IsNull() {
		params.ModifiedBefore = cloudflare.F(mFilterModifiedBefore)
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.AccessInfrastructureTargetListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.VirtualNetworkID.IsNull() {
		params.VirtualNetworkID = cloudflare.F(m.Filter.VirtualNetworkID.ValueString())
	}

	return
}

type ZeroTrustAccessInfrastructureTargetIPDataSourceModel struct {
	IPV4 customfield.NestedObject[ZeroTrustAccessInfrastructureTargetIPIPV4DataSourceModel] `tfsdk:"ipv4" json:"ipv4,computed"`
	IPV6 customfield.NestedObject[ZeroTrustAccessInfrastructureTargetIPIPV6DataSourceModel] `tfsdk:"ipv6" json:"ipv6,computed"`
}

type ZeroTrustAccessInfrastructureTargetIPIPV4DataSourceModel struct {
	IPAddr           types.String `tfsdk:"ip_addr" json:"ip_addr,computed"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
}

type ZeroTrustAccessInfrastructureTargetIPIPV6DataSourceModel struct {
	IPAddr           types.String `tfsdk:"ip_addr" json:"ip_addr,computed"`
	VirtualNetworkID types.String `tfsdk:"virtual_network_id" json:"virtual_network_id,computed"`
}

type ZeroTrustAccessInfrastructureTargetFindOneByDataSourceModel struct {
	CreatedAfter     timetypes.RFC3339 `tfsdk:"created_after" query:"created_after,optional" format:"date-time"`
	CreatedBefore    timetypes.RFC3339 `tfsdk:"created_before" query:"created_before,optional" format:"date-time"`
	Direction        types.String      `tfsdk:"direction" query:"direction,optional"`
	Hostname         types.String      `tfsdk:"hostname" query:"hostname,optional"`
	HostnameContains types.String      `tfsdk:"hostname_contains" query:"hostname_contains,optional"`
	IPLike           types.String      `tfsdk:"ip_like" query:"ip_like,optional"`
	IPV4             types.String      `tfsdk:"ip_v4" query:"ip_v4,optional"`
	IPV6             types.String      `tfsdk:"ip_v6" query:"ip_v6,optional"`
	IPs              *[]types.String   `tfsdk:"ips" query:"ips,optional"`
	IPV4End          types.String      `tfsdk:"ipv4_end" query:"ipv4_end,optional"`
	IPV4Start        types.String      `tfsdk:"ipv4_start" query:"ipv4_start,optional"`
	IPV6End          types.String      `tfsdk:"ipv6_end" query:"ipv6_end,optional"`
	IPV6Start        types.String      `tfsdk:"ipv6_start" query:"ipv6_start,optional"`
	ModifiedAfter    timetypes.RFC3339 `tfsdk:"modified_after" query:"modified_after,optional" format:"date-time"`
	ModifiedBefore   timetypes.RFC3339 `tfsdk:"modified_before" query:"modified_before,optional" format:"date-time"`
	Order            types.String      `tfsdk:"order" query:"order,optional"`
	TargetIDs        *[]types.String   `tfsdk:"target_ids" query:"target_ids,optional"`
	VirtualNetworkID types.String      `tfsdk:"virtual_network_id" query:"virtual_network_id,optional"`
}
