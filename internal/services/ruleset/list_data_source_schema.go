// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"regexp"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*RulesetsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The unique ID of the account.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.MatchRoot("zone_id")),
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			"zone_id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile("^[0-9a-f]{32}$"),
						"value must be a 32-character hexadecimal string",
					),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Maximum number of rulesets to fetch (defaults to 1000).",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"rulesets": schema.SetNestedAttribute{
				Description: "A list of rulesets. The returned information will not include the rules in each ruleset.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectSetType[RulesetsRulesetDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the ruleset.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(
									regexp.MustCompile("^[0-9a-f]{32}$"),
									"value must be a 32-character hexadecimal string",
								),
							},
						},
						"kind": schema.StringAttribute{
							Description: "The kind of the ruleset.\nAvailable values: \"managed\", \"custom\", \"root\", \"zone\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"managed",
									"custom",
									"root",
									"zone",
								),
							},
						},
						"name": schema.StringAttribute{
							Description: "The human-readable name of the ruleset.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"phase": schema.StringAttribute{
							Description: "The phase of the ruleset.\nAvailable values: \"ddos_l4\", \"ddos_l7\", \"http_config_settings\", \"http_custom_errors\", \"http_log_custom_fields\", \"http_ratelimit\", \"http_request_cache_settings\", \"http_request_dynamic_redirect\", \"http_request_firewall_custom\", \"http_request_firewall_managed\", \"http_request_late_transform\", \"http_request_origin\", \"http_request_redirect\", \"http_request_sanitize\", \"http_request_sbfm\", \"http_request_transform\", \"http_response_compression\", \"http_response_firewall_managed\", \"http_response_headers_transform\", \"magic_transit\", \"magic_transit_ids_managed\", \"magic_transit_managed\", \"magic_transit_ratelimit\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ddos_l4",
									"ddos_l7",
									"http_config_settings",
									"http_custom_errors",
									"http_log_custom_fields",
									"http_ratelimit",
									"http_request_cache_settings",
									"http_request_dynamic_redirect",
									"http_request_firewall_custom",
									"http_request_firewall_managed",
									"http_request_late_transform",
									"http_request_origin",
									"http_request_redirect",
									"http_request_sanitize",
									"http_request_sbfm",
									"http_request_transform",
									"http_response_compression",
									"http_response_firewall_managed",
									"http_response_headers_transform",
									"magic_transit",
									"magic_transit_ids_managed",
									"magic_transit_managed",
									"magic_transit_ratelimit",
								),
							},
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the ruleset.",
							Computed:    true,
						},
						"last_updated": schema.StringAttribute{
							Description: "The timestamp of when the ruleset was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"version": schema.StringAttribute{
							Description: "The version of the ruleset.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(
									regexp.MustCompile("^[0-9]+$"),
									"value must be a non-empty string containing only numbers",
								),
							},
						},
					},
				},
			},
			"result": schema.SetNestedAttribute{
				Description:        "A list of rulesets. The returned information will not include the rules in each ruleset.",
				Computed:           true,
				DeprecationMessage: "Use rulesets instead. This attribute will be removed in the next major version of the provider.",
				CustomType:         customfield.NewNestedObjectSetType[RulesetsRulesetDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:        "The unique ID of the ruleset.",
							Computed:           true,
							DeprecationMessage: "Use rulesets.id instead. This attribute will be removed in the next major version of the provider.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(
									regexp.MustCompile("^[0-9a-f]{32}$"),
									"value must be a 32-character hexadecimal string",
								),
							},
						},
						"kind": schema.StringAttribute{
							Description:        "The kind of the ruleset.\nAvailable values: \"managed\", \"custom\", \"root\", \"zone\".",
							Computed:           true,
							DeprecationMessage: "Use rulesets.kind instead. This attribute will be removed in the next major version of the provider.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"managed",
									"custom",
									"root",
									"zone",
								),
							},
						},
						"name": schema.StringAttribute{
							Description:        "The human-readable name of the ruleset.",
							Computed:           true,
							DeprecationMessage: "Use rulesets.name instead. This attribute will be removed in the next major version of the provider.",
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"phase": schema.StringAttribute{
							Description:        "The phase of the ruleset.\nAvailable values: \"ddos_l4\", \"ddos_l7\", \"http_config_settings\", \"http_custom_errors\", \"http_log_custom_fields\", \"http_ratelimit\", \"http_request_cache_settings\", \"http_request_dynamic_redirect\", \"http_request_firewall_custom\", \"http_request_firewall_managed\", \"http_request_late_transform\", \"http_request_origin\", \"http_request_redirect\", \"http_request_sanitize\", \"http_request_sbfm\", \"http_request_transform\", \"http_response_compression\", \"http_response_firewall_managed\", \"http_response_headers_transform\", \"magic_transit\", \"magic_transit_ids_managed\", \"magic_transit_managed\", \"magic_transit_ratelimit\".",
							Computed:           true,
							DeprecationMessage: "Use rulesets.phase instead. This attribute will be removed in the next major version of the provider.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ddos_l4",
									"ddos_l7",
									"http_config_settings",
									"http_custom_errors",
									"http_log_custom_fields",
									"http_ratelimit",
									"http_request_cache_settings",
									"http_request_dynamic_redirect",
									"http_request_firewall_custom",
									"http_request_firewall_managed",
									"http_request_late_transform",
									"http_request_origin",
									"http_request_redirect",
									"http_request_sanitize",
									"http_request_sbfm",
									"http_request_transform",
									"http_response_compression",
									"http_response_firewall_managed",
									"http_response_headers_transform",
									"magic_transit",
									"magic_transit_ids_managed",
									"magic_transit_managed",
									"magic_transit_ratelimit",
								),
							},
						},
						"description": schema.StringAttribute{
							Description:        "An informative description of the ruleset.",
							Computed:           true,
							DeprecationMessage: "Use rulesets.description instead. This attribute will be removed in the next major version of the provider.",
						},
					},
				},
			},
		},
	}
}

func (d *RulesetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *RulesetsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
