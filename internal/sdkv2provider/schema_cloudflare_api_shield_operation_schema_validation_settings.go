package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAPIShieldOperationSchemaValidationSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"operation_id": {
			Description: "Operation ID these settings should apply to",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"mitigation_action": {
			Description: "The mitigation action to apply to this operation",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
