// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PageRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"target": schema.StringAttribute{
				Required: true,
			},
			"actions": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"automatic_https_rewrites": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"bypass_cache_on_cookie": schema.StringAttribute{
						Optional: true,
					},

					"cache_by_device_type": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"cache_deception_armor": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"cache_on_cookie": schema.StringAttribute{
						Optional: true,
					},

					"mirage": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"explicit_cache_control": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"origin_error_page_pass_thru": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"sort_query_string_for_cache": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"respect_strong_etag": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"response_buffering": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					// may not be used with disable_performance
					"rocket_loader": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"true_client_ip_header": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"browser_check": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"email_obfuscation": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"ip_geolocation": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					// may get api errors trying to set this
					"opportunistic_encryption": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"server_side_exclude": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"waf": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},

					"always_use_https": schema.BoolAttribute{
						Default:  booldefault.StaticBool(false),
						Optional: true,
					},

					"disable_apps": schema.BoolAttribute{
						Default:  booldefault.StaticBool(false),
						Optional: true,
					},

					"disable_performance": schema.BoolAttribute{
						Default:  booldefault.StaticBool(false),
						Optional: true,
					},

					"disable_railgun": schema.BoolAttribute{
						Default:  booldefault.StaticBool(false),
						Optional: true,
					},

					"disable_security": schema.BoolAttribute{
						Default:  booldefault.StaticBool(false),
						Optional: true,
					},

					"disable_zaraz": schema.BoolAttribute{
						Default:  booldefault.StaticBool(false),
						Optional: true,
					},

					"browser_cache_ttl": schema.StringAttribute{
						Optional: true,
					},

					"edge_cache_ttl": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtMost(31536000),
						},
					},

					"cache_level": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("bypass", "basic", "simplified", "aggressive", "cache_everything"),
						},
					},

					"forwarding_url": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"url": schema.StringAttribute{
								Required: true,
							},

							"status_code": schema.Int64Attribute{
								Required: true,
								Validators: []validator.Int64{
									int64validator.Between(301, 302),
								},
							},
						},
					},

					"minify": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"js": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("on", "off"),
								},
							},

							"css": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("on", "off"),
								},
							},

							"html": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("on", "off"),
								},
							},
						},
					},

					"host_header_override": schema.StringAttribute{
						Optional: true,
					},

					"polish": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "lossless", "lossy"),
						},
					},

					"resolve_override": schema.StringAttribute{
						Optional: true,
					},

					"security_level": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "essentially_off", "low", "medium", "high", "under_attack"),
						},
					},

					"ssl": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "flexible", "full", "strict", "origin_pull"),
						},
					},

					// "cache_key_fields": {
					// 	Type:     schema.SingleNestedAttribute,
					// 	Optional: true,
					// 	Attributes: map[string]schema.Attribute{
					// 			"cookie": {
					// 				Type:     schema.SingleNestedAttribute,
					// 				Optional: true,
					// 				Attributes: map[string]schema.Attribute{
					// 						"check_presence": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							ElementType:   types.StringType,
					// 						},
					// 						"include": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							ElementType:   types.StringType,
					// 						},
					// 					},
					// 				},
					// 			},

					// 			"header": {
					// 				Type:     schema.SingleNestedAttribute,
					// 				Optional: true,
					// 				Attributes: map[string]schema.Attribute{
					// 						"check_presence": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							ElementType:   types.StringType,
					// 						},
					// 						"exclude": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							ElementType:   types.StringType,
					// 						},
					// 						"include": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							ElementType:   types.StringType,
					// 						},
					// 					},
					// 				},
					// 			},

					// 			"host": {
					// 				Type:     schema.SingleNestedAttribute,
					// 				Required: true,
					// 				Attributes: map[string]schema.Attribute{
					// 						"resolved": {
					// 							Type:     schema.BoolAttribute,
					// 							Optional: true,
					// 							Default:  booldefault.StaticBool(false),
					// 						},
					// 					},
					// 				},
					// 			},

					// 			"query_string": {
					// 				Type:     schema.SingleNestedAttribute,
					// 				Required: true,
					// 				Attributes: map[string]schema.Attribute{
					// 						"exclude": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							Elem: &schema.Schema{
					// 								Type: schema.StringAttribute,
					// 								ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					// 									value := v.(string)
					// 									var diags diag.Diagnostics

					// 									if value == "*" {
					// 										diag := diag.Diagnostic{
					// 											Severity: diag.Error,
					// 											Summary:  "Invalid exclude value",
					// 											Detail:   fmt.Sprintf("full wildcards are not supported for exclude, use ignore=true instead. value: %s", value),
					// 										}
					// 										diags = append(diags, diag)
					// 									}

					// 									return diags
					// 								}},
					// 						},
					// 						"include": {
					// 							Type:     schema.SetAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 							Elem: &schema.Schema{
					// 								Type: schema.StringAttribute,
					// 								ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					// 									value := v.(string)
					// 									var diags diag.Diagnostics

					// 									if value == "*" {
					// 										diag := diag.Diagnostic{
					// 											Severity: diag.Error,
					// 											Summary:  "Invalid include value",
					// 											Detail:   fmt.Sprintf("full wildcards are not supported for include, use ignore=false instead. value: %s", value),
					// 										}
					// 										diags = append(diags, diag)
					// 									}

					// 									return diags
					// 								},
					// 							},
					// 						},
					// 						"ignore": {
					// 							Type:     schema.BoolAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 						},
					// 					},
					// 				},
					// 			},

					// 			"user": {
					// 				Type:     schema.SingleNestedAttribute,
					// 				Required: true,
					// 				Attributes: map[string]schema.Attribute{
					// 						"device_type": {
					// 							Type:     schema.BoolAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 						},
					// 						"geo": {
					// 							Type:     schema.BoolAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 						},
					// 						"lang": {
					// 							Type:     schema.BoolAttribute,
					// 							Optional: true,
					// 							Computed: true,
					// 						},
					// 					},
					// 				},
					// 			},
					// 		},
					// 	},
					// 	"cache_ttl_by_status": {
					// 	Type:     schema.SetAttribute,
					// 	Optional: true,
					// 	Attributes: map[string]schema.Attribute{
					// 			"codes": {
					// 				Type:     schema.StringAttribute,
					// 				Required: true,
					// 			},
					// 			"ttl": {
					// 				Type:     schema.Int64Attribute,
					// 				Required: true,
					// 			},
					// 	},
					// },
				},
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the rule, used to define which Page Rule is processed\nover another. A higher number indicates a higher priority. For example,\nif you have a catch-all Page Rule (rule A: `/images/*`) but want a more\nspecific Page Rule to take precedence (rule B: `/images/special/*`),\nspecify a higher priority for rule B so it overrides rule A.\n",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(1),
			},
			"status": schema.StringAttribute{
				Description: "The status of the Page Rule.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "disabled"),
				},
				Default: stringdefault.StaticString("disabled"),
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the Page Rule was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *PageRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PageRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
