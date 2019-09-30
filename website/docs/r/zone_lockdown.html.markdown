---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone_lockdown"
sidebar_current: "docs-cloudflare-resource-zone-lockdown"
description: |-
  Provides a Cloudflare resource to lock down access to URLs by IP address or IP ranges.
---

# cloudflare_zone_lockdown

Provides a Cloudflare Zone Lockdown resource. Zone Lockdown allows you to define one or more URLs (with wildcard matching on the domain or path) that will only permit access if the request originates from an IP address that matches a safelist of one or more IP addresses and/or IP ranges.

## Example Usage

```hcl
# Restrict access to these endpoints to requests from a known IP address.
resource "cloudflare_zone_lockdown" "endpoint_lockdown" {
  zone_id     = "d41d8cd98f00b204e9800998ecf8427e"
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

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to which the access rule should be added.
* `description` - (Optional) A description about the lockdown entry. Typically used as a reminder or explanation for the lockdown.
* `urls` - (Required) A list of simple wildcard patterns to match requests against. The order of the urls is unimportant.
* `configurations` - (Required) A list of IP addresses or IP ranges to match the request against specified in target, value pairs.  It's a complex value. See description below.   The order of the configuration entries is unimportant.
* `paused` - (Optional) Boolean of whether this zone lockdown is currently paused. Default: false.

**Note:** Either `zone` or `zone_id` is required and `zone` will be resolved to `zone_id` upon creation.

The list item in **configurations** block supports:

* `target` - (Required) The request property to target. Allowed values: "ip", "ip_range"
* `value` - (Required) The value to target. Depends on target's type. IP addresses should just be standard IPv4/IPv6 notation i.e. `198.51.100.4` or `2001:db8::/32` and IP ranges in CIDR format i.e. `198.51.0.0/16`.

## Attributes Reference

The following attributes are exported:

* `id` - The access rule ID.

## Import

Records can be imported using a composite ID formed of zone name and record ID, e.g.

```
$ terraform import cloudflare_zone_lockdown  d41d8cd98f00b204e9800998ecf8427e/37cb64fe4a90adb5ca3afc04f2c82a2f
```

where:

* `d41d8cd98f00b204e9800998ecf8427e` - zone ID
* `37cb64fe4a90adb5ca3afc04f2c82a2f` - zone lockdown ID as returned by [API](https://api.cloudflare.com/#zone-lockdown-list-lockdown-rules)
