---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_rule"
description: Provides a Cloudflare WAF rule resource for a particular zone.
---

# cloudflare_waf_rule

Provides a Cloudflare WAF rule resource for a particular zone. This can be used to configure firewall behaviour for pre-defined firewall rules.

## Example Usage

```hcl
resource "cloudflare_waf_rule" "rule_100000" {
  rule_id = "100000"
  zone_id = "ae36f999674d196762efcc5abb06b345"
  mode = "simulate"
}
```

## Argument Reference

The following arguments are supported:

- `zone_id` - (Required) The DNS zone ID to apply to.
- `rule_id` - (Required) The WAF Rule ID.
- `package_id` - (Optional) The ID of the WAF Rule Package that contains the rule.
- `mode` - (Required) The mode of the rule, can be one of ["block", "challenge", "default", "disable", "simulate"] or ["on", "off"] depending on the WAF Rule type.

## Attributes Reference

The following attributes are exported:

- `id` - The WAF Rule ID, the same as rule_id.
- `package_id` - The ID of the WAF Rule Package that contains the rule.
- `group_id` - The ID of the WAF Rule Group that contains the rule.

## Import

Rules can be imported using a composite ID formed of zone ID and the WAF Rule ID, e.g.

```
$ terraform import cloudflare_waf_rule.100000 ae36f999674d196762efcc5abb06b345/100000
```
