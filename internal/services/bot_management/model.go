// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultEnvelope struct {
	Result BotManagementModel `json:"result"`
}

type BotManagementModel struct {
	ID                           types.String `tfsdk:"id" json:"-,computed"`
	ZoneID                       types.String `tfsdk:"zone_id" path:"zone_id"`
	AutoUpdateModel              types.Bool   `tfsdk:"auto_update_model" json:"auto_update_model"`
	EnableJS                     types.Bool   `tfsdk:"enable_js" json:"enable_js"`
	FightMode                    types.Bool   `tfsdk:"fight_mode" json:"fight_mode"`
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress" json:"optimize_wordpress"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated"`
	SBFMStaticResourceProtection types.Bool   `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score" json:"suppress_session_score,computed_optional"`
}
