package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFallbackDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domains": {
			Required: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"suffix": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The domain suffix to match when resolving locally.",
					},
					"description": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A description of the fallback domain, displayed in the client UI.",
					},
					"dns_server": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "A list of IP addresses to handle domain resolution.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}
