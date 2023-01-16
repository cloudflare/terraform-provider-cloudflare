package sdkv2provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkerRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
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
