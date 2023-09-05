package sdkv2provider

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var defaultTemplateLanguages = []string{
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
	"ru-RU",
	"fa-IR",
}
var waitingRoomQueueingMethod = []string{
	"fifo",
	"random",
	"passthrough",
	"reject",
}

func resourceCloudflareWaitingRoomSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"name": {
			Description: "A unique name to identify the waiting room.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"host": {
			Description: "Host name for which the waiting room will be applied (no wildcards)",
			Type:        schema.TypeString,
			Required:    true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"path": {
			Description: "The path within the host to enable the waiting room on.",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "/",
		},

		"total_active_users": {
			Description: "The total number of active user sessions on the route at a point in time.",
			Type:        schema.TypeInt,
			Required:    true,
		},

		"new_users_per_minute": {
			Description: "The number of new users that will be let into the route every minute.",
			Type:        schema.TypeInt,
			Required:    true,
		},

		"custom_page_html": {
			Description: "This is a templated html file that will be rendered at the edge.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"queueing_method": {
			Description:  fmt.Sprintf("The queueing method used by the waiting room. %s", renderAvailableDocumentationValuesStringSlice(waitingRoomQueueingMethod)),
			Type:         schema.TypeString,
			Default:      "fifo",
			Optional:     true,
			ValidateFunc: validation.StringInSlice(waitingRoomQueueingMethod, false),
		},

		"default_template_language": {
			Description:  fmt.Sprintf("The language to use for the default waiting room page. %s", renderAvailableDocumentationValuesStringSlice(defaultTemplateLanguages)),
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "en-US",
			ValidateFunc: validation.StringInSlice(defaultTemplateLanguages, false),
		},

		"queue_all": {
			Description: "If queue_all is true, then all traffic will be sent to the waiting room.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"disable_session_renewal": {
			Description: "Disables automatic renewal of session cookies.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"suspended": {
			Description: "Suspends the waiting room.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"description": {
			Description: "A description to add more details about the waiting room.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"session_duration": {
			Description: "Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the origin.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     5,
		},

		"json_response_enabled": {
			Description: "If true, requests to the waiting room with the header `Accept: application/json` will receive a JSON response object.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"cookie_suffix": {
			Description: "A cookie suffix to be appended to the Cloudflare waiting room cookie name.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"additional_routes": {
			Description: "A list of additional hostname and paths combination to be applied on the waiting room.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"host": {
						Description: "The additional host name for which the waiting room to be applied on (no wildcards).",
						Required:    true,
						Type:        schema.TypeString,
						StateFunc: func(i interface{}) string {
							return strings.ToLower(i.(string))
						},
					},
					"path": {
						Description: "The path within the additional host to enable the waiting room on.",
						Optional:    true,
						Default:     "/",
						Type:        schema.TypeString,
					},
				},
			},
		},

		"queueing_status_code": {
			Description: "HTTP status code returned to a user while in the queue.",
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     200,
		},
	}
}
