// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
  "context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			"priority": schema.Int64Attribute{
				Description: "The priority of the rule, used to define which Page Rule is processed\nover another. A higher number indicates a higher priority. For example,\nif you have a catch-all Page Rule (rule A: `/images/*`) but want a more\nspecific Page Rule to take precedence (rule B: `/images/special/*`),\nspecify a higher priority for rule B so it overrides rule A.",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(1),
			},
			"status": schema.StringAttribute{
				Description: "The status of the Page Rule.\nAvailable values: \"active\", \"disabled\".",
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
			"target": schema.StringAttribute{
				Required: true,
			},
			"actions": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"always_use_https": schema.BoolAttribute{
						Optional: true,
					},
					"automatic_https_rewrites": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},
					"browser_cache_ttl": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(31536000),
						},
					},
					"browser_check": schema.StringAttribute{
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
					"cache_key_fields": schema.SingleNestedAttribute{
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[PageRuleActionsCacheKeyFieldsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"cookie": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"check_presence": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
									"include": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
								},
							},
							"header": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"check_presence": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
									"include": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
									"exclude": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
								},
							},
							"host": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"resolved": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
							"query_string": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"include": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
									"exclude": schema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
										ElementType: types.StringType,
									},
								},
							},
							"user": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"device_type": schema.BoolAttribute{
										Optional: true,
									},
									"geo": schema.BoolAttribute{
										Optional: true,
									},
									"lang": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
					"cache_level": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("bypass", "basic", "simplified", "aggressive", "cache_everything"),
						},
					},
					"cache_on_cookie": schema.StringAttribute{
						Optional: true,
					},
					"cache_ttl_by_status": schema.DynamicAttribute{
						Optional: true,
					},
					"disable_apps": schema.BoolAttribute{
						Optional: true,
					},
					"disable_performance": schema.BoolAttribute{
						Optional: true,
					},
					"disable_security": schema.BoolAttribute{
						Optional: true,
					},
					"disable_zaraz": schema.BoolAttribute{
						Optional: true,
					},
					"edge_cache_ttl": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(2419200),
						},
					},
					"email_obfuscation": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},
					"forwarding_url": schema.SingleNestedAttribute{
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[PageRuleActionsForwardingURLModel](ctx),
						Attributes: map[string]schema.Attribute{
							"url": schema.StringAttribute{
								Required: true,
							},
							"status_code": schema.Int64Attribute{
								Required: true,
							},
						},
					},
					"host_header_override": schema.StringAttribute{
						Optional: true,
					},
					"ip_geolocation": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},
					"mirage": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},
					"opportunistic_encryption": schema.StringAttribute{
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
					"polish": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "lossless", "lossy"),
						},
					},
					"resolve_override": schema.StringAttribute{
						Optional: true,
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
					"rocket_loader": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},
					"security_level": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "essentially_off", "low", "medium", "high", "under_attack"),
						},
					},
					"sort_query_string_for_cache": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("on", "off"),
						},
					},
					"ssl": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("off", "flexible", "full", "strict", "origin_pull"),
						},
					},
					"true_client_ip_header": schema.StringAttribute{
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
				},
			},
		},
	}
}

func (r *PageRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *PageRuleResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
