package sdkv2provider

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomEventSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"waiting_room_id": {
			Description: "The Waiting Room ID the event should apply to.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"name": {
			Description: "A unique name to identify the event. Only alphanumeric characters, hyphens, and underscores are allowed.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"event_start_time": {
			Description: "ISO 8601 timestamp that marks the start of the event. Must occur at least 1 minute before `event_end_time`.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"event_end_time": {
			Description: "ISO 8601 timestamp that marks the end of the event.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"total_active_users": {
			Description: "The total number of active user sessions on the route at a point in time.",
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"new_users_per_minute": {
			Description: "The number of new users that will be let into the route every minute.",
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"custom_page_html": {
			Description: "This is a templated html file that will be rendered at the edge.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"queueing_method": {
			Description:  fmt.Sprintf("The queueing method used by the waiting room. %s", renderAvailableDocumentationValuesStringSlice(waitingRoomQueueingMethod)),
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(waitingRoomQueueingMethod, false),
		},

		"shuffle_at_event_start": {
			Description: "Users in the prequeue will be shuffled randomly at the `event_start_time`. Requires that `prequeue_start_time` is not null.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},

		"disable_session_renewal": {
			Description: "Disables automatic renewal of session cookies.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"prequeue_start_time": {
			Description: "ISO 8601 timestamp that marks when to begin queueing all users before the event starts. Must occur at least 5 minutes before `event_start_time`.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"suspended": {
			Description: "If suspended, the event is ignored and traffic will be handled based on the waiting room configuration.",
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"description": {
			Description: "A description to let users add more details about the event.",
			Type:        schema.TypeString,
			Optional:    true,
		},

		"session_duration": {
			Description: "Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the origin.",
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"created_on": {
			Description: "Creation time.",
			Type:        schema.TypeString,
			Computed:    true,
		},

		"modified_on": {
			Description: "Last modified time.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
