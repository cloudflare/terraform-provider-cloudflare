// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneSettingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"setting_id": schema.StringAttribute{
				Description: "Setting name",
				Required:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"editable": schema.BoolAttribute{
				Description: "Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "ssl-recommender enrollment setting.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "ID of the zone setting.\navailable values: \"0rtt\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"0rtt",
						"advanced_ddos",
						"aegis",
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
						"mirage",
						"nel",
						"opportunistic_encryption",
						"opportunistic_onion",
						"orange_to_orange",
						"origin_error_page_pass_thru",
						"origin_h2_max_streams",
						"origin_max_http_version",
						"polish",
						"prefetch_preload",
						"privacy_pass",
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
			},
			"modified_on": schema.StringAttribute{
				Description: "last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"time_remaining": schema.Float64Attribute{
				Description: "Value of the zone setting.\nNotes: The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is false.",
				Computed:    true,
			},
			"value": schema.StringAttribute{
				Description: "Current value of the zone setting.\navailable values: \"on\", \"off\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("on", "off"),
				},
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
