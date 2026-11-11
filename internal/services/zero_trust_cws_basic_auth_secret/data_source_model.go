// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_secret

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSBasicAuthSecretResultDataSourceEnvelope struct {
	Result ZeroTrustCWSBasicAuthSecretDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSBasicAuthSecretDataSourceModel struct {
	ID           types.String                                         `tfsdk:"id" path:"secret_id,computed"`
	SecretID     types.String                                         `tfsdk:"secret_id" path:"secret_id,optional"`
	AccountID    types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	CredentialID types.String                                         `tfsdk:"credential_id" path:"credential_id,required"`
	CreatedAt    timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt    timetypes.RFC3339                                    `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Filter       *ZeroTrustCWSBasicAuthSecretFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustCWSBasicAuthSecretDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSBasicAuthCredentialSecretGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSBasicAuthCredentialSecretGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSBasicAuthSecretDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSBasicAuthCredentialSecretListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSBasicAuthCredentialSecretListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSBasicAuthCredentialSecretListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSBasicAuthCredentialSecretListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSBasicAuthSecretFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	Order     types.String `tfsdk:"order" query:"order,computed_optional"`
}
