package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		Description: heredoc.Doc(`
			Provides a Cloudflare Zone Lockdown resource. Zone Lockdown allows
			you to define one or more URLs (with wildcard matching on the domain
			or path) that will only permit access if the request originates
			from an IP address that matches a safelist of one or more IP
			addresses and/or IP ranges.
		`),
	}
}

func resourceCloudflareZoneLockdownCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var err error

	var newZoneLockdown cloudflare.ZoneLockdownCreateParams

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

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Zone Lockdown from struct: %+v", newZoneLockdown))

	var r cloudflare.ZoneLockdown

	r, err = client.CreateZoneLockdown(ctx, cloudflare.ZoneIdentifier(zoneID), newZoneLockdown)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating zone lockdown for zone ID %q: %w", zoneID, err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Zone Lockdown ID: %s", d.Id()))

	return resourceCloudflareZoneLockdownRead(ctx, d, meta)
}

func resourceCloudflareZoneLockdownRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	zoneLockdownResponse, err := client.ZoneLockdown(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())

	tflog.Debug(ctx, fmt.Sprintf("zoneLockdownResponse: %#v", zoneLockdownResponse))
	tflog.Debug(ctx, fmt.Sprintf("zoneLockdownResponse error: %#v", err))

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Zone Lockdown %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding zone lockdown %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Zone Lockdown read configuration: %#v", zoneLockdownResponse))

	d.Set("paused", zoneLockdownResponse.Paused)
	d.Set("priority", zoneLockdownResponse.Priority)
	d.Set("description", zoneLockdownResponse.Description)
	d.Set("urls", zoneLockdownResponse.URLs)
	tflog.Debug(ctx, fmt.Sprintf("read configurations: %#v", d.Get("configurations")))

	configurations := make([]map[string]interface{}, len(zoneLockdownResponse.Configurations))

	for i, entryconfigZoneLockdownConfig := range zoneLockdownResponse.Configurations {
		configurations[i] = map[string]interface{}{
			"target": entryconfigZoneLockdownConfig.Target,
			"value":  entryconfigZoneLockdownConfig.Value,
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Zone Lockdown configuration: %#v", configurations))

	if err := d.Set("configurations", configurations); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Error setting configurations in zone lockdown %q: %s", d.Id(), err))
	}

	return nil
}

func resourceCloudflareZoneLockdownUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var newZoneLockdown cloudflare.ZoneLockdownUpdateParams

	newZoneLockdown.ID = d.Id()
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

	tflog.Info(ctx, fmt.Sprintf("Updating Cloudflare Zone Lockdown from struct: %+v", newZoneLockdown))

	r, err := client.UpdateZoneLockdown(ctx, cloudflare.ZoneIdentifier(zoneID), newZoneLockdown)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating zone lockdown for zone %q: %w", zoneID, err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Update response; resource was empty"))
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Zone Lockdown ID: %s", d.Id()))

	return resourceCloudflareZoneLockdownRead(ctx, d, meta)
}

func resourceCloudflareZoneLockdownDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Zone Lockdown: id %s for zone %s", d.Id(), zoneID))

	_, err := client.DeleteZoneLockdown(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Zone Lockdown: %w", err))
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
		d.Set(consts.ZoneIDSchemaKey, zoneID)
		d.SetId(zoneLockdownID)
	} else {
		return nil, fmt.Errorf("invalid id (%q) specified, should be in format \"zoneID/zoneLockdownId\"", d.Id())
	}

	tflog.Debug(ctx, fmt.Sprintf("zoneID: %s", zoneID))
	tflog.Debug(ctx, fmt.Sprintf("Resource ID : %s", zoneLockdownID))

	resourceCloudflareZoneLockdownRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
