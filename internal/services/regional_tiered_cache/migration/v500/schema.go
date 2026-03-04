package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareRegionalTieredCacheSchema returns the legacy cloudflare_regional_tiered_cache schema (schema_version=0).
// This is used by UpgradeFromLegacyV0 to parse state from the legacy SDKv2 provider.
//
// Schema version: 0 (SDKv2 default - no explicit SchemaVersion set in v4 resource)
// Resource type: cloudflare_regional_tiered_cache (NOT renamed in v5)
//
// This minimal schema includes only the properties needed for state parsing:
// - Required, Optional, Computed flags
// - ElementType for collections (if any)
//
// Intentionally omitted (not needed for reading state):
// - Validators
// - PlanModifiers
// - Descriptions/MarkdownDescription
// - Default values
//
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_regional_tiered_cache.go
func SourceCloudflareRegionalTieredCacheSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			// id: Set to zone_id in v4 provider (see resource.go line 42: d.SetId(zoneID))
			"id": schema.StringAttribute{
				Computed: true,
			},
			// zone_id: Required identifier
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			// value: Required in v4, validates "on" or "off"
			// In v5, this becomes Optional+Computed with default "off"
			"value": schema.StringAttribute{
				Required: true,
			},
			// NOTE: v4 does NOT have these fields (they are new in v5):
			// - editable (types.Bool, computed)
			// - modified_on (timetypes.RFC3339, computed)
		},
	}
}
