package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceWorkerRouteModel represents the v4 cloudflare_worker_route state (schema_version=0).
// Resource type: cloudflare_worker_route (singular) or cloudflare_workers_route (plural)
//
// Key difference from v5: "script_name" field instead of "script".
type SourceWorkerRouteModel struct {
	ID         types.String `tfsdk:"id"`
	ZoneID     types.String `tfsdk:"zone_id"`
	Pattern    types.String `tfsdk:"pattern"`
	ScriptName types.String `tfsdk:"script_name"`
}

// TargetWorkersRouteModel represents the v5 cloudflare_workers_route state.
type TargetWorkersRouteModel struct {
	ID      types.String `tfsdk:"id"`
	ZoneID  types.String `tfsdk:"zone_id"`
	Pattern types.String `tfsdk:"pattern"`
	Script  types.String `tfsdk:"script"`
}
