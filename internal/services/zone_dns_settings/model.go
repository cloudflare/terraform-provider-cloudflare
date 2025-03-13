// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dns_settings

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneDNSSettingsResultEnvelope struct {
Result ZoneDNSSettingsModel `json:"result"`
}

type ZoneDNSSettingsModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
FlattenAllCNAMEs types.Bool `tfsdk:"flatten_all_cnames" json:"flatten_all_cnames,optional"`
FoundationDNS types.Bool `tfsdk:"foundation_dns" json:"foundation_dns,optional"`
MultiProvider types.Bool `tfsdk:"multi_provider" json:"multi_provider,optional"`
NSTTL types.Float64 `tfsdk:"ns_ttl" json:"ns_ttl,optional"`
SecondaryOverrides types.Bool `tfsdk:"secondary_overrides" json:"secondary_overrides,optional"`
ZoneMode types.String `tfsdk:"zone_mode" json:"zone_mode,optional"`
InternalDNS customfield.NestedObject[ZoneDNSSettingsInternalDNSModel] `tfsdk:"internal_dns" json:"internal_dns,computed_optional"`
Nameservers customfield.NestedObject[ZoneDNSSettingsNameserversModel] `tfsdk:"nameservers" json:"nameservers,computed_optional"`
SOA customfield.NestedObject[ZoneDNSSettingsSOAModel] `tfsdk:"soa" json:"soa,computed_optional"`
}

func (m ZoneDNSSettingsModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ZoneDNSSettingsModel) MarshalJSONForUpdate(state ZoneDNSSettingsModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}

type ZoneDNSSettingsInternalDNSModel struct {
ReferenceZoneID types.String `tfsdk:"reference_zone_id" json:"reference_zone_id,optional"`
}

type ZoneDNSSettingsNameserversModel struct {
Type types.String `tfsdk:"type" json:"type,required"`
NSSet types.Int64 `tfsdk:"ns_set" json:"ns_set,optional"`
}

type ZoneDNSSettingsSOAModel struct {
Expire types.Float64 `tfsdk:"expire" json:"expire,required"`
MinTTL types.Float64 `tfsdk:"min_ttl" json:"min_ttl,required"`
MNAME types.String `tfsdk:"mname" json:"mname,required"`
Refresh types.Float64 `tfsdk:"refresh" json:"refresh,required"`
Retry types.Float64 `tfsdk:"retry" json:"retry,required"`
RNAME types.String `tfsdk:"rname" json:"rname,required"`
TTL types.Float64 `tfsdk:"ttl" json:"ttl,required"`
}
