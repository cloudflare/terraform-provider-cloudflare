// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_behavior

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustRiskBehaviorResultDataSourceEnvelope struct {
	Result ZeroTrustRiskBehaviorDataSourceModel `json:"result,computed"`
}

type ZeroTrustRiskBehaviorDataSourceModel struct {
	AccountID types.String                                                               `tfsdk:"account_id" path:"account_id,required"`
	Behaviors customfield.NestedObjectMap[ZeroTrustRiskBehaviorBehaviorsDataSourceModel] `tfsdk:"behaviors" json:"behaviors,computed"`
}

func (m *ZeroTrustRiskBehaviorDataSourceModel) toReadParams(_ context.Context) (params zero_trust.RiskScoringBehaviourGetParams, diags diag.Diagnostics) {
	params = zero_trust.RiskScoringBehaviourGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustRiskBehaviorBehaviorsDataSourceModel struct {
	Description types.String `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	RiskLevel   types.String `tfsdk:"risk_level" json:"risk_level,computed"`
}
