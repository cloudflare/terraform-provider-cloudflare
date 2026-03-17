// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func init() {
	// Provide target schema to migration package (avoids circular import)
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithUpgradeState = (*PagesProjectResource)(nil)

// UpgradeState returns state upgraders for handling schema version migrations.
//
// This handles three upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
//   - build_config, source stored as ARRAYS (ListNestedAttribute) → OBJECTS (SingleNestedAttribute)
//
// 2. v5.16.0 state (schema_version=0) → v5 (version=500): No-op upgrade
//   - v5.16.0 was released with dormant state upgrader (no GetSchemaVersion)
//   - State already has these fields as objects, no transformation needed
//
// 3. v5 state (version=1) → v5 (version=500): No-op upgrade
//
// IMPORTANT: Both v4 and v5.16.0 have schema_version=0, so the version 0
// upgrader must detect the format at runtime by inspecting the raw state.
func (r *PagesProjectResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v5 schema at version=1 for no-op pass-through upgrader
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle BOTH v4 (schema_version=0) AND v5.16.0 (version=0) states
		// PriorSchema is nil because v4 and v5 have incompatible schemas:
		// - v4: build_config, source as ListNestedAttribute (array)
		// - v5: build_config, source as SingleNestedAttribute (object)
		// Handler uses RawState to detect format and process accordingly.
		0: {
			PriorSchema:   nil, // Use RawState - schemas are incompatible
			StateUpgrader: v500.UpgradeFromVersion0,
		},
		// Handle upgrades within v5 series (schema_version=1+) - no-op
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeNoOp,
		},
	}
}
