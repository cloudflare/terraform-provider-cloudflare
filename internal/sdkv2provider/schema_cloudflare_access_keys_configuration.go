package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessKeysConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"key_rotation_interval_days": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Number of days to trigger a rotation of the keys.",
		},
	}
}
