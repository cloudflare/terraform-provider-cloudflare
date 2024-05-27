// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSiteResultEnvelope struct {
	Result WebAnalyticsSiteModel `json:"result,computed"`
}

type WebAnalyticsSiteModel struct {
	SiteTag     types.String `tfsdk:"site_tag" json:"site_tag,computed"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	AutoInstall types.Bool   `tfsdk:"auto_install" json:"auto_install"`
	Host        types.String `tfsdk:"host" json:"host"`
	ZoneTag     types.String `tfsdk:"zone_tag" json:"zone_tag"`
}
