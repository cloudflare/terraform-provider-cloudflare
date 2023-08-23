package sdkv2provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareHostnameTLSSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareHostnameTLSSettingUpdate,
		UpdateContext: resourceCloudflareHostnameTLSSettingUpdate,
		ReadContext:   resourceCloudflareHostnameTLSSettingRead,
		DeleteContext: resourceCloudflareHostnameTLSSettingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareHostnameTLSSettingImport,
		},

		Schema: resourceCloudflareHostnameTLSSettingSchema(),
		Description: heredoc.Doc(`
			Provides a Cloudflare per-hostname TLS setting resource. Used to set TLS settings for hostnames under the specified zone.
		`),
	}
}

func resourceCloudflareHostnameTLSSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	setting := d.Get("setting").(string)
	hostname := d.Get("hostname").(string)

	updateParams := cloudflare.UpdateHostnameTLSSettingParams{
		Setting:  setting,
		Hostname: hostname,
		Value:    d.Get("value").(string),
	}

	_, err := client.UpdateHostnameTLSSetting(ctx, cloudflare.ZoneIdentifier(zoneID), updateParams)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to update hostname tls setting"))
	}

	d.SetId(stringChecksum(zoneID + "/" + hostname + "/" + setting))
	return resourceCloudflareHostnameTLSSettingRead(ctx, d, meta)
}

func resourceCloudflareHostnameTLSSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	zoneIDrc := cloudflare.ZoneIdentifier(zoneID)
	hostname := d.Get("hostname").(string)
	setting := d.Get("setting").(string)

	listParams := cloudflare.ListHostnameTLSSettingsParams{
		Setting:  setting,
		Hostname: []string{hostname},
	}

	record, resultInfo, err := client.ListHostnameTLSSettings(ctx, zoneIDrc, listParams)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			diag.FromErr(fmt.Errorf("hostname tls setting %q not found", hostname))
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding hostname tls setting %q: %w", hostname, err))
	}

	if resultInfo.Count == 0 {
		diag.FromErr(fmt.Errorf("hostname tls setting %q not found", hostname))
	}

	d.SetId(stringChecksum(zoneID + "/" + hostname + "/" + setting))
	d.Set("hostname", record[0].Hostname)
	d.Set("value", record[0].Value)
	d.Set("created_at", record[0].CreatedAt.Format(time.RFC3339Nano))
	d.Set("updated_at", record[0].UpdatedAt.Format(time.RFC3339Nano))

	return nil
}

func resourceCloudflareHostnameTLSSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	zoneIDrc := cloudflare.ZoneIdentifier(zoneID)
	hostname := d.Get("hostname").(string)

	deleteParams := cloudflare.DeleteHostnameTLSSettingParams{
		Setting:  d.Get("setting").(string),
		Hostname: hostname,
	}

	_, err := client.DeleteHostnameTLSSetting(ctx, zoneIDrc, deleteParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting hostname tls setting hostname %q in zone %q: %w", hostname, zoneID, err))
	}
	return nil
}

func resourceCloudflareHostnameTLSSettingImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/hostname/setting\"", d.Id())
	}

	d.SetId(stringChecksum(attributes[0] + "/" + attributes[1] + "/" + attributes[2]))

	resourceCloudflareHostnameTLSSettingRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
