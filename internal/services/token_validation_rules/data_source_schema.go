// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*TokenValidationRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID.",
				Computed:    true,
			},
			"rule_id": schema.StringAttribute{
				Description: "UUID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"action": schema.StringAttribute{
				Description: "Action to take on requests that match operations included in `selector` and fail `expression`.\nAvailable values: \"log\", \"block\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("log", "block"),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
			"last_updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Computed: true,
			},
			"title": schema.StringAttribute{
				Description: "A human-readable name for the rule.",
				Computed:    true,
			},
			"selector": schema.SingleNestedAttribute{
				Description: "Select operations covered by this rule.\n\nFor details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TokenValidationRulesSelectorDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"exclude": schema.ListNestedAttribute{
						Description: "Ignore operations that were otherwise included by `include`.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[TokenValidationRulesSelectorExcludeDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[TokenValidationRulesSelectorIncludeDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Select rules with these IDs.",
						Optional:    true,
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
					"rule_id": schema.StringAttribute{
						Description: "Select rules with these IDs.",
						Optional:    true,
					},
					"token_configuration": schema.ListAttribute{
						Description: "Select rules using any of these token configurations.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *TokenValidationRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TokenValidationRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("rule_id"), path.MatchRoot("filter")),
	}
}
