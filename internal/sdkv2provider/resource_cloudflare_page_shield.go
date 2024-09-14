package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePageShield() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflarePageShieldSchema(),
		CreateContext: resourceCloudflarePageShieldCreate,
		ReadContext:   resourceCloudflarePageShieldRead,
		UpdateContext: resourceCloudflarePageShieldCreate,
		DeleteContext: resourceCloudflarePageShieldDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflarePageShieldImport,
		},
		Description: "Provides the ability to manage Page Shield.",
	}
}

func resourceCloudflarePageShieldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	newPageShield := cloudflare.UpdatePageShieldSettingsParams{
		Enabled:                        cloudflare.BoolPtr(d.Get("enabled").(bool)),
		UseCloudflareReportingEndpoint: cloudflare.BoolPtr(d.Get("use_cloudflare_reporting_endpoint").(bool)),
		UseConnectionURLPath:           cloudflare.BoolPtr(d.Get("use_connection_url_path").(bool)),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Page Shield from struct: %+v", newPageShield))

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	_, err := client.UpdatePageShieldSettings(ctx, cloudflare.ZoneIdentifier(zoneID), newPageShield)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(zoneID)
	return resourceCloudflareAccessRuleRead(ctx, d, meta)
}

func resourceCloudflarePageShieldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	response, err := client.GetPageShieldSettings(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.GetPageShieldSettingsParams{})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(zoneID)
	d.Set("enabled", response.PageShield.Enabled)
	d.Set("use_cloudflare_reporting_endpoint", response.PageShield.UseCloudflareReportingEndpoint)
	d.Set("use_connection_url_path", response.PageShield.UseConnectionURLPath)

	return nil
}

func resourceCloudflarePageShieldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	tflog.Info(ctx, "Disabling Page Shield")

	// There is no DELETE endpoint for schema validation settings,
	// so terraform should reset the state to default settings
	params := cloudflare.UpdatePageShieldSettingsParams{
		Enabled:                        cloudflare.BoolPtr(false),
		UseCloudflareReportingEndpoint: cloudflare.BoolPtr(true),
		UseConnectionURLPath:           cloudflare.BoolPtr(false),
	}

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	_, err := client.UpdatePageShieldSettings(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudflarePageShieldImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set(consts.ZoneIDSchemaKey, d.Id())

	resourceCloudflarePageShieldRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
