package zero_trust_risk_score_integration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RiskScoreIntegrationModel struct {
	AccountID       types.String `tfsdk:"account_id"`
	ID              types.String `tfsdk:"id"`
	IntegrationType types.String `tfsdk:"integration_type"`
	TenantUrl       types.String `tfsdk:"tenant_url"`
	ReferenceID     types.String `tfsdk:"reference_id"`
	Active          types.Bool   `tfsdk:"active"`
	WellKnownUrl    types.String `tfsdk:"well_known_url"`
}
