package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
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
