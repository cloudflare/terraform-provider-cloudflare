package zero_trust_dlp_custom_entry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustDLPCustomEntryResource)(nil)

func (r *ZeroTrustDLPCustomEntryResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
		// Handle upgrades from v5 production state (schema_version=1).
		// Users on v5.0–v5.18 had GetSchemaVersion(1, 500) which stored state
		// at version 1. State is already in v5 format — no transformation needed.
		1: {
			PriorSchema: &targetSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}
