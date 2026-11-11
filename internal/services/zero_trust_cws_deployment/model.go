// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_deployment

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSDeploymentResultEnvelope struct {
	Result ZeroTrustCWSDeploymentModel `json:"result"`
}

type ZeroTrustCWSDeploymentModel struct {
	ID                     types.String             `tfsdk:"id" json:"id,computed"`
	AccountID              types.String             `tfsdk:"account_id" path:"account_id,required"`
	DefaultMode            types.String             `tfsdk:"default_mode" json:"default_mode,required"`
	Name                   types.String             `tfsdk:"name" json:"name,required"`
	BasicAuthCredentialIDs *[]types.String          `tfsdk:"basic_auth_credential_ids" json:"basic_auth_credential_ids,required"`
	PolicySetIDs           *[]types.String          `tfsdk:"policy_set_ids" json:"policy_set_ids,required"`
	Description            types.String             `tfsdk:"description" json:"description,optional"`
	DomainACLID            types.String             `tfsdk:"domain_acl_id" json:"domain_acl_id,optional"`
	Tags                   *map[string]types.String `tfsdk:"tags" json:"tags,optional"`
	OnErrorAction          types.String             `tfsdk:"on_error_action" json:"on_error_action,computed_optional"`
	RegoRaiseError         types.Bool               `tfsdk:"rego_raise_error" json:"rego_raise_error,computed_optional"`
	CreatedAt              timetypes.RFC3339        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt              timetypes.RFC3339        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustCWSDeploymentModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCWSDeploymentModel) MarshalJSONForUpdate(state ZeroTrustCWSDeploymentModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
