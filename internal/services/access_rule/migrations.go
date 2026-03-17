// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_rule/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func init() {
	// Provide target schema to migration package (avoids circular import)
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithUpgradeState = (*AccessRuleResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles three upgrade paths:
// 1. v5.0.0-v5.12.x state (version=0) -> v5 (version=500): No-op upgrade
// 2. v4 state (schema_version=1) -> v5 (version=500): Full transformation
//   - Unwraps configuration array[0] -> configuration object
//   - Initializes new computed fields
//
// 3. v5.18.0 state (version=1) -> v5 (version=500): No-op upgrade
//   - v5.18.0 was released with dormant state upgrader (GetSchemaVersion=1)
//   - State already has configuration as object, no transformation needed
//
// IMPORTANT: Both v4 and v5.18.0 have schema_version=1, so the version 1
// upgrader must detect the format at runtime by inspecting the raw state.
func (r *AccessRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle v5.0.0-v5.12.x resources (version 0 -> 500)
		// No-op upgrade - schema is already in v5 format
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Handle BOTH v4 (schema_version=1) AND v5.18.0+ (version=1) states
		// PriorSchema is nil because v4 and v5 have incompatible schemas:
		// - v4: configuration is array (ListNestedAttribute)
		// - v5: configuration is object (SingleNestedAttribute)
		// Handler uses RawState to detect format and process accordingly.
		1: {
			PriorSchema:   nil, // Use RawState - schemas are incompatible
			StateUpgrader: v500.UpgradeFromVersion1,
		},
	}
}
