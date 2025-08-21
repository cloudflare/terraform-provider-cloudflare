// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
)

var _ resource.ResourceWithUpgradeState = (*SnippetRulesResource)(nil)

func (r *SnippetRulesResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// The main migration is handled by the Grit patterns which convert
			// rules blocks to rules attributes. The state upgrader here just
			// ensures proper state structure is maintained during migration.
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Required: true,
					},
					"rules": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"expression": schema.StringAttribute{
									Required: true,
								},
								"last_updated": schema.StringAttribute{
									Computed: true,
									CustomType: timetypes.RFC3339Type{},
								},
								"snippet_name": schema.StringAttribute{
									Required: true,
								},
								"description": schema.StringAttribute{
									Optional: true,
									Computed: true,
								},
								"enabled": schema.BoolAttribute{
									Optional: true,
									Computed: true,
								},
							},
						},
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Get the prior state data
				var priorStateData SnippetRulesModel
				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// The structure is the same between v4 and v5, the main difference 
				// is that v4 used blocks while v5 uses attributes. The migration tool
				// handles the config transformation, and state can remain the same.
				
				// Set the new state (no transformation needed for state)
				resp.Diagnostics.Append(resp.State.Set(ctx, priorStateData)...)
			},
		},
	}
}
