// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_scripts

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/page_shield"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldScriptsListResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[PageShieldScriptsListResultDataSourceModel] `json:"result,computed"`
}

type PageShieldScriptsListDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Direction types.String `tfsdk:"direction" query:"direction,optional"`
ExcludeURLs types.String `tfsdk:"exclude_urls" query:"exclude_urls,optional"`
Export types.String `tfsdk:"export" query:"export,optional"`
Hosts types.String `tfsdk:"hosts" query:"hosts,optional"`
OrderBy types.String `tfsdk:"order_by" query:"order_by,optional"`
Page types.String `tfsdk:"page" query:"page,optional"`
PageURL types.String `tfsdk:"page_url" query:"page_url,optional"`
PerPage types.Float64 `tfsdk:"per_page" query:"per_page,optional"`
PrioritizeMalicious types.Bool `tfsdk:"prioritize_malicious" query:"prioritize_malicious,optional"`
Status types.String `tfsdk:"status" query:"status,optional"`
URLs types.String `tfsdk:"urls" query:"urls,optional"`
ExcludeCDNCGI types.Bool `tfsdk:"exclude_cdn_cgi" query:"exclude_cdn_cgi,computed_optional"`
ExcludeDuplicates types.Bool `tfsdk:"exclude_duplicates" query:"exclude_duplicates,computed_optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[PageShieldScriptsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *PageShieldScriptsListDataSourceModel) toListParams(_ context.Context) (params page_shield.ScriptListParams, diags diag.Diagnostics) {
  params = page_shield.ScriptListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  if !m.Direction.IsNull() {
    params.Direction = cloudflare.F(page_shield.ScriptListParamsDirection(m.Direction.ValueString()))
  }
  if !m.ExcludeCDNCGI.IsNull() {
    params.ExcludeCDNCGI = cloudflare.F(m.ExcludeCDNCGI.ValueBool())
  }
  if !m.ExcludeDuplicates.IsNull() {
    params.ExcludeDuplicates = cloudflare.F(m.ExcludeDuplicates.ValueBool())
  }
  if !m.ExcludeURLs.IsNull() {
    params.ExcludeURLs = cloudflare.F(m.ExcludeURLs.ValueString())
  }
  if !m.Export.IsNull() {
    params.Export = cloudflare.F(page_shield.ScriptListParamsExport(m.Export.ValueString()))
  }
  if !m.Hosts.IsNull() {
    params.Hosts = cloudflare.F(m.Hosts.ValueString())
  }
  if !m.OrderBy.IsNull() {
    params.OrderBy = cloudflare.F(page_shield.ScriptListParamsOrderBy(m.OrderBy.ValueString()))
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

type PageShieldScriptsListResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AddedAt timetypes.RFC3339 `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
FirstSeenAt timetypes.RFC3339 `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
Host types.String `tfsdk:"host" json:"host,computed"`
LastSeenAt timetypes.RFC3339 `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
URL types.String `tfsdk:"url" json:"url,computed"`
URLContainsCDNCGIPath types.Bool `tfsdk:"url_contains_cdn_cgi_path" json:"url_contains_cdn_cgi_path,computed"`
CryptominingScore types.Int64 `tfsdk:"cryptomining_score" json:"cryptomining_score,computed"`
DataflowScore types.Int64 `tfsdk:"dataflow_score" json:"dataflow_score,computed"`
DomainReportedMalicious types.Bool `tfsdk:"domain_reported_malicious" json:"domain_reported_malicious,computed"`
FetchedAt types.String `tfsdk:"fetched_at" json:"fetched_at,computed"`
FirstPageURL types.String `tfsdk:"first_page_url" json:"first_page_url,computed"`
Hash types.String `tfsdk:"hash" json:"hash,computed"`
JSIntegrityScore types.Int64 `tfsdk:"js_integrity_score" json:"js_integrity_score,computed"`
MagecartScore types.Int64 `tfsdk:"magecart_score" json:"magecart_score,computed"`
MaliciousDomainCategories customfield.List[types.String] `tfsdk:"malicious_domain_categories" json:"malicious_domain_categories,computed"`
MaliciousURLCategories customfield.List[types.String] `tfsdk:"malicious_url_categories" json:"malicious_url_categories,computed"`
MalwareScore types.Int64 `tfsdk:"malware_score" json:"malware_score,computed"`
ObfuscationScore types.Int64 `tfsdk:"obfuscation_score" json:"obfuscation_score,computed"`
PageURLs customfield.List[types.String] `tfsdk:"page_urls" json:"page_urls,computed"`
URLReportedMalicious types.Bool `tfsdk:"url_reported_malicious" json:"url_reported_malicious,computed"`
}
