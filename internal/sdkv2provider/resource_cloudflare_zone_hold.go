package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneHold() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneHoldSchema(),
		CreateContext: resourceCloudflareZoneHoldCreate,
		ReadContext:   resourceCloudflareZoneHoldRead,
		UpdateContext: resourceCloudflareZoneHoldUpdate,
		DeleteContext: resourceCloudflareZoneHoldDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareZoneHoldImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Zone Hold resource that prevents adding
			the hostname to another account for use.
		`),
	}
}

func resourceCloudflareZoneHoldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	hold, err := client.CreateZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.CreateZoneHoldParams{
		IncludeSubdomains: cloudflare.BoolPtr(d.Get("include_subdomains").(bool)),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating zone hold for zone %q: %w", zoneID, err))
	}

	d.SetId(zoneID)

	if hold.HoldAfter != nil {
		d.Set("hold_after", hold.HoldAfter.Format(time.RFC3339Nano))
	}

	return resourceCloudflareZoneHoldRead(ctx, d, meta)
}

func resourceCloudflareZoneHoldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("reading zone hold for zone %q", zoneID))

	r, err := client.GetZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.GetZoneHoldParams{})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("zone hold %s no longer exists", zoneID))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding zone hold %q: %w", zoneID, err))
	}

	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("hold", r.Hold)

	if r.HoldAfter != nil {
		d.Set("hold_after", r.HoldAfter.Format(time.RFC3339Nano))
	}

	d.Set("include_subdomains", r.IncludeSubdomains)

	return nil
}

func resourceCloudflareZoneHoldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.Get("hold").(bool) {
		return resourceCloudflareZoneHoldCreate(ctx, d, meta)
	} else {
		return resourceCloudflareZoneHoldDelete(ctx, d, meta)
	}
}

func resourceCloudflareZoneHoldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("deleting zone hold for zone %q", zoneID))

	_, err := client.DeleteZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.DeleteZoneHoldParams{})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error removing zone hold for zone %q: %w", zoneID, err))
	}

	return nil
}

func resourceCloudflareZoneHoldImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set(consts.ZoneIDSchemaKey, d.Id())

	resourceCloudflareZoneHoldRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
