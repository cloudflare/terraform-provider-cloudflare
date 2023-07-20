package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTieredCache() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTieredCacheSchema(),
		ReadContext:   resourceCloudflareTieredCacheRead,
		UpdateContext: resourceCloudflareTieredCacheUpdate,
		CreateContext: resourceCloudflareTieredCacheUpdate,
		DeleteContext: resourceCloudflareTieredCacheDelete,
		Description: heredoc.Doc(`
			Provides a resource, that manages Cloudflare Tiered Cache settings.
			This allows you to adjust topologies for your zone.
		`),
	}
}

func resourceCloudflareTieredCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var cacheType cloudflare.TieredCacheType
	switch d.Get("cache_type").(string) {
	case "smart":
		cacheType = cloudflare.TieredCacheSmart
		break
	case "generic":
		cacheType = cloudflare.TieredCacheGeneric
		break
	case "off":
		cacheType = cloudflare.TieredCacheOff
		break
	default:
		return diag.FromErr(fmt.Errorf("error updating tiered cache settings: Unsupported cache type requested"))
	}

	_, err := client.SetTieredCache(ctx, cloudflare.ZoneIdentifier(zoneID), cacheType)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating tiered cache settings: %w", err))
	}
	d.SetId(zoneID)
	return resourceCloudflareTieredCacheRead(ctx, d, meta)
}

func resourceCloudflareTieredCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	result, err := client.GetTieredCache(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error retrieving tiered cache settings: %w", err))
	}

	d.SetId(zoneID)
	d.Set("cache_type", result.Type.String())
	return nil
}

func resourceCloudflareTieredCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	_, err := client.DeleteTieredCache(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating deleting tiered cache configuration: %w", err))
	}

	return nil
}
