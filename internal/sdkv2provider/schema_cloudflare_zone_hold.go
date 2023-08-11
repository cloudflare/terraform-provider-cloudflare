package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneHoldSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"hold": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Enablement status of the zone hold.",
		},
		"include_subdomains": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to extend to block any subdomain of the given zone.",
		},
		"hold_after": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The RFC3339 compatible timestamp when to automatically re-enable the zone hold.",
		},
	}
}
