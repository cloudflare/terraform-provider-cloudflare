package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCloudflareAccessApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccessApplicationCreate,
		Read:   resourceCloudflareAccessApplicationRead,
		Update: resourceCloudflareAccessApplicationUpdate,
		Delete: resourceCloudflareAccessApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccessApplicationImport,
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"aud": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"session_duration": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "24h",
				ValidateFunc: validation.StringInSlice([]string{"30m", "6h", "12h", "24h", "168h", "730h"}, false),
			},
		},
	}
}

func resourceCloudflareAccessApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	newAccessApplication := cloudflare.AccessApplication{
		Name:            d.Get("name").(string),
		Domain:          d.Get("domain").(string),
		SessionDuration: d.Get("session_duration").(string),
	}

	log.Printf("[DEBUG] Creating Cloudflare Access Application from struct: %+v", newAccessApplication)

	accessApplication, err := client.CreateAccessApplication(zoneID, newAccessApplication)
	if err != nil {
		return fmt.Errorf("error creating Access Application for zone %q: %s", zoneID, err)
	}

	d.SetId(accessApplication.ID)

	return resourceCloudflareAccessApplicationRead(d, meta)
}

func resourceCloudflareAccessApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	accessApplication, err := client.AccessApplication(zoneID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Access Application %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Access Application %q: %s", d.Id(), err)
	}

	d.Set("aud", accessApplication.AUD)
	d.Set("session_duration", accessApplication.SessionDuration)
	d.Set("domain", accessApplication.Domain)

	return nil
}

func resourceCloudflareAccessApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	updatedAccessApplication := cloudflare.AccessApplication{
		ID:              d.Id(),
		Name:            d.Get("name").(string),
		Domain:          d.Get("domain").(string),
		SessionDuration: d.Get("session_duration").(string),
	}

	log.Printf("[DEBUG] Updating Cloudflare Access Application from struct: %+v", updatedAccessApplication)

	accessApplication, err := client.UpdateAccessApplication(zoneID, updatedAccessApplication)
	if err != nil {
		return fmt.Errorf("error updating Access Application for zone %q: %s", zoneID, err)
	}

	if accessApplication.ID == "" {
		return fmt.Errorf("failed to find Access Application ID in update response; resource was empty")
	}

	return resourceCloudflareAccessApplicationRead(d, meta)
}

func resourceCloudflareAccessApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	appID := d.Id()

	log.Printf("[DEBUG] Deleting Cloudflare Access Application using ID: %s", appID)

	err := client.DeleteAccessApplication(zoneID, appID)
	if err != nil {
		return fmt.Errorf("error deleting Access Application for zone %q: %s", zoneID, err)
	}

	resourceCloudflareAccessApplicationRead(d, meta)

	return nil
}

func resourceCloudflareAccessApplicationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/accessApplicationID\"", d.Id())
	}

	zoneID, accessApplicationID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Access Application: id %s for zone %s", accessApplicationID, zoneID)

	d.Set("zone_id", zoneID)
	d.SetId(accessApplicationID)

	resourceCloudflareAccessApplicationRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
