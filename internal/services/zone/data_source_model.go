// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zones"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultDataSourceEnvelope struct {
Result ZoneDataSourceModel `json:"result,computed"`
}

type ZoneDataSourceModel struct {
ID types.String `tfsdk:"id" path:"zone_id,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,optional"`
ActivatedOn timetypes.RFC3339 `tfsdk:"activated_on" json:"activated_on,computed" format:"date-time"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
DevelopmentMode types.Float64 `tfsdk:"development_mode" json:"development_mode,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
OriginalDnshost types.String `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
OriginalRegistrar types.String `tfsdk:"original_registrar" json:"original_registrar,computed"`
Paused types.Bool `tfsdk:"paused" json:"paused,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
VerificationKey types.String `tfsdk:"verification_key" json:"verification_key,computed"`
NameServers customfield.List[types.String] `tfsdk:"name_servers" json:"name_servers,computed"`
OriginalNameServers customfield.List[types.String] `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
VanityNameServers customfield.List[types.String] `tfsdk:"vanity_name_servers" json:"vanity_name_servers,computed"`
Account customfield.NestedObject[ZoneAccountDataSourceModel] `tfsdk:"account" json:"account,computed"`
Meta customfield.NestedObject[ZoneMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
Owner customfield.NestedObject[ZoneOwnerDataSourceModel] `tfsdk:"owner" json:"owner,computed"`
Filter *ZoneFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZoneDataSourceModel) toReadParams(_ context.Context) (params zones.ZoneGetParams, diags diag.Diagnostics) {
  params = zones.ZoneGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

func (m *ZoneDataSourceModel) toListParams(_ context.Context) (params zones.ZoneListParams, diags diag.Diagnostics) {
  params = zones.ZoneListParams{

  }

  if m.Filter.Account != nil {
    paramsAccount := zones.ZoneListParamsAccount{

    }
    if !m.Filter.Account.ID.IsNull() {
      paramsAccount.ID = cloudflare.F(m.Filter.Account.ID.ValueString())
    }
    if !m.Filter.Account.Name.IsNull() {
      paramsAccount.Name = cloudflare.F(m.Filter.Account.Name.ValueString())
    }
    params.Account = cloudflare.F(paramsAccount)
  }
  if !m.Filter.Direction.IsNull() {
    params.Direction = cloudflare.F(zones.ZoneListParamsDirection(m.Filter.Direction.ValueString()))
  }
  if !m.Filter.Match.IsNull() {
    params.Match = cloudflare.F(zones.ZoneListParamsMatch(m.Filter.Match.ValueString()))
  }
  if !m.Filter.Name.IsNull() {
    params.Name = cloudflare.F(m.Filter.Name.ValueString())
  }
  if !m.Filter.Order.IsNull() {
    params.Order = cloudflare.F(zones.ZoneListParamsOrder(m.Filter.Order.ValueString()))
  }
  if !m.Filter.Status.IsNull() {
    params.Status = cloudflare.F(zones.ZoneListParamsStatus(m.Filter.Status.ValueString()))
  }

  return
}

type ZoneAccountDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZoneMetaDataSourceModel struct {
CDNOnly types.Bool `tfsdk:"cdn_only" json:"cdn_only,computed"`
CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota,computed"`
DNSOnly types.Bool `tfsdk:"dns_only" json:"dns_only,computed"`
FoundationDNS types.Bool `tfsdk:"foundation_dns" json:"foundation_dns,computed"`
PageRuleQuota types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota,computed"`
PhishingDetected types.Bool `tfsdk:"phishing_detected" json:"phishing_detected,computed"`
Step types.Int64 `tfsdk:"step" json:"step,computed"`
}

type ZoneOwnerDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZoneFindOneByDataSourceModel struct {
Account *ZonesAccountDataSourceModel `tfsdk:"account" query:"account,optional"`
Direction types.String `tfsdk:"direction" query:"direction,optional"`
Match types.String `tfsdk:"match" query:"match,computed_optional"`
Name types.String `tfsdk:"name" query:"name,optional"`
Order types.String `tfsdk:"order" query:"order,optional"`
Status types.String `tfsdk:"status" query:"status,optional"`
}
