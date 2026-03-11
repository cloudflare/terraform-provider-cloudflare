// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image_variant

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ImageVariantResource)(nil)

// UpgradeState handles state upgrades for image_variant.
// This resource is new in v5, so we only need a no-op upgrader from version 1 to 500
// for forward compatibility with future schema changes.
func (r *ImageVariantResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// No-op upgrader for v5 state (version 1 -> 500)
		1: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}
