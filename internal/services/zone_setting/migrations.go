// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZoneSettingResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// zone_setting is a special case: it is a one-to-many migration from v4's
// cloudflare_zone_settings_override. The tf-migrate tool deletes the old v4
// state and users run terraform apply to recreate state as new v5 resources.
// Therefore, no v4→v5 state transformation is needed.
//
// Both upgraders are no-ops that just bump the schema version to 500:
// - Version 0: early v5 resources created before version was bumped
// - Version 1: v5 resources at the previous schema version
func (r *ZoneSettingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	noopUpgrader := func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		resp.State.Raw = req.State.Raw
	}
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &targetSchema,
			StateUpgrader: noopUpgrader,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: noopUpgrader,
		},
	}
}
