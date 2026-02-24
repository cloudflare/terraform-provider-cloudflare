package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareWorkersKVSchema returns the source schema for legacy resource.
// Schema version: 0 (SDKv2 default - no explicit SchemaVersion means 0)
// Resource type: cloudflare_workers_kv
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareWorkersKVSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 resources without explicit SchemaVersion default to 0
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"namespace_id": schema.StringAttribute{
				Required: true,
			},
			"key": schema.StringAttribute{
				Required: true,
			},
			"value": schema.StringAttribute{
				Required: true,
			},
		},
	}
}
