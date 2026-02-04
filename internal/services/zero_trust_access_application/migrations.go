package zero_trust_access_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_application/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessApplicationResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessApplicationResource)(nil)

// MoveState handles moves from cloudflare_access_application (v4) to cloudflare_zero_trust_access_application (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_application.example
//	    to   = cloudflare_zero_trust_access_application.example
//	}
func (r *ZeroTrustAccessApplicationResource) MoveState(ctx context.Context) []resource.StateMover {
	v4Schema := v500.SourceAccessApplicationSchema()
	return []resource.StateMover{
		{
			SourceSchema: &v4Schema,
			StateMover:   v500.MoveFromAccessApplication,
		},
	}
}

// UpgradeState handles schema version upgrades.
func (r *ZeroTrustAccessApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := v500.SourceAccessApplicationSchema()
	v5Schema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from earlier v5 versions (no schema changes, just version bump)
		0: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle state moved from cloudflare_access_application (v4 provider)
		// When users run `terraform state mv cloudflare_access_application.x cloudflare_zero_trust_access_application.x`,
		// the old schema_version=2 is preserved, triggering this upgrader.
		2: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}
