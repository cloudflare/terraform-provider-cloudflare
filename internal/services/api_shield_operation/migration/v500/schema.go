package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareAPIShieldOperationSchema returns the legacy cloudflare_api_shield_operation schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
//
// This schema is minimal - it includes only the properties needed for state parsing:
// Required, Optional, Computed, and ElementType. Validators, PlanModifiers, and
// Descriptions are intentionally omitted.
//
// Reference: Reconstructed from git history commit ebcb42f1d
// (v4 files were deleted in commit 595774b9b after internal migration)
func SourceCloudflareAPIShieldOperationSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"method": schema.StringAttribute{
				Required: true,
			},
			"host": schema.StringAttribute{
				Required: true,
			},
			"endpoint": schema.StringAttribute{
				Required: true,
			},
		},
	}
}
