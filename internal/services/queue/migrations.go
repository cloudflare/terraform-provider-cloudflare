// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithUpgradeState = &QueueResource{}

func (r *QueueResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"queue_id": schema.StringAttribute{
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"queue_name": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"consumers_total_count": schema.Float64Attribute{
						Computed: true,
					},
					"created_on": schema.StringAttribute{
						Computed: true,
					},
					"modified_on": schema.StringAttribute{
						Computed: true,
					},
					"producers_total_count": schema.Float64Attribute{
						Computed: true,
					},
					"producers": schema.ListAttribute{
						Computed:    true,
						ElementType: jsontypes.NewNormalizedNull().Type(ctx),
					},
					"consumers": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created_on": schema.StringAttribute{
									Computed: true,
								},
								"environment": schema.StringAttribute{
									Computed: true,
								},
								"queue_name": schema.StringAttribute{
									Computed: true,
								},
								"service": schema.StringAttribute{
									Computed: true,
								},
								"settings": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"batch_size": schema.Float64Attribute{
											Description: "The maximum number of messages to include in a batch.",
											Optional:    true,
										},
										"max_retries": schema.Float64Attribute{
											Description: "The maximum number of retries",
											Optional:    true,
										},
										"max_wait_time_ms": schema.Float64Attribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state QueueModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
