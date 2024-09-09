// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_score_integration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustRiskScoreIntegrationResultDataSourceEnvelope struct {
	Result ZeroTrustRiskScoreIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustRiskScoreIntegrationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustRiskScoreIntegrationDataSourceModel] `json:"result,computed"`
}

type ZeroTrustRiskScoreIntegrationDataSourceModel struct {
	AccountID       types.String                                           `tfsdk:"account_id" path:"account_id,optional"`
	IntegrationID   types.String                                           `tfsdk:"integration_id" path:"integration_id,optional"`
	AccountTag      types.String                                           `tfsdk:"account_tag" json:"account_tag,computed"`
	Active          types.Bool                                             `tfsdk:"active" json:"active,computed"`
	CreatedAt       timetypes.RFC3339                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID              types.String                                           `tfsdk:"id" json:"id,computed"`
	IntegrationType types.String                                           `tfsdk:"integration_type" json:"integration_type,computed"`
	ReferenceID     types.String                                           `tfsdk:"reference_id" json:"reference_id,computed"`
	TenantURL       types.String                                           `tfsdk:"tenant_url" json:"tenant_url,computed"`
	WellKnownURL    types.String                                           `tfsdk:"well_known_url" json:"well_known_url,computed"`
	Filter          *ZeroTrustRiskScoreIntegrationFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustRiskScoreIntegrationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.RiskScoringIntegrationGetParams, diags diag.Diagnostics) {
	params = zero_trust.RiskScoringIntegrationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustRiskScoreIntegrationDataSourceModel) toListParams(_ context.Context) (params zero_trust.RiskScoringIntegrationListParams, diags diag.Diagnostics) {
	params = zero_trust.RiskScoringIntegrationListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustRiskScoreIntegrationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
