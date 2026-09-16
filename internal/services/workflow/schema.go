// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.ResourceWithConfigValidators = (*WorkflowResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
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
			"concurrency": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"limit": schema.Int64Attribute{
						Description: "Maximum number of instances of this workflow that can run concurrently. Additional instances are queued and started as running instances complete. Must not exceed the account concurrency limit.",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 9007199254740991),
						},
					},
				},
			},
			"default_retention": schema.SingleNestedAttribute{
				Description: "Default retention applied to instances of this version when they do not set their own retention.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"error_retention": schema.DynamicAttribute{
						Description: "Specifies the duration in milliseconds or as a string like '5 minutes'.",
						Optional:    true,
						Validators: []validator.Dynamic{
							customvalidator.AllowedSubtypes(basetypes.Int64Type{}, basetypes.StringType{}),
						},
						CustomType:    customfield.NormalizedDynamicType{},
						PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
					},
					"success_retention": schema.DynamicAttribute{
						Description: "Specifies the duration in milliseconds or as a string like '5 minutes'.",
						Optional:    true,
						Validators: []validator.Dynamic{
							customvalidator.AllowedSubtypes(basetypes.Int64Type{}, basetypes.StringType{}),
						},
						CustomType:    customfield.NormalizedDynamicType{},
						PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
					},
				},
			},
			"limits": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"steps": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.Between(1, 9007199254740991),
						},
					},
				},
			},
			"schedules": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cron": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
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
			"instances": schema.MapAttribute{
				Computed:    true,
				CustomType:  customfield.NewMapType[types.Float64](ctx),
				ElementType: types.Float64Type,
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
