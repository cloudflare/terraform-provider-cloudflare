package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// The supported status for a Zone DNSSEC setting
const (
	DNSSECStatusActive   = "active"
	DNSSECStatusPending  = "pending"
	DNSSECStatusDisabled = "disabled"
)

func resourceCloudflareZoneDNSSEC() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareZoneDNSSECSchema(),
		CreateContext: resourceCloudflareZoneDNSSECCreate,
		ReadContext: resourceCloudflareZoneDNSSECRead,
		UpdateContext: resourceCloudflareZoneDNSSECUpdate,
		DeleteContext: resourceCloudflareZoneDNSSECDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCloudflareZoneDNSSECCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Creating Cloudflare Zone DNSSEC: name %s", zoneID)

	currentDNSSEC, err := client.ZoneDNSSECSetting(context.Background(), zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Zone DNSSEC %q: %s", zoneID, err))
	}
	if currentDNSSEC.Status != DNSSECStatusActive && currentDNSSEC.Status != DNSSECStatusPending {
		_, err := client.UpdateZoneDNSSEC(context.Background(), zoneID, cloudflare.ZoneDNSSECUpdateOptions{Status: DNSSECStatusActive})

		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating zone DNSSEC %q: %s", zoneID, err))
		}

	}

	d.SetId(zoneID)

	return resourceCloudflareZoneDNSSECRead(d, meta)
}

func resourceCloudflareZoneDNSSECRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get("zone_id").(string)

	// In the event zoneID isn't populated at this point, we're likely to be
	// performing an import so set the zoneID to the d.Id() from the passthrough.
	if zoneID == "" {
		zoneID = d.Id()
	}

	dnssec, err := client.ZoneDNSSECSetting(context.Background(), zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Zone DNSSEC %q: %s", zoneID, err))
	}

	if dnssec.Status == DNSSECStatusDisabled {
		return diag.FromErr(fmt.Errorf("zone DNSSEC %q: already disabled", zoneID))
	}

	d.Set("zone_id", zoneID)
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

// Just returning remote state since changing the zone ID would force a new resource
func resourceCloudflareZoneDNSSECUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceCloudflareZoneRead(d, meta)
}

func resourceCloudflareZoneDNSSECDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting Cloudflare Zone DNSSEC: id %s", zoneID)

	_, err := client.UpdateZoneDNSSEC(context.Background(), zoneID, cloudflare.ZoneDNSSECUpdateOptions{Status: DNSSECStatusDisabled})

	if err != nil {
		if strings.Contains(err.Error(), "DNSSEC is already disabled") {
			log.Printf("[INFO] Zone DNSSEC %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Zone DNSSEC: %s", err))
	}

	return nil
}
