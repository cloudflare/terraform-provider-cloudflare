package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareCustomHostnameFallbackOriginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
		},
		"origin": {
			Type:     schema.TypeString,
			Required: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
