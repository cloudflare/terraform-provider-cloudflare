package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareArgoTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"secret": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			ForceNew:  true,
		},
		"cname": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
