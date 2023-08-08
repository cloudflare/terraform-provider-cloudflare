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
			Description: "If true, the zone hold will be enabled. If false, the zone hold will be disabled.",
		},
		"include_subdomains": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If provided, the zone hold will extend to block any subdomain of the given zone, as well as SSL4SaaS Custom Hostnames.",
		},
		"hold_after": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "If provided, the hold will be temporarily disabled, then automatically re-enabled by the system at the time specified in this RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.",
		},
	}
}
