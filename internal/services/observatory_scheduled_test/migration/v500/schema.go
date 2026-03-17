package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareObservatoryScheduledTestSchema returns the legacy cloudflare_observatory_scheduled_test schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
//
// The v4 schema has no explicit SchemaVersion field, so it defaults to 0 (SDKv2 default).
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_observatory_scheduled.go
//
// This is a minimal schema containing only the properties needed for state parsing:
// - Required, Optional, Computed flags
// - No validators, plan modifiers, or descriptions (not needed for reading state)
func SourceCloudflareObservatoryScheduledTestSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			// id is computed by the provider (set to URL value in v4)
			"id": schema.StringAttribute{
				Computed: true,
			},
			// zone_id is required
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			// url is required (used as the test identifier)
			"url": schema.StringAttribute{
				Required: true,
			},
			// region is required in v4, becomes computed+optional in v5
			"region": schema.StringAttribute{
				Required: true,
			},
			// frequency is required in v4, becomes computed+optional in v5
			"frequency": schema.StringAttribute{
				Required: true,
			},
		},
	}
}
