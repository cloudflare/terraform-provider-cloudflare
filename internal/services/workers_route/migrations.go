// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*WorkersRouteResource)(nil)

func (r *WorkersRouteResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}
