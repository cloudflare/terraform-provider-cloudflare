package cloudflare

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudflareAuthenticatedOriginPulls() *schema.Resource {
	return &schema.Resource{
		// Pointing `Create` to the `Update `method is intentional. Argo
		// settings are always present, it's just whether or not the value
		// is "on" or "off".
		Create: resourceCloudflareAuthenticatedOriginPullsUpdate,
		Read:   resourceCloudflareAuthenticatedOriginPullsRead,
		Update: resourceCloudflareAuthenticatedOriginPullsUpdate,
		Delete: resourceCloudflareAuthenticatedOriginPullsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAuthenticatedOriginPullsImport,
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
				Required:     true,
			},
		},
	}
}

func resourceCloudflareAuthenticatedOriginPullsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	checksum := stringChecksum(fmt.Sprintf("%s/AOP", zoneID))
	d.SetId(checksum)

	if _, ok := d.GetOk("status"); ok {
		res, err := client.GetAuthenticatedOriginPullsStatus(zoneID)
		if err != nil {
			return errors.Wrap(err, "failed to get global AOP setting")
		}
		d.Set("status", res.Value)
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	status := d.Get("status").(string)

	var globalAOP cloudflare.AuthenticatedOriginPulls
	var err error
	if status == "on" {
		globalAOP, err = client.SetAuthenticatedOriginPullsStatus(zoneID, true)
	} else {
		globalAOP, err = client.SetAuthenticatedOriginPullsStatus(zoneID, false)
	}
	if err != nil {
		return errors.Wrap(err, "failed to update smart routing setting")
	}
	log.Printf("[DEBUG] Global AOP set to: %s", globalAOP.Value)

	return resourceCloudflareAuthenticatedOriginPullsRead(d, meta)
}

func resourceCloudflareAuthenticatedOriginPullsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[DEBUG] Resetting Global AOP set to 'off'")

	_, err := client.SetAuthenticatedOriginPullsStatus(zoneID, false)
	if err != nil {
		return errors.Wrap(err, "failed to update tiered caching setting")
	}

	return nil
}

func resourceCloudflareAuthenticatedOriginPullsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	id := stringChecksum(fmt.Sprintf("%s/AOP", zoneID))
	d.SetId(id)
	d.Set("zone_id", zoneID)

	return []*schema.ResourceData{d}, nil
}
