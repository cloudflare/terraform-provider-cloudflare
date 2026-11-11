// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_rego_module

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSRegoModulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSRegoModulesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSRegoModulesDataSourceModel struct {
	AccountID   types.String                                                               `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String                                                               `tfsdk:"name" query:"name,optional"`
	PackagePath types.String                                                               `tfsdk:"package_path" query:"package_path,optional"`
	Direction   types.String                                                               `tfsdk:"direction" query:"direction,computed_optional"`
	Order       types.String                                                               `tfsdk:"order" query:"order,computed_optional"`
	MaxItems    types.Int64                                                                `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[ZeroTrustCWSRegoModulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSRegoModulesDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSRegoModuleListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSRegoModuleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSRegoModuleListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSRegoModuleListParamsOrder(m.Order.ValueString()))
	}
	if !m.PackagePath.IsNull() {
		params.PackagePath = cloudflare.F(m.PackagePath.ValueString())
	}

	return
}

type ZeroTrustCWSRegoModulesResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	PackagePath types.String      `tfsdk:"package_path" json:"package_path,computed"`
	RegoSource  types.String      `tfsdk:"rego_source" json:"rego_source,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
