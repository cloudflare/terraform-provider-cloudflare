package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAccessCACertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"account_id"},
		},
		"application_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"aud": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
