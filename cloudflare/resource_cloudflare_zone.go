package cloudflare

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareZoneCreate,
		Read:   resourceCloudflareZoneRead,
		Update: resourceCloudflareZoneUpdate,
		Delete: resourceCloudflareZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^([a-zA-Z0-9][\\-a-zA-Z0-9]*\\.)+[\\-a-zA-Z0-9]{2,20}$"), ""),
			},
			"jump_start": {
				Type:     schema.TypeBool,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"paused": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vanity_name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"plan": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Free Website", "Business Website", "Pro Website", "Enterprise Website"}, false),
			},
			"meta": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"wildcard_proxiable": {
							Type: schema.TypeBool,
						},
						"phishing_detected": {
							Type: schema.TypeBool,
						},
					},
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
		},
	}
}

func resourceCloudflareZoneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	zoneName := d.Get("zone").(string)
	jumpstart := d.Get("jump_start").(bool)
	organization := cloudflare.Organization{
		ID: client.OrganizationID,
	}

	log.Printf("[INFO] Creating Cloudflare Zone: name %s", zoneName)

	zone, err := client.CreateZone(zoneName, jumpstart, organization)

	if err != nil {
		return fmt.Errorf("Error creating zone %q: %s", zoneName, err)
	}

	d.SetId(zone.ID)

	if paused, ok := d.GetOk("paused"); ok {
		if paused.(bool) == true {
			_, err := client.ZoneSetPaused(zone.ID, paused.(bool))

			if err != nil {
				return fmt.Errorf("Error updating zone_id %q: %s", zone.ID, err)
			}
		}
	}

	if plan, ok := d.GetOk("plan"); ok {
		if err := setRatePlan(plan.(string), client, zone.ID); err != nil {
			return err
		}
	}

	return resourceCloudflareZoneRead(d, meta)
}

func resourceCloudflareZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	zone, err := client.ZoneDetails(zoneID)

	log.Printf("[DEBUG] ZoneDetails: %#v", zone)
	log.Printf("[DEBUG] ZoneDetails error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Zone %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Zone %q: %s", d.Id(), err)
	}

	d.Set("paused", zone.Paused)
	d.Set("vanity_name_servers", zone.VanityNS)
	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("name_servers", zone.NameServers)
	d.Set("meta", flattenMeta(d, zone.Meta))
	d.Set("zone", zone.Name)
	d.Set("plan", zone.Plan.Name)

	return nil
}

func resourceCloudflareZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	log.Printf("[INFO] Updating Cloudflare Zone: id %s", zoneID)

	if paused, ok := d.GetOkExists("paused"); ok {
		log.Printf("[DEBUG] _ paused")

		_, err := client.ZoneSetPaused(zoneID, paused.(bool))

		if err != nil {
			return fmt.Errorf("Error updating zone_id %q: %s", zoneID, err)
		}
	}

	if plan, ok := d.GetOk("plan"); ok {
		if err := setRatePlan(plan.(string), client, zoneID); err != nil {
			return err
		}
	}

	return resourceCloudflareZoneRead(d, meta)
}

func resourceCloudflareZoneDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	log.Printf("[INFO] Deleting Cloudflare Zone: id %s", zoneID)

	_, err := client.DeleteZone(zoneID)

	if err != nil {
		return fmt.Errorf("Error deleting Cloudflare Zone: %s", err)
	}

	return nil
}

func flattenMeta(d *schema.ResourceData, meta cloudflare.ZoneMeta) map[string]interface{} {
	cfg := map[string]interface{}{}

	cfg["wildcard_proxiable"] = strconv.FormatBool(meta.WildcardProxiable)
	cfg["phishing_detected"] = strconv.FormatBool(meta.PhishingDetected)

	log.Printf("[DEBUG] flattenMeta %#v", cfg)

	return cfg
}

func setRatePlan(planName string, client *cloudflare.API, zoneID string) error {
	ratePlan, err := getZonePlanIDFor(client, zoneID, planName)
	if err != nil {
		return fmt.Errorf("Error fetching planName %s for zone %q: %s", planName, zoneID, err)
	}
	if _, err := client.ZoneSetRatePlan(zoneID, *ratePlan); err != nil {
		return fmt.Errorf("Error setting planName %s for zone %q: %s", planName, zoneID, err)
	}
	return nil
}

func getZonePlanIDFor(client *cloudflare.API, zoneID, planName string) (*cloudflare.ZoneRatePlan, error) {
	plans, err := client.AvailableZoneRatePlans(zoneID)
	if err != nil {
		return nil, err
	}
	for _, p := range plans {
		if strings.EqualFold(p.Name, planName) {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("plan %s not found amongst the available plans", planName)
}
