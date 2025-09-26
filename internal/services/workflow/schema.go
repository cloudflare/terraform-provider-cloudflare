// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*WorkflowResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"workflow_name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"class_name": schema.StringAttribute{
				Required: true,
			},
			"script_name": schema.StringAttribute{
				Required: true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"is_deleted": schema.Float64Attribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"terminator_running": schema.Float64Attribute{
				Computed: true,
			},
			"triggered_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"version_id": schema.StringAttribute{
				Computed: true,
			},
			"instances": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[WorkflowInstancesModel](ctx),
				Attributes: map[string]schema.Attribute{
					"complete": schema.Float64Attribute{
						Computed: true,
					},
					"errored": schema.Float64Attribute{
						Computed: true,
					},
					"paused": schema.Float64Attribute{
						Computed: true,
					},
					"queued": schema.Float64Attribute{
						Computed: true,
					},
					"running": schema.Float64Attribute{
						Computed: true,
					},
					"terminated": schema.Float64Attribute{
						Computed: true,
					},
					"waiting": schema.Float64Attribute{
						Computed: true,
					},
					"waiting_for_pause": schema.Float64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *WorkflowResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkflowResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
