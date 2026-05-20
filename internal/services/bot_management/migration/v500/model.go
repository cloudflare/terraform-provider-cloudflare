package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareBotManagementModel represents the legacy cloudflare_bot_management resource state from v4.x provider.
// Schema version: 0 (SDKv2 provider)
// Resource type: cloudflare_bot_management
type SourceCloudflareBotManagementModel struct {
	ID                           types.String `tfsdk:"id"`
	ZoneID                       types.String `tfsdk:"zone_id"`
	AIBotsProtection             types.String `tfsdk:"ai_bots_protection"`
	AutoUpdateModel              types.Bool   `tfsdk:"auto_update_model"`
	EnableJS                     types.Bool   `tfsdk:"enable_js"`
	FightMode                    types.Bool   `tfsdk:"fight_mode"`
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated"`
	SBFMStaticResourceProtection types.Bool   `tfsdk:"sbfm_static_resource_protection"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score"`
	UsingLatestModel             types.Bool   `tfsdk:"using_latest_model"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetBotManagementModel represents the current cloudflare_bot_management resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_bot_management
type TargetBotManagementModel struct {
	ID                           types.String                                                       `tfsdk:"id"`
	ZoneID                       types.String                                                       `tfsdk:"zone_id"`
	AIBotsProtection             types.String                                                       `tfsdk:"ai_bots_protection"`
	AutoUpdateModel              types.Bool                                                         `tfsdk:"auto_update_model"`
	BmCookieEnabled              types.Bool                                                         `tfsdk:"bm_cookie_enabled"`
	CfRobotsVariant              types.String                                                       `tfsdk:"cf_robots_variant"`
	CrawlerProtection            types.String                                                       `tfsdk:"crawler_protection"`
	EnableJS                     types.Bool                                                         `tfsdk:"enable_js"`
	FightMode                    types.Bool                                                         `tfsdk:"fight_mode"`
	IsRobotsTXTManaged           types.Bool                                                         `tfsdk:"is_robots_txt_managed"`
	OptimizeWordpress            types.Bool                                                         `tfsdk:"optimize_wordpress"`
	SBFMDefinitelyAutomated      types.String                                                       `tfsdk:"sbfm_definitely_automated"`
	SBFMLikelyAutomated          types.String                                                       `tfsdk:"sbfm_likely_automated"`
	SBFMStaticResourceProtection types.Bool                                                         `tfsdk:"sbfm_static_resource_protection"`
	SBFMVerifiedBots             types.String                                                       `tfsdk:"sbfm_verified_bots"`
	SuppressSessionScore         types.Bool                                                         `tfsdk:"suppress_session_score"`
	UsingLatestModel             types.Bool                                                         `tfsdk:"using_latest_model"`
	StaleZoneConfiguration       customfield.NestedObject[TargetBotManagementStaleZoneConfigModel]  `tfsdk:"stale_zone_configuration"`
}

// TargetBotManagementStaleZoneConfigModel represents the nested stale_zone_configuration in v5.
type TargetBotManagementStaleZoneConfigModel struct {
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated"`
	SBFMStaticResourceProtection types.String `tfsdk:"sbfm_static_resource_protection"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score"`
	FightMode                    types.Bool   `tfsdk:"fight_mode"`
}
