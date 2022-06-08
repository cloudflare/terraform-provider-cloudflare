package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareZoneDNSSEC() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareZoneDNSSECRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Description: "The zone identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flags": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"digest_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"digest_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"digest": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ds": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_tag": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudflareZoneDNSSECRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get("zone_id").(string)

	tflog.Debug(ctx, fmt.Sprintf("Reading Zone DNSSEC %s", zoneID))

	dnssec, err := client.ZoneDNSSECSetting(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Zone DNSSEC %q: %w", zoneID, err))
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

	d.SetId(stringChecksum(dnssec.ModifiedOn.String()))

	return nil
}
