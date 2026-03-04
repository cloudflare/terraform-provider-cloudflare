package regional_hostname

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_hostname/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*RegionalHostnameResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Both v4 and v5 used schema_version=0 (neither set an explicit Version).
//
//   - Slot 0: no-op upgrader (safely bumps existing v5 users from 0→1)
//
// Testing: schema returns version 500
//   - Slot 0: v4→v5 no-op (schema identical, just version bump)
//   - Slot 1: v5 no-op (existing v5 users bumped to 500)
func (r *RegionalHostnameResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceRegionalHostnameSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 provider with version=1
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
