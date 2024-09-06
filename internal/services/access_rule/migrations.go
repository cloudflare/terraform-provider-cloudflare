// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AccessRuleResource)(nil)

func (r *AccessRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
