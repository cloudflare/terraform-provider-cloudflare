// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_secret

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSBasicAuthSecretsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSBasicAuthSecretsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSBasicAuthSecretsDataSourceModel struct {
	AccountID    types.String                                                                    `tfsdk:"account_id" path:"account_id,required"`
	CredentialID types.String                                                                    `tfsdk:"credential_id" path:"credential_id,required"`
	Direction    types.String                                                                    `tfsdk:"direction" query:"direction,computed_optional"`
	Order        types.String                                                                    `tfsdk:"order" query:"order,computed_optional"`
	MaxItems     types.Int64                                                                     `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[ZeroTrustCWSBasicAuthSecretsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSBasicAuthSecretsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSBasicAuthCredentialSecretListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSBasicAuthCredentialSecretListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSBasicAuthCredentialSecretListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSBasicAuthCredentialSecretListParamsOrder(m.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSBasicAuthSecretsResultDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
}
