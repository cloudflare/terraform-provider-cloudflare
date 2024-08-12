// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithUpgradeState = &WorkersCustomDomainResource{}

func (r *WorkersCustomDomainResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Hostname of the Worker Domain.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"hostname": schema.StringAttribute{
						Description:   "Hostname of the Worker Domain.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifer of the account.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"environment": schema.StringAttribute{
						Description: "Worker environment associated with the zone and hostname.",
						Required:    true,
					},
					"service": schema.StringAttribute{
						Description: "Worker service associated with the zone and hostname.",
						Required:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "Identifier of the zone.",
						Required:    true,
					},
					"zone_name": schema.StringAttribute{
						Description: "Name of the zone.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state WorkersCustomDomainModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
