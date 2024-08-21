package zero_trust_risk_behavior

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZeroTrustRiskBehaviorModel struct {
	AccountID types.String                         `tfsdk:"account_id"`
	Behaviors []ZeroTrustRiskBehaviorBehaviorModel `tfsdk:"behavior"`
}

type ZeroTrustRiskBehaviorBehaviorModel struct {
	Enabled   types.Bool   `tfsdk:"enabled"`
	Name      types.String `tfsdk:"name"`
	RiskLevel types.String `tfsdk:"risk_level"`
}
