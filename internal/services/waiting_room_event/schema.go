// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r WaitingRoomEventResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"waiting_room_id": schema.StringAttribute{
				Required: true,
			},
			"event_end_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks the end of the event.",
				Required:    true,
			},
			"event_start_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks the start of the event. At this time, queued users will be processed with the event's configuration. The start time must be at least one minute before `event_end_time`.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "A unique name to identify the event. Only alphanumeric characters, hyphens and underscores are allowed.",
				Required:    true,
			},
			"custom_page_html": schema.StringAttribute{
				Description: "If set, the event will override the waiting room's `custom_page_html` property while it is active. If null, the event will inherit it.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A note that you can use to add more details about the event.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"disable_session_renewal": schema.BoolAttribute{
				Description: "If set, the event will override the waiting room's `disable_session_renewal` property while it is active. If null, the event will inherit it.",
				Optional:    true,
			},
			"new_users_per_minute": schema.Int64Attribute{
				Description: "If set, the event will override the waiting room's `new_users_per_minute` property while it is active. If null, the event will inherit it. This can only be set if the event's `total_active_users` property is also set.",
				Optional:    true,
			},
			"prequeue_start_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks when to begin queueing all users before the event starts. The prequeue must start at least five minutes before `event_start_time`.",
				Optional:    true,
			},
			"queueing_method": schema.StringAttribute{
				Description: "If set, the event will override the waiting room's `queueing_method` property while it is active. If null, the event will inherit it.",
				Optional:    true,
			},
			"session_duration": schema.Int64Attribute{
				Description: "If set, the event will override the waiting room's `session_duration` property while it is active. If null, the event will inherit it.",
				Optional:    true,
			},
			"shuffle_at_event_start": schema.BoolAttribute{
				Description: "If enabled, users in the prequeue will be shuffled randomly at the `event_start_time`. Requires that `prequeue_start_time` is not null. This is useful for situations when many users will join the event prequeue at the same time and you want to shuffle them to ensure fairness. Naturally, it makes the most sense to enable this feature when the `queueing_method` during the event respects ordering such as **fifo**, or else the shuffling may be unnecessary.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"suspended": schema.BoolAttribute{
				Description: "Suspends or allows an event. If set to `true`, the event is ignored and traffic will be handled based on the waiting room configuration.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"total_active_users": schema.Int64Attribute{
				Description: "If set, the event will override the waiting room's `total_active_users` property while it is active. If null, the event will inherit it. This can only be set if the event's `new_users_per_minute` property is also set.",
				Optional:    true,
			},
		},
	}
}
