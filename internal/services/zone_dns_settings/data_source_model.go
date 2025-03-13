// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dns_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneDNSSettingsResultDataSourceEnvelope struct {
	Result ZoneDNSSettingsDataSourceModel `json:"result,computed"`
}

type ZoneDNSSettingsDataSourceModel struct {
	ZoneID             types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	FlattenAllCNAMEs   types.Bool                                                          `tfsdk:"flatten_all_cnames" json:"flatten_all_cnames,computed"`
	FoundationDNS      types.Bool                                                          `tfsdk:"foundation_dns" json:"foundation_dns,computed"`
	MultiProvider      types.Bool                                                          `tfsdk:"multi_provider" json:"multi_provider,computed"`
	NSTTL              types.Float64                                                       `tfsdk:"ns_ttl" json:"ns_ttl,computed"`
	SecondaryOverrides types.Bool                                                          `tfsdk:"secondary_overrides" json:"secondary_overrides,computed"`
	ZoneMode           types.String                                                        `tfsdk:"zone_mode" json:"zone_mode,computed"`
	InternalDNS        customfield.NestedObject[ZoneDNSSettingsInternalDNSDataSourceModel] `tfsdk:"internal_dns" json:"internal_dns,computed"`
	Nameservers        customfield.NestedObject[ZoneDNSSettingsNameserversDataSourceModel] `tfsdk:"nameservers" json:"nameservers,computed"`
	SOA                customfield.NestedObject[ZoneDNSSettingsSOADataSourceModel]         `tfsdk:"soa" json:"soa,computed"`
}

func (m *ZoneDNSSettingsDataSourceModel) toReadParams(_ context.Context) (params dns.SettingZoneGetParams, diags diag.Diagnostics) {
	params = dns.SettingZoneGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ZoneDNSSettingsInternalDNSDataSourceModel struct {
	ReferenceZoneID types.String `tfsdk:"reference_zone_id" json:"reference_zone_id,computed"`
}

type ZoneDNSSettingsNameserversDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	NSSet types.Int64  `tfsdk:"ns_set" json:"ns_set,computed"`
}

type ZoneDNSSettingsSOADataSourceModel struct {
	Expire  types.Float64 `tfsdk:"expire" json:"expire,computed"`
	MinTTL  types.Float64 `tfsdk:"min_ttl" json:"min_ttl,computed"`
	MNAME   types.String  `tfsdk:"mname" json:"mname,computed"`
	Refresh types.Float64 `tfsdk:"refresh" json:"refresh,computed"`
	Retry   types.Float64 `tfsdk:"retry" json:"retry,computed"`
	RNAME   types.String  `tfsdk:"rname" json:"rname,computed"`
	TTL     types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
}
