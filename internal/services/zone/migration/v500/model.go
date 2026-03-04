package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareZoneModel represents the legacy cloudflare_zone state from v4.x provider.
// Schema version: 0 (SDKv2 provider; no explicit version set)
// Resource type: cloudflare_zone
type SourceCloudflareZoneModel struct {
	ID                types.String `tfsdk:"id"`
	AccountID         types.String `tfsdk:"account_id"`
	Zone              types.String `tfsdk:"zone"`              // Renamed to name in v5
	JumpStart         types.Bool   `tfsdk:"jump_start"`        // Removed in v5
	Paused            types.Bool   `tfsdk:"paused"`
	VanityNameServers types.List   `tfsdk:"vanity_name_servers"`
	Plan              types.String `tfsdk:"plan"`              // Changed to computed nested object in v5
	Meta              types.Map    `tfsdk:"meta"`              // Incompatible type in v5; drop from state
	Status            types.String `tfsdk:"status"`
	Type              types.String `tfsdk:"type"`
	NameServers       types.List   `tfsdk:"name_servers"`
	VerificationKey   types.String `tfsdk:"verification_key"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetZoneModel represents the current cloudflare_zone state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_zone
//
// Note: Duplicates ZoneModel from the parent package to keep the migration
// package self-contained.
type TargetZoneModel struct {
	ID                  types.String                                         `tfsdk:"id"`
	Name                types.String                                         `tfsdk:"name"`
	Account             *TargetZoneAccountModel                              `tfsdk:"account"`
	Paused              types.Bool                                           `tfsdk:"paused"`
	Type                types.String                                         `tfsdk:"type"`
	VanityNameServers   customfield.List[types.String]                       `tfsdk:"vanity_name_servers"`
	ActivatedOn         timetypes.RFC3339                                    `tfsdk:"activated_on"`
	CNAMESuffix         types.String                                         `tfsdk:"cname_suffix"`
	CreatedOn           timetypes.RFC3339                                    `tfsdk:"created_on"`
	DevelopmentMode     types.Float64                                        `tfsdk:"development_mode"`
	ModifiedOn          timetypes.RFC3339                                    `tfsdk:"modified_on"`
	OriginalDnshost     types.String                                         `tfsdk:"original_dnshost"`
	OriginalRegistrar   types.String                                         `tfsdk:"original_registrar"`
	Status              types.String                                         `tfsdk:"status"`
	VerificationKey     types.String                                         `tfsdk:"verification_key"`
	NameServers         customfield.List[types.String]                       `tfsdk:"name_servers"`
	OriginalNameServers customfield.List[types.String]                       `tfsdk:"original_name_servers"`
	Permissions         customfield.List[types.String]                       `tfsdk:"permissions"`
	Meta                customfield.NestedObject[TargetZoneMetaModel]        `tfsdk:"meta"`
	Owner               customfield.NestedObject[TargetZoneOwnerModel]       `tfsdk:"owner"`
	Plan                customfield.NestedObject[TargetZonePlanModel]        `tfsdk:"plan"`
	Tenant              customfield.NestedObject[TargetZoneTenantModel]      `tfsdk:"tenant"`
	TenantUnit          customfield.NestedObject[TargetZoneTenantUnitModel]  `tfsdk:"tenant_unit"`
}

// TargetZoneAccountModel represents the nested account object in v5.
type TargetZoneAccountModel struct {
	ID types.String `tfsdk:"id"`
}

// TargetZoneMetaModel represents the nested meta object in v5.
type TargetZoneMetaModel struct {
	CDNOnly                types.Bool  `tfsdk:"cdn_only"`
	CustomCertificateQuota types.Int64 `tfsdk:"custom_certificate_quota"`
	DNSOnly                types.Bool  `tfsdk:"dns_only"`
	FoundationDNS          types.Bool  `tfsdk:"foundation_dns"`
	PageRuleQuota          types.Int64 `tfsdk:"page_rule_quota"`
	PhishingDetected       types.Bool  `tfsdk:"phishing_detected"`
	Step                   types.Int64 `tfsdk:"step"`
}

// TargetZoneOwnerModel represents the nested owner object in v5.
type TargetZoneOwnerModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// TargetZonePlanModel represents the nested plan object in v5.
type TargetZonePlanModel struct {
	ID                types.String  `tfsdk:"id"`
	CanSubscribe      types.Bool    `tfsdk:"can_subscribe"`
	Currency          types.String  `tfsdk:"currency"`
	ExternallyManaged types.Bool    `tfsdk:"externally_managed"`
	Frequency         types.String  `tfsdk:"frequency"`
	IsSubscribed      types.Bool    `tfsdk:"is_subscribed"`
	LegacyDiscount    types.Bool    `tfsdk:"legacy_discount"`
	LegacyID          types.String  `tfsdk:"legacy_id"`
	Name              types.String  `tfsdk:"name"`
	Price             types.Float64 `tfsdk:"price"`
}

// TargetZoneTenantModel represents the nested tenant object in v5.
type TargetZoneTenantModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// TargetZoneTenantUnitModel represents the nested tenant_unit object in v5.
type TargetZoneTenantUnitModel struct {
	ID types.String `tfsdk:"id"`
}
