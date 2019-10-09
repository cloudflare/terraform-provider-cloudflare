---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_packages"
sidebar_current: "docs-cloudflare-datasource-waf-packages"
description: |-
  List available Cloudflare WAF Packages.
---

# cloudflare_waf_packages

Use this data source to look up [WAF Rule Packages][1].

## Example Usage

The example below matches all `high` sensitivity WAF Rule Packages, with a `challenge` action mode and an `anomaly` detection mode, that contain the word `example`. The matched WAF Rule Packages are then returned as output.

```hcl
data "cloudflare_waf_packages" "test" {
  filter {
    name   = ".*example.*"
    detection_mode = "anomaly"
    sensitivity = "high"
    action_mode = "challenge"
  }
}

output "waf_packages" {
  value = data.cloudflare_waf_packages.test.packages
}
```

## Argument Reference

- `filter` - (Optional) One or more values used to look up WAF Rule Packages. If more than one value is given all
values must match in order to be included, see below for full list.

**filter**

- `name` - (Optional) A regular expression matching the name of the WAF Rule Packages to lookup.
- `detection_mode` - (Optional) Detection mode of the WAF Rule Packages to lookup.
- `sensitivity` - (Optional) Sensitivity of the WAF Rule Packages to lookup. Valid values: high, medium, low and off.
- `action_mode` - (Optional) Action mode of the WAF Rule Packages to lookup. Valid values: simulate, block and challenge.

## Attributes Reference

- `packages` - A map of WAF Rule Packages details. Full list below:

**packages**

- `id` - The WAF Rule Package ID
- `name` - The WAF Rule Package name
- `description` - The WAF Rule Package description
- `detection_mode` - The WAF Rule Package detection mode
- `sensitivity` - The WAF Rule Package sensitivity
- `action_mode` - The WAF Rule Package action mode

[1]: https://api.cloudflare.com/#waf-rule-packages-properties
