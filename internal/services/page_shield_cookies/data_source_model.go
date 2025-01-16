// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_cookies

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldCookiesResultDataSourceEnvelope struct {
	Result PageShieldCookiesDataSourceModel `json:"result,computed"`
}

type PageShieldCookiesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PageShieldCookiesDataSourceModel] `json:"result,computed"`
}

type PageShieldCookiesDataSourceModel struct {
	CookieID          types.String                               `tfsdk:"cookie_id" path:"cookie_id,optional"`
	ZoneID            types.String                               `tfsdk:"zone_id" path:"zone_id,optional"`
	DomainAttribute   types.String                               `tfsdk:"domain_attribute" json:"domain_attribute,computed"`
	ExpiresAttribute  timetypes.RFC3339                          `tfsdk:"expires_attribute" json:"expires_attribute,computed" format:"date-time"`
	FirstSeenAt       timetypes.RFC3339                          `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
	Host              types.String                               `tfsdk:"host" json:"host,computed"`
	HTTPOnlyAttribute types.Bool                                 `tfsdk:"http_only_attribute" json:"http_only_attribute,computed"`
	ID                types.String                               `tfsdk:"id" json:"id,computed"`
	LastSeenAt        timetypes.RFC3339                          `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	MaxAgeAttribute   types.Int64                                `tfsdk:"max_age_attribute" json:"max_age_attribute,computed"`
	Name              types.String                               `tfsdk:"name" json:"name,computed"`
	PathAttribute     types.String                               `tfsdk:"path_attribute" json:"path_attribute,computed"`
	SameSiteAttribute types.String                               `tfsdk:"same_site_attribute" json:"same_site_attribute,computed"`
	SecureAttribute   types.Bool                                 `tfsdk:"secure_attribute" json:"secure_attribute,computed"`
	Type              types.String                               `tfsdk:"type" json:"type,computed"`
	PageURLs          customfield.List[types.String]             `tfsdk:"page_urls" json:"page_urls,computed"`
	Filter            *PageShieldCookiesFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *PageShieldCookiesDataSourceModel) toReadParams(_ context.Context) (params page_shield.CookieGetParams, diags diag.Diagnostics) {
	params = page_shield.CookieGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *PageShieldCookiesDataSourceModel) toListParams(_ context.Context) (params page_shield.CookieListParams, diags diag.Diagnostics) {
	params = page_shield.CookieListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(page_shield.CookieListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Domain.IsNull() {
		params.Domain = cloudflare.F(m.Filter.Domain.ValueString())
	}
	if !m.Filter.Export.IsNull() {
		params.Export = cloudflare.F(page_shield.CookieListParamsExport(m.Filter.Export.ValueString()))
	}
	if !m.Filter.Hosts.IsNull() {
		params.Hosts = cloudflare.F(m.Filter.Hosts.ValueString())
	}
	if !m.Filter.HTTPOnly.IsNull() {
		params.HTTPOnly = cloudflare.F(m.Filter.HTTPOnly.ValueBool())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(page_shield.CookieListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.Page.IsNull() {
		params.Page = cloudflare.F(m.Filter.Page.ValueString())
	}
	if !m.Filter.PageURL.IsNull() {
		params.PageURL = cloudflare.F(m.Filter.PageURL.ValueString())
	}
	if !m.Filter.Path.IsNull() {
		params.Path = cloudflare.F(m.Filter.Path.ValueString())
	}
	if !m.Filter.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.Filter.PerPage.ValueFloat64())
	}
	if !m.Filter.SameSite.IsNull() {
		params.SameSite = cloudflare.F(page_shield.CookieListParamsSameSite(m.Filter.SameSite.ValueString()))
	}
	if !m.Filter.Secure.IsNull() {
		params.Secure = cloudflare.F(m.Filter.Secure.ValueBool())
	}
	if !m.Filter.Type.IsNull() {
		params.Type = cloudflare.F(page_shield.CookieListParamsType(m.Filter.Type.ValueString()))
	}

	return
}

type PageShieldCookiesFindOneByDataSourceModel struct {
	ZoneID    types.String  `tfsdk:"zone_id" path:"zone_id,required"`
	Direction types.String  `tfsdk:"direction" query:"direction,optional"`
	Domain    types.String  `tfsdk:"domain" query:"domain,optional"`
	Export    types.String  `tfsdk:"export" query:"export,optional"`
	Hosts     types.String  `tfsdk:"hosts" query:"hosts,optional"`
	HTTPOnly  types.Bool    `tfsdk:"http_only" query:"http_only,optional"`
	Name      types.String  `tfsdk:"name" query:"name,optional"`
	OrderBy   types.String  `tfsdk:"order_by" query:"order_by,optional"`
	Page      types.String  `tfsdk:"page" query:"page,optional"`
	PageURL   types.String  `tfsdk:"page_url" query:"page_url,optional"`
	Path      types.String  `tfsdk:"path" query:"path,optional"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page,optional"`
	SameSite  types.String  `tfsdk:"same_site" query:"same_site,optional"`
	Secure    types.Bool    `tfsdk:"secure" query:"secure,optional"`
	Type      types.String  `tfsdk:"type" query:"type,optional"`
}
