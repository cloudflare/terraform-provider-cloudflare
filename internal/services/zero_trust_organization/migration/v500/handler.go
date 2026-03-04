package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV1 handles state upgrades from v5 with explicit version=1 to version 500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
//
// Note: This upgrader may not be used in practice since published v5 releases used
// version 0 (no Version field), not version 1. Kept for completeness.
func UpgradeFromV1(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading zero_trust_organization state from version 1 to version 500 (no-op)")
	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw
}

// UpgradeFromV0 handles state upgrades from published v5 releases (schema_version=0)
// to current v5 (version=500).
//
// Published v5 releases (before Version field was added) defaulted to version 0.
// These states already use v5 schema format (login_design as object, etc.) and
// just need a version bump to 500.
//
// Note: v4→v5 migration is NOT handled here. It's handled by MoveState which
// transforms the state during resource rename (cloudflare_access_organization →
// cloudflare_zero_trust_organization) and outputs state at version 500.
func UpgradeFromV0(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading zero_trust_organization from published v5 (version 0) to version 500")
	// Published v5 already uses v5 schema format
	// Just copy state through for version bump
	resp.State.Raw = req.State.Raw
}
