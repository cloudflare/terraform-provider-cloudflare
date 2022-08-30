package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareEmailRoutingCatchAll() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareEmailRoutingRuleSchema(),
		ReadContext:   resourceCloudflareEmailRoutingCatchAllRead,
		CreateContext: resourceCloudflareEmailRoutingCatchAllUpdate,
		UpdateContext: resourceCloudflareEmailRoutingCatchAllUpdate,
		DeleteContext: resourceCloudflareEmailRoutingCatchAllDelete,
	}
}

func resourceCloudflareEmailRoutingCatchAllRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	res, err := client.GetEmailRoutingCatchAllRule(ctx, cloudflare.AccountIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading email routing catch all rule %q: %w", d.Id(), err))
	}
	d.SetId(res.Tag)
	d.Set("name", res.Name)
	d.Set("enabled", res.Enabled)
	return nil
}

func resourceCloudflareEmailRoutingCatchAllUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	createParams := cloudflare.EmailRoutingCatchAllRule{}
	createParams.Name = d.Get("name").(string)
	createParams.Enabled = cloudflare.BoolPtr(d.Get("enabled").(bool))

	createParams.Matchers, createParams.Actions = buildMatchersAndActions(d)

	_, err := client.UpdateEmailRoutingCatchAllRule(ctx, cloudflare.ZoneIdentifier(zoneID), createParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating email routing catch all rule %q: %w", createParams.Name, err))
	}

	return resourceCloudflareEmailRoutingCatchAllRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingCatchAllDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	createParams := cloudflare.EmailRoutingCatchAllRule{}
	createParams.Name = d.Get("name").(string)
	createParams.Enabled = cloudflare.BoolPtr(false)

	createParams.Matchers, createParams.Actions = buildMatchersAndActions(d)

	_, err := client.UpdateEmailRoutingCatchAllRule(ctx, cloudflare.ZoneIdentifier(zoneID), createParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error delete email routing catch all rule %q: %w", d.Id(), err))
	}

	return nil
}