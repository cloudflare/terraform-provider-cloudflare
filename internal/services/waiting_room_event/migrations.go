// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*WaitingRoomEventResource)(nil)

func (r *WaitingRoomEventResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
