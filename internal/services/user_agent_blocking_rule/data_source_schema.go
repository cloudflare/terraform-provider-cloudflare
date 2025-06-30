// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UserAgentBlockingRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique identifier of the User Agent Blocking rule.",
				Computed:    true,
			},
			"ua_rule_id": schema.StringAttribute{
				Description: "The unique identifier of the User Agent Blocking rule.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Defines an identifier.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the rule.",
				Computed:    true,
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.\nAvailable values: \"block\", \"challenge\", \"js_challenge\", \"managed_challenge\".",
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
			"configuration": schema.SingleNestedAttribute{
				Description: "The configuration object for the current rule.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[UserAgentBlockingRuleConfigurationDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "A string to search for in the description of existing rules.",
						Optional:    true,
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the rule is currently paused.",
						Optional:    true,
					},
					"user_agent": schema.StringAttribute{
						Description: "A string to search for in the user agent values of existing rules.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *UserAgentBlockingRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *UserAgentBlockingRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("ua_rule_id"), path.MatchRoot("filter")),
	}
}
