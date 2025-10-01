// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_network_hostname_route

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustNetworkHostnameRouteResultDataSourceEnvelope struct {
	Result ZeroTrustNetworkHostnameRouteDataSourceModel `json:"result,computed"`
}

type ZeroTrustNetworkHostnameRouteDataSourceModel struct {
	ID              types.String                                           `tfsdk:"id" path:"hostname_route_id,computed"`
	HostnameRouteID types.String                                           `tfsdk:"hostname_route_id" path:"hostname_route_id,optional"`
	AccountID       types.String                                           `tfsdk:"account_id" path:"account_id,required"`
	Comment         types.String                                           `tfsdk:"comment" json:"comment,computed"`
	CreatedAt       timetypes.RFC3339                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt       timetypes.RFC3339                                      `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Hostname        types.String                                           `tfsdk:"hostname" json:"hostname,computed"`
	TunnelID        types.String                                           `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName      types.String                                           `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
	Filter          *ZeroTrustNetworkHostnameRouteFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustNetworkHostnameRouteDataSourceModel) toReadParams(_ context.Context) (params zero_trust.NetworkHostnameRouteGetParams, diags diag.Diagnostics) {
	params = zero_trust.NetworkHostnameRouteGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustNetworkHostnameRouteDataSourceModel) toListParams(_ context.Context) (params zero_trust.NetworkHostnameRouteListParams, diags diag.Diagnostics) {
	params = zero_trust.NetworkHostnameRouteListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.ID.IsNull() {
		params.ID = cloudflare.F(m.Filter.ID.ValueString())
	}
	if !m.Filter.Comment.IsNull() {
		params.Comment = cloudflare.F(m.Filter.Comment.ValueString())
	}
	if !m.Filter.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(m.Filter.ExistedAt.ValueString())
	}
	if !m.Filter.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Filter.Hostname.ValueString())
	}
	if !m.Filter.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.Filter.IsDeleted.ValueBool())
	}
	if !m.Filter.TunnelID.IsNull() {
		params.TunnelID = cloudflare.F(m.Filter.TunnelID.ValueString())
	}

	return
}

type ZeroTrustNetworkHostnameRouteFindOneByDataSourceModel struct {
	ID        types.String `tfsdk:"id" query:"id,optional"`
	Comment   types.String `tfsdk:"comment" query:"comment,optional"`
	ExistedAt types.String `tfsdk:"existed_at" query:"existed_at,optional"`
	Hostname  types.String `tfsdk:"hostname" query:"hostname,optional"`
	IsDeleted types.Bool   `tfsdk:"is_deleted" query:"is_deleted,computed_optional"`
	TunnelID  types.String `tfsdk:"tunnel_id" query:"tunnel_id,optional"`
}
