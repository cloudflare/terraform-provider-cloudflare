// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &FirewallRulesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &FirewallRulesDataSource{}

func (r FirewallRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
							Description: "The unique identifier of the firewall rule.",
							Computed:    true,
						},
						"action": schema.StringAttribute{
							Description: "The action to apply to a matched request. The `log` action is only available on an Enterprise plan.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("block", "challenge", "js_challenge", "managed_challenge", "allow", "log", "bypass"),
							},
						},
						"paused": schema.BoolAttribute{
							Description: "When true, indicates that the firewall rule is currently paused.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the firewall rule.",
							Computed:    true,
						},
						"priority": schema.Float64Attribute{
							Description: "The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.",
							Computed:    true,
						},
						"products": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"ref": schema.StringAttribute{
							Description: "A short reference tag. Allows you to select related firewall rules.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *FirewallRulesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *FirewallRulesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
