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

type ZeroTrustCWSDeploymentResultDataSourceEnvelope struct {
	Result ZeroTrustCWSDeploymentDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSDeploymentDataSourceModel struct {
	ID                     types.String                                    `tfsdk:"id" path:"deployment_id,computed"`
	DeploymentID           types.String                                    `tfsdk:"deployment_id" path:"deployment_id,optional"`
	AccountID              types.String                                    `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt              timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DefaultMode            types.String                                    `tfsdk:"default_mode" json:"default_mode,computed"`
	Description            types.String                                    `tfsdk:"description" json:"description,computed"`
	DomainACLID            types.String                                    `tfsdk:"domain_acl_id" json:"domain_acl_id,computed"`
	Name                   types.String                                    `tfsdk:"name" json:"name,computed"`
	OnErrorAction          types.String                                    `tfsdk:"on_error_action" json:"on_error_action,computed"`
	RegoRaiseError         types.Bool                                      `tfsdk:"rego_raise_error" json:"rego_raise_error,computed"`
	UpdatedAt              timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	BasicAuthCredentialIDs customfield.Set[types.String]                   `tfsdk:"basic_auth_credential_ids" json:"basic_auth_credential_ids,computed"`
	PolicySetIDs           customfield.Set[types.String]                   `tfsdk:"policy_set_ids" json:"policy_set_ids,computed"`
	Tags                   customfield.Map[types.String]                   `tfsdk:"tags" json:"tags,computed"`
	Filter                 *ZeroTrustCWSDeploymentFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustCWSDeploymentDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSDeploymentGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSDeploymentGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSDeploymentDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSDeploymentListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSDeploymentListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSDeploymentListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSDeploymentListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSDeploymentFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
	Order     types.String `tfsdk:"order" query:"order,computed_optional"`
}
