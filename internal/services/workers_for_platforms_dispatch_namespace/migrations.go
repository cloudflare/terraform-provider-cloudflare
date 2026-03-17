// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_dispatch_namespace/migration/v500"
)

var _ resource.ResourceWithMoveState = (*WorkersForPlatformsDispatchNamespaceResource)(nil)
var _ resource.ResourceWithUpgradeState = (*WorkersForPlatformsDispatchNamespaceResource)(nil)

// MoveState handles moves from cloudflare_workers_for_platforms_namespace (deprecated v4) to
// cloudflare_workers_for_platforms_dispatch_namespace (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_workers_for_platforms_namespace.example
//	    to   = cloudflare_workers_for_platforms_dispatch_namespace.example
//	}
func (r *WorkersForPlatformsDispatchNamespaceResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceWorkersForPlatformsNamespaceSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveFromWorkersForPlatformsNamespace,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_workers_for_platforms_dispatch_namespace.
// Version 0 handles state from v4 provider (both resource type variants had schema version 0).
// Version 1 handles v5 state that needs a version bump (no-op).
func (r *WorkersForPlatformsDispatchNamespaceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceWorkersForPlatformsNamespaceSchema()

	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 state (both cloudflare_workers_for_platforms_dispatch_namespace
		// and cloudflare_workers_for_platforms_namespace had schema version 0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle upgrades from v5 state at version 1 (no schema changes, just version bump)
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
