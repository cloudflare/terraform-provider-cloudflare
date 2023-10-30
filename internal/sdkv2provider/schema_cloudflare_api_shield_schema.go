package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// API Shield Schema Terraform Schema.
func resourceCloudflareAPIShieldSchemaSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the schema",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"kind": {
			Description: "Kind of schema",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Default:     "openapi_v3",
		},
		"source": {
			Description: "Schema file bytes",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"validation_enabled": {
			Description: "Flag whether schema is enabled for validation",
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}
