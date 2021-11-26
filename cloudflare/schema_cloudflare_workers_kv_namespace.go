package cloudflare

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWorkersKVNamespaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"title": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
