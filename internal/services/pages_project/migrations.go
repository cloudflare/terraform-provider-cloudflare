// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*PagesProjectResource)(nil)

// UpgradeState returns state upgraders for handling schema version migrations.
// Version 0: v4 provider schema (pre-5.x) - blocks stored as lists (SDKv2 style)
// Version 1/500: v5 provider schema - single nested attributes
//
// to safely bump existing v5 users from version 0 to 1 without triggering the
// v4→v5 transformation (which would fail on v5-format state).
func (r *PagesProjectResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourcePagesProjectSchemaV0(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 provider (schema_version=0)
		0: {
			PriorSchema:   sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle upgrades within v5 series (schema_version=1+) - no-op
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeNoOp,
		},
	}
}
