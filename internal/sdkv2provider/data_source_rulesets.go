package sdkv2provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRulesetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   consts.AccountIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   consts.ZoneIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Name of the ruleset.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Brief summary of the ruleset and its intended use.",
		},
		"kind": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(cloudflare.RulesetKindValues(), false),
			Description:  fmt.Sprintf("Type of Ruleset to create. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetKindValues())),
		},
		"phase": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(cloudflare.RulesetPhaseValues(), false),
			Description:  fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
		},
		"rules": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of rules to apply to the ruleset.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Unique rule identifier.",
					},
					"version": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Version of the ruleset to deploy.",
					},
					"ref": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Rule reference.",
					},
					"enabled": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether the rule is active.",
					},
					"action": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
						Description:  fmt.Sprintf("Action to perform in the ruleset rule. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionValues())),
					},
					"expression": {
						Description: "Criteria for an HTTP request to trigger the ruleset rule action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions",
						Type:        schema.TypeString,
						Required:    true,
					},
					"description": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Brief summary of the ruleset rule and its intended use.",
					},
					"last_updated": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The most recent update to this rule.",
					},
					"action_parameters": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "List of parameters that configure the behavior of the ruleset rule action.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"id": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Identifier of the action parameter to modify.",
								},
								"products": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: fmt.Sprintf("Products to target with the actions. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetActionParameterProductValues())),
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"phases": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"uri": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of URI properties to configure for the ruleset rule when performing URL rewrite transformations.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"path": {
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Description: "URI path configuration when performing a URL rewrite.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"value": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Static string value of the updated URI path or query string component.",
														},
														"expression": {
															Description: "Expression that defines the updated (dynamic) value of the URI path or query string component. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"query": {
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Description: "Query string configuration when performing a URL rewrite.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"value": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Static string value of the updated URI path or query string component.",
														},
														"expression": {
															Description: "Expression that defines the updated (dynamic) value of the URI path or query string component. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"origin": {
												Type:     schema.TypeBool,
												Optional: true,
											},
										},
									},
								},
								"headers": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "List of HTTP header modifications to perform in the ruleset rule.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Name of the HTTP request header to target.",
											},
											"value": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Static value to provide as the HTTP request header value. Conflicts with `\"expression\"`.",
											},
											"expression": {
												Description: "Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions. Conflicts with `\"value\"`.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"operation": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: fmt.Sprintf("Action to perform on the HTTP request header. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionParametersHTTPHeaderOperationValues())),
											},
										},
									},
								},
								"increment": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"version": {
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
									Description: "Version of the ruleset to deploy.",
								},
								"ruleset": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Which ruleset ID to target.",
								},
								"rulesets": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "List of managed WAF rule IDs to target. Only valid when the `\"action\"` is set to skip",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"rules": {
									Type:        schema.TypeMap,
									Optional:    true,
									Description: "Map of managed WAF rule ID to comma-delimited string of ruleset rule IDs. Example: `rules = { \"efb7b8c949ac4650a09736fc376e9aee\" = \"5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760\" }`",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"overrides": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of override configurations to apply to the ruleset.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"enabled": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Defines if the current ruleset-level override enables or disables the ruleset.",
												Deprecated:  "Use `status` instead. Continuing to use `enabled` will result in an inconsistent state for your Ruleset configuration.",
											},
											"status": {
												Type:         schema.TypeString,
												Optional:     true,
												Default:      "",
												ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
												Description:  fmt.Sprintf("Defines if the current ruleset-level override enables or disables the ruleset. %s", renderAvailableDocumentationValuesStringSlice([]string{"enabled", "disabled"})),
											},
											"action": {
												Type:         schema.TypeString,
												Optional:     true,
												ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
												Description:  fmt.Sprintf("Action to perform in the rule-level override. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionValues())),
											},
											"sensitivity_level": {
												Type:         schema.TypeString,
												Optional:     true,
												ValidateFunc: validation.StringInSlice([]string{"default", "medium", "low", "eoff"}, false),
												Description:  fmt.Sprintf("Sensitivity level to override for all ruleset rules. %s", renderAvailableDocumentationValuesStringSlice([]string{"default", "medium", "low", "eoff"})),
											},
											"categories": {
												Type:        schema.TypeList,
												Optional:    true,
												Description: "List of tag-based overrides.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"category": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Tag name to apply the ruleset rule override to.",
														},
														"action": {
															Type:         schema.TypeString,
															Optional:     true,
															ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
															Description:  fmt.Sprintf("Action to perform in the tag-level override. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionValues())),
														},
														"enabled": {
															Type:        schema.TypeBool,
															Optional:    true,
															Description: "Defines if the current tag-level override enables or disables the ruleset rules with the specified tag.",
															Deprecated:  "Use `status` instead. Continuing to use `enabled` will result in an inconsistent state for your Ruleset configuration.",
														},
														"status": {
															Type:         schema.TypeString,
															Optional:     true,
															Default:      "",
															ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
															Description:  fmt.Sprintf("Defines if the current tag-level override enables or disables the ruleset rules with the specified tag. %s", renderAvailableDocumentationValuesStringSlice([]string{"enabled", "disabled"})),
														},
													},
												},
											},
											"rules": {
												Type:        schema.TypeList,
												Optional:    true,
												Description: "List of rule-based overrides.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"id": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Rule ID to apply the override to.",
														},
														"action": {
															Type:         schema.TypeString,
															Optional:     true,
															ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
															Description:  fmt.Sprintf("Action to perform in the rule-level override. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionValues())),
														},
														"enabled": {
															Type:        schema.TypeBool,
															Optional:    true,
															Description: "Defines if the current rule-level override enables or disables the rule.",
															Deprecated:  "Use `status` instead. Continuing to use `enabled` will result in an inconsistent state for your Ruleset configuration.",
														},
														"status": {
															Type:         schema.TypeString,
															Optional:     true,
															Default:      "",
															ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
															Description:  fmt.Sprintf("Defines if the current rule-level override enables or disables the rule. %s", renderAvailableDocumentationValuesStringSlice([]string{"enabled", "disabled"})),
														},
														"score_threshold": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "Anomaly score threshold to apply in the ruleset rule override. Only applicable to modsecurity-based rulesets.",
														},
														"sensitivity_level": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Sensitivity level for a ruleset rule override.",
														},
													},
												},
											},
										},
									},
								},
								"matched_data": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of properties to configure WAF payload logging.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"public_key": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Public key to use within WAF Ruleset payload logging to view the HTTP request parameters. You can generate a public key [using the `matched-data-cli` command-line tool](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/command-line/generate-key-pair) or [in the Cloudflare dashboard](https://developers.cloudflare.com/waf/managed-rulesets/payload-logging/configure)",
											},
										},
									},
								},
								"response": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "List of parameters that configure the response given to end users",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"status_code": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "HTTP status code to send in the response.",
											},
											"content_type": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "HTTP content type to send in the response.",
											},
											"content": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Body content to include in the response.",
											},
										},
									},
								},
								"content": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Content of the custom error response",
								},
								"content_type": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Content-Type of the custom error response",
								},
								"status_code": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP status code of the custom error response",
								},
								"host_header": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Host Header that request origin receives.",
								},
								"origin": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of properties to change request origin.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"host": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Origin Hostname where request is sent.",
											},
											"port": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "Origin Port where request is sent.",
											},
										},
									},
								},
								"sni": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of properties to manange Server Name Indication.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"value": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Value to define for SNI.",
											},
										},
									},
								},
								"request_fields": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "List of request headers to include as part of custom fields logging, in lowercase.",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"response_fields": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "List of response headers to include as part of custom fields logging, in lowercase.",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"cookie_fields": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "List of cookie values to include as part of custom fields logging.",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"cache": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Whether to cache if expression matches.",
								},
								"automatic_https_rewrites": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off Cloudflare Automatic HTTPS rewrites.",
								},
								"autominify": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Indicate which file extensions to minify automatically.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"html": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "HTML minification.",
											},
											"css": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "SSL minification.",
											},
											"js": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "JS minification.",
											},
										},
									},
								},
								"bic": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Inspect the visitor's browser for headers commonly associated with spammers and certain bots.",
								},
								"disable_apps": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn off all active Cloudflare Apps.",
								},
								"disable_zaraz": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn off zaraz feature.",
								},
								"disable_railgun": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn off railgun feature of the Cloudflare Speed app.",
								},
								"email_obfuscation": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off the Cloudflare Email Obfuscation feature of the Cloudflare Scrape Shield app.",
								},
								"mirage": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off Cloudflare Mirage of the Cloudflare Speed app.",
								},
								"opportunistic_encryption": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off the Cloudflare Opportunistic Encryption feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.",
								},
								"polish": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Apply options from the Polish feature of the Cloudflare Speed app.",
								},
								"rocket_loader": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off Cloudflare Rocket Loader in the Cloudflare Speed app.",
								},
								"security_level": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Control options for the Security Level feature from the Security app.",
								},
								"server_side_excludes": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off the Server Side Excludes feature of the Cloudflare Scrape Shield app.",
								},
								"ssl": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Control options for the SSL feature of the Edge Certificates tab in the Cloudflare SSL/TLS app.",
								},
								"sxg": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off the SXG feature.",
								},
								"hotlink_protection": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Turn on or off the hotlink protection feature.",
								},
								"edge_ttl": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of edge TTL parameters to apply to the request.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"mode": {
												Type:        schema.TypeString,
												Required:    true,
												Description: "Mode of the edge TTL.",
											},
											"default": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "Default edge TTL",
											},
											"status_code_ttl": {
												Type:        schema.TypeList,
												Optional:    true,
												Description: "Edge TTL for the status codes.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"status_code": {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "Status code for which the edge TTL is applied. Conflicts with \"status_code_range\".",
														},
														"status_code_range": {
															Type:        schema.TypeList,
															Optional:    true,
															Description: "Status code range for which the edge TTL is applied. Conflicts with \"status_code\".",
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"from": {
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Description: "From status code.",
																	},
																	"to": {
																		Type:        schema.TypeInt,
																		Optional:    true,
																		Description: "To status code.",
																	},
																},
															},
														},
														"value": {
															Type:        schema.TypeInt,
															Required:    true,
															Description: "Status code edge TTL value.",
														},
													},
												},
											},
										},
									},
								},
								"browser_ttl": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of browser TTL parameters to apply to the request.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"mode": {
												Type:        schema.TypeString,
												Required:    true,
												Description: "Mode of the browser TTL.",
											},
											"default": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "Default browser TTL.",
											},
										},
									},
								},
								"serve_stale": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "List of serve stale parameters to apply to the request.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"disable_stale_while_updating": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Disable stale while updating.",
											},
										},
									},
								},
								"read_timeout": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Sets the timeout value for reading content from an origin server.",
								},
								"respect_strong_etags": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Respect strong ETags.",
								},
								"cache_key": {
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Description: "List of cache key parameters to apply to the request.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"cache_by_device_type": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Cache by device type. Conflicts with \"custom_key.user.device_type\".",
											},
											"ignore_query_strings_order": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Ignore query strings order.",
											},
											"cache_deception_armor": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Cache deception armor.",
											},
											"custom_key": {
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Description: "Custom key parameters for the request.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"query_string": {
															Type:        schema.TypeList,
															Optional:    true,
															MaxItems:    1,
															Description: "Query string parameters for the custom key.",
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"include": {
																		Type:        schema.TypeList,
																		Optional:    true,
																		Description: "List of query string parameters to include in the custom key. Conflicts with \"exclude\".",
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																	"exclude": {
																		Type:        schema.TypeList,
																		Optional:    true,
																		Description: "List of query string parameters to exclude from the custom key. Conflicts with \"include\".",
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
														"header": {
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Description: "Header parameters for the custom key.",
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"include": {
																		Type:        schema.TypeList,
																		Optional:    true,
																		Description: "List of headers to include in the custom key.",
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																	"check_presence": {
																		Type:        schema.TypeList,
																		Optional:    true,
																		Description: "List of headers to check for presence in the custom key.",
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																	"exclude_origin": {
																		Type:        schema.TypeBool,
																		Optional:    true,
																		Description: "Exclude the origin header from the custom key.",
																	},
																},
															},
														},
														"cookie": {
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Description: "Cookie parameters for the custom key.",
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"include": {
																		Type:        schema.TypeList,
																		Optional:    true,
																		Description: "List of cookies to include in the custom key.",
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																	"check_presence": {
																		Type:        schema.TypeList,
																		Optional:    true,
																		Description: "List of cookies to check for presence in the custom key.",
																		Elem: &schema.Schema{
																			Type: schema.TypeString,
																		},
																	},
																},
															},
														},
														"user": {
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Description: "User parameters for the custom key.",
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"device_type": {
																		Type:        schema.TypeBool,
																		Optional:    true,
																		Description: "Add device type to the custom key. Conflicts with \"cache_key.cache_by_device_type\".",
																	},
																	"geo": {
																		Type:        schema.TypeBool,
																		Optional:    true,
																		Description: "Add geo data to the custom key.",
																	},
																	"lang": {
																		Type:        schema.TypeBool,
																		Optional:    true,
																		Description: "Add language data to the custom key.",
																	},
																},
															},
														},
														"host": {
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Description: "Host parameters for the custom key.",
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"resolved": {
																		Type:        schema.TypeBool,
																		Optional:    true,
																		Description: "Resolve hostname to IP address.",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								"origin_cache_control": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Sets a more compliant mode for parsing Cache Control headers",
								},
								"origin_error_page_passthru": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Pass-through error page for origin.",
								},
								"from_list": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "Use a list to lookup information for the action.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Description: "Name of the list.",
												Required:    true,
											},
											"key": {
												Type:        schema.TypeString,
												Description: "Expression to use for the list lookup.",
												Required:    true,
											},
										},
									},
								},
								"from_value": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "Use a value to lookup information for the action.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"status_code": {
												Type:        schema.TypeInt,
												Description: "Status code for redirect.",
												Optional:    true,
											},
											"target_url": {
												Type:        schema.TypeList,
												Optional:    true,
												MaxItems:    1,
												Description: "Target URL for redirect.",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"value": {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Static value to provide as the HTTP request header value. Conflicts with `\"expression\"`.",
														},
														"expression": {
															Description: "Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions. Conflicts with `\"value\"`.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"preserve_query_string": {
												Type:        schema.TypeBool,
												Description: "Preserve query string for redirect URL.",
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
					"ratelimit": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "List of parameters that configure HTTP rate limiting behaviour.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"characteristics": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "List of parameters that define how Cloudflare tracks the request rate for this rule.",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"period": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "The period of time to consider (in seconds) when evaluating the request rate.",
								},
								"requests_per_period": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "The number of requests over the period of time that will trigger the Rate Limiting rule.",
								},
								"score_per_period": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "The maximum aggregate score over the period of time that will trigger Rate Limiting rule.",
								},
								"score_response_header_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Name of HTTP header in the response, set by the origin server, with the score for the current request.",
								},
								"mitigation_timeout": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Once the request rate is reached, the Rate Limiting rule blocks further requests for the period of time defined in this field.",
								},
								"counting_expression": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Criteria for counting HTTP requests to trigger the Rate Limiting action. Uses the Firewall Rules expression language based on Wireshark display filters. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions.",
								},
								"requests_to_origin": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Whether to include requests to origin within the Rate Limiting count.",
								},
							},
						},
					},
					"exposed_credential_check": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "List of parameters that configure exposed credential checks.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"username_expression": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Firewall Rules expression language based on Wireshark display filters for where to check for the \"username\" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).",
								},
								"password_expression": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Firewall Rules expression language based on Wireshark display filters for where to check for the \"password\" value. Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language).",
								},
							},
						},
					},
					"logging": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "List parameters to configure how the rule generates logs.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Override the default logging behavior when a rule is matched.",
									Deprecated:  "Use `status` instead. Continuing to use `enabled` will result in an inconsistent state for your Ruleset configuration.",
								},
								"status": {
									Type:         schema.TypeString,
									Optional:     true,
									Default:      "",
									ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
									Description:  fmt.Sprintf("Override the default logging behavior when a rule is matched. %s", renderAvailableDocumentationValuesStringSlice([]string{"enabled", "disabled"})),
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareRulesets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareRulesetsRead,

		Description: heredoc.Doc(`
			Use this datasource to lookup Rulesets in an account or zone.
		`),

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description:  consts.AccountIDSchemaDescription,
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"zone_id", "account_id"},
			},
			"zone_id": {
				Description:  consts.ZoneIDSchemaDescription,
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"zone_id", "account_id"},
			},

			"include_rules": {
				Description: "Include rule data in response",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the Ruleset to target.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the ruleset.",
						},
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: fmt.Sprintf("Point in the request/response lifecycle where the ruleset will be created. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
						},
						"kind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: fmt.Sprintf("Type of Ruleset to create. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetKindValues())),
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Version of the ruleset to filter on.",
						},
					},
				},
			},

			"rulesets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID of the ruleset.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the ruleset.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Brief summary of the ruleset and its intended use.",
						},
						"kind": {
							Type:        schema.TypeString,
							Required:    true,
							Description: fmt.Sprintf("Type of Ruleset. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetKindValues())),
						},
						"version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Version of the ruleset.",
						},
						"phase": {
							Type:        schema.TypeString,
							Required:    true,
							Description: fmt.Sprintf("Point in the request/response lifecycle where the ruleset executes. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetPhaseValues())),
						},
						"rules": resourceCloudflareRulesetSchema()["rules"],
					},
				},
			},
		},
	}
}

func dataSourceCloudflareRulesetsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	accountID := d.Get("account_id").(string)
	includeRules := d.Get("include_rules").(bool)
	filter, err := expandFilterRulesets(d.Get("filter"))
	if err != nil {
		return diag.FromErr(err)
	}

	var identifier *cloudflare.ResourceContainer
	if accountID != "" {
		identifier = cloudflare.AccountIdentifier(accountID)
	} else {
		identifier = cloudflare.ZoneIdentifier(zoneID)
	}

	rulesetsList, err := client.ListRulesets(ctx, identifier, cloudflare.ListRulesetsParams{})
	if err != nil {
		return diag.FromErr(err)
	}

	rulesets := make([]interface{}, 0)
	rulesetIds := make([]string, 0)
	for _, ruleset := range rulesetsList {
		if filter.ID != "" && filter.ID != ruleset.ID {
			continue
		}
		if filter.Phase != "" && filter.Phase != ruleset.Phase {
			continue
		}
		if filter.Kind != "" && filter.Kind != ruleset.Kind {
			continue
		}
		if filter.Version != "" && filter.Version != cloudflare.String(ruleset.Version) {
			continue
		}
		if filter.Name != nil && !filter.Name.Match([]byte(ruleset.Name)) {
			continue
		}

		rulesetIds = append(rulesetIds, ruleset.ID)
		resultRuleset := map[string]interface{}{
			"id":          ruleset.ID,
			"name":        ruleset.Name,
			"description": ruleset.Description,
			"kind":        ruleset.Kind,
			"version":     ruleset.Version,
			"phase":       ruleset.Phase,
		}

		if includeRules {
			fullRuleset, err := client.GetRuleset(ctx, identifier, ruleset.ID)
			if err != nil {
				return diag.FromErr(err)
			}

			rules := make([]interface{}, 0)
			for _, rule := range fullRuleset.Rules {
				fullRulesetRule := map[string]interface{}{
					"id":           rule.ID,
					"version":      rule.Version,
					"action":       rule.Action,
					"expression":   rule.Expression,
					"description":  rule.Description,
					"last_updated": rule.LastUpdated.String(),
					"ref":          rule.Ref,
					"enabled":      rule.Enabled,
				}

				if rule.RateLimit != nil {
					rl := make([]interface{}, 0)
					fullRulesetRule["ratelimit"] = append(rl, map[string]interface{}{
						"characteristics":     rule.RateLimit.Characteristics,
						"requests_per_period": rule.RateLimit.RequestsPerPeriod,
						"period":              rule.RateLimit.Period,
						"mitigation_timeout":  rule.RateLimit.MitigationTimeout,
						"counting_expression": rule.RateLimit.CountingExpression,
						"requests_to_origin":  rule.RateLimit.RequestsToOrigin,
					})
				}

				if rule.ExposedCredentialCheck != nil {
					ecc := make([]interface{}, 0)
					fullRulesetRule["exposed_credential_check"] = append(ecc, map[string]interface{}{
						"username_expression": rule.ExposedCredentialCheck.UsernameExpression,
						"password_expression": rule.ExposedCredentialCheck.PasswordExpression,
					})
				}
				if rule.Logging != nil {
					lg := make([]interface{}, 0)
					fullRulesetRule["logging"] = append(lg, map[string]interface{}{
						"enabled": rule.Logging.Enabled,
						"status":  loggingStatus(rule.Logging),
					})
				}
				rules = append(rules, fullRulesetRule)
			}
			resultRuleset["rules"] = rules
		}

		rulesets = append(rulesets, resultRuleset)
	}

	err = d.Set("rulesets", rulesets)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting Ruleset Ids: %w", err))
	}

	d.SetId(stringListChecksum(rulesetIds))
	return nil
}

func loggingStatus(logging *cloudflare.RulesetRuleLogging) string {
	if logging == nil || logging.Enabled == nil {
		return "default"
	}
	if *logging.Enabled {
		return "enabled"
	}
	return "disabled"
}

func expandFilterRulesets(d interface{}) (*searchFilterRulesets, error) {
	cfg := d.([]interface{})
	filter := &searchFilterRulesets{}
	if len(cfg) == 0 || cfg[0] == nil {
		return filter, nil
	}

	m := cfg[0].(map[string]interface{})
	if name, ok := m["name"]; ok {
		match, err := regexp.Compile(name.(string))
		if err != nil {
			return nil, err
		}

		filter.Name = match
	}

	if kind, ok := m["kind"]; ok {
		filter.Kind = kind.(string)
	}
	if phase, ok := m["phase"]; ok {
		filter.Phase = phase.(string)
	}
	if id, ok := m["id"]; ok {
		filter.ID = id.(string)
	}
	if version, ok := m["version"]; ok {
		filter.Version = version.(string)
	}

	return filter, nil
}

type searchFilterRulesets struct {
	Name    *regexp.Regexp
	Kind    string
	Phase   string
	ID      string
	Version string
}
