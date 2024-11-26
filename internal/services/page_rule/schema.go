// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			"actions": schema.ListNestedAttribute{
				Description: "The set of actions to perform if the targets of this rule match the\nrequest. Actions can redirect to another URL or override settings, but\nnot both.\n",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "If enabled, any `http://`` URL is converted to `https://` through a\n301 redirect.\n",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"always_use_https",
									"automatic_https_rewrites",
									"browser_cache_ttl",
									"browser_check",
									"bypass_cache_on_cookie",
									"cache_by_device_type",
									"cache_deception_armor",
									"cache_key",
									"cache_level",
									"cache_on_cookie",
									"disable_apps",
									"disable_performance",
									"disable_security",
									"disable_zaraz",
									"edge_cache_ttl",
									"email_obfuscation",
									"explicit_cache_control",
									"forwarding_url",
									"host_header_override",
									"ip_geolocation",
									"mirage",
									"opportunistic_encryption",
									"origin_error_page_pass_thru",
									"polish",
									"resolve_override",
									"respect_strong_etag",
									"response_buffering",
									"rocket_loader",
									"security_level",
									"sort_query_string_for_cache",
									"ssl",
									"true_client_ip_header",
									"waf",
								),
							},
						},
						"value": schema.StringAttribute{
							Description: "The status of Automatic HTTPS Rewrites.\n",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("on", "off"),
							},
						},
					},
				},
			},
			"targets": schema.ListNestedAttribute{
				Description: "The rule targets to evaluate on each request.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"constraint": schema.SingleNestedAttribute{
							Description: "String constraint.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"operator": schema.StringAttribute{
									Description: "The matches operator can use asterisks and pipes as wildcard and 'or' operators.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"matches",
											"contains",
											"equals",
											"not_equal",
											"not_contain",
										),
									},
									Default: stringdefault.StaticString("contains"),
								},
								"value": schema.StringAttribute{
									Description: "The URL pattern to match against the current request. The pattern may contain up to four asterisks ('*') as placeholders.",
									Required:    true,
								},
							},
						},
						"target": schema.StringAttribute{
							Description: "A target based on the URL of the request.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("url"),
							},
						},
					},
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
