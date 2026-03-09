package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareAuthenticatedOriginPullsSchema returns the source schema for legacy resource.
// Schema version: 0 (default, not explicitly specified in v4)
// Resource type: cloudflare_authenticated_origin_pulls
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// The v4 resource handled three modes:
// 1. Global AOP: zone_id + enabled
// 2. Per-Zone AOP: zone_id + authenticated_origin_pulls_certificate + enabled
// 3. Per-Hostname AOP: zone_id + hostname + authenticated_origin_pulls_certificate + enabled
//
// Only mode 3 (Per-Hostname) resources migrate to this v5 resource.
// Modes 1 and 2 migrate to cloudflare_authenticated_origin_pulls_settings instead.
func SourceCloudflareAuthenticatedOriginPullsSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"hostname": schema.StringAttribute{
				Optional: true,
			},
			"authenticated_origin_pulls_certificate": schema.StringAttribute{
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}
