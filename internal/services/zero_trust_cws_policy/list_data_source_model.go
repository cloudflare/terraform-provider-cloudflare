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

type ZeroTrustCWSPoliciesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSPoliciesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSPoliciesDataSourceModel struct {
	AccountID   types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String                                                            `tfsdk:"name" query:"name,optional"`
	PackagePath types.String                                                            `tfsdk:"package_path" query:"package_path,optional"`
	PolicySetID types.String                                                            `tfsdk:"policy_set_id" query:"policy_set_id,optional"`
	Direction   types.String                                                            `tfsdk:"direction" query:"direction,computed_optional"`
	Order       types.String                                                            `tfsdk:"order" query:"order,computed_optional"`
	MaxItems    types.Int64                                                             `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[ZeroTrustCWSPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSPoliciesDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSPolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSPolicyListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSPolicyListParamsOrder(m.Order.ValueString()))
	}
	if !m.PackagePath.IsNull() {
		params.PackagePath = cloudflare.F(m.PackagePath.ValueString())
	}
	if !m.PolicySetID.IsNull() {
		params.PolicySetID = cloudflare.F(m.PolicySetID.ValueString())
	}

	return
}

type ZeroTrustCWSPoliciesResultDataSourceModel struct {
	ID          types.String                  `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Entrypoint  types.String                  `tfsdk:"entrypoint" json:"entrypoint,computed"`
	Filter      types.String                  `tfsdk:"filter" json:"filter,computed"`
	Name        types.String                  `tfsdk:"name" json:"name,computed"`
	PackagePath types.String                  `tfsdk:"package_path" json:"package_path,computed"`
	RegoSource  types.String                  `tfsdk:"rego_source" json:"rego_source,computed"`
	Tags        customfield.Map[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt   timetypes.RFC3339             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String                  `tfsdk:"description" json:"description,computed"`
}
