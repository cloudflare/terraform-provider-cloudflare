package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"host": {
			Type:     schema.TypeString,
			Required: true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"path": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"total_active_users": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"new_users_per_minute": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"custom_page_html": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"default_template_language": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "en-US",
			ValidateFunc: validation.StringInSlice([]string{
				"de-DE",
				"es-ES",
				"en-US",
				"fr-FR",
				"id-ID",
				"it-IT",
				"ja-JP",
				"ko-KR",
				"nl-NL",
				"pl-PL",
				"pt-BR",
				"tr-TR",
				"zh-CN",
				"zh-TW",
			}, false),
		},

		"queue_all": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"disable_session_renewal": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"suspended": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"session_duration": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"json_response_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}
