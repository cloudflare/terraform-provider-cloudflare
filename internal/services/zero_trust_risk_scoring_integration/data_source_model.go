// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_scoring_integration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustRiskScoringIntegrationResultDataSourceEnvelope struct {
	Result ZeroTrustRiskScoringIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustRiskScoringIntegrationDataSourceModel struct {
	ID              types.String      `tfsdk:"id" path:"integration_id,computed"`
	IntegrationID   types.String      `tfsdk:"integration_id" path:"integration_id,optional"`
	AccountID       types.String      `tfsdk:"account_id" path:"account_id,required"`
	AccountTag      types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	Active          types.Bool        `tfsdk:"active" json:"active,computed"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IntegrationType types.String      `tfsdk:"integration_type" json:"integration_type,computed"`
	ReferenceID     types.String      `tfsdk:"reference_id" json:"reference_id,computed"`
	TenantURL       types.String      `tfsdk:"tenant_url" json:"tenant_url,computed"`
	WellKnownURL    types.String      `tfsdk:"well_known_url" json:"well_known_url,computed"`
}

func (m *ZeroTrustRiskScoringIntegrationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.RiskScoringIntegrationGetParams, diags diag.Diagnostics) {
	params = zero_trust.RiskScoringIntegrationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
