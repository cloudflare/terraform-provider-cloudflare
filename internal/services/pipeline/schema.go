// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PipelineResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Indicates a unique identifier for this pipeline.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Specifies the public ID of the account.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Specifies the name of the Pipeline.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"sql": schema.StringAttribute{
				Description:   "Specifies SQL for the Pipeline processing flow.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"failure_reason": schema.StringAttribute{
				Description: "Indicates the reason for the failure of the Pipeline.",
				Computed:    true,
			},
			"modified_at": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: "Indicates the current status of the Pipeline.",
				Computed:    true,
			},
			"tables": schema.ListNestedAttribute{
				Description: "List of streams and sinks used by this pipeline.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PipelineTablesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the connection (stream or sink).",
							Computed:    true,
						},
						"latest": schema.Int64Attribute{
							Description: "Latest available version of the connection.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the connection.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Type of the connection.\nAvailable values: \"stream\", \"sink\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("stream", "sink"),
							},
						},
						"version": schema.Int64Attribute{
							Description: "Current version of the connection used by this pipeline.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *PipelineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PipelineResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
