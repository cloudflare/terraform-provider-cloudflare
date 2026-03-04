package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudfareTurnstileWidgetSchema returns the legacy cloudflare_turnstile_widget schema.
// Schema version: 0 (implicit in v4 - no Version field set)
// Resource type: cloudflare_turnstile_widget
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Reference: cloudflare-terraform-v4/internal/framework/service/turnstile/schema.go
func SourceCloudfareTurnstileWidgetSchema() schema.Schema {
	return schema.Schema{
		// Version: 0 (implicit - v4 schema had no explicit version)
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true, // v4 had both Computed and Optional
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"secret": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			// v4: SetAttribute (unordered)
			// v5: ListAttribute (ordered)
			"domains": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"mode": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"bot_fight_mode": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"offlabel": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
