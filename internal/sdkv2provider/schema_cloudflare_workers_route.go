package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"pattern": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The [route pattern](https://developers.cloudflare.com/workers/about/routes/) to associate the Worker with.",
		},

		"script_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Worker script name to invoke for requests that match the route pattern.",
		},
	}
}
