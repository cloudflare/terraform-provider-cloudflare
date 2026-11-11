// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSPolicyResultDataSourceEnvelope struct {
	Result ZeroTrustCWSPolicyDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSPolicyDataSourceModel struct {
	ID          types.String                                `tfsdk:"id" path:"policy_id,computed"`
	PolicyID    types.String                                `tfsdk:"policy_id" path:"policy_id,optional"`
	AccountID   types.String                                `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt   timetypes.RFC3339                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String                                `tfsdk:"description" json:"description,computed"`
	Entrypoint  types.String                                `tfsdk:"entrypoint" json:"entrypoint,computed"`
	Filter      types.String                                `tfsdk:"filter" json:"filter,computed"`
	Name        types.String                                `tfsdk:"name" json:"name,computed"`
	PackagePath types.String                                `tfsdk:"package_path" json:"package_path,computed"`
	RegoSource  types.String                                `tfsdk:"rego_source" json:"rego_source,computed"`
	UpdatedAt   timetypes.RFC3339                           `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Tags        customfield.Map[types.String]               `tfsdk:"tags" json:"tags,computed"`
	FindOneBy   *ZeroTrustCWSPolicyFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *ZeroTrustCWSPolicyDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSPolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSPolicyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSPolicyDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSPolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.FindOneBy.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSPolicyListParamsDirection(m.FindOneBy.Direction.ValueString()))
	}
	if !m.FindOneBy.Name.IsNull() {
		params.Name = cloudflare.F(m.FindOneBy.Name.ValueString())
	}
	if !m.FindOneBy.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSPolicyListParamsOrder(m.FindOneBy.Order.ValueString()))
	}
	if !m.FindOneBy.PackagePath.IsNull() {
		params.PackagePath = cloudflare.F(m.FindOneBy.PackagePath.ValueString())
	}
	if !m.FindOneBy.PolicySetID.IsNull() {
		params.PolicySetID = cloudflare.F(m.FindOneBy.PolicySetID.ValueString())
	}

	return
}

type ZeroTrustCWSPolicyFindOneByDataSourceModel struct {
	Direction   types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name        types.String `tfsdk:"name" query:"name,optional"`
	Order       types.String `tfsdk:"order" query:"order,computed_optional"`
	PackagePath types.String `tfsdk:"package_path" query:"package_path,optional"`
	PolicySetID types.String `tfsdk:"policy_set_id" query:"policy_set_id,optional"`
}
