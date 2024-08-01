// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"urls": schema.ListAttribute{
						Description: "The URLs to include in the current WAF override. You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns.",
						Required:    true,
						ElementType: types.StringType,
					},
					"configurations": schema.SingleNestedAttribute{
						Description: "A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of `ip` or `ip_range` configurations.",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"target": schema.StringAttribute{
								Description: "The configuration target. You must set the target to `ip` when specifying an IP address in the Zone Lockdown rule.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("ip", "ip_range"),
								},
							},
							"value": schema.StringAttribute{
								Description: "The IP address to match. This address will be compared to the IP address of incoming requests.",
								Optional:    true,
							},
						},
					},
					"created_on": schema.StringAttribute{
						Description: "The timestamp of when the rule was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the rule.",
						Computed:    true,
					},
					"modified_on": schema.StringAttribute{
						Description: "The timestamp of when the rule was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the rule is currently paused.",
						Computed:    true,
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
