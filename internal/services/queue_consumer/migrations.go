// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*QueueConsumerResource)(nil)

func (r *QueueConsumerResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}
