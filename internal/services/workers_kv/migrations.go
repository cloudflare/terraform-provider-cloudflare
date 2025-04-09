// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*WorkersKVResource)(nil)

func (r *WorkersKVResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}
