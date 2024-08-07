---
page_title: "cloudflare_zero_trust_device_certificates Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a Cloudflare device policy certificates resource. Device
  policy certificate resources enable client device certificate
  generation.
---

# cloudflare_zero_trust_device_certificates (Resource)

Provides a Cloudflare device policy certificates resource. Device
policy certificate resources enable client device certificate
generation.

## Example Usage

```terraform
resource "cloudflare_zero_trust_device_certificates" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  enabled = true
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) `true` if certificate generation is enabled.
- `zone_id` (String) The zone identifier to target for the resource.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zero_trust_device_certificates.example <zone_id>
```
