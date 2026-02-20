package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceV4AccessRuleSchema returns the source schema for legacy cloudflare_access_rule resource.
// Schema version: 1 (v4 schema version after v0→v1 migration)
// Resource type: cloudflare_access_rule
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Key differences from v5 schema:
// - configuration is ListNestedAttribute (TypeList MaxItems:1 in SDKv2)
// - No created_on, modified_on, allowed_modes, scope fields (new in v5)
func SourceV4AccessRuleSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 1, // CRITICAL: Must match actual v4 schema version (after v0→v1 migration)
		Attributes: map[string]schema.Attribute{
			// Resource identifier (implicit in SDKv2 but present in state)
			"id": schema.StringAttribute{
				Computed: true,
			},

			// Identifiers (ExactlyOneOf in v4, but we don't need validation here)
			"account_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},

			// Required fields
			"mode": schema.StringAttribute{
				Required: true,
			},

			// Nested configuration - stored as LIST in v4 (TypeList MaxItems:1)
			// This is the CRITICAL difference from v5 (which uses SingleNestedAttribute)
			"configuration": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"target": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},

			// Optional fields
			"notes": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}
