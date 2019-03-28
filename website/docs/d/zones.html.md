---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zones"
sidebar_current: "docs-cloudflare-datasource-zones"
description: |-
  Get information on a Cloudflare Zones.
---

# cloudflare_zones

Use this data source to look up [Zone][1] records.

## Example Usage

The example below matches all `active` zones that begin with `example.` and are not paused. The matched zones are then
locked down using the `cloudflare_zone_lockdown` resource.

```hcl
data "cloudflare_zones" "test" {
  filter {
    name   = "example.*"
    status = "active"
    paused = false
  }
}

resource "cloudflare_zone_lockdown" "endpoint_lockdown" {
  zone        = "${lookup(data.cloudflare_zones.test.zones[0], "name")}"
  paused      = "false"
  description = "Restrict access to these endpoints to requests from a known IP address"
  urls = [
    "api.mysite.com/some/endpoint*",
  ]
  configurations = [
    {
      "target" = "ip"
      "value" = "198.51.100.4"
    },
  ]
}
```

## Argument Reference
- `filter` - (Required) One or more values used to look up zone records. If more than one value is given all
values must match in order to be included, see below for full list.

**filter**

- `name` - (Optional) A regular expression matching the zone to lookup.
- `status` - (Optional) Status of the zone to lookup. Valid values: active, pending, initializing, moved, deleted, deactivated and read only.
- `paused` - (Optional) Paused status of the zone to lookup. Valid values are `true` or `false`.

## Attributes Reference

- `zones` - A map of zone details. Full list below:

**zones**

- `id` - The zone ID
- `name` - Zone name

[1]: https://api.cloudflare.com/#zone-properties
