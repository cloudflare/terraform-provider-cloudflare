// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZonesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZonesResultDataSourceModel] `json:"result,computed"`
}

type ZonesDataSourceModel struct {
	Direction types.String                                             `tfsdk:"direction" query:"direction"`
	Name      types.String                                             `tfsdk:"name" query:"name"`
	Order     types.String                                             `tfsdk:"order" query:"order"`
	Status    types.String                                             `tfsdk:"status" query:"status"`
	Account   *ZonesAccountDataSourceModel                             `tfsdk:"account" query:"account"`
	Match     types.String                                             `tfsdk:"match" query:"match,computed_optional"`
	MaxItems  types.Int64                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZonesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZonesDataSourceModel) toListParams() (params zones.ZoneListParams, diags diag.Diagnostics) {
	params = zones.ZoneListParams{}

	if m.Account != nil {
		paramsAccount := zones.ZoneListParamsAccount{}
		if !m.Account.ID.IsNull() {
			paramsAccount.ID = cloudflare.F(m.Account.ID.ValueString())
		}
		if !m.Account.Name.IsNull() {
			paramsAccount.Name = cloudflare.F(m.Account.Name.ValueString())
		}
		params.Account = cloudflare.F(paramsAccount)
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zones.ZoneListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Match.IsNull() {
		params.Match = cloudflare.F(zones.ZoneListParamsMatch(m.Match.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zones.ZoneListParamsOrder(m.Order.ValueString()))
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(zones.ZoneListParamsStatus(m.Status.ValueString()))
	}

	return
}

type ZonesAccountDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed_optional"`
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
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
	NameServers         types.List                                            `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalDnshost     types.String                                          `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalNameServers types.List                                            `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	OriginalRegistrar   types.String                                          `tfsdk:"original_registrar" json:"original_registrar,computed"`
	Owner               customfield.NestedObject[ZonesOwnerDataSourceModel]   `tfsdk:"owner" json:"owner,computed"`
	VanityNameServers   *[]types.String                                       `tfsdk:"vanity_name_servers" json:"vanity_name_servers,computed_optional"`
}

type ZonesMetaDataSourceModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only,computed_optional"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota,computed_optional"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only,computed_optional"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns,computed_optional"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota,computed_optional"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected,computed_optional"`
	Step                   types.Int64 `tfsdk:"step" json:"step,computed_optional"`
}

type ZonesOwnerDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed_optional"`
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
	Type types.String `tfsdk:"type" json:"type,computed_optional"`
}
