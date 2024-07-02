// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r WaitingRoomRulesResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"waiting_room_id": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"rule_id": schema.StringAttribute{
						Description:   "The ID of the rule.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"action": schema.StringAttribute{
						Description: "The action to take when the expression matches.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("bypass_waiting_room"),
						},
					},
					"expression": schema.StringAttribute{
						Description: "Criteria defining when there is a match for the current rule.",
						Required:    true,
					},
					"description": schema.StringAttribute{
						Description: "The description of the rule.",
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString(""),
					},
					"enabled": schema.BoolAttribute{
						Description: "When set to true, the rule is enabled.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state WaitingRoomRulesModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
