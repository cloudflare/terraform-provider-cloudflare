// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WaitingRoomResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"host": schema.StringAttribute{
				Description: "The host name to which the waiting room will be applied (no wildcards). Please do not include the scheme (http:// or https://). The host and path combination must be unique.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "A unique name to identify the waiting room. Only alphanumeric characters, hyphens and underscores are allowed.",
				Required:    true,
			},
			"new_users_per_minute": schema.Int64Attribute{
				Description: "Sets the number of new users that will be let into the route every minute. This value is used as baseline for the number of users that are let in per minute. So it is possible that there is a little more or little less traffic coming to the route based on the traffic patterns at that time around the world.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(200, 2147483647),
				},
			},
			"total_active_users": schema.Int64Attribute{
				Description: "Sets the total number of active user sessions on the route at a point in time. A route is a combination of host and path on which a waiting room is available. This value is used as a baseline for the total number of active user sessions on the route. It is possible to have a situation where there are more or less active users sessions on the route based on the traffic patterns at that time around the world.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(200, 2147483647),
				},
			},
			"cookie_suffix": schema.StringAttribute{
				Description: "Appends a '_' + a custom suffix to the end of Cloudflare Waiting Room's cookie name(__cf_waitingroom). If `cookie_suffix` is \"abcd\", the cookie name will be `__cf_waitingroom_abcd`. This field is required if using `additional_routes`.",
				Optional:    true,
			},
			"custom_page_html": schema.StringAttribute{
				Description: "Only available for the Waiting Room Advanced subscription. This is a template html file that will be rendered at the edge. If no custom_page_html is provided, the default waiting room will be used. The template is based on mustache ( https://mustache.github.io/ ). There are several variables that are evaluated by the Cloudflare edge:\n1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to take when wait time is not available, for instance when queue_all is **true**.\n2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five minutes. Alternatively, you can use:\n3. {{`waitTime`}} Number of minutes of estimated wait for a user.\n4. {{`waitTimeHours`}} Number of hours of estimated wait for a user (`Math.floor(waitTime/60)`).\n5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value (`waitTime%60`).\n6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the queue.\n\nTo view the full list of variables, look at the `cfWaitingRoom` object described under the `json_response_enabled` property in other Waiting Room API calls.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"default_template_language": schema.StringAttribute{
				Description: "The language of the default page template. If no default_template_language is provided, then `en-US` (English) will be used.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"en-US",
						"es-ES",
						"de-DE",
						"fr-FR",
						"it-IT",
						"ja-JP",
						"ko-KR",
						"pt-BR",
						"zh-CN",
						"zh-TW",
						"nl-NL",
						"pl-PL",
						"id-ID",
						"tr-TR",
						"ar-EG",
						"ru-RU",
						"fa-IR",
					),
				},
				Default: stringdefault.StaticString("en-US"),
			},
			"description": schema.StringAttribute{
				Description: "A note that you can use to add more details about the waiting room.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"disable_session_renewal": schema.BoolAttribute{
				Description: "Only available for the Waiting Room Advanced subscription. Disables automatic renewal of session cookies. If `true`, an accepted user will have session_duration minutes to browse the site. After that, they will have to go through the waiting room again. If `false`, a user's session cookie will be automatically renewed on every request.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"json_response_enabled": schema.BoolAttribute{
				Description: "Only available for the Waiting Room Advanced subscription. If `true`, requests to the waiting room with the header `Accept: application/json` will receive a JSON response object with information on the user's status in the waiting room as opposed to the configured static HTML page. This JSON response object has one property `cfWaitingRoom` which is an object containing the following fields:\n1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room (always **true**).\n2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are accurate. If **false**, they are not available.\n3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating the current estimated time in minutes the user will wait in the waiting room. When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.\n4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and `waitTimeKnown` is **true**. Integer indicating the current estimated maximum wait time for the 25% of users that gain entry the fastest (25th percentile).\n5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and `waitTimeKnown` is **true**. Integer indicating the current estimated maximum wait time for the 50% of users that gain entry the fastest (50th percentile). In other words, half of the queued users are expected to let into the origin website before `waitTime50Percentile` and half are expected to be let in after it.\n6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and `waitTimeKnown` is **true**. Integer indicating the current estimated maximum wait time for the 75% of users that gain entry the fastest (75th percentile).\n7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display **unavailable**.\n8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently full and not accepting new users at the moment.\n9. `queueAll`: Boolean indicating if all users will be queued in the waiting room and no one will be let into the origin website.\n10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the user's last attempt to leave the waiting room and be let into the origin website. The user is able to make another attempt after `refreshIntervalSeconds` past this time. If the user makes a request too soon, it will be ignored and `lastUpdated` will not change.\n11. `refreshIntervalSeconds`: Integer indicating the number of seconds after `lastUpdated` until the user is able to make another attempt to leave the waiting room and be let into the origin website. When the `queueingMethod` is `reject`, there is no specified refresh time — it will always be **zero**.\n12. `queueingMethod`: The queueing method currently used by the waiting room. It is either **fifo**, **random**, **passthrough**, or **reject**.\n13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO (First-In-First-Out) queue.\n14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue where users gain access randomly.\n15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a passthrough queue. Keep in mind that when passthrough is enabled, this JSON response will only exist when `queueAll` is **true** or `isEventPrequeueing` is **true** because in all other cases requests will go directly to the origin.\n16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.\n17. `isEventActive`: Boolean indicating if an event is currently occurring. Events are able to change a waiting room's behavior during a specified period of time. For additional information, look at the event properties `prequeue_start_time`, `event_start_time`, and `event_end_time` in the documentation for creating waiting room events. Events are considered active between these start and end times, as well as during the prequeueing period if it exists.\n18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean indicating if an event is currently prequeueing users before it starts.\n19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**. Integer indicating the number of minutes until the event starts.\n20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart` formatted in English for users. If `isEventPrequeueing` is **false**, `timeUntilEventStartFormatted` will display **unavailable**.\n21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer indicating the number of minutes until the event ends.\n22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd` formatted in English for users. If `isEventActive` is **false**, `timeUntilEventEndFormatted` will display **unavailable**.\n23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean indicating if the users in the prequeue are shuffled randomly when the event starts.\n\nAn example cURL to a waiting room could be:\n\n\tcurl -X GET \"https://example.com/waitingroom\" \\\n\t\t-H \"Accept: application/json\"\n\nIf `json_response_enabled` is **true** and the request hits the waiting room, an example JSON response when `queueingMethod` is **fifo** and no event is active could be:\n\n\t{\n\t\t\"cfWaitingRoom\": {\n\t\t\t\"inWaitingRoom\": true,\n\t\t\t\"waitTimeKnown\": true,\n\t\t\t\"waitTime\": 10,\n\t\t\t\"waitTime25Percentile\": 0,\n\t\t\t\"waitTime50Percentile\": 0,\n\t\t\t\"waitTime75Percentile\": 0,\n\t\t\t\"waitTimeFormatted\": \"10 minutes\",\n\t\t\t\"queueIsFull\": false,\n\t\t\t\"queueAll\": false,\n\t\t\t\"lastUpdated\": \"2020-08-03T23:46:00.000Z\",\n\t\t\t\"refreshIntervalSeconds\": 20,\n\t\t\t\"queueingMethod\": \"fifo\",\n\t\t\t\"isFIFOQueue\": true,\n\t\t\t\"isRandomQueue\": false,\n\t\t\t\"isPassthroughQueue\": false,\n\t\t\t\"isRejectQueue\": false,\n\t\t\t\"isEventActive\": false,\n\t\t\t\"isEventPrequeueing\": false,\n\t\t\t\"timeUntilEventStart\": 0,\n\t\t\t\"timeUntilEventStartFormatted\": \"unavailable\",\n\t\t\t\"timeUntilEventEnd\": 0,\n\t\t\t\"timeUntilEventEndFormatted\": \"unavailable\",\n\t\t\t\"shuffleAtEventStart\": false\n\t\t}\n\t}\n\nIf `json_response_enabled` is **true** and the request hits the waiting room, an example JSON response when `queueingMethod` is **random** and an event is active could be:\n\n\t{\n\t\t\"cfWaitingRoom\": {\n\t\t\t\"inWaitingRoom\": true,\n\t\t\t\"waitTimeKnown\": true,\n\t\t\t\"waitTime\": 10,\n\t\t\t\"waitTime25Percentile\": 5,\n\t\t\t\"waitTime50Percentile\": 10,\n\t\t\t\"waitTime75Percentile\": 15,\n\t\t\t\"waitTimeFormatted\": \"5 minutes to 15 minutes\",\n\t\t\t\"queueIsFull\": false,\n\t\t\t\"queueAll\": false,\n\t\t\t\"lastUpdated\": \"2020-08-03T23:46:00.000Z\",\n\t\t\t\"refreshIntervalSeconds\": 20,\n\t\t\t\"queueingMethod\": \"random\",\n\t\t\t\"isFIFOQueue\": false,\n\t\t\t\"isRandomQueue\": true,\n\t\t\t\"isPassthroughQueue\": false,\n\t\t\t\"isRejectQueue\": false,\n\t\t\t\"isEventActive\": true,\n\t\t\t\"isEventPrequeueing\": false,\n\t\t\t\"timeUntilEventStart\": 0,\n\t\t\t\"timeUntilEventStartFormatted\": \"unavailable\",\n\t\t\t\"timeUntilEventEnd\": 15,\n\t\t\t\"timeUntilEventEndFormatted\": \"15 minutes\",\n\t\t\t\"shuffleAtEventStart\": true\n\t\t}\n\t}.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"path": schema.StringAttribute{
				Description: "Sets the path within the host to enable the waiting room on. The waiting room will be enabled for all subpaths as well. If there are two waiting rooms on the same subpath, the waiting room for the most specific path will be chosen. Wildcards and query parameters are not supported.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("/"),
			},
			"queue_all": schema.BoolAttribute{
				Description: "If queue_all is `true`, all the traffic that is coming to a route will be sent to the waiting room. No new traffic can get to the route once this field is set and estimated time will become unavailable.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"queueing_method": schema.StringAttribute{
				Description: "Sets the queueing method used by the waiting room. Changing this parameter from the **default** queueing method is only available for the Waiting Room Advanced subscription. Regardless of the queueing method, if `queue_all` is enabled or an event is prequeueing, users in the waiting room will not be accepted to the origin. These users will always see a waiting room page that refreshes automatically. The valid queueing methods are:\n1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in the order they arrived.\n2. `random`: Random queue where customers gain access randomly, regardless of arrival time.\n3. `passthrough`: Users will pass directly through the waiting room and into the origin website. As a result, any configured limits will not be respected while this is enabled. This method can be used as an alternative to disabling a waiting room (with `suspended`) so that analytics are still reported. This can be used if you wish to allow all traffic normally, but want to restrict traffic during a waiting room event, or vice versa.\n4. `reject`: Users will be immediately rejected from the waiting room. As a result, no users will reach the origin website while this is enabled. This can be used if you wish to reject all traffic while performing maintenance, block traffic during a specified period of time (an event), or block traffic while events are not occurring. Consider a waiting room used for vaccine distribution that only allows traffic during sign-up events, and otherwise blocks all traffic. For this case, the waiting room uses `reject`, and its events override this with `fifo`, `random`, or `passthrough`. When this queueing method is enabled and neither `queueAll` is enabled nor an event is prequeueing, the waiting room page **will not refresh automatically**.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"fifo",
						"random",
						"passthrough",
						"reject",
					),
				},
				Default: stringdefault.StaticString("fifo"),
			},
			"queueing_status_code": schema.Int64Attribute{
				Description: "HTTP status code returned to a user while in the queue.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.OneOf(
						200,
						202,
						429,
					),
				},
				Default: int64default.StaticInt64(200),
			},
			"session_duration": schema.Int64Attribute{
				Description: "Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the route. If a user is not seen by Cloudflare again in that time period, they will be treated as a new user that visits the route.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
				Default: int64default.StaticInt64(5),
			},
			"suspended": schema.BoolAttribute{
				Description: "Suspends or allows traffic going to the waiting room. If set to `true`, the traffic will not go to the waiting room.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"enabled_origin_commands": schema.ListAttribute{
				Description: "A list of enabled origin commands.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("revoke"),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"additional_routes": schema.ListNestedAttribute{
				Description: "Only available for the Waiting Room Advanced subscription. Additional hostname and path combinations to which this waiting room will be applied. There is an implied wildcard at the end of the path. The hostname and path combination must be unique to this and all other waiting rooms.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[WaitingRoomAdditionalRoutesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.StringAttribute{
							Description: "The hostname to which this waiting room will be applied (no wildcards). The hostname must be the primary domain, subdomain, or custom hostname (if using SSL for SaaS) of this zone. Please do not include the scheme (http:// or https://).",
							Optional:    true,
						},
						"path": schema.StringAttribute{
							Description: "Sets the path within the host to enable the waiting room on. The waiting room will be enabled for all subpaths as well. If there are two waiting rooms on the same subpath, the waiting room for the most specific path will be chosen. Wildcards and query parameters are not supported.",
							Computed:    true,
							Optional:    true,
							Default:     stringdefault.StaticString("/"),
						},
					},
				},
			},
			"cookie_attributes": schema.SingleNestedAttribute{
				Description: "Configures cookie attributes for the waiting room cookie. This encrypted cookie stores a user's status in the waiting room, such as queue position.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WaitingRoomCookieAttributesModel](ctx),
				Attributes: map[string]schema.Attribute{
					"samesite": schema.StringAttribute{
						Description: "Configures the SameSite attribute on the waiting room cookie. Value `auto` will be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled. Note that when using value `none`, the secure attribute cannot be set to `never`.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"auto",
								"lax",
								"none",
								"strict",
							),
						},
						Default: stringdefault.StaticString("auto"),
					},
					"secure": schema.StringAttribute{
						Description: "Configures the Secure attribute on the waiting room cookie. Value `always` indicates that the Secure attribute will be set in the Set-Cookie header, `never` indicates that the Secure attribute will not be set, and `auto` will set the Secure attribute depending if **Always Use HTTPS** is enabled.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"auto",
								"always",
								"never",
							),
						},
						Default: stringdefault.StaticString("auto"),
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"next_event_prequeue_start_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks when the next event will begin queueing.",
				Computed:    true,
			},
			"next_event_start_time": schema.StringAttribute{
				Description: "An ISO 8601 timestamp that marks when the next event will start.",
				Computed:    true,
			},
		},
	}
}

func (r *WaitingRoomResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WaitingRoomResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
