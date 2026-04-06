package queue

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*QueueResource)(nil)

// UpgradeState handles schema version upgrades for cloudflare_queue.
//
// Schema version history:
// - v4 (SDKv2): schema_version=0
// - v5 production (v5.0–v5.18): schema_version=1 (GetSchemaVersion(1, 500) returned 1)
// - v5 current: schema_version=500
//
// Upgrade paths:
// 1. v4 SDKv2 (schema_version=0) → v5 (500): Full transformation
//    - "name" → "queue_name"
//    - "id" → "queue_id"
//
// 2. v5 production (schema_version=1) → v5 (500): No-op
//    - State is already in v5 format; just bumps the version number.
func (r *QueueResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareQueueSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from the legacy SDKv2 cloudflare_queue (schema_version=0).
		// Transforms: name → queue_name, id → queue_id (copy).
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
		// Handle upgrades from v5 production state (schema_version=1).
		// Users on v5.0–v5.18 had GetSchemaVersion(1, 500) which stored state
		// at version 1. State is already in v5 format — no transformation needed.
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
