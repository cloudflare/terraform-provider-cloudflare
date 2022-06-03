package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelVirtualNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"is_default_network": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"comment": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
