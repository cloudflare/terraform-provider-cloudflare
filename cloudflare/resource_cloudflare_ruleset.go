package cloudflare

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudflareRuleset() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareRulesetCreate,
		Read:   resourceCloudflareRulesetRead,
		Update: resourceCloudflareRulesetUpdate,
		Delete: resourceCloudflareRulesetDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareRulesetImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"zone_id"},
			},
			"zone_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"account_id"},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kind": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(cloudflare.RulesetKindValues(), false),
			},
			"phase": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(cloudflare.RulesetPhaseValues(), false),
			},
			"shareable_entitlement_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"action": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
						},
						"expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"action_parameters": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"products": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"uri": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"path": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"expression": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"query": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"expression": {
																Type:     schema.TypeString,
																Optional: true,
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
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"expression": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"operation": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"increment": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ruleset": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"rulesets": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"rules": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"overrides": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"action": {
													Type:         schema.TypeString,
													Optional:     true,
													ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
												},
												"categories": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"category": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"action": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
															},
															"enabled": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
												"rules": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"action": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(cloudflare.RulesetRuleActionValues(), false),
															},
															"enabled": {
																Type:     schema.TypeBool,
																Optional: true,
															},
															"score_threshold": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"sensitivity_level": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"matched_data": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"public_key": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"ratelimit": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"characteristics": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"period": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"requests_per_period": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mitigation_timeout": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mitigation_expression": {
										Type:     schema.TypeString,
										Optional: true,
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

func resourceCloudflareRulesetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	zoneID := d.Get("zone_id").(string)

	rs := cloudflare.Ruleset{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Kind:        d.Get("kind").(string),
		Phase:       d.Get("phase").(string),
	}

	rules, err := buildRulesetRulesFromResource(d.Get("phase").(string), d.Get("rules"))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error building ruleset from resource"))
	}

	if len(rules) > 0 {
		rs.Rules = rules
	}

	var ruleset cloudflare.Ruleset
	if accountID != "" {
		ruleset, err = client.CreateAccountRuleset(context.Background(), accountID, rs)
	} else {
		ruleset, err = client.CreateZoneRuleset(context.Background(), zoneID, rs)
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error creating ruleset %s", d.Get("name").(string)))
	}

	rulesetEntryPoint := cloudflare.Ruleset{
		Description: d.Get("description").(string),
		Rules:       rules,
	}

	if accountID != "" {
		_, err = client.UpdateAccountRulesetPhase(context.Background(), accountID, rs.Phase, rulesetEntryPoint)
	} else {
		_, err = client.UpdateZoneRulesetPhase(context.Background(), zoneID, rs.Phase, rulesetEntryPoint)
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error updating ruleset phase entrypoint %s", d.Get("name").(string)))
	}

	d.SetId(ruleset.ID)

	return resourceCloudflareRulesetRead(d, meta)
}

func resourceCloudflareRulesetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, errors.New("Import is not yet supported for Rulesets")
}

func resourceCloudflareRulesetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	zoneID := d.Get("zone_id").(string)

	var ruleset cloudflare.Ruleset
	var err error

	if accountID != "" {
		ruleset, err = client.GetAccountRuleset(context.Background(), accountID, d.Id())
	} else {
		ruleset, err = client.GetZoneRuleset(context.Background(), zoneID, d.Id())
	}

	if err != nil {
		if strings.Contains(err.Error(), "could not find ruleset") {
			log.Printf("[INFO] Ruleset %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return errors.Wrap(err, fmt.Sprintf("error reading ruleset ID: %s", d.Id()))
	}

	d.Set("name", ruleset.Name)
	d.Set("description", ruleset.Description)

	if err := d.Set("rules", buildStateFromRulesetRules(ruleset.Rules)); err != nil {
		log.Fatalf("failed to set rules: %s", err)
	}

	return nil
}

func resourceCloudflareRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	zoneID := d.Get("zone_id").(string)

	rules, err := buildRulesetRulesFromResource(d.Get("phase").(string), d.Get("rules"))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error building ruleset from resource"))
	}

	description := d.Get("description").(string)
	if accountID != "" {
		_, err = client.UpdateAccountRuleset(context.Background(), accountID, d.Id(), description, rules)
	} else {
		_, err = client.UpdateZoneRuleset(context.Background(), zoneID, d.Id(), description, rules)
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error updating ruleset with ID %q", d.Id()))
	}

	return resourceCloudflareRulesetRead(d, meta)
}

func resourceCloudflareRulesetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	zoneID := d.Get("zone_id").(string)
	var err error

	if accountID != "" {
		err = client.DeleteAccountRuleset(context.Background(), accountID, d.Id())
	} else {
		err = client.DeleteZoneRuleset(context.Background(), zoneID, d.Id())
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error deleting ruleset with ID %q", d.Id()))
	}

	return nil
}

// buildStateFromRulesetRules receives the current ruleset rules and returns an
// interface for the state file
func buildStateFromRulesetRules(rules []cloudflare.RulesetRule) interface{} {
	var rulesData []map[string]interface{}
	for _, r := range rules {
		rule := map[string]interface{}{
			"id":         r.ID,
			"expression": r.Expression,
			"action":     r.Action,
			"enabled":    r.Enabled,
		}

		if r.Description != "" {
			rule["description"] = r.Description
		}

		if !reflect.ValueOf(r.ActionParameters).IsNil() {
			var actionParameters []map[string]interface{}
			var overrides []map[string]interface{}
			var idBasedOverrides []map[string]interface{}
			var categoryBasedOverrides []map[string]interface{}
			var headers []map[string]interface{}
			var uri []map[string]interface{}
			var matchedData []map[string]interface{}
			actionParameterRules := make(map[string]string)

			if !reflect.ValueOf(r.ActionParameters.Overrides).IsNil() {
				for _, overrideRule := range r.ActionParameters.Overrides.Rules {
					idBasedOverrides = append(idBasedOverrides, map[string]interface{}{
						"id":                overrideRule.ID,
						"action":            overrideRule.Action,
						"enabled":           overrideRule.Enabled,
						"score_threshold":   overrideRule.ScoreThreshold,
						"sensitivity_level": overrideRule.SensitivityLevel,
					})
				}

				for _, overrideRule := range r.ActionParameters.Overrides.Categories {
					categoryBasedOverrides = append(categoryBasedOverrides, map[string]interface{}{
						"category": overrideRule.Category,
						"action":   overrideRule.Action,
						"enabled":  overrideRule.Enabled,
					})
				}

				overrides = append(overrides, map[string]interface{}{
					"categories": categoryBasedOverrides,
					"rules":      idBasedOverrides,
					"enabled":    r.ActionParameters.Overrides.Enabled,
					"action":     r.ActionParameters.Overrides.Action,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.Headers).IsNil() {
				sortedHeaders := make([]string, 0, len(r.ActionParameters.Headers))
				for headerName := range r.ActionParameters.Headers {
					sortedHeaders = append(sortedHeaders, headerName)
				}
				sort.Strings(sortedHeaders)

				for _, headerName := range sortedHeaders {
					headers = append(headers, map[string]interface{}{
						"name":       headerName,
						"value":      r.ActionParameters.Headers[headerName].Value,
						"expression": r.ActionParameters.Headers[headerName].Expression,
						"operation":  r.ActionParameters.Headers[headerName].Operation,
					})
				}
			}

			if !reflect.ValueOf(r.ActionParameters.URI).IsNil() {
				var query []map[string]interface{}
				var path []map[string]interface{}
				originValue := false
				if r.ActionParameters.URI.Origin {
					originValue = true
				}

				if !reflect.ValueOf(r.ActionParameters.URI.Query).IsNil() {
					query = append(query, map[string]interface{}{
						"value":      r.ActionParameters.URI.Query.Value,
						"expression": r.ActionParameters.URI.Query.Expression,
					})
				}

				if !reflect.ValueOf(r.ActionParameters.URI.Path).IsNil() {
					path = append(path, map[string]interface{}{
						"value":      r.ActionParameters.URI.Path.Value,
						"expression": r.ActionParameters.URI.Path.Expression,
					})
				}

				uri = append(uri, map[string]interface{}{
					"origin": originValue,
					"query":  query,
					"path":   path,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.MatchedData).IsNil() {
				matchedData = append(matchedData, map[string]interface{}{
					"public_key": r.ActionParameters.MatchedData.PublicKey,
				})
			}

			if !reflect.ValueOf(r.ActionParameters.Rules).IsNil() {
				for k, v := range r.ActionParameters.Rules {
					actionParameterRules[k] = strings.Join(v, ",")
				}
			}

			actionParameters = append(actionParameters, map[string]interface{}{
				"id":           r.ActionParameters.ID,
				"increment":    r.ActionParameters.Increment,
				"headers":      headers,
				"overrides":    overrides,
				"products":     r.ActionParameters.Products,
				"ruleset":      r.ActionParameters.Ruleset,
				"rulesets":     r.ActionParameters.Rulesets,
				"rules":        actionParameterRules,
				"uri":          uri,
				"matched_data": matchedData,
				"version":      r.ActionParameters.Version,
			})

			rule["action_parameters"] = actionParameters
		}

		if !reflect.ValueOf(r.RateLimit).IsNil() {
			var rateLimit []map[string]interface{}

			rateLimit = append(rateLimit, map[string]interface{}{
				"characteristics":     r.RateLimit.Characteristics,
				"period":              r.RateLimit.Period,
				"requests_per_period": r.RateLimit.RequestsPerPeriod,
				"mitigation_timeout":  r.RateLimit.MitigationTimeout,
			})

			rule["ratelimit"] = rateLimit
		}

		rulesData = append(rulesData, rule)
	}

	return rulesData
}

// receives the resource config and builds a ruleset rule array
func buildRulesetRulesFromResource(phase string, r interface{}) ([]cloudflare.RulesetRule, error) {
	var rulesetRules []cloudflare.RulesetRule

	rules, ok := r.([]interface{})
	if !ok {
		return nil, errors.New("unable to create interface array type assertion")
	}

	for _, v := range rules {
		var rule cloudflare.RulesetRule

		resourceRule, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("unable to create interface map type assertion for rule")
		}

		if len(resourceRule["action_parameters"].([]interface{})) > 0 {
			rule.ActionParameters = &cloudflare.RulesetRuleActionParameters{}
			for _, parameter := range resourceRule["action_parameters"].([]interface{}) {
				for pKey, pValue := range parameter.(map[string]interface{}) {
					switch pKey {
					case "id":
						rule.ActionParameters.ID = pValue.(string)
					case "version":
						rule.ActionParameters.Version = pValue.(string)
					case "ruleset":
						rule.ActionParameters.Ruleset = pValue.(string)
					case "rulesets":
						var rulesetsValues []string
						for _, v := range pValue.(*schema.Set).List() {
							rulesetsValues = append(rulesetsValues, v.(string))
						}
						rule.ActionParameters.Rulesets = rulesetsValues
					case "rules":
						apRules := make(map[string][]string)
						for name, data := range pValue.(map[string]interface{}) {
							// regex (not string.Split) needs to be used here to account for
							// whitespace from the end user in the map value.
							split := regexp.MustCompile("\\s*,\\s*").Split(data.(string), -1)
							ruleValues := []string{}

							for i := range split {
								ruleValues = append(ruleValues, split[i])
							}

							apRules[name] = ruleValues
						}

						rule.ActionParameters.Rules = apRules
					case "increment":
						rule.ActionParameters.Increment = pValue.(int)
					case "overrides":
						var overrideConfiguration cloudflare.RulesetRuleActionParametersOverrides
						var categories []cloudflare.RulesetRuleActionParametersCategories
						var rules []cloudflare.RulesetRuleActionParametersRules

						for _, overrideParamValue := range pValue.([]interface{}) {
							if phase != string(cloudflare.RulesetPhaseDDoSL7) {
								overrideConfiguration.Enabled = &[]bool{overrideParamValue.(map[string]interface{})["enabled"].(bool)}[0]
							}

							if val, ok := overrideParamValue.(map[string]interface{})["action"]; ok {
								overrideConfiguration.Action = val.(string)
							}

							// Category based overrides
							if val, ok := overrideParamValue.(map[string]interface{})["categories"]; ok {
								for _, category := range val.([]interface{}) {
									cData := category.(map[string]interface{})
									categories = append(categories, cloudflare.RulesetRuleActionParametersCategories{
										Category: cData["category"].(string),
										Action:   cData["action"].(string),
										Enabled:  cData["enabled"].(bool),
									})
								}
							}

							// Rule ID based overrides
							if val, ok := overrideParamValue.(map[string]interface{})["rules"]; ok {
								for _, rule := range val.([]interface{}) {
									rData := rule.(map[string]interface{})

									var enabled *bool
									if phase != string(cloudflare.RulesetPhaseDDoSL7) {
										enabled = &[]bool{rData["enabled"].(bool)}[0]
									}

									rules = append(rules, cloudflare.RulesetRuleActionParametersRules{
										ID:               rData["id"].(string),
										Action:           rData["action"].(string),
										Enabled:          enabled,
										ScoreThreshold:   rData["score_threshold"].(int),
										SensitivityLevel: rData["sensitivity_level"].(string),
									})
								}
							}
						}

						if len(categories) > 0 || len(rules) > 0 {
							overrideConfiguration.Categories = categories
							overrideConfiguration.Rules = rules
						}

						if !reflect.DeepEqual(overrideConfiguration, cloudflare.RulesetRuleActionParametersOverrides{}) {
							rule.ActionParameters.Overrides = &overrideConfiguration
						}

					case "matched_data":
						for i := range pValue.([]interface{}) {
							rule.ActionParameters.MatchedData = &cloudflare.RulesetRuleActionParametersMatchedData{
								PublicKey: pValue.([]interface{})[i].(map[string]interface{})["public_key"].(string),
							}
						}

					case "uri":
						for _, uriValue := range pValue.([]interface{}) {
							if val, ok := uriValue.(map[string]interface{})["path"]; ok && len(val.([]interface{})) > 0 {
								uriPathConfig := val.([]interface{})[0].(map[string]interface{})
								rule.ActionParameters.URI = &cloudflare.RulesetRuleActionParametersURI{
									Path: &cloudflare.RulesetRuleActionParametersURIPath{
										Value:      uriPathConfig["value"].(string),
										Expression: uriPathConfig["expression"].(string),
									},
								}
							}

							if val, ok := uriValue.(map[string]interface{})["query"]; ok && len(val.([]interface{})) > 0 {
								uriQueryConfig := val.([]interface{})[0].(map[string]interface{})
								rule.ActionParameters.URI = &cloudflare.RulesetRuleActionParametersURI{
									Query: &cloudflare.RulesetRuleActionParametersURIQuery{
										Value:      uriQueryConfig["value"].(string),
										Expression: uriQueryConfig["expression"].(string),
									},
								}
							}
						}

					case "headers":
						headers := make(map[string]cloudflare.RulesetRuleActionParametersHTTPHeader)
						for _, headerList := range pValue.([]interface{}) {
							name := headerList.(map[string]interface{})["name"].(string)

							headers[name] = cloudflare.RulesetRuleActionParametersHTTPHeader{
								Value:      headerList.(map[string]interface{})["value"].(string),
								Expression: headerList.(map[string]interface{})["expression"].(string),
								Operation:  headerList.(map[string]interface{})["operation"].(string),
							}
						}

						rule.ActionParameters.Headers = headers

					default:
						log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for action parameters: %s", pKey)
					}
				}
			}
		}

		if len(resourceRule["ratelimit"].([]interface{})) > 0 {
			rule.RateLimit = &cloudflare.RulesetRuleRateLimit{}
			for _, parameter := range resourceRule["ratelimit"].([]interface{}) {
				for pKey, pValue := range parameter.(map[string]interface{}) {
					switch pKey {
					case "characteristics":
						characteristicKeys := make([]string, 0)
						for _, v := range pValue.(*schema.Set).List() {
							characteristicKeys = append(characteristicKeys, v.(string))
						}
						rule.RateLimit.Characteristics = characteristicKeys
					case "period":
						rule.RateLimit.Period = pValue.(int)
					case "requests_per_period":
						rule.RateLimit.RequestsPerPeriod = pValue.(int)
					case "mitigation_timeout":
						rule.RateLimit.MitigationTimeout = pValue.(int)
					case "mitigation_expression":
						rule.RateLimit.MitigationExpression = pValue.(string)

					default:
						log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for ratelimit: %s", pKey)
					}
				}
			}
		}

		rule.Action = resourceRule["action"].(string)
		rule.Enabled = resourceRule["enabled"].(bool)

		if resourceRule["expression"] != nil {
			rule.Expression = resourceRule["expression"].(string)
		}

		if resourceRule["description"] != nil {
			rule.Description = resourceRule["description"].(string)
		}

		rulesetRules = append(rulesetRules, rule)
	}

	return rulesetRules, nil
}
