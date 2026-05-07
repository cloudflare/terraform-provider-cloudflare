// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*UserGroupResource)(nil)

func (r *UserGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
