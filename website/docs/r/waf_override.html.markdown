---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_override"
sidebar_current: "docs-cloudflare-resource-waf-override"
description: |-
  Provides a Cloudflare WAF Override resource.
---

# cloudflare_waf_override

Provides a Cloudflare WAF override resource. This enables the ability to toggle
WAF rules and groups on or off based on URIs.

## Example Usage

```hcl
resource "cloudflare_waf_override" "shop_ecxample" {
  zone_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  urls    = [
    "example.com/no-waf-here",
    "example.com/another/path/*"
  ]

  # Disable rule ID 100015.
  rules = {
    "100015": "disable"
  }

  # Set to Cloudflare default action for group ID ea8687e59929c1fd05ba97574ad43f77.
  groups = {
    "ea8687e59929c1fd05ba97574ad43f77": "default"
  }

  # Update the actions for when a matching rule is encountered.
  rewrite_action = {
    "default": "block",
    "challenge": "block",
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone to which the WAF override condition should be added.
* `urls` - (Required) An array of URLs to apply the WAF override to.
* `rules` - (Required) A list of WAF rule ID to rule action you intend to apply.
* `paused` - (Optional) Whether this package is currently paused.
* `description` - (Optional) Description of what the WAF override does.
* `priority` - (Optional) Relative priority of this configuration when multiple configurations match a single URL.
* `groups` - (Optional) Similar to `rules`; which WAF groups you want to alter.
* `rewrite_action` - (Optional) When a WAF rule matches, substitute its configured action for a different action specified by this definition.

## Import

WAF Overrides can be imported using a composite ID formed of zone
ID and override ID.

```
$ terraform import cloudflare_waf_override.my_example_waf_override 3abe5b950053dbddf1516d89f9ef1e8a/9d4e66d7649c178663bf62e06dbacb23
```
