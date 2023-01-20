package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareWAFGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"package_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "on",
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		},
	}
}
