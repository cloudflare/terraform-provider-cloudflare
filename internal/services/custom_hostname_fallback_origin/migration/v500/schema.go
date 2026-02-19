package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCustomHostnameFallbackOriginSchema returns the source schema for legacy resource.
// Schema version: 0 (v4 had no explicit version)
// Resource type: cloudflare_custom_hostname_fallback_origin
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Note: This schema is IDENTICAL to v5 except for omitting validators/modifiers.
// No field changes occurred between v4 and v5.
func SourceCustomHostnameFallbackOriginSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 had no explicit version, defaults to 0
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"origin": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"errors": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
		},
	}
}
