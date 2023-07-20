package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAuthenticatedOriginPullsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify a hostname to enable Per-Hostname Authenticated Origin Pulls on, using the provided certificate.",
		},
		"authenticated_origin_pulls_certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The ID of an uploaded Authenticated Origin Pulls certificate. If no hostname is provided, this certificate will be used zone wide as Per-Zone Authenticated Origin Pulls.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Whether to enable Authenticated Origin Pulls on the given zone or hostname.",
		},
	}
}
