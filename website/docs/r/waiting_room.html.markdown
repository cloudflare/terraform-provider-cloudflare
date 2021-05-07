---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waiting_room"
sidebar_current: "docs-cloudflare-resource-waiting-room"
description: |-
  Provides a Cloudflare resource to create and modify a waiting room.
---

# cloudflare_waiting_room

Provides a Cloudflare Waiting Room resource.

## Example Usage

```hcl
resource "cloudflare_waiting_room" "example" {
    zone_id              = "ae36f999674d196762efcc5abb06b345"
    name                 = "foo"
    host                 = "foo.bar.com"
    path                 = "/"
    new_users_per_minute = 200
    total_active_users   = 200
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to apply to.
* `name` - (Required) A unique name to identify the waiting room.
* `host` - (Required) Host name for which the waiting room will be applied (no wildcards).
* `path` - (Required) The path within the host to enable the waiting room on. Default: "/".
* `total_active_users` - (Required) The total number of active user sessions on the route at a point in time.
* `new_users_per_minute` - (Required) The number of new users that will be let into the route every minute.
* `custom_page_html` - (Optional) This a templated html file that will be rendered at the edge.
* `queue_all` - (Optional) If queue_all is true all the traffic that is coming to a route will be sent to the waiting room. Default: false.
* `disable_session_renewal` - (Optional) Disables automatic renewal of session cookies. Default: false.
* `suspended` - (Optional) If suspended, the traffic doesn't go to the waiting room. Default: false.
* `description` - (Optional) A description to let users add more details about the waiting room.
* `session_duration` - (Optional) Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the route. Default: 5

## Attributes Reference

The following attributes are exported:

* `id` - The waiting room ID.

## Import

Waiting rooms can be imported using a composite ID formed of zone ID and waiting room ID, e.g.

```
$ terraform import cloudflare_waiting_room.default ae36f999674d196762efcc5abb06b345/d41d8cd98f00b204e9800998ecf8427e
```

where:

* `ae36f999674d196762efcc5abb06b345` - the zone ID
* `d41d8cd98f00b204e9800998ecf8427e` - waiting room ID as returned by [API](https://api.cloudflare.com/#waiting-room-list-waiting-rooms)
