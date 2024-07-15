// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r FirewallRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"id": schema.StringAttribute{
						Description:   "The unique identifier of the firewall rule.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"path_id": schema.StringAttribute{
						Description:   "The unique identifier of the firewall rule.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"action": schema.StringAttribute{
						Description: "The action to apply to a matched request. The `log` action is only available on an Enterprise plan.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge", "allow", "log", "bypass"),
						},
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the firewall rule is currently paused.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the firewall rule.",
						Computed:    true,
					},
					"priority": schema.Float64Attribute{
						Description: "The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 2147483647),
						},
					},
					"products": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"ref": schema.StringAttribute{
						Description: "A short reference tag. Allows you to select related firewall rules.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state FirewallRuleModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
