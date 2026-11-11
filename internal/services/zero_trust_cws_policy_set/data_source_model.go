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

type ZeroTrustCWSPolicySetResultDataSourceEnvelope struct {
	Result ZeroTrustCWSPolicySetDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSPolicySetDataSourceModel struct {
	ID                 types.String                                                              `tfsdk:"id" path:"policy_set_id,computed"`
	PolicySetID        types.String                                                              `tfsdk:"policy_set_id" path:"policy_set_id,optional"`
	AccountID          types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt          timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description        types.String                                                              `tfsdk:"description" json:"description,computed"`
	Name               types.String                                                              `tfsdk:"name" json:"name,computed"`
	UpdatedAt          timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ExcludeWorkloadIDs customfield.Set[types.String]                                             `tfsdk:"exclude_workload_ids" json:"exclude_workload_ids,computed"`
	IncludeWorkloadIDs customfield.Set[types.String]                                             `tfsdk:"include_workload_ids" json:"include_workload_ids,computed"`
	Tags               customfield.Map[types.String]                                             `tfsdk:"tags" json:"tags,computed"`
	Policies           customfield.NestedObjectSet[ZeroTrustCWSPolicySetPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Filter             *ZeroTrustCWSPolicySetFindOneByDataSourceModel                            `tfsdk:"filter"`
}

func (m *ZeroTrustCWSPolicySetDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSPolicySetGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSPolicySetGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSPolicySetDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSPolicySetListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSPolicySetListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.DeploymentID.IsNull() {
		params.DeploymentID = cloudflare.F(m.Filter.DeploymentID.ValueString())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSPolicySetListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSPolicySetListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSPolicySetPoliciesDataSourceModel struct {
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,computed"`
	PolicyID types.String `tfsdk:"policy_id" json:"policy_id,computed"`
}

type ZeroTrustCWSPolicySetFindOneByDataSourceModel struct {
	DeploymentID types.String `tfsdk:"deployment_id" query:"deployment_id,optional"`
	Direction    types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name         types.String `tfsdk:"name" query:"name,optional"`
	Order        types.String `tfsdk:"order" query:"order,computed_optional"`
}
