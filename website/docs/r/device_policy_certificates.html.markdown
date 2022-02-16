---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_device_policy_certificates"
sidebar_current: "docs-cloudflare-resource-device-policy-certificates"
description: |-
  Provides a Cloudflare Device Policy Certificates resource.
---

# cloudflare_device_policy_certificates

Provides a Cloudflare device policy certificates resource. Device policy certificate resources enable client device certificate generation.

## Example Usage

```hcl
resource "cloudflare_device_policy_certificates" "client_certificates" {
  zone_id     = "1d5fdc9e88c8a8c4518b068cd94331fe"
  enabled     = true
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The zone where certificate generation is allowed.
* `enabled` - (Required) True if certificate generation is enabled.
## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the device policy certificates setting.

## Import

Device policy certificate settings can be imported using the zone ID.

```
$ terraform import cloudflare_device_policy_certificates.client_certificates cb029e245cfdd66dc8d2e570d5dd3322
```
