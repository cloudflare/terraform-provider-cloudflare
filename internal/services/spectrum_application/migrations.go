package spectrum_application

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/spectrum_application/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*SpectrumApplicationResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Both v4 and v5 used schema_version=0 (neither set an explicit Version).
//
//   - Slot 0: no-op upgrader (safely bumps existing v5 users from 0→1)
//
// Testing: schema returns version 500
//   - Slot 0: v4→v5 full transformation (v4 state has schema_version=0)
//   - Slot 1: v5 no-op (v5 users already bumped to version=1 in prod)
func (r *SpectrumApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceSpectrumApplicationSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 provider (schema_version=0)
		// Full transformation: arrays→objects, origin_port_range→origin_port, etc.
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 provider with version=1 (production dormant state)
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
