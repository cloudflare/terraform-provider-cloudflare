package queue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*QueueResource)(nil)

func (r *QueueResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareQueueSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from the legacy SDKv2 cloudflare_queue (schema_version=0).
		// Transforms: name → queue_name, id → queue_id (copy).
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
	}
}
