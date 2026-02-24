package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareZoneSchema returns the legacy cloudflare_zone schema (schema_version=0).
// This is used by UpgradeFromLegacy to parse state from the legacy SDKv2 provider.
// Reference: https://github.com/cloudflare/terraform-provider-cloudflare/blob/v4/internal/sdkv2provider/schema_cloudflare_zone.go
func SourceCloudflareZoneSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			// Renamed to name in v5
			"zone": schema.StringAttribute{
				Required: true,
			},
			// Removed in v5; no equivalent field
			"jump_start": schema.BoolAttribute{
				Optional: true,
			},
			"paused": schema.BoolAttribute{
				Optional: true,
			},
			"vanity_name_servers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			// Changed to computed-only SingleNestedAttribute in v5; dropped from state
			"plan": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			// TypeMap[Bool] in v4; incompatible with SingleNestedAttribute in v5; dropped from state
			"meta": schema.MapAttribute{
				ElementType: types.BoolType,
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},
			"name_servers": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"verification_key": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
