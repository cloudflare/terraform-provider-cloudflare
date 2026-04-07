// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_dynamic_routing

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*AIGatewayDynamicRoutingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"gateway_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"elements": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"outputs": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"next": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
								"false": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
								"true": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
								"element_id": schema.StringAttribute{
									Optional: true,
								},
								"fallback": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
								"success": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
						},
						"type": schema.StringAttribute{
							Description: `Available values: "start", "conditional", "percentage", "rate", "model", "end".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"start",
									"conditional",
									"percentage",
									"rate",
									"model",
									"end",
								),
							},
						},
						"properties": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"conditions": schema.StringAttribute{
									Optional:   true,
									CustomType: jsontypes.NormalizedType{},
								},
								"key": schema.StringAttribute{
									Optional: true,
								},
								"limit": schema.Float64Attribute{
									Optional: true,
								},
								"limit_type": schema.StringAttribute{
									Description: `Available values: "count", "cost".`,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("count", "cost"),
									},
								},
								"window": schema.Float64Attribute{
									Optional: true,
								},
								"model": schema.StringAttribute{
									Optional: true,
								},
								"ai_gateway_dynamic_routing_provider": schema.StringAttribute{
									Optional: true,
								},
								"retries": schema.Float64Attribute{
									Optional: true,
								},
								"timeout": schema.Float64Attribute{
									Optional: true,
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
			"deployment": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingDeploymentModel](ctx),
				Attributes: map[string]schema.Attribute{
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"deployment_id": schema.StringAttribute{
						Computed: true,
					},
					"version_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"route": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"account_tag": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"deployment": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteDeploymentModel](ctx),
						Attributes: map[string]schema.Attribute{
							"created_at": schema.StringAttribute{
								Computed: true,
							},
							"deployment_id": schema.StringAttribute{
								Computed: true,
							},
							"version_id": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"elements": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[AIGatewayDynamicRoutingRouteElementsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"outputs": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsOutputsModel](ctx),
									Attributes: map[string]schema.Attribute{
										"next": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsOutputsNextModel](ctx),
											Attributes: map[string]schema.Attribute{
												"element_id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"false": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsOutputsFalseModel](ctx),
											Attributes: map[string]schema.Attribute{
												"element_id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"true": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsOutputsTrueModel](ctx),
											Attributes: map[string]schema.Attribute{
												"element_id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"element_id": schema.StringAttribute{
											Computed: true,
										},
										"fallback": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsOutputsFallbackModel](ctx),
											Attributes: map[string]schema.Attribute{
												"element_id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"success": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsOutputsSuccessModel](ctx),
											Attributes: map[string]schema.Attribute{
												"element_id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
									},
								},
								"type": schema.StringAttribute{
									Description: `Available values: "start", "conditional", "percentage", "rate", "model", "end".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"start",
											"conditional",
											"percentage",
											"rate",
											"model",
											"end",
										),
									},
								},
								"properties": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteElementsPropertiesModel](ctx),
									Attributes: map[string]schema.Attribute{
										"conditions": schema.StringAttribute{
											Computed:   true,
											CustomType: jsontypes.NormalizedType{},
										},
										"key": schema.StringAttribute{
											Computed: true,
										},
										"limit": schema.Float64Attribute{
											Computed: true,
										},
										"limit_type": schema.StringAttribute{
											Description: `Available values: "count", "cost".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("count", "cost"),
											},
										},
										"window": schema.Float64Attribute{
											Computed: true,
										},
										"model": schema.StringAttribute{
											Computed: true,
										},
										"ai_gateway_dynamic_routing_provider": schema.StringAttribute{
											Computed: true,
										},
										"retries": schema.Float64Attribute{
											Computed: true,
										},
										"timeout": schema.Float64Attribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"gateway_id": schema.StringAttribute{
						Computed: true,
					},
					"modified_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"version": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingRouteVersionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"active": schema.StringAttribute{
								Description: `Available values: "true", "false".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("true", "false"),
								},
							},
							"created_at": schema.StringAttribute{
								Computed: true,
							},
							"data": schema.StringAttribute{
								Computed: true,
							},
							"version_id": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"version": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingVersionModel](ctx),
				Attributes: map[string]schema.Attribute{
					"active": schema.StringAttribute{
						Description: `Available values: "true", "false".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("true", "false"),
						},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"data": schema.StringAttribute{
						Computed: true,
					},
					"version_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *AIGatewayDynamicRoutingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AIGatewayDynamicRoutingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
