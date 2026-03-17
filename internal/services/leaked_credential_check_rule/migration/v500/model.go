package v500

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLeakedCredentialCheckRuleModel struct {
	ID       types.String `tfsdk:"id"`
	ZoneID   types.String `tfsdk:"zone_id"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type TargetLeakedCredentialCheckRuleModel struct {
	ID       types.String `tfsdk:"id"`
	ZoneID   types.String `tfsdk:"zone_id"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}
