package cloudflare

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomEventSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"waiting_room_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"event_start_time": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"event_end_time": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"total_active_users": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"new_users_per_minute": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"custom_page_html": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"queueing_method": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"shuffle_at_event_start": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},

		"disable_session_renewal": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"prequeue_start_time": {
			Type:     schema.TypeString,
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

		"created_on": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
