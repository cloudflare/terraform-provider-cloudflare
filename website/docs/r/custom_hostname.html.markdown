---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_custom_hostname"
sidebar_current: "docs-cloudflare-resource-custom-hostname"
description: !-
  Provides a Cloudflare custom hostname resource.
---

# cloudflare_custom_hostname

Provides a Cloudflare custom hostname (also known as SSL for SaaS) resource.

## Example Usage

```hcl
resource "cloudflare_custom_hostname" "example_hostname" {
  zone_id  = "d41d8cd98f00b204e9800998ecf8427e"
  hostname = "hostname.example.com"
  ssl {
    method = "txt"
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID where the custom hostname should be assigned.
* `hostname` - (Required) Hostname you intend to request a certificate for.
* `custom_origin_server` - (Optional) The custom origin server used for certificates.
* `custom_origin_sni` - (Optional) The [custom origin SNI](https://developers.cloudflare.com/ssl/ssl-for-saas/hostname-specific-behavior/custom-origin) used for certificates.
* `ssl` - (Required) SSL configuration of the certificate. See further notes below.

**ssl** block supports:

* `method` - (Required) Domain control validation (DCV) method used for this
  hostname. Valid values are `"txt"`, `"http"` and `"email"`.
* `type` - (Required) Level of validation to be used for this hostname. Domain validation ("dv") must be used.
* `wildcard` - (Required) Indicates whether the certificate covers a wildcard.
* `custom_certificate` - (Optional) If a custom uploaded certificate is used.
* `custom_key` - (Optional) The key for a custom uploaded certificate.
* `settings` - (Required) SSL/TLS settings for the certificate. See further notes below.

**settings** block supports:

* `http2` - (Optional) Whether or not HTTP2 should be supported. Valid values are `"on"` or `"off"`.
* `tls13` - (Optional) Whether or not TLSv1.3 should be supported. Valid values are `"on"` or `"off"`.
* `min_tls_version` - (Optional) Lowest version of TLS this certificate should
  support. Valid values are `"1.0"`, `"1.1"`, `"1.2"` and `"1.3"`.
* `ciphers` - (Optional) List of SSL/TLS ciphers to associate with this certificate.
* `early_hints` - (Optional) Whether or not early hints should be supported. Valid values are `"on"` or `"off"`.

## Attributes Reference

The following attributes are exported:

* `ownership_verification.type` - Domain control validation (DCV) method used
  for the hostname.
* `ownership_verification.value` - Domain control validation (DCV) value for
  confirming ownership. Example, "_cf-custom-hostname.example.com`
* `ownership_verification.name` - Domain control validation (DCV) name
  confirming ownership. Example, "03f28e11-fa64-4966-bb1e-dd2423e16f36"`
* `ownership_verification_http.http_url` - Domain control validation (DCV) URL for
  confirming ownership. Example, `http://hostname.example.com/.well-known/cf-custom-hostname-challenge/643395f9-de80-42f5-a2a0-e03ff60cf2a7`
* `ownership_verification_http.http_body` - Domain control validation (DCV) body for
  confirming ownership. Example, `03f28e11-fa64-4966-bb1e-dd2423e16f36`

## Import

Custom hostname certificates can be imported using a composite ID formed of the zone ID and [hostname ID](https://api.cloudflare.com/#custom-hostname-for-a-zone-properties),
separated by a "/" e.g.

```
$ terraform import cloudflare_custom_hostname.example d41d8cd98f00b204e9800998ecf8427e/0d89c70d-ad9f-4843-b99f-6cc0252067e9
```
