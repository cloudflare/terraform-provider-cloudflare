package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDevicePolicyCertificates() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDevicePolicyCertificatesSchema(),
		CreateContext: resourceCloudflareDevicePolicyCertificateUpdate,
		ReadContext:   resourceCloudflareDevicePolicyCertificateRead,
		UpdateContext: resourceCloudflareDevicePolicyCertificateUpdate,
		DeleteContext: resourceCloudflareDevicePolicyCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDevicePolicyCertificateImport,
		},
	}
}

func resourceCloudflareDevicePolicyCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	enabled := d.Get("enabled").(bool)

	log.Printf("[DEBUG] Updating Cloudflare device policy certificate: zoneID=%s enabled=%t", zoneID, enabled)

	_, err := client.UpdateDeviceClientCertificatesZone(ctx, zoneID, enabled)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Cloudflare device policy certificate %q: %w", zoneID, err))
	}

	d.SetId(zoneID)
	return nil
}

func resourceCloudflareDevicePolicyCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	enabled, err := client.GetDeviceClientCertificatesZone(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading device policy certificate setting %q: %w", zoneID, err))
	}

	d.SetId(zoneID)
	d.Set("enabled", enabled.Result.Enabled)
	return nil
}

func resourceCloudflareDevicePolicyCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	log.Printf("[DEBUG] Importing Cloudflare device policy certificate setting: zoneID=%s", zoneID)

	d.SetId(zoneID)
	d.Set("zone_id", zoneID)

	resourceCloudflareDevicePolicyCertificateRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareDevicePolicyCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
