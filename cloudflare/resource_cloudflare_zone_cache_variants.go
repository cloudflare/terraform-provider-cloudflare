package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneCacheVariants() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareZoneCacheVariantsSchema(),
		Create: resourceCloudflareZoneCacheVariantsCreate,
		Read:   resourceCloudflareZoneCacheVariantsRead,
		Update: resourceCloudflareZoneCacheVariantsUpdate,
		Delete: resourceCloudflareZoneCacheVariantsDelete,
	}
}

func resourceCloudflareZoneCacheVariantsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Reading Zone Cache Variants in zone %q", d.Id())

	zoneCacheVariants, err := client.ZoneCacheVariants(context.Background(), d.Id())

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Zone %q not found", d.Id())
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("Error reading cache variants for zone %q: %w", d.Id(), err)
		}
	}

	value := zoneCacheVariants.Value
	valueMap := map[string][]string{
		"avif": value.Avif,
		"bmp":  value.Bmp,
		"gif":  value.Gif,
		"jpeg": value.Jpeg,
		"jpg":  value.Jpg,
		"jp2":  value.Jp2,
		"jpg2": value.Jpg2,
		"png":  value.Png,
		"tif":  value.Tif,
		"tiff": value.Tiff,
		"webp": value.Webp,
	}

	for k, v := range valueMap {
		if err := d.Set(k, v); err != nil {
			return fmt.Errorf("failed to set %v: %w", k, err)
		}
	}

	return nil
}

func resourceCloudflareZoneCacheVariantsCreate(d *schema.ResourceData, meta interface{}) error {
	zoneID := d.Get("zone_id").(string)
	d.SetId(zoneID)

	log.Printf("[INFO] Creating Zone Cache Variants for zone ID: %q", d.Id())

	return resourceCloudflareZoneCacheVariantsUpdate(d, meta)
}

func resourceCloudflareZoneCacheVariantsUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*cloudflare.API)

	variantsValue := cacheVariantsValuesFromResource(d)
	log.Printf("[INFO] Setting Zone Cache Variants to struct: %+v for zone ID: %q", variantsValue, d.Id())

	_, err := client.UpdateZoneCacheVariants(context.Background(), d.Id(), variantsValue)

	if err != nil {
		return fmt.Errorf("error setting cache variants for zone %q: %w", d.Id(), err)
	}

	return resourceCloudflareZoneCacheVariantsRead(d, meta)
}

func resourceCloudflareZoneCacheVariantsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting Zone Cache Variants for zone ID: %q", d.Id())

	err := client.DeleteZoneCacheVariants(context.Background(), d.Id())

	if err != nil {
		return fmt.Errorf("error deleting cache variants for zone %v: %w", d.Id(), err)
	}

	return nil
}

func cacheVariantsValuesFromResource(d *schema.ResourceData) cloudflare.ZoneCacheVariantsValues {
	variantsValue := cloudflare.ZoneCacheVariantsValues{}

	if value, ok := d.GetOk("avif"); ok {
		variantsValue.Avif = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("bmp"); ok {
		variantsValue.Bmp = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("gif"); ok {
		variantsValue.Gif = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("jpeg"); ok {
		variantsValue.Jpeg = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("jpg"); ok {
		variantsValue.Jpg = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("jp2"); ok {
		variantsValue.Jp2 = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("jpg2"); ok {
		variantsValue.Jpg2 = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("png"); ok {
		variantsValue.Png = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("tif"); ok {
		variantsValue.Tif = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("tiff"); ok {
		variantsValue.Tiff = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	if value, ok := d.GetOk("webp"); ok {
		variantsValue.Webp = expandInterfaceToStringList(value.(*schema.Set).List())
	}

	return variantsValue
}
