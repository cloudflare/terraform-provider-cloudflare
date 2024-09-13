// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"math"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
				Description:   "The unique ID of the ruleset.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"kind": schema.StringAttribute{
				Description: "The kind of the ruleset.",
				Required:    true,
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
				Required:    true,
			},
			"phase": schema.StringAttribute{
				Description: "The phase of the ruleset.",
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
						"http_request_select_configuration",
						"http_request_transform",
						"http_response_compression",
						"http_response_firewall_managed",
						"http_response_headers_transform",
						"magic_transit",
						"magic_transit_ids_managed",
						"magic_transit_managed",
					),
				},
			},
			"rules": schema.ListNestedAttribute{
				Description: "The list of rules in the ruleset.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"last_updated": schema.StringAttribute{
							Description: "The timestamp of when the rule was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"version": schema.StringAttribute{
							Description: "The version of the rule.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "The unique ID of the rule.",
							Computed:    true,
							Optional:    true,
						},
						"action": schema.StringAttribute{
							Description: "The action to perform when the rule matches.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"block",
									"challenge",
									"compress_response",
									"execute",
									"js_challenge",
									"log",
									"managed_challenge",
									"redirect",
									"rewrite",
									"route",
									"score",
									"serve_error",
									"set_config",
									"skip",
									"set_cache_settings",
									"log_custom_field",
									"ddos_dynamic",
									"force_connection_close",
								),
							},
						},
						"action_parameters": schema.SingleNestedAttribute{
							Description: "The parameters configuring the rule's action.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersModel](ctx),
							Attributes: map[string]schema.Attribute{
								"response": schema.SingleNestedAttribute{
									Description: "The response to show when the block is applied.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersResponseModel](ctx),
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description: "The content to return.",
											Required:    true,
										},
										"content_type": schema.StringAttribute{
											Description: "The type of the content to return.",
											Required:    true,
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
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectListType[RulesetRulesActionParametersAlgorithmsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "Name of compression algorithm to enable.",
												Computed:    true,
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"none",
														"auto",
														"default",
														"gzip",
														"brotli",
													),
												},
											},
										},
									},
								},
								"id": schema.StringAttribute{
									Description: "The ID of the ruleset to execute.",
									Computed:    true,
									Optional:    true,
								},
								"matched_data": schema.SingleNestedAttribute{
									Description: "The configuration to use for matched data logging.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersMatchedDataModel](ctx),
									Attributes: map[string]schema.Attribute{
										"public_key": schema.StringAttribute{
											Description: "The public key to encrypt matched data logs with.",
											Required:    true,
										},
									},
								},
								"overrides": schema.SingleNestedAttribute{
									Description: "A set of overrides to apply to the target ruleset.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersOverridesModel](ctx),
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description: "An action to override all rules with. This option has lower precedence than rule and category overrides.",
											Computed:    true,
											Optional:    true,
										},
										"categories": schema.ListNestedAttribute{
											Description: "A list of category-level overrides. This option has the second-highest precedence after rule-level overrides.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectListType[RulesetRulesActionParametersOverridesCategoriesModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"category": schema.StringAttribute{
														Description: "The name of the category to override.",
														Required:    true,
													},
													"action": schema.StringAttribute{
														Description: "The action to override rules in the category with.",
														Computed:    true,
														Optional:    true,
													},
													"enabled": schema.BoolAttribute{
														Description: "Whether to enable execution of rules in the category.",
														Computed:    true,
														Optional:    true,
													},
													"sensitivity_level": schema.StringAttribute{
														Description: "The sensitivity level to use for rules in the category.",
														Computed:    true,
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
											Computed:    true,
											Optional:    true,
										},
										"rules": schema.ListNestedAttribute{
											Description: "A list of rule-level overrides. This option has the highest precedence.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectListType[RulesetRulesActionParametersOverridesRulesModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "The ID of the rule to override.",
														Required:    true,
													},
													"action": schema.StringAttribute{
														Description: "The action to override the rule with.",
														Computed:    true,
														Optional:    true,
													},
													"enabled": schema.BoolAttribute{
														Description: "Whether to enable execution of the rule.",
														Computed:    true,
														Optional:    true,
													},
													"score_threshold": schema.Int64Attribute{
														Description: "The score threshold to use for the rule.",
														Computed:    true,
														Optional:    true,
													},
													"sensitivity_level": schema.StringAttribute{
														Description: "The sensitivity level to use for the rule.",
														Computed:    true,
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
											Description: "A sensitivity level to set for all rules. This option has lower precedence than rule and category overrides and is only applicable for DDoS phases.",
											Computed:    true,
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
									Description: "Serve a redirect based on a bulk list lookup.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersFromListModel](ctx),
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description: "Expression that evaluates to the list lookup key.",
											Computed:    true,
											Optional:    true,
										},
										"name": schema.StringAttribute{
											Description: "The name of the list to match against.",
											Computed:    true,
											Optional:    true,
										},
									},
								},
								"from_value": schema.SingleNestedAttribute{
									Description: "Serve a redirect based on the request properties.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersFromValueModel](ctx),
									Attributes: map[string]schema.Attribute{
										"preserve_query_string": schema.BoolAttribute{
											Description: "Keep the query string of the original request.",
											Computed:    true,
											Optional:    true,
										},
										"status_code": schema.Float64Attribute{
											Description: "The status code to be used for the redirect.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.Float64{
												float64validator.OneOf(
													301,
													302,
													303,
													307,
													308,
												),
											},
										},
										"target_url": schema.SingleNestedAttribute{
											Description: "The URL to redirect the request to.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersFromValueTargetURLModel](ctx),
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description: "The URL to redirect the request to.",
													Computed:    true,
													Optional:    true,
												},
												"expression": schema.StringAttribute{
													Description: "An expression to evaluate to get the URL to redirect the request to.",
													Computed:    true,
													Optional:    true,
												},
											},
										},
									},
								},
								"headers": schema.MapNestedAttribute{
									Description: "Map of request headers to modify.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectMapType[RulesetRulesActionParametersHeadersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"operation": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("remove", "set"),
												},
											},
											"value": schema.StringAttribute{
												Description: "Static value for the header.",
												Computed:    true,
												Optional:    true,
											},
											"expression": schema.StringAttribute{
												Description: "Expression for the header value.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
								"uri": schema.SingleNestedAttribute{
									Description: "URI to rewrite the request to.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersURIModel](ctx),
									Attributes: map[string]schema.Attribute{
										"path": schema.SingleNestedAttribute{
											Description: "Path portion rewrite.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersURIPathModel](ctx),
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description: "Predefined replacement value.",
													Computed:    true,
													Optional:    true,
												},
												"expression": schema.StringAttribute{
													Description: "Expression to evaluate for the replacement value.",
													Computed:    true,
													Optional:    true,
												},
											},
										},
										"query": schema.SingleNestedAttribute{
											Description: "Query portion rewrite.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersURIQueryModel](ctx),
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description: "Predefined replacement value.",
													Computed:    true,
													Optional:    true,
												},
												"expression": schema.StringAttribute{
													Description: "Expression to evaluate for the replacement value.",
													Computed:    true,
													Optional:    true,
												},
											},
										},
									},
								},
								"host_header": schema.StringAttribute{
									Description: "Rewrite the HTTP Host header.",
									Computed:    true,
									Optional:    true,
								},
								"origin": schema.SingleNestedAttribute{
									Description: "Override the IP/TCP destination.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersOriginModel](ctx),
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description: "Override the resolved hostname.",
											Computed:    true,
											Optional:    true,
										},
										"port": schema.Float64Attribute{
											Description: "Override the destination port.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.Float64{
												float64validator.Between(1, 65535),
											},
										},
									},
								},
								"sni": schema.SingleNestedAttribute{
									Description: "Override the Server Name Indication (SNI).",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersSNIModel](ctx),
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description: "The SNI override.",
											Required:    true,
										},
									},
								},
								"increment": schema.Int64Attribute{
									Description: "Increment contains the delta to change the score and can be either positive or negative.",
									Computed:    true,
									Optional:    true,
								},
								"content": schema.StringAttribute{
									Description: "Error response content.",
									Computed:    true,
									Optional:    true,
								},
								"content_type": schema.StringAttribute{
									Description: "Content-type header to set with the response.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"application/json",
											"text/xml",
											"text/plain",
											"text/html",
										),
									},
								},
								"status_code": schema.Float64Attribute{
									Description: "The status code to use for the error.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.Float64{
										float64validator.Between(400, 999),
									},
								},
								"automatic_https_rewrites": schema.BoolAttribute{
									Description: "Turn on or off Automatic HTTPS Rewrites.",
									Computed:    true,
									Optional:    true,
								},
								"autominify": schema.SingleNestedAttribute{
									Description: "Select which file extensions to minify automatically.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersAutominifyModel](ctx),
									Attributes: map[string]schema.Attribute{
										"css": schema.BoolAttribute{
											Description: "Minify CSS files.",
											Computed:    true,
											Optional:    true,
										},
										"html": schema.BoolAttribute{
											Description: "Minify HTML files.",
											Computed:    true,
											Optional:    true,
										},
										"js": schema.BoolAttribute{
											Description: "Minify JS files.",
											Computed:    true,
											Optional:    true,
										},
									},
								},
								"bic": schema.BoolAttribute{
									Description: "Turn on or off Browser Integrity Check.",
									Computed:    true,
									Optional:    true,
								},
								"disable_apps": schema.BoolAttribute{
									Description: "Turn off all active Cloudflare Apps.",
									Computed:    true,
									Optional:    true,
								},
								"disable_rum": schema.BoolAttribute{
									Description: "Turn off Real User Monitoring (RUM).",
									Computed:    true,
									Optional:    true,
								},
								"disable_zaraz": schema.BoolAttribute{
									Description: "Turn off Zaraz.",
									Computed:    true,
									Optional:    true,
								},
								"email_obfuscation": schema.BoolAttribute{
									Description: "Turn on or off Email Obfuscation.",
									Computed:    true,
									Optional:    true,
								},
								"fonts": schema.BoolAttribute{
									Description: "Turn on or off Cloudflare Fonts.",
									Computed:    true,
									Optional:    true,
								},
								"hotlink_protection": schema.BoolAttribute{
									Description: "Turn on or off the Hotlink Protection.",
									Computed:    true,
									Optional:    true,
								},
								"mirage": schema.BoolAttribute{
									Description: "Turn on or off Mirage.",
									Computed:    true,
									Optional:    true,
								},
								"opportunistic_encryption": schema.BoolAttribute{
									Description: "Turn on or off Opportunistic Encryption.",
									Computed:    true,
									Optional:    true,
								},
								"polish": schema.StringAttribute{
									Description: "Configure the Polish level.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"off",
											"lossless",
											"lossy",
										),
									},
								},
								"rocket_loader": schema.BoolAttribute{
									Description: "Turn on or off Rocket Loader",
									Computed:    true,
									Optional:    true,
								},
								"security_level": schema.StringAttribute{
									Description: "Configure the Security Level.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
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
									Description: "Turn on or off Server Side Excludes.",
									Computed:    true,
									Optional:    true,
								},
								"ssl": schema.StringAttribute{
									Description: "Configure the SSL level.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
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
									Description: "Turn on or off Signed Exchanges (SXG).",
									Computed:    true,
									Optional:    true,
								},
								"phases": schema.ListAttribute{
									Description: "A list of phases to skip the execution of. This option is incompatible with the ruleset and rulesets options.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.List{
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
												"http_request_select_configuration",
												"http_request_transform",
												"http_response_compression",
												"http_response_firewall_managed",
												"http_response_headers_transform",
												"magic_transit",
												"magic_transit_ids_managed",
												"magic_transit_managed",
											),
										),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"products": schema.ListAttribute{
									Description: "A list of legacy security products to skip the execution of.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.List{
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
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"ruleset": schema.StringAttribute{
									Description: "A ruleset to skip the execution of. This option is incompatible with the rulesets, rules and phases options.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("current"),
									},
								},
								"rulesets": schema.ListAttribute{
									Description: "A list of ruleset IDs to skip the execution of. This option is incompatible with the ruleset and phases options.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"additional_cacheable_ports": schema.ListAttribute{
									Description: "List of additional ports that caching can be enabled on.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewListType[types.Int64](ctx),
									ElementType: types.Int64Type,
								},
								"browser_ttl": schema.SingleNestedAttribute{
									Description: "Specify how long client browsers should cache the response. Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersBrowserTTLModel](ctx),
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description: "Determines which browser ttl mode to use.",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"respect_origin",
													"bypass_by_default",
													"override_origin",
												),
											},
										},
										"default": schema.Int64Attribute{
											Description: "The TTL (in seconds) if you choose override_origin mode.",
											Computed:    true,
											Optional:    true,
										},
									},
								},
								"cache": schema.BoolAttribute{
									Description: "Mark whether the request’s response from origin is eligible for caching. Caching itself will still depend on the cache-control header and your other caching configurations.",
									Computed:    true,
									Optional:    true,
								},
								"cache_key": schema.SingleNestedAttribute{
									Description: "Define which components of the request are included or excluded from the cache key Cloudflare uses to store the response in cache.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyModel](ctx),
									Attributes: map[string]schema.Attribute{
										"cache_by_device_type": schema.BoolAttribute{
											Description: "Separate cached content based on the visitor’s device type",
											Computed:    true,
											Optional:    true,
										},
										"cache_deception_armor": schema.BoolAttribute{
											Description: "Protect from web cache deception attacks while allowing static assets to be cached",
											Computed:    true,
											Optional:    true,
										},
										"custom_key": schema.SingleNestedAttribute{
											Description: "Customize which components of the request are included or excluded from the cache key.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyModel](ctx),
											Attributes: map[string]schema.Attribute{
												"cookie": schema.SingleNestedAttribute{
													Description: "The cookies to include in building the cache key.",
													Computed:    true,
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyCookieModel](ctx),
													Attributes: map[string]schema.Attribute{
														"check_presence": schema.ListAttribute{
															Description: "Checks for the presence of these cookie names. The presence of these cookies is used in building the cache key.",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
														"include": schema.ListAttribute{
															Description: "Include these cookies' names and their values.",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
													},
												},
												"header": schema.SingleNestedAttribute{
													Description: "The header names and values to include in building the cache key.",
													Computed:    true,
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyHeaderModel](ctx),
													Attributes: map[string]schema.Attribute{
														"check_presence": schema.ListAttribute{
															Description: "Checks for the presence of these header names. The presence of these headers is used in building the cache key.",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
														"contains": schema.MapAttribute{
															Description: "For each header name and list of values combination, check if the request header contains any of the values provided. The presence of the request header and whether any of the values provided are contained in the request header value is used in building the cache key.",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
															ElementType: types.ListType{
																ElemType: types.StringType,
															},
														},
														"exclude_origin": schema.BoolAttribute{
															Description: "Whether or not to include the origin header. A value of true will exclude the origin header in the cache key.",
															Computed:    true,
															Optional:    true,
														},
														"include": schema.ListAttribute{
															Description: "Include these headers' names and their values.",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewListType[types.String](ctx),
															ElementType: types.StringType,
														},
													},
												},
												"host": schema.SingleNestedAttribute{
													Description: "Whether to use the original host or the resolved host in the cache key.",
													Computed:    true,
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyHostModel](ctx),
													Attributes: map[string]schema.Attribute{
														"resolved": schema.BoolAttribute{
															Description: "Use the resolved host in the cache key. A value of true will use the resolved host, while a value or false will use the original host.",
															Computed:    true,
															Optional:    true,
														},
													},
												},
												"query_string": schema.SingleNestedAttribute{
													Description: "Use the presence or absence of parameters in the query string to build the cache key.",
													Computed:    true,
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringModel](ctx),
													Attributes: map[string]schema.Attribute{
														"exclude": schema.SingleNestedAttribute{
															Description: "build the cache key using all query string parameters EXCECPT these excluded parameters",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeModel](ctx),
															Attributes: map[string]schema.Attribute{
																"all": schema.BoolAttribute{
																	Description: "Exclude all query string parameters from use in building the cache key.",
																	Computed:    true,
																	Optional:    true,
																},
																"list": schema.ListAttribute{
																	Description: "A list of query string parameters NOT used to build the cache key. All parameters present in the request but missing in this list will be used to build the cache key.",
																	Computed:    true,
																	Optional:    true,
																	CustomType:  customfield.NewListType[types.String](ctx),
																	ElementType: types.StringType,
																},
															},
														},
														"include": schema.SingleNestedAttribute{
															Description: "build the cache key using a list of query string parameters that ARE in the request.",
															Computed:    true,
															Optional:    true,
															CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeModel](ctx),
															Attributes: map[string]schema.Attribute{
																"all": schema.BoolAttribute{
																	Description: "Use all query string parameters in the cache key.",
																	Computed:    true,
																	Optional:    true,
																},
																"list": schema.ListAttribute{
																	Description: "A list of query string parameters used to build the cache key.",
																	Computed:    true,
																	Optional:    true,
																	CustomType:  customfield.NewListType[types.String](ctx),
																	ElementType: types.StringType,
																},
															},
														},
													},
												},
												"user": schema.SingleNestedAttribute{
													Description: "Characteristics of the request user agent used in building the cache key.",
													Computed:    true,
													Optional:    true,
													CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheKeyCustomKeyUserModel](ctx),
													Attributes: map[string]schema.Attribute{
														"device_type": schema.BoolAttribute{
															Description: "Use the user agent's device type in the cache key.",
															Computed:    true,
															Optional:    true,
														},
														"geo": schema.BoolAttribute{
															Description: "Use the user agents's country in the cache key.",
															Computed:    true,
															Optional:    true,
														},
														"lang": schema.BoolAttribute{
															Description: "Use the user agent's language in the cache key.",
															Computed:    true,
															Optional:    true,
														},
													},
												},
											},
										},
										"ignore_query_strings_order": schema.BoolAttribute{
											Description: "Treat requests with the same query parameters the same, regardless of the order those query parameters are in. A value of true ignores the query strings' order.",
											Computed:    true,
											Optional:    true,
										},
									},
								},
								"cache_reserve": schema.SingleNestedAttribute{
									Description: "Mark whether the request's response from origin is eligible for  Cache Reserve (requires a Cache Reserve add-on plan).",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersCacheReserveModel](ctx),
									Attributes: map[string]schema.Attribute{
										"eligible": schema.BoolAttribute{
											Description: "Determines whether cache reserve is enabled. If this is true and a request meets eligibility criteria, Cloudflare will write the resource to cache reserve.",
											Required:    true,
										},
										"minimum_file_size": schema.Int64Attribute{
											Description: "The minimum file size eligible for store in cache reserve.",
											Required:    true,
										},
									},
								},
								"edge_ttl": schema.SingleNestedAttribute{
									Description: "TTL (Time to Live) specifies the maximum time to cache a resource in the Cloudflare edge network.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersEdgeTTLModel](ctx),
									Attributes: map[string]schema.Attribute{
										"default": schema.Int64Attribute{
											Description: "The TTL (in seconds) if you choose override_origin mode.",
											Required:    true,
											Validators: []validator.Int64{
												int64validator.Between(0, math.MaxInt64),
											},
										},
										"mode": schema.StringAttribute{
											Description: "edge ttl options",
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
											Description: "List of single status codes, or status code ranges to apply the selected mode",
											Required:    true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"value": schema.Int64Attribute{
														Description: "Time to cache a response (in seconds). A value of 0 is equivalent to setting the Cache-Control header with the value \"no-cache\". A value of -1 is equivalent to setting Cache-Control header with the value of \"no-store\".",
														Required:    true,
													},
													"status_code_range": schema.SingleNestedAttribute{
														Description: "The range of status codes used to apply the selected mode.",
														Computed:    true,
														Optional:    true,
														CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeModel](ctx),
														Attributes: map[string]schema.Attribute{
															"from": schema.Int64Attribute{
																Description: "response status code lower bound",
																Required:    true,
															},
															"to": schema.Int64Attribute{
																Description: "response status code upper bound",
																Required:    true,
															},
														},
													},
													"status_code_value": schema.Int64Attribute{
														Description: "Set the ttl for responses with this specific status code",
														Computed:    true,
														Optional:    true,
													},
												},
											},
										},
									},
								},
								"origin_cache_control": schema.BoolAttribute{
									Description: "When enabled, Cloudflare will aim to strictly adhere to RFC 7234.",
									Computed:    true,
									Optional:    true,
								},
								"origin_error_page_passthru": schema.BoolAttribute{
									Description: "Generate Cloudflare error pages from issues sent from the origin server. When on, error pages will trigger for issues from the origin",
									Computed:    true,
									Optional:    true,
								},
								"read_timeout": schema.Int64Attribute{
									Description: "Define a timeout value between two successive read operations to your origin server. Historically, the timeout value between two read options from Cloudflare to an origin server is 100 seconds. If you are attempting to reduce HTTP 524 errors because of timeouts from an origin server, try increasing this timeout value.",
									Computed:    true,
									Optional:    true,
								},
								"respect_strong_etags": schema.BoolAttribute{
									Description: "Specify whether or not Cloudflare should respect strong ETag (entity tag) headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.",
									Computed:    true,
									Optional:    true,
								},
								"serve_stale": schema.SingleNestedAttribute{
									Description: "Define if Cloudflare should serve stale content while getting the latest content from the origin. If on, Cloudflare will not serve stale content while getting the latest content from the origin.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[RulesetRulesActionParametersServeStaleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"disable_stale_while_updating": schema.BoolAttribute{
											Description: "Defines whether Cloudflare should serve stale content while updating. If true, Cloudflare will not serve stale content while getting the latest content from the origin.",
											Required:    true,
										},
									},
								},
								"cookie_fields": schema.ListNestedAttribute{
									Description: "The cookie fields to log.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectListType[RulesetRulesActionParametersCookieFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the field.",
												Required:    true,
											},
										},
									},
								},
								"request_fields": schema.ListNestedAttribute{
									Description: "The request fields to log.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectListType[RulesetRulesActionParametersRequestFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the field.",
												Required:    true,
											},
										},
									},
								},
								"response_fields": schema.ListNestedAttribute{
									Description: "The response fields to log.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectListType[RulesetRulesActionParametersResponseFieldsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the field.",
												Required:    true,
											},
										},
									},
								},
							},
						},
						"categories": schema.ListAttribute{
							Description: "The categories of the rule.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
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
						"expression": schema.StringAttribute{
							Description: "The expression defining which traffic will match the rule.",
							Computed:    true,
							Optional:    true,
						},
						"logging": schema.SingleNestedAttribute{
							Description: "An object configuring the rule's logging behavior.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[RulesetRulesLoggingModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether to generate a log when the rule matches.",
									Required:    true,
								},
							},
						},
						"ref": schema.StringAttribute{
							Description: "The reference of the rule (the rule ID by default).",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "An informative description of the ruleset.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"last_updated": schema.StringAttribute{
				Description: "The timestamp of when the ruleset was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.StringAttribute{
				Description: "The version of the ruleset.",
				Computed:    true,
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
