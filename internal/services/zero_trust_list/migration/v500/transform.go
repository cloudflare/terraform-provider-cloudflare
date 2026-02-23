package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts v4 cloudflare_teams_list state to v5 cloudflare_zero_trust_list state.
//
// Merges v4 "items" (string array) and "items_with_description" (object array)
// into unified v5 "items" (set of {value, description} objects).
// items_with_description entries come first (preserving v4 API order).
func Transform(ctx context.Context, source SourceTeamsListModel) (*TargetZeroTrustListModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetZeroTrustListModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Type:        source.Type,
		Name:        source.Name,
		Description: source.Description,
		// Computed fields — set to null, refreshed from API
		CreatedAt:   timetypes.NewRFC3339Null(),
		ListCount:   types.Float64Null(),
		UpdatedAt:   timetypes.NewRFC3339Null(),
	}

	// Merge items_with_description (first) and items (second) into unified items
	var items []*TargetItemModel

	// items_with_description → objects with value and description
	for _, iwd := range source.ItemsWithDescription {
		items = append(items, &TargetItemModel{
			Value:       iwd.Value,
			Description: iwd.Description,
		})
	}

	// items (string array) → objects with value only, description null
	for _, item := range source.Items {
		if item.IsNull() || item.IsUnknown() {
			continue
		}
		items = append(items, &TargetItemModel{
			Value:       item,
			Description: types.StringNull(),
		})
	}

	if len(items) > 0 {
		target.Items = &items
	}

	return target, diags
}
