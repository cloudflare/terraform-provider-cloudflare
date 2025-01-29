// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_scoring_integration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustRiskScoringIntegrationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustRiskScoringIntegrationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustRiskScoringIntegrationsDataSourceModel struct {
	AccountID types.String                                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustRiskScoringIntegrationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustRiskScoringIntegrationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.RiskScoringIntegrationListParams, diags diag.Diagnostics) {
	params = zero_trust.RiskScoringIntegrationListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustRiskScoringIntegrationsResultDataSourceModel struct {
	ID              types.String      `tfsdk:"id" json:"id,computed"`
	AccountTag      types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	Active          types.Bool        `tfsdk:"active" json:"active,computed"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IntegrationType types.String      `tfsdk:"integration_type" json:"integration_type,computed"`
	ReferenceID     types.String      `tfsdk:"reference_id" json:"reference_id,computed"`
	TenantURL       types.String      `tfsdk:"tenant_url" json:"tenant_url,computed"`
	WellKnownURL    types.String      `tfsdk:"well_known_url" json:"well_known_url,computed"`
}
