package sdkv2provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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
		url := fmt.Sprintf("/zones/%s/hold", zoneID)
		if d.Get("include_subdomains") != nil {
			url = fmt.Sprintf("%s?include_subdomains=%v", url, d.Get("include_subdomains").(bool))
		}
		data := map[string]interface{}{
			"hold": hold,
		}
		_, err := client.Raw(ctx, "PUT", url, data, http.Header{})
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating zone Hold for zone ID %q: %w", zoneID, err))
		}
	} else {
		url := fmt.Sprintf("/zones/%s/hold", zoneID)
		if d.Get("hold_after") != nil {
			url = fmt.Sprintf("%s?hold_after=%v", url, d.Get("hold_after").(string))
		}
		_, err := client.Raw(ctx, "DELETE", url, nil, http.Header{})
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

	r, err := client.Raw(ctx, "GET", fmt.Sprintf("/zones/%s/hold", zoneID), nil, http.Header{})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Zone Hold %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding zone Hold %q: %w", d.Id(), err))
	}

	var response ZoneHold
	err = json.Unmarshal(r, &response)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing response for id %q", d.Id()))
	}

	tflog.Debug(ctx, fmt.Sprintf("zoneHoldResponse: %#v", response))

	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, d.Id())
	d.Set("hold", response.Hold)
	d.Set("hold_after", response.HoldAfter)
	d.Set("include_subdomains", response.IncludeSubdomains)
	return nil
}

func resourceCloudflareZoneHoldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hold := d.Get("hold").(bool)

	if hold {
		url := fmt.Sprintf("/zones/%s/hold", zoneID)
		if d.Get("include_subdomains") != nil {
			url = fmt.Sprintf("%s?include_subdomains=%v", url, d.Get("include_subdomains").(bool))
		}
		data := map[string]interface{}{
			"hold": hold,
		}
		_, err := client.Raw(ctx, "PUT", url, data, http.Header{})
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating zone Hold for zone ID %q: %w", zoneID, err))
		}
	} else {
		url := fmt.Sprintf("/zones/%s/hold", zoneID)
		if d.Get("hold_after") != nil {
			url = fmt.Sprintf("%s?hold_after=%v", url, d.Get("hold_after").(string))
		}
		_, err := client.Raw(ctx, "DELETE", url, nil, http.Header{})
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

	url := fmt.Sprintf("/zones/%s/hold", zoneID)
	if d.Get("hold_after") != nil {
		url = fmt.Sprintf("%s?hold_after=%v", url, d.Get("hold_after").(string))
	}
	_, err := client.Raw(ctx, "DELETE", url, nil, http.Header{})
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
