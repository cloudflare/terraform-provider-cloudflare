// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"regexp"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*RulesetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"kind": schema.StringAttribute{
				Description: "The kind of the ruleset.\nAvailable values: \"managed\", \"custom\", \"root\", \"zone\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"managed",
						"custom",
						"root",
						"zone",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The human-readable name of the ruleset.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"phase": schema.StringAttribute{
				Description: "The phase of the ruleset.\nAvailable values: \"ddos_l4\", \"ddos_l7\", \"http_config_settings\", \"http_custom_errors\", \"http_log_custom_fields\", \"http_ratelimit\", \"http_request_cache_settings\", \"http_request_dynamic_redirect\", \"http_request_firewall_custom\", \"http_request_firewall_managed\", \"http_request_late_transform\", \"http_request_origin\", \"http_request_redirect\", \"http_request_sanitize\", \"http_request_sbfm\", \"http_request_transform\", \"http_response_compression\", \"http_response_firewall_managed\", \"http_response_headers_transform\", \"magic_transit\", \"magic_transit_ids_managed\", \"magic_transit_managed\", \"magic_transit_ratelimit\".",
				Required:    true,
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "An informative description of the ruleset.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"rules": schema.ListNestedAttribute{
				Description: "The list of rules in the ruleset.",
				Computed:    true,
				Optional:    true,
				Default:     listdefault.StaticValue(customfield.NewObjectListMust(ctx, []RulesetRulesModel{}).ListValue),
				CustomType:  customfield.NewNestedObjectListType[RulesetRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the rule.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(
									regexp.MustCompile("^[0-9a-f]{32}$"),
									"value must be a 32-character hexadecimal string",
								),
							},
						},
						"action": schema.StringAttribute{
							Description: "The action to perform when the rule matches.\nAvailable values: \"block\", \"challenge\", \"compress_response\", \"ddos_dynamic\", \"execute\", \"force_connection_close\", \"js_challenge\", \"log\", \"log_custom_field\", \"managed_challenge\", \"redirect\", \"rewrite\", \"route\", \"score\", \"serve_error\", \"set_cache_settings\", \"set_config\", \"skip\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"block",
									"challenge",
									"compress_response",
									"ddos_dynamic",
									"execute",
									"force_connection_close",
									"js_challenge",
									"log",
									"log_custom_field",
									"managed_challenge",
									"redirect",
									"rewrite",
									"route",
									"score",
									"serve_error",
									"set_cache_settings",
									"set_config",
									"skip",
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile("^[a-z_]+$"),
									"value must be a non-empty string containing only lowercase characters and underscores",
								),
							},
						},
						"action_parameters": schema.SingleNestedAttribute{
							Description: "The parameters configuring the rule's action.",
							Computed:    true,
							Optional:    true,
							Default:     objectdefault.StaticValue(customfield.NewObjectMust(ctx, &RulesetRulesActionParametersModel{}).ObjectValue),
							CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersModel](ctx),
							Attributes: map[string]schema.Attribute{
								"response": schema.SingleNestedAttribute{
									Description: "The response to show when the block is applied.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"block",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersResponseModel](ctx),
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description: "The content to return.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
										"content_type": schema.StringAttribute{
											Description: "The type of the content to return.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
										"status_code": schema.Int64Attribute{
											Description: "The status code to return.",
											Required:    true,
											Validators: []validator.Int64{
												int64validator.Between(400, 499),
											},
										},
									},
								},
								"algorithms": schema.ListNestedAttribute{
									Description: "Custom order for compression algorithms.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"compress_response",
										),
										listvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersAlgorithmsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "Name of the compression algorithm to enable.\nAvailable values: \"none\", \"auto\", \"default\", \"gzip\", \"brotli\", \"zstd\".",
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"none",
														"auto",
														"default",
														"gzip",
														"brotli",
														"zstd",
													),
												},
											},
										},
									},
								},
								"id": schema.StringAttribute{
									Description: "The ID of the ruleset to execute.",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"execute",
										),
										stringvalidator.RegexMatches(
											regexp.MustCompile("^[0-9a-f]{32}$"),
											"value must be a 32-character hexadecimal string",
										),
									},
								},
								"matched_data": schema.SingleNestedAttribute{
									Description: "The configuration to use for matched data logging.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"execute",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersMatchedDataModel](ctx),
									Attributes: map[string]schema.Attribute{
										"public_key": schema.StringAttribute{
											Description: "The public key to encrypt matched data logs with.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
								},
								"overrides": schema.SingleNestedAttribute{
									Description: "A set of overrides to apply to the target ruleset.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"execute",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersOverridesModel](ctx),
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description: "An action to override all rules with. This option has lower precedence than rule and category overrides.",
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.AtLeastOneOf(
													path.MatchRelative().AtParent().AtName("categories"),
													path.MatchRelative().AtParent().AtName("enabled"),
													path.MatchRelative().AtParent().AtName("rules"),
													path.MatchRelative().AtParent().AtName("sensitivity_level"),
												),
												stringvalidator.RegexMatches(
													regexp.MustCompile("^[a-z_]+$"),
													"value must be a non-empty string containing only lowercase characters and underscores",
												),
											},
										},
										"categories": schema.ListNestedAttribute{
											Description: "A list of category-level overrides. This option has the second-highest precedence after rule-level overrides.",
											Optional:    true,
											Validators: []validator.List{
												listvalidator.SizeAtLeast(1),
											},
											CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersOverridesCategoriesModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"category": schema.StringAttribute{
														Description: "The name of the category to override.",
														Required:    true,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
													"action": schema.StringAttribute{
														Description: "The action to override rules in the category with.",
														Optional:    true,
														Validators: []validator.String{
															stringvalidator.AtLeastOneOf(
																path.MatchRelative().AtParent().AtName("enabled"),
																path.MatchRelative().AtParent().AtName("sensitivity_level"),
															),
															stringvalidator.RegexMatches(
																regexp.MustCompile("^[a-z_]+$"),
																"value must be a non-empty string containing only lowercase characters and underscores",
															),
														},
													},
													"enabled": schema.BoolAttribute{
														Description: "Whether to enable execution of rules in the category.",
														Optional:    true,
													},
													"sensitivity_level": schema.StringAttribute{
														Description: "The sensitivity level to use for rules in the category. This option is only applicable for DDoS phases.\nAvailable values: \"default\", \"medium\", \"low\", \"eoff\".",
														Optional:    true,
														Validators: []validator.String{
															stringvalidator.OneOfCaseInsensitive(
																"default",
																"medium",
																"low",
																"eoff",
															),
														},
													},
												},
											},
										},
										"enabled": schema.BoolAttribute{
											Description: "Whether to enable execution of all rules. This option has lower precedence than rule and category overrides.",
											Optional:    true,
										},
										"rules": schema.ListNestedAttribute{
											Description: "A list of rule-level overrides. This option has the highest precedence.",
											Optional:    true,
											Validators: []validator.List{
												listvalidator.SizeAtLeast(1),
											},
											CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersOverridesRulesModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "The ID of the rule to override.",
														Required:    true,
														Validators: []validator.String{
															stringvalidator.RegexMatches(
																regexp.MustCompile("^[0-9a-f]{32}$"),
																"value must be a 32-character hexadecimal string",
															),
														},
													},
													"action": schema.StringAttribute{
														Description: "The action to override the rule with.",
														Optional:    true,
														Validators: []validator.String{
															stringvalidator.AtLeastOneOf(
																path.MatchRelative().AtParent().AtName("enabled"),
																path.MatchRelative().AtParent().AtName("score_threshold"),
																path.MatchRelative().AtParent().AtName("sensitivity_level"),
															),
															stringvalidator.RegexMatches(
																regexp.MustCompile("^[a-z_]+$"),
																"value must be a non-empty string containing only lowercase characters and underscores",
															),
														},
													},
													"enabled": schema.BoolAttribute{
														Description: "Whether to enable execution of the rule.",
														Optional:    true,
													},
													"score_threshold": schema.Int64Attribute{
														Description: "The score threshold to use for the rule.",
														Optional:    true,
													},
													"sensitivity_level": schema.StringAttribute{
														Description: "The sensitivity level to use for the rule. This option is only applicable for DDoS phases.\nAvailable values: \"default\", \"medium\", \"low\", \"eoff\".",
														Optional:    true,
														Validators: []validator.String{
															stringvalidator.OneOfCaseInsensitive(
																"default",
																"medium",
																"low",
																"eoff",
															),
														},
													},
												},
											},
										},
										"sensitivity_level": schema.StringAttribute{
											Description: "A sensitivity level to set for all rules. This option has lower precedence than rule and category overrides and is only applicable for DDoS phases.\nAvailable values: \"default\", \"medium\", \"low\", \"eoff\".",
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"default",
													"medium",
													"low",
													"eoff",
												),
											},
										},
									},
								},
								"from_list": schema.SingleNestedAttribute{
									Description: "A redirect based on a bulk list lookup.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"redirect",
										),
										objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("from_value")),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersFromListModel](ctx),
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description: "An expression that evaluates to the list lookup key.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
										"name": schema.StringAttribute{
											Description: "The name of the list to match against.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.RegexMatches(
													regexp.MustCompile("^[a-zA-Z0-9_]+$"),
													"value must be a non-empty string containing only alphanumeric characters and underscores",
												),
											},
										},
									},
								},
								"from_value": schema.SingleNestedAttribute{
									Description: "A redirect based on the request properties.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"redirect",
										),
										objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("from_list")),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersFromValueModel](ctx),
									Attributes: map[string]schema.Attribute{
										"preserve_query_string": schema.BoolAttribute{
											Description: "Whether to keep the query string of the original request.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"status_code": schema.Int64Attribute{
											Description: "The status code to use for the redirect.",
											Optional:    true,
											Validators: []validator.Int64{
												int64validator.OneOf(
													301,
													302,
													303,
													307,
													308,
												),
											},
										},
										"target_url": schema.SingleNestedAttribute{
											Description: "A URL to redirect the request to.",
											Required:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersFromValueTargetURLModel](ctx),
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description: "A URL to redirect the request to.",
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("expression")),
														stringvalidator.LengthAtLeast(1),
													},
												},
												"expression": schema.StringAttribute{
													Description: "An expression that evaluates to a URL to redirect the request to.",
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
									},
								},
								"headers": schema.MapNestedAttribute{
									Description: "A map of headers to rewrite.",
									Optional:    true,
									Validators: []validator.Map{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"rewrite",
										),
										mapvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectMapType[RulesetRulesActionParametersHeadersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"operation": schema.StringAttribute{
												Description: "The operation to perform on the header.\nAvailable values: \"add\", \"set\", \"remove\".",
												Required:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"add",
														"set",
														"remove",
													),
												},
											},
											"value": schema.StringAttribute{
												Description: "A static value for the header.",
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("expression")),
													stringvalidator.LengthAtLeast(1),
												},
											},
											"expression": schema.StringAttribute{
												Description: "An expression that evaluates to a value for the header.",
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("value")),
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
									},
								},
								"uri": schema.SingleNestedAttribute{
									Description: "A URI rewrite.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"rewrite",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersURIModel](ctx),
									Attributes: map[string]schema.Attribute{
										"path": schema.SingleNestedAttribute{
											Description: "A URI path rewrite.",
											Optional:    true,
											Validators: []validator.Object{
												objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("query")),
											},
											CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersURIPathModel](ctx),
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description: "A value to rewrite the URI path to.",
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("expression")),
														stringvalidator.LengthAtLeast(1),
													},
												},
												"expression": schema.StringAttribute{
													Description: "An expression that evaluates to a value to rewrite the URI path to.",
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
										"query": schema.SingleNestedAttribute{
											Description: "A URI query rewrite.",
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersURIQueryModel](ctx),
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description: "A value to rewrite the URI query to.",
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("expression")),
													},
												},
												"expression": schema.StringAttribute{
													Description: "An expression that evaluates to a value to rewrite the URI query to.",
													Optional:    true,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
									},
								},
								"host_header": schema.StringAttribute{
									Description: "A value to rewrite the HTTP host header to.",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"route",
										),
										stringvalidator.LengthAtLeast(1),
									},
								},
								"origin": schema.SingleNestedAttribute{
									Description: "An origin to route to.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"route",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersOriginModel](ctx),
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description: "A resolved host to route to.",
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("port")),
												stringvalidator.LengthAtLeast(1),
											},
										},
										"port": schema.Int64Attribute{
											Description: "A destination port to route to.",
											Optional:    true,
											Validators: []validator.Int64{
												int64validator.Between(1, 65535),
											},
										},
									},
								},
								"sni": schema.SingleNestedAttribute{
									Description: "A Server Name Indication (SNI) override.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"route",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersSNIModel](ctx),
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description: "A value to override the SNI to.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
								},
								"increment": schema.Int64Attribute{
									Description: "A delta to change the score by, which can be either positive or negative.",
									Optional:    true,
									Validators: []validator.Int64{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"score",
										),
									},
								},
								"asset_name": schema.StringAttribute{
									Description: "The name of a custom asset to serve as the response.",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"serve_error",
										),
										stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("content")),
										stringvalidator.LengthAtLeast(1),
									},
								},
								"content": schema.StringAttribute{
									Description: "The response content.",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"serve_error",
										),
										stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("asset_name")),
										stringvalidator.LengthAtLeast(1),
									},
								},
								"content_type": schema.StringAttribute{
									Description: "The content type header to set with the error response.\nAvailable values: \"application/json\", \"text/html\", \"text/plain\", \"text/xml\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"serve_error",
										),
										stringvalidator.OneOfCaseInsensitive(
											"application/json",
											"text/html",
											"text/plain",
											"text/xml",
										),
									},
								},
								"status_code": schema.Int64Attribute{
									Description: "The status code to use for the error.",
									Optional:    true,
									Validators: []validator.Int64{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"serve_error",
										),
										int64validator.Between(400, 999),
									},
								},
								"automatic_https_rewrites": schema.BoolAttribute{
									Description: "Whether to enable Automatic HTTPS Rewrites.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"autominify": schema.SingleNestedAttribute{
									Description: "Which file extensions to minify automatically.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersAutominifyModel](ctx),
									Attributes: map[string]schema.Attribute{
										"css": schema.BoolAttribute{
											Description: "Whether to minify CSS files.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"html": schema.BoolAttribute{
											Description: "Whether to minify HTML files.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"js": schema.BoolAttribute{
											Description: "Whether to minify JavaScript files.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
									},
								},
								"bic": schema.BoolAttribute{
									Description: "Whether to enable Browser Integrity Check (BIC).",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"disable_apps": schema.BoolAttribute{
									Description: "Whether to disable Cloudflare Apps.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										boolvalidator.Equals(true),
									},
								},
								"disable_rum": schema.BoolAttribute{
									Description: "Whether to disable Real User Monitoring (RUM).",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										boolvalidator.Equals(true),
									},
								},
								"disable_zaraz": schema.BoolAttribute{
									Description: "Whether to disable Zaraz.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										boolvalidator.Equals(true),
									},
								},
								"email_obfuscation": schema.BoolAttribute{
									Description: "Whether to enable Email Obfuscation.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"fonts": schema.BoolAttribute{
									Description: "Whether to enable Cloudflare Fonts.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"hotlink_protection": schema.BoolAttribute{
									Description: "Whether to enable Hotlink Protection.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"mirage": schema.BoolAttribute{
									Description: "Whether to enable Mirage.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"opportunistic_encryption": schema.BoolAttribute{
									Description: "Whether to enable Opportunistic Encryption.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"polish": schema.StringAttribute{
									Description: "The Polish level to configure.\nAvailable values: \"off\", \"lossless\", \"lossy\", \"webp\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										stringvalidator.OneOfCaseInsensitive(
											"off",
											"lossless",
											"lossy",
											"webp",
										),
									},
								},
								"request_body_buffering": schema.StringAttribute{
									Description: "The request body buffering mode to configure.\nAvailable values: \"none\", \"standard\", \"full\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										stringvalidator.OneOfCaseInsensitive(
											"none",
											"standard",
											"full",
										),
									},
								},
								"response_body_buffering": schema.StringAttribute{
									Description: "The response body buffering mode to configure.\nAvailable values: \"none\", \"standard\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										stringvalidator.OneOfCaseInsensitive(
											"none",
											"standard",
										),
									},
								},
								"rocket_loader": schema.BoolAttribute{
									Description: "Whether to enable Rocket Loader.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"security_level": schema.StringAttribute{
									Description: "The Security Level to configure.\nAvailable values: \"off\", \"essentially_off\", \"low\", \"medium\", \"high\", \"under_attack\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										stringvalidator.OneOfCaseInsensitive(
											"off",
											"essentially_off",
											"low",
											"medium",
											"high",
											"under_attack",
										),
									},
								},
								"server_side_excludes": schema.BoolAttribute{
									Description: "Whether to enable Server-Side Excludes.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								"ssl": schema.StringAttribute{
									Description: "The SSL level to configure.\nAvailable values: \"off\", \"flexible\", \"full\", \"strict\", \"origin_pull\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
										stringvalidator.OneOfCaseInsensitive(
											"off",
											"flexible",
											"full",
											"strict",
											"origin_pull",
										),
									},
								},
								"sxg": schema.BoolAttribute{
									Description: "Whether to enable Signed Exchanges (SXG).",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_config",
										),
									},
								},
								// "phase": schema.StringAttribute{
								// 	Description: "A phase to skip the execution of. This option is only compatible with the products option.\nAvailable values: \"current\".",
								// 	Optional:    true,
								// 	Validators: []validator.String{
								// 		customvalidator.RequiresOtherStringAttributeToBe(
								// 			path.MatchRelative().AtParent().AtParent().AtName("action"),
								// 			"skip",
								// 		),
								// 		stringvalidator.OneOfCaseInsensitive("current"),
								// 	},
								// },
								"phases": schema.ListAttribute{
									Description: "A list of phases to skip the execution of. This option is incompatible with the rulesets option.\nAvailable values: \"ddos_l4\", \"ddos_l7\", \"http_config_settings\", \"http_custom_errors\", \"http_log_custom_fields\", \"http_ratelimit\", \"http_request_cache_settings\", \"http_request_dynamic_redirect\", \"http_request_firewall_custom\", \"http_request_firewall_managed\", \"http_request_late_transform\", \"http_request_origin\", \"http_request_redirect\", \"http_request_sanitize\", \"http_request_sbfm\", \"http_request_transform\", \"http_response_compression\", \"http_response_firewall_managed\", \"http_response_headers_transform\", \"magic_transit\", \"magic_transit_ids_managed\", \"magic_transit_managed\", \"magic_transit_ratelimit\".",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"skip",
										),
										listvalidator.SizeAtLeast(1),
										listvalidator.ValueStringsAre(
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
										),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"products": schema.ListAttribute{
									Description: "A list of legacy security products to skip the execution of.\nAvailable values: \"bic\", \"hot\", \"rateLimit\", \"securityLevel\", \"uaBlock\", \"waf\", \"zoneLockdown\".",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"skip",
										),
										listvalidator.SizeAtLeast(1),
										listvalidator.ValueStringsAre(
											stringvalidator.OneOfCaseInsensitive(
												"bic",
												"hot",
												"rateLimit",
												"securityLevel",
												"uaBlock",
												"waf",
												"zoneLockdown",
											),
										),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"rules": schema.MapAttribute{
									Description: "A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of. This option is incompatible with the ruleset option.",
									Optional:    true,
									Validators: []validator.Map{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"skip",
										),
										mapvalidator.SizeAtLeast(1),
										mapvalidator.ValueListsAre(
											listvalidator.SizeAtLeast(1),
											listvalidator.ValueStringsAre(
												stringvalidator.RegexMatches(
													regexp.MustCompile("^[0-9a-f]{32}$"),
													"value must be a 32-character hexadecimal string",
												),
											),
										),
									},
									CustomType: customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"ruleset": schema.StringAttribute{
									Description: "A ruleset to skip the execution of. This option is incompatible with the rulesets option.\nAvailable values: \"current\".",
									Optional:    true,
									Validators: []validator.String{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"skip",
										),
										stringvalidator.OneOfCaseInsensitive("current"),
									},
								},
								"rulesets": schema.ListAttribute{
									Description: "A list of ruleset IDs to skip the execution of. This option is incompatible with the ruleset and phases options.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"skip",
										),
										listvalidator.SizeAtLeast(1),
										listvalidator.ValueStringsAre(
											stringvalidator.RegexMatches(
												regexp.MustCompile("^[0-9a-f]{32}$"),
												"value must be a 32-character hexadecimal string",
											),
										),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"additional_cacheable_ports": schema.ListAttribute{
									Description: "A list of additional ports that caching should be enabled on.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
										listvalidator.SizeAtLeast(1),
										listvalidator.ValueInt64sAre(int64validator.Between(1, 65535)),
									},
									CustomType:  customfield.NewListType[types.Int64](ctx),
									ElementType: types.Int64Type,
								},
								"browser_ttl": schema.SingleNestedAttribute{
									Description: "How long client browsers should cache the response. Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersBrowserTTLModel](ctx),
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description: "The browser TTL mode.\nAvailable values: \"respect_origin\", \"bypass_by_default\", \"override_origin\", \"bypass\".",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"respect_origin",
													"bypass_by_default",
													"override_origin",
													"bypass",
												),
											},
										},
										"default": schema.Int64Attribute{
											Description: "The browser TTL (in seconds) if you choose the \"override_origin\" mode.",
											Optional:    true,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
									},
								},
								"cache": schema.BoolAttribute{
									Description: "Whether the request's response from the origin is eligible for caching. Caching itself will still depend on the cache control header and your other caching configurations.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
								},
								"cache_key": schema.SingleNestedAttribute{
									Description: "Which components of the request are included in or excluded from the cache key Cloudflare uses to store the response in cache.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyModel](ctx),
									Attributes: map[string]schema.Attribute{
										"cache_by_device_type": schema.BoolAttribute{
											Description: "Whether to separate cached content based on the visitor's device type.",
											Optional:    true,
										},
										"cache_deception_armor": schema.BoolAttribute{
											Description: "Whether to protect from web cache deception attacks, while allowing static assets to be cached.",
											Optional:    true,
										},
										"custom_key": schema.SingleNestedAttribute{
											Description: "Which components of the request are included or excluded from the cache key.",
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyModel](ctx),
											Attributes: map[string]schema.Attribute{
												"cookie": schema.SingleNestedAttribute{
													Description: "Which cookies to include in the cache key.",
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyCookieModel](ctx),
													Attributes: map[string]schema.Attribute{
														"check_presence": schema.ListAttribute{
															Description: "A list of cookies to check for the presence of. The presence of these cookies is included in the cache key.",
															Optional:    true,
															Validators: []validator.List{
																listvalidator.SizeAtLeast(1),
															},
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
														"include": schema.ListAttribute{
															Description: "A list of cookies to include in the cache key.",
															Optional:    true,
															Validators: []validator.List{
																listvalidator.SizeAtLeast(1),
															},
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
													},
												},
												"header": schema.SingleNestedAttribute{
													Description: "Which headers to include in the cache key.",
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyHeaderModel](ctx),
													Attributes: map[string]schema.Attribute{
														"check_presence": schema.ListAttribute{
															Description: "A list of headers to check for the presence of. The presence of these headers is included in the cache key.",
															Optional:    true,
															Validators: []validator.List{
																listvalidator.SizeAtLeast(1),
															},
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
														"contains": schema.MapAttribute{
															Description: "A mapping of header names to a list of values. If a header is present in the request and contains any of the values provided, its value is included in the cache key.",
															Optional:    true,
															Validators: []validator.Map{
																mapvalidator.SizeAtLeast(1),
																mapvalidator.ValueListsAre(listvalidator.SizeAtLeast(1)),
															},
															CustomType: customfield.NewMapType[customfield.List[types.String]](ctx),
															ElementType: types.ListType{
																ElemType: types.StringType,
															},
														},
														"exclude_origin": schema.BoolAttribute{
															Description: "Whether to exclude the origin header in the cache key.",
															Optional:    true,
														},
														"include": schema.ListAttribute{
															Description: "A list of headers to include in the cache key.",
															Optional:    true,
															Validators: []validator.List{
																listvalidator.SizeAtLeast(1),
															},
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
													},
												},
												"host": schema.SingleNestedAttribute{
													Description: "How to use the host in the cache key.",
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyHostModel](ctx),
													Attributes: map[string]schema.Attribute{
														"resolved": schema.BoolAttribute{
															Description: "Whether to use the resolved host in the cache key.",
															Optional:    true,
														},
													},
												},
												"query_string": schema.SingleNestedAttribute{
													Description: "Which query string parameters to include in or exclude from the cache key.",
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringModel](ctx),
													Attributes: map[string]schema.Attribute{
														"include": schema.SingleNestedAttribute{
															Description: "Which query string parameters to include in the cache key.",
															Optional:    true,
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("exclude")),
															},
															CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeModel](ctx),
															Attributes: map[string]schema.Attribute{
																"list": schema.ListAttribute{
																	Description: "A list of query string parameters to include in the cache key.",
																	Optional:    true,
																	Validators: []validator.List{
																		listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("all")),
																		listvalidator.SizeAtLeast(1),
																	},
																	CustomType:  customfield.NewListType[types.String](ctx),
																	ElementType: types.StringType,
																},
																"all": schema.BoolAttribute{
																	Description: "Whether to include all query string parameters in the cache key.",
																	Optional:    true,
																	Validators: []validator.Bool{
																		boolvalidator.Equals(true),
																	},
																},
															},
														},
														"exclude": schema.SingleNestedAttribute{
															Description: "Which query string parameters to exclude from the cache key.",
															Optional:    true,
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("include")),
															},
															CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeModel](ctx),
															Attributes: map[string]schema.Attribute{
																"list": schema.ListAttribute{
																	Description: "A list of query string parameters to exclude from the cache key.",
																	Optional:    true,
																	Validators: []validator.List{
																		listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("all")),
																		listvalidator.SizeAtLeast(1),
																	},
																	CustomType:  customfield.NewListType[types.String](ctx),
																	ElementType: types.StringType,
																},
																"all": schema.BoolAttribute{
																	Description: "Whether to exclude all query string parameters from the cache key.",
																	Optional:    true,
																	Validators: []validator.Bool{
																		boolvalidator.Equals(true),
																	},
																},
															},
														},
													},
												},
												"user": schema.SingleNestedAttribute{
													Description: "How to use characteristics of the request user agent in the cache key.",
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyUserModel](ctx),
													Attributes: map[string]schema.Attribute{
														"device_type": schema.BoolAttribute{
															Description: "Whether to use the user agent's device type in the cache key.",
															Optional:    true,
														},
														"geo": schema.BoolAttribute{
															Description: "Whether to use the user agents's country in the cache key.",
															Optional:    true,
														},
														"lang": schema.BoolAttribute{
															Description: "Whether to use the user agent's language in the cache key.",
															Optional:    true,
														},
													},
												},
											},
										},
										"ignore_query_strings_order": schema.BoolAttribute{
											Description: "Whether to treat requests with the same query parameters the same, regardless of the order those query parameters are in.",
											Optional:    true,
										},
									},
								},
								"cache_reserve": schema.SingleNestedAttribute{
									Description: "Settings to determine whether the request's response from origin is eligible for Cache Reserve (requires a Cache Reserve add-on plan).",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersCacheReserveModel](ctx),
									Attributes: map[string]schema.Attribute{
										"eligible": schema.BoolAttribute{
											Description: "Whether Cache Reserve is enabled. If this is true and a request meets eligibility criteria, Cloudflare will write the resource to Cache Reserve.",
											Required:    true,
										},
										"minimum_file_size": schema.Int64Attribute{
											Description: "The minimum file size eligible for storage in Cache Reserve.",
											Optional:    true,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
									},
								},
								"edge_ttl": schema.SingleNestedAttribute{
									Description: "How long the Cloudflare edge network should cache the response.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersEdgeTTLModel](ctx),
									Attributes: map[string]schema.Attribute{
										"default": schema.Int64Attribute{
											Description: "The edge TTL (in seconds) if you choose the \"override_origin\" mode.",
											Optional:    true,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
										"mode": schema.StringAttribute{
											Description: "The edge TTL mode.\nAvailable values: \"respect_origin\", \"bypass_by_default\", \"override_origin\".",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"respect_origin",
													"bypass_by_default",
													"override_origin",
												),
											},
										},
										"status_code_ttl": schema.ListNestedAttribute{
											Description: "A list of TTLs to apply to specific status codes or status code ranges.",
											Optional:    true,
											Validators: []validator.List{
												listvalidator.SizeAtLeast(1),
											},
											CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersEdgeTTLStatusCodeTTLModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"status_code_range": schema.SingleNestedAttribute{
														Description: "A range of status codes to apply the TTL to.",
														Optional:    true,
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("status_code")),
														},
														CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeModel](ctx),
														Attributes: map[string]schema.Attribute{
															"from": schema.Int64Attribute{
																Description: "The lower bound of the range.",
																Optional:    true,
																Validators: []validator.Int64{
																	int64validator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("to")),
																	int64validator.Between(100, 999),
																},
															},
															"to": schema.Int64Attribute{
																Description: "The upper bound of the range.",
																Optional:    true,
																Validators: []validator.Int64{
																	int64validator.Between(100, 999),
																},
															},
														},
													},
													"status_code": schema.Int64Attribute{
														Description: "A single status code to apply the TTL to.",
														Optional:    true,
														Validators: []validator.Int64{
															int64validator.Between(100, 999),
														},
													},
													"value": schema.Int64Attribute{
														Description: "The time to cache the response for (in seconds). A value of 0 is equivalent to setting the cache control header with the value \"no-cache\". A value of -1 is equivalent to setting the cache control header with the value of \"no-store\".",
														Required:    true,
													},
												},
											},
										},
									},
								},
								"origin_cache_control": schema.BoolAttribute{
									Description: "Whether Cloudflare will aim to strictly adhere to RFC 7234.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
								},
								"origin_error_page_passthru": schema.BoolAttribute{
									Description: "Whether to generate Cloudflare error pages for issues from the origin server.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
								},
								"read_timeout": schema.Int64Attribute{
									Description: "A timeout value between two successive read operations to use for your origin server. Historically, the timeout value between two read options from Cloudflare to an origin server is 100 seconds. If you are attempting to reduce HTTP 524 errors because of timeouts from an origin server, try increasing this timeout value.",
									Optional:    true,
									Validators: []validator.Int64{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
										int64validator.Between(100, 6000),
									},
								},
								"respect_strong_etags": schema.BoolAttribute{
									Description: "Whether Cloudflare should respect strong ETag (entity tag) headers. If false, Cloudflare converts strong ETag headers to weak ETag headers.",
									Optional:    true,
									Validators: []validator.Bool{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
								},
								"serve_stale": schema.SingleNestedAttribute{
									Description: "When to serve stale content from cache.",
									Optional:    true,
									Validators: []validator.Object{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"set_cache_settings",
										),
									},
									CustomType: customfield.NewNestedObjectType[RulesetRulesActionParametersServeStaleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"disable_stale_while_updating": schema.BoolAttribute{
											Description: "Whether Cloudflare should disable serving stale content while getting the latest content from the origin.",
											Optional:    true,
										},
									},
								},
								"cookie_fields": schema.ListNestedAttribute{
									Description: "The cookie fields to log.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"log_custom_field",
										),
										listvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersCookieFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the cookie.",
												Required:    true,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
									},
								},
								"raw_response_fields": schema.ListNestedAttribute{
									Description: "The raw response fields to log.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"log_custom_field",
										),
										listvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersRawResponseFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the response header.",
												Required:    true,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
											"preserve_duplicates": schema.BoolAttribute{
												Description: "Whether to log duplicate values of the same header.",
												Computed:    true,
												Optional:    true,
												Default:     booldefault.StaticBool(false),
											},
										},
									},
								},
								"request_fields": schema.ListNestedAttribute{
									Description: "The raw request fields to log.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"log_custom_field",
										),
										listvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersRequestFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the header.",
												Required:    true,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
									},
								},
								"response_fields": schema.ListNestedAttribute{
									Description: "The transformed response fields to log.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"log_custom_field",
										),
										listvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersResponseFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the response header.",
												Required:    true,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
											"preserve_duplicates": schema.BoolAttribute{
												Description: "Whether to log duplicate values of the same header.",
												Computed:    true,
												Optional:    true,
												Default:     booldefault.StaticBool(false),
											},
										},
									},
								},
								"transformed_request_fields": schema.ListNestedAttribute{
									Description: "The transformed request fields to log.",
									Optional:    true,
									Validators: []validator.List{
										customvalidator.RequiresOtherStringAttributeToBe(
											path.MatchRelative().AtParent().AtParent().AtName("action"),
											"log_custom_field",
										),
										listvalidator.SizeAtLeast(1),
									},
									CustomType: customfield.NewNestedObjectListType[RulesetRulesActionParametersTransformedRequestFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the header.",
												Required:    true,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
									},
								},
							},
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the rule.",
							Computed:    true,
							Optional:    true,
							Default:     stringdefault.StaticString(""),
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the rule should be executed.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(true),
						},
						"exposed_credential_check": schema.SingleNestedAttribute{
							Description: "Configuration for exposed credential checking.",
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[RulesetRulesExposedCredentialCheckModel](ctx),
							Attributes: map[string]schema.Attribute{
								"password_expression": schema.StringAttribute{
									Description: "An expression that selects the password used in the credentials check.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
								"username_expression": schema.StringAttribute{
									Description: "An expression that selects the user ID used in the credentials check.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						"expression": schema.StringAttribute{
							Description: "The expression defining which traffic will match the rule.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"logging": schema.SingleNestedAttribute{
							Description: "An object configuring the rule's logging behavior.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.Object{
								objectvalidator.AlsoRequires(path.MatchRelative().AtName("enabled")),
							},
							CustomType: customfield.NewNestedObjectType[RulesetRulesLoggingModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether to generate a log when the rule matches.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"ratelimit": schema.SingleNestedAttribute{
							Description: "An object configuring the rule's rate limit behavior.",
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[RulesetRulesRatelimitModel](ctx),
							Attributes: map[string]schema.Attribute{
								"characteristics": schema.ListAttribute{
									Description: "Characteristics of the request on which the rate limit counter will be incremented.",
									Required:    true,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"period": schema.Int64Attribute{
									Description: "Period in seconds over which the counter is being incremented.",
									Required:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
								"counting_expression": schema.StringAttribute{
									Description: "An expression that defines when the rate limit counter should be incremented. It defaults to the same as the rule's expression.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
								"mitigation_timeout": schema.Int64Attribute{
									Description: "Period of time in seconds after which the action will be disabled following its first execution.",
									Computed:    true,
									Optional:    true,
								},
								"requests_per_period": schema.Int64Attribute{
									Description: "The threshold of requests per period after which the action will be executed for the first time.",
									Optional:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
									},
								},
								"requests_to_origin": schema.BoolAttribute{
									Description: "Whether counting is only performed when an origin is reached.",
									Computed:    true,
									Optional:    true,
									Default:     booldefault.StaticBool(false),
								},
								"score_per_period": schema.Int64Attribute{
									Description: "The score threshold per period for which the action will be executed the first time.",
									Optional:    true,
								},
								"score_response_header_name": schema.StringAttribute{
									Description: "A response header name provided by the origin, which contains the score to increment rate limit counter with.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						"ref": schema.StringAttribute{
							Description: "The reference of the rule (the rule's ID by default).",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
					},
				},
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
	}
}

func (r *RulesetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *RulesetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
