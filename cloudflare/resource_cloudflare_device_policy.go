package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDevicePolicyCertificates() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareDevicePolicyCertificatesSchema(),
		Create: resourceCloudflareDevicePolicyCertificateUpdate,
		Read:   resourceCloudflareDevicePolicyCertificateRead,
		Update: resourceCloudflareDevicePolicyCertificateUpdate,
		Delete: resourceCloudflareDevicePolicyCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareDevicePolicyCertificateImport,
		},
	}
}

func resourceCloudflareDevicePolicyCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	enabled := d.Get("enabled").(bool)

	log.Printf("[DEBUG] Updating Cloudflare device policy certificate: zoneID=%s enabled=%t", zoneID, enabled)

	_, err := client.UpdateDeviceClientCertificatesZone(context.Background(), zoneID, enabled)
	if err != nil {
		return fmt.Errorf("error updating Cloudflare device policy certificate %q: %w", zoneID, err)
	}

	d.SetId(zoneID)
	return nil
}

func resourceCloudflareDevicePolicyCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	enabled, err := client.GetDeviceClientCertificatesZone(context.Background(), zoneID)
	if err != nil {
		return fmt.Errorf("error reading device policy certificate setting %q: %w", zoneID, err)
	}

	d.SetId(zoneID)
	d.Set("enabled", enabled.Result.Enable)
	return nil
}

func resourceCloudflareDevicePolicyCertificateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	log.Printf("[DEBUG] Importing Cloudflare device policy certificate setting: zoneID=%s", zoneID)

	d.SetId(zoneID)
	d.Set("zone_id", zoneID)
	err := resourceCloudflareDevicePolicyCertificateRead(d, meta)
	return []*schema.ResourceData{d}, err
}

func resourceCloudflareDevicePolicyCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
