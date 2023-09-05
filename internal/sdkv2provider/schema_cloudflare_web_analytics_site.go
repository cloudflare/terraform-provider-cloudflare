package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWebAnalyticsSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"site_tag": {
			Description: "The Web Analytics site tag.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"host": {
			Description:  "The hostname to use for gray-clouded sites.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"zone_tag"},
			ForceNew:     true,
		},

		"zone_tag": {
			Description:  "The zone identifier for orange-clouded sites.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"host"},
			ForceNew:     true,
		},

		"site_token": {
			Description: "The token for the Web Analytics site.",
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
		},

		"snippet": {
			Description: "The encoded JS snippet to add to your site's HTML page if auto_install is false.",
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
		},

		"auto_install": {
			Description: "Whether Cloudflare will automatically inject the JavaScript snippet for orange-clouded sites.",
			Type:        schema.TypeBool,
			Required:    true,
			ForceNew:    true,
		},

		"ruleset_id": {
			Description: "The ID for the ruleset associated to this Web Analytics Site.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
