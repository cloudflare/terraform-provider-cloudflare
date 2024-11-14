// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*FirewallRuleResource)(nil)

func (r *FirewallRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
