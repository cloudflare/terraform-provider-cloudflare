// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneSettingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zone Settings Read",
				"Zone Settings Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Setting name",
				Computed:    true,
			},
			"setting_id": schema.StringAttribute{
				Description: "Setting name",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"editable": schema.BoolAttribute{
				Description: "For `id: 0rtt` or `id: advanced_ddos` or `id: always_online` or `id: always_use_https` or `id: automatic_https_rewrites` or `id: brotli` or `id: browser_cache_ttl` or `id: browser_check` or `id: cache_level` or `id: challenge_ttl` or `id: china_network_enabled` or `id: content_converter` or `id: ciphers` or `id: cname_flattening` or `id: development_mode` or `id: early_hints` or `id: edge_cache_ttl` or `id: email_obfuscation` or `id: h2_prioritization` or `id: hotlink_protection` or `id: http2` or `id: http3` or `id: image_resizing` or `id: ip_geolocation` or `id: ipv6` or `id: max_upload` or `id: min_tls_version` or `id: mirage` or `id: nel` or `id: opportunistic_encryption` or `id: opportunistic_onion` or `id: orange_to_orange` or `id: origin_error_page_pass_thru` or `id: polish` or `id: prefetch_preload` or `id: pre_render` or `id: privacy_pass` or `id: proxy_read_timeout` or `id: pseudo_ipv4` or `id: redirects_for_ai_training` or `id: replace_insecure_js` or `id: response_buffering` or `id: rocket_loader` or `id: automatic_platform_optimization` or `id: search_for_agents` or `id: security_header` or `id: security_level` or `id: server_side_exclude` or `id: sha1_support` or `id: sort_query_string_for_cache` or `id: ssl` or `id: tls_1_2_only` or `id: tls_1_3` or `id: tls_client_auth` or `id: transformations` or `id: transformations_allowed_origins` or `id: true_client_ip_header` or `id: waf` or `id: webmcp_enabled` or `id: webmcp_mcp_url` or `id: webmcp_packs` or `id: webp` or `id: avif` or `id: websockets`: Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "For `id: ssl_recommender`: ssl-recommender enrollment setting.",
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
			"value": schema.StringAttribute{
				Description: "For `id: 0rtt` or `id: advanced_ddos` or `id: always_online` or `id: always_use_https` or `id: automatic_https_rewrites` or `id: brotli` or `id: browser_check` or `id: cache_level` or `id: challenge_ttl` or `id: china_network_enabled` or `id: content_converter` or `id: ciphers` or `id: cname_flattening` or `id: development_mode` or `id: early_hints` or `id: edge_cache_ttl` or `id: email_obfuscation` or `id: h2_prioritization` or `id: hotlink_protection` or `id: http2` or `id: http3` or `id: image_resizing` or `id: ip_geolocation` or `id: ipv6` or `id: max_upload` or `id: min_tls_version` or `id: mirage` or `id: nel` or `id: opportunistic_encryption` or `id: opportunistic_onion` or `id: orange_to_orange` or `id: origin_error_page_pass_thru` or `id: polish` or `id: prefetch_preload` or `id: pre_render` or `id: privacy_pass` or `id: pseudo_ipv4` or `id: redirects_for_ai_training` or `id: replace_insecure_js` or `id: response_buffering` or `id: rocket_loader` or `id: automatic_platform_optimization` or `id: search_for_agents` or `id: security_header` or `id: security_level` or `id: server_side_exclude` or `id: sha1_support` or `id: sort_query_string_for_cache` or `id: ssl` or `id: tls_1_2_only` or `id: tls_1_3` or `id: tls_client_auth` or `id: transformations` or `id: true_client_ip_header` or `id: waf` or `id: webmcp_enabled` or `id: webp` or `id: avif` or `id: websockets`: Current value of the zone setting.\n\nFor `id: aegis`: Value of the zone setting.\n\nFor `id: browser_cache_ttl`: Value of the zone setting in seconds.\nMinimum values by plan:\n- Free: 1 second\n- Pro: 1 second\n- Business: 1 second\n- Enterprise: 1 second\nSetting a TTL of 0 is equivalent to selecting `Respect Existing Headers` and is allowed for all plans.\n\nFor `id: origin_h2_max_streams`: Value of the Origin H2 Max Streams Setting.\n\nFor `id: origin_max_http_version`: Value of the Origin Max HTTP Version Setting.\n\nFor `id: proxy_read_timeout`: Value of the zone setting.\nNotes: Value must be between 1 and 6000\n\nFor `id: transformations_allowed_origins`: Comma-separated list of allowed origins.\nRefer to the [Image Transformations](https://developers.cloudflare.com/images/transform-images/sources/) and [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started) documentation for more information.\n\nFor `id: webmcp_mcp_url`: Root-relative path (with optional query string) of this zone's MCP\nendpoint, for example \"/api/mcp\". Surfaced to the injected bridge as the\ndata-mcp-url attribute. Absolute and protocol-relative URLs, fragments,\nand characters outside RFC 3986 paths and queries are rejected. An empty\nvalue clears the override so the bridge uses /mcp.\n\nFor `id: webmcp_packs`: Comma-separated list of WebMCP tool-pack names to activate for this\nzone (for example, \"c2pa,mcp-server-client\"), surfaced to the injected\nbridge as the data-packs attribute. An empty value clears the override\nso the bridge uses its built-in default pack set.",
				Computed:    true,
			},
		},
	}
}

func (d *ZoneSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneSettingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
