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

func dataSourceCloudflareZoneDNSSEC() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareZoneDNSSECRead,

		Schema: map[string]*schema.Schema{
			consts.ZoneIDSchemaKey: {
				Description: "The zone identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the Zone DNSSEC.",
			},
			"flags": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Zone DNSSEC flags.",
			},
			"algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone DNSSEC algorithm.",
			},
			"key_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key type used for Zone DNSSEC.",
			},
			"digest_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Digest Type for Zone DNSSEC.",
			},
			"digest_algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Digest algorithm use for Zone DNSSEC.",
			},
			"digest": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone DNSSEC digest.",
			},
			"ds": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DS for the Zone DNSSEC.",
			},
			"key_tag": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Key Tag for the Zone DNSSEC.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public Key for the Zone DNSSEC.",
			},
		},
		Description: "Use this data source to look up Zone DNSSEC settings.",
	}
}

func dataSourceCloudflareZoneDNSSECRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Reading Zone DNSSEC %s", zoneID))

	dnssec, err := client.ZoneDNSSECSetting(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding Zone DNSSEC %q: %w", zoneID, err))
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

	d.SetId(stringChecksum(dnssec.ModifiedOn.String()))

	return nil
}
