// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ListItemResource)(nil)

func (r *ListItemResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

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
