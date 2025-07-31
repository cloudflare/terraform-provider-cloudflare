// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*CustomPagesResource)(nil)

func (r *CustomPagesResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// We don't know whether this was created with the V4 or V5 provider;
			// When the schema changed in V5, the schema version wasn't incremented.
			// So we define a schema that can work for with BOTH schemas
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_id": schema.StringAttribute{
						Optional: true,
					},
					"zone_id": schema.StringAttribute{
						Optional: true,
					},
					"url": schema.StringAttribute{
						Required: true,
					},
					// type is present in v4 & not v5
					"type": schema.StringAttribute{
						Optional: true,
					},
					// identifier is present in v5, not v4
					"identifier": schema.StringAttribute{
						Optional: true,
					},
					// state is optional in v4, required in v5
					"state": schema.StringAttribute{
						Optional: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {

				var priorStateData struct {
					ID         types.String `tfsdk:"id"`
					Identifier types.String `tfsdk:"identifier"`
					AccountID  types.String `tfsdk:"account_id"`
					ZoneID     types.String `tfsdk:"zone_id"`
					State      types.String `tfsdk:"state"`
					URL        types.String `tfsdk:"url"`
					Type       types.String `tfsdk:"type"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// initialize
				var newState = CustomPagesModel{
					ID:         priorStateData.ID,
					Identifier: priorStateData.Identifier,
					AccountID:  priorStateData.AccountID,
					ZoneID:     priorStateData.ZoneID,
					State:      priorStateData.State,
					URL:        priorStateData.URL,
				}

				if newState.Identifier.IsNull() {
					// Prior schema is from V4, which had "type" instead of "identifier"
					newState.Identifier = priorStateData.Type
				} else {

				}

				if newState.State.IsNull() {
					newState.State = types.StringValue("default")
				}

				// Marshal the upgraded state
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
			},
		},
	}
}
