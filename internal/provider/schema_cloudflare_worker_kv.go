package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkerKVSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"namespace_id": {
			Type:     schema.TypeString,
			ForceNew: true,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
