---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_filter"
sidebar_current: "docs-cloudflare-resource-filter"
description: |-
  Provides a Cloudflare Filter expression that can be referenced across multiple features.
---

# cloudflare_filter

Filter expressions that can be referenced across multiple features, e.g. [Firewall Rule](firewall_rule.html). The expression format is similar to [Wireshark Display Filter](https://www.wireshark.org/docs/man-pages/wireshark-filter.html).

## Example Usage

```hcl
resource "cloudflare_filter" "wordpress" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  description = "Wordpress break-in attempts that are outside of the office"
  expression = "(http.request.uri.path ~ \"*wp-login.php\" or http.request.uri.path ~ \"*xmlrpc.php\") and ip.addr ne 192.0.2.1"
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Optional) The DNS zone to which the Filter should be added. Will be resolved to `zone_id` upon creation.
* `zone_id` - (Optional) The DNS zone to which the Filter should be added.
* `paused` - (Optional) Whether this filter is currently paused. Boolean value.
* `expression` - (Required) The filter expression to be used.
* `description` - (Optional) A note that you can use to describe the purpose of the filter.
* `ref` - (Optional) Short reference tag to quickly select related rules.

## Attributes Reference

The following attributes are exported:

* `id` - Filter identifier.
* `zone_id` - The DNS zone ID.

## Import

Filter can be imported using a composite ID formed of zone ID and filter ID, e.g.

```
$ terraform import cloudflare_filter.default d41d8cd98f00b204e9800998ecf8427e/9e107d9d372bb6826bd81d3542a419d6
```

where:

* `d41d8cd98f00b204e9800998ecf8427e` - zone ID
* `9e107d9d372bb6826bd81d3542a419d6` - filter ID as returned by [API](https://api.cloudflare.com/#zone-firewall-filters)
