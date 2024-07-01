// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r ZoneLockdownResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "The unique identifier of the Zone Lockdown rule.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "The timestamp of when the rule was created.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the rule.",
						Computed:    true,
					},
					"modified_on": schema.StringAttribute{
						Description: "The timestamp of when the rule was last modified.",
						Computed:    true,
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the rule is currently paused.",
						Computed:    true,
					},
					"urls": schema.ListAttribute{
						Description: "The URLs to include in the rule definition. You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns.",
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state ZoneLockdownModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
