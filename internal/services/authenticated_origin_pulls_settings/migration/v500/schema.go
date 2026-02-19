package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareAuthenticatedOriginPullsSchema returns the source schema for legacy resource.
// Schema version: 0 (v4 schema version)
// Resource type: cloudflare_authenticated_origin_pulls
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareAuthenticatedOriginPullsSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 schema version
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
