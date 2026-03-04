package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 SourceVirtualNetworkModel into the v5 target model.
//
// Field mapping:
//   - id                 → id                 (unchanged)
//   - account_id         → account_id         (unchanged)
//   - is_default         → is_default         (deprecated, pass-through)
//   - name               → name               (unchanged)
//   - comment            → comment            (defaults to "" if null)
//   - is_default_network → is_default_network (defaults to false if null)
//   - created_at         → created_at         (computed, pass-through)
//   - deleted_at         → deleted_at         (computed, pass-through)
func Transform(ctx context.Context, source *SourceVirtualNetworkModel) (*TargetVirtualNetworkModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetVirtualNetworkModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		IsDefault: source.IsDefault,
		Name:      source.Name,
		CreatedAt: source.CreatedAt,
		DeletedAt: source.DeletedAt,
	}

	// Ensure comment has its default value of "" if null or missing.
	// v5 schema declares Default: stringdefault.StaticString("") but older v4 state
	// may not have this field populated.
	if source.Comment.IsNull() || source.Comment.IsUnknown() {
		target.Comment = types.StringValue("")
	} else {
		target.Comment = source.Comment
	}

	// Ensure is_default_network has its default value of false if null or missing.
	// v5 schema declares Default: booldefault.StaticBool(false) but older v4 state
	// may not have this field populated.
	if source.IsDefaultNetwork.IsNull() || source.IsDefaultNetwork.IsUnknown() {
		target.IsDefaultNetwork = types.BoolValue(false)
	} else {
		target.IsDefaultNetwork = source.IsDefaultNetwork
	}

	return target, diags
}
