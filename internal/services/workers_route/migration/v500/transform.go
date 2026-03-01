package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts v4 cloudflare_worker_route state to v5 cloudflare_workers_route state.
// The only transformation is: script_name → script.
func Transform(ctx context.Context, source SourceWorkerRouteModel) (*TargetWorkersRouteModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetWorkersRouteModel{
		ID:      source.ID,
		ZoneID:  source.ZoneID,
		Pattern: source.Pattern,
		Script:  source.ScriptName, // script_name → script
	}

	// If script_name was null/empty, script should be null
	if source.ScriptName.IsNull() || source.ScriptName.IsUnknown() {
		target.Script = types.StringNull()
	}

	return target, diags
}
