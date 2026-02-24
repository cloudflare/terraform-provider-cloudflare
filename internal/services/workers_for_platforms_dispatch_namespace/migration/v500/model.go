// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceWorkersForPlatformsNamespaceModel represents the v4 state structure.
// Both cloudflare_workers_for_platforms_namespace (deprecated) and
// cloudflare_workers_for_platforms_dispatch_namespace (current v4) had identical schemas.
type SourceWorkersForPlatformsNamespaceModel struct {
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
}

// TargetWorkersForPlatformsDispatchNamespaceModel represents the v5 state structure.
type TargetWorkersForPlatformsDispatchNamespaceModel struct {
	ID             types.String      `tfsdk:"id"`
	NamespaceName  types.String      `tfsdk:"namespace_name"`
	AccountID      types.String      `tfsdk:"account_id"`
	Name           types.String      `tfsdk:"name"`
	CreatedBy      types.String      `tfsdk:"created_by"`
	CreatedOn      timetypes.RFC3339 `tfsdk:"created_on"`
	ModifiedBy     types.String      `tfsdk:"modified_by"`
	ModifiedOn     timetypes.RFC3339 `tfsdk:"modified_on"`
	NamespaceID    types.String      `tfsdk:"namespace_id"`
	ScriptCount    types.Int64       `tfsdk:"script_count"`
	TrustedWorkers types.Bool        `tfsdk:"trusted_workers"`
}
