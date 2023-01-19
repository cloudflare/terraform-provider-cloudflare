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

func resourceCloudflareEmailRoutingCatchAll() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareEmailRoutingCatchAllSchema(),
		ReadContext:   resourceCloudflareEmailRoutingCatchAllRead,
		CreateContext: resourceCloudflareEmailRoutingCatchAllUpdate,
		UpdateContext: resourceCloudflareEmailRoutingCatchAllUpdate,
		DeleteContext: resourceCloudflareEmailRoutingCatchAllDelete,
		Description: heredoc.Doc(`
			Provides a resource for managing Email Routing Addresses catch all behaviour.
		`),
	}
}

func resourceCloudflareEmailRoutingCatchAllRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

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
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	updateParams := cloudflare.EmailRoutingCatchAllRule{
		Name:    d.Get("name").(string),
		Enabled: cloudflare.BoolPtr(d.Get("enabled").(bool)),
	}

	updateParams.Matchers, updateParams.Actions = buildMatchersAndActions(d)

	_, err := client.UpdateEmailRoutingCatchAllRule(ctx, cloudflare.ZoneIdentifier(zoneID), updateParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating email routing catch all rule %q: %w", updateParams.Name, err))
	}

	return resourceCloudflareEmailRoutingCatchAllRead(ctx, d, meta)
}

func resourceCloudflareEmailRoutingCatchAllDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	deleteParams := cloudflare.EmailRoutingCatchAllRule{
		Name:    d.Get("name").(string),
		Enabled: cloudflare.BoolPtr(false),
	}
	deleteParams.Matchers, deleteParams.Actions = buildMatchersAndActions(d)

	_, err := client.UpdateEmailRoutingCatchAllRule(ctx, cloudflare.ZoneIdentifier(zoneID), deleteParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error delete email routing catch all rule %q: %w", d.Id(), err))
	}

	return nil
}
