package sdkv2provider

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRulesetSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
									"phases": {
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
						"exposed_credential_check": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username_expression": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"password_expression": {
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

func resourceCloudflareRulesetStateUpgradeV0ToV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	if rawState["ratelimit"] == nil {
		return rawState, nil
	}

	rawState["ratelimit"].([]map[string]interface{})[0]["counting_expression"] = rawState["ratelimit"].([]map[string]interface{})[0]["mitigation_expression"]
	delete(rawState["ratelimit"].([]map[string]interface{})[0], "mitigation_expression")

	return rawState, nil
}
