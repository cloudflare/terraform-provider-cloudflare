package provider

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRulesetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"account_id"},
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
		"shareable_entitlement_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of entitlement that is shareable between entities.",
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
											},
											"action": {
												Type:         schema.TypeString,
												Optional:     true,
												ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
												Description:  fmt.Sprintf("Action to perform in the rule-level override. %s", renderAvailableDocumentationValuesStringSlice(cloudflare.RulesetRuleActionValues())),
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
								},
							},
						},
					},
				},
			},
		},
	}
}
