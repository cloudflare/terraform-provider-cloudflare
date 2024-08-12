// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &FirewallRuleDataSource{}

func (d *FirewallRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"action": schema.StringAttribute{
				Description: "The action to apply to a matched request. The `log` action is only available on an Enterprise plan.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"block",
						"challenge",
						"js_challenge",
						"managed_challenge",
						"allow",
						"log",
						"bypass",
					),
				},
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the firewall rule.",
				Computed:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the firewall rule is currently paused.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the firewall rule.",
				Computed:    true,
				Optional:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 2147483647),
				},
			},
			"ref": schema.StringAttribute{
				Description: "A short reference tag. Allows you to select related firewall rules.",
				Computed:    true,
				Optional:    true,
			},
			"products": schema.ListAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"zoneLockdown",
							"uaBlock",
							"bic",
							"hot",
							"securityLevel",
							"rateLimit",
							"waf",
						),
					),
				},
				ElementType: types.StringType,
			},
		},
	}
}

func (d *FirewallRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
