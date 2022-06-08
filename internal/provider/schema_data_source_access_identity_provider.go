package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessIdentityProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:  "The account identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"zone_id", "account_id"},
		},
		"zone_id": {
			Description:  "The zone identifier to target for the resource.",
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
