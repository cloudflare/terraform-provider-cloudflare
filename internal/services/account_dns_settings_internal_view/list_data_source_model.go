// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_dns_settings_internal_view

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/dns"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountDNSSettingsInternalViewsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[AccountDNSSettingsInternalViewsResultDataSourceModel] `json:"result,computed"`
}

type AccountDNSSettingsInternalViewsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Order types.String `tfsdk:"order" query:"order,optional"`
ZoneID types.String `tfsdk:"zone_id" query:"zone_id,optional"`
ZoneName types.String `tfsdk:"zone_name" query:"zone_name,optional"`
Name *AccountDNSSettingsInternalViewsNameDataSourceModel `tfsdk:"name" query:"name,optional"`
Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
Match types.String `tfsdk:"match" query:"match,computed_optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[AccountDNSSettingsInternalViewsResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountDNSSettingsInternalViewsDataSourceModel) toListParams(_ context.Context) (params dns.SettingAccountViewListParams, diags diag.Diagnostics) {
  params = dns.SettingAccountViewListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  if !m.Direction.IsNull() {
    params.Direction = cloudflare.F(dns.SettingAccountViewListParamsDirection(m.Direction.ValueString()))
  }
  if !m.Match.IsNull() {
    params.Match = cloudflare.F(dns.SettingAccountViewListParamsMatch(m.Match.ValueString()))
  }
  if m.Name != nil {
    paramsName := dns.SettingAccountViewListParamsName{

    }
    if !m.Name.Contains.IsNull() {
      paramsName.Contains = cloudflare.F(m.Name.Contains.ValueString())
    }
    if !m.Name.Endswith.IsNull() {
      paramsName.Endswith = cloudflare.F(m.Name.Endswith.ValueString())
    }
    if !m.Name.Exact.IsNull() {
      paramsName.Exact = cloudflare.F(m.Name.Exact.ValueString())
    }
    if !m.Name.Startswith.IsNull() {
      paramsName.Startswith = cloudflare.F(m.Name.Startswith.ValueString())
    }
    params.Name = cloudflare.F(paramsName)
  }
  if !m.Order.IsNull() {
    params.Order = cloudflare.F(dns.SettingAccountViewListParamsOrder(m.Order.ValueString()))
  }
  if !m.ZoneID.IsNull() {
    params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
  }
  if !m.ZoneName.IsNull() {
    params.ZoneName = cloudflare.F(m.ZoneName.ValueString())
  }

  return
}

type AccountDNSSettingsInternalViewsNameDataSourceModel struct {
Contains types.String `tfsdk:"contains" json:"contains,optional"`
Endswith types.String `tfsdk:"endswith" json:"endswith,optional"`
Exact types.String `tfsdk:"exact" json:"exact,optional"`
Startswith types.String `tfsdk:"startswith" json:"startswith,optional"`
}

type AccountDNSSettingsInternalViewsResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
CreatedTime timetypes.RFC3339 `tfsdk:"created_time" json:"created_time,computed" format:"date-time"`
ModifiedTime timetypes.RFC3339 `tfsdk:"modified_time" json:"modified_time,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
Zones customfield.List[types.String] `tfsdk:"zones" json:"zones,computed"`
}
