package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareOriginCACertificateSchema returns the source schema for legacy cloudflare_origin_ca_certificate resource.
// Schema version: 0 (SDKv2 default - no explicit SchemaVersion set in v4)
// Resource type: cloudflare_origin_ca_certificate
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_origin_ca_certificate.go
func SourceCloudflareOriginCACertificateSchema() schema.Schema {
	return schema.Schema{
		// Version 0 is implicit for SDKv2 resources without explicit SchemaVersion
		// We omit the Version field to use the default value of 0
		Attributes: map[string]schema.Attribute{
			// Resource identifier (implicit in SDKv2 but present in state)
			"id": schema.StringAttribute{
				Computed: true,
			},

			// Required user-provided fields
			"csr": schema.StringAttribute{
				Required: true,
			},
			"request_type": schema.StringAttribute{
				Required: true,
			},
			"hostnames": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
			},

			// Optional user-provided fields
			"requested_validity": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"min_days_for_renewal": schema.Int64Attribute{
				Optional: true,
			},

			// Computed fields (API-provided, read-only)
			"certificate": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
