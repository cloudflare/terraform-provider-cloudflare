---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waiting_room_event"
sidebar_current: "docs-cloudflare-resource-waiting-room-event"
description: |-
  Provides a Cloudflare resource to create and modify a waiting room event.
---

# cloudflare_waiting_room_event

Provides a Cloudflare Waiting Room Event resource.

## Example Usage

```hcl
resource "cloudflare_waiting_room_vent" "example" {
    zone_id              = "ae36f999674d196762efcc5abb06b345"
    waiting_room_id      = "d41d8cd98f00b204e9800998ecf8427e"
    name                 = "foo"
    event_start_time     = "2006-01-02T15:04:05Z"
    event_end_time       = "2006-01-02T20:04:05Z"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The zone ID to apply to.
* `waiting_room_id` - (Required) The Waiting Room ID the event should apply to.
* `name` - (Required) A unique name to identify the event. Only alphanumeric characters, hyphens, and underscores are allowed.
* `event_start_time` (Required) ISO 8601 timestamp that marks the start of the event. At this time, queued users will be processed with the event's configuration. Must occur at least 1 minute before event_end_time.
* `event_end_time` (Required) ISO 8601 timestamp that marks the end of the event.
* `total_active_users` - (Optional) The total number of active user sessions on the route at a point in time.
* `new_users_per_minute` - (Optional) The number of new users that will be let into the route every minute.
* `custom_page_html` - (Optional) This a templated html file that will be rendered at the edge.
* `queueing_method` - (Optional) The queueing method to be used by the waiting room during the event. If not specified, the event will inherit it from the waiting room.
* `shuffle_at_event_start` - (Optional) Users in the prequeue will be shuffled randomly at the `event_start_time`. Requires that `prequeue_start_time` is not null. Default: false.
* `disable_session_renewal` - (Optional) Disables automatic renewal of session cookies. If not specified, the event will inherit it from the waiting room.
* `prequeue_start_time` - (Optional) ISO 8601 timestamp that marks when to begin queueing all users before the event starts. Must occur at least 5 minutes before event_start_time.
* `suspended` - (Optional) If suspended, the traffic doesn't go to the waiting room. Default: false.
* `description` - (Optional) A description to let users add more details about the waiting room event.
* `session_duration` - (Optional) Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the route. Default: 5

## Attributes Reference

The following attributes are exported:

* `id` - The waiting room event ID.

## Import

Waiting room events can be imported using a composite ID formed of zone ID, waiting room ID, and waiting room event ID, e.g.

```
$ terraform import cloudflare_waiting_room_event.default ae36f999674d196762efcc5abb06b345/d41d8cd98f00b204e9800998ecf8427e/25756b2dfe6e378a06b033b670413757
```

where:

* `ae36f999674d196762efcc5abb06b345` - the zone ID
* `d41d8cd98f00b204e9800998ecf8427e` - waiting room ID as returned by [API](https://api.cloudflare.com/#waiting-room-list-waiting-rooms)
* `25756b2dfe6e378a06b033b670413757` - waiting room event ID as returned by [API](https://api.cloudflare.com/#waiting-room-list-events)
