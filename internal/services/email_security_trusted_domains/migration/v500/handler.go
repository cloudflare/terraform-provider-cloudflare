package v500

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UpgradeFromV0 converts state from schema_version=0 to schema_version=500.
//
// Transformations:
//   - id: Int64 → String (decimal representation of the prior int)
//   - body[0] nested fields → top-level (is_recent / is_regex / is_similarity /
//     pattern / comments)
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var prior SourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	next := TargetModel{
		AccountID:    prior.AccountID,
		CreatedAt:    prior.CreatedAt,
		LastModified: prior.LastModified,
		// v0 had both nested-body and top-level fields; default to top-level.
		IsRecent:     prior.IsRecent,
		IsRegex:      prior.IsRegex,
		IsSimilarity: prior.IsSimilarity,
		Pattern:      prior.Pattern,
		Comments:     prior.Comments,
	}

	if !prior.ID.IsNull() && !prior.ID.IsUnknown() {
		next.ID = types.StringValue(strconv.FormatInt(prior.ID.ValueInt64(), 10))
	}

	// If body[0] was set, its nested values take precedence.
	if prior.Body != nil && len(*prior.Body) > 0 {
		b := (*prior.Body)[0]
		next.IsRecent = b.IsRecent
		next.IsRegex = b.IsRegex
		next.IsSimilarity = b.IsSimilarity
		next.Pattern = b.Pattern
		if !b.Comments.IsNull() {
			next.Comments = b.Comments
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &next)...)
}
