// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue_consumer/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*QueueConsumerResource)(nil)

// UpgradeState handles schema version upgrades for cloudflare_queue_consumer.
//
// cloudflare_queue_consumer is a v5-native resource with no v4 predecessor.
// Both upgraders are no-ops — the schema has not changed, only the version number.
//
// Version history:
//   - 0: Original v5 state (before schema versioning was applied)
//   - 1: Dormant production version
//   - 500: Active migration version
func (r *QueueConsumerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareQueueConsumerSchema()

	return map[int64]resource.StateUpgrader{
		// Handle v5 states created before explicit schema versioning (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle dormant production v5 states (schema_version=1) -> v500 (no-op)
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
