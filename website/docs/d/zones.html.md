---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zones"
sidebar_current: "docs-cloudflare-datasource-zones"
description: |-
  Get information on a Cloudflare Zones.
---

# cloudflare_zones

Use this data source to look up [Zone][1] records.

## Example usage

Given you have the following zones in Cloudflare.

- example.com
- example.net
- not-example.com

```hcl
# Look for a single domain that you know exists using an exact match.
# API request will be for zones?name=example.com. Will not match not-example.com
# or example.net.
data "cloudflare_zones" "example" {
  filter {
    name = "example.com"
  }
}
```

```hcl
# Look for all zones which include "example".
# API request will be for zones?name=contains:example. Will return all three
# zones.
data "cloudflare_zones" "example" {
  filter {
    name        = "example"
    lookup_type = "contains"
  }
}
```

```hcl
# Look for all zones which include "example" but start with "not-".
# API request will be for zones?name=contains:example. Will perform client side
# filtering using the provided regex and will only match one domain.
data "cloudflare_zones" "example" {
  filter {
    name        = "example"
    lookup_type = "contains"
    match       = "^not-"
  }
}
```

### Example usage with other resources

The example below matches all zones which have "example" in their value, end
with ".com" and are active. The matched zone is then referenced in the zone
lockdown resource.

```hcl
data "cloudflare_zones" "test" {
  filter {
    name        = "example"
    lookup_type = "contains"
    match       = ".com$"
    status      = "active"
  }
}

resource "cloudflare_zone_lockdown" "endpoint_lockdown" {
  zone        = "${lookup(data.cloudflare_zones.test.zones[0], "name")}"
  paused      = "false"
  description = "Restrict access to these endpoints to requests from a known IP address"
  urls = [
    "api.mysite.com/some/endpoint*",
  ]
  configurations {
    target = "ip"
    value = "198.51.100.4"
  }
}
```

## Argument Reference
- `filter` - (Required) One or more values used to look up zone records. If more than one value is given all
values must match in order to be included, see below for full list.

**filter**

- `name` - (Optional) A string value to search for.
- `lookup_type` - (Optional) The type of search to perform for the `name` value
  when querying the zone API. Valid values: `"exact"` and `"contains"`. Defaults
  to `"exact"`.
- `match` - (Optional) A RE2 compatible regular expression to filter the
  results. This is performed client side whereas the `name` and `lookup_type`
  are performed on the Cloudflare server side.
- `status` - (Optional) Status of the zone to lookup. Valid values: `"active"`,
  `"pending"`, `"initializing"`, `"moved"`, `"deleted"`, `"deactivated"` and
  `"read only"`.
- `paused` - (Optional) Paused status of the zone to lookup. Valid values are
  `true` or `false`.

## Attributes Reference

- `zones` - A map of zone details. Full list below:

**zones**

- `id` - The zone ID
- `name` - Zone name

[1]: https://api.cloudflare.com/#zone-properties
