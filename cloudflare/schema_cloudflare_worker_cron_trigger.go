package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkerCronTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"script_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"schedules": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
