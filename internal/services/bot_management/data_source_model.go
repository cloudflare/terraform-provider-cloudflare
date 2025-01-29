// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/bot_management"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultDataSourceEnvelope struct {
	Result BotManagementDataSourceModel `json:"result,computed"`
}

type BotManagementDataSourceModel struct {
	ZoneID                       types.String                                                                 `tfsdk:"zone_id" path:"zone_id,required"`
	AIBotsProtection             types.String                                                                 `tfsdk:"ai_bots_protection" json:"ai_bots_protection,computed"`
	AutoUpdateModel              types.Bool                                                                   `tfsdk:"auto_update_model" json:"auto_update_model,computed"`
	EnableJS                     types.Bool                                                                   `tfsdk:"enable_js" json:"enable_js,computed"`
	FightMode                    types.Bool                                                                   `tfsdk:"fight_mode" json:"fight_mode,computed"`
	OptimizeWordpress            types.Bool                                                                   `tfsdk:"optimize_wordpress" json:"optimize_wordpress,computed"`
	SBFMDefinitelyAutomated      types.String                                                                 `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated,computed"`
	SBFMLikelyAutomated          types.String                                                                 `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated,computed"`
	SBFMStaticResourceProtection types.Bool                                                                   `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection,computed"`
	SBFMVerifiedBots             types.String                                                                 `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots,computed"`
	SuppressSessionScore         types.Bool                                                                   `tfsdk:"suppress_session_score" json:"suppress_session_score,computed"`
	UsingLatestModel             types.Bool                                                                   `tfsdk:"using_latest_model" json:"using_latest_model,computed"`
	StaleZoneConfiguration       customfield.NestedObject[BotManagementStaleZoneConfigurationDataSourceModel] `tfsdk:"stale_zone_configuration" json:"stale_zone_configuration,computed"`
}

func (m *BotManagementDataSourceModel) toReadParams(_ context.Context) (params bot_management.BotManagementGetParams, diags diag.Diagnostics) {
	params = bot_management.BotManagementGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type BotManagementStaleZoneConfigurationDataSourceModel struct {
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress" json:"optimize_wordpress,computed"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated" json:"sbfm_definitely_automated,computed"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated" json:"sbfm_likely_automated,computed"`
	SBFMStaticResourceProtection types.String `tfsdk:"sbfm_static_resource_protection" json:"sbfm_static_resource_protection,computed"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots" json:"sbfm_verified_bots,computed"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score" json:"suppress_session_score,computed"`
	FightMode                    types.Bool   `tfsdk:"fight_mode" json:"fight_mode,computed"`
}
