package managed_transforms

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/managed_transforms/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ManagedTransformsResource)(nil)
var _ resource.ResourceWithMoveState = (*ManagedTransformsResource)(nil)

// MoveState handles moves from cloudflare_managed_headers to cloudflare_managed_transforms.
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_managed_headers.example
//	    to   = cloudflare_managed_transforms.example
//	}
func (r *ManagedTransformsResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceManagedHeadersSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// v4 (cloudflare_managed_headers) used schema_version=0, v5 uses schema_version=1.
//
//   - Slot 0: no-op upgrader (bumps existing v5 users from 0→1)
//
// Testing: schema returns version 500
//   - Slot 0: v4→v5 full transformation (v4 state from `terraform state mv`)
//   - Slot 1: v5 no-op (existing v5 users bumped to 500)
func (r *ManagedTransformsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceManagedHeadersSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 managed_headers (schema_version=0, via `terraform state mv`)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 managed_transforms (version=1)
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
