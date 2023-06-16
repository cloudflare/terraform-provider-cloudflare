package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
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
		"config_src": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"local", "cloudflare"}, false),
			Description:  fmt.Sprintf("Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel on the Zero Trust dashboard or using tunnel_config, tunnel_route or tunnel_virtual_network resources. %s", renderAvailableDocumentationValuesStringSlice([]string{"local", "cloudflare"})),
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
