// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &FirewallRuleDataSource{}
var _ datasource.DataSourceWithValidateConfig = &FirewallRuleDataSource{}

func (r FirewallRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"path_id": schema.StringAttribute{
				Description: "The unique identifier of the firewall rule.",
				Optional:    true,
			},
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"query_id": schema.StringAttribute{
				Description: "The unique identifier of the firewall rule.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the firewall rule.",
				Optional:    true,
			},
			"action": schema.StringAttribute{
				Description: "The action to apply to a matched request. The `log` action is only available on an Enterprise plan.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge", "allow", "log", "bypass"),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The unique identifier of the filter.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the filter.",
						Optional:    true,
					},
					"expression": schema.StringAttribute{
						Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
						Optional:    true,
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the filter is currently paused.",
						Optional:    true,
					},
					"ref": schema.StringAttribute{
						Description: "A short reference tag. Allows you to select related filters.",
						Optional:    true,
					},
					"deleted": schema.BoolAttribute{
						Description: "When true, indicates that the firewall rule was deleted.",
						Optional:    true,
					},
				},
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the firewall rule is currently paused.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the firewall rule.",
				Optional:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.",
				Optional:    true,
			},
			"products": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("zoneLockdown", "uaBlock", "bic", "hot", "securityLevel", "rateLimit", "waf"),
				},
			},
			"ref": schema.StringAttribute{
				Description: "A short reference tag. Allows you to select related firewall rules.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"id": schema.StringAttribute{
						Description: "The unique identifier of the firewall rule.",
						Optional:    true,
					},
					"action": schema.StringAttribute{
						Description: "The action to search for. Must be an exact match.",
						Optional:    true,
					},
					"description": schema.StringAttribute{
						Description: "A case-insensitive string to find in the description.",
						Optional:    true,
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"paused": schema.BoolAttribute{
						Description: "When true, indicates that the firewall rule is currently paused.",
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of firewall rules per page.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *FirewallRuleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *FirewallRuleDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
