package cloudflare

import (
	"fmt"
	"log"

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

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
						"advanced_ddos": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Computed: true, //Defaults to on for Business+ plans, off otherwise
						},

						"always_online": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"brotli": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"browser_cache_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  14400,
							/*valid values: 30, 60, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000
							notes: The minimum TTL available depends on the plan level of the zone. (Enterprise = 30, Business = 1800, Pro = 1800, Free = 1800)*/
						},

						"browser_check": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
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
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"origin_error_page_pass_thru": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"sort_query_string_for_cache": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"email_obfuscation": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"hotlink_protection": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"ip_geolocation": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"ipv6": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"websockets": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
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
										Type:     schema.TypeBool, // on/off
										Optional: true,
										Default:  false,
									},

									"html": {
										Type:     schema.TypeBool, // on/off
										Optional: true,
										Default:  false,
									},

									"js": {
										Type:     schema.TypeBool, // on/off
										Optional: true,
										Default:  false,
									},
								},
							},
						},

						"mobile_redirect": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"mirage": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"opportunistic_encryption": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"polish": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"off", "lossless", "lossy"}, false),
						},

						"webp": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"prefetch_preload": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"privacy_pass": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"response_buffering": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
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
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"ssl": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"off", "flexible", "full", "strict"}, false), // depends on plan
						},

						"tls_client_auth": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"true_client_ip_header": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"waf": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"tls_1_2_only": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"tls_1_3": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Computed: true, // default depends on plan level
						},

						"automatic_https_rewrites": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  true,
						},

						"http2": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
						},

						"pseudo_ipv4": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "off",
							ValidateFunc: validation.StringInSlice([]string{"off", "add_header", "overwrite_header"}, false),
						},

						"always_use_https": {
							Type:     schema.TypeBool, // on/off
							Optional: true,
							Default:  false,
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

func resourceCloudFlareRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newRecord := cloudflare.DNSRecord{
		Type:     d.Get("type").(string),
		Name:     d.Get("name").(string),
		Content:  d.Get("value").(string),
		Proxied:  d.Get("proxied").(bool),
		ZoneName: d.Get("domain").(string),
	}

	if priority, ok := d.GetOk("priority"); ok {
		newRecord.Priority = priority.(int)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		newRecord.TTL = ttl.(int)
	}

	// Validate value based on type
	if err := validateRecordName(newRecord.Type, newRecord.Content); err != nil {
		return fmt.Errorf("Error validating record name %q: %s", newRecord.Name, err)
	}

	// Validate type
	if err := validateRecordType(newRecord.Type, newRecord.Proxied); err != nil {
		return fmt.Errorf("Error validating record type %q: %s", newRecord.Type, err)
	}

	zoneId, err := client.ZoneIDByName(newRecord.ZoneName)
	if err != nil {
		return fmt.Errorf("Error finding zone %q: %s", newRecord.ZoneName, err)
	}

	d.Set("zone_id", zoneId)
	newRecord.ZoneID = zoneId

	log.Printf("[DEBUG] CloudFlare Record create configuration: %#v", newRecord)

	r, err := client.CreateDNSRecord(zoneId, newRecord)
	if err != nil {
		return fmt.Errorf("Failed to create record: %s", err)
	}

	// In the Event that the API returns an empty DNS Record, we verify that the
	// ID returned is not the default ""
	if r.Result.ID == "" {
		return fmt.Errorf("Failed to find record in Creat response; Record was empty")
	}

	d.SetId(r.Result.ID)

	log.Printf("[INFO] CloudFlare Record ID: %s", d.Id())

	return resourceCloudFlareRecordRead(d, meta)
}

func resourceCloudFlareRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	domain := d.Get("domain").(string)

	// not all settings are visible to all users, so this might be a subset
	// assume (for now) that user can see/do everything
	record, err := client.ZoneSettings(d.Id())
	if err != nil {
		return err
	}

	d.SetId(record.ID)
	d.Set("hostname", record.Name)
	d.Set("type", record.Type)
	d.Set("value", record.Content)
	d.Set("ttl", record.TTL)
	d.Set("priority", record.Priority)
	d.Set("proxied", record.Proxied)
	d.Set("zone_id", zoneId)

	return nil
}

func resourceCloudFlareRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	updateRecord := cloudflare.DNSRecord{
		ID:       d.Id(),
		Type:     d.Get("type").(string),
		Name:     d.Get("name").(string),
		Content:  d.Get("value").(string),
		ZoneName: d.Get("domain").(string),
		Proxied:  false,
	}

	if priority, ok := d.GetOk("priority"); ok {
		updateRecord.Priority = priority.(int)
	}

	if proxied, ok := d.GetOk("proxied"); ok {
		updateRecord.Proxied = proxied.(bool)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		updateRecord.TTL = ttl.(int)
	}

	zoneId, err := client.ZoneIDByName(updateRecord.ZoneName)
	if err != nil {
		return fmt.Errorf("Error finding zone %q: %s", updateRecord.ZoneName, err)
	}

	updateRecord.ZoneID = zoneId

	log.Printf("[DEBUG] CloudFlare Record update configuration: %#v", updateRecord)
	err = client.UpdateDNSRecord(zoneId, d.Id(), updateRecord)
	if err != nil {
		return fmt.Errorf("Failed to update CloudFlare Record: %s", err)
	}

	return resourceCloudFlareRecordRead(d, meta)
}

func resourceCloudFlareRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	domain := d.Get("domain").(string)

	zoneId, err := client.ZoneIDByName(domain)
	if err != nil {
		return fmt.Errorf("Error finding zone %q: %s", domain, err)
	}

	log.Printf("[INFO] Deleting CloudFlare Record: %s, %s", domain, d.Id())

	err = client.DeleteDNSRecord(zoneId, d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting CloudFlare Record: %s", err)
	}

	return nil
}
