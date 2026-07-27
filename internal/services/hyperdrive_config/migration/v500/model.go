package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareHyperdriveConfigModel represents the source cloudflare_hyperdrive_config state structure.
// This corresponds to schema_version=0 from the legacy (framework) cloudflare provider v4.
// Used by UpgradeFromV0 to parse legacy state.
type SourceCloudflareHyperdriveConfigModel struct {
	ID        types.String                              `tfsdk:"id"`
	AccountID types.String                              `tfsdk:"account_id"`
	Name      types.String                              `tfsdk:"name"`
	Origin    *SourceCloudflareHyperdriveConfigOriginModel `tfsdk:"origin"`
	Caching   types.Object                              `tfsdk:"caching"`
}

// SourceCloudflareHyperdriveConfigOriginModel represents the v4 origin nested object.
type SourceCloudflareHyperdriveConfigOriginModel struct {
	Database           types.String `tfsdk:"database"`
	Password           types.String `tfsdk:"password"`
	Host               types.String `tfsdk:"host"`
	Port               types.Int64  `tfsdk:"port"`
	Scheme             types.String `tfsdk:"scheme"`
	User               types.String `tfsdk:"user"`
	AccessClientID     types.String `tfsdk:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret"`
}

// TargetHyperdriveConfigModel represents the target cloudflare_hyperdrive_config state structure (v500).
// Must match the v5 HyperdriveConfigModel structure exactly.
type TargetHyperdriveConfigModel struct {
	ID                    types.String                       `tfsdk:"id"`
	AccountID             types.String                       `tfsdk:"account_id"`
	Name                  types.String                       `tfsdk:"name"`
	Origin                *TargetHyperdriveConfigOriginModel `tfsdk:"origin"`
	OriginConnectionLimit types.Int64                        `tfsdk:"origin_connection_limit"`
	Caching               *TargetHyperdriveConfigCachingModel `tfsdk:"caching"`
	MTLS                  *TargetHyperdriveConfigMTLSModel   `tfsdk:"mtls"`
	CreatedOn             timetypes.RFC3339                  `tfsdk:"created_on"`
	ModifiedOn            timetypes.RFC3339                  `tfsdk:"modified_on"`
}

// TargetHyperdriveConfigOriginModel represents the target origin nested object (v500).
// Must match HyperdriveConfigOriginModel structure exactly.
type TargetHyperdriveConfigOriginModel struct {
	Database           types.String `tfsdk:"database"`
	Host               types.String `tfsdk:"host"`
	Password           types.String `tfsdk:"password"`
	Port               types.Int64  `tfsdk:"port"`
	Scheme             types.String `tfsdk:"scheme"`
	User               types.String `tfsdk:"user"`
	AccessClientID     types.String `tfsdk:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret"`
	ServiceID          types.String `tfsdk:"service_id"`
}

// TargetHyperdriveConfigCachingModel represents the target caching nested object (v500).
// Must match HyperdriveConfigCachingModel structure exactly.
type TargetHyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate"`
}

// TargetHyperdriveConfigMTLSModel represents the target mtls nested object (v500).
// Must match HyperdriveConfigMTLSModel structure exactly.
type TargetHyperdriveConfigMTLSModel struct {
	CACertificateID   types.String `tfsdk:"ca_certificate_id"`
	MTLSCertificateID types.String `tfsdk:"mtls_certificate_id"`
	Sslmode           types.String `tfsdk:"sslmode"`
}
