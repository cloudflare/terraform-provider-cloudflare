// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BotManagementResultEnvelope struct {
	Result BotManagementAPIModel `json:"result"`
}

// BotManagementAPIModel represents the API response/request model with pointers for problematic boolean fields
type BotManagementAPIModel struct {
	AIBotsProtection             *string                                      `json:"ai_bots_protection,omitempty"`
	AutoUpdateModel              *bool                                        `json:"auto_update_model,omitempty"`
	BmCookieEnabled              *bool                                        `json:"bm_cookie_enabled,omitempty"`
	CrawlerProtection            *string                                      `json:"crawler_protection,omitempty"`
	EnableJS                     *bool                                        `json:"enable_js,omitempty"`
	FightMode                    *bool                                        `json:"fight_mode,omitempty"`
	IsRobotsTXTManaged           *bool                                        `json:"is_robots_txt_managed,omitempty"`
	OptimizeWordpress            *bool                                        `json:"optimize_wordpress,omitempty"`
	SBFMDefinitelyAutomated      *string                                      `json:"sbfm_definitely_automated,omitempty"`
	SBFMLikelyAutomated          *string                                      `json:"sbfm_likely_automated,omitempty"`
	SBFMStaticResourceProtection *bool                                        `json:"sbfm_static_resource_protection,omitempty"`
	SBFMVerifiedBots             *string                                      `json:"sbfm_verified_bots,omitempty"`
	SuppressSessionScore         *bool                                        `json:"suppress_session_score,omitempty"`
	UsingLatestModel             *bool                                        `json:"using_latest_model,omitempty"`
	StaleZoneConfiguration       *BotManagementStaleZoneConfigurationAPIModel `json:"stale_zone_configuration,omitempty"`
}

type BotManagementStaleZoneConfigurationAPIModel struct {
	OptimizeWordpress            *bool   `json:"optimize_wordpress,omitempty"`
	SBFMDefinitelyAutomated      *string `json:"sbfm_definitely_automated,omitempty"`
	SBFMLikelyAutomated          *string `json:"sbfm_likely_automated,omitempty"`
	SBFMStaticResourceProtection *bool   `json:"sbfm_static_resource_protection,omitempty"`
	SBFMVerifiedBots             *string `json:"sbfm_verified_bots,omitempty"`
	SuppressSessionScore         *bool   `json:"suppress_session_score,omitempty"`
	FightMode                    *bool   `json:"fight_mode,omitempty"`
}

// BotManagementModel represents the Terraform state model
type BotManagementModel struct {
	ID                           types.String                                                       `tfsdk:"id"`
	ZoneID                       types.String                                                       `tfsdk:"zone_id"`
	AIBotsProtection             types.String                                                       `tfsdk:"ai_bots_protection"`
	AutoUpdateModel              types.Bool                                                         `tfsdk:"auto_update_model"`
	BmCookieEnabled              types.Bool                                                         `tfsdk:"bm_cookie_enabled"`
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
	StaleZoneConfiguration       customfield.NestedObject[BotManagementStaleZoneConfigurationModel] `tfsdk:"stale_zone_configuration"`
}

// Helper functions for ToAPIModel conversion
func setBoolField(field types.Bool, target **bool) {
	if !field.IsNull() && !field.IsUnknown() {
		val := field.ValueBool()
		*target = &val
	}
}

func setStringField(field types.String, target **string) {
	if !field.IsNull() && !field.IsUnknown() {
		val := field.ValueString()
		*target = &val
	}
}

// ToAPIModel converts the Terraform state model to the API model for requests
func (m BotManagementModel) ToAPIModel() BotManagementAPIModel {
	api := BotManagementAPIModel{}

	// Convert boolean fields to pointers
	setBoolField(m.AutoUpdateModel, &api.AutoUpdateModel)
	setBoolField(m.BmCookieEnabled, &api.BmCookieEnabled)
	setBoolField(m.EnableJS, &api.EnableJS)
	setBoolField(m.FightMode, &api.FightMode)
	setBoolField(m.IsRobotsTXTManaged, &api.IsRobotsTXTManaged)
	setBoolField(m.OptimizeWordpress, &api.OptimizeWordpress)
	setBoolField(m.SBFMStaticResourceProtection, &api.SBFMStaticResourceProtection)
	setBoolField(m.SuppressSessionScore, &api.SuppressSessionScore)
	setBoolField(m.UsingLatestModel, &api.UsingLatestModel)

	// Convert string fields to pointers
	setStringField(m.AIBotsProtection, &api.AIBotsProtection)
	setStringField(m.CrawlerProtection, &api.CrawlerProtection)
	setStringField(m.SBFMDefinitelyAutomated, &api.SBFMDefinitelyAutomated)
	setStringField(m.SBFMLikelyAutomated, &api.SBFMLikelyAutomated)
	setStringField(m.SBFMVerifiedBots, &api.SBFMVerifiedBots)

	return api
}

// Helper functions to reduce repetition
func updateBoolField(field *types.Bool, apiValue *bool) {
	if apiValue != nil {
		*field = types.BoolValue(*apiValue)
	} else if field.IsUnknown() {
		*field = types.BoolNull()
	}
}

func updateStringField(field *types.String, apiValue *string) {
	if apiValue != nil {
		*field = types.StringValue(*apiValue)
	} else if field.IsUnknown() {
		*field = types.StringNull()
	}
}

// Helper function for boolean-to-string conversion in nested config
func updateBoolToStringField(field *types.String, apiValue *bool) {
	if apiValue != nil {
		*field = types.StringValue(fmt.Sprintf("%t", *apiValue))
	} else if field.IsUnknown() {
		*field = types.StringNull()
	}
}

// Helper function to update nested stale zone configuration
func updateStaleZoneConfig(staleConfig *BotManagementStaleZoneConfigurationModel, api *BotManagementStaleZoneConfigurationAPIModel) {
	if api == nil {
		return
	}

	// Boolean fields
	updateBoolField(&staleConfig.OptimizeWordpress, api.OptimizeWordpress)
	updateBoolField(&staleConfig.SuppressSessionScore, api.SuppressSessionScore)
	updateBoolField(&staleConfig.FightMode, api.FightMode)

	// String fields
	updateStringField(&staleConfig.SBFMDefinitelyAutomated, api.SBFMDefinitelyAutomated)
	updateStringField(&staleConfig.SBFMLikelyAutomated, api.SBFMLikelyAutomated)
	updateStringField(&staleConfig.SBFMVerifiedBots, api.SBFMVerifiedBots)

	// Boolean-to-string conversion for SBFMStaticResourceProtection
	updateBoolToStringField(&staleConfig.SBFMStaticResourceProtection, api.SBFMStaticResourceProtection)
}

// UpdateFromAPIModel updates the Terraform state model from API response, preserving existing values for missing fields
func (m *BotManagementModel) UpdateFromAPIModel(api BotManagementAPIModel) {
	// Update boolean fields
	updateBoolField(&m.AutoUpdateModel, api.AutoUpdateModel)
	updateBoolField(&m.BmCookieEnabled, api.BmCookieEnabled)
	updateBoolField(&m.EnableJS, api.EnableJS)
	updateBoolField(&m.FightMode, api.FightMode)
	updateBoolField(&m.IsRobotsTXTManaged, api.IsRobotsTXTManaged)
	updateBoolField(&m.OptimizeWordpress, api.OptimizeWordpress)
	updateBoolField(&m.SBFMStaticResourceProtection, api.SBFMStaticResourceProtection)
	updateBoolField(&m.SuppressSessionScore, api.SuppressSessionScore)
	updateBoolField(&m.UsingLatestModel, api.UsingLatestModel)

	// Update string fields
	updateStringField(&m.AIBotsProtection, api.AIBotsProtection)
	updateStringField(&m.CrawlerProtection, api.CrawlerProtection)
	updateStringField(&m.SBFMDefinitelyAutomated, api.SBFMDefinitelyAutomated)
	updateStringField(&m.SBFMLikelyAutomated, api.SBFMLikelyAutomated)
	updateStringField(&m.SBFMVerifiedBots, api.SBFMVerifiedBots)

	// Handle nested stale zone configuration
	if api.StaleZoneConfiguration != nil {
		// Convert nested API model to state model using helper function
		staleConfig := BotManagementStaleZoneConfigurationModel{}
		updateStaleZoneConfig(&staleConfig, api.StaleZoneConfiguration)

		nestedObj, _ := customfield.NewObject[BotManagementStaleZoneConfigurationModel](context.TODO(), &staleConfig)
		m.StaleZoneConfiguration = nestedObj
	} else if m.StaleZoneConfiguration.IsUnknown() {
		m.StaleZoneConfiguration = customfield.NullObject[BotManagementStaleZoneConfigurationModel](context.TODO())
	}
}

type BotManagementStaleZoneConfigurationModel struct {
	OptimizeWordpress            types.Bool   `tfsdk:"optimize_wordpress"`
	SBFMDefinitelyAutomated      types.String `tfsdk:"sbfm_definitely_automated"`
	SBFMLikelyAutomated          types.String `tfsdk:"sbfm_likely_automated"`
	SBFMStaticResourceProtection types.String `tfsdk:"sbfm_static_resource_protection"`
	SBFMVerifiedBots             types.String `tfsdk:"sbfm_verified_bots"`
	SuppressSessionScore         types.Bool   `tfsdk:"suppress_session_score"`
	FightMode                    types.Bool   `tfsdk:"fight_mode"`
}
