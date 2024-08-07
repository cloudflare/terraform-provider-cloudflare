// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &UserAgentBlockingRulesDataSource{}

func (d *UserAgentBlockingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
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
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the User Agent Blocking rule.",
							Computed:    true,
						},
						"configuration": schema.SingleNestedAttribute{
							Description: "The configuration object for the current rule.",
							Computed:    true,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"target": schema.StringAttribute{
									Description: "The configuration target for this rule. You must set the target to `ua` for User Agent Blocking rules.",
									Computed:    true,
									Optional:    true,
								},
								"value": schema.StringAttribute{
									Description: "The exact user agent string to match. This value will be compared to the received `User-Agent` HTTP header value.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the rule.",
							Computed:    true,
							Optional:    true,
						},
						"mode": schema.StringAttribute{
							Description: "The action to apply to a matched request.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge"),
							},
						},
						"paused": schema.BoolAttribute{
							Description: "When true, indicates that the rule is currently paused.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UserAgentBlockingRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
