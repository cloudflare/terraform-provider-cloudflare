// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 WorkersKVNamespace state to v5 format
// This is a simple pass-through migration - all fields are preserved as-is
func Transform(ctx context.Context, source SourceWorkersKVNamespaceModel) (*TargetWorkersKVNamespaceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetWorkersKVNamespaceModel{
		// Direct field copies - no transformations needed
		ID:        source.ID,
		AccountID: source.AccountID,
		Title:     source.Title,

		// SupportsURLEncoding is a new computed field in v5
		// Leave it null - provider will populate it on first refresh
		SupportsURLEncoding: types.BoolNull(),
	}

	return target, diags
}
