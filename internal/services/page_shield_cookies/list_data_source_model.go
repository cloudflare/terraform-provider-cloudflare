// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_cookies

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldCookiesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PageShieldCookiesListResultDataSourceModel] `json:"result,computed"`
}

type PageShieldCookiesListDataSourceModel struct {
	ZoneID    types.String                                                             `tfsdk:"zone_id" path:"zone_id,required"`
	Direction types.String                                                             `tfsdk:"direction" query:"direction,optional"`
	Domain    types.String                                                             `tfsdk:"domain" query:"domain,optional"`
	Export    types.String                                                             `tfsdk:"export" query:"export,optional"`
	Hosts     types.String                                                             `tfsdk:"hosts" query:"hosts,optional"`
	HTTPOnly  types.Bool                                                               `tfsdk:"http_only" query:"http_only,optional"`
	Name      types.String                                                             `tfsdk:"name" query:"name,optional"`
	OrderBy   types.String                                                             `tfsdk:"order_by" query:"order_by,optional"`
	Page      types.String                                                             `tfsdk:"page" query:"page,optional"`
	PageURL   types.String                                                             `tfsdk:"page_url" query:"page_url,optional"`
	Path      types.String                                                             `tfsdk:"path" query:"path,optional"`
	PerPage   types.Float64                                                            `tfsdk:"per_page" query:"per_page,optional"`
	SameSite  types.String                                                             `tfsdk:"same_site" query:"same_site,optional"`
	Secure    types.Bool                                                               `tfsdk:"secure" query:"secure,optional"`
	Type      types.String                                                             `tfsdk:"type" query:"type,optional"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[PageShieldCookiesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *PageShieldCookiesListDataSourceModel) toListParams(_ context.Context) (params page_shield.CookieListParams, diags diag.Diagnostics) {
	params = page_shield.CookieListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(page_shield.CookieListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Domain.IsNull() {
		params.Domain = cloudflare.F(m.Domain.ValueString())
	}
	if !m.Export.IsNull() {
		params.Export = cloudflare.F(page_shield.CookieListParamsExport(m.Export.ValueString()))
	}
	if !m.Hosts.IsNull() {
		params.Hosts = cloudflare.F(m.Hosts.ValueString())
	}
	if !m.HTTPOnly.IsNull() {
		params.HTTPOnly = cloudflare.F(m.HTTPOnly.ValueBool())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(page_shield.CookieListParamsOrderBy(m.OrderBy.ValueString()))
	}
	if !m.Page.IsNull() {
		params.Page = cloudflare.F(m.Page.ValueString())
	}
	if !m.PageURL.IsNull() {
		params.PageURL = cloudflare.F(m.PageURL.ValueString())
	}
	if !m.Path.IsNull() {
		params.Path = cloudflare.F(m.Path.ValueString())
	}
	if !m.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.PerPage.ValueFloat64())
	}
	if !m.SameSite.IsNull() {
		params.SameSite = cloudflare.F(page_shield.CookieListParamsSameSite(m.SameSite.ValueString()))
	}
	if !m.Secure.IsNull() {
		params.Secure = cloudflare.F(m.Secure.ValueBool())
	}
	if !m.Type.IsNull() {
		params.Type = cloudflare.F(page_shield.CookieListParamsType(m.Type.ValueString()))
	}

	return
}

type PageShieldCookiesListResultDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	FirstSeenAt       timetypes.RFC3339              `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
	Host              types.String                   `tfsdk:"host" json:"host,computed"`
	LastSeenAt        timetypes.RFC3339              `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	Name              types.String                   `tfsdk:"name" json:"name,computed"`
	Type              types.String                   `tfsdk:"type" json:"type,computed"`
	DomainAttribute   types.String                   `tfsdk:"domain_attribute" json:"domain_attribute,computed"`
	ExpiresAttribute  timetypes.RFC3339              `tfsdk:"expires_attribute" json:"expires_attribute,computed" format:"date-time"`
	HTTPOnlyAttribute types.Bool                     `tfsdk:"http_only_attribute" json:"http_only_attribute,computed"`
	MaxAgeAttribute   types.Int64                    `tfsdk:"max_age_attribute" json:"max_age_attribute,computed"`
	PageURLs          customfield.List[types.String] `tfsdk:"page_urls" json:"page_urls,computed"`
	PathAttribute     types.String                   `tfsdk:"path_attribute" json:"path_attribute,computed"`
	SameSiteAttribute types.String                   `tfsdk:"same_site_attribute" json:"same_site_attribute,computed"`
	SecureAttribute   types.Bool                     `tfsdk:"secure_attribute" json:"secure_attribute,computed"`
}
