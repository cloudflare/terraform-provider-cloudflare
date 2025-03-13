// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*UserAgentBlockingRuleResource)(nil)

func (r *UserAgentBlockingRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
