---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_group"
sidebar_current: "docs-cloudflare-resource-waf-group"
description: |-
  Provides a Cloudflare WAF rule group resource for a particular zone.
---

# cloudflare_waf_group

Provides a Cloudflare WAF rule group resource for a particular zone. This can be used to configure firewall behaviour for pre-defined firewall groups.

## Example Usage

```hcl
resource "cloudflare_waf_group" "honey_pot" {
  group_id = "de677e5818985db1285d0e80225f06e5"
  zone_id = "ae36f999674d196762efcc5abb06b345"
  mode = "on"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to apply to.
* `group_id` - (Required) The WAF Rule Group ID.
* `package_id` - (Optional) The ID of the WAF Rule Package that contains the group.
* `mode` - (Optional) The mode of the group, can be one of ["on", "off"].


## Attributes Reference

The following attributes are exported:

* `id` - The WAF Rule Group ID, the same as group_id.
* `package_id` - The ID of the WAF Rule Package that contains the group.

## Import

Rules can be imported using a composite ID formed of zone ID and the WAF Rule Group ID, e.g.

```
$ terraform import cloudflare_waf_group.honey_pot ae36f999674d196762efcc5abb06b345/de677e5818985db1285d0e80225f06e5
```
