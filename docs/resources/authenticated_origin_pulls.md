---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_authenticated_origin_pulls"
description: Provides a Cloudflare Authenticated Origin Pulls resource.
---

# cloudflare_authenticated_origin_pulls

Provides a Cloudflare Authenticated Origin Pulls resource. An `cloudflare_authenticated_origin_pulls` resource is required to use Per-Zone or Per-Hostname Authenticated Origin Pulls.

## Example Usage

The arguments that you provide determine which form of Authenticated Origin Pulls to use:

```hcl
# Authenticated Origin Pulls
resource "cloudflare_authenticated_origin_pulls" "my_aop" {
  zone_id     = "${var.cloudflare_zone_id}"
  enabled     = true
}

# Per-Zone Authenticated Origin Pulls
resource "cloudflare_authenticated_origin_pulls_certificate" "my_per_zone_aop_cert" {
  zone_id     = "${var.cloudflare_zone_id}"
  certificate = "-----INSERT CERTIFICATE-----"
  private_key = "-----INSERT PRIVATE KEY-----"
  type        = "per-zone"
}

resource "cloudflare_authenticated_origin_pulls" "my_per_zone_aop" {
  zone_id                                   = "${var.cloudflare_zone_id}"
  authenticated_origin_pulls_certificate    = cloudflare_authenticated_origin_pulls_certificate.my_per_zone_aop_cert.id
  enabled                                   = true
}

# Per-Hostname Authenticated Origin Pulls
resource "cloudflare_authenticated_origin_pulls_certificate" "my_per_hostname_aop_cert" {
  zone_id     = "${var.cloudflare_zone_id}"
  certificate = "-----INSERT CERTIFICATE-----"
  private_key = "-----INSERT PRIVATE KEY-----"
  type        = "per-hostname"
}

resource "cloudflare_authenticated_origin_pulls" "my_per_hostname_aop" {
  zone_id                                   = "${var.cloudflare_zone_id}"
  authenticated_origin_pulls_certificate    = cloudflare_authenticated_origin_pulls_certificate.my_per_hostname_aop_cert.id
  hostname                                  = "aop.example.com"
  enabled                                   = true
}
```

## Argument Reference

The following arguments are supported:

- `zone_id` - (Required) The zone ID to upload the certificate to.
- `authenticated_origin_pulls_certificate` - (Optional) The id of an uploaded Authenticated Origin Pulls certificate. If no hostname is provided, this certificate will be used zone wide as Per-Zone Authenticated Origin Pulls.
- `hostname` - (Optional) Specify a hostname to enable Per-Hostname Authenticated Origin Pulls on, using the provided certificate.
- `enabled` - (Required) Whether or not to enable Authenticated Origin Pulls on the given zone or hostname.

## Import

Authenticated Origin Pull configuration can be imported using a composite ID formed of the zone ID, the form of Authenticated Origin Pulls, and the certificate ID, with each section filled or left blank e.g.

```
# Import Authenticated Origin Pull configuration
$ terraform import cloudflare_authenticated_origin_pulls_certificate.my_aop 023e105f4ecef8ad9ca31a8372d0c353//

# Import Per-Zone Authenticated Origin Pull configuration
$ terraform import cloudflare_authenticated_origin_pulls_certificate.my_per_zone_aop 023e105f4ecef8ad9ca31a8372d0c353/2458ce5a-0c35-4c7f-82c7-8e9487d3ff60/

# Import Per-Hostname Authenticated Origin Pull configuration
$ terraform import cloudflare_authenticated_origin_pulls_certificate.my_per_hostname_aop 023e105f4ecef8ad9ca31a8372d0c353/2458ce5a-0c35-4c7f-82c7-8e9487d3ff60/aop.example.com
```
