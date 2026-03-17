// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AISearchTokenResource)(nil)

func (r *AISearchTokenResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}
