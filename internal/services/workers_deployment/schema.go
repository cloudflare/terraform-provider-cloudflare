// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*WorkersDeploymentResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"script_name": schema.StringAttribute{
				Description:   "Name of the script.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"strategy": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("percentage"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"versions": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"percentage": schema.Float64Attribute{
							Required: true,
							Validators: []validator.Float64{
								float64validator.Between(0.01, 100),
							},
						},
						"version_id": schema.StringAttribute{
							Required: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"annotations": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[WorkersDeploymentAnnotationsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"workers_message": schema.StringAttribute{
						Description: "Human-readable message about the deployment. Truncated to 100 bytes.",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"author_email": schema.StringAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"deployments": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[WorkersDeploymentDeploymentsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"strategy": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("percentage"),
							},
						},
						"versions": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[WorkersDeploymentDeploymentsVersionsModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"percentage": schema.Float64Attribute{
										Computed: true,
										Validators: []validator.Float64{
											float64validator.Between(0.01, 100),
										},
									},
									"version_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"annotations": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[WorkersDeploymentDeploymentsAnnotationsModel](ctx),
							Attributes: map[string]schema.Attribute{
								"workers_message": schema.StringAttribute{
									Description: "Human-readable message about the deployment. Truncated to 100 bytes.",
									Computed:    true,
								},
							},
						},
						"author_email": schema.StringAttribute{
							Computed: true,
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"source": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *WorkersDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkersDeploymentResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
