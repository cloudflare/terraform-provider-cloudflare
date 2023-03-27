package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		
		"script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of worker script to attach the domain to.",
		},
		
		"environment": {
			Type:        schema.TypeString,
			Default: "production",
			Description: "The [route pattern](https://developers.cloudflare.com/workers/about/routes/) to associate the Worker with.",
		},
	}
}