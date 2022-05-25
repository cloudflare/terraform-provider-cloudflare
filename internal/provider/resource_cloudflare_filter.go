package provider

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFilter() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareFilterSchema(),
		CreateContext: resourceCloudflareFilterCreate,
		ReadContext:   resourceCloudflareFilterRead,
		UpdateContext: resourceCloudflareFilterUpdate,
		DeleteContext: resourceCloudflareFilterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareFilterImport,
		},
	}
}

func resourceCloudflareFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	var err error

	var newFilter cloudflare.Filter

	if paused, ok := d.GetOk("paused"); ok {
		newFilter.Paused = paused.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newFilter.Description = description.(string)
	}

	if expression, ok := d.GetOk("expression"); ok {
		newFilter.Expression = expression.(string)
	}

	if ref, ok := d.GetOk("ref"); ok {
		newFilter.Ref = ref.(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Filter from struct: %+v", newFilter))

	r, err := client.CreateFilters(ctx, zoneID, []cloudflare.Filter{newFilter})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Filter for zone %q: %w", zoneID, err))
	}

	if len(r) == 0 {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r[0].ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Filter ID: %s", d.Id()))

	return resourceCloudflareFilterRead(ctx, d, meta)
}

func resourceCloudflareFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	tflog.Debug(ctx, fmt.Sprintf("Getting a Filter record for zone %q, id %s", zoneID, d.Id()))
	filter, err := client.Filter(ctx, zoneID, d.Id())

	tflog.Debug(ctx, fmt.Sprintf("filter: %#v", filter))
	tflog.Debug(ctx, fmt.Sprintf("filter error: %#v", err))

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			tflog.Info(ctx, fmt.Sprintf("Filter %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Filter %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Filter read configuration: %#v", filter))

	d.Set("paused", filter.Paused)
	d.Set("description", filter.Description)
	d.Set("expression", filter.Expression)
	d.Set("ref", filter.Ref)

	return nil
}

func resourceCloudflareFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	var newFilter cloudflare.Filter
	newFilter.ID = d.Id()

	if paused, ok := d.GetOk("paused"); ok {
		newFilter.Paused = paused.(bool)
	}

	if description, ok := d.GetOk("description"); ok {
		newFilter.Description = description.(string)
	}

	if expression, ok := d.GetOk("expression"); ok {
		newFilter.Expression = expression.(string)
	}

	if ref, ok := d.GetOk("ref"); ok {
		newFilter.Ref = ref.(string)
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Filter from struct: %+v", newFilter))

	r, err := client.UpdateFilter(ctx, zoneID, newFilter)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Filter for zone %q: %w", zoneID, err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Update response; resource was empty"))
	}

	return resourceCloudflareFilterRead(ctx, d, meta)
}

func resourceCloudflareFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Filter: id %s for zone %s", d.Id(), zoneID))

	err := client.DeleteFilter(ctx, zoneID, d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Filter: %w", err))
	}

	return nil
}

func resourceCloudflareFilterImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/filterID\"", d.Id())
	}

	zoneID, filterID := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Filter: id %s for zone %s", filterID, zoneID))

	d.Set("zone_id", zoneID)
	d.SetId(filterID)

	resourceCloudflareFilterRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
