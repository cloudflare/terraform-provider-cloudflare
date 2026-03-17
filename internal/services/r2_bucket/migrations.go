// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*R2BucketResource)(nil)

func (r *R2BucketResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := v500.V4Schema(ctx)
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle state from v4 Plugin Framework provider (version=0)
		// v4 r2_bucket had 4 fields: id, name, account_id, location
		// v5 adds: jurisdiction, storage_class, creation_date
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle state from v5 Plugin Framework provider (version=1)
		// This is a no-op upgrade for stepping stone compatibility
		1: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}
