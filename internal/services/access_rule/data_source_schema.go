// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &AccessRuleDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessRuleDataSource{}

func (r AccessRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
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
					"egs_pagination": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"json": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"page": schema.Float64Attribute{
										Description: "The page number of paginated results.",
										Computed:    true,
										Optional:    true,
									},
									"per_page": schema.Float64Attribute{
										Description: "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
						},
					},
					"filters": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"configuration_target": schema.StringAttribute{
								Description: "The target to search in existing rules.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("ip", "ip_range", "asn", "country"),
								},
							},
							"configuration_value": schema.StringAttribute{
								Description: "The target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`.\nNotes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.",
								Optional:    true,
							},
							"match": schema.StringAttribute{
								Description: "When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("any", "all"),
								},
							},
							"mode": schema.StringAttribute{
								Description: "The action to apply to a matched request.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("block", "challenge", "whitelist", "js_challenge", "managed_challenge"),
								},
							},
							"notes": schema.StringAttribute{
								Description: "The string to search for in the notes of existing IP Access rules.\nNotes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.",
								Optional:    true,
							},
						},
					},
					"order": schema.StringAttribute{
						Description: "The field used to sort returned rules.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("configuration.target", "configuration.value", "mode"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Requested page within paginated list of results.",
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Maximum number of results requested.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *AccessRuleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessRuleDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
