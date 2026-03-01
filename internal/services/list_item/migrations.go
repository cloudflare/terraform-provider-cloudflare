// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ListItemResource)(nil)

func (r *ListItemResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing upgraders only
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// Test mode (TF_MIG_TEST=1): full StateUpgrader migration
	sourceSchema := v500.SourceListItemSchema()
	v1Schema := v500.SourceListItemV1Schema()
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 SDKv2 (schema_version=0) to v500
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle upgrades from v4.52.5 framework (schema_version=1) to v500
		// v4.52.5 uses ListNestedBlock for hostname/redirect, v5 uses SingleNestedAttribute
		1: {
			PriorSchema:   &v1Schema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
