// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustTunnelCloudflaredsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredsDataSourceModel struct {
	AccountID     types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	ExcludePrefix types.String                                                                   `tfsdk:"exclude_prefix" query:"exclude_prefix,optional"`
	ExistedAt     types.String                                                                   `tfsdk:"existed_at" query:"existed_at,optional"`
	IncludePrefix types.String                                                                   `tfsdk:"include_prefix" query:"include_prefix,optional"`
	IsDeleted     types.Bool                                                                     `tfsdk:"is_deleted" query:"is_deleted,optional"`
	Name          types.String                                                                   `tfsdk:"name" query:"name,optional"`
	Status        types.String                                                                   `tfsdk:"status" query:"status,optional"`
	UUID          types.String                                                                   `tfsdk:"uuid" query:"uuid,optional"`
	WasActiveAt   timetypes.RFC3339                                                              `tfsdk:"was_active_at" query:"was_active_at,optional" format:"date-time"`
	WasInactiveAt timetypes.RFC3339                                                              `tfsdk:"was_inactive_at" query:"was_inactive_at,optional" format:"date-time"`
	MaxItems      types.Int64                                                                    `tfsdk:"max_items"`
	Result        customfield.NestedObjectList[ZeroTrustTunnelCloudflaredsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustTunnelCloudflaredsDataSourceModel) toListParams(_ context.Context) (params zero_trust.TunnelCloudflaredListParams, diags diag.Diagnostics) {
	mWasActiveAt, errs := m.WasActiveAt.ValueRFC3339Time()
	diags.Append(errs...)
	mWasInactiveAt, errs := m.WasInactiveAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.TunnelCloudflaredListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ExcludePrefix.IsNull() {
		params.ExcludePrefix = cloudflare.F(m.ExcludePrefix.ValueString())
	}
	if !m.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(m.ExistedAt.ValueString())
	}
	if !m.IncludePrefix.IsNull() {
		params.IncludePrefix = cloudflare.F(m.IncludePrefix.ValueString())
	}
	if !m.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.IsDeleted.ValueBool())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(zero_trust.TunnelCloudflaredListParamsStatus(m.Status.ValueString()))
	}
	if !m.UUID.IsNull() {
		params.UUID = cloudflare.F(m.UUID.ValueString())
	}
	if !m.WasActiveAt.IsNull() {
		params.WasActiveAt = cloudflare.F(mWasActiveAt)
	}
	if !m.WasInactiveAt.IsNull() {
		params.WasInactiveAt = cloudflare.F(mWasInactiveAt)
	}

	return
}

type ZeroTrustTunnelCloudflaredsResultDataSourceModel struct {
	ID              types.String                                                                        `tfsdk:"id" json:"id,computed"`
	AccountTag      types.String                                                                        `tfsdk:"account_tag" json:"account_tag,computed"`
	Connections     customfield.NestedObjectList[ZeroTrustTunnelCloudflaredsConnectionsDataSourceModel] `tfsdk:"connections" json:"connections,computed"`
	ConnsActiveAt   timetypes.RFC3339                                                                   `tfsdk:"conns_active_at" json:"conns_active_at,computed" format:"date-time"`
	ConnsInactiveAt timetypes.RFC3339                                                                   `tfsdk:"conns_inactive_at" json:"conns_inactive_at,computed" format:"date-time"`
	CreatedAt       timetypes.RFC3339                                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt       timetypes.RFC3339                                                                   `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Metadata        jsontypes.Normalized                                                                `tfsdk:"metadata" json:"metadata,computed"`
	Name            types.String                                                                        `tfsdk:"name" json:"name,computed"`
	RemoteConfig    types.Bool                                                                          `tfsdk:"remote_config" json:"remote_config,computed"`
	Status          types.String                                                                        `tfsdk:"status" json:"status,computed"`
	TunType         types.String                                                                        `tfsdk:"tun_type" json:"tun_type,computed"`
}

type ZeroTrustTunnelCloudflaredsConnectionsDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	ClientID           types.String      `tfsdk:"client_id" json:"client_id,computed"`
	ClientVersion      types.String      `tfsdk:"client_version" json:"client_version,computed"`
	ColoName           types.String      `tfsdk:"colo_name" json:"colo_name,computed"`
	IsPendingReconnect types.Bool        `tfsdk:"is_pending_reconnect" json:"is_pending_reconnect,computed"`
	OpenedAt           timetypes.RFC3339 `tfsdk:"opened_at" json:"opened_at,computed" format:"date-time"`
	OriginIP           types.String      `tfsdk:"origin_ip" json:"origin_ip,computed"`
	UUID               types.String      `tfsdk:"uuid" json:"uuid,computed"`
}
