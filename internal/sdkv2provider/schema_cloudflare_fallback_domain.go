package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareFallbackDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"domains": {
			Required: true,
			Type:     schema.TypeSet,
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
		"policy_id": {
			Optional:    true,
			Type:        schema.TypeString,
			Description: "The settings policy for which to configure this fallback domain policy.",
		},
	}
}
