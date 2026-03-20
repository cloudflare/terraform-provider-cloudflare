package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4 DLP profile with type="predefined") state
// to target (current v5 predefined profile) state.
// This function is shared by both UpgradeFromV0 and MoveState handlers.
func Transform(ctx context.Context, source SourceCloudflareDLPProfileModel) (*TargetZeroTrustDLPPredefinedProfileModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zero_trust_dlp_predefined_profile migration.",
		)
		return nil, diags
	}

	target := &TargetZeroTrustDLPPredefinedProfileModel{
		AccountID: source.AccountID,
	}

	// v4 id → v5 profile_id + id
	target.ProfileID = source.ID
	target.ID = source.ID

	// v4 name → v5 name (now computed, but preserve from state)
	target.Name = source.Name

	// Extract enabled entry IDs from v4 entries
	target.EnabledEntries = extractEnabledEntryIDs(source.Entry)

	// Direct copies
	target.AllowedMatchCount = source.AllowedMatchCount
	target.OCREnabled = source.OCREnabled

	// New optional fields — set to defaults
	target.AIContextEnabled = types.BoolValue(false)
	target.ConfidenceThreshold = types.StringValue("low")

	// Computed fields — set to null, will refresh from API
	target.OpenAccess = types.BoolNull()
	target.Entries = customfield.NullObjectList[TargetPredefinedProfileEntriesModel](ctx)

	return target, diags
}

// extractEnabledEntryIDs collects the IDs of all enabled entries from v4 state.
//
// When the v4 state has entry blocks but none are enabled, this returns a pointer to
// an empty slice (not nil). This is important for state upgrade correctness: the v5
// config produced by tf-migrate writes `enabled_entries = []` for the all-disabled case,
// and the upgraded state must match that exactly — a nil pointer would cause a one-time
// plan diff of `+ enabled_entries = []` on the first plan after the moved block fires.
//
// When the v4 state has no entry blocks at all (len == 0), nil is returned because the
// resource was never configured with entries and the user's config will not have
// enabled_entries set.
func extractEnabledEntryIDs(entries []SourceEntryModel) *[]types.String {
	if len(entries) == 0 {
		return nil
	}

	var enabledIDs []types.String
	for _, entry := range entries {
		if !entry.Enabled.IsNull() && entry.Enabled.ValueBool() {
			if !entry.ID.IsNull() && !entry.ID.IsUnknown() {
				enabledIDs = append(enabledIDs, entry.ID)
			}
		}
	}

	// Always return a non-nil pointer when v4 entry blocks existed. If none were
	// enabled, return a pointer to an empty slice so the state matches `enabled_entries = []`
	// in the migrated config, avoiding a spurious plan diff after state upgrade.
	if enabledIDs == nil {
		enabledIDs = []types.String{}
	}
	return &enabledIDs
}
