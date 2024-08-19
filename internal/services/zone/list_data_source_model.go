// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZonesResultListDataSourceEnvelope struct {
	Result *[]*ZonesResultDataSourceModel `json:"result,computed"`
}

type ZonesDataSourceModel struct {
	Direction types.String                   `tfsdk:"direction" query:"direction"`
	Name      types.String                   `tfsdk:"name" query:"name"`
	Order     types.String                   `tfsdk:"order" query:"order"`
	Status    types.String                   `tfsdk:"status" query:"status"`
	Account   *ZonesAccountDataSourceModel   `tfsdk:"account" query:"account"`
	Match     types.String                   `tfsdk:"match" query:"match"`
	MaxItems  types.Int64                    `tfsdk:"max_items"`
	Result    *[]*ZonesResultDataSourceModel `tfsdk:"result"`
}

type ZonesAccountDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}

type ZonesResultDataSourceModel struct {
	ID                  types.String                                          `tfsdk:"id" json:"id,computed"`
	Account             customfield.NestedObject[ZonesAccountDataSourceModel] `tfsdk:"account" json:"account,computed"`
	ActivatedOn         timetypes.RFC3339                                     `tfsdk:"activated_on" json:"activated_on,computed"`
	CreatedOn           timetypes.RFC3339                                     `tfsdk:"created_on" json:"created_on,computed"`
	DevelopmentMode     types.Float64                                         `tfsdk:"development_mode" json:"development_mode,computed"`
	Meta                customfield.NestedObject[ZonesMetaDataSourceModel]    `tfsdk:"meta" json:"meta,computed"`
	ModifiedOn          timetypes.RFC3339                                     `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                types.String                                          `tfsdk:"name" json:"name,computed"`
	NameServers         *[]types.String                                       `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalDnshost     types.String                                          `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalNameServers *[]types.String                                       `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	OriginalRegistrar   types.String                                          `tfsdk:"original_registrar" json:"original_registrar,computed"`
	Owner               customfield.NestedObject[ZonesOwnerDataSourceModel]   `tfsdk:"owner" json:"owner,computed"`
	VanityNameServers   *[]types.String                                       `tfsdk:"vanity_name_servers" json:"vanity_name_servers"`
}

type ZonesMetaDataSourceModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected"`
	Step                   types.Int64 `tfsdk:"step" json:"step"`
}

type ZonesOwnerDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}
