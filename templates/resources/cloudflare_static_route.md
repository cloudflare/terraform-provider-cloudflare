---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_static_route"
description: Provides a resource which manages Cloudflare static routes for Magic Transit or Magic WAN.
---

# cloudflare_static_route

Provides a resource, that manages Cloudflare static routes for Magic Transit or Magic WAN.
Static routes are used to route traffic through GRE tunnels.

## Example Usage

```hcl
resource "cloudflare_static_route" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  description = "New route for new prefix 192.0.2.0/24"
  prefix = "192.0.2.0/24"
  nexthop = "10.0.0.0"
  priority = 100
  weight = 10
  colo_names = [
    "den01"
  ]
  colo_regions = [
    "APAC"
  ]
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account where the static route is being created.
- `description` - (Optional) Description of the static route.
- `prefix` - (Required) Your network prefix using CIDR notation.
- `nexthop` - (Required) The nexthop IP address where traffic will be routed to.
- `priority` - (Required) The priority for the static route.
- `weight` - (Optional) The optional weight for ECMP routes.
- `colo_names` - (Optional) Optional list of Cloudflare colocation names for this static route.
- `colo_regions` - (Optional) Optional list of Cloudflare colocation regions for this static route.

## Import

An existing static route can be imported using the account ID and static route ID

```
$ terraform import cloudflare_static_route.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322
```
