// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r FilterResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"id": schema.StringAttribute{
						Description: "The unique identifier of the filter.",
						Optional:    true,
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the filter.",
						Computed:    true,
					},
					"expression": schema.StringAttribute{
						Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
						Computed:    true,
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the filter is currently paused.",
						Computed:    true,
					},
					"ref": schema.StringAttribute{
						Description: "A short reference tag. Allows you to select related filters.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state FilterModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
