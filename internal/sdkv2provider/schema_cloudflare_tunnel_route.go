package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"tunnel_id": {
			Description: "The ID of the tunnel that will service the tunnel route.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"network": {
			Description: "The IPv4 or IPv6 network that should use this tunnel route, in CIDR notation.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"comment": {
			Description: "Description of the tunnel route.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"virtual_network_id": {
			Description: "The ID of the virtual network for which this route is being added; uses the default virtual network of the account if none is provided.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
	}
}
