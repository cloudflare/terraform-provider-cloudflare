// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_basic_auth_credential

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSBasicAuthCredentialsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSBasicAuthCredentialsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSBasicAuthCredentialsDataSourceModel struct {
	AccountID    types.String                                                                        `tfsdk:"account_id" path:"account_id,required"`
	DeploymentID types.String                                                                        `tfsdk:"deployment_id" query:"deployment_id,optional"`
	Username     types.String                                                                        `tfsdk:"username" query:"username,optional"`
	Direction    types.String                                                                        `tfsdk:"direction" query:"direction,computed_optional"`
	Order        types.String                                                                        `tfsdk:"order" query:"order,computed_optional"`
	MaxItems     types.Int64                                                                         `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[ZeroTrustCWSBasicAuthCredentialsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSBasicAuthCredentialsDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSBasicAuthCredentialListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSBasicAuthCredentialListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.DeploymentID.IsNull() {
		params.DeploymentID = cloudflare.F(m.DeploymentID.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSBasicAuthCredentialListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSBasicAuthCredentialListParamsOrder(m.Order.ValueString()))
	}
	if !m.Username.IsNull() {
		params.Username = cloudflare.F(m.Username.ValueString())
	}

	return
}

type ZeroTrustCWSBasicAuthCredentialsResultDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Username  types.String      `tfsdk:"username" json:"username,computed"`
}
