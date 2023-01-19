package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneDNSSECSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
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
		"modified_on": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Zone DNSSEC updated time.",
		},
	}
}
