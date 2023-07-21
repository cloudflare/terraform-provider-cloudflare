package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareRegionalTieredCache() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareRegionalTieredCacheSchema(),
		CreateContext: resourceCloudflareRegionalTieredCacheUpdate,
		ReadContext:   resourceCloudflareRegionalTieredCacheRead,
		UpdateContext: resourceCloudflareRegionalTieredCacheUpdate,
		DeleteContext: resourceCloudflareRegionalTieredCacheDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareRegionalTieredCacheImport,
		},
		Description: heredoc.Doc(`
			Instructs Cloudflare to check a regional hub data center on the way to your upper tier.
			This can help improve performance for smart and custom tiered cache topologies.
		`),
	}
}

func resourceCloudflareRegionalTieredCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	regionalTieredCache, err := client.GetRegionalTieredCache(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.GetRegionalTieredCacheParams{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to get tiered caching setting"))
	}

	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("value", regionalTieredCache.Value)
	return nil
}

func resourceCloudflareRegionalTieredCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	value := d.Get("value").(string)

	tflog.Debug(ctx, fmt.Sprintf("Setting Regional Tiered Cache value to: %s", value))

	_, err := client.UpdateRegionalTieredCache(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateRegionalTieredCacheParams{Value: value})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to update Regional Tiered Cache value"))
	}

	return resourceCloudflareRegionalTieredCacheRead(ctx, d, meta)
}

func resourceCloudflareRegionalTieredCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Resetting Regional Tiered Cache value to 'off'"))

	_, err := client.UpdateRegionalTieredCache(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateRegionalTieredCacheParams{Value: "off"})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to reset Regional Tiered Cache value"))
	}

	return nil
}

func resourceCloudflareRegionalTieredCacheImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set(consts.ZoneIDSchemaKey, d.Id())

	resourceCloudflareRegionalTieredCacheRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
