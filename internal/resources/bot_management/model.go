// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultEnvelope struct {
	Result BotManagementModel `json:"result,computed"`
}

type BotManagementModel struct {
	ZoneID                       types.String `tfsdk:"zone_id" path:"zone_id"`
	EnableJS                     types.Bool   `tfsdk:"enable_js" json:"enable_js"`
	FightMode                    types.Bool   `tfsdk:"fight_mode" json:"fight_mode"`
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress" json:"optimize_wordpress"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated"`
	SBFMStaticResourceProtection types.Bool   `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated"`
	AutoUpdateModel              types.Bool   `tfsdk:"auto_update_model" json:"auto_update_model"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score" json:"suppress_session_score"`
}
