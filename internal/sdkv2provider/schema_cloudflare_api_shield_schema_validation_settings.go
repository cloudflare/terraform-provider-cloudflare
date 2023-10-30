package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAPIShieldSchemaValidationSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"validation_default_mitigation_action": {
			Description: "The default mitigation action used when there is no mitigation action defined on the operation",
			Type:        schema.TypeString,
			Required:    true,
		},
		"validation_override_mitigation_action": {
			Description: "When set, this overrides both zone level and operation level mitigation actions",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
