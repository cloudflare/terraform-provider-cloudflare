package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareArgo() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareArgoSchema(),
		CreateContext: resourceCloudflareArgoUpdate,
		ReadContext: resourceCloudflareArgoRead,
		UpdateContext: resourceCloudflareArgoUpdate,
		DeleteContext: resourceCloudflareArgoDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareArgoImport,
		},
	}
}

func resourceCloudflareArgoRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	tieredCaching := d.Get("tiered_caching").(string)
	smartRouting := d.Get("smart_routing").(string)

	log.Printf("[DEBUG] zone ID: %s", zoneID)

	checksum := stringChecksum(fmt.Sprintf("%s/argo", zoneID))
	d.SetId(checksum)
	d.Set("zone_id", zoneID)

	if tieredCaching != "" {
		tieredCaching, err := client.ArgoTieredCaching(context.Background(), zoneID)
		if err != nil {
			return errors.Wrap(err, "failed to get tiered caching setting")
		}

		d.Set("tiered_caching", tieredCaching.Value)
	}

	if smartRouting != "" {
		smartRouting, err := client.ArgoSmartRouting(context.Background(), zoneID)
		if err != nil {
			return errors.Wrap(err, "failed to get smart routing setting")
		}

		d.Set("smart_routing", smartRouting.Value)
	}

	return nil
}

func resourceCloudflareArgoUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	tieredCaching := d.Get("tiered_caching").(string)
	smartRouting := d.Get("smart_routing").(string)

	if smartRouting != "" {
		argoSmartRouting, err := client.UpdateArgoSmartRouting(context.Background(), zoneID, smartRouting)
		if err != nil {
			return errors.Wrap(err, "failed to update smart routing setting")
		}
		log.Printf("[DEBUG] Argo Smart Routing set to: %s", argoSmartRouting.Value)
	}

	if tieredCaching != "" {
		argoTieredCaching, err := client.UpdateArgoTieredCaching(context.Background(), zoneID, tieredCaching)
		if err != nil {
			return errors.Wrap(err, "failed to update tiered caching setting")
		}
		log.Printf("[DEBUG] Argo Tiered Caching set to: %s", argoTieredCaching.Value)
	}

	return resourceCloudflareArgoRead(d, meta)
}

func resourceCloudflareArgoDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[DEBUG] Resetting Argo values to 'off'")

	_, smartRoutingErr := client.UpdateArgoSmartRouting(context.Background(), zoneID, "off")
	if smartRoutingErr != nil {
		return errors.Wrap(smartRoutingErr, "failed to update smart routing setting")
	}

	_, tieredCachingErr := client.UpdateArgoTieredCaching(context.Background(), zoneID, "off")
	if tieredCachingErr != nil {
		return errors.Wrap(tieredCachingErr, "failed to update tiered caching setting")
	}

	return nil
}

func resourceCloudflareArgoImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	id := stringChecksum(fmt.Sprintf("%s/argo", zoneID))
	d.SetId(id)
	d.Set("zone_id", zoneID)

	resourceCloudflareArgoRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
