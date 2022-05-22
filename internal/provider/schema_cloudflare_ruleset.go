package provider

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRulesetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
						Optional: true,
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
								"response": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"status_code": {
												Type:     schema.TypeInt,
												Optional: true,
											},
											"content_type": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"content": {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								"host_header": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"origin": {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"host": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"port": {
												Type:     schema.TypeInt,
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
								"counting_expression": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"requests_to_origin": {
									Type:     schema.TypeBool,
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
					"logging": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	}
}
