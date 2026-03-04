package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) snippet_rules state to target (current v5) state.
//
// The snippet_rules migration is straightforward — no field renames or type changes.
// The main structural change (blocks → list) is handled at the HCL config level by tf-migrate.
// For state, the rules are already stored as a JSON array in v4 state.
//
// Key actions:
// - Copy zone_id and all rule fields directly
// - Set computed fields (id, last_updated) to null (will refresh from API)
func Transform(ctx context.Context, source SourceSnippetRulesModel) (*TargetSnippetRulesModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for snippet_rules migration.",
		)
		return nil, diags
	}

	target := &TargetSnippetRulesModel{
		ZoneID: source.ZoneID,
	}

	if len(source.Rules) > 0 {
		rules := make([]*TargetSnippetRuleModel, len(source.Rules))
		for i, r := range source.Rules {
			rules[i] = &TargetSnippetRuleModel{
				Expression:  r.Expression,
				SnippetName: r.SnippetName,
				Enabled:     r.Enabled,
				Description: r.Description,
				// Computed fields — set to null, will refresh from API
				ID:          types.StringNull(),
				LastUpdated: timetypes.NewRFC3339Null(),
			}
		}
		target.Rules = &rules
	}

	return target, diags
}
