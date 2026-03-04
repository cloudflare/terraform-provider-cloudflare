// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_pages/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*CustomPagesResource)(nil)

// UpgradeState handles schema version upgrades for cloudflare_custom_pages.
// Version 0 handles state from v4 provider (schema version 0).
// Version 1 handles v5 state that needs a version bump (no-op).
func (r *CustomPagesResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 state (dual-format detection: v4 or early v5)
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// Handle upgrades from v5 state at version 1 (no schema changes, just version bump)
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
