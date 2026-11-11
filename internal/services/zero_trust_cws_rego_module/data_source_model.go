// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_rego_module

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSRegoModuleResultDataSourceEnvelope struct {
	Result ZeroTrustCWSRegoModuleDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSRegoModuleDataSourceModel struct {
	ID           types.String                                    `tfsdk:"id" path:"rego_module_id,computed"`
	RegoModuleID types.String                                    `tfsdk:"rego_module_id" path:"rego_module_id,optional"`
	AccountID    types.String                                    `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt    timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name         types.String                                    `tfsdk:"name" json:"name,computed"`
	PackagePath  types.String                                    `tfsdk:"package_path" json:"package_path,computed"`
	RegoSource   types.String                                    `tfsdk:"rego_source" json:"rego_source,computed"`
	UpdatedAt    timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Filter       *ZeroTrustCWSRegoModuleFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustCWSRegoModuleDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSRegoModuleGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSRegoModuleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSRegoModuleDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSRegoModuleListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSRegoModuleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSRegoModuleListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSRegoModuleListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.PackagePath.IsNull() {
		params.PackagePath = cloudflare.F(m.Filter.PackagePath.ValueString())
	}

	return
}

type ZeroTrustCWSRegoModuleFindOneByDataSourceModel struct {
	Direction   types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name        types.String `tfsdk:"name" query:"name,optional"`
	Order       types.String `tfsdk:"order" query:"order,computed_optional"`
	PackagePath types.String `tfsdk:"package_path" query:"package_path,optional"`
}
