// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZoneSettingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the zone setting.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"0rtt",
						"advanced_ddos",
						"always_online",
						"always_use_https",
						"automatic_https_rewrites",
						"brotli",
						"browser_cache_ttl",
						"browser_check",
						"cache_level",
						"challenge_ttl",
						"ciphers",
						"cname_flattening",
						"development_mode",
						"early_hints",
						"edge_cache_ttl",
						"email_obfuscation",
						"h2_prioritization",
						"hotlink_protection",
						"http2",
						"http3",
						"image_resizing",
						"ip_geolocation",
						"ipv6",
						"max_upload",
						"min_tls_version",
						"minify",
						"mirage",
						"mobile_redirect",
						"nel",
						"opportunistic_encryption",
						"opportunistic_onion",
						"orange_to_orange",
						"origin_error_page_pass_thru",
						"polish",
						"prefetch_preload",
						"proxy_read_timeout",
						"pseudo_ipv4",
						"replace_insecure_js",
						"response_buffering",
						"rocket_loader",
						"automatic_platform_optimization",
						"security_header",
						"security_level",
						"server_side_exclude",
						"sha1_support",
						"sort_query_string_for_cache",
						"ssl",
						"ssl_recommender",
						"tls_1_2_only",
						"tls_1_3",
						"tls_client_auth",
						"true_client_ip_header",
						"waf",
						"webp",
						"websockets",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"setting_id": schema.StringAttribute{
				Description:   "Setting name",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "ssl-recommender enrollment setting.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"value": schema.StringAttribute{
				Description: "Current value of the zone setting.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("on", "off"),
				},
				Default: stringdefault.StaticString("off"),
			},
		},
	}
}

func (r *ZoneSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneSettingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}