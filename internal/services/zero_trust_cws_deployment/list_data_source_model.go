// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_deployment

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSDeploymentsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSDeploymentsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSDeploymentsDataSourceModel struct {
	AccountID types.String                                                               `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String                                                               `tfsdk:"name" query:"name,optional"`
	Direction types.String                                                               `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String                                                               `tfsdk:"order" query:"order,computed_optional"`
	MaxItems  types.Int64                                                                `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustCWSDeploymentsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSDeploymentsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSDeploymentListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSDeploymentListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSDeploymentListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSDeploymentListParamsOrder(m.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSDeploymentsResultDataSourceModel struct {
	ID                     types.String                  `tfsdk:"id" json:"id,computed"`
	BasicAuthCredentialIDs customfield.Set[types.String] `tfsdk:"basic_auth_credential_ids" json:"basic_auth_credential_ids,computed"`
	CreatedAt              timetypes.RFC3339             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DefaultMode            types.String                  `tfsdk:"default_mode" json:"default_mode,computed"`
	Name                   types.String                  `tfsdk:"name" json:"name,computed"`
	OnErrorAction          types.String                  `tfsdk:"on_error_action" json:"on_error_action,computed"`
	PolicySetIDs           customfield.Set[types.String] `tfsdk:"policy_set_ids" json:"policy_set_ids,computed"`
	RegoRaiseError         types.Bool                    `tfsdk:"rego_raise_error" json:"rego_raise_error,computed"`
	Tags                   customfield.Map[types.String] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt              timetypes.RFC3339             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description            types.String                  `tfsdk:"description" json:"description,computed"`
	DomainACLID            types.String                  `tfsdk:"domain_acl_id" json:"domain_acl_id,computed"`
}
