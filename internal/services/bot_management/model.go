// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultEnvelope struct {
	Result BotManagementModel `json:"result"`
}

type BotManagementModel struct {
	ID                           types.String                                                       `tfsdk:"id" json:"-,computed"`
	ZoneID                       types.String                                                       `tfsdk:"zone_id" path:"zone_id,required"`
	AIBotsProtection             types.String                                                       `tfsdk:"ai_bots_protection" json:"ai_bots_protection,optional"`
	AutoUpdateModel              types.Bool                                                         `tfsdk:"auto_update_model" json:"auto_update_model,optional"`
	EnableJS                     types.Bool                                                         `tfsdk:"enable_js" json:"enable_js,optional"`
	FightMode                    types.Bool                                                         `tfsdk:"fight_mode" json:"fight_mode,optional"`
	OptimizeWordpress            types.Bool                                                         `tfsdk:"optimize_wordpress" json:"optimize_wordpress,optional"`
	SBFMDefinitelyAutomated      types.String                                                       `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated,optional"`
	SBFMLikelyAutomated          types.String                                                       `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated,optional"`
	SBFMStaticResourceProtection types.Bool                                                         `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection,optional"`
	SBFMVerifiedBots             types.String                                                       `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots,optional"`
	SuppressSessionScore         types.Bool                                                         `tfsdk:"suppress_session_score" json:"suppress_session_score,computed_optional"`
	UsingLatestModel             types.Bool                                                         `tfsdk:"using_latest_model" json:"using_latest_model,computed"`
	StaleZoneConfiguration       customfield.NestedObject[BotManagementStaleZoneConfigurationModel] `tfsdk:"stale_zone_configuration" json:"stale_zone_configuration,computed"`
}

func (m BotManagementModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BotManagementModel) MarshalJSONForUpdate(state BotManagementModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type BotManagementStaleZoneConfigurationModel struct {
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress" json:"optimize_wordpress,computed"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated,computed"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated,computed"`
	SBFMStaticResourceProtection types.String `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection,computed"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots,computed"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score" json:"suppress_session_score,computed"`
	FightMode                    types.Bool   `tfsdk:"fight_mode" json:"fight_mode,computed"`
}
