package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessIdentityProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"zone_id", "account_id"},
		},
		"zone_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"zone_id", "account_id"},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
