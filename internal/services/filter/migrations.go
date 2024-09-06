// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*FilterResource)(nil)

func (r *FilterResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
