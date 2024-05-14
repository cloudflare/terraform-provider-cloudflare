package risk_behavior

import "github.com/hashicorp/terraform-plugin-framework/types"

type RiskBehaviorModel struct {
	AccountID types.String                `tfsdk:"account_id"`
	Behaviors []RiskBehaviorBehaviorModel `tfsdk:"behavior"`
}

type RiskBehaviorBehaviorModel struct {
	Enabled   types.Bool   `tfsdk:"enabled"`
	Name      types.String `tfsdk:"name"`
	RiskLevel types.String `tfsdk:"risk_level"`
}
