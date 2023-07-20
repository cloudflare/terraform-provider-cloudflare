package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDevicePolicyCertificates() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDevicePolicyCertificatesSchema(),
		CreateContext: resourceCloudflareDevicePolicyCertificateUpdate,
		ReadContext:   resourceCloudflareDevicePolicyCertificateRead,
		UpdateContext: resourceCloudflareDevicePolicyCertificateUpdate,
		DeleteContext: resourceCloudflareDevicePolicyCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDevicePolicyCertificateImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare device policy certificates resource. Device
			policy certificate resources enable client device certificate
			generation.
		`),
	}
}

func resourceCloudflareDevicePolicyCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	enabled := d.Get("enabled").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare device policy certificate: zoneID=%s enabled=%t", zoneID, enabled))

	_, err := client.UpdateDeviceClientCertificatesZone(ctx, zoneID, enabled)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Cloudflare device policy certificate %q: %w", zoneID, err))
	}

	d.SetId(zoneID)
	return nil
}

func resourceCloudflareDevicePolicyCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	enabled, err := client.GetDeviceClientCertificatesZone(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading device policy certificate setting %q: %w", zoneID, err))
	}

	d.SetId(zoneID)
	d.Set("enabled", enabled.Result.Enabled)
	return nil
}

func resourceCloudflareDevicePolicyCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID := d.Id()

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare device policy certificate setting: zoneID=%s", zoneID))

	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)

	resourceCloudflareDevicePolicyCertificateRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareDevicePolicyCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
