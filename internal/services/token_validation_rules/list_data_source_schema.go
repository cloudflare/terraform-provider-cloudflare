// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_rules

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

var _ datasource.DataSourceWithConfigValidators = (*TokenValidationRulesListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"action": schema.StringAttribute{
				Description: "Action to take on requests that match operations included in `selector` and fail `expression`.\nAvailable values: \"log\", \"block\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("log", "block"),
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Toggle rule on or off.",
				Optional:    true,
			},
			"host": schema.StringAttribute{
				Description: "Select rules with this host in `include`.",
				Optional:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "Select rules with this host in `include`.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Select rules with these IDs.",
				Optional:    true,
			},
			"rule_id": schema.StringAttribute{
				Description: "Select rules with these IDs.",
				Optional:    true,
			},
			"token_configuration": schema.ListAttribute{
				Description: "Select rules using any of these token configurations.",
				Optional:    true,
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[TokenValidationRulesListResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Description: "Action to take on requests that match operations included in `selector` and fail `expression`.\nAvailable values: \"log\", \"block\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("log", "block"),
							},
						},
						"description": schema.StringAttribute{
							Description: "A human-readable description that gives more details than `title`.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Toggle rule on or off.",
							Computed:    true,
						},
						"expression": schema.StringAttribute{
							Description: "Rule expression. Requests that fail to match this expression will be subject to `action`.\n\nFor details on expressions, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).",
							Computed:    true,
						},
						"selector": schema.SingleNestedAttribute{
							Description: "Select operations covered by this rule.\n\nFor details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[TokenValidationRulesListSelectorDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"exclude": schema.ListNestedAttribute{
									Description: "Ignore operations that were otherwise included by `include`.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[TokenValidationRulesListSelectorExcludeDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"operation_ids": schema.ListAttribute{
												Description: "Excluded operation IDs.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
								"include": schema.ListNestedAttribute{
									Description: "Select all matching operations.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[TokenValidationRulesListSelectorIncludeDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"host": schema.ListAttribute{
												Description: "Included hostnames.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
							},
						},
						"title": schema.StringAttribute{
							Description: "A human-readable name for the rule.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "UUID.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"last_updated": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *TokenValidationRulesListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *TokenValidationRulesListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
