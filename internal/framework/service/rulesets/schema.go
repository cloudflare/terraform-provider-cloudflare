package rulesets

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/modifiers/defaults"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *RulesetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Cloudflare Ruleset Engine](https://developers.cloudflare.com/firewall/cf-rulesets)
			allows you to create and deploy rules and rulesets.

			The engine syntax, inspired by the Wireshark Display Filter language, is the
			same syntax used in custom Firewall Rules. Cloudflare uses the Ruleset Engine
			in different products, allowing you to configure several products using the same
			basic syntax.
		`),
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: consts.IDSchemaDescription,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot(consts.ZoneIDSchemaKey)),
					),
				},
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot(consts.AccountIDSchemaKey)),
					),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the ruleset.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Brief summary of the ruleset and its intended use.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Computed: true,
			},
			"kind": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(cfv1.RulesetKindValues()...),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: fmt.Sprintf("Type of Ruleset to create. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetKindValues())),
			},
			"phase": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(cfv1.RulesetPhaseValues()...),
					sbfmDeprecationWarningValidator{},
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetPhaseValues())),
			},
		},
		Blocks: map[string]schema.Block{
			"rules": schema.ListNestedBlock{
				MarkdownDescription: "List of rules to apply to the ruleset.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						consts.IDSchemaKey: schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Unique rule identifier.",
						},
						"version": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Version of the ruleset to deploy.",
						},
						"ref": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Rule reference.",
						},
						"enabled": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Whether the rule is active.",
							PlanModifiers: []planmodifier.Bool{
								defaults.DefaultBool(true),
							},
						},
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Brief summary of the ruleset rule and its intended use.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"expression": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Criteria for an HTTP request to trigger the ruleset rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
						},
						"action": schema.StringAttribute{
							MarkdownDescription: fmt.Sprintf("Action to perform in the ruleset rule. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetRuleActionValues())),
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(cfv1.RulesetRuleActionValues()...),
							},
							Optional: true,
						},
						"last_updated": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The most recent update to this rule.",
						},
					},
					Blocks: map[string]schema.Block{
						"action_parameters": schema.ListNestedBlock{
							MarkdownDescription: "List of parameters that configure the behavior of the ruleset rule action.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"additional_cacheable_ports": schema.SetAttribute{
										ElementType:         types.Int64Type,
										Optional:            true,
										MarkdownDescription: "Specifies uncommon ports to allow cacheable assets to be served from.",
									},
									"automatic_https_rewrites": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off Cloudflare Automatic HTTPS rewrites.",
									},
									"bic": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Inspect the visitor's browser for headers commonly associated with spammers and certain bots.",
									},
									"cache": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Whether to cache if expression matches.",
									},
									"content": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Content of the custom error response.",
									},
									"content_type": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Content-Type of the custom error response.",
									},
									"cookie_fields": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of cookie values to include as part of custom fields logging.",
									},
									"disable_apps": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn off all active Cloudflare Apps.",
									},
									"disable_railgun": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn off railgun feature of the Cloudflare Speed app.",
									},
									"disable_zaraz": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn off zaraz feature.",
									},
									"email_obfuscation": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the Cloudflare Email Obfuscation feature of the Cloudflare Scrape Shield app.",
									},
									"host_header": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Host Header that request origin receives.",
									},
									"hotlink_protection": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the hotlink protection feature.",
									},
									consts.IDSchemaKey: schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Identifier of the action parameter to modify.",
									},
									"increment": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "",
									},
									"mirage": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off Cloudflare Mirage of the Cloudflare Speed app.",
									},
									"opportunistic_encryption": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the Cloudflare Opportunistic Encryption feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.",
									},
									"origin_cache_control": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Enable or disable the use of a more compliant Cache Control parsing mechanism, enabled by default for most zones.",
									},
									"phases": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetPhaseValues())),
									},
									"polish": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Apply options from the Polish feature of the Cloudflare Speed app.",
									},
									"products": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: fmt.Sprintf("Products to target with the actions. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetActionParameterProductValues())),
									},
									"read_timeout": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Specifies a maximum timeout for reading content from an origin server.",
									},
									"request_fields": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of request headers to include as part of custom fields logging, in lowercase.",
									},
									"respect_strong_etags": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Respect strong ETags.",
									},
									"response_fields": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of response headers to include as part of custom fields logging, in lowercase.",
									},
									"rocket_loader": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off Cloudflare Rocket Loader in the Cloudflare Speed app.",
									},
									"rules": schema.MapAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "Map of managed WAF rule ID to comma-delimited string of ruleset rule IDs. Example: `rules = { \"efb7b8c949ac4650a09736fc376e9aee\" = \"5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760\" }`.",
									},
									"ruleset": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Which ruleset ID to target.",
									},
									"rulesets": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of managed WAF rule IDs to target. Only valid when the `\"action\"` is set to skip.",
									},
									"security_level": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Control options for the Security Level feature from the Security app.",
									},
									"server_side_excludes": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the Server Side Excludes feature of the Cloudflare Scrape Shield app.",
									},
									"ssl": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Control options for the SSL feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.",
									},
									"status_code": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "HTTP status code of the custom error response.",
									},
									"sxg": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Turn on or off the SXG feature.",
									},
									"origin_error_page_passthru": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Pass-through error page for origin.",
									},
									"version": schema.StringAttribute{
										Computed:            true,
										Optional:            true,
										MarkdownDescription: "Version of the ruleset to deploy.",
									},
								},
								Blocks: map[string]schema.Block{
									"algorithms": schema.ListNestedBlock{
										MarkdownDescription: "Compression algorithms to use in order of preference.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: fmt.Sprintf("Name of the compression algorithm to use. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"gzip", "brotli", "auto", "default", "none"})),
													Validators: []validator.String{
														stringvalidator.OneOf("gzip", "brotli", "auto", "default", "none"),
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtLeast(1),
										},
									},
									"uri": schema.ListNestedBlock{
										MarkdownDescription: "List of URI properties to configure for the ruleset rule when performing URL rewrite transformations.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"origin": schema.BoolAttribute{
													Optional: true,
												},
											},
											Blocks: map[string]schema.Block{
												"path": schema.ListNestedBlock{
													MarkdownDescription: "URI path configuration when performing a URL rewrite.",
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Static string value of the updated URI path or query string component.",
															},
															"expression": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Expression that defines the updated (dynamic) value of the URI path or query string component. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
															},
														},
													},
													Validators: []validator.List{
														listvalidator.SizeAtMost(1),
													},
												},
												"query": schema.ListNestedBlock{
													MarkdownDescription: "Query string configuration when performing a URL rewrite.",
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Static string value of the updated URI path or query string component.",
															},
															"expression": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Expression that defines the updated (dynamic) value of the URI path or query string component. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
															},
														},
													},
													Validators: []validator.List{
														listvalidator.SizeAtMost(1),
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"headers": schema.ListNestedBlock{
										MarkdownDescription: "List of HTTP header modifications to perform in the ruleset rule. Note: Headers are order dependent and must be provided sorted alphabetically ascending based on the `name` value.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Name of the HTTP request header to target.",
												},
												"value": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Static value to provide as the HTTP request header value.",
													Validators: []validator.String{
														stringvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("expression"))),
													},
												},
												"expression": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
													Validators: []validator.String{
														stringvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("value"))),
													},
												},
												"operation": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: fmt.Sprintf("Action to perform on the HTTP request header. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetRuleActionParametersHTTPHeaderOperationValues())),
												},
											},
										},
									},
									"matched_data": schema.ListNestedBlock{
										MarkdownDescription: "List of properties to configure WAF payload logging.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"public_key": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Public key to use within WAF Ruleset payload logging to view the HTTP request parameters. You can generate a public key [using the `matched-data-cli` command-line tool](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/command-line/generate-key-pair) or [in the Cloudflare dashboard](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/configure).",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"response": schema.ListNestedBlock{
										MarkdownDescription: "List of parameters that configure the response given to end users.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"status_code": schema.Int64Attribute{
													Optional:            true,
													MarkdownDescription: "HTTP status code to send in the response.",
												},
												"content_type": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "HTTP content type to send in the response.",
												},
												"content": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Body content to include in the response.",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"autominify": schema.ListNestedBlock{
										MarkdownDescription: "Indicate which file extensions to minify automatically.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"html": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "HTML minification.",
												},
												"css": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "CSS minification.",
												},
												"js": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "JS minification.",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"edge_ttl": schema.ListNestedBlock{
										MarkdownDescription: "List of edge TTL parameters to apply to the request.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"mode": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("override_origin", "respect_origin", "bypass_by_default")},
													MarkdownDescription: fmt.Sprintf("Mode of the edge TTL. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"override_origin", "respect_origin", "bypass_by_default"})),
												},
												"default": schema.Int64Attribute{
													Optional:            true,
													Validators:          []validator.Int64{int64validator.AtLeast(1)},
													MarkdownDescription: "Default edge TTL.",
												},
											},
											Validators: []validator.Object{EdgeTTLValidator{}},
											Blocks: map[string]schema.Block{
												"status_code_ttl": schema.ListNestedBlock{
													MarkdownDescription: "Edge TTL for the status codes.",
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"status_code": schema.Int64Attribute{
																Optional:            true,
																MarkdownDescription: "Status code for which the edge TTL is applied.",
																Validators: []validator.Int64{
																	int64validator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("status_code_range"))),
																},
															},
															"value": schema.Int64Attribute{
																Optional:            true,
																MarkdownDescription: "Status code edge TTL value.",
															},
														},
														Blocks: map[string]schema.Block{
															"status_code_range": schema.ListNestedBlock{
																MarkdownDescription: "Status code range for which the edge TTL is applied.",
																NestedObject: schema.NestedBlockObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.Int64Attribute{
																			Optional:            true,
																			MarkdownDescription: "From status code.",
																		},
																		"to": schema.Int64Attribute{
																			Optional:            true,
																			MarkdownDescription: "To status code.",
																		},
																	},
																},
																Validators: []validator.List{
																	listvalidator.SizeAtMost(1),
																	listvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("status_code"))),
																},
															},
														},
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"browser_ttl": schema.ListNestedBlock{
										MarkdownDescription: "List of browser TTL parameters to apply to the request.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"mode": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("override_origin", "respect_origin", "bypass")},
													MarkdownDescription: fmt.Sprintf("Mode of the browser TTL. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"override_origin", "respect_origin", "bypass"})),
												},
												"default": schema.Int64Attribute{
													Optional:            true,
													Validators:          []validator.Int64{int64validator.AtLeast(1)},
													MarkdownDescription: "Default browser TTL. This value is required when override_origin is set",
												},
											},
											Validators: []validator.Object{BrowserTTLValidator{}},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"serve_stale": schema.ListNestedBlock{
										MarkdownDescription: "List of serve stale parameters to apply to the request.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"disable_stale_while_updating": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "Disable stale while updating.",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"cache_key": schema.ListNestedBlock{
										MarkdownDescription: "List of cache key parameters to apply to the request.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"cache_by_device_type": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "Cache by device type.",
													Validators: []validator.Bool{
														boolvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("custom_key").AtAnyListIndex().AtName("user"))),
													},
												},
												"ignore_query_strings_order": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "Ignore query strings order.",
												},
												"cache_deception_armor": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "Cache deception armor.",
												},
											},
											Blocks: map[string]schema.Block{
												"custom_key": schema.ListNestedBlock{
													MarkdownDescription: "Custom key parameters for the request.",
													NestedObject: schema.NestedBlockObject{
														Blocks: map[string]schema.Block{
															"query_string": schema.ListNestedBlock{
																MarkdownDescription: "Query string parameters for the custom key.",
																NestedObject: schema.NestedBlockObject{
																	Attributes: map[string]schema.Attribute{
																		"include": schema.SetAttribute{
																			ElementType:         types.StringType,
																			Optional:            true,
																			MarkdownDescription: "List of query string parameters to include in the custom key.",
																			Validators: []validator.Set{
																				setvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("exclude"))),
																			},
																		},
																		"exclude": schema.SetAttribute{
																			ElementType:         types.StringType,
																			Optional:            true,
																			MarkdownDescription: "List of query string parameters to exclude from the custom key.",
																			Validators: []validator.Set{
																				setvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("include"))),
																			},
																		},
																	},
																},
																Validators: []validator.List{
																	listvalidator.SizeAtMost(1),
																},
															},
															"header": schema.ListNestedBlock{
																MarkdownDescription: "Header parameters for the custom key.",
																NestedObject: schema.NestedBlockObject{
																	Attributes: map[string]schema.Attribute{
																		"include": schema.SetAttribute{
																			ElementType:         types.StringType,
																			Optional:            true,
																			MarkdownDescription: "List of headers to include in the custom key.",
																		},
																		"check_presence": schema.SetAttribute{
																			ElementType:         types.StringType,
																			Optional:            true,
																			MarkdownDescription: "List of headers to check for presence in the custom key.",
																		},
																		"exclude_origin": schema.BoolAttribute{
																			Computed:            true,
																			Optional:            true,
																			MarkdownDescription: "Exclude the origin header from the custom key.",
																			PlanModifiers: []planmodifier.Bool{
																				defaults.DefaultBool(false),
																			},
																		},
																	},
																},
																Validators: []validator.List{
																	listvalidator.SizeAtMost(1),
																},
															},
															"cookie": schema.ListNestedBlock{
																MarkdownDescription: "Cookie parameters for the custom key.",
																NestedObject: schema.NestedBlockObject{
																	Attributes: map[string]schema.Attribute{
																		"include": schema.SetAttribute{
																			ElementType:         types.StringType,
																			Optional:            true,
																			MarkdownDescription: "List of cookies to include in the custom key.",
																		},
																		"check_presence": schema.SetAttribute{
																			ElementType:         types.StringType,
																			Optional:            true,
																			MarkdownDescription: "List of cookies to check for presence in the custom key.",
																		},
																	},
																},
																Validators: []validator.List{
																	listvalidator.SizeAtMost(1),
																},
															},
															"user": schema.ListNestedBlock{
																MarkdownDescription: "User parameters for the custom key.",
																NestedObject: schema.NestedBlockObject{
																	Attributes: map[string]schema.Attribute{
																		"device_type": schema.BoolAttribute{
																			Optional:            true,
																			MarkdownDescription: "Add device type to the custom key.",
																		},
																		"geo": schema.BoolAttribute{
																			Optional:            true,
																			MarkdownDescription: "Add geo data to the custom key.",
																		},
																		"lang": schema.BoolAttribute{
																			Optional:            true,
																			MarkdownDescription: "Add language data to the custom key.",
																		},
																	},
																},
																Validators: []validator.List{
																	listvalidator.SizeAtMost(1),
																	listvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtParent().AtParent().AtName("cache_by_device_type"))),
																},
															},
															"host": schema.ListNestedBlock{
																MarkdownDescription: "Host parameters for the custom key.",
																NestedObject: schema.NestedBlockObject{
																	Attributes: map[string]schema.Attribute{
																		"resolved": schema.BoolAttribute{
																			Optional:            true,
																			MarkdownDescription: "Resolve hostname to IP address.",
																		},
																	},
																},
																Validators: []validator.List{
																	listvalidator.SizeAtMost(1),
																},
															},
														},
													},
													Validators: []validator.List{
														listvalidator.SizeAtMost(1),
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"from_list": schema.ListNestedBlock{
										MarkdownDescription: "Use a list to lookup information for the action.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Name of the list.",
												},
												"key": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Expression to use for the list lookup.",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"from_value": schema.ListNestedBlock{
										MarkdownDescription: "Use a value to lookup information for the action.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"status_code": schema.Int64Attribute{
													Optional:            true,
													MarkdownDescription: "Status code for redirect.",
												},
												"preserve_query_string": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "Preserve query string for redirect URL.",
												},
											},
											Blocks: map[string]schema.Block{
												"target_url": schema.ListNestedBlock{
													MarkdownDescription: "Target URL for redirect.",
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Static value to provide as the HTTP request header value.",
																Validators: []validator.String{
																	stringvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("expression"))),
																},
															},
															"expression": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
																Validators: []validator.String{
																	stringvalidator.ConflictsWith(path.Expression(path.MatchRelative().AtParent().AtName("value"))),
																},
															},
														},
													},
													Validators: []validator.List{
														listvalidator.SizeAtMost(1),
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"overrides": schema.ListNestedBlock{
										MarkdownDescription: "List of override configurations to apply to the ruleset.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Optional:            true,
													MarkdownDescription: "Defines if the current ruleset-level override enables or disables the ruleset.",
												},
												"action": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: fmt.Sprintf("Action to perform in the rule-level override. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetRuleActionValues())),
												},
												"sensitivity_level": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: fmt.Sprintf("Sensitivity level to override for all ruleset rules. %s.", utils.RenderAvailableDocumentationValuesStringSlice([]string{"default", "medium", "low", "eoff"})),
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive("default", "medium", "low", "eoff"),
													},
												},
											},
											Blocks: map[string]schema.Block{
												"categories": schema.ListNestedBlock{
													MarkdownDescription: "List of tag-based overrides.",
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"category": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Tag name to apply the ruleset rule override to.",
															},
															"action": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: fmt.Sprintf("Action to perform in the tag-level override. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetRuleActionValues())),
																Validators: []validator.String{
																	stringvalidator.OneOfCaseInsensitive(cfv1.RulesetRuleActionValues()...),
																},
															},
															"enabled": schema.BoolAttribute{
																Optional:            true,
																MarkdownDescription: "Defines if the current tag-level override enables or disables the ruleset rules with the specified tag.",
															},
														},
													},
												},
												"rules": schema.ListNestedBlock{
													MarkdownDescription: "List of rule-based overrides.",
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															consts.IDSchemaKey: schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Rule ID to apply the override to.",
															},
															"action": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: fmt.Sprintf("Action to perform in the rule-level override. %s.", utils.RenderAvailableDocumentationValuesStringSlice(cfv1.RulesetRuleActionValues())),
																Validators: []validator.String{
																	stringvalidator.OneOfCaseInsensitive(cfv1.RulesetRuleActionValues()...),
																},
															},
															"enabled": schema.BoolAttribute{
																Optional:            true,
																MarkdownDescription: "Defines if the current rule-level override enables or disables the rule.",
															},
															"score_threshold": schema.Int64Attribute{
																Optional:            true,
																MarkdownDescription: "Anomaly score threshold to apply in the ruleset rule override. Only applicable to modsecurity-based rulesets.",
															},
															"sensitivity_level": schema.StringAttribute{
																Optional:            true,
																MarkdownDescription: "Sensitivity level for a ruleset rule override.",
															},
														},
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"origin": schema.ListNestedBlock{
										MarkdownDescription: "List of properties to change request origin.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Origin Hostname where request is sent.",
												},
												"port": schema.Int64Attribute{
													Optional:            true,
													MarkdownDescription: "Origin Port where request is sent.",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
									"sni": schema.ListNestedBlock{
										MarkdownDescription: "List of properties to manange Server Name Indication.",
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Value to define for SNI.",
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
						"ratelimit": schema.ListNestedBlock{
							MarkdownDescription: "List of parameters that configure HTTP rate limiting behaviour.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"characteristics": schema.SetAttribute{
										ElementType:         types.StringType,
										Optional:            true,
										MarkdownDescription: "List of parameters that define how Cloudflare tracks the request rate for this rule.",
									},
									"period": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "The period of time to consider (in seconds) when evaluating the request rate.",
									},
									"requests_per_period": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "The number of requests over the period of time that will trigger the Rate Limiting rule.",
									},
									"score_per_period": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "The maximum aggregate score over the period of time that will trigger Rate Limiting rule.",
									},
									"score_response_header_name": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Name of HTTP header in the response, set by the origin server, with the score for the current request.",
									},
									"mitigation_timeout": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Once the request rate is reached, the Rate Limiting rule blocks further requests for the period of time defined in this field.",
									},
									"counting_expression": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Criteria for counting HTTP requests to trigger the Rate Limiting action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
									},
									"requests_to_origin": schema.BoolAttribute{
										Optional:            true,
										Computed:            true,
										Default:             booldefault.StaticBool(false),
										MarkdownDescription: "Whether to include requests to origin within the Rate Limiting count.",
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
						"exposed_credential_check": schema.ListNestedBlock{
							MarkdownDescription: "List of parameters that configure exposed credential checks.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"username_expression": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: `Firewall Rules expression language based on Wireshark display filters for where to check for the "username" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).`,
									},
									"password_expression": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: `Firewall Rules expression language based on Wireshark display filters for where to check for the "password" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).`,
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
						"logging": schema.ListNestedBlock{
							MarkdownDescription: "List parameters to configure how the rule generates logs. Only valid for skip action.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Optional:            true,
										MarkdownDescription: "Override the default logging behavior when a rule is matched.",
									},
								},
							},
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
						},
					},
				},
			},
		},
	}
}
