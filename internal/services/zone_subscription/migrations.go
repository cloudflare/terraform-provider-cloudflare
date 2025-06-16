// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZoneSubscriptionResource)(nil)

func (r *ZoneSubscriptionResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
