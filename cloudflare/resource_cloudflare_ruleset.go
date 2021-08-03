package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"path": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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
									"increment": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ruleset": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"current"}, false),
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

	rules, err := buildRulesetRulesFromResource(d.Get("rules"))
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

	for i, rule := range rules {
		if rule.Action == string(cloudflare.RulesetRuleActionExecute) {
			rulesetEntryPoint := cloudflare.Ruleset{
				Description: d.Get("description").(string),
				Rules:       []cloudflare.RulesetRule{rules[i]},
			}

			if accountID != "" {
				_, err = client.UpdateAccountRulesetPhase(context.Background(), accountID, rs.Phase, rulesetEntryPoint)
			} else {
				_, err = client.UpdateZoneRulesetPhase(context.Background(), zoneID, rs.Phase, rulesetEntryPoint)
			}

			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error updating ruleset phase entrypoint %s", d.Get("name").(string)))
			}
		}
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
	d.Set("rules", buildStateFromRulesetRules(ruleset.Rules))

	return nil
}

func resourceCloudflareRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	zoneID := d.Get("zone_id").(string)

	rules, err := buildRulesetRulesFromResource(d.Get("rules"))
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

// receives the current rules and returns an interface for the state file
func buildStateFromRulesetRules(r []cloudflare.RulesetRule) interface{} {
	var ruleset []interface{}
	var rulesetRule map[string]interface{}

	for _, rule := range r {
		rulesetRule = make(map[string]interface{})

		rulesetRule["expression"] = rule.Expression
		rulesetRule["action"] = rule.Action
		if rule.Description != "" {
			rulesetRule["description"] = rule.Description
		}

		if rule.Enabled {
			rulesetRule["enabled"] = "true"
		} else {
			rulesetRule["enabled"] = "false"
		}

		ruleset = append(ruleset, rulesetRule)
	}

	return ruleset
}

// receives the resource config and builds a ruleset rule array
func buildRulesetRulesFromResource(r interface{}) ([]cloudflare.RulesetRule, error) {
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

		rule.ActionParameters = &cloudflare.RulesetRuleActionParameters{}
		for _, parameter := range resourceRule["action_parameters"].([]interface{}) {
			for pKey, pValue := range parameter.(map[string]interface{}) {
				switch pKey {
				case "id":
					rule.ActionParameters.ID = pValue.(string)
				case "overrides":
					categories := []cloudflare.RulesetRuleActionParametersCategories{}
					rules := []cloudflare.RulesetRuleActionParametersRules{}

					for _, overrideParamValue := range pValue.([]interface{}) {
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
								rules = append(rules, cloudflare.RulesetRuleActionParametersRules{
									ID:             rData["id"].(string),
									Action:         rData["action"].(string),
									Enabled:        rData["enabled"].(bool),
									ScoreThreshold: rData["score_threshold"].(int),
								})
							}
						}
					}

					if len(categories) > 0 || len(rules) > 0 {
						rule.ActionParameters.Overrides = &cloudflare.RulesetRuleActionParametersOverrides{
							Categories: categories,
							Rules:      rules,
						}
					}

				default:
					log.Printf("[DEBUG] unknown key encountered in buildRulesetRulesFromResource for action parameters: %s", pKey)
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
