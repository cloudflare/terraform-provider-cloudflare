package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareZoneDNSSECSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
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
		"modified_on": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}
