package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareMTLSCertificateSchema returns the legacy cloudflare_mtls_certificate schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_mtls_certificate.go
//
// Note: This minimal schema is used only for reading v4 state during migration.
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareMTLSCertificateSchema() schema.Schema {
	return schema.Schema{
		// CRITICAL: Must match actual v4 schema version.
		// SDKv2 resources default to schema_version=0 when SchemaVersion is not set.
		// cloudflare_mtls_certificate in v4 has no SchemaVersion set → version 0.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"ca": schema.BoolAttribute{
				Required: true,
			},
			"certificates": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"private_key": schema.StringAttribute{
				Optional: true,
			},
			// Computed fields from the v4 provider
			"issuer": schema.StringAttribute{
				Computed: true,
			},
			"signature": schema.StringAttribute{
				Computed: true,
			},
			"serial_number": schema.StringAttribute{
				Computed: true,
			},
			"uploaded_on": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
