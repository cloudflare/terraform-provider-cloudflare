// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monetization_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MonetizationRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Required:    true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "The zone's payment rules, in the order they are evaluated. Empty when the zone has no payment rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MonetizationRulesRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also accepted. The address must pass wallet screening whenever its rule is deployed or patched.",
							Computed:    true,
						},
						"expression": schema.StringAttribute{
							Description: "Wirefilter expression identifying the requests that require payment. Forwarded verbatim to the Rulesets API, which validates its syntax.",
							Computed:    true,
						},
						"price": schema.StringAttribute{
							Description: `Price in the smallest indivisible unit of the configured payment token, encoded as a decimal string. Must be a canonical decimal integer in [1000, 100000000]: no leading zeros, and a bare JSON number is rejected. The price is required for the fixed-price schemes ("exact" and "upto") and must be at least 1000, the smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the scheme is "origin_controlled" the origin server sets pricing dynamically and the rule carries no price: the field must be omitted — any provided value, including "0", is rejected because it would not be enforced. Responses always encode this as a string and omit it for "origin_controlled" rules; the pattern matches exactly the set of values the server accepts.`,
							Computed:    true,
						},
						"scheme": schema.StringAttribute{
							Description: "X402 payment scheme. \"exact\" requires the specified payment amount; \"upto\" permits a payment up to the specified amount. Both fixed-price schemes require the price field.\nAvailable values: \"exact\", \"upto\", \"origin_controlled\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"exact",
									"upto",
									"origin_controlled",
								),
							},
						},
						"id": schema.StringAttribute{
							Description: "The server-assigned unique ID of the payment rule. Stable across full-ruleset replacements.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *MonetizationRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MonetizationRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
