// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_integration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbIntegrationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCasbIntegrationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCasbIntegrationsDataSourceModel struct {
	AccountID   types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	Application types.String                                                                 `tfsdk:"application" query:"application,optional"`
	Direction   types.String                                                                 `tfsdk:"direction" query:"direction,optional"`
	DLPEnabled  types.Bool                                                                   `tfsdk:"dlp_enabled" query:"dlp_enabled,optional"`
	Order       types.String                                                                 `tfsdk:"order" query:"order,optional"`
	Page        types.Int64                                                                  `tfsdk:"page" query:"page,optional"`
	PageSize    types.Int64                                                                  `tfsdk:"page_size" query:"page_size,optional"`
	Search      types.String                                                                 `tfsdk:"search" query:"search,optional"`
	Status      types.String                                                                 `tfsdk:"status" query:"status,optional"`
	UseCases    types.String                                                                 `tfsdk:"use_cases" query:"use_cases,optional"`
	MaxItems    types.Int64                                                                  `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[ZeroTrustCasbIntegrationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCasbIntegrationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CasbIntegrationListParams, diags diag.Diagnostics) {
	params = zero_trust.CasbIntegrationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Application.IsNull() {
		params.Application = cloudflare.F(m.Application.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CasbIntegrationListParamsDirection(m.Direction.ValueString()))
	}
	if !m.DLPEnabled.IsNull() {
		params.DLPEnabled = cloudflare.F(m.DLPEnabled.ValueBool())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CasbIntegrationListParamsOrder(m.Order.ValueString()))
	}
	if !m.Page.IsNull() {
		params.Page = cloudflare.F(m.Page.ValueInt64())
	}
	if !m.PageSize.IsNull() {
		params.PageSize = cloudflare.F(m.PageSize.ValueInt64())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(zero_trust.CasbIntegrationListParamsStatus(m.Status.ValueString()))
	}
	if !m.UseCases.IsNull() {
		params.UseCases = cloudflare.F(m.UseCases.ValueString())
	}

	return
}

type ZeroTrustCasbIntegrationsResultDataSourceModel struct {
	ID          types.String                  `tfsdk:"id" json:"id,computed"`
	Application customfield.Map[types.String] `tfsdk:"application" json:"application,computed"`
	Created     timetypes.RFC3339             `tfsdk:"created" json:"created,computed" format:"date-time"`
	IsPaused    types.Bool                    `tfsdk:"is_paused" json:"is_paused,computed"`
	Name        types.String                  `tfsdk:"name" json:"name,computed"`
	Status      types.String                  `tfsdk:"status" json:"status,computed"`
	Updated     timetypes.RFC3339             `tfsdk:"updated" json:"updated,computed" format:"date-time"`
}
