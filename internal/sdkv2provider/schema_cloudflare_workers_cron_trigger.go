package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerCronTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"script_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Worker script to target for the schedules.",
		},
		"schedules": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "Cron expressions to execute the Worker script.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
