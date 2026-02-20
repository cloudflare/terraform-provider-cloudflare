package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceAccessMutualTLSCertificateSchema returns the minimal schema for the legacy cloudflare_access_mutual_tls_certificate resource.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_access_mutual_tls_certificate
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceAccessMutualTLSCertificateSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"certificate": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"associated_hostnames": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
			"fingerprint": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
