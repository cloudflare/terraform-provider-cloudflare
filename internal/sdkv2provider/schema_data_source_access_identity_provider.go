package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessIdentityProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:  "The account identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"zone_id", "account_id"},
		},
		consts.ZoneIDSchemaKey: {
			Description:  "The zone identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"zone_id", "account_id"},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Access Identity Provider name to search for.",
		},
		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Access Identity Provider Type.",
		},
	}
}
