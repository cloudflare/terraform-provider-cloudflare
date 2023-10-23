package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareZoneSettingsOverrideSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"settings": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: resourceCloudflareZoneSettingsSchema,
			},
		},

		"initial_settings": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: resourceCloudflareZoneSettingsSchema,
			},
		},

		"initial_settings_read_at": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"readonly_settings": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"zone_status": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"zone_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

var resourceCloudflareZoneSettingsSchema = map[string]*schema.Schema{
	"always_online": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"brotli": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"browser_cache_ttl": {
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
		ValidateFunc: validation.IntInSlice([]int{0, 30, 60, 120, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800,
			43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800,
			16070400, 31536000}),
		// minimum TTL available depends on the plan level of the zone.
		// - Respect existing headers = 0
		// - Enterprise = 30
		// - Business, Pro, Free = 1800
	},

	"browser_check": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"cache_level": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"aggressive", "basic", "simplified"}, false),
	},

	"ciphers": {
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},

	"challenge_ttl": {
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
		ValidateFunc: validation.IntInSlice([]int{300, 900, 1800, 2700, 3600, 7200, 10800, 14400, 28800, 57600,
			86400, 604800, 2592000, 31536000}),
	},

	"development_mode": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"origin_error_page_pass_thru": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"sort_query_string_for_cache": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"email_obfuscation": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"hotlink_protection": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"ip_geolocation": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"ipv6": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"websockets": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"minify": {
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"css": {
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					Required:     true,
				},

				"html": {
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					Required:     true,
				},

				"js": {
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					Required:     true,
				},
			},
		},
	},

	"mobile_redirect": {
		Type:     schema.TypeList, // on/off
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// which parameters are mandatory is not specified
				"mobile_subdomain": {
					Type:     schema.TypeString,
					Required: true,
				},

				"strip_uri": {
					Type:     schema.TypeBool,
					Required: true,
				},

				"status": {
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
					Required:     true,
				},
			},
		},
	},

	"mirage": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"opportunistic_encryption": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"opportunistic_onion": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"origin_max_http_version": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),
	},

	"polish": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"off", "lossless", "lossy"}, false),
	},

	"webp": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
	},

	"prefetch_preload": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"privacy_pass": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"response_buffering": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"rocket_loader": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"on", "off", "manual"}, false),
	},

	"security_header": {
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"enabled": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},

				"preload": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},

				"max_age": {
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},

				"include_subdomains": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},

				"nosniff": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},
			},
		},
	},

	"security_level": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"off", "essentially_off", "low", "medium", "high", "under_attack"}, false),
	},

	"server_side_exclude": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"ssl": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict", "origin_pull"}, false), // depends on plan
	},

	"universal_ssl": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
	},

	"tls_client_auth": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"true_client_ip_header": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"waf": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"min_tls_version": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"1.0", "1.1", "1.2", "1.3"}, false),
		Optional:     true,
		Computed:     true,
	},

	"tls_1_2_only": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
		Deprecated:   "tls_1_2_only has been deprecated in favour of using `min_tls_version = \"1.2\"` instead.",
	},

	"tls_1_3": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off", "zrt"}, false),
		Optional:     true,
		Computed:     true, // default depends on plan level
	},

	"automatic_https_rewrites": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"http2": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"http3": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"pseudo_ipv4": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"off", "add_header", "overwrite_header"}, false),
	},

	"always_use_https": {
		// may cause an error: HTTP status 400: content "{\"success\":false,\"errors\":[{\"code\":1016,\"message\":\"An unknown error has occurred\"}],\"messages\":[],\"result\":null}"
		// but it still gets set at the API
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"cname_flattening": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"flatten_at_root", "flatten_all", "flatten_none"}, false),
		Optional:     true,
		Computed:     true,
	},

	"max_upload": {
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"h2_prioritization": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off", "custom"}, false),
		Optional:     true,
		Computed:     true,
	},

	"image_resizing": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off", "open"}, false),
		Optional:     true,
		Computed:     true,
	},

	"zero_rtt": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"orange_to_orange": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"filter_logs_to_cloudflare": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"log_to_cloudflare": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"visitor_ip": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"proxy_read_timeout": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"binary_ast": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"early_hints": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"fonts": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},
}
