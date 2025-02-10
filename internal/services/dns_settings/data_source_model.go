// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSSettingsResultDataSourceEnvelope struct {
	Result DNSSettingsDataSourceModel `json:"result,computed"`
}

type DNSSettingsDataSourceModel struct {
	AccountID    types.String                                                     `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID       types.String                                                     `tfsdk:"zone_id" path:"zone_id,optional"`
	ZoneDefaults customfield.NestedObject[DNSSettingsZoneDefaultsDataSourceModel] `tfsdk:"zone_defaults" json:"zone_defaults,computed"`
}

func (m *DNSSettingsDataSourceModel) toReadParams(_ context.Context) (params dns.SettingGetParams, diags diag.Diagnostics) {
	params = dns.SettingGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type DNSSettingsZoneDefaultsDataSourceModel struct {
	FlattenAllCNAMEs   types.Bool                                                                  `tfsdk:"flatten_all_cnames" json:"flatten_all_cnames,computed"`
	FoundationDNS      types.Bool                                                                  `tfsdk:"foundation_dns" json:"foundation_dns,computed"`
	InternalDNS        customfield.NestedObject[DNSSettingsZoneDefaultsInternalDNSDataSourceModel] `tfsdk:"internal_dns" json:"internal_dns,computed"`
	MultiProvider      types.Bool                                                                  `tfsdk:"multi_provider" json:"multi_provider,computed"`
	Nameservers        customfield.NestedObject[DNSSettingsZoneDefaultsNameserversDataSourceModel] `tfsdk:"nameservers" json:"nameservers,computed"`
	NSTTL              types.Float64                                                               `tfsdk:"ns_ttl" json:"ns_ttl,computed"`
	SecondaryOverrides types.Bool                                                                  `tfsdk:"secondary_overrides" json:"secondary_overrides,computed"`
	SOA                customfield.NestedObject[DNSSettingsZoneDefaultsSOADataSourceModel]         `tfsdk:"soa" json:"soa,computed"`
	ZoneMode           types.String                                                                `tfsdk:"zone_mode" json:"zone_mode,computed"`
}

type DNSSettingsZoneDefaultsInternalDNSDataSourceModel struct {
	ReferenceZoneID types.String `tfsdk:"reference_zone_id" json:"reference_zone_id,computed"`
}

type DNSSettingsZoneDefaultsNameserversDataSourceModel struct {
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type DNSSettingsZoneDefaultsSOADataSourceModel struct {
	Expire  types.Float64 `tfsdk:"expire" json:"expire,computed"`
	MinTTL  types.Float64 `tfsdk:"min_ttl" json:"min_ttl,computed"`
	MNAME   types.String  `tfsdk:"mname" json:"mname,computed"`
	Refresh types.Float64 `tfsdk:"refresh" json:"refresh,computed"`
	Retry   types.Float64 `tfsdk:"retry" json:"retry,computed"`
	RNAME   types.String  `tfsdk:"rname" json:"rname,computed"`
	TTL     types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
}
