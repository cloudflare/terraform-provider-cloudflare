package sdkv2provider

import (
	"context"
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceCloudflareWaitingRoomSettingsRead,
		UpdateContext: resourceCloudflareWaitingRoomSettingsUpdate,
		CreateContext: resourceCloudflareWaitingRoomSettingsUpdate,
		DeleteContext: schema.NoopContext,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWaitingRoomSettingsImport,
		},

		Schema:      resourceCloudflareWaitingRoomSettingsSchema(),
		Description: "Configure zone-wide settings for Cloudflare waiting rooms.",
	}
}

func buildWaitingRoomSettingsUpdate(d *schema.ResourceData) cloudflare.UpdateWaitingRoomSettingsParams {
	searchEngineCrawlerBypass := d.Get("search_engine_crawler_bypass").(bool)
	return cloudflare.UpdateWaitingRoomSettingsParams{
		SearchEngineCrawlerBypass: &searchEngineCrawlerBypass,
	}
}

func resourceCloudflareWaitingRoomSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	waitingRoomSettings, err := client.GetWaitingRoomSettings(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error getting waiting room settings for zone %q: %w", zoneID, err))
	}
	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("search_engine_crawler_bypass", waitingRoomSettings.SearchEngineCrawlerBypass)
	return nil
}

func resourceCloudflareWaitingRoomSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	waitingRoomSettings := buildWaitingRoomSettingsUpdate(d)

	_, err := client.UpdateWaitingRoomSettings(ctx, cloudflare.ZoneIdentifier(zoneID), waitingRoomSettings)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating waiting room settings for zone %q: %w", zoneID, err))
	}

	return resourceCloudflareWaitingRoomSettingsRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomSettingsImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	resourceCloudflareWaitingRoomSettingsRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
