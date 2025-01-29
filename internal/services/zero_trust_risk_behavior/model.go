// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_behavior

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustRiskBehaviorResultEnvelope struct {
	Result ZeroTrustRiskBehaviorModel `json:"result"`
}

type ZeroTrustRiskBehaviorModel struct {
	AccountID types.String                                    `tfsdk:"account_id" path:"account_id,required"`
	Behaviors *map[string]ZeroTrustRiskBehaviorBehaviorsModel `tfsdk:"behaviors" json:"behaviors,required"`
}

func (m ZeroTrustRiskBehaviorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustRiskBehaviorModel) MarshalJSONForUpdate(state ZeroTrustRiskBehaviorModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustRiskBehaviorBehaviorsModel struct {
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	RiskLevel types.String `tfsdk:"risk_level" json:"risk_level,required"`
}
