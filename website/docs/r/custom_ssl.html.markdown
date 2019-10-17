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
  custom_ssl_options = {
    "certificate" = "-----INSERT CERTIFICATE-----"
    "private_key" = "-----INSERT PRIVATE KEY-----"
    "bundle_method" = "ubiquitous",
    "geo_restrictions" = "us",
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

**custom_ssl_options** block supports:

* `certificate` - (Required) Certificate certificate and the intermediate(s)
* `private_key` - (Required) Certificate's private key
* `bundle_method` - (Optional) Method of building intermediate certificate chain. A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. Valid values are `ubiquitous` (default), `optimal`, `force`.
* `geo_restrictions` - (Optional) Specifies the region where your private key can be held locally. Valid values are `us`, `eu`, `highest_security`.
* `type` - (Optional) Whether to enable support for legacy clients which do not include SNI in the TLS handshake. Valid values are `legacy_custom` (default), `sni_custom`.

## Import

Custom SSL Certs can be imported using a composite ID formed of the zone ID and [certificate ID](https://api.cloudflare.com/#custom-ssl-for-a-zone-properties),
separated by a "/" e.g.

```
$ terraform import cloudflare_custom_ssl.default 1d5fdc9e88c8a8c4518b068cd94331fe/0123f0ab-9cde-45b2-80bd-4da3010f1337
```
