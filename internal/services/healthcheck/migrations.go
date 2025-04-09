// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*HealthcheckResource)(nil)

func (r *HealthcheckResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}
