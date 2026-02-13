package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareBotManagementSchema returns the source schema for legacy cloudflare_bot_management resource.
// Schema version: 0 (SDKv2 provider - v4.x)
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareBotManagementSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"ai_bots_protection": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"auto_update_model": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"enable_js": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"fight_mode": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"optimize_wordpress": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"sbfm_definitely_automated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sbfm_likely_automated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sbfm_static_resource_protection": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"sbfm_verified_bots": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"suppress_session_score": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"using_latest_model": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}
