package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceR2BucketSchema returns the source schema for legacy cloudflare_r2_bucket resource.
// Schema version: 0 (Framework resource without explicit versioning in v4)
// Resource type: cloudflare_r2_bucket
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceR2BucketSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 Framework resource had no explicit version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"location": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
