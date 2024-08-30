// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultEnvelope struct {
	Result ZoneModel `json:"result"`
}

type ZoneModel struct {
	ID                  types.String                             `tfsdk:"id" json:"id,computed"`
	Name                types.String                             `tfsdk:"name" json:"name"`
	Account             *ZoneAccountModel                        `tfsdk:"account" json:"account"`
	VanityNameServers   *[]types.String                          `tfsdk:"vanity_name_servers" json:"vanity_name_servers"`
	Type                types.String                             `tfsdk:"type" json:"type,computed_optional"`
	ActivatedOn         timetypes.RFC3339                        `tfsdk:"activated_on" json:"activated_on,computed" format:"date-time"`
	CreatedOn           timetypes.RFC3339                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DevelopmentMode     types.Float64                            `tfsdk:"development_mode" json:"development_mode,computed"`
	ModifiedOn          timetypes.RFC3339                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	OriginalDnshost     types.String                             `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalRegistrar   types.String                             `tfsdk:"original_registrar" json:"original_registrar,computed"`
	Paused              types.Bool                               `tfsdk:"paused" json:"paused,computed"`
	Status              types.String                             `tfsdk:"status" json:"status,computed"`
	NameServers         types.List                               `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalNameServers types.List                               `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	Meta                customfield.NestedObject[ZoneMetaModel]  `tfsdk:"meta" json:"meta,computed"`
	Owner               customfield.NestedObject[ZoneOwnerModel] `tfsdk:"owner" json:"owner,computed"`
}

type ZoneAccountModel struct {
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
