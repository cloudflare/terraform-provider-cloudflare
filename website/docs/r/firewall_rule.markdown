---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_firewall_rule"
sidebar_current: "docs-cloudflare-resource-firewall-rule"
description: |-
  Define Firewall rule using filter expression for more control over how traffic is matched to the rule.
---

# cloudflare_firewall_rule

Define Firewall rules using filter expressions for more control over how traffic is matched to the rule.
A filter expression permits selecting traffic by multiple criteria allowing greater freedom in rule creation.

Filter expressions needs to be created first before using Firewall Rule. See [Filter](filter.html).

## Example Usage

```hcl
resource "cloudflare_filter" "wordpress" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  description = "Wordpress break-in attempts that are outside of the office"
  expression = "(http.request.uri.path ~ \".*wp-login.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1"
}

resource "cloudflare_firewall_rule" "wordpress" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  description = "Block wordpress break-in attempts"
  filter_id = cloudflare_filter.wordpress.id
  action = "block"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone to which the Filter should be added.
* `action` - (Required) The action to apply to a matched request. Allowed values: "block", "challenge", "allow", "js_challenge". Enterprise plan also allows "log".
* `priority` - (Optional) The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.
* `paused` - (Optional) Whether this filter based firewall rule is currently paused. Boolean value.
* `description` - (Optional) A description of the rule to help identify it.

## Attributes Reference

The following attributes are exported:

* `id` - Firewall Rule identifier.

## Import

Firewall Rule can be imported using a composite ID formed of zone ID and rule ID, e.g.

```
$ terraform import cloudflare_filter.default d41d8cd98f00b204e9800998ecf8427e/9e107d9d372bb6826bd81d3542a419d6
```

where:

* `d41d8cd98f00b204e9800998ecf8427e` - zone ID
* `9e107d9d372bb6826bd81d3542a419d6` - rule ID as returned by [API](https://api.cloudflare.com/#zone-firewall-filter-rules)
