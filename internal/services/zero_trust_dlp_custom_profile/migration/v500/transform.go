package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4 DLP profile) state to target (current v5 custom profile) state.
// This function is shared by both UpgradeFromV0 and MoveState handlers.
func Transform(ctx context.Context, source SourceCloudflareDLPProfileModel) (*TargetZeroTrustDLPCustomProfileModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zero_trust_dlp_custom_profile migration.",
		)
		return nil, diags
	}

	target := &TargetZeroTrustDLPCustomProfileModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Name:        source.Name,
		Description: source.Description,
	}

	// Convert entry[] → entries
	target.Entries = transformEntries(source.Entry)

	// Preserve context_awareness from v4 state if present.
	// Although deprecated in v5, we must keep user-set values.
	// tf-migrate ensures the v5 config also has context_awareness (adds default if absent),
	// so state and config stay aligned and no null is sent to the API.
	target.ContextAwareness = transformContextAwareness(source.ContextAwareness)

	// Direct copies
	target.AllowedMatchCount = source.AllowedMatchCount
	target.OCREnabled = source.OCREnabled

	// New optional fields — set to defaults
	target.AIContextEnabled = types.BoolValue(false)
	target.ConfidenceThreshold = types.StringValue("low")

	// shared_entries not present in v4
	target.SharedEntries = nil

	// Computed fields — set to null, will refresh from API
	target.CreatedAt = timetypes.NewRFC3339Null()
	target.UpdatedAt = timetypes.NewRFC3339Null()
	target.OpenAccess = types.BoolNull()
	target.Type = types.StringNull()

	return target, diags
}

// transformEntries converts v4 entry (TypeSet stored as list) to v5 entries (set of objects).
func transformEntries(sourceEntries []SourceEntryModel) *[]*TargetEntriesModel {
	if len(sourceEntries) == 0 {
		return nil
	}

	entries := make([]*TargetEntriesModel, 0, len(sourceEntries))
	for _, se := range sourceEntries {
		entry := &TargetEntriesModel{
			Enabled: se.Enabled,
			Name:    se.Name,
		}

		// Do NOT copy entry.id → entries.entry_id.
		// The v4 entry.id is an API-generated identifier that doesn't map to the v5
		// entry_id field (which is for user-specified cross-references). The v5 config
		// doesn't set entry_id, so copying it would create a plan diff.

		// entry.pattern (list MaxItems:1) → entries.pattern (single object)
		if len(se.Pattern) > 0 {
			p := se.Pattern[0]
			pattern := &TargetPatternModel{
				Regex: p.Regex,
			}
			// Convert empty string validation to null to avoid plan diff
			if !p.Validation.IsNull() && !p.Validation.IsUnknown() && p.Validation.ValueString() != "" {
				pattern.Validation = p.Validation
			} else {
				pattern.Validation = types.StringNull()
			}
			entry.Pattern = pattern
		}

		entries = append(entries, entry)
	}

	return &entries
}

// transformContextAwareness converts v4 context_awareness (list MaxItems:1) to v5 single nested object.
func transformContextAwareness(source []SourceContextAwarenessModel) *TargetContextAwarenessModel {
	if len(source) == 0 {
		return &TargetContextAwarenessModel{
			Enabled: types.BoolValue(false),
			Skip: &TargetContextAwarenessSkipModel{
				Files: types.BoolValue(false),
			},
		}
	}

	ca := source[0]
	target := &TargetContextAwarenessModel{
		Enabled: ca.Enabled,
	}

	if len(ca.Skip) > 0 {
		target.Skip = &TargetContextAwarenessSkipModel{
			Files: ca.Skip[0].Files,
		}
	}

	return target
}
