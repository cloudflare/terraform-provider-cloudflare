// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/moq"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MoqRelayResultDataSourceEnvelope struct {
	Result MoqRelayDataSourceModel `json:"result,computed"`
}

type MoqRelayDataSourceModel struct {
	ID        types.String                                            `tfsdk:"id" path:"relay_id,computed"`
	RelayID   types.String                                            `tfsdk:"relay_id" path:"relay_id,optional"`
	AccountID types.String                                            `tfsdk:"account_id" path:"account_id,required"`
	Created   timetypes.RFC3339                                       `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                                       `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                                            `tfsdk:"name" json:"name,computed"`
	Status    types.String                                            `tfsdk:"status" json:"status,computed"`
	UID       types.String                                            `tfsdk:"uid" json:"uid,computed"`
	Config    customfield.NestedObject[MoqRelayConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Filter    *MoqRelayFindOneByDataSourceModel                       `tfsdk:"filter"`
}

func (m *MoqRelayDataSourceModel) toReadParams(_ context.Context) (params moq.RelayGetParams, diags diag.Diagnostics) {
	params = moq.RelayGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MoqRelayDataSourceModel) toListParams(_ context.Context) (params moq.RelayListParams, diags diag.Diagnostics) {
	mFilterCreatedAfter, errs := m.Filter.CreatedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mFilterCreatedBefore, errs := m.Filter.CreatedBefore.ValueRFC3339Time()
	diags.Append(errs...)

	params = moq.RelayListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Asc.IsNull() {
		params.Asc = cloudflare.F(m.Filter.Asc.ValueBool())
	}
	if !m.Filter.CreatedAfter.IsNull() {
		params.CreatedAfter = cloudflare.F(mFilterCreatedAfter)
	}
	if !m.Filter.CreatedBefore.IsNull() {
		params.CreatedBefore = cloudflare.F(mFilterCreatedBefore)
	}
	if !m.Filter.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.Filter.PerPage.ValueInt64())
	}

	return
}

type MoqRelayConfigDataSourceModel struct {
	LingeringSubscribe customfield.NestedObject[MoqRelayConfigLingeringSubscribeDataSourceModel] `tfsdk:"lingering_subscribe" json:"lingering_subscribe,computed"`
	OriginFallback     customfield.NestedObject[MoqRelayConfigOriginFallbackDataSourceModel]     `tfsdk:"origin_fallback" json:"origin_fallback,computed"`
}

type MoqRelayConfigLingeringSubscribeDataSourceModel struct {
	Enabled      types.Bool  `tfsdk:"enabled" json:"enabled,computed"`
	MaxTimeoutMs types.Int64 `tfsdk:"max_timeout_ms" json:"max_timeout_ms,computed"`
}

type MoqRelayConfigOriginFallbackDataSourceModel struct {
	Enabled types.Bool                                                                       `tfsdk:"enabled" json:"enabled,computed"`
	Origins customfield.NestedObjectList[MoqRelayConfigOriginFallbackOriginsDataSourceModel] `tfsdk:"origins" json:"origins,computed"`
}

type MoqRelayConfigOriginFallbackOriginsDataSourceModel struct {
	URL types.String `tfsdk:"url" json:"url,computed"`
}

type MoqRelayFindOneByDataSourceModel struct {
	Asc           types.Bool        `tfsdk:"asc" query:"asc,computed_optional"`
	CreatedAfter  timetypes.RFC3339 `tfsdk:"created_after" query:"created_after,optional" format:"date-time"`
	CreatedBefore timetypes.RFC3339 `tfsdk:"created_before" query:"created_before,optional" format:"date-time"`
	PerPage       types.Int64       `tfsdk:"per_page" query:"per_page,optional"`
}
