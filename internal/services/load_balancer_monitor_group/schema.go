// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*LoadBalancerMonitorGroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the Monitor Group to use for checking the health of origins within this pool.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "A short description of the monitor group",
				Required:    true,
			},
			"members": schema.ListNestedAttribute{
				Description: "List of monitors in this group",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Description: "Whether this monitor is enabled in the group",
							Required:    true,
						},
						"monitor_id": schema.StringAttribute{
							Description: "The ID of the Monitor to use for checking the health of origins within this pool.",
							Required:    true,
						},
						"monitoring_only": schema.BoolAttribute{
							Description: "Whether this monitor is used for monitoring only (does not affect pool health)",
							Required:    true,
						},
						"must_be_healthy": schema.BoolAttribute{
							Description: "Whether this monitor must be healthy for the pool to be considered healthy",
							Required:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "The timestamp of when the monitor was added to the group",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"updated_at": schema.StringAttribute{
							Description: "The timestamp of when the monitor group member was last updated",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the monitor group was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the monitor group was last updated",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *LoadBalancerMonitorGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *LoadBalancerMonitorGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
