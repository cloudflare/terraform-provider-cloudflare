// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group_members

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*UserGroupMembersResource)(nil)

func (r *UserGroupMembersResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
