package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"hostname": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Hostname of the Worker Domain",
		},

		"service": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of worker script to attach the domain to.",
		},

		"environment": {
			Type:        schema.TypeString,
			Default:     "production",
			Optional:    true,
			Description: "The name of the Worker environment.",
		},
	}
}
