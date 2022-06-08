package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAccessServiceTokenSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"account_id"},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"client_secret": {
			Type:      schema.TypeString,
			Computed:  true,
			Sensitive: true,
			ForceNew:  true,
		},
		"expires_at": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"min_days_for_renewal": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
}
