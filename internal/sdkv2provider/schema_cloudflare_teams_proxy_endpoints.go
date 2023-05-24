package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsProxyEndpointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the teams proxy endpoint.",
		},
		"subdomain": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The FQDN that proxy clients should be pointed at.",
		},
		"ips": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Required:    true,
			Description: "The networks CIDRs that will be allowed to initiate proxy connections.",
		},
	}
}
