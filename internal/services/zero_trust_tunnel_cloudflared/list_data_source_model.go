// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
	ExistedAt     timetypes.RFC3339                                                              `tfsdk:"existed_at" query:"existed_at,optional" format:"date-time"`
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

func (m *ZeroTrustTunnelCloudflaredsDataSourceModel) toListParams(_ context.Context) (params zero_trust.TunnelListParams, diags diag.Diagnostics) {
	mExistedAt, errs := m.ExistedAt.ValueRFC3339Time()
	diags.Append(errs...)
	mWasActiveAt, errs := m.WasActiveAt.ValueRFC3339Time()
	diags.Append(errs...)
	mWasInactiveAt, errs := m.WasInactiveAt.ValueRFC3339Time()
	diags.Append(errs...)

	params = zero_trust.TunnelListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ExcludePrefix.IsNull() {
		params.ExcludePrefix = cloudflare.F(m.ExcludePrefix.ValueString())
	}
	if !m.ExistedAt.IsNull() {
		params.ExistedAt = cloudflare.F(mExistedAt)
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
		params.Status = cloudflare.F(zero_trust.TunnelListParamsStatus(m.Status.ValueString()))
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
}
