---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_waf_package"
sidebar_current: "docs-cloudflare-resource-waf-package"
description: |-
  Provides a Cloudflare WAF rule package resource for a particular zone.
---

# cloudflare_waf_package

Provides a Cloudflare WAF rule package resource for a particular zone. This can be used to configure firewall behaviour for pre-defined firewall packages.

## Example Usage

```hcl
resource "cloudflare_waf_package" "owasp" {
  package_id = "a25a9a7e9c00afc1fb2e0245519d725b"
  zone_id = "ae36f999674d196762efcc5abb06b345"
  sensitivity = "medium"
  action_mode = "simulate"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to apply to.
* `package_id` - (Required) The WAF Package ID.
* `sensitivity` - (Required) The sensitivity of the package, can be one of ["high", "medium", "low", "off"].
* `action_mode` - (Required) The action mode of the package, can be one of ["block", "challenge", "simulate"].


## Attributes Reference

The following attributes are exported:

* `id` - The WAF Package ID, the same as package_id.

## Import

Packages can be imported using a composite ID formed of zone ID and the WAF Package ID, e.g.

```
$ terraform import cloudflare_waf_package.owasp ae36f999674d196762efcc5abb06b345/a25a9a7e9c00afc1fb2e0245519d725b
```
