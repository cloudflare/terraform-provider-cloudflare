package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflarePageRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"target": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: suppressEquivalentURLs,
		},

		"actions": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				SchemaVersion: 1,
				Schema: map[string]*schema.Schema{
					// may get api errors trying to set this
					"automatic_https_rewrites": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"bypass_cache_on_cookie": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"cache_by_device_type": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"cache_deception_armor": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"cache_on_cookie": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"mirage": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"explicit_cache_control": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"origin_error_page_pass_thru": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"sort_query_string_for_cache": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"respect_strong_etag": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"response_buffering": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					// may not be used with disable_performance
					"rocket_loader": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"true_client_ip_header": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"browser_check": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"email_obfuscation": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"ip_geolocation": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					// may get api errors trying to set this
					"opportunistic_encryption": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"server_side_exclude": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},

					"waf": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					},
					// end on/off fields

					// unitary fields
					// getting api errors trying to set this
					"always_use_https": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},

					"disable_apps": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},

					// may not be used with rocket loader
					// n.b. ConflictsWith doesn't seem to work on nested schemas
					"disable_performance": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},

					"disable_railgun": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},

					"disable_security": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},

					"disable_zaraz": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},

					"browser_cache_ttl": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"edge_cache_ttl": {
						Type:         schema.TypeInt,
						Optional:     true,
						ValidateFunc: validation.IntAtMost(31536000),
					},

					"cache_level": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"bypass", "basic", "simplified", "aggressive", "cache_everything"}, false),
					},

					"forwarding_url": {
						Type:     schema.TypeList,
						Optional: true,
						MinItems: 1,
						MaxItems: 1,
						Elem: &schema.Resource{
							SchemaVersion: 1,
							Schema: map[string]*schema.Schema{
								"url": {
									Type:     schema.TypeString,
									Required: true,
								},

								"status_code": {
									Type:         schema.TypeInt,
									Required:     true,
									ValidateFunc: validation.IntBetween(301, 302),
								},
							},
						},
					},

					"minify": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							SchemaVersion: 1,
							Schema: map[string]*schema.Schema{
								"js": {
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
								},

								"css": {
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
								},

								"html": {
									Type:         schema.TypeString,
									Required:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
								},
							},
						},
					},

					"host_header_override": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"polish": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"off", "lossless", "lossy"}, false),
					},

					"resolve_override": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"security_level": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"off", "essentially_off", "low", "medium", "high", "under_attack"}, false),
					},

					"ssl": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict", "origin_pull"}, false),
					},

					"cache_key_fields": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						MinItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"cookie": {
									Type:     schema.TypeList,
									Optional: true,
									MinItems: 1,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"check_presence": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"include": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
										},
									},
								},

								"header": {
									Type:     schema.TypeList,
									Optional: true,
									MinItems: 1,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"check_presence": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"exclude": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
											"include": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												},
											},
										},
									},
								},

								"host": {
									Type:     schema.TypeList,
									Required: true,
									MinItems: 1,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"resolved": {
												Type:     schema.TypeBool,
												Optional: true,
												Default:  false,
											},
										},
									},
								},

								"query_string": {
									Type:     schema.TypeList,
									Required: true,
									MinItems: 1,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"exclude": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
													ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
														value := v.(string)
														var diags diag.Diagnostics

														if value == "*" {
															diag := diag.Diagnostic{
																Severity: diag.Error,
																Summary:  "Invalid exclude value",
																Detail:   fmt.Sprintf("full wildcards are not supported for exclude, use ignore=true instead. value: %s", value),
															}
															diags = append(diags, diag)
														}

														return diags
													}},
											},
											"include": {
												Type:     schema.TypeSet,
												Optional: true,
												Computed: true,
												Elem: &schema.Schema{
													Type: schema.TypeString,
													ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
														value := v.(string)
														var diags diag.Diagnostics

														if value == "*" {
															diag := diag.Diagnostic{
																Severity: diag.Error,
																Summary:  "Invalid include value",
																Detail:   fmt.Sprintf("full wildcards are not supported for include, use ignore=false instead. value: %s", value),
															}
															diags = append(diags, diag)
														}

														return diags
													},
												},
											},
											"ignore": {
												Type:     schema.TypeBool,
												Optional: true,
												Computed: true,
											},
										},
									},
								},

								"user": {
									Type:     schema.TypeList,
									Required: true,
									MinItems: 1,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"device_type": {
												Type:     schema.TypeBool,
												Optional: true,
												Computed: true,
											},
											"geo": {
												Type:     schema.TypeBool,
												Optional: true,
												Computed: true,
											},
											"lang": {
												Type:     schema.TypeBool,
												Optional: true,
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
					"cache_ttl_by_status": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"codes": {
									Type:     schema.TypeString,
									Required: true,
								},
								"ttl": {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					},
				},
			},
		},

		"priority": {
			Type:     schema.TypeInt,
			Default:  1,
			Optional: true,
		},

		"status": {
			Type:         schema.TypeString,
			Default:      "active",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"active", "disabled"}, false),
		},
	}
}
