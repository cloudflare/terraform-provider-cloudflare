package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareHostnameTLSSettingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"setting": {
			Type:        schema.TypeString,
			Description: "TLS setting name.",
			Required:    true,
			ForceNew:    true,
		},
		"hostname": {
			Type:        schema.TypeString,
			Description: "Hostname that belongs to this zone name.",
			Required:    true,
			ForceNew:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "TLS setting value.",
			Required:    true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
