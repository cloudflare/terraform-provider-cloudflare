package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkerKVSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "Name of the KV pair.",
		},
		"namespace_id": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "The ID of the Workers KV namespace in which you want to create the KV pair.",
		},
		"value": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Value of the KV pair.",
		},
	}
}
