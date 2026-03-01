package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareCustomPagesSchema returns the legacy cloudflare_custom_pages
// schema from v4.x SDKv2 provider.
//
// Schema version: 0 (v4 SDKv2 resources without explicit SchemaVersion field default to 0)
// Resource type: cloudflare_custom_pages
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_custom_pages.go
func SourceCloudflareCustomPagesSchema() schema.Schema {
	return schema.Schema{
		// Note: No explicit Version field here.
		// v4 SDKv2 resources without SchemaVersion default to 0.
		// This is registered in migrations.go as: 0: { PriorSchema: ... }
		Attributes: map[string]schema.Attribute{
			// id field (Computed in both v4 and v5)
			"id": schema.StringAttribute{
				Computed: true,
			},
			// type field (renamed to identifier in v5)
			"type": schema.StringAttribute{
				Required: true,
			},
			// state field (Optional in v4, Required in v5)
			"state": schema.StringAttribute{
				Optional: true,
			},
			// url field (Required in v4, Optional+Computed in v5)
			"url": schema.StringAttribute{
				Required: true,
			},
			// zone_id (Optional, mutually exclusive with account_id)
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			// account_id (Optional, mutually exclusive with zone_id)
			"account_id": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}
