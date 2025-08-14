// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
)

var _ resource.ResourceWithUpgradeState = (*RegionalHostnameResource)(nil)

func (r *RegionalHostnameResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// PriorSchema includes fields that can be handled with typed structs
			// but excludes timeouts which we'll handle via RawState
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"hostname": schema.StringAttribute{
						Required: true,
					},
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"routing": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"region_key": schema.StringAttribute{
						Required: true,
					},
					"created_on": schema.StringAttribute{
						Computed: true,
						CustomType: timetypes.RFC3339Type{},
					},
					// Note: intentionally omitting timeouts so it gets handled via RawState
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Get typed data for unchanged fields
				var priorStateData struct {
					ID        types.String      `tfsdk:"id"`
					Hostname  types.String      `tfsdk:"hostname"`
					ZoneID    types.String      `tfsdk:"zone_id"`
					Routing   types.String      `tfsdk:"routing"`
					RegionKey types.String      `tfsdk:"region_key"`
					CreatedOn timetypes.RFC3339 `tfsdk:"created_on"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Initialize new state with unchanged fields
				newState := RegionalHostnameModel{
					ID:        priorStateData.ID,
					Hostname:  priorStateData.Hostname,
					ZoneID:    priorStateData.ZoneID,
					Routing:   priorStateData.Routing,
					RegionKey: priorStateData.RegionKey,
					CreatedOn: priorStateData.CreatedOn,
				}

				// Handle routing default value - v4 didn't have this field
				if newState.Routing.IsNull() || newState.Routing.ValueString() == "" {
					newState.Routing = types.StringValue("dns")
				}

				// Handle timeouts removal from RawState - we intentionally ignore timeouts
				// The timeouts field from v4 will simply be dropped as it's not included
				// in the new state model. No special handling needed since we're not
				// including it in the target state.

				// Marshal the upgraded state (timeouts will be completely omitted)
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
