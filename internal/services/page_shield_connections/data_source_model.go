// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_connections

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/page_shield"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldConnectionsResultDataSourceEnvelope struct {
Result PageShieldConnectionsDataSourceModel `json:"result,computed"`
}

type PageShieldConnectionsDataSourceModel struct {
ConnectionID types.String `tfsdk:"connection_id" path:"connection_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
AddedAt timetypes.RFC3339 `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
DomainReportedMalicious types.Bool `tfsdk:"domain_reported_malicious" json:"domain_reported_malicious,computed"`
FirstPageURL types.String `tfsdk:"first_page_url" json:"first_page_url,computed"`
FirstSeenAt timetypes.RFC3339 `tfsdk:"first_seen_at" json:"first_seen_at,computed" format:"date-time"`
Host types.String `tfsdk:"host" json:"host,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
LastSeenAt timetypes.RFC3339 `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
URL types.String `tfsdk:"url" json:"url,computed"`
URLContainsCDNCGIPath types.Bool `tfsdk:"url_contains_cdn_cgi_path" json:"url_contains_cdn_cgi_path,computed"`
URLReportedMalicious types.Bool `tfsdk:"url_reported_malicious" json:"url_reported_malicious,computed"`
MaliciousDomainCategories customfield.List[types.String] `tfsdk:"malicious_domain_categories" json:"malicious_domain_categories,computed"`
MaliciousURLCategories customfield.List[types.String] `tfsdk:"malicious_url_categories" json:"malicious_url_categories,computed"`
PageURLs customfield.List[types.String] `tfsdk:"page_urls" json:"page_urls,computed"`
}

func (m *PageShieldConnectionsDataSourceModel) toReadParams(_ context.Context) (params page_shield.ConnectionGetParams, diags diag.Diagnostics) {
  params = page_shield.ConnectionGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}
