// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_group/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessGroupResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessGroupResource)(nil)

// MoveState handles moves from cloudflare_access_group (v0) to cloudflare_zero_trust_access_group (v500).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_group.example
//	    to   = cloudflare_zero_trust_access_group.example
//	}
func (r *ZeroTrustAccessGroupResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceV4ZeroTrustAccessGroupSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_zero_trust_access_group.
// This is triggered when users manually run `terraform state mv` (Terraform < 1.8).
func (r *ZeroTrustAccessGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
		1: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}
