// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_connections

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldConnectionsResultDataSourceEnvelope struct {
	Result PageShieldConnectionsDataSourceModel `json:"result,computed"`
}

type PageShieldConnectionsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PageShieldConnectionsDataSourceModel] `json:"result,computed"`
}

type PageShieldConnectionsDataSourceModel struct {
	ConnectionID              types.String                                   `tfsdk:"connection_id" path:"connection_id,optional"`
	ZoneID                    types.String                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	AddedAt                   timetypes.RFC3339                              `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	DomainReportedMalicious   types.Bool                                     `tfsdk:"domain_reported_malicious" json:"domain_reported_malicious,computed"`
	FirstPageURL              types.String                                   `tfsdk:"first_page_url" json:"first_page_url,computed"`
	FirstSeenAt               timetypes.RFC3339                              `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
	Host                      types.String                                   `tfsdk:"host" json:"host,computed"`
	ID                        types.String                                   `tfsdk:"id" json:"id,computed"`
	LastSeenAt                timetypes.RFC3339                              `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	URL                       types.String                                   `tfsdk:"url" json:"url,computed"`
	URLContainsCDNCGIPath     types.Bool                                     `tfsdk:"url_contains_cdn_cgi_path" json:"url_contains_cdn_cgi_path,computed"`
	URLReportedMalicious      types.Bool                                     `tfsdk:"url_reported_malicious" json:"url_reported_malicious,computed"`
	MaliciousDomainCategories customfield.List[types.String]                 `tfsdk:"malicious_domain_categories" json:"malicious_domain_categories,computed"`
	MaliciousURLCategories    customfield.List[types.String]                 `tfsdk:"malicious_url_categories" json:"malicious_url_categories,computed"`
	PageURLs                  customfield.List[types.String]                 `tfsdk:"page_urls" json:"page_urls,computed"`
	Filter                    *PageShieldConnectionsFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *PageShieldConnectionsDataSourceModel) toReadParams(_ context.Context) (params page_shield.ConnectionGetParams, diags diag.Diagnostics) {
	params = page_shield.ConnectionGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *PageShieldConnectionsDataSourceModel) toListParams(_ context.Context) (params page_shield.ConnectionListParams, diags diag.Diagnostics) {
	params = page_shield.ConnectionListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(page_shield.ConnectionListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.ExcludeCDNCGI.IsNull() {
		params.ExcludeCDNCGI = cloudflare.F(m.Filter.ExcludeCDNCGI.ValueBool())
	}
	if !m.Filter.ExcludeURLs.IsNull() {
		params.ExcludeURLs = cloudflare.F(m.Filter.ExcludeURLs.ValueString())
	}
	if !m.Filter.Export.IsNull() {
		params.Export = cloudflare.F(page_shield.ConnectionListParamsExport(m.Filter.Export.ValueString()))
	}
	if !m.Filter.Hosts.IsNull() {
		params.Hosts = cloudflare.F(m.Filter.Hosts.ValueString())
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(page_shield.ConnectionListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.Page.IsNull() {
		params.Page = cloudflare.F(m.Filter.Page.ValueString())
	}
	if !m.Filter.PageURL.IsNull() {
		params.PageURL = cloudflare.F(m.Filter.PageURL.ValueString())
	}
	if !m.Filter.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.Filter.PerPage.ValueFloat64())
	}
	if !m.Filter.PrioritizeMalicious.IsNull() {
		params.PrioritizeMalicious = cloudflare.F(m.Filter.PrioritizeMalicious.ValueBool())
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(m.Filter.Status.ValueString())
	}
	if !m.Filter.URLs.IsNull() {
		params.URLs = cloudflare.F(m.Filter.URLs.ValueString())
	}

	return
}

type PageShieldConnectionsFindOneByDataSourceModel struct {
	ZoneID              types.String  `tfsdk:"zone_id" path:"zone_id,required"`
	Direction           types.String  `tfsdk:"direction" query:"direction,optional"`
	ExcludeCDNCGI       types.Bool    `tfsdk:"exclude_cdn_cgi" query:"exclude_cdn_cgi,optional"`
	ExcludeURLs         types.String  `tfsdk:"exclude_urls" query:"exclude_urls,optional"`
	Export              types.String  `tfsdk:"export" query:"export,optional"`
	Hosts               types.String  `tfsdk:"hosts" query:"hosts,optional"`
	OrderBy             types.String  `tfsdk:"order_by" query:"order_by,optional"`
	Page                types.String  `tfsdk:"page" query:"page,optional"`
	PageURL             types.String  `tfsdk:"page_url" query:"page_url,optional"`
	PerPage             types.Float64 `tfsdk:"per_page" query:"per_page,optional"`
	PrioritizeMalicious types.Bool    `tfsdk:"prioritize_malicious" query:"prioritize_malicious,optional"`
	Status              types.String  `tfsdk:"status" query:"status,optional"`
	URLs                types.String  `tfsdk:"urls" query:"urls,optional"`
}
