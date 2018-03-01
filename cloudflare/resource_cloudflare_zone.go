package cloudflare

import (
	"fmt"
	"log"

	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudFlareZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlareZoneCreate,
		Read:   resourceCloudFlareZoneRead,
		Update: resourceCloudFlareZoneUpdate,
		Delete: resourceCloudFlareZoneDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudFlareZoneImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"with_existing_records": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},

			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//TODO remove defaults looking for the one that causes: * cloudflare_zone.test: error from makeRequest: HTTP status 400: content "{\"success\":false,\"errors\":[{\"code\":1016,\"message\":\"An unknown error has occurred\"}],\"messages\":[],\"result\":null}"

						"advanced_ddos": {
							Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional: true,
							Computed: true, //Defaults to on for Business+ plans, off otherwise
						},

						"always_online": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"brotli": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"browser_cache_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							/*valid values: 30, 60, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000
							notes: The minimum TTL available depends on the plan level of the zone. (Enterprise = 30, Business = 1800, Pro = 1800, Free = 1800)*/
						},

						"browser_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"cache_level": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "aggressive",
							ValidateFunc: validation.StringInSlice([]string{"aggressive", "basic", "simplified"}, false),
						},

						"challenge_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1800,
							/* valid values: 300, 900, 1800, 2700, 3600, 7200, 10800, 14400, 28800, 57600, 86400, 604800, 2592000, 31536000 */
						},

						"development_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"origin_error_page_pass_thru": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"sort_query_string_for_cache": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"email_obfuscation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"hotlink_protection": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"ip_geolocation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"ipv6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"websockets": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
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
										Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
										Optional: true,
										Default:  false,
									},

									"html": {
										Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
										Optional: true,
										Default:  false,
									},

									"js": {
										Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
										Optional: true,
										Default:  false,
									},
								},
							},
						},

						"mobile_redirect": { // TODO: received type is : map[string]interface {}{"mobile_subdomain":interface {}(nil), "strip_uri":false, "status":"off"}
							Type:     schema.TypeString, // on/off
							Optional: true,
							Default:  "",
						},

						"mirage": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"opportunistic_encryption": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"polish": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"off", "lossless", "lossy"}, false),
						},

						"webp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "",
						},

						"prefetch_preload": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"privacy_pass": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"response_buffering": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"rocket_loader": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"on", "off", "manual"}, false),
						},

						"security_header": {
							// TODO theres a nested "struct_transport_security" layer - why?
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// TODO not specified if these fields are required (test this)
									"enabled": {
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
							Default:      "medium",
							ValidateFunc: validation.StringInSlice([]string{"essentially_off", "low", "medium", "high", "under_attack"}, false),
						},

						"server_side_exclude": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"ssl": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict"}, false), // depends on plan
						},

						"tls_client_auth": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "on",
						},

						"true_client_ip_header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"waf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
						},

						"tls_1_2_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional:     true,
							Default:      "off",
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
							Default:      "off",
						},

						"http2": {
							Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional: true,
							Default:  "off",
						},

						"pseudo_ipv4": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"off", "add_header", "overwrite_header"}, false),
						},

						"always_use_https": {
							Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional: true,
							Default:  "off",
						},

						"sha1_support": {
							Type: schema.TypeString, ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
							Optional: true,
							Default:  "off",
						},

						"cname_flattening": {
							Type:     schema.TypeString, // enum
							Optional: true,
							Computed: true,
						},

						"max_upload": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  100,
						},

						"edge_cache_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  7200,
						},
					},
				},
			},

			"editable_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudFlareZoneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	//TODO change to zone settings
	newZone, err := client.CreateZone(d.Get("name").(string), d.Get("with_existing_records").(bool), cloudflare.Organization{})
	if err != nil {
		log.Printf("[WARN] Error creating zone: %q", err.Error())
		zoneId, err := client.ZoneIDByName(d.Get("name").(string))
		if err != nil {
			return fmt.Errorf("couldn't find zone %q while trying to import it: %q", d.Get("name").(string), err)
		}
		d.SetId(zoneId)
	} else {
		d.SetId(newZone.ID)
	}

	log.Printf("[INFO] CloudFlare New Zone ID: %s", d.Id())

	return resourceCloudFlareZoneUpdate(d, meta)
}

func resourceCloudFlareZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	if cfg, ok := d.GetOk("settings"); ok && len(cfg.([]interface{})) > 0 {
		settingsCfg := cfg.([]interface{})[0].(map[string]interface{})
		zoneSettings, err := expandZoneSettings(settingsCfg)
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] CloudFlare Zone Settings update configuration: %#v", zoneSettings)

		_, err = client.UpdateZoneSettings(d.Id(), zoneSettings)
		if err != nil {
			return err
		}
	}

	return resourceCloudFlareZoneRead(d, meta)
}

func expandZoneSettings(cfg map[string]interface{}) ([]cloudflare.ZoneSetting, error) {
	zoneSettings := make([]cloudflare.ZoneSetting, 0)

	for k, v := range cfg {
		var zoneSettingValue interface{}

		if strValue, ok := v.(string); ok {
			//empty string means we didnt set a value
			if strValue != "" {
				zoneSettingValue = strValue
			} else {
				continue
			}
		} else if intValue, ok := v.(int); ok {
			zoneSettingValue = intValue // passthrough
		} else if listValue, ok := v.([]interface{}); ok && k == "minify" {
			if len(listValue) > 0 {
				zoneSettingValue = listValue[0].(map[string]interface{})
			} else {
				continue
			}
		} else if listValue, ok := v.([]interface{}); ok && k == "security_header" {
			if len(listValue) > 0 {
				zoneSettingValue = listValue[0].(map[string]interface{})
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

func resourceCloudFlareZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	zone, err := client.ZoneDetails(d.Id())

	// not all settings are visible to all users, so this might be a subset
	// assume (for now) that user can see/do everything
	zoneSettings, err := client.ZoneSettings(d.Id())
	if err != nil {
		// TODO on 404, set id blank
		return err
	}

	log.Printf("[DEBUG] Read CloudFlareZone Settings: %#v", zoneSettings)

	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("name_servers", zone.NameServers)
	d.Set("created_on", zone.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", zone.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("settings", flattenZoneSettings(zoneSettings.Result))
	d.Set("editable_settings", flattenEditableZoneSettings(zoneSettings.Result))

	return nil
}

func flattenZoneSettings(settings []cloudflare.ZoneSetting) []map[string]interface{} {
	cfg := map[string]interface{}{}
	for _, s := range settings {
		if s.ID == "minify" {
			//TODO
		} else if s.ID == "security_header" {
			//TODO
		} else if s.ID == "mobile_redirect" {
			//TODO
		} else if strValue, ok := s.Value.(string); ok {
			log.Printf("[DEBUG] Found string zone setting %q: %q", s.ID, strValue)
			cfg[s.ID] = strValue
		} else if floatValue, ok := s.Value.(float64); ok {
			log.Printf("[DEBUG] Found int zone setting %q: %d", s.ID, int(floatValue))
			cfg[s.ID] = int(floatValue)
		} else {
			log.Printf("[DEBUG] Unexpected value type found in API zone settings - %q : %#v", s.ID, s.Value)
		}
	}
	// TODO if new settings are found in the api, this will fail

	log.Printf("[DEBUG] Flattened CloudFlare Zone Settings: %#v", cfg)

	return []map[string]interface{}{cfg}
}

func flattenEditableZoneSettings(settings []cloudflare.ZoneSetting) []string {
	ids := make([]string, 0)
	for _, zs := range settings {
		if zs.Editable {
			ids = append(ids, zs.ID)
		}
	}
	log.Printf("[DEBUG] Flattened CloudFlare Editable Zone Settings: %#v", ids)

	return ids
}

func resourceCloudFlareZoneDelete(d *schema.ResourceData, meta interface{}) error {
	/* TODO: keeping this off while testing
	client := meta.(*cloudflare.API)

	_, err := client.DeleteZone(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting CloudFlare zone: %s", err)
	}*/

	return nil
}

func resourceCloudFlareZoneImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	d.Set("name", d.Id())
	zoneId, err := client.ZoneIDByName(d.Id())
	if err != nil {
		return nil, fmt.Errorf("couldn't find zone %q while trying to import it: %q", d.Id(), err)
	}
	d.SetId(zoneId)
	// with_existing_records is not readable, so on import this always has to be false
	d.Set("with_existing_records", false)
	return []*schema.ResourceData{d}, nil
}
