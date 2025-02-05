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

type ZeroTrustTunnelCloudflaredResultDataSourceEnvelope struct {
	Result ZeroTrustTunnelCloudflaredDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredDataSourceModel struct {
	ID              types.String                                                                       `tfsdk:"id" json:"-,computed"`
	TunnelID        types.String                                                                       `tfsdk:"tunnel_id" path:"tunnel_id,optional"`
	AccountID       types.String                                                                       `tfsdk:"account_id" path:"account_id,required"`
	AccountTag      types.String                                                                       `tfsdk:"account_tag" json:"account_tag,computed"`
	ConnsActiveAt   timetypes.RFC3339                                                                  `tfsdk:"conns_active_at" json:"conns_active_at,computed" format:"date-time"`
	ConnsInactiveAt timetypes.RFC3339                                                                  `tfsdk:"conns_inactive_at" json:"conns_inactive_at,computed" format:"date-time"`
	CreatedAt       timetypes.RFC3339                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DeletedAt       timetypes.RFC3339                                                                  `tfsdk:"deleted_at" json:"deleted_at,computed" format:"date-time"`
	Name            types.String                                                                       `tfsdk:"name" json:"name,computed"`
	RemoteConfig    types.Bool                                                                         `tfsdk:"remote_config" json:"remote_config,computed"`
	Status          types.String                                                                       `tfsdk:"status" json:"status,computed"`
	TunType         types.String                                                                       `tfsdk:"tun_type" json:"tun_type,computed"`
	Connections     customfield.NestedObjectList[ZeroTrustTunnelCloudflaredConnectionsDataSourceModel] `tfsdk:"connections" json:"connections,computed"`
	Metadata        jsontypes.Normalized                                                               `tfsdk:"metadata" json:"metadata,computed"`
	Filter          *ZeroTrustTunnelCloudflaredFindOneByDataSourceModel                                `tfsdk:"filter"`
}

func (m *ZeroTrustTunnelCloudflaredDataSourceModel) toReadParams(_ context.Context) (params zero_trust.TunnelGetParams, diags diag.Diagnostics) {
	params = zero_trust.TunnelGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustTunnelCloudflaredDataSourceModel) toListParams(_ context.Context) (params zero_trust.TunnelListParams, diags diag.Diagnostics) {
	mFilterExistedAt, errs := m.Filter.ExistedAt.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterWasActiveAt, errs := m.Filter.WasActiveAt.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterWasInactiveAt, errs := m.Filter.WasInactiveAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.TunnelListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.ExcludePrefix.IsNull() {
		params.ExcludePrefix = cloudflare.F(m.Filter.ExcludePrefix.ValueString())
	}
	if !m.Filter.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(mFilterExistedAt)
	}
	if !m.Filter.IncludePrefix.IsNull() {
		params.IncludePrefix = cloudflare.F(m.Filter.IncludePrefix.ValueString())
	}
	if !m.Filter.IsDeleted.IsNull() {
		params.IsDeleted = cloudflare.F(m.Filter.IsDeleted.ValueBool())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(zero_trust.TunnelListParamsStatus(m.Filter.Status.ValueString()))
	}
	if !m.Filter.UUID.IsNull() {
		params.UUID = cloudflare.F(m.Filter.UUID.ValueString())
	}
	if !m.Filter.WasActiveAt.IsNull() {
		params.WasActiveAt = cloudflare.F(mFilterWasActiveAt)
	}
	if !m.Filter.WasInactiveAt.IsNull() {
		params.WasInactiveAt = cloudflare.F(mFilterWasInactiveAt)
	}

	return
}

type ZeroTrustTunnelCloudflaredConnectionsDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	ClientID           types.String      `tfsdk:"client_id" json:"client_id,computed"`
	ClientVersion      types.String      `tfsdk:"client_version" json:"client_version,computed"`
	ColoName           types.String      `tfsdk:"colo_name" json:"colo_name,computed"`
	IsPendingReconnect types.Bool        `tfsdk:"is_pending_reconnect" json:"is_pending_reconnect,computed"`
	OpenedAt           timetypes.RFC3339 `tfsdk:"opened_at" json:"opened_at,computed" format:"date-time"`
	OriginIP           types.String      `tfsdk:"origin_ip" json:"origin_ip,computed"`
	UUID               types.String      `tfsdk:"uuid" json:"uuid,computed"`
}

type ZeroTrustTunnelCloudflaredFindOneByDataSourceModel struct {
	ExcludePrefix types.String      `tfsdk:"exclude_prefix" query:"exclude_prefix,optional"`
	ExistedAt     timetypes.RFC3339 `tfsdk:"existed_at" query:"existed_at,optional" format:"date-time"`
	IncludePrefix types.String      `tfsdk:"include_prefix" query:"include_prefix,optional"`
	IsDeleted     types.Bool        `tfsdk:"is_deleted" query:"is_deleted,optional"`
	Name          types.String      `tfsdk:"name" query:"name,optional"`
	Status        types.String      `tfsdk:"status" query:"status,optional"`
	UUID          types.String      `tfsdk:"uuid" query:"uuid,optional"`
	WasActiveAt   timetypes.RFC3339 `tfsdk:"was_active_at" query:"was_active_at,optional" format:"date-time"`
	WasInactiveAt timetypes.RFC3339 `tfsdk:"was_inactive_at" query:"was_inactive_at,optional" format:"date-time"`
}
