// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_dns_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountDNSSettingsResultEnvelope struct {
	Result AccountDNSSettingsModel `json:"result"`
}

type AccountDNSSettingsModel struct {
	AccountID    types.String                         `tfsdk:"account_id" path:"account_id,required"`
	ZoneDefaults *AccountDNSSettingsZoneDefaultsModel `tfsdk:"zone_defaults" json:"zone_defaults,optional"`
}

func (m AccountDNSSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountDNSSettingsModel) MarshalJSONForUpdate(state AccountDNSSettingsModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type AccountDNSSettingsZoneDefaultsModel struct {
	FlattenAllCNAMEs   types.Bool                                      `tfsdk:"flatten_all_cnames" json:"flatten_all_cnames,optional"`
	FoundationDNS      types.Bool                                      `tfsdk:"foundation_dns" json:"foundation_dns,optional"`
	InternalDNS        *AccountDNSSettingsZoneDefaultsInternalDNSModel `tfsdk:"internal_dns" json:"internal_dns,optional"`
	MultiProvider      types.Bool                                      `tfsdk:"multi_provider" json:"multi_provider,optional"`
	Nameservers        *AccountDNSSettingsZoneDefaultsNameserversModel `tfsdk:"nameservers" json:"nameservers,optional"`
	NSTTL              types.Float64                                   `tfsdk:"ns_ttl" json:"ns_ttl,optional"`
	SecondaryOverrides types.Bool                                      `tfsdk:"secondary_overrides" json:"secondary_overrides,optional"`
	SOA                *AccountDNSSettingsZoneDefaultsSOAModel         `tfsdk:"soa" json:"soa,optional"`
	ZoneMode           types.String                                    `tfsdk:"zone_mode" json:"zone_mode,optional"`
}

type AccountDNSSettingsZoneDefaultsInternalDNSModel struct {
	ReferenceZoneID types.String `tfsdk:"reference_zone_id" json:"reference_zone_id,optional"`
}

type AccountDNSSettingsZoneDefaultsNameserversModel struct {
	Type types.String `tfsdk:"type" json:"type,required"`
}

type AccountDNSSettingsZoneDefaultsSOAModel struct {
	Expire  types.Float64 `tfsdk:"expire" json:"expire,required"`
	MinTTL  types.Float64 `tfsdk:"min_ttl" json:"min_ttl,required"`
	MNAME   types.String  `tfsdk:"mname" json:"mname,required"`
	Refresh types.Float64 `tfsdk:"refresh" json:"refresh,required"`
	Retry   types.Float64 `tfsdk:"retry" json:"retry,required"`
	RNAME   types.String  `tfsdk:"rname" json:"rname,required"`
	TTL     types.Float64 `tfsdk:"ttl" json:"ttl,required"`
}
