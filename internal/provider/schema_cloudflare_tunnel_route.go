package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tunnel_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"network": {
			Type:     schema.TypeString,
			Required: true,
		},
		"comment": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
