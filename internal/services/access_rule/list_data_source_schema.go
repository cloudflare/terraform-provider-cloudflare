// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
				Description: "The direction used to sort returned rules.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.",
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
				Description: "The string to search for in the notes of existing IP Access rules.\nNotes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.",
				Optional:    true,
			},
			"order": schema.StringAttribute{
				Description: "The field used to sort returned rules.",
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
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"target": schema.StringAttribute{
						Description: "The target to search in existing rules.",
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
						Description: "The target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`.\nNotes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.",
						Optional:    true,
					},
				},
			},
			"match": schema.StringAttribute{
				Description: "When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.",
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
									Description: "The configuration target. You must set the target to `ip` when specifying an IP address in the rule.",
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
									Description: "The IP address to match. This address will be compared to the IP address of incoming requests.",
									Computed:    true,
								},
							},
						},
						"mode": schema.StringAttribute{
							Description: "The action to apply to a matched request.",
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
									Description: "Identifier",
									Computed:    true,
								},
								"email": schema.StringAttribute{
									Description: "The contact email address of the user.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The scope of the rule.",
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
