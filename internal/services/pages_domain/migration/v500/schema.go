package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourcePagesDomainSchema returns the source schema for legacy pages_domain resource.
// Schema version: 0 (SDKv2 default - no explicit version in v4)
// Resource type: cloudflare_pages_domain
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// v4 schema reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_pages_domain.go
func SourcePagesDomainSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			// Required fields
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"domain": schema.StringAttribute{
				Required: true,
			},
			"project_name": schema.StringAttribute{
				Required: true,
			},
			// Computed fields
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
