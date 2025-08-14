// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

/*

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*LoadBalancerPoolResource)(nil)

// UpgradeState handles schema version 0 state that has already been migrated from v4 to v5
// by the cmd/migrate tool. This function only handles v5 state structure.
// All v4→v5 transformations are handled in cmd/migrate/state.go
func (r *LoadBalancerPoolResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// PriorSchema represents the v5 schema version 0 structure
			// This ensures we're receiving state that's already been migrated to v5 format
			PriorSchema: &schema.Schema{
				Version: 0,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"origins": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Required: true,
								},
								"name": schema.StringAttribute{
									Required: true,
								},
								"enabled": schema.BoolAttribute{
									Optional: true,
									Computed: true,
									Default:  booldefault.StaticBool(true),
								},
								"weight": schema.Float64Attribute{
									Optional: true,
									Computed: true,
									Default:  float64default.StaticFloat64(1.0),
								},
								"header": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"host": schema.ListAttribute{
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
								"virtual_network_id": schema.StringAttribute{
									Optional: true,
								},
								"disabled_at": schema.StringAttribute{
									Computed: true,
									CustomType: timetypes.RFC3339Type{},
								},
							},
						},
					},
					"enabled": schema.BoolAttribute{
						Optional: true,
						Computed: true,
						Default:  booldefault.StaticBool(true),
					},
					"minimum_origins": schema.Int64Attribute{
						Optional: true,
						Computed: true,
						Default:  int64default.StaticInt64(1),
					},
					"latitude": schema.Float64Attribute{
						Optional: true,
					},
					"longitude": schema.Float64Attribute{
						Optional: true,
					},
					"description": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString(""),
					},
					"monitor": schema.StringAttribute{
						Optional: true,
					},
					"notification_email": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString(""),
					},
					"notification_filter": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
							"disable": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
							"healthy": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"check_regions": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"load_shedding": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"default_percent": schema.Float64Attribute{
								Optional: true,
								Computed: true,
								Default:  float64default.StaticFloat64(0),
							},
							"default_policy": schema.StringAttribute{
								Optional: true,
								Computed: true,
								Default:  stringdefault.StaticString(""),
							},
							"session_percent": schema.Float64Attribute{
								Optional: true,
								Computed: true,
								Default:  float64default.StaticFloat64(0),
							},
							"session_policy": schema.StringAttribute{
								Optional: true,
								Computed: true,
								Default:  stringdefault.StaticString(""),
							},
						},
					},
					"origin_steering": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"policy": schema.StringAttribute{
								Optional: true,
								Computed: true,
								Default:  stringdefault.StaticString("random"),
							},
						},
					},
					"created_on": schema.StringAttribute{
						Computed: true,
					},
					"modified_on": schema.StringAttribute{
						Computed: true,
					},
					"disabled_at": schema.StringAttribute{
						Computed: true,
						CustomType: timetypes.RFC3339Type{},
					},
					"networks": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// At this point, the PriorSchema has validated that the state is in v5 format
				// Since cmd/migrate has already transformed v4→v5, we just need to pass the state through
				// to upgrade from schema version 0 to version 1
				
				var priorState LoadBalancerPoolModel
				
				// Get the prior state - this will use the PriorSchema to parse it
				diags := req.State.Get(ctx, &priorState)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
				
				// The state is already in the correct format, just pass it through
				// The only change is the schema version number (0 → 1)
				resp.Diagnostics.Append(resp.State.Set(ctx, priorState)...)
			},
		},
	}
}
*/
