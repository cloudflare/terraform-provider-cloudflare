// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_policy_set

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSPolicySetResultEnvelope struct {
	Result ZeroTrustCWSPolicySetModel `json:"result"`
}

type ZeroTrustCWSPolicySetModel struct {
	ID                 types.String                           `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                           `tfsdk:"account_id" path:"account_id,required"`
	Name               types.String                           `tfsdk:"name" json:"name,required"`
	Policies           *[]*ZeroTrustCWSPolicySetPoliciesModel `tfsdk:"policies" json:"policies,required"`
	Description        types.String                           `tfsdk:"description" json:"description,optional"`
	Tags               *map[string]types.String               `tfsdk:"tags" json:"tags,optional"`
	ExcludeWorkloadIDs customfield.Set[types.String]          `tfsdk:"exclude_workload_ids" json:"exclude_workload_ids,computed_optional"`
	IncludeWorkloadIDs customfield.Set[types.String]          `tfsdk:"include_workload_ids" json:"include_workload_ids,computed_optional"`
	CreatedAt          timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt          timetypes.RFC3339                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSPolicySetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSPolicySetModel) MarshalJSONForUpdate(state ZeroTrustCWSPolicySetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustCWSPolicySetPoliciesModel struct {
	PolicyID types.String `tfsdk:"policy_id" json:"policy_id,required"`
	Enforce  types.Bool   `tfsdk:"enforce" json:"enforce,computed_optional"`
}
