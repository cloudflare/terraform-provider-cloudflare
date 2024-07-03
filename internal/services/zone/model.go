// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultEnvelope struct {
	Result ZoneModel `json:"result,computed"`
}

type ZoneModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	Type                types.String      `tfsdk:"type" json:"type"`
	Account             *ZoneAccountModel `tfsdk:"account" json:"account"`
	Name                types.String      `tfsdk:"name" json:"name"`
	Plan                *ZonePlanModel    `tfsdk:"plan" json:"plan"`
	VanityNameServers   *[]types.String   `tfsdk:"vanity_name_servers" json:"vanity_name_servers"`
	ActivatedOn         types.String      `tfsdk:"activated_on" json:"activated_on,computed"`
	CreatedOn           types.String      `tfsdk:"created_on" json:"created_on,computed"`
	DevelopmentMode     types.Float64     `tfsdk:"development_mode" json:"development_mode,computed"`
	ModifiedOn          types.String      `tfsdk:"modified_on" json:"modified_on,computed"`
	NameServers         *[]types.String   `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalDnshost     types.String      `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalNameServers *[]types.String   `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	OriginalRegistrar   types.String      `tfsdk:"original_registrar" json:"original_registrar,computed"`
}

type ZoneAccountModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZonePlanModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZoneMetaModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected"`
	Step                   types.Int64 `tfsdk:"step" json:"step"`
}

type ZoneOwnerModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}
