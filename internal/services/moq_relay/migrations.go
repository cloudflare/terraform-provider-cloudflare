// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*MoqRelayResource)(nil)

func (r *MoqRelayResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
