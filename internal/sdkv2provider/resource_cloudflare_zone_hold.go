package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ZoneHold struct {
	Hold              bool   `json:"hold,omitempty"`
	IncludeSubdomains bool   `json:"include_subdomains,omitempty"`
	HoldAfter         string `json:"hold_after,omitempty"`
}

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
			Provides a Cloudflare Zone Hold resource. Zone holds prevent other teams in your organization from adding zones that are already active in another account.
		`),
	}
}

func resourceCloudflareZoneHoldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hold := d.Get("hold").(bool)

	if hold {
		params := cloudflare.CreateZoneHoldParams{}
		if d.Get("include_subdomains") != nil {
			sub := d.Get("include_subdomains").(bool)
			params.IncludeSubdomains = &sub
		}
		_, err := client.CreateZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating zone Hold for zone ID %q: %w", zoneID, err))
		}
	} else {
		params := cloudflare.DeleteZoneHoldParams{}
		if d.Get("hold_after") != nil {
			holdAfter := d.Get("hold_after").(string)
			t, err := time.Parse(time.RFC3339, holdAfter)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error parsing hold_after %q: %w", holdAfter, err))
			}
			params.HoldAfter = &t
		}
		_, err := client.DeleteZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error removing zone Hold for zone ID %q: %w", zoneID, err))
		}
	}

	d.SetId(zoneID)
	tflog.Info(ctx, fmt.Sprintf("Cloudflare Zone Hold ID: %s", d.Id()))

	return resourceCloudflareZoneHoldRead(ctx, d, meta)
}

func resourceCloudflareZoneHoldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Reading Cloudflare Zone Hold: id %s for zone %s", d.Id(), zoneID))

	r, err := client.GetZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.GetZoneHoldParams{})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Zone Hold %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding zone Hold %q: %w", d.Id(), err))
	}

	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, d.Id())
	d.Set("hold", r.Hold)
	d.Set("hold_after", r.HoldAfter.Format(time.RFC3339))
	d.Set("include_subdomains", r.IncludeSubdomains)
	return nil
}

func resourceCloudflareZoneHoldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hold := d.Get("hold").(bool)

	if hold {
		params := cloudflare.CreateZoneHoldParams{}
		if d.Get("include_subdomains") != nil {
			sub := d.Get("include_subdomains").(bool)
			params.IncludeSubdomains = &sub
		}
		_, err := client.CreateZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating zone Hold for zone ID %q: %w", zoneID, err))
		}
	} else {
		params := cloudflare.DeleteZoneHoldParams{}
		if d.Get("hold_after") != nil {
			holdAfter := d.Get("hold_after").(string)
			t, err := time.Parse(time.RFC3339, holdAfter)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error parsing hold_after %q: %w", holdAfter, err))
			}
			params.HoldAfter = &t
		}
		_, err := client.DeleteZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error removing zone Hold for zone ID %q: %w", zoneID, err))
		}
	}

	d.SetId(zoneID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Zone Hold ID: %s", d.Id()))

	return resourceCloudflareZoneHoldRead(ctx, d, meta)
}

func resourceCloudflareZoneHoldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Zone Hold: id %s for zone %s", d.Id(), zoneID))

	params := cloudflare.DeleteZoneHoldParams{}
	if d.Get("hold_after") != nil {
		holdAfter := d.Get("hold_after").(string)
		t, err := time.Parse(time.RFC3339, holdAfter)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error parsing hold_after %q: %w", holdAfter, err))
		}
		params.HoldAfter = &t
	}
	_, err := client.DeleteZoneHold(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error removing zone Hold for zone ID %q: %w", zoneID, err))
	}

	return nil
}

func resourceCloudflareZoneHoldImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set(consts.ZoneIDSchemaKey, d.Id())
	resourceCloudflareZoneHoldRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
