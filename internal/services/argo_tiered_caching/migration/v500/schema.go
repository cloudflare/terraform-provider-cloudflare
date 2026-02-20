package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareArgoSchema returns the source schema for legacy cloudflare_argo resource.
// Schema version: 0 (SDK v2 default)
// Resource type: cloudflare_argo
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareArgoSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDK v2 default version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"smart_routing": schema.StringAttribute{
				Optional: true, // "on" or "off"
			},
			"tiered_caching": schema.StringAttribute{
				Optional: true, // "on" or "off"
			},
		},
	}
}
