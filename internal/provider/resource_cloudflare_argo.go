package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareArgo() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareArgoSchema(),
		CreateContext: resourceCloudflareArgoUpdate,
		ReadContext:   resourceCloudflareArgoRead,
		UpdateContext: resourceCloudflareArgoUpdate,
		DeleteContext: resourceCloudflareArgoDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareArgoImport,
		},
	}
}

func resourceCloudflareArgoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	tieredCaching := d.Get("tiered_caching").(string)
	smartRouting := d.Get("smart_routing").(string)

	log.Printf("[DEBUG] zone ID: %s", zoneID)

	checksum := stringChecksum(fmt.Sprintf("%s/argo", zoneID))
	d.SetId(checksum)
	d.Set("zone_id", zoneID)

	if tieredCaching != "" {
		tieredCaching, err := client.ArgoTieredCaching(ctx, zoneID)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to get tiered caching setting"))
		}

		d.Set("tiered_caching", tieredCaching.Value)
	}

	if smartRouting != "" {
		smartRouting, err := client.ArgoSmartRouting(ctx, zoneID)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to get smart routing setting"))
		}

		d.Set("smart_routing", smartRouting.Value)
	}

	return nil
}

func resourceCloudflareArgoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	tieredCaching := d.Get("tiered_caching").(string)
	smartRouting := d.Get("smart_routing").(string)

	if smartRouting != "" {
		argoSmartRouting, err := client.UpdateArgoSmartRouting(ctx, zoneID, smartRouting)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to update smart routing setting"))
		}
		log.Printf("[DEBUG] Argo Smart Routing set to: %s", argoSmartRouting.Value)
	}

	if tieredCaching != "" {
		argoTieredCaching, err := client.UpdateArgoTieredCaching(ctx, zoneID, tieredCaching)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to update tiered caching setting"))
		}
		log.Printf("[DEBUG] Argo Tiered Caching set to: %s", argoTieredCaching.Value)
	}

	return resourceCloudflareArgoRead(ctx, d, meta)
}

func resourceCloudflareArgoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[DEBUG] Resetting Argo values to 'off'")

	_, smartRoutingErr := client.UpdateArgoSmartRouting(ctx, zoneID, "off")
	if smartRoutingErr != nil {
		return diag.FromErr(errors.Wrap(smartRoutingErr, "failed to update smart routing setting"))
	}

	_, tieredCachingErr := client.UpdateArgoTieredCaching(ctx, zoneID, "off")
	if tieredCachingErr != nil {
		return diag.FromErr(errors.Wrap(tieredCachingErr, "failed to update tiered caching setting"))
	}

	return nil
}

func resourceCloudflareArgoImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	id := stringChecksum(fmt.Sprintf("%s/argo", zoneID))
	d.SetId(id)
	d.Set("zone_id", zoneID)

	resourceCloudflareArgoRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
