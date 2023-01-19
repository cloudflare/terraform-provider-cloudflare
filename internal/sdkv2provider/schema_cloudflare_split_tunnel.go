package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareSplitTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"mode": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  fmt.Sprintf("The mode of the split tunnel policy. %s", renderAvailableDocumentationValuesStringSlice([]string{"include", "exclude"})),
			ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),
		},
		"tunnels": {
			Required:    true,
			Type:        schema.TypeSet,
			Description: "The value of the tunnel attributes.",
			Elem:        tunnelSetResource,
		},
		"policy_id": {
			Optional:    true,
			Type:        schema.TypeString,
			Description: "The settings policy for which to configure this split tunnel policy.",
		},
	}
}

var tunnelSetResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"address": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The address for the tunnel.",
		},
		"host": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The domain name for the tunnel.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A description for the tunnel.",
		},
	},
}
