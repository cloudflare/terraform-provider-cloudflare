package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/idna"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// we enforce the use of the Cloudflare API 'legacy_id' field until the mapping of plan is fixed in cloudflare-go
const (
	planIDFree       = "free"
	planIDPro        = "pro"
	planIDBusiness   = "business"
	planIDEnterprise = "enterprise"

	planIDPartnerFree            = "partners_free"
	planIDPartnerPro             = "partners_pro"
	planIDPartnerBusiness        = "partners_business"
	planIDPartnerEnterprise      = "partners_enterprise"
	planIDPartnerWorkers         = "partners_workers_ss"
	planIDPartnerImageEnterprise = "image_resizing_enterprise"
)

// we keep a private map and we will have a function to check and validate the descriptive name from the RatePlan API with the legacy_id
var idForName = map[string]string{
	"Free Website":       planIDFree,
	"Pro Website":        planIDPro,
	"Business Website":   planIDBusiness,
	"Enterprise Website": planIDEnterprise,

	"Partners Free Plan":         planIDPartnerFree,
	"Partners Professional Plan": planIDPartnerPro,
	"Partners Business Plan":     planIDPartnerBusiness,
	"Partners Enterprise Plan":   planIDPartnerEnterprise,
	"Cloudflare Workers":         planIDPartnerWorkers,
	"Image Resizing Ent":         planIDPartnerImageEnterprise,
}

// maintain a mapping for the subscription API term for rate plans to
// the one we are expecting end users to use.
var subscriptionIDOfRatePlans = map[string]string{
	planIDFree:       "CF_FREE",
	planIDPro:        "CF_PRO",
	planIDBusiness:   "CF_BIZ",
	planIDEnterprise: "CF_ENT",

	planIDPartnerFree:            "PARTNERS_FREE",
	planIDPartnerPro:             "PARTNERS_PRO",
	planIDPartnerBusiness:        "PARTNERS_BIZ",
	planIDPartnerEnterprise:      "PARTNERS_ENT",
	planIDPartnerWorkers:         "PARTNERS_WORKERS_SS",
	planIDPartnerImageEnterprise: "IMAGE_RESIZING_ENT",
}

func resourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareZoneSchema(),
		Create: resourceCloudflareZoneCreate,
		Read:   resourceCloudflareZoneRead,
		Update: resourceCloudflareZoneUpdate,
		Delete: resourceCloudflareZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCloudflareZoneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	zoneName := d.Get("zone").(string)
	jumpstart := d.Get("jump_start").(bool)
	zoneType := d.Get("type").(string)
	account := cloudflare.Account{
		ID: client.AccountID,
	}

	log.Printf("[INFO] Creating Cloudflare Zone: name %s", zoneName)

	zone, err := client.CreateZone(context.Background(), zoneName, jumpstart, account, zoneType)

	if err != nil {
		return fmt.Errorf("error creating zone %q: %s", zoneName, err)
	}

	d.SetId(zone.ID)

	if paused, ok := d.GetOk("paused"); ok {
		if paused.(bool) == true {
			_, err := client.ZoneSetPaused(context.Background(), zone.ID, paused.(bool))

			if err != nil {
				return fmt.Errorf("error updating zone_id %q: %s", zone.ID, err)
			}
		}
	}

	if plan, ok := d.GetOk("plan"); ok {
		if err := setRatePlan(client, zone.ID, plan.(string), true, d); err != nil {
			return err
		}
	}

	if ztype, ok := d.GetOk("type"); ok {
		_, err := client.ZoneSetType(context.Background(), zone.ID, ztype.(string))
		if err != nil {
			return fmt.Errorf("error setting type on zone ID %q: %s", zone.ID, err)
		}
	}

	return resourceCloudflareZoneRead(d, meta)
}

func resourceCloudflareZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	zone, err := client.ZoneDetails(context.Background(), zoneID)

	log.Printf("[DEBUG] ZoneDetails: %#v", zone)
	log.Printf("[DEBUG] ZoneDetails error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Zone %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Zone %q: %s", d.Id(), err)
	}

	// In the cases where the zone isn't completely setup yet, we need to
	// check the `status` field and should it be pending, use the `Name`
	// from `zone.PlanPending` instead to account for paid plans.
	var plan string
	if zone.Status == "pending" && zone.PlanPending.Name != "" {
		plan = zone.PlanPending.Name
	} else {
		plan = zone.Plan.Name
	}

	d.Set("paused", zone.Paused)
	d.Set("vanity_name_servers", zone.VanityNS)
	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("name_servers", zone.NameServers)
	d.Set("meta", flattenMeta(d, zone.Meta))
	d.Set("zone", zone.Name)
	d.Set("plan", planIDForName(plan))
	d.Set("verification_key", zone.VerificationKey)

	return nil
}

func resourceCloudflareZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()
	zone, _ := client.ZoneDetails(context.Background(), zoneID)

	log.Printf("[INFO] Updating Cloudflare Zone: id %s", zoneID)

	if paused, ok := d.GetOkExists("paused"); ok && d.HasChange("paused") {
		log.Printf("[DEBUG] _ paused")

		_, err := client.ZoneSetPaused(context.Background(), zoneID, paused.(bool))

		if err != nil {
			return fmt.Errorf("error setting paused for zone ID %q: %s", zoneID, err)
		}
	}

	if ztype, ok := d.GetOkExists("type"); ok && d.HasChange("type") {
		_, err := client.ZoneSetType(context.Background(), zoneID, ztype.(string))

		if err != nil {
			return fmt.Errorf("error setting type for on zone ID %q: %s", zoneID, err)
		}
	}

	// In the cases where the zone isn't completely setup yet, we need to
	// check the `status` field and should it be pending, use the `LegacyID`
	// from `zone.PlanPending` instead to account for paid plans.
	if zone.Status == "pending" && zone.PlanPending.Name != "" {
		d.Set("plan", zone.PlanPending.LegacyID)
	}

	if change := d.HasChange("plan"); change {
		// If we're upgrading from a free plan, we need to use POST (not PUT) as the
		// the subscription needs to be created, not modified despite the resource
		// already existing.
		existingPlan, newPlan := d.GetChange("plan")
		wasFreePlan := existingPlan.(string) == "free"
		planID := newPlan.(string)

		if err := setRatePlan(client, zoneID, planID, wasFreePlan, d); err != nil {
			return err
		}
	}

	return resourceCloudflareZoneRead(d, meta)
}

func resourceCloudflareZoneDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()

	log.Printf("[INFO] Deleting Cloudflare Zone: id %s", zoneID)

	_, err := client.DeleteZone(context.Background(), zoneID)

	if err != nil {
		return fmt.Errorf("error deleting Cloudflare Zone: %s", err)
	}

	return nil
}

func flattenMeta(d *schema.ResourceData, meta cloudflare.ZoneMeta) map[string]interface{} {
	cfg := map[string]interface{}{}

	cfg["wildcard_proxiable"] = meta.WildcardProxiable
	cfg["phishing_detected"] = meta.PhishingDetected

	log.Printf("[DEBUG] flattenMeta %#v", cfg)

	return cfg
}

// setRatePlan handles the internals of creating or updating a zone
// subscription rate plan.
func setRatePlan(client *cloudflare.API, zoneID, planID string, isNewPlan bool, d *schema.ResourceData) error {
	if isNewPlan {
		// A free rate plan is the default so no need to explicitly make another
		// HTTP call to set it.
		if subscriptionIDOfRatePlans[planID] != planIDFree {
			if err := client.ZoneSetPlan(context.Background(), zoneID, subscriptionIDOfRatePlans[planID]); err != nil {
				return fmt.Errorf("error setting plan %s for zone %q: %s", planID, zoneID, err)
			}
		}
	} else {
		if err := client.ZoneUpdatePlan(context.Background(), zoneID, subscriptionIDOfRatePlans[planID]); err != nil {
			return fmt.Errorf("error updating plan %s for zone %q: %s", planID, zoneID, err)
		}
	}

	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		zone, _ := client.ZoneDetails(context.Background(), zoneID)

		if zone.Plan.LegacyID != planID {
			return resource.RetryableError(fmt.Errorf("plan ID change has not yet propagated"))
		}

		return nil
	})
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

// zoneDiffFunc is a DiffSuppressFunc that accepts two strings and then converts
// them to unicode before performing the comparison whether or not the value has
// changed. This ensures that zones which could be either are evaluated
// consistently and align with what the Cloudflare API returns.
func zoneDiffFunc(k, old, new string, d *schema.ResourceData) bool {
	var p *idna.Profile
	p = idna.New()
	unicodeOld, _ := p.ToUnicode(old)
	unicodeNew, _ := p.ToUnicode(new)

	return unicodeOld == unicodeNew
}
