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

Authenticated Origin Pull certificates can be imported using a composite ID formed of the zone ID, the form of Authenticated Origin Pulls, and the certificate ID, e.g.

```
# Import Per-Zone Authenticated Origin Pull certificate
$ terraform import cloudflare_authenticated_origin_pulls_certificate.2458ce5a-0c35-4c7f-82c7-8e9487d3ff60 023e105f4ecef8ad9ca31a8372d0c353/per-zone/2458ce5a-0c35-4c7f-82c7-8e9487d3ff60

# Import Per-Hostname Authenticated Origin Pull certificate
$ terraform import cloudflare_authenticated_origin_pulls_certificate.2458ce5a-0c35-4c7f-82c7-8e9487d3ff60 023e105f4ecef8ad9ca31a8372d0c353/per-hostname/2458ce5a-0c35-4c7f-82c7-8e9487d3ff60
```
