package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareSplitTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"mode": {
			Type:         schema.TypeString,
			Required:     true,
			Description:  "The mode of the split tunnel policy. Either 'include' or 'exclude'.",
			ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),
		},
		"tunnels": {
			Required: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
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
			},
		},
	}
}
