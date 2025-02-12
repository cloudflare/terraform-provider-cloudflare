// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/rum"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSiteResultDataSourceEnvelope struct {
	Result WebAnalyticsSiteDataSourceModel `json:"result,computed"`
}

type WebAnalyticsSiteDataSourceModel struct {
	ID          types.String                                                       `tfsdk:"id" json:"-,computed"`
	SiteID      types.String                                                       `tfsdk:"site_id" path:"site_id,optional"`
	AccountID   types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	AutoInstall types.Bool                                                         `tfsdk:"auto_install" json:"auto_install,computed"`
	Created     timetypes.RFC3339                                                  `tfsdk:"created" json:"created,computed" format:"date-time"`
	SiteTag     types.String                                                       `tfsdk:"site_tag" json:"site_tag,computed"`
	SiteToken   types.String                                                       `tfsdk:"site_token" json:"site_token,computed"`
	Snippet     types.String                                                       `tfsdk:"snippet" json:"snippet,computed"`
	Rules       customfield.NestedObjectList[WebAnalyticsSiteRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
	Ruleset     customfield.NestedObject[WebAnalyticsSiteRulesetDataSourceModel]   `tfsdk:"ruleset" json:"ruleset,computed"`
	Filter      *WebAnalyticsSiteFindOneByDataSourceModel                          `tfsdk:"filter"`
}

func (m *WebAnalyticsSiteDataSourceModel) toReadParams(_ context.Context) (params rum.SiteInfoGetParams, diags diag.Diagnostics) {
	params = rum.SiteInfoGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WebAnalyticsSiteDataSourceModel) toListParams(_ context.Context) (params rum.SiteInfoListParams, diags diag.Diagnostics) {
	params = rum.SiteInfoListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(rum.SiteInfoListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}

	return
}

type WebAnalyticsSiteRulesDataSourceModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	Created   timetypes.RFC3339              `tfsdk:"created" json:"created,computed" format:"date-time"`
	Host      types.String                   `tfsdk:"host" json:"host,computed"`
	Inclusive types.Bool                     `tfsdk:"inclusive" json:"inclusive,computed"`
	IsPaused  types.Bool                     `tfsdk:"is_paused" json:"is_paused,computed"`
	Paths     customfield.List[types.String] `tfsdk:"paths" json:"paths,computed"`
	Priority  types.Float64                  `tfsdk:"priority" json:"priority,computed"`
}

type WebAnalyticsSiteRulesetDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	ZoneName types.String `tfsdk:"zone_name" json:"zone_name,computed"`
	ZoneTag  types.String `tfsdk:"zone_tag" json:"zone_tag,computed"`
}

type WebAnalyticsSiteFindOneByDataSourceModel struct {
	OrderBy types.String `tfsdk:"order_by" query:"order_by,optional"`
}
