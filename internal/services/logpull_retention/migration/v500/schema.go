package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareLogpullRetentionSchema returns the legacy cloudflare_logpull_retention schema.
// This is used by UpgradeState to parse state from the legacy SDKv2 provider (v4.x).
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_logpull_retention.go
//
// This minimal schema includes only the properties needed for state parsing:
// - Required, Optional, Computed flags
// - ElementType for collections
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareLogpullRetentionSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			// Source v4 field: "enabled" (will be renamed to "flag" in v5)
			"enabled": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}
