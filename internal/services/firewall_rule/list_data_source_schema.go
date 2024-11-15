// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*FirewallRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"action": schema.StringAttribute{
				Description: "The action to search for. Must be an exact match.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A case-insensitive string to find in the description.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the firewall rule.",
				Optional:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the firewall rule is currently paused.",
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
				CustomType:  customfield.NewNestedObjectListType[FirewallRulesResultDataSourceModel](ctx),
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
						"filter": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[FirewallRulesFilterDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The unique identifier of the filter.",
									Computed:    true,
								},
								"description": schema.StringAttribute{
									Description: "An informative summary of the filter.",
									Computed:    true,
								},
								"expression": schema.StringAttribute{
									Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
									Computed:    true,
								},
								"paused": schema.BoolAttribute{
									Description: "When true, indicates that the filter is currently paused.",
									Computed:    true,
								},
								"ref": schema.StringAttribute{
									Description: "A short reference tag. Allows you to select related filters.",
									Computed:    true,
								},
								"deleted": schema.BoolAttribute{
									Description: "When true, indicates that the firewall rule was deleted.",
									Computed:    true,
								},
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
							Validators: []validator.Float64{
								float64validator.Between(0, 2147483647),
							},
						},
						"products": schema.ListAttribute{
							Computed: true,
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
							CustomType:  customfield.NewListType[types.String](ctx),
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

func (d *FirewallRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *FirewallRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
