package leaked_credential_check_rule

import "github.com/hashicorp/terraform-plugin-framework/types"

type LeakedCredentialCheckRulesModel struct {
	ZoneID   types.String `tfsdk:"zone_id"`
	ID       types.String `tfsdk:"id"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}
