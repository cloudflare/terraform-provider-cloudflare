// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_credential

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSBasicAuthCredentialResultDataSourceEnvelope struct {
	Result ZeroTrustCWSBasicAuthCredentialDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSBasicAuthCredentialDataSourceModel struct {
	ID           types.String                                             `tfsdk:"id" path:"credential_id,computed"`
	CredentialID types.String                                             `tfsdk:"credential_id" path:"credential_id,optional"`
	AccountID    types.String                                             `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt    timetypes.RFC3339                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt    timetypes.RFC3339                                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Username     types.String                                             `tfsdk:"username" json:"username,computed"`
	Filter       *ZeroTrustCWSBasicAuthCredentialFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustCWSBasicAuthCredentialDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSBasicAuthCredentialGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSBasicAuthCredentialGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSBasicAuthCredentialDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSBasicAuthCredentialListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSBasicAuthCredentialListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.DeploymentID.IsNull() {
		params.DeploymentID = cloudflare.F(m.Filter.DeploymentID.ValueString())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSBasicAuthCredentialListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSBasicAuthCredentialListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Username.IsNull() {
		params.Username = cloudflare.F(m.Filter.Username.ValueString())
	}

	return
}

type ZeroTrustCWSBasicAuthCredentialFindOneByDataSourceModel struct {
	DeploymentID types.String `tfsdk:"deployment_id" query:"deployment_id,optional"`
	Direction    types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Order        types.String `tfsdk:"order" query:"order,computed_optional"`
	Username     types.String `tfsdk:"username" query:"username,optional"`
}
