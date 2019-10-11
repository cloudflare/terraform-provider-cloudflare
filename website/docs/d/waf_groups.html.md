---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_groups"
sidebar_current: "docs-cloudflare-datasource-waf-groups"
description: |-
  List available Cloudflare WAF Groups.
---

# cloudflare_waf_groups

Use this data source to look up [WAF Rule Groups][1].

## Example Usage

The example below matches all WAF Rule Groups that contain the word `example` and are currently `on`. The matched WAF Rule Groups are then returned as output.

```hcl
data "cloudflare_waf_groups" "test" {
  filter {
    name   = ".*example.*"
    mode = "on"
  }
}

output "waf_groups" {
  value = data.cloudflare_waf_groups.test.groups
}
```

## Argument Reference

- `zone_id` - (Required) The ID of the DNS zone in which to search for the WAF Rule Groups.
- `package_id` - (Optional) The ID of the WAF Rule Package in which to search for the WAF Rule Groups.
- `filter` - (Optional) One or more values used to look up WAF Rule Groups. If more than one value is given all
values must match in order to be included, see below for full list.

**filter**

- `name` - (Optional) A regular expression matching the name of the WAF Rule Groups to lookup.
- `mode` - (Optional) Mode of the WAF Rule Groups to lookup. Valid values: on and off.

## Attributes Reference

- `groups` - A map of WAF Rule Groups details. Full list below:

**groups**

- `id` - The WAF Rule Group ID
- `name` - The WAF Rule Group name
- `description` - The WAF Rule Group description
- `mode` - The WAF Rule Group mode
- `rules_count` - The number of rules in the WAF Rule Group
- `modified_rules_count` - The number of modified rules in the WAF Rule Group
- `package_id` - The ID of the WAF Rule Package that contains the WAF Rule Group

[1]: https://api.cloudflare.com/#waf-rule-groups-properties
