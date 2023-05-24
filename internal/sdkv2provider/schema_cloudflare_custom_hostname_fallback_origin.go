package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareCustomHostnameFallbackOriginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
		},
		"origin": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Status of the fallback origin's activation.",
		},
	}
}
