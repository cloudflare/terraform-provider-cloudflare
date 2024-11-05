// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_scripts

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/page_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldScriptsResultDataSourceEnvelope struct {
	Result PageShieldScriptsDataSourceModel `json:"result,computed"`
}

type PageShieldScriptsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PageShieldScriptsDataSourceModel] `json:"result,computed"`
}

type PageShieldScriptsDataSourceModel struct {
	ScriptID                  types.String                                 `tfsdk:"script_id" path:"script_id,optional"`
	ZoneID                    types.String                                 `tfsdk:"zone_id" path:"zone_id,optional"`
	Versions                  *[]*PageShieldScriptsVersionsDataSourceModel `tfsdk:"versions" json:"versions,optional"`
	AddedAt                   timetypes.RFC3339                            `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	CryptominingScore         types.Int64                                  `tfsdk:"cryptomining_score" json:"cryptomining_score,computed"`
	DataflowScore             types.Int64                                  `tfsdk:"dataflow_score" json:"dataflow_score,computed"`
	DomainReportedMalicious   types.Bool                                   `tfsdk:"domain_reported_malicious" json:"domain_reported_malicious,computed"`
	FetchedAt                 types.String                                 `tfsdk:"fetched_at" json:"fetched_at,computed"`
	FirstPageURL              types.String                                 `tfsdk:"first_page_url" json:"first_page_url,computed"`
	FirstSeenAt               timetypes.RFC3339                            `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
	Hash                      types.String                                 `tfsdk:"hash" json:"hash,computed"`
	Host                      types.String                                 `tfsdk:"host" json:"host,computed"`
	ID                        types.String                                 `tfsdk:"id" json:"id,computed"`
	JSIntegrityScore          types.Int64                                  `tfsdk:"js_integrity_score" json:"js_integrity_score,computed"`
	LastSeenAt                timetypes.RFC3339                            `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	MagecartScore             types.Int64                                  `tfsdk:"magecart_score" json:"magecart_score,computed"`
	MalwareScore              types.Int64                                  `tfsdk:"malware_score" json:"malware_score,computed"`
	ObfuscationScore          types.Int64                                  `tfsdk:"obfuscation_score" json:"obfuscation_score,computed"`
	URL                       types.String                                 `tfsdk:"url" json:"url,computed"`
	URLContainsCDNCGIPath     types.Bool                                   `tfsdk:"url_contains_cdn_cgi_path" json:"url_contains_cdn_cgi_path,computed"`
	URLReportedMalicious      types.Bool                                   `tfsdk:"url_reported_malicious" json:"url_reported_malicious,computed"`
	MaliciousDomainCategories customfield.List[types.String]               `tfsdk:"malicious_domain_categories" json:"malicious_domain_categories,computed"`
	MaliciousURLCategories    customfield.List[types.String]               `tfsdk:"malicious_url_categories" json:"malicious_url_categories,computed"`
	PageURLs                  customfield.List[types.String]               `tfsdk:"page_urls" json:"page_urls,computed"`
	Filter                    *PageShieldScriptsFindOneByDataSourceModel   `tfsdk:"filter"`
}

func (m *PageShieldScriptsDataSourceModel) toReadParams(_ context.Context) (params page_shield.ScriptGetParams, diags diag.Diagnostics) {
	params = page_shield.ScriptGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *PageShieldScriptsDataSourceModel) toListParams(_ context.Context) (params page_shield.ScriptListParams, diags diag.Diagnostics) {
	params = page_shield.ScriptListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(page_shield.ScriptListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.ExcludeCDNCGI.IsNull() {
		params.ExcludeCDNCGI = cloudflare.F(m.Filter.ExcludeCDNCGI.ValueBool())
	}
	if !m.Filter.ExcludeDuplicates.IsNull() {
		params.ExcludeDuplicates = cloudflare.F(m.Filter.ExcludeDuplicates.ValueBool())
	}
	if !m.Filter.ExcludeURLs.IsNull() {
		params.ExcludeURLs = cloudflare.F(m.Filter.ExcludeURLs.ValueString())
	}
	if !m.Filter.Export.IsNull() {
		params.Export = cloudflare.F(page_shield.ScriptListParamsExport(m.Filter.Export.ValueString()))
	}
	if !m.Filter.Hosts.IsNull() {
		params.Hosts = cloudflare.F(m.Filter.Hosts.ValueString())
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(page_shield.ScriptListParamsOrderBy(m.Filter.OrderBy.ValueString()))
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

type PageShieldScriptsVersionsDataSourceModel struct {
	CryptominingScore types.Int64  `tfsdk:"cryptomining_score" json:"cryptomining_score,computed"`
	DataflowScore     types.Int64  `tfsdk:"dataflow_score" json:"dataflow_score,computed"`
	FetchedAt         types.String `tfsdk:"fetched_at" json:"fetched_at,computed"`
	Hash              types.String `tfsdk:"hash" json:"hash,computed"`
	JSIntegrityScore  types.Int64  `tfsdk:"js_integrity_score" json:"js_integrity_score,computed"`
	MagecartScore     types.Int64  `tfsdk:"magecart_score" json:"magecart_score,computed"`
	MalwareScore      types.Int64  `tfsdk:"malware_score" json:"malware_score,computed"`
	ObfuscationScore  types.Int64  `tfsdk:"obfuscation_score" json:"obfuscation_score,computed"`
}

type PageShieldScriptsFindOneByDataSourceModel struct {
	ZoneID              types.String  `tfsdk:"zone_id" path:"zone_id,required"`
	Direction           types.String  `tfsdk:"direction" query:"direction,optional"`
	ExcludeCDNCGI       types.Bool    `tfsdk:"exclude_cdn_cgi" query:"exclude_cdn_cgi,computed_optional"`
	ExcludeDuplicates   types.Bool    `tfsdk:"exclude_duplicates" query:"exclude_duplicates,computed_optional"`
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
