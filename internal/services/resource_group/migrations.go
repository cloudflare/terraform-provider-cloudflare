// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ResourceGroupResource)(nil)

func (r *ResourceGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
