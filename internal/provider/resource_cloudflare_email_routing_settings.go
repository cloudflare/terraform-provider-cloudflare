package provider

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareEmailRoutingSettings() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareEmailRoutingSettingsSchema(),
		ReadContext:   resourceCloudflareEmailRoutingSettingsRead,
		CreateContext: resourceCloudflareEmailRoutingSettingsCreate,
		DeleteContext: resourceCloudflareEmailRoutingSettingsDelete,
	}
}

func resourceCloudflareEmailRoutingSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	res, err := client.GetEmailRoutingSettings(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error getting email routing settings %q: %w", zoneID, err))
	}

	d.SetId(res.Tag)
	d.Set("name", res.Name)
	d.Set("enabled", res.Enabled)
	d.Set("created", res.Created.String())
	d.Set("modified", res.Modified.String())
	d.Set("skip_wizard", res.SkipWizard)
	d.Set("status", res.Status)

	return nil
}

func resourceCloudflareEmailRoutingSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	_, err := client.EnableEmailRouting(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error enabling email routing %q: %w", zoneID, err))
	}

	return resourceCloudflareEmailRoutingSettingsRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	_, err := client.EnableEmailRouting(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error disabling email routing %q: %w", zoneID, err))
	}

	return nil
}
