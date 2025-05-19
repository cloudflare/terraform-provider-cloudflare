// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneResultEnvelope struct {
	Result ZoneModel `json:"result"`
}

type ZoneModel struct {
	ID                  types.String                                  `tfsdk:"id" json:"id,computed"`
	Name                types.String                                  `tfsdk:"name" json:"name,required"`
	Account             *ZoneAccountModel                             `tfsdk:"account" json:"account,required"`
	VanityNameServers   *[]types.String                               `tfsdk:"vanity_name_servers" json:"vanity_name_servers,optional"`
	Type                types.String                                  `tfsdk:"type" json:"type,computed_optional"`
	ActivatedOn         timetypes.RFC3339                             `tfsdk:"activated_on" json:"activated_on,computed" format:"date-time"`
	CNAMESuffix         types.String                                  `tfsdk:"cname_suffix" json:"cname_suffix,computed"`
	CreatedOn           timetypes.RFC3339                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DevelopmentMode     types.Float64                                 `tfsdk:"development_mode" json:"development_mode,computed"`
	ModifiedOn          timetypes.RFC3339                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	OriginalDnshost     types.String                                  `tfsdk:"original_dnshost" json:"original_dnshost,computed"`
	OriginalRegistrar   types.String                                  `tfsdk:"original_registrar" json:"original_registrar,computed"`
	Paused              types.Bool                                    `tfsdk:"paused" json:"paused,computed"`
	Status              types.String                                  `tfsdk:"status" json:"status,computed"`
	VerificationKey     types.String                                  `tfsdk:"verification_key" json:"verification_key,computed"`
	NameServers         customfield.List[types.String]                `tfsdk:"name_servers" json:"name_servers,computed"`
	OriginalNameServers customfield.List[types.String]                `tfsdk:"original_name_servers" json:"original_name_servers,computed"`
	Permissions         customfield.List[types.String]                `tfsdk:"permissions" json:"permissions,computed"`
	Meta                customfield.NestedObject[ZoneMetaModel]       `tfsdk:"meta" json:"meta,computed"`
	Owner               customfield.NestedObject[ZoneOwnerModel]      `tfsdk:"owner" json:"owner,computed"`
	Plan                customfield.NestedObject[ZonePlanModel]       `tfsdk:"plan" json:"plan,computed"`
	Tenant              customfield.NestedObject[ZoneTenantModel]     `tfsdk:"tenant" json:"tenant,computed"`
	TenantUnit          customfield.NestedObject[ZoneTenantUnitModel] `tfsdk:"tenant_unit" json:"tenant_unit,computed"`
}

func (m ZoneModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneModel) MarshalJSONForUpdate(state ZoneModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ZoneAccountModel struct {
	ID types.String `tfsdk:"id" json:"id,optional"`
}

type ZoneMetaModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only" json:"cdn_only,computed"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota" json:"custom_certificate_quota,computed"`
	DNSOnly                types.Bool  `tfsdk:"dns_only" json:"dns_only,computed"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns" json:"foundation_dns,computed"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota" json:"page_rule_quota,computed"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected" json:"phishing_detected,computed"`
	Step                   types.Int64 `tfsdk:"step" json:"step,computed"`
}

type ZoneOwnerModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZonePlanModel struct {
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

type ZoneTenantModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZoneTenantUnitModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
