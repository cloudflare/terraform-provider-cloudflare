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

func resourceCloudflareHostnameTLSSettingCiphers() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareHostnameTLSSettingCiphersUpdate,
		UpdateContext: resourceCloudflareHostnameTLSSettingCiphersUpdate,
		ReadContext:   resourceCloudflareHostnameTLSSettingCiphersRead,
		DeleteContext: resourceCloudflareHostnameTLSSettingCiphersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareHostnameTLSSettingCiphersImport,
		},

		Schema: resourceCloudflareHostnameTLSSettingCiphersSchema(),
		Description: heredoc.Doc(`
			Provides a Cloudflare per-hostname TLS setting resource, specifically for ciphers suites. Used to set ciphers suites for hostnames under the specified zone.
		`),
	}
}

func resourceCloudflareHostnameTLSSettingCiphersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hostname := d.Get("hostname").(string)

	updateParams := cloudflare.UpdateHostnameTLSSettingCiphersParams{
		Hostname: hostname,
	}
	if value, ok := d.GetOk("value"); ok {
		updateParams.Value = expandInterfaceToStringList(value.([]interface{}))
	}

	_, err := client.UpdateHostnameTLSSettingCiphers(ctx, cloudflare.ZoneIdentifier(zoneID), updateParams)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to update hostname tls setting"))
	}

	d.SetId(stringChecksum(fmt.Sprintf("%s/%s", zoneID, hostname)))
	return resourceCloudflareHostnameTLSSettingCiphersRead(ctx, d, meta)
}

func resourceCloudflareHostnameTLSSettingCiphersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	zoneIDrc := cloudflare.ZoneIdentifier(zoneID)
	hostname := d.Get("hostname").(string)

	listParams := cloudflare.ListHostnameTLSSettingsCiphersParams{
		Hostname: []string{hostname},
	}

	record, resultInfo, err := client.ListHostnameTLSSettingsCiphers(ctx, zoneIDrc, listParams)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			diag.FromErr(fmt.Errorf("hostname tls setting %q not found", hostname))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding hostname tls setting %q: %w", hostname, err))
	}
	if resultInfo.Count == 0 {
		diag.FromErr(fmt.Errorf("hostname tls setting %q not found", hostname))
	}

	d.SetId(stringChecksum(fmt.Sprintf("%s/%s", zoneID, hostname)))
	d.Set("hostname", record[0].Hostname)
	d.Set("value", record[0].Value)
	d.Set("created_at", record[0].CreatedAt.Format(time.RFC3339Nano))
	d.Set("updated_at", record[0].UpdatedAt.Format(time.RFC3339Nano))

	return nil
}

func resourceCloudflareHostnameTLSSettingCiphersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hostname := d.Get("hostname").(string)

	deleteParams := cloudflare.DeleteHostnameTLSSettingCiphersParams{
		Hostname: hostname,
	}

	_, err := client.DeleteHostnameTLSSettingCiphers(ctx, cloudflare.ZoneIdentifier(zoneID), deleteParams)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting hostname tls setting hostname %q in zone %q: %w", hostname, zoneID, err))
	}
	return nil
}

func resourceCloudflareHostnameTLSSettingCiphersImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/hostname\"", d.Id())
	}

	d.SetId(stringChecksum(attributes[0] + "/" + attributes[1]))
	resourceCloudflareHostnameTLSSettingCiphersRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
