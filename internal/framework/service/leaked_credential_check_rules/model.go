package leaked_credential_check_rules

import "github.com/hashicorp/terraform-plugin-framework/types"

type LeakedCredentialCheckRulesModel struct {
	ZoneID types.String        `tfsdk:"zone_id"`
	Rules  []LCCRuleValueModel `tfsdk:"rule"`
}

type LCCRuleValueModel struct {
	ID       types.String `tfsdk:"id"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}
