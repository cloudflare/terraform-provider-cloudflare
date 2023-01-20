package sdkv2provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// The supported status for a Zone DNSSEC setting.
const (
	DNSSECStatusActive   = "active"
	DNSSECStatusPending  = "pending"
	DNSSECStatusDisabled = "disabled"
)

func resourceCloudflareZoneDNSSEC() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareZoneDNSSECSchema(),
		CreateContext: resourceCloudflareZoneDNSSECCreate,
		ReadContext:   resourceCloudflareZoneDNSSECRead,
		UpdateContext: resourceCloudflareZoneDNSSECUpdate,
		DeleteContext: resourceCloudflareZoneDNSSECDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: heredoc.Doc(`Provides a Cloudflare resource to create and modify zone DNSSEC settings.`),
	}
}

func resourceCloudflareZoneDNSSECCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare Zone DNSSEC: name %s", zoneID))

	currentDNSSEC, err := client.ZoneDNSSECSetting(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Zone DNSSEC %q: %w", zoneID, err))
	}
	if currentDNSSEC.Status != DNSSECStatusActive && currentDNSSEC.Status != DNSSECStatusPending {
		_, err := client.UpdateZoneDNSSEC(ctx, zoneID, cloudflare.ZoneDNSSECUpdateOptions{Status: DNSSECStatusActive})

		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating zone DNSSEC %q: %w", zoneID, err))
		}
	}

	d.SetId(zoneID)

	return resourceCloudflareZoneDNSSECRead(ctx, d, meta)
}

func resourceCloudflareZoneDNSSECRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	// In the event zoneID isn't populated at this point, we're likely to be
	// performing an import so set the zoneID to the d.Id() from the passthrough.
	if zoneID == "" {
		zoneID = d.Id()
	}

	dnssec, err := client.ZoneDNSSECSetting(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Zone DNSSEC %q: %w", zoneID, err))
	}

	if dnssec.Status == DNSSECStatusDisabled {
		return diag.FromErr(fmt.Errorf("zone DNSSEC %q: already disabled", zoneID))
	}

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("status", dnssec.Status)
	d.Set("flags", dnssec.Flags)
	d.Set("algorithm", dnssec.Algorithm)
	d.Set("key_type", dnssec.KeyType)
	d.Set("digest_type", dnssec.DigestType)
	d.Set("digest_algorithm", dnssec.DigestAlgorithm)
	d.Set("digest", dnssec.Digest)
	d.Set("ds", dnssec.DS)
	d.Set("key_tag", dnssec.KeyTag)
	d.Set("public_key", dnssec.PublicKey)
	d.Set("modified_on", dnssec.ModifiedOn.Format(time.RFC1123Z))

	return nil
}

// Just returning remote state since changing the zone ID would force a new resource.
func resourceCloudflareZoneDNSSECUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceCloudflareZoneRead(ctx, d, meta)
}

func resourceCloudflareZoneDNSSECDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Zone DNSSEC: id %s", zoneID))

	_, err := client.UpdateZoneDNSSEC(ctx, zoneID, cloudflare.ZoneDNSSECUpdateOptions{Status: DNSSECStatusDisabled})

	if err != nil {
		if strings.Contains(err.Error(), "DNSSEC is already disabled") {
			tflog.Info(ctx, fmt.Sprintf("Zone DNSSEC %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Zone DNSSEC: %w", err))
	}

	return nil
}
