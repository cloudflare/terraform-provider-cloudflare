---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_authenticated_origin_pulls_certificate"
sidebar_current: "docs-cloudflare-resource-authenticated-origin-pulls-certificate"
description: |-
  Provides a Cloudflare Authenticated Origin Pulls certificate resource.
---

# cloudflare_authenticated_origin_pulls_certificate

Provides a Cloudflare Authenticated Origin Pulls certificate resource. An uploaded client certificate is required to use Per-Zone or Per-Hostname Authenticated Origin Pulls.

## Example Usage

```hcl
# Per-Zone Authenticated Origin Pulls certificate
resource "cloudflare_authenticated_origin_pulls_certificate" "my_per_zone_aop_cert" {
  zone_id     = "${var.cloudflare_zone_id}"
  certificate = "-----INSERT CERTIFICATE-----"
  private_key = "-----INSERT PRIVATE KEY-----"
  type        = "per-zone"
}

# Per-Hostname Authenticated Origin Pulls certificate
resource "cloudflare_authenticated_origin_pulls_certificate" "my_per_hostname_aop_cert" {
  zone_id     = "${var.cloudflare_zone_id}"
  certificate = "-----INSERT CERTIFICATE-----"
  private_key = "-----INSERT PRIVATE KEY-----"
  type        = "per-hostname"
}
```

## Argument Reference

The following arguments are supported:

- `zone_id` - (Required) The zone ID to upload the certificate to.
- `certificate` - (Required) The public client certificate.
- `private_key` - (Required) The private key of the client certificate.
- `type` - (Required) The form of Authenticated Origin Pulls to upload the certificate to.

## Import
