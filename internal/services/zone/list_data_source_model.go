// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZonesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZonesResultDataSourceModel] `json:"result,computed"`
}

type ZonesDataSourceModel struct {
	Direction types.String                                             `tfsdk:"direction" query:"direction,optional"`
	Name      types.String                                             `tfsdk:"name" query:"name,optional"`
	Order     types.String                                             `tfsdk:"order" query:"order,optional"`
	Status    types.String                                             `tfsdk:"status" query:"status,optional"`
	Account   *ZonesAccountDataSourceModel                             `tfsdk:"account" query:"account,optional"`
	Match     types.String                                             `tfsdk:"match" query:"match,computed_optional"`
	MaxItems  types.Int64                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZonesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZonesDataSourceModel) toListParams(_ context.Context) (params zones.ZoneListParams, diags diag.Diagnostics) {
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
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZonesResultDataSourceModel struct {
	ID                  types.String                                             `tfsdk:"id" json:"id,computed"`
	Account             customfield.NestedObject[ZonesAccountDataSourceModel]    `tfsdk:"account" json:"account,computed"`
	ActivatedOn         timetypes.RFC3339                                        `tfsdk:"activated_on" json:"activated_on,computed" format:"date-time"`
	CreatedOn           timetypes.RFC3339                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DevelopmentMode     types.Float64                                            `tfsdk:"development_mode" json:"development_mode,computed"`
	Meta                customfield.NestedObject[ZonesMetaDataSourceModel]       `tfsdk:"meta" json:"meta,computed"`
	ModifiedOn          timetypes.RFC3339                                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name                types.String                                             `tfsdk:"name" json:"name,computed"`
	NameServers         customfield.List[types.String]                           `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalDnshost     types.String                                             `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalNameServers customfield.List[types.String]                           `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	OriginalRegistrar   types.String                                             `tfsdk:"original_registrar" json:"original_registrar,computed"`
	Owner               customfield.NestedObject[ZonesOwnerDataSourceModel]      `tfsdk:"owner" json:"owner,computed"`
	Plan                customfield.NestedObject[ZonesPlanDataSourceModel]       `tfsdk:"plan" json:"plan,computed"`
	CNAMESuffix         types.String                                             `tfsdk:"cname_suffix" json:"cname_suffix,computed"`
	Paused              types.Bool                                               `tfsdk:"paused" json:"paused,computed"`
	Permissions         customfield.List[types.String]                           `tfsdk:"permissions" json:"permissions,computed"`
	Status              types.String                                             `tfsdk:"status" json:"status,computed"`
	Tenant              customfield.NestedObject[ZonesTenantDataSourceModel]     `tfsdk:"tenant" json:"tenant,computed"`
	TenantUnit          customfield.NestedObject[ZonesTenantUnitDataSourceModel] `tfsdk:"tenant_unit" json:"tenant_unit,computed"`
	Type                types.String                                             `tfsdk:"type" json:"type,computed"`
	VanityNameServers   customfield.List[types.String]                           `tfsdk:"vanity_name_servers" json:"vanity_name_servers,computed"`
	VerificationKey     types.String                                             `tfsdk:"verification_key" json:"verification_key,computed"`
}

type ZonesMetaDataSourceModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only,computed"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota,computed"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only,computed"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns,computed"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota,computed"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected,computed"`
	Step                   types.Int64 `tfsdk:"step" json:"step,computed"`
}

type ZonesOwnerDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZonesPlanDataSourceModel struct {
	ID                types.String  `tfsdk:"id" json:"id,computed"`
	CanSubscribe      types.Bool    `tfsdk:"can_subscribe" json:"can_subscribe,computed"`
	Currency          types.String  `tfsdk:"currency" json:"currency,computed"`
	ExternallyManaged types.Bool    `tfsdk:"externally_managed" json:"externally_managed,computed"`
	Frequency         types.String  `tfsdk:"frequency" json:"frequency,computed"`
	IsSubscribed      types.Bool    `tfsdk:"is_subscribed" json:"is_subscribed,computed"`
	LegacyDiscount    types.Bool    `tfsdk:"legacy_discount" json:"legacy_discount,computed"`
	LegacyID          types.String  `tfsdk:"legacy_id" json:"legacy_id,computed"`
	Name              types.String  `tfsdk:"name" json:"name,computed"`
	Price             types.Float64 `tfsdk:"price" json:"price,computed"`
}

type ZonesTenantDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZonesTenantUnitDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
