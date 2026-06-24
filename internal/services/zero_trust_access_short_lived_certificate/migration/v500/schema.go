package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceAccessCACertificateSchema returns the minimal schema for the legacy cloudflare_access_ca_certificate resource.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_access_ca_certificate
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceAccessCACertificateSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"application_id": schema.StringAttribute{
				Required: true,
			},
			"aud": schema.StringAttribute{
				Computed: true,
			},
			"public_key": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
