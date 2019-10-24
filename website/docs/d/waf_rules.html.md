---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_rules"
sidebar_current: "docs-cloudflare-datasource-waf-rules"
description: |-
  List available Cloudflare WAF Rules.
---

# cloudflare_waf_rules

Use this data source to look up [WAF Rules][1].

## Example Usage

The example below matches all WAF Rules that are in the group of ID `de677e5818985db1285d0e80225f06e5`, contain `example` in their description, and are currently `on`. The matched WAF Rules are then returned as output.

```hcl
data "cloudflare_waf_rules" "test" {
  zone_id    = "ae36f999674d196762efcc5abb06b345"
  package_id = "a25a9a7e9c00afc1fb2e0245519d725b"

  filter {
    description = ".*example.*"
    mode        = "on"
    group_id    = "de677e5818985db1285d0e80225f06e5"
  }
}

output "waf_rules" {
  value = data.cloudflare_waf_rules.test.rules
}
```

## Argument Reference

- `zone_id` - (Required) The ID of the DNS zone in which to search for the WAF Rules.
- `package_id` - (Optional) The ID of the WAF Rule Package in which to search for the WAF Rules.
- `filter` - (Optional) One or more values used to look up WAF Rules. If more than one value is given all
values must match in order to be included, see below for full list.

**filter**

- `description` - (Optional) A regular expression matching the description of the WAF Rules to lookup.
- `mode` - (Optional) Mode of the WAF Rules to lookup. Valid values: on and off.
- `group_id` - (Optional) The ID of the WAF Rule Group in which the WAF Rules to lookup have to be.

## Attributes Reference

- `rules` - A map of WAF Rules details. Full list below:

**rules**

- `id` - The WAF Rule ID
- `description` - The WAF Rule description
- `priority` - The WAF Rule priority
- `mode` - The WAF Rule mode
- `group_id` - The ID of the WAF Rule Group that contains the WAF Rule
- `group_name` - The Name of the WAF Rule Group that contains the WAF Rule
- `package_id` - The ID of the WAF Rule Package that contains the WAF Rule

[1]: https://api.cloudflare.com/#waf-rule-groups-properties
