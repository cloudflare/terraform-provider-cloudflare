package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareCustomSSLModel represents the legacy resource state from v4.x provider.
// Schema version: 1 (the final schema version after the v0→v1 upgrade in the SDKv2 provider).
// Resource type: cloudflare_custom_ssl
//
// In v4, the custom_ssl_options fields were nested inside a TypeList MaxItems:1 block.
// By the time we receive this state for upgrading, the v4 provider has already
// run its own v0→v1 upgrader (converting custom_ssl_options from TypeMap to TypeList),
// so we only need to model the v1 (TypeList) shape here.
type SourceCloudflareCustomSSLModel struct {
	ID                types.String                   `tfsdk:"id"`
	ZoneID            types.String                   `tfsdk:"zone_id"`
	CustomSSLOptions  []SourceCustomSSLOptionsModel  `tfsdk:"custom_ssl_options"`
	CustomSSLPriority []SourceCustomSSLPriorityModel `tfsdk:"custom_ssl_priority"`
	Hosts             types.List                     `tfsdk:"hosts"`
	Issuer            types.String                   `tfsdk:"issuer"`
	Signature         types.String                   `tfsdk:"signature"`
	Status            types.String                   `tfsdk:"status"`
	UploadedOn        types.String                   `tfsdk:"uploaded_on"`
	ModifiedOn        types.String                   `tfsdk:"modified_on"`
	ExpiresOn         types.String                   `tfsdk:"expires_on"`
	Priority          types.Int64                    `tfsdk:"priority"`
}

// SourceCustomSSLOptionsModel represents the custom_ssl_options block from v4.
// This is a TypeList MaxItems:1 block containing the certificate fields.
type SourceCustomSSLOptionsModel struct {
	Certificate     types.String `tfsdk:"certificate"`
	PrivateKey      types.String `tfsdk:"private_key"`
	BundleMethod    types.String `tfsdk:"bundle_method"`
	GeoRestrictions types.String `tfsdk:"geo_restrictions"`
	Type            types.String `tfsdk:"type"`
}

// SourceCustomSSLPriorityModel represents the custom_ssl_priority block from v4.
// This is a write-only reprioritization block that does not exist in v5.
type SourceCustomSSLPriorityModel struct {
	ID       types.String `tfsdk:"id"`
	Priority types.Int64  `tfsdk:"priority"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetCustomSSLModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_custom_ssl
//
// Note: This duplicates the model in the parent package (custom_ssl.CustomSSLModel)
// to keep the migration package self-contained.
type TargetCustomSSLModel struct {
	ID                 types.String                                                `tfsdk:"id"`
	ZoneID             types.String                                                `tfsdk:"zone_id"`
	Type               types.String                                                `tfsdk:"type"`
	Certificate        types.String                                                `tfsdk:"certificate"`
	PrivateKey         types.String                                                `tfsdk:"private_key"`
	CustomCsrID        types.String                                                `tfsdk:"custom_csr_id"`
	Policy             types.String                                                `tfsdk:"policy"`
	GeoRestrictions    *TargetCustomSSLGeoRestrictionsModel                        `tfsdk:"geo_restrictions"`
	BundleMethod       types.String                                                `tfsdk:"bundle_method"`
	Deploy             types.String                                                `tfsdk:"deploy"`
	ExpiresOn          timetypes.RFC3339                                           `tfsdk:"expires_on"`
	Issuer             types.String                                                `tfsdk:"issuer"`
	ModifiedOn         timetypes.RFC3339                                           `tfsdk:"modified_on"`
	PolicyRestrictions types.String                                                `tfsdk:"policy_restrictions"`
	Priority           types.Float64                                               `tfsdk:"priority"`
	Signature          types.String                                                `tfsdk:"signature"`
	Status             types.String                                                `tfsdk:"status"`
	UploadedOn         timetypes.RFC3339                                           `tfsdk:"uploaded_on"`
	Hosts              customfield.List[types.String]                              `tfsdk:"hosts"`
	KeylessServer      customfield.NestedObject[TargetCustomSSLKeylessServerModel] `tfsdk:"keyless_server"`
}

// TargetCustomSSLGeoRestrictionsModel represents the geo_restrictions nested object in v5.
// In v4, geo_restrictions was a plain TypeString. In v5, it is a SingleNestedAttribute.
type TargetCustomSSLGeoRestrictionsModel struct {
	Label types.String `tfsdk:"label"`
}

// TargetCustomSSLKeylessServerModel represents the keyless_server computed field in v5.
// This field did not exist in v4 and will be populated from the API after migration.
type TargetCustomSSLKeylessServerModel struct {
	ID          types.String                                                      `tfsdk:"id"`
	CreatedOn   timetypes.RFC3339                                                 `tfsdk:"created_on"`
	Enabled     types.Bool                                                        `tfsdk:"enabled"`
	Host        types.String                                                      `tfsdk:"host"`
	ModifiedOn  timetypes.RFC3339                                                 `tfsdk:"modified_on"`
	Name        types.String                                                      `tfsdk:"name"`
	Permissions customfield.List[types.String]                                    `tfsdk:"permissions"`
	Port        types.Float64                                                     `tfsdk:"port"`
	Status      types.String                                                      `tfsdk:"status"`
	Tunnel      customfield.NestedObject[TargetCustomSSLKeylessServerTunnelModel] `tfsdk:"tunnel"`
}

// TargetCustomSSLKeylessServerTunnelModel represents the keyless_server.tunnel nested object.
type TargetCustomSSLKeylessServerTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip"`
	VnetID    types.String `tfsdk:"vnet_id"`
}
