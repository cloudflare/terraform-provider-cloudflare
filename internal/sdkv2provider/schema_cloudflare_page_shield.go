package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePageShieldSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When true, indicates that Page Shield is enabled.",
			Default:     true,
		},
		"use_cloudflare_reporting_endpoint": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When true, CSP reports will be sent to https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report",
		},
		"use_connection_url_path": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "When true, the paths associated with connections URLs will also be analyzed.",
		},
	}
}
