// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustListResultDataSourceEnvelope struct {
	Result ZeroTrustListDataSourceModel `json:"result,computed"`
}

type ZeroTrustListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustListDataSourceModel] `json:"result,computed"`
}

type ZeroTrustListDataSourceModel struct {
	AccountID   types.String                           `tfsdk:"account_id" path:"account_id"`
	ListID      types.String                           `tfsdk:"list_id" path:"list_id"`
	CreatedAt   timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed"`
	ListCount   types.Float64                          `tfsdk:"list_count" json:"count,computed"`
	UpdatedAt   timetypes.RFC3339                      `tfsdk:"updated_at" json:"updated_at,computed"`
	Description types.String                           `tfsdk:"description" json:"description"`
	ID          types.String                           `tfsdk:"id" json:"id"`
	Name        types.String                           `tfsdk:"name" json:"name"`
	Type        types.String                           `tfsdk:"type" json:"type"`
	Filter      *ZeroTrustListFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustListDataSourceModel) toReadParams() (params zero_trust.GatewayListGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayListGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustListDataSourceModel) toListParams() (params zero_trust.GatewayListListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayListListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Type.IsNull() {
		params.Type = cloudflare.F(zero_trust.GatewayListListParamsType(m.Filter.Type.ValueString()))
	}

	return
}

type ZeroTrustListFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Type      types.String `tfsdk:"type" query:"type"`
}
