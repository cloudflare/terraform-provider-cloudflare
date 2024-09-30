// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_scoring_integration

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustRiskScoringIntegrationResultEnvelope struct {
	Result ZeroTrustRiskScoringIntegrationModel `json:"result"`
}

type ZeroTrustRiskScoringIntegrationModel struct {
	ID              types.String      `tfsdk:"id" json:"id,computed"`
	AccountID       types.String      `tfsdk:"account_id" path:"account_id,required"`
	IntegrationType types.String      `tfsdk:"integration_type" json:"integration_type,required"`
	TenantURL       types.String      `tfsdk:"tenant_url" json:"tenant_url,required"`
	Active          types.Bool        `tfsdk:"active" json:"active,optional"`
	ReferenceID     types.String      `tfsdk:"reference_id" json:"reference_id,optional"`
	AccountTag      types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	WellKnownURL    types.String      `tfsdk:"well_known_url" json:"well_known_url,computed"`
}
