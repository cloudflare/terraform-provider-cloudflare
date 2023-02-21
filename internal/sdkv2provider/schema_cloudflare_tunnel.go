package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "A user-friendly name chosen when the tunnel is created.",
		},
		"secret": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			ForceNew:    true,
			Description: "32 or more bytes, encoded as a base64 string. The Create Argo Tunnel endpoint sets this as the tunnel's password. Anyone wishing to run the tunnel needs this password.",
		},
		"cname": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Usable CNAME for accessing the Tunnel.",
		},
		"tunnel_token": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Token used by a connector to authenticate and run the tunnel.",
			Sensitive:   true,
		},
	}
}
