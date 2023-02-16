package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareDeviceSettingsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"default": {
			Description: "Whether the policy refers to the default account policy.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"name": {
			Description: "Name of the policy.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"precedence": {
			Description: "The precedence of the policy. Lower values indicate higher precedence.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"match": {
			Description: "Wirefilter expression to match a device against when evaluating whether this policy should take effect for that device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enabled": {
			Description: "Whether the policy is enabled (cannot be set for default policies).",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
		},
		"disable_auto_fallback": {
			Description: "Whether to disable auto fallback for this policy.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"captive_portal": {
			Description: "The captive portal value for this policy.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     180,
		},
		"allow_mode_switch": {
			Description: "Whether to allow mode switch for this policy.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"switch_locked": {
			Description: "Enablement of the ZT client switch lock.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"allow_updates": {
			Description: "Whether to allow updates under this policy.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"auto_connect": {
			Description: "The amount of time in minutes to reconnect after having been disabled.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"allowed_to_leave": {
			Description: "Whether to allow devices to leave the organization.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
		},
		"support_url": {
			Description: "The support URL that will be opened when sending feedback.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"service_mode_v2_mode": {
			Description: "The service mode.",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "warp",
		},
		"service_mode_v2_port": {
			Description:  "The port to use for the proxy service mode.",
			Type:         schema.TypeInt,
			Optional:     true,
			RequiredWith: []string{"service_mode_v2_mode"},
		},
		"exclude_office_ips": {
			Description: "Whether to add Microsoft IPs to split tunnel exclusions.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}
