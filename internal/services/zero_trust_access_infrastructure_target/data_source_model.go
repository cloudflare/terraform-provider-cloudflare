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

type ZeroTrustAccessInfrastructureTargetResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessInfrastructureTargetDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessInfrastructureTargetDataSourceModel struct {
	AccountID  types.String                                                                   `tfsdk:"account_id" path:"account_id,optional"`
	TargetID   types.String                                                                   `tfsdk:"target_id" path:"target_id,optional"`
	CreatedAt  timetypes.RFC3339                                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Hostname   types.String                                                                   `tfsdk:"hostname" json:"hostname,computed"`
	ID         types.String                                                                   `tfsdk:"id" json:"id,computed"`
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
	mFilterCreatedAfter, errs := m.Filter.CreatedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterModifiedAfter, errs := m.Filter.ModifiedAfter.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.AccessInfrastructureTargetListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.CreatedAfter.IsNull() {
		params.CreatedAfter = cloudflare.F(mFilterCreatedAfter)
	}
	if !m.Filter.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Filter.Hostname.ValueString())
	}
	if !m.Filter.HostnameContains.IsNull() {
		params.HostnameContains = cloudflare.F(m.Filter.HostnameContains.ValueString())
	}
	if !m.Filter.IPV4.IsNull() {
		params.IPV4 = cloudflare.F(m.Filter.IPV4.ValueString())
	}
	if !m.Filter.IPV6.IsNull() {
		params.IPV6 = cloudflare.F(m.Filter.IPV6.ValueString())
	}
	if !m.Filter.ModifiedAfter.IsNull() {
		params.ModifiedAfter = cloudflare.F(mFilterModifiedAfter)
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
	AccountID        types.String      `tfsdk:"account_id" path:"account_id,required"`
	CreatedAfter     timetypes.RFC3339 `tfsdk:"created_after" query:"created_after,optional" format:"date-time"`
	Hostname         types.String      `tfsdk:"hostname" query:"hostname,optional"`
	HostnameContains types.String      `tfsdk:"hostname_contains" query:"hostname_contains,optional"`
	IPV4             types.String      `tfsdk:"ip_v4" query:"ip_v4,optional"`
	IPV6             types.String      `tfsdk:"ip_v6" query:"ip_v6,optional"`
	ModifiedAfter    timetypes.RFC3339 `tfsdk:"modified_after" query:"modified_after,optional" format:"date-time"`
	VirtualNetworkID types.String      `tfsdk:"virtual_network_id" query:"virtual_network_id,optional"`
}
