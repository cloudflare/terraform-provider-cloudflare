// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UserAgentBlockingRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A string to search for in the description of existing rules.",
				Optional:    true,
			},
			"description_search": schema.StringAttribute{
				Description: "A string to search for in the description of existing rules.",
				Optional:    true,
			},
			"ua_search": schema.StringAttribute{
				Description: "A string to search for in the user agent values of existing rules.",
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
				CustomType:  customfield.NewNestedObjectListType[UserAgentBlockingRulesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the User Agent Blocking rule.",
							Computed:    true,
						},
						"configuration": schema.SingleNestedAttribute{
							Description: "The configuration object for the current rule.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[UserAgentBlockingRulesConfigurationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"target": schema.StringAttribute{
									Description: "The configuration target for this rule. You must set the target to `ua` for User Agent Blocking rules.",
									Computed:    true,
								},
								"value": schema.StringAttribute{
									Description: "The exact user agent string to match. This value will be compared to the received `User-Agent` HTTP header value.",
									Computed:    true,
								},
							},
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the rule.",
							Computed:    true,
						},
						"mode": schema.StringAttribute{
							Description: "The action to apply to a matched request.\navailable values: \"block\", \"challenge\", \"js_challenge\", \"managed_challenge\"",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"block",
									"challenge",
									"js_challenge",
									"managed_challenge",
								),
							},
						},
						"paused": schema.BoolAttribute{
							Description: "When true, indicates that the rule is currently paused.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UserAgentBlockingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *UserAgentBlockingRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
