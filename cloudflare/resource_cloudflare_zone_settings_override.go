package cloudflare

import (
	"fmt"
	"log"

	"strings"

	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudFlareZoneSettingsOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlareZoneSettingsOverrideCreate,
		Read:   resourceCloudFlareZoneSettingsOverrideRead,
		Update: resourceCloudFlareZoneSettingsOverrideUpdate,
		Delete: resourceCloudFlareZoneSettingsOverrideDelete,

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: resourceCloudFlareZoneSettingsSchema,
				},
			},

			"initial_settings": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: resourceCloudFlareZoneSettingsSchema,
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
		},
	}
}

var resourceCloudFlareZoneSettingsSchema = map[string]*schema.Schema{
	"advanced_ddos": {
		Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional: true,
		Computed: true, //Defaults to on for Business+ plans, off otherwise
	},

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
		ValidateFunc: validateIntInSlice([]int{30, 60, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800,
			43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800,
			16070400, 31536000}),
		// minimum TTL available depends on the plan level of the zone. (Enterprise = 30, Business = 1800, Pro = 1800, Free = 1800)
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

	"challenge_ttl": {
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
		ValidateFunc: validateIntInSlice([]int{300, 900, 1800, 2700, 3600, 7200, 10800, 14400, 28800, 57600,
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

	"polish": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"off", "lossless", "lossy"}, false),
	},

	"webp": {
		Type:     schema.TypeString,
		Computed: true, // setting this causes API errors, conflict with polish
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
		ValidateFunc: validation.StringInSlice([]string{"essentially_off", "low", "medium", "high", "under_attack"}, false),
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
		ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict"}, false), // depends on plan
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

	"tls_1_2_only": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional:     true,
		Computed:     true,
	},

	"tls_1_3": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
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
		Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional: true,
		Computed: true,
	},

	"pseudo_ipv4": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringInSlice([]string{"off", "add_header", "overwrite_header"}, false),
	},

	"always_use_https": {
		Type:     schema.TypeString,
		Computed: true, // setting default 'off' caused HTTP status 400: content "{\"success\":false,\"errors\":[{\"code\":1016,\"message\":\"An unknown error has occurred\"}],
	},

	"sha1_support": {
		Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Optional: true,
		Computed: true,
	},

	"cname_flattening": {
		Type:     schema.TypeString, // enum
		Optional: true,
		Computed: true,
	},

	"max_upload": {
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"edge_cache_ttl": {
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

func resourceCloudFlareZoneSettingsOverrideCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	zoneId, err := client.ZoneIDByName(d.Get("name").(string))
	if err != nil {
		return fmt.Errorf("couldn't find zone %q while trying to import it: %q", d.Get("name").(string), err)
	}
	d.SetId(zoneId)

	log.Printf("[INFO] Creating zone settings resource for zone ID: %s", d.Id())

	// do extra initial read to get initial_settings before updating
	zoneSettings, err := client.ZoneSettings(d.Id())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading initial settings for zone %q", d.Id()))
	}

	log.Printf("[DEBUG] Read CloudFlareZone initial settings: %#v", zoneSettings)

	d.Set("initial_settings", flattenZoneSettings(zoneSettings.Result))
	d.Set("initial_settings_read_at", time.Now().UTC().Format(time.RFC3339Nano))

	log.Printf("[DEBUG] Saved CloudFlareZone initial settings: %#v", d.Get("initial_settings"))

	return resourceCloudFlareZoneSettingsOverrideUpdate(d, meta)
}

func resourceCloudFlareZoneSettingsOverrideUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	if cfg, ok := d.GetOk("settings"); ok && cfg != nil && len(cfg.([]interface{})) > 0 {
		settingsCfg := cfg.([]interface{})[0].(map[string]interface{})
		zoneSettings, err := expandZoneSettings(settingsCfg)
		if err != nil {
			return err
		}
		if browserTTL, ok := d.GetOk("settings.0.browser_cache_ttl"); ok {
			newZoneSetting := cloudflare.ZoneSetting{
				ID:    "browser_cache_ttl",
				Value: browserTTL.(int),
			}
			zoneSettings = append(zoneSettings, newZoneSetting)
		}

		log.Printf("[DEBUG] CloudFlare Zone Settings update configuration: %#v", zoneSettings)

		_, err = client.UpdateZoneSettings(d.Id(), zoneSettings)
		if err != nil {
			return err
		}
	}

	return resourceCloudFlareZoneSettingsOverrideRead(d, meta)
}

func expandZoneSettings(cfg map[string]interface{}) ([]cloudflare.ZoneSetting, error) {
	zoneSettings := make([]cloudflare.ZoneSetting, 0)

	// can't distinguish between empty and unset values in here
	for k, v := range cfg {
		var zoneSettingValue interface{}

		if k == "webp" || k == "always_use_https" {
			// errors in the api when setting these values, ignore for now
			continue
		} else if k == "browser_cache_ttl" {
			// need to distinguish explicit 0 from not set
			continue
		} else if listValue, ok := v.([]interface{}); ok && (k == "minify" || k == "mobile_redirect") {
			if len(listValue) > 0 && listValue != nil {
				zoneSettingValue = listValue[0].(map[string]interface{})
			} else {
				continue
			}
		} else if listValue, ok := v.([]interface{}); ok && k == "security_header" {
			if len(listValue) > 0 && listValue != nil {
				val := map[string]interface{}{
					"strict_transport_security": listValue[0].(map[string]interface{}),
				}
				zoneSettingValue = val
			} else {
				continue
			}
		} else if strValue, ok := v.(string); ok {
			// default mapping for non-empty string fields
			//empty string means we didnt set a value
			if strValue != "" {
				zoneSettingValue = strValue
			} else {
				continue
			}
		} else if intValue, ok := v.(int); ok {
			// default mapping for non-empty int fields
			if intValue != 0 {
				zoneSettingValue = intValue // passthrough
			} else {
				continue
			}
		} else {
			return nil, fmt.Errorf("unknown zone setting specified %q = %#v", k, v)
		}
		newZoneSetting := cloudflare.ZoneSetting{
			ID:    k,
			Value: zoneSettingValue,
		}
		zoneSettings = append(zoneSettings, newZoneSetting)
	}
	return zoneSettings, nil
}

func resourceCloudFlareZoneSettingsOverrideRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	zone, err := client.ZoneDetails(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Zone %q not found", d.Id())
			d.SetId("")
			return nil
		} else {
			return errors.Wrap(err, fmt.Sprintf("Error reading zone %q", d.Id()))
		}
	}

	// not all settings are visible to all users, so this might be a subset
	// assume (for now) that user can see/do everything
	zoneSettings, err := client.ZoneSettings(d.Id())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading settings for zone %q", d.Id()))
	}

	log.Printf("[DEBUG] Read CloudFlareZone Settings: %#v", zoneSettings)

	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("settings", flattenZoneSettings(zoneSettings.Result))
	d.Set("readonly_settings", flattenReadOnlyZoneSettings(zoneSettings.Result))

	return nil
}

func flattenZoneSettings(settings []cloudflare.ZoneSetting) []map[string]interface{} {
	cfg := map[string]interface{}{}
	for _, s := range settings {
		if !settingInSchema(s.ID) {
			log.Printf("[WARN] Value not in schema returned from API zone settings (is it new?) - %q : %#v", s.ID, s.Value)
		} else if s.ID == "minify" || s.ID == "mobile_redirect" {
			cfg[s.ID] = []interface{}{s.Value.(map[string]interface{})}
		} else if s.ID == "security_header" {
			cfg[s.ID] = []interface{}{s.Value.(map[string]interface{})["strict_transport_security"]}
		} else if strValue, ok := s.Value.(string); ok {
			log.Printf("[DEBUG] Found string zone setting %q: %q", s.ID, strValue)
			cfg[s.ID] = strValue
		} else if floatValue, ok := s.Value.(float64); ok {
			log.Printf("[DEBUG] Found int zone setting %q: %d", s.ID, int(floatValue))
			cfg[s.ID] = int(floatValue)
		} else {
			log.Printf("[WARN] Unexpected value type found in API zone settings - %q : %#v", s.ID, s.Value)
		}
	}

	log.Printf("[DEBUG] Flattened CloudFlare Zone Settings: %#v", cfg)

	return []map[string]interface{}{cfg}
}

func settingInSchema(val string) bool {
	for k, _ := range resourceCloudFlareZoneSettingsSchema {
		if val == k {
			return true
		}
	}
	return false
}

func flattenReadOnlyZoneSettings(settings []cloudflare.ZoneSetting) []string {
	ids := make([]string, 0)
	for _, zs := range settings {
		if !zs.Editable {
			ids = append(ids, zs.ID)
		}
	}
	log.Printf("[DEBUG] Flattened CloudFlare Read Only Zone Settings: %#v", ids)

	return ids
}

func resourceCloudFlareZoneSettingsOverrideDelete(d *schema.ResourceData, meta interface{}) error {
	// we cannot delete settings independently of the zone, which is why the resources have to be combined

	d.Set("settings", d.Get("initial_settings"))

	return resourceCloudFlareZoneSettingsOverrideUpdate(d, meta)
}
