// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 WorkersForPlatforms namespace state to v5 format.
//
// Key transformation: copies id to namespace_name.
// In v5, the provider uses namespace_name for all Read and Delete API calls:
//
//	r.client.WorkersForPlatforms.Dispatch.Namespaces.Get(ctx, data.NamespaceName.ValueString(), ...)
//	r.client.WorkersForPlatforms.Dispatch.Namespaces.Delete(ctx, data.NamespaceName.ValueString(), ...)
//
// New computed fields are set to null; the provider will populate them on the first refresh.
func Transform(ctx context.Context, source SourceWorkersForPlatformsNamespaceModel) (*TargetWorkersForPlatformsDispatchNamespaceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetWorkersForPlatformsDispatchNamespaceModel{
		// Direct field copies
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,

		// Key transformation: namespace_name must equal id.
		// v4 stored the namespace identifier only in id; v5 requires it in namespace_name too.
		NamespaceName: source.ID,

		// New computed fields - set to null; provider will populate on first refresh
		CreatedBy:      types.StringNull(),
		CreatedOn:      timetypes.NewRFC3339Null(),
		ModifiedBy:     types.StringNull(),
		ModifiedOn:     timetypes.NewRFC3339Null(),
		NamespaceID:    types.StringNull(),
		ScriptCount:    types.Int64Null(),
		TrustedWorkers: types.BoolNull(),
	}

	return target, diags
}
