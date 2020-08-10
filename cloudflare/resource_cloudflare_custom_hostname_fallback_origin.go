package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomHostnameFallbackOrigin() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareCustomHostnameFallbackOriginCreate,
		Read:   resourceCloudflareCustomHostnameFallbackOriginRead,
		Update: resourceCloudflareCustomHostnameFallbackOriginUpdate,
		Delete: resourceCloudflareCustomHostnameFallbackOriginDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareCustomHostnameFallbackOriginImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudflareCustomHostnameFallbackOriginRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	customHostnameFallbackOrigin, err := client.CustomHostnameFallbackOrigin(zoneID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error reading custom hostname fallback origin %q", zoneID))
	}

	d.Set("origin", customHostnameFallbackOrigin.Origin)
	d.Set("status", customHostnameFallbackOrigin.Status)

	return nil
}

func resourceCloudflareCustomHostnameFallbackOriginDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	err := client.DeleteCustomHostnameFallbackOrigin(zoneID)
	if err != nil {
		return errors.Wrap(err, "failed to delete custom hostname fallback origin")
	}

	return nil
}

func resourceCloudflareCustomHostnameFallbackOriginCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	origin := d.Get("origin").(string)

	fallbackOrigin := cloudflare.CustomHostnameFallbackOrigin{
		Origin: origin,
	}

	_, err := client.UpdateCustomHostnameFallbackOrigin(zoneID, fallbackOrigin)
	if err != nil {
		return errors.Wrap(err, "failed to create custom hostname fallback origin")
	}

	d.SetId(zoneID + "-custom_hostname_fallback_origin")

	return resourceCloudflareCustomHostnameFallbackOriginRead(d, meta)
}

func resourceCloudflareCustomHostnameFallbackOriginUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	origin := d.Get("origin").(string)

	fallbackOrigin := cloudflare.CustomHostnameFallbackOrigin{
		Origin: origin,
	}

	_, err := client.UpdateCustomHostnameFallbackOrigin(zoneID, fallbackOrigin)
	if err != nil {
		return errors.Wrap(err, "failed to update custom hostname fallback origin")
	}

	d.SetId(zoneID + "-custom_hostname_fallback_origin")

	return resourceCloudflareCustomHostnameFallbackOriginRead(d, meta)
}

func resourceCloudflareCustomHostnameFallbackOriginImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/origin\"", d.Id())
	}

	zoneID, origin := idAttr[0], idAttr[1]

	log.Printf("[DEBUG] Importing Cloudflare Custom Hostname Fallback Origin: origin %s for zone %s", origin, zoneID)

	d.Set("zone_id", zoneID)
	d.Set("origin", origin)
	d.SetId(zoneID + "-custom_hostname_fallback_origin")

	return []*schema.ResourceData{d}, nil
}
