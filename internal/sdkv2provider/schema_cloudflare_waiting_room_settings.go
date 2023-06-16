package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"search_engine_crawler_bypass": {
			Description: "Whether to allow verified search engine crawlers to bypass all waiting rooms on this zone.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
	}
}
