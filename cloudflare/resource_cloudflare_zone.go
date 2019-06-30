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

// we enforce the use of the Cloudflare API 'legacy_id' field until the mapping of plan is fixed in cloudflare-go
const (
	planIDFree       = "free"
	planIDPro        = "pro"
	planIDBusiness   = "business"
	planIDEnterprise = "enterprise"
)

// we keep a private map and we will have a function to check and validate the descriptive name from the RatePlan API with the legacy_id
var idForName = map[string]string{
	"Free Website":       planIDFree,
	"Pro Website":        planIDPro,
	"Business Website":   planIDBusiness,
	"Enterprise Website": planIDEnterprise,
}

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
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{planIDFree, planIDPro, planIDBusiness, planIDEnterprise}, false),
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
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"full", "partial"}, false),
				Default:      "full",
				Optional:     true,
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
	zoneType := d.Get("type").(string)
	organization := cloudflare.Organization{
		ID: client.OrganizationID,
	}

	log.Printf("[INFO] Creating Cloudflare Zone: name %s", zoneName)

	zone, err := client.CreateZone(zoneName, jumpstart, organization, zoneType)

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
		if err := setRatePlan(client, zone.ID, plan.(string)); err != nil {
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
	d.Set("plan", planIDForName(zone.Plan.Name))

	return nil
}

func resourceCloudflareZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	log.Printf("[INFO] Updating Cloudflare Zone: id %s", zoneID)

	if paused, ok := d.GetOkExists("paused"); ok && d.HasChange("paused") {
		log.Printf("[DEBUG] _ paused")

		_, err := client.ZoneSetPaused(zoneID, paused.(bool))

		if err != nil {
			return fmt.Errorf("Error updating zone_id %q: %s", zoneID, err)
		}
	}

	if plan, ok := d.GetOk("plan"); ok {
		if err := setRatePlan(client, zoneID, plan.(string)); err != nil {
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

func setRatePlan(client *cloudflare.API, zoneID string, planID string) error {
	plan, err := getAvailableZonePlan(client, zoneID, planID)
	if err != nil {
		return fmt.Errorf("Error fetching plans %s for zone %q: %s", planID, zoneID, err)
	}
	log.Printf("[DEBUG] ratePlan = %#v", plan)
	if _, err := client.ZoneSetPlan(zoneID, *plan); err != nil {
		return fmt.Errorf("Error setting plan %s for zone %q: %s", planID, zoneID, err)
	}
	return nil
}

func getAvailableZonePlan(client *cloudflare.API, zoneID, planID string) (*cloudflare.ZonePlan, error) {
	plans, err := client.AvailableZonePlans(zoneID)
	if err != nil {
		return nil, err
	}
	for _, p := range plans {
		if strings.EqualFold(p.LegacyID, planID) {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("plan '%s' not found amongst the available plans", planID)
}

func planIDForName(name string) string {
	if p, ok := idForName[name]; ok {
		return p
	}
	return ""
}

func planNameForID(id string) string {
	for k, v := range idForName {
		if strings.EqualFold(id, v) {
			return k
		}
	}
	return ""
}
