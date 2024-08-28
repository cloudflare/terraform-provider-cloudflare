// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSiteResultEnvelope struct {
	Result WebAnalyticsSiteModel `json:"result"`
}

type WebAnalyticsSiteModel struct {
	ID          types.String                                             `tfsdk:"id" json:"-,computed"`
	SiteTag     types.String                                             `tfsdk:"site_tag" json:"site_tag,computed"`
	AccountID   types.String                                             `tfsdk:"account_id" path:"account_id"`
	AutoInstall types.Bool                                               `tfsdk:"auto_install" json:"auto_install"`
	Host        types.String                                             `tfsdk:"host" json:"host"`
	ZoneTag     types.String                                             `tfsdk:"zone_tag" json:"zone_tag"`
	Created     timetypes.RFC3339                                        `tfsdk:"created" json:"created,computed" format:"date-time"`
	SiteToken   types.String                                             `tfsdk:"site_token" json:"site_token,computed"`
	Snippet     types.String                                             `tfsdk:"snippet" json:"snippet,computed"`
	Rules       customfield.NestedObjectList[WebAnalyticsSiteRulesModel] `tfsdk:"rules" json:"rules,computed"`
	Ruleset     customfield.NestedObject[WebAnalyticsSiteRulesetModel]   `tfsdk:"ruleset" json:"ruleset,computed"`
}

type WebAnalyticsSiteRulesModel struct {
	ID        types.String      `tfsdk:"id" json:"id"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Host      types.String      `tfsdk:"host" json:"host"`
	Inclusive types.Bool        `tfsdk:"inclusive" json:"inclusive"`
	IsPaused  types.Bool        `tfsdk:"is_paused" json:"is_paused"`
	Paths     *[]types.String   `tfsdk:"paths" json:"paths"`
	Priority  types.Float64     `tfsdk:"priority" json:"priority"`
}

type WebAnalyticsSiteRulesetModel struct {
	ID       types.String `tfsdk:"id" json:"id"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled"`
	ZoneName types.String `tfsdk:"zone_name" json:"zone_name"`
	ZoneTag  types.String `tfsdk:"zone_tag" json:"zone_tag"`
}