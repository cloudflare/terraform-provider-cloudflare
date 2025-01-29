// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_connections

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldConnectionsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PageShieldConnectionsListResultDataSourceModel] `json:"result,computed"`
}

type PageShieldConnectionsListDataSourceModel struct {
	ZoneID              types.String                                                                 `tfsdk:"zone_id" path:"zone_id,required"`
	Direction           types.String                                                                 `tfsdk:"direction" query:"direction,optional"`
	ExcludeCDNCGI       types.Bool                                                                   `tfsdk:"exclude_cdn_cgi" query:"exclude_cdn_cgi,optional"`
	ExcludeURLs         types.String                                                                 `tfsdk:"exclude_urls" query:"exclude_urls,optional"`
	Export              types.String                                                                 `tfsdk:"export" query:"export,optional"`
	Hosts               types.String                                                                 `tfsdk:"hosts" query:"hosts,optional"`
	OrderBy             types.String                                                                 `tfsdk:"order_by" query:"order_by,optional"`
	Page                types.String                                                                 `tfsdk:"page" query:"page,optional"`
	PageURL             types.String                                                                 `tfsdk:"page_url" query:"page_url,optional"`
	PerPage             types.Float64                                                                `tfsdk:"per_page" query:"per_page,optional"`
	PrioritizeMalicious types.Bool                                                                   `tfsdk:"prioritize_malicious" query:"prioritize_malicious,optional"`
	Status              types.String                                                                 `tfsdk:"status" query:"status,optional"`
	URLs                types.String                                                                 `tfsdk:"urls" query:"urls,optional"`
	MaxItems            types.Int64                                                                  `tfsdk:"max_items"`
	Result              customfield.NestedObjectList[PageShieldConnectionsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *PageShieldConnectionsListDataSourceModel) toListParams(_ context.Context) (params page_shield.ConnectionListParams, diags diag.Diagnostics) {
	params = page_shield.ConnectionListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(page_shield.ConnectionListParamsDirection(m.Direction.ValueString()))
	}
	if !m.ExcludeCDNCGI.IsNull() {
		params.ExcludeCDNCGI = cloudflare.F(m.ExcludeCDNCGI.ValueBool())
	}
	if !m.ExcludeURLs.IsNull() {
		params.ExcludeURLs = cloudflare.F(m.ExcludeURLs.ValueString())
	}
	if !m.Export.IsNull() {
		params.Export = cloudflare.F(page_shield.ConnectionListParamsExport(m.Export.ValueString()))
	}
	if !m.Hosts.IsNull() {
		params.Hosts = cloudflare.F(m.Hosts.ValueString())
	}
	if !m.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(page_shield.ConnectionListParamsOrderBy(m.OrderBy.ValueString()))
	}
	if !m.Page.IsNull() {
		params.Page = cloudflare.F(m.Page.ValueString())
	}
	if !m.PageURL.IsNull() {
		params.PageURL = cloudflare.F(m.PageURL.ValueString())
	}
	if !m.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.PerPage.ValueFloat64())
	}
	if !m.PrioritizeMalicious.IsNull() {
		params.PrioritizeMalicious = cloudflare.F(m.PrioritizeMalicious.ValueBool())
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(m.Status.ValueString())
	}
	if !m.URLs.IsNull() {
		params.URLs = cloudflare.F(m.URLs.ValueString())
	}

	return
}

type PageShieldConnectionsListResultDataSourceModel struct {
	ID                        types.String                   `tfsdk:"id" json:"id,computed"`
	AddedAt                   timetypes.RFC3339              `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	FirstSeenAt               timetypes.RFC3339              `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
	Host                      types.String                   `tfsdk:"host" json:"host,computed"`
	LastSeenAt                timetypes.RFC3339              `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	URL                       types.String                   `tfsdk:"url" json:"url,computed"`
	URLContainsCDNCGIPath     types.Bool                     `tfsdk:"url_contains_cdn_cgi_path" json:"url_contains_cdn_cgi_path,computed"`
	DomainReportedMalicious   types.Bool                     `tfsdk:"domain_reported_malicious" json:"domain_reported_malicious,computed"`
	FirstPageURL              types.String                   `tfsdk:"first_page_url" json:"first_page_url,computed"`
	MaliciousDomainCategories customfield.List[types.String] `tfsdk:"malicious_domain_categories" json:"malicious_domain_categories,computed"`
	MaliciousURLCategories    customfield.List[types.String] `tfsdk:"malicious_url_categories" json:"malicious_url_categories,computed"`
	PageURLs                  customfield.List[types.String] `tfsdk:"page_urls" json:"page_urls,computed"`
	URLReportedMalicious      types.Bool                     `tfsdk:"url_reported_malicious" json:"url_reported_malicious,computed"`
}
