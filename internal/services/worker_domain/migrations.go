// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r WorkerDomainResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Identifer of the Worker Domain.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"environment": schema.StringAttribute{
						Description: "Worker environment associated with the zone and hostname.",
						Required:    true,
					},
					"hostname": schema.StringAttribute{
						Description: "Hostname of the Worker Domain.",
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
				var state WorkerDomainModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
