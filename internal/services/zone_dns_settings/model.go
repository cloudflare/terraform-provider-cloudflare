// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dns_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneDNSSettingsResultEnvelope struct {
	Result ZoneDNSSettingsModel `json:"result"`
}

type ZoneDNSSettingsModel struct {
	ZoneID             types.String                     `tfsdk:"zone_id" path:"zone_id,required"`
	FlattenAllCNAMEs   types.Bool                       `tfsdk:"flatten_all_cnames" json:"flatten_all_cnames,optional"`
	FoundationDNS      types.Bool                       `tfsdk:"foundation_dns" json:"foundation_dns,optional"`
	MultiProvider      types.Bool                       `tfsdk:"multi_provider" json:"multi_provider,optional"`
	NSTTL              types.Float64                    `tfsdk:"ns_ttl" json:"ns_ttl,optional"`
	SecondaryOverrides types.Bool                       `tfsdk:"secondary_overrides" json:"secondary_overrides,optional"`
	ZoneMode           types.String                     `tfsdk:"zone_mode" json:"zone_mode,optional"`
	InternalDNS        *ZoneDNSSettingsInternalDNSModel `tfsdk:"internal_dns" json:"internal_dns,optional"`
	Nameservers        *ZoneDNSSettingsNameserversModel `tfsdk:"nameservers" json:"nameservers,optional"`
	SOA                *ZoneDNSSettingsSOAModel         `tfsdk:"soa" json:"soa,optional"`
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
	NSSet types.Int64  `tfsdk:"ns_set" json:"ns_set,optional"`
	Type  types.String `tfsdk:"type" json:"type,optional"`
}

type ZoneDNSSettingsSOAModel struct {
	Expire  types.Float64 `tfsdk:"expire" json:"expire,optional"`
	MinTTL  types.Float64 `tfsdk:"min_ttl" json:"min_ttl,optional"`
	MNAME   types.String  `tfsdk:"mname" json:"mname,optional"`
	Refresh types.Float64 `tfsdk:"refresh" json:"refresh,optional"`
	Retry   types.Float64 `tfsdk:"retry" json:"retry,optional"`
	RNAME   types.String  `tfsdk:"rname" json:"rname,optional"`
	TTL     types.Float64 `tfsdk:"ttl" json:"ttl,optional"`
}
