// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultEnvelope struct {
	Result BotManagementModel `json:"result"`
}

type BotManagementModel struct {
	ID                           types.String                                                       `tfsdk:"id" json:"-,computed"`
	ZoneID                       types.String                                                       `tfsdk:"zone_id" path:"zone_id,required"`
	AIBotsProtection             types.String                                                       `tfsdk:"ai_bots_protection" json:"ai_bots_protection,computed_optional"`
	AutoUpdateModel              types.Bool                                                         `tfsdk:"auto_update_model" json:"auto_update_model,computed_optional,decode_null_to_zero"`
	CrawlerProtection            types.String                                                       `tfsdk:"crawler_protection" json:"crawler_protection,computed_optional"`
	EnableJS                     types.Bool                                                         `tfsdk:"enable_js" json:"enable_js,computed_optional,decode_null_to_zero"`
	FightMode                    types.Bool                                                         `tfsdk:"fight_mode" json:"fight_mode,computed_optional,decode_null_to_zero"`
	IsRobotsTXTManaged           types.Bool                                                         `tfsdk:"is_robots_txt_managed" json:"is_robots_txt_managed,computed_optional,decode_null_to_zero"`
	OptimizeWordpress            types.Bool                                                         `tfsdk:"optimize_wordpress" json:"optimize_wordpress,computed_optional,decode_null_to_zero"`
	SBFMDefinitelyAutomated      types.String                                                       `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated,computed_optional"`
	SBFMLikelyAutomated          types.String                                                       `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated,computed_optional"`
	SBFMStaticResourceProtection types.Bool                                                         `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection,computed_optional,decode_null_to_zero"`
	SBFMVerifiedBots             types.String                                                       `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots,computed_optional"`
	SuppressSessionScore         types.Bool                                                         `tfsdk:"suppress_session_score" json:"suppress_session_score,computed_optional,decode_null_to_zero"`
	UsingLatestModel             types.Bool                                                         `tfsdk:"using_latest_model" json:"using_latest_model,computed"`
	StaleZoneConfiguration       customfield.NestedObject[BotManagementStaleZoneConfigurationModel] `tfsdk:"stale_zone_configuration" json:"stale_zone_configuration,computed"`
}

func (m BotManagementModel) MarshalJSON() (data []byte, err error) {
	return apijsoncustom.MarshalRoot(m)
}

func (m BotManagementModel) MarshalJSONForUpdate(state BotManagementModel) (data []byte, err error) {
	return apijsoncustom.MarshalForUpdate(m, state)
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
