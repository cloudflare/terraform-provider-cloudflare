// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AISearchNamespacesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"search": schema.StringAttribute{
				Description: "Filter namespaces whose name or description contains this string (case-insensitive).",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AISearchNamespacesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Description: "Optional description for the namespace. Max 256 characters.",
							Computed:    true,
						},
						"public_endpoint_id": schema.StringAttribute{
							Computed: true,
						},
						"public_endpoint_params": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AISearchNamespacesPublicEndpointParamsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"authorized_hosts": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"chat_completions_endpoint": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchNamespacesPublicEndpointParamsChatCompletionsEndpointDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"disabled": schema.BoolAttribute{
											Description: "Disable chat completions endpoint for this public endpoint",
											Computed:    true,
										},
									},
								},
								"custom_domains": schema.ListAttribute{
									Description: "Custom domain hostnames that alias this public endpoint. GET and create responses return the current set; on update (PUT) this field is only echoed back when supplied in the request body, otherwise it is null (omit it to leave domains unchanged).",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"default_domain_enabled": schema.BoolAttribute{
									Description: "When false, the instance is reachable only via a registered custom domain and the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404. Requires at least one custom domain. Defaults to true. public_endpoint_params is replaced wholesale on update, so resend default_domain_enabled on every update to keep the default host off — omitting it resets to true.",
									Computed:    true,
								},
								"enabled": schema.BoolAttribute{
									Computed: true,
								},
								"instances_allowed": schema.ListAttribute{
									Description: "Instance IDs exposed through the namespace public endpoint. Empty means nothing is searchable. Every ID must be an existing instance in this namespace, and the list cannot exceed the account's multi-instance search limit.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"mcp": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchNamespacesPublicEndpointParamsMcpDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"disabled": schema.BoolAttribute{
											Description: "Disable MCP endpoint for this public endpoint",
											Computed:    true,
										},
									},
								},
								"rate_limit": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchNamespacesPublicEndpointParamsRateLimitDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"period_ms": schema.Int64Attribute{
											Computed: true,
											Validators: []validator.Int64{
												int64validator.Between(60000, 3600000),
											},
										},
										"requests": schema.Int64Attribute{
											Computed: true,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},
										"technique": schema.StringAttribute{
											Description: `Available values: "fixed", "sliding".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
											},
										},
									},
								},
								"search_endpoint": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AISearchNamespacesPublicEndpointParamsSearchEndpointDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"disabled": schema.BoolAttribute{
											Description: "Disable search endpoint for this public endpoint",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *AISearchNamespacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AISearchNamespacesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
