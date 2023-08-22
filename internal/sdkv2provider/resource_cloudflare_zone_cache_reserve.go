package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cacheReserveEnabled  = "on"
	cacheReserveDisabled = "off"
)

var cacheReserveNotFoundError *cloudflare.NotFoundError

func resourceCloudflareZoneCacheReserve() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneCacheReserveSchema(),
		CreateContext: resourceCloudflareZoneCacheReserveCreate,
		ReadContext:   resourceCloudflareZoneCacheReserveRead,
		UpdateContext: resourceCloudflareZoneCacheReserveUpdate,
		DeleteContext: resourceCloudflareZoneCacheReserveDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareZoneCacheReserveImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Cache Reserve resource. Cache Reserve can
			increase cache lifetimes by automatically storing all cacheable
			files in Cloudflare's persistent object storage buckets.

			Note: Using Cache Reserve without Tiered Cache is not recommended.
		`),
	}
}

func resourceCloudflareZoneCacheReserveCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	// Ensure an unique resource ID to differentiate it from a Zone ID.
	d.SetId(stringChecksum(fmt.Sprintf("%s/cache-reserve", zoneID)))

	return resourceCloudflareZoneCacheReserveUpdate(ctx, d, meta)
}

func resourceCloudflareZoneCacheReserveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, "reading Cache Reserve", map[string]interface{}{
		"zone_id": zoneID,
	})

	params := cloudflare.GetCacheReserveParams{}
	output, err := client.GetCacheReserve(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		// Zone does not exist?
		if errors.As(err, &cacheReserveNotFoundError) {
			tflog.Warn(ctx, "zone could not be found", map[string]interface{}{
				"zone_id": zoneID,
			})
			d.SetId("")
			return nil
		}
		return diag.Errorf("unable to read Cache Reserve for zone %q: %s", zoneID, err)
	}

	d.Set("enabled", output.Value == cacheReserveEnabled)

	return nil
}

func resourceCloudflareZoneCacheReserveUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	params := cloudflare.UpdateCacheReserveParams{
		Value: cacheReserveDisabled,
	}
	if value, ok := d.GetOk("enabled"); ok && value.(bool) {
		params.Value = cacheReserveEnabled
	}

	tflog.Info(ctx, "setting Cache Reserve", map[string]interface{}{
		"zone_id":       zoneID,
		"cache_reserve": params.Value == cacheReserveEnabled,
	})

	_, err := client.UpdateCacheReserve(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.Errorf("unable to set Cache Reserve for zone %q: %s", zoneID, err)
	}

	return resourceCloudflareZoneCacheReserveRead(ctx, d, meta)
}

func resourceCloudflareZoneCacheReserveDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, "deleting Cache Reserve", map[string]interface{}{
		"zone_id": zoneID,
	})

	// The Cache Reserve does not have a concept of being added or removed,
	// it's either enabled or disabled, and as such, deleting a Cache
	// Reserve simply means disabling it, which is also the default.
	params := cloudflare.UpdateCacheReserveParams{
		Value: cacheReserveDisabled,
	}
	_, err := client.UpdateCacheReserve(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		// Zone does not exist or already had been deleted?
		if errors.As(err, &cacheReserveNotFoundError) {
			tflog.Warn(ctx, "zone could not be found", map[string]interface{}{
				"zone_id": zoneID,
			})
			return nil
		}
		return diag.Errorf("unable to delete Cache Reserve for zone %q: %s", zoneID, err)
	}

	return nil
}

func resourceCloudflareZoneCacheReserveImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	d.SetId(stringChecksum(fmt.Sprintf("%s/cache-reserve", zoneID)))
	d.Set(consts.ZoneIDSchemaKey, zoneID)

	resourceCloudflareZoneCacheReserveRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
