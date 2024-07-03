// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultDataSourceEnvelope struct {
	Result ZoneDataSourceModel `json:"result,computed"`
}

type ZoneResultListDataSourceEnvelope struct {
	Result *[]*ZoneDataSourceModel `json:"result,computed"`
}

type ZoneDataSourceModel struct {
	ZoneID              types.String                  `tfsdk:"zone_id" path:"zone_id"`
	ID                  types.String                  `tfsdk:"id" json:"id"`
	Account             *ZoneAccountDataSourceModel   `tfsdk:"account" json:"account"`
	ActivatedOn         types.String                  `tfsdk:"activated_on" json:"activated_on"`
	CreatedOn           types.String                  `tfsdk:"created_on" json:"created_on"`
	DevelopmentMode     types.Float64                 `tfsdk:"development_mode" json:"development_mode"`
	Meta                *ZoneMetaDataSourceModel      `tfsdk:"meta" json:"meta"`
	ModifiedOn          types.String                  `tfsdk:"modified_on" json:"modified_on"`
	Name                types.String                  `tfsdk:"name" json:"name"`
	NameServers         types.String                  `tfsdk:"name_servers" json:"name_servers"`
	OriginalDnshost     types.String                  `tfsdk:"original_dnshost" json:"original_dnshost"`
	OriginalNameServers types.String                  `tfsdk:"original_name_servers" json:"original_name_servers"`
	OriginalRegistrar   types.String                  `tfsdk:"original_registrar" json:"original_registrar"`
	Owner               *ZoneOwnerDataSourceModel     `tfsdk:"owner" json:"owner"`
	VanityNameServers   types.String                  `tfsdk:"vanity_name_servers" json:"vanity_name_servers"`
	FindOneBy           *ZoneFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type ZoneAccountDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}

type ZoneMetaDataSourceModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected"`
	Step                   types.Int64 `tfsdk:"step" json:"step"`
}

type ZoneOwnerDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}

type ZoneFindOneByDataSourceModel struct {
	Account   *ZoneFindOneByAccountDataSourceModel `tfsdk:"account" query:"account"`
	Direction types.String                         `tfsdk:"direction" query:"direction"`
	Match     types.String                         `tfsdk:"match" query:"match"`
	Name      types.String                         `tfsdk:"name" query:"name"`
	Order     types.String                         `tfsdk:"order" query:"order"`
	Page      types.Float64                        `tfsdk:"page" query:"page"`
	PerPage   types.Float64                        `tfsdk:"per_page" query:"per_page"`
	Status    types.String                         `tfsdk:"status" query:"status"`
}

type ZoneFindOneByAccountDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}
