package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// For bot_management, v4 and v5 share the same field names and types for all v4 fields.
// The transformation is a direct copy of all v4 fields, with new v5-only fields set to null.
func Transform(ctx context.Context, source SourceCloudflareBotManagementModel) (*TargetBotManagementModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for bot_management migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Direct copy of all v4 fields (same names and types in v5)
	target := &TargetBotManagementModel{
		ID:                           source.ID,
		ZoneID:                       source.ZoneID,
		AIBotsProtection:             source.AIBotsProtection,
		AutoUpdateModel:              source.AutoUpdateModel,
		EnableJS:                     source.EnableJS,
		FightMode:                    source.FightMode,
		OptimizeWordpress:            source.OptimizeWordpress,
		SBFMDefinitelyAutomated:      source.SBFMDefinitelyAutomated,
		SBFMLikelyAutomated:          source.SBFMLikelyAutomated,
		SBFMStaticResourceProtection: source.SBFMStaticResourceProtection,
		SBFMVerifiedBots:             source.SBFMVerifiedBots,
		SuppressSessionScore:         source.SuppressSessionScore,
		UsingLatestModel:             source.UsingLatestModel,
	}

	// New v5-only fields: set to null (will be refreshed from API on next read)
	target.BmCookieEnabled = types.BoolNull()
	target.CfRobotsVariant = types.StringNull()
	target.CrawlerProtection = types.StringNull()
	target.IsRobotsTXTManaged = types.BoolNull()
	target.StaleZoneConfiguration = customfield.NullObject[TargetBotManagementStaleZoneConfigModel](ctx)

	return target, diags
}
