package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceWorkerRouteSchema returns the v4 cloudflare_worker_route schema.
// Used by MoveState for reading v4 state during singular→plural rename.
func SourceWorkerRouteSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"pattern": schema.StringAttribute{
				Optional: true,
			},
			"script_name": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

// UnionV0Schema returns a PriorSchema for version 0 that can parse both V4 and V5 state.
// V4 has "script_name", V5 has "script". Both are simple strings — no type conflicts.
func UnionV0Schema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"pattern": schema.StringAttribute{
				Optional: true,
			},
			// V4 field
			"script_name": schema.StringAttribute{
				Optional: true,
			},
			// V5 field
			"script": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}
