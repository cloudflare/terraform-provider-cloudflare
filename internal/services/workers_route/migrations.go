// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*WorkersRouteResource)(nil)

func (r *WorkersRouteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// Handle v4 -> v5 migration: script_name -> script
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"pattern": schema.StringAttribute{
						Required: true,
					},
					"script_name": schema.StringAttribute{
						Optional: true,
					},
					// Also include "script" in case this is from v5 already
					"script": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData struct {
					ID         types.String `tfsdk:"id"`
					ZoneID     types.String `tfsdk:"zone_id"`
					Pattern    types.String `tfsdk:"pattern"`
					ScriptName types.String `tfsdk:"script_name"`
					Script     types.String `tfsdk:"script"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Initialize new state
				var newState = WorkersRouteModel{
					ID:      priorStateData.ID,
					ZoneID:  priorStateData.ZoneID,
					Pattern: priorStateData.Pattern,
					Script:  priorStateData.Script,
				}

				// If script is null but script_name exists, migrate it
				if newState.Script.IsNull() && !priorStateData.ScriptName.IsNull() {
					newState.Script = priorStateData.ScriptName
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
