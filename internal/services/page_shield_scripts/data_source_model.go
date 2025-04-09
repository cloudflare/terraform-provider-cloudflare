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

type PageShieldScriptsResultDataSourceEnvelope struct {
Result PageShieldScriptsDataSourceModel `json:"result,computed"`
}

type PageShieldScriptsDataSourceModel struct {
ScriptID types.String `tfsdk:"script_id" path:"script_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
AddedAt timetypes.RFC3339 `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
CryptominingScore types.Int64 `tfsdk:"cryptomining_score" json:"cryptomining_score,computed"`
DataflowScore types.Int64 `tfsdk:"dataflow_score" json:"dataflow_score,computed"`
DomainReportedMalicious types.Bool `tfsdk:"domain_reported_malicious" json:"domain_reported_malicious,computed"`
FetchedAt types.String `tfsdk:"fetched_at" json:"fetched_at,computed"`
FirstPageURL types.String `tfsdk:"first_page_url" json:"first_page_url,computed"`
FirstSeenAt timetypes.RFC3339 `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
Hash types.String `tfsdk:"hash" json:"hash,computed"`
Host types.String `tfsdk:"host" json:"host,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
JSIntegrityScore types.Int64 `tfsdk:"js_integrity_score" json:"js_integrity_score,computed"`
LastSeenAt timetypes.RFC3339 `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
MagecartScore types.Int64 `tfsdk:"magecart_score" json:"magecart_score,computed"`
MalwareScore types.Int64 `tfsdk:"malware_score" json:"malware_score,computed"`
ObfuscationScore types.Int64 `tfsdk:"obfuscation_score" json:"obfuscation_score,computed"`
URL types.String `tfsdk:"url" json:"url,computed"`
URLContainsCDNCGIPath types.Bool `tfsdk:"url_contains_cdn_cgi_path" json:"url_contains_cdn_cgi_path,computed"`
URLReportedMalicious types.Bool `tfsdk:"url_reported_malicious" json:"url_reported_malicious,computed"`
MaliciousDomainCategories customfield.List[types.String] `tfsdk:"malicious_domain_categories" json:"malicious_domain_categories,computed"`
MaliciousURLCategories customfield.List[types.String] `tfsdk:"malicious_url_categories" json:"malicious_url_categories,computed"`
PageURLs customfield.List[types.String] `tfsdk:"page_urls" json:"page_urls,computed"`
Versions customfield.NestedObjectList[PageShieldScriptsVersionsDataSourceModel] `tfsdk:"versions" json:"versions,computed"`
}

func (m *PageShieldScriptsDataSourceModel) toReadParams(_ context.Context) (params page_shield.ScriptGetParams, diags diag.Diagnostics) {
  params = page_shield.ScriptGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type PageShieldScriptsVersionsDataSourceModel struct {
CryptominingScore types.Int64 `tfsdk:"cryptomining_score" json:"cryptomining_score,computed"`
DataflowScore types.Int64 `tfsdk:"dataflow_score" json:"dataflow_score,computed"`
FetchedAt types.String `tfsdk:"fetched_at" json:"fetched_at,computed"`
Hash types.String `tfsdk:"hash" json:"hash,computed"`
JSIntegrityScore types.Int64 `tfsdk:"js_integrity_score" json:"js_integrity_score,computed"`
MagecartScore types.Int64 `tfsdk:"magecart_score" json:"magecart_score,computed"`
MalwareScore types.Int64 `tfsdk:"malware_score" json:"malware_score,computed"`
ObfuscationScore types.Int64 `tfsdk:"obfuscation_score" json:"obfuscation_score,computed"`
}
