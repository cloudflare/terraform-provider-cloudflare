// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ListResource)(nil)

func (r *ListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceListSchema()
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from schema_version=0 (both v4 and v5 state).
		// The published v5 provider writes schema_version=0 (no Version field set),
		// so both v4 and v5 state arrive here. The handler detects the format.
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
	}
}
