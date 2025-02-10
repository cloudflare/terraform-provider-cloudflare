// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSSettingsResultEnvelope struct {
	Result DNSSettingsModel `json:"result"`
}

type DNSSettingsModel struct {
	AccountID    types.String                                           `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID       types.String                                           `tfsdk:"zone_id" path:"zone_id,optional"`
	ZoneDefaults customfield.NestedObject[DNSSettingsZoneDefaultsModel] `tfsdk:"zone_defaults" json:"zone_defaults,computed_optional"`
}

func (m DNSSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSSettingsModel) MarshalJSONForUpdate(state DNSSettingsModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type DNSSettingsZoneDefaultsModel struct {
	FlattenAllCNAMEs   types.Bool                                                        `tfsdk:"flatten_all_cnames" json:"flatten_all_cnames,optional"`
	FoundationDNS      types.Bool                                                        `tfsdk:"foundation_dns" json:"foundation_dns,optional"`
	InternalDNS        customfield.NestedObject[DNSSettingsZoneDefaultsInternalDNSModel] `tfsdk:"internal_dns" json:"internal_dns,computed_optional"`
	MultiProvider      types.Bool                                                        `tfsdk:"multi_provider" json:"multi_provider,optional"`
	Nameservers        customfield.NestedObject[DNSSettingsZoneDefaultsNameserversModel] `tfsdk:"nameservers" json:"nameservers,computed_optional"`
	NSTTL              types.Float64                                                     `tfsdk:"ns_ttl" json:"ns_ttl,optional"`
	SecondaryOverrides types.Bool                                                        `tfsdk:"secondary_overrides" json:"secondary_overrides,optional"`
	SOA                customfield.NestedObject[DNSSettingsZoneDefaultsSOAModel]         `tfsdk:"soa" json:"soa,computed_optional"`
	ZoneMode           types.String                                                      `tfsdk:"zone_mode" json:"zone_mode,optional"`
}

type DNSSettingsZoneDefaultsInternalDNSModel struct {
	ReferenceZoneID types.String `tfsdk:"reference_zone_id" json:"reference_zone_id,optional"`
}

type DNSSettingsZoneDefaultsNameserversModel struct {
	Type types.String `tfsdk:"type" json:"type,required"`
}

type DNSSettingsZoneDefaultsSOAModel struct {
	Expire  types.Float64 `tfsdk:"expire" json:"expire,required"`
	MinTTL  types.Float64 `tfsdk:"min_ttl" json:"min_ttl,required"`
	MNAME   types.String  `tfsdk:"mname" json:"mname,required"`
	Refresh types.Float64 `tfsdk:"refresh" json:"refresh,required"`
	Retry   types.Float64 `tfsdk:"retry" json:"retry,required"`
	RNAME   types.String  `tfsdk:"rname" json:"rname,required"`
	TTL     types.Float64 `tfsdk:"ttl" json:"ttl,required"`
}
