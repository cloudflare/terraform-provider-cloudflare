package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerSecretSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"script_name": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "The name of the Worker script to associate the secret with.",
		},
		"name": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "The name of the Worker secret.",
		},
		"secret_text": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			ForceNew:    true,
			Description: "The text of the Worker secret.",
		},
	}
}
