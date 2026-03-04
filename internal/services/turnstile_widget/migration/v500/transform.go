package v500

import (
	"context"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
)

// Transform converts source (v4) state to target (v5) state.
//
// Key transformations:
// 1. domains: types.Set → *[]types.String (with alphabetical sorting)
// 2. sitekey: Copy from ID (in v4, sitekey WAS the ID; in v5 it's a separate field)
// 3. All other fields: direct pass-through (types match)
// 4. New computed fields (created_on, modified_on, etc.): set to null (provider will refresh)
func Transform(ctx context.Context, source SourceCloudfareTurnstileWidgetModel) (*TargetTurnstileWidgetModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for turnstile_widget migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for turnstile_widget migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Mode.IsNull() || source.Mode.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"mode is required for turnstile_widget migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies (pass-through fields)
	target := &TargetTurnstileWidgetModel{
		ID:           source.ID,
		AccountID:    source.AccountID,
		Name:         source.Name,
		Mode:         source.Mode,
		Region:       source.Region,
		BotFightMode: source.BotFightMode,
		Offlabel:     source.OffLabel, // Note: field name case change (OffLabel → Offlabel)
		Secret:       source.Secret,
	}

	// Step 3: Transform domains (Set → *[]types.String with sorting)
	// CRITICAL: The Cloudflare API returns domains in alphabetical order.
	// Since v5 uses ListAttribute (ordered) instead of SetAttribute (unordered),
	// we MUST sort domains to prevent drift on every apply.
	if !source.Domains.IsNull() && !source.Domains.IsUnknown() {
		domains, transformDiags := convertSetToSortedStringSlice(ctx, source.Domains)
		diags.Append(transformDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Domains = &domains
	} else {
		// Empty or null domains - preserve as nil
		target.Domains = nil
	}

	// Step 4: Handle computed fields
	// sitekey: In v4, the sitekey WAS the ID. In v5, it's a separate field.
	// Copy the ID to sitekey to preserve this critical identifier.
	target.Sitekey = source.ID

	// Set other computed fields to null (will be refreshed from API on first read)
	// These are NEW fields in v5 that didn't exist in v4
	target.CreatedOn = timetypes.NewRFC3339Null()   // NEW in v5
	target.ModifiedOn = timetypes.NewRFC3339Null()  // NEW in v5
	target.ClearanceLevel = types.StringNull()      // NEW in v5
	target.EphemeralID = types.BoolNull()           // NEW in v5

	return target, diags
}

// convertSetToSortedStringSlice converts a types.Set to a sorted []types.String slice.
//
// This is critical for turnstile_widget because:
// - v4 uses SetAttribute (unordered)
// - v5 uses ListAttribute (ordered)
// - The Cloudflare API returns domains alphabetically
// - Without sorting, every terraform apply would show drift
func convertSetToSortedStringSlice(ctx context.Context, set types.Set) ([]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Extract Set to native Go []string
	// IMPORTANT: Extract directly to []string, NOT []attr.Value
	// See quick_reference_lessons_learned.md - "Cannot use attr.Value attr.Value" error
	var rawStrings []string
	diags.Append(set.ElementsAs(ctx, &rawStrings, false)...)
	if diags.HasError() {
		return nil, diags
	}

	// Sort alphabetically to match API ordering
	sort.Strings(rawStrings)

	// Convert to []types.String
	result := make([]types.String, 0, len(rawStrings))
	for _, str := range rawStrings {
		result = append(result, types.StringValue(str))
	}

	return result, diags
}
