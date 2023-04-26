package sdkv2provider

import (
	"context"
	"errors"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareRegionalHostname() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareRegionalHostnameCreate,
		ReadContext:   resourceCloudflareRegionalHostnameRead,
		UpdateContext: resourceCloudflareRegionalHostnameUpdate,
		DeleteContext: resourceCloudflareRegionalHostnameDelete,
		Description:   heredoc.Doc("Provides a Data Localization Suite Regional Hostname."),
		SchemaVersion: 1,
		Schema:        resourceCloudflareRegionalHostnameSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
	}
}

func resourceCloudflareRegionalHostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := cloudflare.ZoneIdentifier(d.Get(consts.ZoneIDSchemaKey).(string))
	newHostname := cloudflare.CreateDataLocalizationRegionalHostnameParams{
		Hostname:  d.Get("hostname").(string),
		RegionKey: d.Get("region_key").(string),
	}

	r, err := client.CreateDataLocalizationRegionalHostname(ctx, zoneID, newHostname)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(r.Hostname)
	d.Set("hostname", r.Hostname)
	d.Set("region_key", r.RegionKey)
	d.Set("created_on", r.CreatedOn.Format(time.RFC3339Nano))
	return nil
}

func resourceCloudflareRegionalHostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := cloudflare.ZoneIdentifier(d.Get(consts.ZoneIDSchemaKey).(string))
	hostname := d.Get("hostname").(string)

	r, err := client.GetDataLocalizationRegionalHostname(ctx, zoneID, hostname)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(r.Hostname)
	d.Set("hostname", r.Hostname)
	d.Set("region_key", r.RegionKey)
	d.Set("created_on", r.CreatedOn.Format(time.RFC3339Nano))
	return nil
}

func resourceCloudflareRegionalHostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := cloudflare.ZoneIdentifier(d.Get(consts.ZoneIDSchemaKey).(string))

	toUpdate := cloudflare.UpdateDataLocalizationRegionalHostnameParams{
		Hostname:  d.Get("hostname").(string),
		RegionKey: d.Get("region_key").(string),
	}

	if toUpdate.Hostname == "" {
		return diag.FromErr(errors.New("hostname must not be empty"))
	}

	if toUpdate.RegionKey == "" {
		return diag.FromErr(errors.New("region must not be empty"))
	}

	r, err := client.UpdateDataLocalizationRegionalHostname(ctx, zoneID, toUpdate)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(r.Hostname)
	d.Set("hostname", r.Hostname)
	d.Set("region_key", r.RegionKey)
	d.Set("created_on", r.CreatedOn.Format(time.RFC3339Nano))
	return nil
}

func resourceCloudflareRegionalHostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := cloudflare.ZoneIdentifier(d.Get(consts.ZoneIDSchemaKey).(string))

	toDelete := d.Get("hostname").(string)

	if toDelete == "" {
		return diag.FromErr(errors.New("hostname must not be empty"))
	}

	err := client.DeleteDataLocalizationRegionalHostname(ctx, zoneID, toDelete)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
