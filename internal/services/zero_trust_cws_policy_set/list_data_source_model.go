// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy_set

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSPolicySetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSPolicySetsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSPolicySetsDataSourceModel struct {
	AccountID    types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	DeploymentID types.String                                                              `tfsdk:"deployment_id" query:"deployment_id,optional"`
	Name         types.String                                                              `tfsdk:"name" query:"name,optional"`
	Direction    types.String                                                              `tfsdk:"direction" query:"direction,computed_optional"`
	Order        types.String                                                              `tfsdk:"order" query:"order,computed_optional"`
	MaxItems     types.Int64                                                               `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[ZeroTrustCWSPolicySetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSPolicySetsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSPolicySetListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSPolicySetListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.DeploymentID.IsNull() {
		params.DeploymentID = cloudflare.F(m.DeploymentID.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSPolicySetListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSPolicySetListParamsOrder(m.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSPolicySetsResultDataSourceModel struct {
	ID                 types.String                                                               `tfsdk:"id" json:"id,computed"`
	CreatedAt          timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExcludeWorkloadIDs customfield.Set[types.String]                                              `tfsdk:"exclude_workload_ids" json:"exclude_workload_ids,computed"`
	IncludeWorkloadIDs customfield.Set[types.String]                                              `tfsdk:"include_workload_ids" json:"include_workload_ids,computed"`
	Name               types.String                                                               `tfsdk:"name" json:"name,computed"`
	Policies           customfield.NestedObjectSet[ZeroTrustCWSPolicySetsPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Tags               customfield.Map[types.String]                                              `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt          timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description        types.String                                                               `tfsdk:"description" json:"description,computed"`
}

type ZeroTrustCWSPolicySetsPoliciesDataSourceModel struct {
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,computed"`
	PolicyID types.String `tfsdk:"policy_id" json:"policy_id,computed"`
}
