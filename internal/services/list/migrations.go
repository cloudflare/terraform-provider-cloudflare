// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ListResource)(nil)

func (r *ListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceListSchema()
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 state (schema_version=0).
		// v4 SDKv2 provider used schema_version=0 (default).
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle upgrades from v5 state (schema_version=1).
		// Users must upgrade to the v5 release with Version: 1 before running tf-migrate.
		// No-op upgrade that just bumps the version to 500.
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
