package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneLockdown() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneLockdownSchema(),
		CreateContext: resourceCloudflareZoneLockdownCreate,
		ReadContext:   resourceCloudflareZoneLockdownRead,
		UpdateContext: resourceCloudflareZoneLockdownUpdate,
		DeleteContext: resourceCloudflareZoneLockdownDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareZoneLockdownImport,
		},
	}
}

func resourceCloudflareZoneLockdownCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	var err error

	var newZoneLockdown cloudflare.ZoneLockdown

	if paused, ok := d.GetOk("paused"); ok {
		newZoneLockdown.Paused = paused.(bool)
	}

	if priority, ok := d.GetOk("priority"); ok {
		newZoneLockdown.Priority = priority.(int)
	}

	if description, ok := d.GetOk("description"); ok {
		newZoneLockdown.Description = description.(string)
	}

	if urls, ok := d.GetOk("urls"); ok {
		newZoneLockdown.URLs = expandInterfaceToStringList(urls.(*schema.Set).List())
	}

	if configurations, ok := d.GetOk("configurations"); ok {
		newZoneLockdown.Configurations = expandZoneLockdownConfig(configurations.(*schema.Set))
	}

	log.Printf("[DEBUG] Creating Cloudflare Zone Lockdown from struct: %+v", newZoneLockdown)

	var r *cloudflare.ZoneLockdownResponse

	r, err = client.CreateZoneLockdown(ctx, zoneID, newZoneLockdown)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating zone lockdown for zone ID %q: %s", zoneID, err))
	}

	if r.Result.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.Result.ID)

	log.Printf("[INFO] Cloudflare Zone Lockdown ID: %s", d.Id())

	return resourceCloudflareZoneLockdownRead(ctx, d, meta)
}

func resourceCloudflareZoneLockdownRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	zoneLockdownResponse, err := client.ZoneLockdown(ctx, zoneID, d.Id())

	log.Printf("[DEBUG] zoneLockdownResponse: %#v", zoneLockdownResponse)
	log.Printf("[DEBUG] zoneLockdownResponse error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Zone Lockdown %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding zone lockdown %q: %s", d.Id(), err))
	}

	log.Printf("[DEBUG] Cloudflare Zone Lockdown read configuration: %#v", zoneLockdownResponse)

	d.Set("paused", zoneLockdownResponse.Result.Paused)
	d.Set("priority", zoneLockdownResponse.Result.Priority)
	d.Set("description", zoneLockdownResponse.Result.Description)
	d.Set("urls", zoneLockdownResponse.Result.URLs)
	log.Printf("[DEBUG] read configurations: %#v", d.Get("configurations"))

	configurations := make([]map[string]interface{}, len(zoneLockdownResponse.Result.Configurations))

	for i, entryconfigZoneLockdownConfig := range zoneLockdownResponse.Result.Configurations {
		configurations[i] = map[string]interface{}{
			"target": entryconfigZoneLockdownConfig.Target,
			"value":  entryconfigZoneLockdownConfig.Value,
		}
	}
	log.Printf("[DEBUG] Cloudflare Zone Lockdown configuration: %#v", configurations)

	if err := d.Set("configurations", configurations); err != nil {
		log.Printf("[WARN] Error setting configurations in zone lockdown %q: %s", d.Id(), err)
	}

	return nil
}

func resourceCloudflareZoneLockdownUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	var newZoneLockdown cloudflare.ZoneLockdown

	if paused, ok := d.GetOk("paused"); ok {
		newZoneLockdown.Paused = paused.(bool)
	}

	if priority, ok := d.GetOk("priority"); ok {
		newZoneLockdown.Priority = priority.(int)
	}

	if description, ok := d.GetOk("description"); ok {
		newZoneLockdown.Description = description.(string)
	}

	if urls, ok := d.GetOk("urls"); ok {
		newZoneLockdown.URLs = expandInterfaceToStringList(urls.(*schema.Set).List())
	}

	if configurations, ok := d.GetOk("configurations"); ok {
		newZoneLockdown.Configurations = expandZoneLockdownConfig(configurations.(*schema.Set))
	}

	log.Printf("[INFO] Updating Cloudflare Zone Lockdown from struct: %+v", newZoneLockdown)

	r, err := client.UpdateZoneLockdown(ctx, zoneID, d.Id(), newZoneLockdown)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating zone lockdown for zone %q: %s", d.Get("zone").(string), err))
	}

	if r.Result.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Update response; resource was empty"))
	}

	d.SetId(r.Result.ID)

	log.Printf("[INFO] Cloudflare Zone Lockdown ID: %s", d.Id())

	return resourceCloudflareZoneLockdownRead(ctx, d, meta)
}

func resourceCloudflareZoneLockdownDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting Cloudflare Zone Lockdown: id %s for zone %s", d.Id(), zoneID)

	_, err := client.DeleteZoneLockdown(ctx, zoneID, d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Zone Lockdown: %s", err))
	}

	return nil
}

func expandZoneLockdownConfig(configs *schema.Set) []cloudflare.ZoneLockdownConfig {
	configArray := make([]cloudflare.ZoneLockdownConfig, configs.Len())
	for i, entry := range configs.List() {
		e := entry.(map[string]interface{})
		configArray[i] = cloudflare.ZoneLockdownConfig{
			Target: e["target"].(string),
			Value:  e["value"].(string),
		}
	}
	return configArray
}

func resourceCloudflareZoneLockdownImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var zoneLockdownID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		zoneLockdownID = idAttr[1]
		d.Set("zone_id", zoneID)
		d.SetId(zoneLockdownID)
	} else {
		return nil, fmt.Errorf("invalid id (%q) specified, should be in format \"zoneID/zoneLockdownId\"", d.Id())
	}

	log.Printf("[DEBUG] zoneID: %s", zoneID)
	log.Printf("[DEBUG] Resource ID : %s", zoneLockdownID)

	resourceCloudflareZoneLockdownRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
