// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &UserAgentBlockingRuleDataSource{}

func (d *UserAgentBlockingRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the User Agent Blocking rule.",
				Computed:    true,
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the rule.",
				Optional:    true,
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge"),
				},
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the rule is currently paused.",
				Optional:    true,
			},
			"configuration": schema.SingleNestedAttribute{
				Description: "The configuration object for the current rule.",
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
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
						Validators: []validator.Float64{
							float64validator.AtLeast(1),
						},
					},
					"per_page": schema.Float64Attribute{
						Description: "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(1, 1000),
						},
					},
					"ua_search": schema.StringAttribute{
						Description: "A string to search for in the user agent values of existing rules.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *UserAgentBlockingRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
