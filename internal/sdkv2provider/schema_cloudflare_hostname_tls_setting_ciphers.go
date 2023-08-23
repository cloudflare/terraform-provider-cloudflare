package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareHostnameTLSSettingCiphersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
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
			Type:        schema.TypeList,
			Description: "Ciphers suites value.",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"ports": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeInt},
			Description: "Ports to use within the IP rule.",
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
