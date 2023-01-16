package sdkv2provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkersKVNamespaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},
		"title": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Title value of the Worker KV Namespace.",
		},
	}
}
