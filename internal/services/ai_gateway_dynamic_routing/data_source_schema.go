// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_dynamic_routing

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AIGatewayDynamicRoutingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"gateway_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
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
			"name": schema.StringAttribute{
				Computed: true,
			},
			"deployment": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingDeploymentDataSourceModel](ctx),
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
				CustomType: customfield.NewNestedObjectListType[AIGatewayDynamicRoutingElementsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"outputs": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsOutputsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"next": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsOutputsNextDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"false": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsOutputsFalseDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"true": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsOutputsTrueDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsOutputsFallbackDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"element_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"success": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsOutputsSuccessDataSourceModel](ctx),
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
							CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingElementsPropertiesDataSourceModel](ctx),
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
			"version": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayDynamicRoutingVersionDataSourceModel](ctx),
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

func (d *AIGatewayDynamicRoutingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AIGatewayDynamicRoutingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
