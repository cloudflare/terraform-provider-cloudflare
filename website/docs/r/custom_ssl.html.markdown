---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_custom_ssl"
sidebar_current: "docs-cloudflare-resource-custom-ssl"
description: !-
  Provides a Cloudflare custom ssl resource.
---

# cloudflare_custom_ssl

Provides a Cloudflare custom ssl resource.

## Example Usage

```hcl
# Add a custom ssl certificate to the domain
resource "cloudflare_custom_ssl" "foossl" {
  zone_id = "${var.cloudflare_zone_id}"
  custom_ssl_options = "${var.cloudflare_custom_ssl_options}"
}

variable "cloudflare_custom_ssl_options" {
  type = "map"
  default = {
    "certificate" = "-----INSERT CERTIFICATE-----"
    "private_key" = "-----INSERT PRIVATE KEY-----"
    "bundle_method" = "ubiquitous",
    "geo_restrictions.label" = "usd",
    "type" = "legacy_custom"
  }
}

variable "cloudflare_zone_id" {
  type = "string"
  default = "1d5fdc9e88c8a8c4518b068cd94331fe"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone id to the custom ssl cert should be added.
* `custom_ssl_options` - (Required) The certificate, private key and associated optional parameters, such as bundle_method, geo_restrictions, and type.
 
## Import

Custom SSL Certs can be imported using a composite ID formed of the zone id and certificate id,
separated by a "/" e.g.

```
$ terraform import cloudflare_custom_ssl.default 1d5fdc9e88c8a8c4518b068cd94331fe/c671356fb0ef68a9d746e3c9ef84ec3e
```
