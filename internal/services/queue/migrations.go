package queue

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*QueueResource)(nil)

func (r *QueueResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing pass-through upgrader
		targetSchema := ResourceSchema(ctx)
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
