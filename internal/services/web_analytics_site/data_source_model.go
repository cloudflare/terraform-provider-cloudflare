// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSiteResultDataSourceEnvelope struct {
	Result WebAnalyticsSiteDataSourceModel `json:"result,computed"`
}

type WebAnalyticsSiteResultListDataSourceEnvelope struct {
	Result *[]*WebAnalyticsSiteDataSourceModel `json:"result,computed"`
}

type WebAnalyticsSiteDataSourceModel struct {
	AccountID   types.String                              `tfsdk:"account_id" path:"account_id"`
	SiteID      types.String                              `tfsdk:"site_id" path:"site_id"`
	AutoInstall types.Bool                                `tfsdk:"auto_install" json:"auto_install"`
	Created     types.String                              `tfsdk:"created" json:"created"`
	Rules       *[]*WebAnalyticsSiteRulesDataSourceModel  `tfsdk:"rules" json:"rules"`
	Ruleset     *WebAnalyticsSiteRulesetDataSourceModel   `tfsdk:"ruleset" json:"ruleset"`
	SiteTag     types.String                              `tfsdk:"site_tag" json:"site_tag"`
	SiteToken   types.String                              `tfsdk:"site_token" json:"site_token"`
	Snippet     types.String                              `tfsdk:"snippet" json:"snippet"`
	FindOneBy   *WebAnalyticsSiteFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type WebAnalyticsSiteRulesDataSourceModel struct {
	ID        types.String  `tfsdk:"id" json:"id"`
	Created   types.String  `tfsdk:"created" json:"created,computed"`
	Host      types.String  `tfsdk:"host" json:"host"`
	Inclusive types.Bool    `tfsdk:"inclusive" json:"inclusive"`
	IsPaused  types.Bool    `tfsdk:"is_paused" json:"is_paused"`
	Paths     types.String  `tfsdk:"paths" json:"paths"`
	Priority  types.Float64 `tfsdk:"priority" json:"priority"`
}

type WebAnalyticsSiteRulesetDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled"`
	ZoneName types.String `tfsdk:"zone_name" json:"zone_name"`
	ZoneTag  types.String `tfsdk:"zone_tag" json:"zone_tag"`
}

type WebAnalyticsSiteFindOneByDataSourceModel struct {
	AccountID types.String  `tfsdk:"account_id" path:"account_id"`
	OrderBy   types.String  `tfsdk:"order_by" query:"order_by"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}
