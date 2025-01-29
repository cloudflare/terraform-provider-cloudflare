// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/rum"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSitesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WebAnalyticsSitesResultDataSourceModel] `json:"result,computed"`
}

type WebAnalyticsSitesDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	OrderBy   types.String                                                         `tfsdk:"order_by" query:"order_by,optional"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WebAnalyticsSitesResultDataSourceModel] `tfsdk:"result"`
}

func (m *WebAnalyticsSitesDataSourceModel) toListParams(_ context.Context) (params rum.SiteInfoListParams, diags diag.Diagnostics) {
	params = rum.SiteInfoListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(rum.SiteInfoListParamsOrderBy(m.OrderBy.ValueString()))
	}

	return
}

type WebAnalyticsSitesResultDataSourceModel struct {
	AutoInstall types.Bool                                                          `tfsdk:"auto_install" json:"auto_install,computed"`
	Created     timetypes.RFC3339                                                   `tfsdk:"created" json:"created,computed" format:"date-time"`
	Rules       customfield.NestedObjectList[WebAnalyticsSitesRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
	Ruleset     customfield.NestedObject[WebAnalyticsSitesRulesetDataSourceModel]   `tfsdk:"ruleset" json:"ruleset,computed"`
	SiteTag     types.String                                                        `tfsdk:"site_tag" json:"site_tag,computed"`
	SiteToken   types.String                                                        `tfsdk:"site_token" json:"site_token,computed"`
	Snippet     types.String                                                        `tfsdk:"snippet" json:"snippet,computed"`
}

type WebAnalyticsSitesRulesDataSourceModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	Created   timetypes.RFC3339              `tfsdk:"created" json:"created,computed" format:"date-time"`
	Host      types.String                   `tfsdk:"host" json:"host,computed"`
	Inclusive types.Bool                     `tfsdk:"inclusive" json:"inclusive,computed"`
	IsPaused  types.Bool                     `tfsdk:"is_paused" json:"is_paused,computed"`
	Paths     customfield.List[types.String] `tfsdk:"paths" json:"paths,computed"`
	Priority  types.Float64                  `tfsdk:"priority" json:"priority,computed"`
}

type WebAnalyticsSitesRulesetDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	ZoneName types.String `tfsdk:"zone_name" json:"zone_name,computed"`
	ZoneTag  types.String `tfsdk:"zone_tag" json:"zone_tag,computed"`
}
