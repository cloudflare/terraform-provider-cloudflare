// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_network_hostname_route

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustNetworkHostnameRoutesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustNetworkHostnameRoutesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustNetworkHostnameRoutesDataSourceModel struct {
	AccountID types.String                                                                      `tfsdk:"account_id" path:"account_id,required"`
	Comment   types.String                                                                      `tfsdk:"comment" query:"comment,optional"`
	ExistedAt types.String                                                                      `tfsdk:"existed_at" query:"existed_at,optional"`
	Hostname  types.String                                                                      `tfsdk:"hostname" query:"hostname,optional"`
	ID        types.String                                                                      `tfsdk:"id" query:"id,optional"`
	TunnelID  types.String                                                                      `tfsdk:"tunnel_id" query:"tunnel_id,optional"`
	IsDeleted types.Bool                                                                        `tfsdk:"is_deleted" query:"is_deleted,computed_optional"`
	MaxItems  types.Int64                                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustNetworkHostnameRoutesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustNetworkHostnameRoutesDataSourceModel) toListParams(_ context.Context) (params zero_trust.NetworkHostnameRouteListParams, diags diag.Diagnostics) {
	params = zero_trust.NetworkHostnameRouteListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Comment.IsNull() {
		params.Comment = cloudflare.F(m.Comment.ValueString())
	}
	if !m.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(m.ExistedAt.ValueString())
	}
	if !m.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Hostname.ValueString())
	}
	if !m.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.IsDeleted.ValueBool())
	}
	if !m.TunnelID.IsNull() {
		params.TunnelID = cloudflare.F(m.TunnelID.ValueString())
	}

	return
}

type ZeroTrustNetworkHostnameRoutesResultDataSourceModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	Comment    types.String      `tfsdk:"comment" json:"comment,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt  timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Hostname   types.String      `tfsdk:"hostname" json:"hostname,computed"`
	TunnelID   types.String      `tfsdk:"tunnel_id" json:"tunnel_id,computed"`
	TunnelName types.String      `tfsdk:"tunnel_name" json:"tunnel_name,computed"`
}
