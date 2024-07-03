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
var _ datasource.DataSourceWithValidateConfig = &UserAgentBlockingRulesDataSource{}

func (r UserAgentBlockingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"page": schema.Float64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.",
				Computed:    true,
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
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the User Agent Blocking rule.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the rule.",
							Computed:    true,
						},
						"mode": schema.StringAttribute{
							Description: "The action to apply to a matched request.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge"),
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

func (r *UserAgentBlockingRulesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *UserAgentBlockingRulesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
