package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccessIdentityProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:  consts.AccountIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{consts.ZoneIDSchemaKey, consts.AccountIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:  consts.ZoneIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{consts.ZoneIDSchemaKey, consts.AccountIDSchemaKey},
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
