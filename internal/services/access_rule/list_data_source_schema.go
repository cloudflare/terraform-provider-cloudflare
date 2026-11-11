// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AccessRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Firewall Access Rules Read",
				"Account Firewall Access Rules Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Defines the direction used to sort returned rules.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.\nAvailable values: \"block\", \"challenge\", \"whitelist\", \"js_challenge\", \"managed_challenge\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"block",
						"challenge",
						"whitelist",
						"js_challenge",
						"managed_challenge",
					),
				},
			},
			"notes": schema.StringAttribute{
				Description: "Defines the string to search for in the notes of existing IP Access rules.\nNotes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.",
				Optional:    true,
			},
			"order": schema.StringAttribute{
				Description: "Defines the field used to sort returned rules.\nAvailable values: \"configuration.target\", \"configuration.value\", \"mode\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"configuration.target",
						"configuration.value",
						"mode",
					),
				},
			},
			"configuration": schema.SingleNestedAttribute{
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AccessRulesFilterConfigurationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"target": schema.StringAttribute{
						Description: "Defines the target to search in existing rules.\nAvailable values: \"ip\", \"ip_range\", \"asn\", \"country\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ip",
								"ip_range",
								"asn",
								"country",
							),
						},
					},
					"value": schema.StringAttribute{
						Description: "Defines the target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`.\nNotes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.",
						Optional:    true,
					},
				},
			},
			"match": schema.StringAttribute{
				Description: "Defines the search requirements. When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.\nAvailable values: \"any\", \"all\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("any", "all"),
				},
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
				CustomType:  customfield.NewNestedObjectListType[AccessRulesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the IP Access rule.",
							Computed:    true,
						},
						"allowed_modes": schema.ListAttribute{
							Description: "The available actions that a rule can apply to a matched request.",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"block",
										"challenge",
										"whitelist",
										"js_challenge",
										"managed_challenge",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"configuration": schema.SingleNestedAttribute{
							Description: "The rule configuration.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[AccessRulesConfigurationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"target": schema.StringAttribute{
									Description: "For `target: ip`: The configuration target. You must set the target to `ip` when specifying an IP address in the rule.\n\nFor `target: ip6`: The configuration target. You must set the target to `ip6` when specifying an IPv6 address in the rule.\n\nFor `target: ip_range`: The configuration target. You must set the target to `ip_range` when specifying an IP address range in the rule.\n\nFor `target: asn`: The configuration target. You must set the target to `asn` when specifying an Autonomous System Number (ASN) in the rule.\n\nFor `target: country`: The configuration target. You must set the target to `country` when specifying a country code in the rule.\nAvailable values: \"ip\", \"ip6\", \"ip_range\", \"asn\", \"country\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"ip",
											"ip6",
											"ip_range",
											"asn",
											"country",
										),
									},
								},
								"value": schema.StringAttribute{
									Description: "For `target: ip`: The IP address to match. This address will be compared to the IP address of incoming requests.\n\nFor `target: ip6`: The IPv6 address to match.\n\nFor `target: ip_range`: The IP address range to match. You can only use prefix lengths `/16` and `/24` for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.\n\nFor `target: asn`: The AS number to match.\n\nFor `target: country`: The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).",
									Computed:    true,
								},
							},
						},
						"mode": schema.StringAttribute{
							Description: "The action to apply to a matched request.\nAvailable values: \"block\", \"challenge\", \"whitelist\", \"js_challenge\", \"managed_challenge\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"block",
									"challenge",
									"whitelist",
									"js_challenge",
									"managed_challenge",
								),
							},
						},
						"created_on": schema.StringAttribute{
							Description: "The timestamp of when the rule was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"modified_on": schema.StringAttribute{
							Description: "The timestamp of when the rule was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"notes": schema.StringAttribute{
							Description: "An informative summary of the rule, typically used as a reminder or explanation.",
							Computed:    true,
						},
						"scope": schema.SingleNestedAttribute{
							Description: "All zones owned by the user will have the rule applied.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[AccessRulesScopeDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Defines an identifier.",
									Computed:    true,
								},
								"email": schema.StringAttribute{
									Description: "The contact email address of the user.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "Defines the scope of the rule.\nAvailable values: \"user\", \"organization\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("user", "organization"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *AccessRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccessRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
