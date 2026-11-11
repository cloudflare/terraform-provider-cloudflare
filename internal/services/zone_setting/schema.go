// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZoneSettingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zone Settings Read",
				"Zone Settings Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Setting name",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"setting_id": schema.StringAttribute{
				Description:   "Setting name",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.DynamicAttribute{
				Description: "Current value of the zone setting.",
				Required:    true,
				CustomType:  customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{
					customfield.NormalizeDynamicPlanModifier(),
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "ssl-recommender enrollment setting.",
				Computed:    true,
				Optional:    true,
			},
			"editable": schema.BoolAttribute{
				Description: "For `id: 0rtt` or `id: advanced_ddos` or `id: always_online` or `id: always_use_https` or `id: automatic_https_rewrites` or `id: brotli` or `id: browser_cache_ttl` or `id: browser_check` or `id: cache_level` or `id: challenge_ttl` or `id: china_network_enabled` or `id: content_converter` or `id: ciphers` or `id: cname_flattening` or `id: development_mode` or `id: early_hints` or `id: edge_cache_ttl` or `id: email_obfuscation` or `id: h2_prioritization` or `id: hotlink_protection` or `id: http2` or `id: http3` or `id: image_resizing` or `id: ip_geolocation` or `id: ipv6` or `id: max_upload` or `id: min_tls_version` or `id: mirage` or `id: nel` or `id: opportunistic_encryption` or `id: opportunistic_onion` or `id: orange_to_orange` or `id: origin_error_page_pass_thru` or `id: polish` or `id: prefetch_preload` or `id: pre_render` or `id: privacy_pass` or `id: proxy_read_timeout` or `id: pseudo_ipv4` or `id: redirects_for_ai_training` or `id: replace_insecure_js` or `id: response_buffering` or `id: rocket_loader` or `id: automatic_platform_optimization` or `id: search_for_agents` or `id: security_header` or `id: security_level` or `id: server_side_exclude` or `id: sha1_support` or `id: sort_query_string_for_cache` or `id: ssl` or `id: tls_1_2_only` or `id: tls_1_3` or `id: tls_client_auth` or `id: transformations` or `id: transformations_allowed_origins` or `id: true_client_ip_header` or `id: waf` or `id: webmcp_enabled` or `id: webmcp_mcp_url` or `id: webmcp_packs` or `id: webp` or `id: avif` or `id: websockets`: Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "For `id: 0rtt` or `id: advanced_ddos` or `id: always_online` or `id: always_use_https` or `id: automatic_https_rewrites` or `id: brotli` or `id: browser_cache_ttl` or `id: browser_check` or `id: cache_level` or `id: challenge_ttl` or `id: china_network_enabled` or `id: content_converter` or `id: ciphers` or `id: cname_flattening` or `id: development_mode` or `id: early_hints` or `id: edge_cache_ttl` or `id: email_obfuscation` or `id: h2_prioritization` or `id: hotlink_protection` or `id: http2` or `id: http3` or `id: image_resizing` or `id: ip_geolocation` or `id: ipv6` or `id: max_upload` or `id: min_tls_version` or `id: mirage` or `id: nel` or `id: opportunistic_encryption` or `id: opportunistic_onion` or `id: orange_to_orange` or `id: origin_error_page_pass_thru` or `id: polish` or `id: prefetch_preload` or `id: pre_render` or `id: privacy_pass` or `id: proxy_read_timeout` or `id: pseudo_ipv4` or `id: redirects_for_ai_training` or `id: replace_insecure_js` or `id: response_buffering` or `id: rocket_loader` or `id: automatic_platform_optimization` or `id: search_for_agents` or `id: security_header` or `id: security_level` or `id: server_side_exclude` or `id: sha1_support` or `id: sort_query_string_for_cache` or `id: ssl` or `id: tls_1_2_only` or `id: tls_1_3` or `id: tls_client_auth` or `id: transformations` or `id: transformations_allowed_origins` or `id: true_client_ip_header` or `id: waf` or `id: webmcp_enabled` or `id: webmcp_mcp_url` or `id: webmcp_packs` or `id: webp` or `id: avif` or `id: websockets`: last time this setting was modified.\n\nFor `id: aegis` or `id: origin_h2_max_streams` or `id: origin_max_http_version`: Last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"time_remaining": schema.Float64Attribute{
				Description: "For `id: development_mode`: Value of the zone setting.\nNotes: The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is false.",
				Computed:    true,
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
