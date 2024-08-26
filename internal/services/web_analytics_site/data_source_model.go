// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/rum"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSiteResultDataSourceEnvelope struct {
	Result WebAnalyticsSiteDataSourceModel `json:"result,computed"`
}

type WebAnalyticsSiteResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WebAnalyticsSiteDataSourceModel] `json:"result,computed"`
}

type WebAnalyticsSiteDataSourceModel struct {
	AccountID   types.String                              `tfsdk:"account_id" path:"account_id"`
	SiteID      types.String                              `tfsdk:"site_id" path:"site_id"`
	Created     timetypes.RFC3339                         `tfsdk:"created" json:"created,computed"`
	AutoInstall types.Bool                                `tfsdk:"auto_install" json:"auto_install"`
	SiteTag     types.String                              `tfsdk:"site_tag" json:"site_tag"`
	SiteToken   types.String                              `tfsdk:"site_token" json:"site_token"`
	Snippet     types.String                              `tfsdk:"snippet" json:"snippet"`
	Rules       *[]*WebAnalyticsSiteRulesDataSourceModel  `tfsdk:"rules" json:"rules"`
	Ruleset     *WebAnalyticsSiteRulesetDataSourceModel   `tfsdk:"ruleset" json:"ruleset"`
	Filter      *WebAnalyticsSiteFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WebAnalyticsSiteDataSourceModel) toReadParams() (params rum.SiteInfoGetParams, diags diag.Diagnostics) {
	params = rum.SiteInfoGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WebAnalyticsSiteDataSourceModel) toListParams() (params rum.SiteInfoListParams, diags diag.Diagnostics) {
	params = rum.SiteInfoListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(rum.SiteInfoListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}

	return
}

type WebAnalyticsSiteRulesDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed"`
	Host      types.String      `tfsdk:"host" json:"host"`
	Inclusive types.Bool        `tfsdk:"inclusive" json:"inclusive"`
	IsPaused  types.Bool        `tfsdk:"is_paused" json:"is_paused"`
	Paths     *[]types.String   `tfsdk:"paths" json:"paths"`
	Priority  types.Float64     `tfsdk:"priority" json:"priority"`
}

type WebAnalyticsSiteRulesetDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled"`
	ZoneName types.String `tfsdk:"zone_name" json:"zone_name"`
	ZoneTag  types.String `tfsdk:"zone_tag" json:"zone_tag"`
}

type WebAnalyticsSiteFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	OrderBy   types.String `tfsdk:"order_by" query:"order_by"`
}
