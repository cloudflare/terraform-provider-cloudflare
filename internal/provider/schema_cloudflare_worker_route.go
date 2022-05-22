package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkerRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"pattern": {
			Type:     schema.TypeString,
			Required: true,
		},

		"script_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
